package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"sss/GetSmscd/handler"
	"sss/GetSmscd/subscriber"

	GetSmsCD "sss/GetSmscd/proto/GetSmscd"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetSmscd"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetSmsCD.RegisterGetSmscdHandler(service.Server(), new(handler.GetSmscd))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSmscd", service.Server(), new(subscriber.GetSmscd))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSmscd", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
