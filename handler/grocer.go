package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	grocer "grocer/proto"
)

type Grocer struct{}

// Return a new handler
func New() *Grocer {
	return &Grocer{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Grocer) GetGroceryList(ctx context.Context, req *grocer.Request, rsp *grocer.Response) error {
	log.Info("Received Grocer.Call request")
	rsp.Msg = "Hello Client: " + req.Name + " This Items Are Delivered To Address: " + req.Address + " "
	rsp.Msg = rsp.Msg + " Products: {  'Items: [{ 'Bar Soap ', 'Qty' : 2 } , { 'Tuna Canned', 'Qty' : 3 } ]  }"
	rsp.Code =200 //always has success code 200
	return nil
}

func (e *Grocer) Greetings(ctx context.Context, req *grocer.Request, rsp *grocer.Response) error {
	log.Info("Received Grocer.Call request")
	rsp.Msg = "Good day! Welcome to 7/11 Store Web Service"
	rsp.Code =200 //always has success code 200
	return nil
}

func (e *Grocer) Call(ctx context.Context, req *grocer.Request, rsp *grocer.Response) error {
	log.Info("Received Grocer.Call request")
	rsp.Msg = "Requested Params " + req.Name
	return nil
}


// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Grocer) Stream(ctx context.Context, req *grocer.StreamingRequest, stream grocer.Grocer_StreamStream) error {
	log.Infof("Received Grocer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&grocer.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Grocer) PingPong(ctx context.Context, stream grocer.Grocer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&grocer.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
