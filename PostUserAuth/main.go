package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/PostUserAuth/handler"
	POSTUserAuth "sss/PostUserAuth/proto/PostUserAuth"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTUserAuth.RegisterPostUserAuthHandler(service.Server(), new(handler.PostUserAuth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
