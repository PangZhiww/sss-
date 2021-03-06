package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	DELETESession "sss/DeleteSession/proto/DeleteSession"
)

type DeleteSession struct{}

func (e *DeleteSession) Handle(ctx context.Context, msg *DELETESession.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *DELETESession.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
