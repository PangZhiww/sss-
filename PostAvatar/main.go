package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/PostAvatar/handler"
	POSTAvatar "sss/PostAvatar/proto/PostAvatar"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostAvatar"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTAvatar.RegisterPostAvatarHandler(service.Server(), new(handler.PostAvatar))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
