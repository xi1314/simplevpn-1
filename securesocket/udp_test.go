package securesocket

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestSecureUdp(t *testing.T) {
	data := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	port := ":0"
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		t.Fatal(err)
	}
	ln, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		t.Fatal(err)
	}
	Addr := ln.LocalAddr()
	defer ln.Close()
	cipher, err := NewCipher("aes-256-cfb", "password")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		secureConn := NewPacketConn(ln, cipher.Copy())
		buf := make([]byte, 1500)
		n, src, err := secureConn.ReadFrom(buf)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(buf[:n], data) {
			t.Error("UDP server recieved data does not match.")
		}
		_, err = secureConn.WriteTo(buf[:n], src)
		if err != nil {
			t.Fatal(err)
		}
		return
	}()
	time.Sleep(100 * time.Millisecond)
	ln2, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		t.Fatal(err)
	}
	secureconn := NewPacketConn(ln2, cipher.Copy())
	_, err = secureconn.WriteTo(data, Addr)
	if err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 1500)
	n, _, err := secureconn.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf[:n], data) {
		t.Error("UDP client read does not match.")
	}

}
