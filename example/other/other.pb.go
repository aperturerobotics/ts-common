// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.6.2
// source: github.com/aperturerobotics/common/example/other/other.proto

package example_other

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
	strings "strings"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
)

// OtherMsg is a different message from ExampleMsg.
type OtherMsg struct {
	unknownFields []byte
	// FooField is an example field.
	FooField uint32 `protobuf:"varint,1,opt,name=foo_field,json=fooField,proto3" json:"fooField,omitempty"`
}

func (x *OtherMsg) Reset() {
	*x = OtherMsg{}
}

func (*OtherMsg) ProtoMessage() {}

func (x *OtherMsg) GetFooField() uint32 {
	if x != nil {
		return x.FooField
	}
	return 0
}

func (m *OtherMsg) CloneVT() *OtherMsg {
	if m == nil {
		return (*OtherMsg)(nil)
	}
	r := new(OtherMsg)
	r.FooField = m.FooField
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OtherMsg) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *OtherMsg) EqualVT(that *OtherMsg) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.FooField != that.FooField {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OtherMsg) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*OtherMsg)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// MarshalProtoJSON marshals the OtherMsg message to JSON.
func (x *OtherMsg) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.FooField != 0 || s.HasField("fooField") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("fooField")
		s.WriteUint32(x.FooField)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the OtherMsg to JSON.
func (x *OtherMsg) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the OtherMsg message from JSON.
func (x *OtherMsg) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "foo_field", "fooField":
			s.AddField("foo_field")
			x.FooField = s.ReadUint32()
		}
	})
}

// UnmarshalJSON unmarshals the OtherMsg from JSON.
func (x *OtherMsg) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

func (m *OtherMsg) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OtherMsg) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OtherMsg) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.FooField != 0 {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(m.FooField))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OtherMsg) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FooField != 0 {
		n += 1 + protobuf_go_lite.SizeOfVarint(uint64(m.FooField))
	}
	n += len(m.unknownFields)
	return n
}

func (x *OtherMsg) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("OtherMsg { ")
	if x.FooField != 0 {
		sb.WriteString(" foo_field: ")
		sb.WriteString(strconv.FormatUint(uint64(x.FooField), 10))
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *OtherMsg) String() string {
	return x.MarshalProtoText()
}
func (m *OtherMsg) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OtherMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OtherMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FooField", wireType)
			}
			m.FooField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FooField |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
