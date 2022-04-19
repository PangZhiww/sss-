package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"image"
	"image/png"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	//"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/service/grpc"
	"net/http"
	// 调用area的proto
	GETAREA "sss/GetArea/proto/GetArea"
	//IhomeWeb "path/to/service/proto/IhomeWeb"

	GetImageCD "sss/GetImageCd/proto/GetImageCd"
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

// GetArea 获取地区信息
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

// GetImageCd 获取验证码图片
func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取验证码图片 GetImageCd api/v1.0/imagecode/:uuid")

	// 创建服务
	server := grpc.NewService()
	server.Init()

	// 调用服务
	GetImageCdClient := GetImageCD.NewGetImageCdService("go.micro.srv.GetImageCd", server.Client())

	// 获取uuid
	uuid := ps.ByName("uuid")

	rsp, err := GetImageCdClient.GetImageCd(context.TODO(), &GetImageCD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 接收图片信息的 图片格式
	var img image.RGBA

	img.Stride = int(rsp.Stride)
	img.Pix = []uint8(rsp.Pix)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	var image captcha.Image
	image.RGBA = &img

	// 将图片发送给浏览器
	png.Encode(w, image)

}

// GetSession 获取session信息
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取session信息 GetSession api/v1.0/session")

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  utils.RECODE_SESSIONERR,
		"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
	}

	// 回传数据的时候是直接发送过去的 并没有设置数据格式 所以需要设置
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetIndex 获取首页轮播图信息
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取首页轮播图信息 GetIndex api/v1.0/house/index")

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
	}

	// 回传数据的时候是直接发送过去的 并没有设置数据格式 所以需要设置
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
