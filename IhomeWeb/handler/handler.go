package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	models "sss/IhomeWeb/model"

	//"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/service/grpc"
	"net/http"
	// 调用area的proto
	GETAREA "sss/GetArea/proto/GetArea"
	//IhomeWeb "path/to/service/proto/IhomeWeb"
)

/*
func IhomeWebCall(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	IhomeWebClient := IhomeWeb.NewIhomeWebService("go.micro.srv.IhomeWeb", client.DefaultClient)
	rsp, err := IhomeWebClient.Call(context.TODO(), &IhomeWeb.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


*/

func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("请求地区信息 GetArea api/v1.0/areas")

	// 创建服务 获取句柄
	server := grpc.NewService()
	// 服务初始化
	server.Init()

	// call the backend service 调用服务返回句柄
	IhomeWebClient := GETAREA.NewGetAreaService("go.micro.srv.GetArea", server.Client())
	// 调用服务返回数据
	rsp, err := IhomeWebClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 接受数据
	// 准备接收切片
	area_list := []models.Area{}
	// 循环接收数据
	for _, value := range rsp.Data {
		tmp := models.Area{Id: int(value.Aid), Name: value.Aname}
		area_list = append(area_list, tmp)
	}

	// we want to augment the response 返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.Errmsg,
		"data":   area_list,
	}

	// 回传数据的时候是直接发送过去的 并没有设置数据格式 所以需要设置
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
