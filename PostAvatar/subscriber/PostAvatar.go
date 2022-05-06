package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	POSTAvatar "sss/PostAvatar/proto/PostAvatar"
)

type PostAvatar struct{}

func (e *PostAvatar) Handle(ctx context.Context, msg *POSTAvatar.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *POSTAvatar.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
