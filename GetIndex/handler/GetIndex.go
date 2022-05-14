package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GetIndex "sss/GetIndex/proto/GetIndex"
)

type GetIndex struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetIndex) Call(ctx context.Context, req *GetIndex.Request, rsp *GetIndex.Response) error {
	log.Log("Received GetIndex.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetIndex) Stream(ctx context.Context, req *GetIndex.StreamingRequest, stream GetIndex.GetIndex_StreamStream) error {
	log.Logf("Received GetIndex.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GetIndex.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetIndex) PingPong(ctx context.Context, stream GetIndex.GetIndex_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GetIndex.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
