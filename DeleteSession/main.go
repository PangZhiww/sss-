package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/DeleteSession/handler"
	DELETESession "sss/DeleteSession/proto/DeleteSession"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.DeleteSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	DELETESession.RegisterDeleteSessionHandler(service.Server(), new(handler.DeleteSession))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
