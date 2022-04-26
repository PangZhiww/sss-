package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/GetUserinfo/handler"
	GetUserInfo "sss/GetUserinfo/proto/GetUserinfo"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserinfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetUserInfo.RegisterGetUserinfoHandler(service.Server(), new(handler.GetUserinfo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
