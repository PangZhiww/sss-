package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	GETUserHouses "sss/GetUserHouses/proto/GetUserHouses"
)

type GetUserHouses struct{}

func (e *GetUserHouses) Handle(ctx context.Context, msg *GETUserHouses.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GETUserHouses.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
