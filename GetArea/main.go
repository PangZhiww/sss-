package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/GetArea/handler"
	"sss/GetArea/subscriber"

	GetArEA "sss/GetArea/proto/GetArea"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()



	// Register Handler
	GetArEA.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetArea", service.Server(), new(subscriber.GetArea))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetArea", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
