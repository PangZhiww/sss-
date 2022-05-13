package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	POSTHousesImage "sss/PostHousesImage/proto/PostHousesImage"
)

type PostHousesImage struct{}

// PostHousesImage is a single request handler called via client.Call or the generated client code
func (e *PostHousesImage) PostHousesImage(ctx context.Context, req *POSTHousesImage.Request, rsp *POSTHousesImage.Response) error {
	log.Log("Received PostHousesImage.Call request")

	return nil
}
