// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/af_packet.api.json

/*
Package af_packet is a generated from VPP binary API module 'af_packet'.

 The af_packet module consists of:
	  8 messages
	  4 services
*/
package af_packet

import api "git.fd.io/govpp.git/api"
import bytes "bytes"
import context "context"
import strconv "strconv"
import struc "github.com/lunixbochs/struc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = strconv.Itoa
var _ = struc.Pack

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

const (
	// ModuleName is the name of this module.
	ModuleName = "af_packet"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x206563c
)

/* Messages */

// AfPacketCreate represents VPP binary API message 'af_packet_create':
type AfPacketCreate struct {
	HostIfName      []byte `struc:"[64]byte"`
	HwAddr          []byte `struc:"[6]byte"`
	UseRandomHwAddr uint8
}

func (*AfPacketCreate) GetMessageName() string {
	return "af_packet_create"
}
func (*AfPacketCreate) GetCrcString() string {
	return "6d5d30d6"
}
func (*AfPacketCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketCreateReply represents VPP binary API message 'af_packet_create_reply':
type AfPacketCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*AfPacketCreateReply) GetMessageName() string {
	return "af_packet_create_reply"
}
func (*AfPacketCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*AfPacketCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDelete represents VPP binary API message 'af_packet_delete':
type AfPacketDelete struct {
	HostIfName []byte `struc:"[64]byte"`
}

func (*AfPacketDelete) GetMessageName() string {
	return "af_packet_delete"
}
func (*AfPacketDelete) GetCrcString() string {
	return "3efceda3"
}
func (*AfPacketDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketDeleteReply represents VPP binary API message 'af_packet_delete_reply':
type AfPacketDeleteReply struct {
	Retval int32
}

func (*AfPacketDeleteReply) GetMessageName() string {
	return "af_packet_delete_reply"
}
func (*AfPacketDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AfPacketDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDetails represents VPP binary API message 'af_packet_details':
type AfPacketDetails struct {
	SwIfIndex  uint32
	HostIfName []byte `struc:"[64]byte"`
}

func (*AfPacketDetails) GetMessageName() string {
	return "af_packet_details"
}
func (*AfPacketDetails) GetCrcString() string {
	return "057205fa"
}
func (*AfPacketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AfPacketDump represents VPP binary API message 'af_packet_dump':
type AfPacketDump struct{}

func (*AfPacketDump) GetMessageName() string {
	return "af_packet_dump"
}
func (*AfPacketDump) GetCrcString() string {
	return "51077d14"
}
func (*AfPacketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketSetL4CksumOffload represents VPP binary API message 'af_packet_set_l4_cksum_offload':
type AfPacketSetL4CksumOffload struct {
	SwIfIndex uint8
	Set       uint8
}

func (*AfPacketSetL4CksumOffload) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload"
}
func (*AfPacketSetL4CksumOffload) GetCrcString() string {
	return "86538585"
}
func (*AfPacketSetL4CksumOffload) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AfPacketSetL4CksumOffloadReply represents VPP binary API message 'af_packet_set_l4_cksum_offload_reply':
type AfPacketSetL4CksumOffloadReply struct {
	Retval int32
}

func (*AfPacketSetL4CksumOffloadReply) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload_reply"
}
func (*AfPacketSetL4CksumOffloadReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AfPacketSetL4CksumOffloadReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*AfPacketCreate)(nil), "af_packet.AfPacketCreate")
	api.RegisterMessage((*AfPacketCreateReply)(nil), "af_packet.AfPacketCreateReply")
	api.RegisterMessage((*AfPacketDelete)(nil), "af_packet.AfPacketDelete")
	api.RegisterMessage((*AfPacketDeleteReply)(nil), "af_packet.AfPacketDeleteReply")
	api.RegisterMessage((*AfPacketDetails)(nil), "af_packet.AfPacketDetails")
	api.RegisterMessage((*AfPacketDump)(nil), "af_packet.AfPacketDump")
	api.RegisterMessage((*AfPacketSetL4CksumOffload)(nil), "af_packet.AfPacketSetL4CksumOffload")
	api.RegisterMessage((*AfPacketSetL4CksumOffloadReply)(nil), "af_packet.AfPacketSetL4CksumOffloadReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AfPacketCreate)(nil),
		(*AfPacketCreateReply)(nil),
		(*AfPacketDelete)(nil),
		(*AfPacketDeleteReply)(nil),
		(*AfPacketDetails)(nil),
		(*AfPacketDump)(nil),
		(*AfPacketSetL4CksumOffload)(nil),
		(*AfPacketSetL4CksumOffloadReply)(nil),
	}
}

// Service represents VPP binary API services in af_packet module.
type Service interface {
	DumpAfPacket(ctx context.Context, in *AfPacketDump) ([]*AfPacketDetails, error)
	AfPacketCreate(ctx context.Context, in *AfPacketCreate) (*AfPacketCreateReply, error)
	AfPacketDelete(ctx context.Context, in *AfPacketDelete) (*AfPacketDeleteReply, error)
	AfPacketSetL4CksumOffload(ctx context.Context, in *AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error)
}

type service struct {
	ch api.Channel
}

func NewService(ch api.Channel) Service {
	return &service{ch}
}

func (c *service) DumpAfPacket(ctx context.Context, in *AfPacketDump) ([]*AfPacketDetails, error) {
	var dump []*AfPacketDetails
	req := c.ch.SendMultiRequest(in)
	for {
		m := new(AfPacketDetails)
		stop, err := req.ReceiveReply(m)
		if stop {
			break
		}
		if err != nil {
			return nil, err
		}
		dump = append(dump, m)
	}
	return dump, nil
}

func (c *service) AfPacketCreate(ctx context.Context, in *AfPacketCreate) (*AfPacketCreateReply, error) {
	out := new(AfPacketCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) AfPacketDelete(ctx context.Context, in *AfPacketDelete) (*AfPacketDeleteReply, error) {
	out := new(AfPacketDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) AfPacketSetL4CksumOffload(ctx context.Context, in *AfPacketSetL4CksumOffload) (*AfPacketSetL4CksumOffloadReply, error) {
	out := new(AfPacketSetL4CksumOffloadReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
