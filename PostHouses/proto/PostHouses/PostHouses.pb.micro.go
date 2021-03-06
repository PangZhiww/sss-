// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/PostHouses/PostHouses.proto

package PostHouses

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

// Client API for PostHouses service

type PostHousesService interface {
	PostHouses(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type postHousesService struct {
	c    client.Client
	name string
}

func NewPostHousesService(name string, c client.Client) PostHousesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.PostHouses"
	}
	return &postHousesService{
		c:    c,
		name: name,
	}
}

func (c *postHousesService) PostHouses(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PostHouses.PostHouses", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostHouses service

type PostHousesHandler interface {
	PostHouses(context.Context, *Request, *Response) error
}

func RegisterPostHousesHandler(s server.Server, hdlr PostHousesHandler, opts ...server.HandlerOption) error {
	type postHouses interface {
		PostHouses(ctx context.Context, in *Request, out *Response) error
	}
	type PostHouses struct {
		postHouses
	}
	h := &postHousesHandler{hdlr}
	return s.Handle(s.NewHandler(&PostHouses{h}, opts...))
}

type postHousesHandler struct {
	PostHousesHandler
}

func (h *postHousesHandler) PostHouses(ctx context.Context, in *Request, out *Response) error {
	return h.PostHousesHandler.PostHouses(ctx, in, out)
}
