package handler

import (
	"context"
	"fmt"

	POSTHousesImage "sss/PostHousesImage/proto/PostHousesImage"
)

type PostHousesImage struct{}

// PostHousesImage is a single request handler called via client.Call or the generated client code
func (e *PostHousesImage) PostHousesImage(ctx context.Context, req *POSTHousesImage.Request, rsp *POSTHousesImage.Response) error {

	fmt.Println("PostHousesImage 上传房屋图片流程 /api/v1.0/houses/:id/images ")

	return nil
}
