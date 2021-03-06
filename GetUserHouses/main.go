package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/GetUserHouses/handler"
	GETUserHouses "sss/GetUserHouses/proto/GetUserHouses"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETUserHouses.RegisterGetUserHousesHandler(service.Server(), new(handler.GetUserHouses))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
