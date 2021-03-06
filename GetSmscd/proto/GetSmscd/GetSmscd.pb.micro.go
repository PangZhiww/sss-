// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetSmscd/GetSmscd.proto

package GetSmscd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GetSmscd service

type GetSmscdService interface {
	GetSmscd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetSmscd_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetSmscd_PingPongService, error)
}

type getSmscdService struct {
	c    client.Client
	name string
}

func NewGetSmscdService(name string, c client.Client) GetSmscdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.GetSmscd"
	}
	return &getSmscdService{
		c:    c,
		name: name,
	}
}

func (c *getSmscdService) GetSmscd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetSmscd.GetSmscd", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getSmscdService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetSmscd_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetSmscd.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getSmscdServiceStream{stream}, nil
}

type GetSmscd_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getSmscdServiceStream struct {
	stream client.Stream
}

func (x *getSmscdServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getSmscdServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmscdServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmscdServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getSmscdService) PingPong(ctx context.Context, opts ...client.CallOption) (GetSmscd_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetSmscd.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getSmscdServicePingPong{stream}, nil
}

type GetSmscd_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getSmscdServicePingPong struct {
	stream client.Stream
}

func (x *getSmscdServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getSmscdServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmscdServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmscdServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getSmscdServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetSmscd service

type GetSmscdHandler interface {
	GetSmscd(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetSmscd_StreamStream) error
	PingPong(context.Context, GetSmscd_PingPongStream) error
}

func RegisterGetSmscdHandler(s server.Server, hdlr GetSmscdHandler, opts ...server.HandlerOption) error {
	type getSmscd interface {
		GetSmscd(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetSmscd struct {
		getSmscd
	}
	h := &getSmscdHandler{hdlr}
	return s.Handle(s.NewHandler(&GetSmscd{h}, opts...))
}

type getSmscdHandler struct {
	GetSmscdHandler
}

func (h *getSmscdHandler) GetSmscd(ctx context.Context, in *Request, out *Response) error {
	return h.GetSmscdHandler.GetSmscd(ctx, in, out)
}

func (h *getSmscdHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetSmscdHandler.Stream(ctx, m, &getSmscdStreamStream{stream})
}

type GetSmscd_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getSmscdStreamStream struct {
	stream server.Stream
}

func (x *getSmscdStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getSmscdStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmscdStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmscdStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getSmscdHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetSmscdHandler.PingPong(ctx, &getSmscdPingPongStream{stream})
}

type GetSmscd_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getSmscdPingPongStream struct {
	stream server.Stream
}

func (x *getSmscdPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getSmscdPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getSmscdPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getSmscdPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getSmscdPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
