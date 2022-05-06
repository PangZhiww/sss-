package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	POSTAvatar "sss/PostAvatar/proto/PostAvatar"
)

type PostAvatar struct{}

// PostAvatar is a single request handler called via client.Call or the generated client code
func (e *PostAvatar) PostAvatar(ctx context.Context, req *POSTAvatar.Request, rsp *POSTAvatar.Response) error {
	log.Log("Received PostAvatar.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
