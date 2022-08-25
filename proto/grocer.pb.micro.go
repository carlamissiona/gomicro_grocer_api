// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/grocer.proto

package grocer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Grocer service

func NewGrocerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Grocer service

type GrocerService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetGroceryList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Greetings(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Grocer_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Grocer_PingPongService, error)
}

type grocerService struct {
	c    client.Client
	name string
}

func NewGrocerService(name string, c client.Client) GrocerService {
	return &grocerService{
		c:    c,
		name: name,
	}
}

func (c *grocerService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Grocer.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grocerService) GetGroceryList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Grocer.GetGroceryList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grocerService) Greetings(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Grocer.Greetings", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grocerService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Grocer_StreamService, error) {
	req := c.c.NewRequest(c.name, "Grocer.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &grocerServiceStream{stream}, nil
}

type Grocer_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type grocerServiceStream struct {
	stream client.Stream
}

func (x *grocerServiceStream) Close() error {
	return x.stream.Close()
}

func (x *grocerServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *grocerServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *grocerServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *grocerServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grocerService) PingPong(ctx context.Context, opts ...client.CallOption) (Grocer_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Grocer.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &grocerServicePingPong{stream}, nil
}

type Grocer_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type grocerServicePingPong struct {
	stream client.Stream
}

func (x *grocerServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *grocerServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *grocerServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *grocerServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *grocerServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *grocerServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Grocer service

type GrocerHandler interface {
	Call(context.Context, *Request, *Response) error
	GetGroceryList(context.Context, *Request, *Response) error
	Greetings(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Grocer_StreamStream) error
	PingPong(context.Context, Grocer_PingPongStream) error
}

func RegisterGrocerHandler(s server.Server, hdlr GrocerHandler, opts ...server.HandlerOption) error {
	type grocer interface {
		Call(ctx context.Context, in *Request, out *Response) error
		GetGroceryList(ctx context.Context, in *Request, out *Response) error
		Greetings(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Grocer struct {
		grocer
	}
	h := &grocerHandler{hdlr}
	return s.Handle(s.NewHandler(&Grocer{h}, opts...))
}

type grocerHandler struct {
	GrocerHandler
}

func (h *grocerHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.GrocerHandler.Call(ctx, in, out)
}

func (h *grocerHandler) GetGroceryList(ctx context.Context, in *Request, out *Response) error {
	return h.GrocerHandler.GetGroceryList(ctx, in, out)
}

func (h *grocerHandler) Greetings(ctx context.Context, in *Request, out *Response) error {
	return h.GrocerHandler.Greetings(ctx, in, out)
}

func (h *grocerHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GrocerHandler.Stream(ctx, m, &grocerStreamStream{stream})
}

type Grocer_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type grocerStreamStream struct {
	stream server.Stream
}

func (x *grocerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *grocerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *grocerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *grocerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *grocerStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *grocerHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GrocerHandler.PingPong(ctx, &grocerPingPongStream{stream})
}

type Grocer_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type grocerPingPongStream struct {
	stream server.Stream
}

func (x *grocerPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *grocerPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *grocerPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *grocerPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *grocerPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *grocerPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
