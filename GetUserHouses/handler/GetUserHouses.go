package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GetUserHouses "sss/GetUserHouses/proto/GetUserHouses"
)

type GetUserHouses struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetUserHouses) Call(ctx context.Context, req *GetUserHouses.Request, rsp *GetUserHouses.Response) error {
	log.Log("Received GetUserHouses.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetUserHouses) Stream(ctx context.Context, req *GetUserHouses.StreamingRequest, stream GetUserHouses.GetUserHouses_StreamStream) error {
	log.Logf("Received GetUserHouses.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GetUserHouses.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetUserHouses) PingPong(ctx context.Context, stream GetUserHouses.GetUserHouses_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GetUserHouses.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
