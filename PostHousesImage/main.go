package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/PostHousesImage/handler"

	POSTHousesImage "sss/PostHousesImage/proto/PostHousesImage"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostHousesImage"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTHousesImage.RegisterPostHousesImageHandler(service.Server(), new(handler.PostHousesImage))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
