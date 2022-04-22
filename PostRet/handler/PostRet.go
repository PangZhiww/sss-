package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	PostRET "sss/PostRet/proto/PostRet"
)

type PostRet struct{}

// PostRet is a single request handler called via client.Call or the generated client code
func (e *PostRet) PostRet(ctx context.Context, req *PostRET.Request, rsp *PostRET.Response) error {
	log.Log("Received PostRet.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *PostRet) Stream(ctx context.Context, req *PostRET.StreamingRequest, stream PostRET.PostRet_StreamStream) error {
	log.Logf("Received PostRet.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&PostRET.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *PostRet) PingPong(ctx context.Context, stream PostRET.PostRet_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&PostRET.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
