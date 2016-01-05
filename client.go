package main

import (
	"fmt"
	ss "github.com/lixin9311/simplevpn/securesocket"
	"github.com/lixin9311/simplevpn/tap"
	"io"
	"log"
	"net"
)

func runClient(conf *Config) error {
	test := net.ParseIP(conf.Server.Ip)
	v6 := false
	if test.To4() == nil {
		v6 = true
	}
	cipher, err := ss.NewCipher(conf.User.Method, conf.User.Password)
	if err != nil {
		return err
	}
	ifce, err := tap.NewTAP()
	if err != nil {
		log.Println("Failed to create TAP device:", err)
		return err
	}
	defer ifce.Close()
	if !v6 {
		ip, ip_mask, err := net.ParseCIDR(conf.Server.Ip + "/32")
		if err != nil {
			log.Println("Failed to parse ip:", err)
			return err
		}
		ip_mask.IP = ip
		err = tap.Bypass(ip_mask)
		if err != nil {
			log.Println("[Client]: Failed to bypass server address from route, please manually fix that:", err)
		}
		defer tap.Unbypass()
	}
	// reg with server
	var dst string
	if v6 {
		dst = fmt.Sprintf("[%s]:%d", conf.Server.Ip, conf.Server.Port)
	} else {
		dst = fmt.Sprintf("%s:%d", conf.Server.Ip, conf.Server.Port)
	}
	conn, err := net.Dial("tcp", dst)
	if err != nil {
		return err
	}
	secureconn := ss.NewConn(conn, cipher.Copy())
	c := ss.NewPacketStreamConn(secureconn)
	defer c.Close()
	auth := new(Auth)
	auth.Type = Auth_Hello
	mac := ifce.MacAddr()
	auth.MacAddr = mac[:]
	data, err := auth.Marshal()
	if err != nil {
		log.Println("[Client]: Failed to marshal data:", err)
		return err
	}
	_, err = c.Write(data)
	if err != nil {
		log.Println("[Client]: Failed to write socket: ", err)
		return err
	}
	buf := make([]byte, 2048)
	n, err := c.Read(buf)
	if err != nil {
		log.Println("[Client]: Failed to recieve config: ", err)
		return err
	}
	err = auth.Unmarshal(buf[:n])
	if err != nil {
		log.Println("[Client]: Failed to decode recieved config: ", err)
		return err
	}
	if auth.Type != Auth_Welcome {
		return fmt.Errorf("[Client]: Unexpected response type: %s.", Auth_MessageType_name[int32(auth.Type)])
	}
	ip, ip_mask, err := net.ParseCIDR(auth.IP)
	if err != nil {
		log.Println("[Client]: Failed to parse CIDR from response:", err)
		return err
	}
	ip_mask.IP = ip
	err = ifce.SetIP(ip_mask)
	if err != nil {
		log.Println("Failed to set IP address:", err)
		return err
	}
	ip, ip_mask, err = net.ParseCIDR("0.0.0.0/1")
	if err != nil {
		log.Println("Failed to parse address:", err)
	}
	ip_mask.IP = ip
	ip = net.ParseIP(auth.GateWay)
	err = ifce.AddRoute(ip, ip_mask)
	if err != nil {
		log.Println("Failed to set default route, please manually fix that:", err)
	}
	ip, ip_mask, err = net.ParseCIDR("128.0.0.0/1")
	if err != nil {
		log.Println("Failed to parse address:", err)
	}
	ip_mask.IP = ip
	ip = net.ParseIP(auth.GateWay)
	err = ifce.AddRoute(ip, ip_mask)
	if err != nil {
		log.Println("Failed to set default route, please manually fix that", err)
	}
	defer func() {
		ip, ip_mask, _ = net.ParseCIDR("0.0.0.0/1")
		ip_mask.IP = ip
		ip = net.ParseIP(auth.GateWay)
		ifce.DelRoute(ip, ip_mask)
		ip, ip_mask, _ = net.ParseCIDR("128.0.0.0/1")
		ip_mask.IP = ip
		ip = net.ParseIP(auth.GateWay)
		ifce.DelRoute(ip, ip_mask)
	}()
	go PipeThenClose(ifce, c)
	PipeThenClose(c, ifce)
	return nil
}

func PipeThenClose(src, dst io.ReadWriteCloser) {
	defer dst.Close()
	buf := make([]byte, MaxPacketSize)
	for {
		n, err := src.Read(buf)
		// read may return EOF with n > 0
		// should always process n > 0 bytes before handling error
		if n > 0 {
			// Note: avoid overwrite err returned by Read.
			if _, err := dst.Write(buf[0:n]); err != nil {
				log.Println("write:", err)
				break
			}
		}
		if err != nil {
			// Always "use of closed network connection", but no easy way to
			// identify this specific error. So just leave the error along for now.
			// More info here: https://code.google.com/p/go/issues/detail?id=4373
			/*
				if bool(Debug) && err != io.EOF {
					Debug.Println("read:", err)
				}
			*/
			log.Println("read:", err)
			break
		}
	}
}
