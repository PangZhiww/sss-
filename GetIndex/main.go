package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/GetIndex/handler"
	GETIndex "sss/GetIndex/proto/GetIndex"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetIndex"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETIndex.RegisterGetIndexHandler(service.Server(), new(handler.GetIndex))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
