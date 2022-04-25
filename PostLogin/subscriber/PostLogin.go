package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	POSTLogin "sss/PostLogin/proto/PostLogin"
)

type PostLogin struct{}

func (e *PostLogin) Handle(ctx context.Context, msg *POSTLogin.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *POSTLogin.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
