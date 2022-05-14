package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"sss/GetHouseInfo/handler"
	GETHouseInfo "sss/GetHouseInfo/proto/GetHouseInfo"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetHouseInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETHouseInfo.RegisterGetHouseInfoHandler(service.Server(), new(handler.GetHouseInfo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
