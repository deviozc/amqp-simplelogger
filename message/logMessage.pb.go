// Code generated by protoc-gen-go.
// source: logMessage.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	logMessage.proto

It has these top-level messages:
	LogMessage
*/
package message

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LogMessage struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	Timestamp        *int64  `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	Type             *int32  `protobuf:"varint,3,req,name=type,def=1" json:"type,omitempty"`
	FromService      *string `protobuf:"bytes,4,req,name=fromService" json:"fromService,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}

const Default_LogMessage_Type int32 = 1

func (m *LogMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *LogMessage) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *LogMessage) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_LogMessage_Type
}

func (m *LogMessage) GetFromService() string {
	if m != nil && m.FromService != nil {
		return *m.FromService
	}
	return ""
}

func init() {
}
