package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/PostHouses/handler"
	POSTHouses "sss/PostHouses/proto/PostHouses"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTHouses.RegisterPostHousesHandler(service.Server(), new(handler.PostHouses))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
