package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GETUserHouses "sss/GetUserHouses/proto/GetUserHouses"
)

type GetUserHouses struct{}

// GetUserHouses is a single request handler called via client.Call or the generated client code
func (e *GetUserHouses) GetUserHouses(ctx context.Context, req *GETUserHouses.Request, rsp *GETUserHouses.Response) error {
	log.Log("Received GetUserHouses.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
