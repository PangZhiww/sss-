package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	POSTHouses "sss/PostHouses/proto/PostHouses"
)

type PostHouses struct{}

func (e *PostHouses) Handle(ctx context.Context, msg *POSTHouses.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *POSTHouses.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
