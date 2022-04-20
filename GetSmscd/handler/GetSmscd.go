package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GetSmsCD "sss/GetSmscd/proto/GetSmscd"
)

type GetSmscd struct{}

// GetSmscd is a single request handler called via client.Call or the generated client code
func (e *GetSmscd) GetSmscd(ctx context.Context, req *GetSmsCD.Request, rsp *GetSmsCD.Response) error {
	log.Log("Received GetSmscd.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetSmscd) Stream(ctx context.Context, req *GetSmsCD.StreamingRequest, stream GetSmsCD.GetSmscd_StreamStream) error {
	log.Logf("Received GetSmscd.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GetSmsCD.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetSmscd) PingPong(ctx context.Context, stream GetSmsCD.GetSmscd_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GetSmsCD.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
