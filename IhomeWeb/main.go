package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"net/http"
	"sss/IhomeWeb/handler"
	_ "sss/IhomeWeb/model"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// 使用路由中间件来映射页面
	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))

	// 获取地区请求
	rou.GET("/api/v1.0/areas",handler.GetArea)


	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))
	service.Handle("/", rou)

	// register call handler
	//service.HandleFunc("/IhomeWeb/call", handler.IhomeWebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
