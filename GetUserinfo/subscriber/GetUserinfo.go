package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	GetUserInfo "sss/GetUserinfo/proto/GetUserinfo"
)

type GetUserinfo struct{}

func (e *GetUserinfo) Handle(ctx context.Context, msg *GetUserInfo.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GetUserInfo.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
