// Code generated by protoc-gen-gogo.
// source: simplevpn.proto
// DO NOT EDIT!

/*
	Package simplevpn is a generated protocol buffer package.

	It is generated from these files:
		simplevpn.proto

	It has these top-level messages:
		User
		Auth
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User_UserType int32

const (
	User_Admin User_UserType = 0
	User_User  User_UserType = 1
)

var User_UserType_name = map[int32]string{
	0: "Admin",
	1: "User",
}
var User_UserType_value = map[string]int32{
	"Admin": 0,
	"User":  1,
}

func (x User_UserType) String() string {
	return proto.EnumName(User_UserType_name, int32(x))
}

type Auth_MessageType int32

const (
	Auth_Hello   Auth_MessageType = 0
	Auth_Welcome Auth_MessageType = 1
)

var Auth_MessageType_name = map[int32]string{
	0: "Hello",
	1: "Welcome",
}
var Auth_MessageType_value = map[string]int32{
	"Hello":   0,
	"Welcome": 1,
}

func (x Auth_MessageType) String() string {
	return proto.EnumName(Auth_MessageType_name, int32(x))
}

type User struct {
	Type      User_UserType `protobuf:"varint,1,opt,name=Type,proto3,enum=simplevpn.User_UserType" json:"Type,omitempty"`
	Name      string        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email     string        `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Password  string        `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	PublicKey []byte        `protobuf:"bytes,5,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Extension []byte        `protobuf:"bytes,6,opt,name=Extension,proto3" json:"Extension,omitempty"`
	// temp delete those
	PubBase64 string `protobuf:"bytes,7,opt,name=PubBase64,proto3" json:"PubBase64,omitempty"`
	PriBase64 string `protobuf:"bytes,8,opt,name=PriBase64,proto3" json:"PriBase64,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

type Auth struct {
	Type      Auth_MessageType `protobuf:"varint,1,opt,name=Type,proto3,enum=simplevpn.Auth_MessageType" json:"Type,omitempty"`
	MacAddr   []byte           `protobuf:"bytes,2,opt,name=MacAddr,proto3" json:"MacAddr,omitempty"`
	IP        string           `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`
	GateWay   string           `protobuf:"bytes,4,opt,name=GateWay,proto3" json:"GateWay,omitempty"`
	DNS       string           `protobuf:"bytes,5,opt,name=DNS,proto3" json:"DNS,omitempty"`
	MTU       int32            `protobuf:"varint,6,opt,name=MTU,proto3" json:"MTU,omitempty"`
	Extension []byte           `protobuf:"bytes,7,opt,name=Extension,proto3" json:"Extension,omitempty"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}

func init() {
	proto.RegisterType((*User)(nil), "simplevpn.User")
	proto.RegisterType((*Auth)(nil), "simplevpn.Auth")
	proto.RegisterEnum("simplevpn.User_UserType", User_UserType_name, User_UserType_value)
	proto.RegisterEnum("simplevpn.Auth_MessageType", Auth_MessageType_name, Auth_MessageType_value)
}
func (m *User) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *User) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintSimplevpn(data, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.Email) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.Email)))
		i += copy(data[i:], m.Email)
	}
	if len(m.Password) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.Password)))
		i += copy(data[i:], m.Password)
	}
	if m.PublicKey != nil {
		if len(m.PublicKey) > 0 {
			data[i] = 0x2a
			i++
			i = encodeVarintSimplevpn(data, i, uint64(len(m.PublicKey)))
			i += copy(data[i:], m.PublicKey)
		}
	}
	if m.Extension != nil {
		if len(m.Extension) > 0 {
			data[i] = 0x32
			i++
			i = encodeVarintSimplevpn(data, i, uint64(len(m.Extension)))
			i += copy(data[i:], m.Extension)
		}
	}
	if len(m.PubBase64) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.PubBase64)))
		i += copy(data[i:], m.PubBase64)
	}
	if len(m.PriBase64) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.PriBase64)))
		i += copy(data[i:], m.PriBase64)
	}
	return i, nil
}

func (m *Auth) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Auth) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintSimplevpn(data, i, uint64(m.Type))
	}
	if m.MacAddr != nil {
		if len(m.MacAddr) > 0 {
			data[i] = 0x12
			i++
			i = encodeVarintSimplevpn(data, i, uint64(len(m.MacAddr)))
			i += copy(data[i:], m.MacAddr)
		}
	}
	if len(m.IP) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.IP)))
		i += copy(data[i:], m.IP)
	}
	if len(m.GateWay) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.GateWay)))
		i += copy(data[i:], m.GateWay)
	}
	if len(m.DNS) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintSimplevpn(data, i, uint64(len(m.DNS)))
		i += copy(data[i:], m.DNS)
	}
	if m.MTU != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintSimplevpn(data, i, uint64(m.MTU))
	}
	if m.Extension != nil {
		if len(m.Extension) > 0 {
			data[i] = 0x3a
			i++
			i = encodeVarintSimplevpn(data, i, uint64(len(m.Extension)))
			i += copy(data[i:], m.Extension)
		}
	}
	return i, nil
}

func encodeFixed64Simplevpn(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Simplevpn(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSimplevpn(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovSimplevpn(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	if m.PublicKey != nil {
		l = len(m.PublicKey)
		if l > 0 {
			n += 1 + l + sovSimplevpn(uint64(l))
		}
	}
	if m.Extension != nil {
		l = len(m.Extension)
		if l > 0 {
			n += 1 + l + sovSimplevpn(uint64(l))
		}
	}
	l = len(m.PubBase64)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	l = len(m.PriBase64)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	return n
}

func (m *Auth) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovSimplevpn(uint64(m.Type))
	}
	if m.MacAddr != nil {
		l = len(m.MacAddr)
		if l > 0 {
			n += 1 + l + sovSimplevpn(uint64(l))
		}
	}
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	l = len(m.GateWay)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	l = len(m.DNS)
	if l > 0 {
		n += 1 + l + sovSimplevpn(uint64(l))
	}
	if m.MTU != 0 {
		n += 1 + sovSimplevpn(uint64(m.MTU))
	}
	if m.Extension != nil {
		l = len(m.Extension)
		if l > 0 {
			n += 1 + l + sovSimplevpn(uint64(l))
		}
	}
	return n
}

func sovSimplevpn(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSimplevpn(x uint64) (n int) {
	return sovSimplevpn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSimplevpn
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (User_UserType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extension", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extension = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubBase64", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubBase64 = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriBase64", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriBase64 = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSimplevpn(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSimplevpn
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Auth) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSimplevpn
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Auth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Auth_MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MacAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MacAddr = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateWay", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateWay = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DNS", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DNS = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MTU", wireType)
			}
			m.MTU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MTU |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extension", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSimplevpn
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extension = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSimplevpn(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSimplevpn
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSimplevpn(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSimplevpn
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSimplevpn
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSimplevpn
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSimplevpn
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSimplevpn(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSimplevpn = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSimplevpn   = fmt.Errorf("proto: integer overflow")
)
