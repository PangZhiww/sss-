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
	rou.GET("/api/v1.0/areas", handler.GetArea)
	// 获取首页轮播图
	rou.GET("/api/v1.0/house/index", handler.GetIndex)
	// 获取验证码图片
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	// 获取短信验证码
	rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmscd)
	// 获取短信验证码
	rou.POST("/api/v1.0/users", handler.PostRet)
	// 获取session
	rou.GET("/api/v1.0/session", handler.GetSession)
	// 登陆
	rou.POST("/api/v1.0/sessions", handler.PostLogin)

	// register html handler 映射前端页面
	//service.Handle("/", http.FileServer(http.Dir("html")))
	service.Handle("/", rou)

	// register call handler
	//service.HandleFunc("/IhomeWeb/call", handler.IhomeWebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
