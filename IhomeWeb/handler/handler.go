package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/service/grpc"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"regexp"
	DELETESession "sss/DeleteSession/proto/DeleteSession"
	GETAREA "sss/GetArea/proto/GetArea"
	GetImageCD "sss/GetImageCd/proto/GetImageCd"
	GetSESSION "sss/GetSession/proto/GetSession"
	GetSmsCD "sss/GetSmscd/proto/GetSmscd"
	GETUserHouses "sss/GetUserHouses/proto/GetUserHouses"
	GetUserInfo "sss/GetUserinfo/proto/GetUserinfo"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	POSTAvatar "sss/PostAvatar/proto/PostAvatar"
	POSTHouses "sss/PostHouses/proto/PostHouses"
	POSTHousesImage "sss/PostHousesImage/proto/PostHousesImage"
	POSTLogin "sss/PostLogin/proto/PostLogin"
	PostRET "sss/PostRet/proto/PostRet"
	POSTUserAuth "sss/PostUserAuth/proto/PostUserAuth"
	"time"
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
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
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
	areaList := []models.Area{}
	// 循环接收数据
	for _, value := range rsp.Data {
		tmp := models.Area{Id: int(value.Aid), Name: value.Aname}
		areaList = append(areaList, tmp)
	}

	// we want to augment the response 返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.Errmsg,
		"data":   areaList,
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

// GetSmscd 获取短信验证码
func GetSmscd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Println(" GetSmscd 获取短信验证码 /api/v1.0/smscode/:mobile ")
	// 通过传入参数URL 下 Query 获取前端在 url 里的参数
	//fmt.Println(r.URL.Query()) // map[id:[e242f5ef-5145-49a1-a77c-de29c973d250] text:[22222]]
	// 获取参数
	text := r.URL.Query()["text"][0]
	id := r.URL.Query()["id"][0]
	mobile := ps.ByName("mobile")

	// 通过正则进行手机号的判断
	// 创建正则条件
	mobileReg := regexp.MustCompile(`0?(13|14|15|18|17)[0-9]{9}`)
	// 通过条件判断字符串是否匹配 返会true或false
	bl := mobileReg.MatchString(mobile)
	// 如果手机号不匹配那就不调用服务，直接返回错误
	if bl == false {
		// we want to augment the response 创建返回数据的map
		response := map[string]interface{}{
			"error":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}

		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json 发送数据
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	// 创建并初始化服务
	server := grpc.NewService()
	server.Init()

	// call the backend service 调用服务
	GetSmscdClient := GetSmsCD.NewGetSmscdService("go.micro.srv.GetSmscd", server.Client())
	rsp, err := GetSmscdClient.GetSmscd(context.TODO(), &GetSmsCD.Request{
		Mobile:   mobile,
		Imagestr: text,
		Uuid:     id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response 创建返回数据的map
	response := map[string]interface{}{
		"error":  rsp.Error,
		"errmsg": rsp.Errmsg,
	}

	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json 发送数据
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// PostRet 注册请求
func PostRet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("PostRet 注册请求 api/v1.0/users")

	// 服务创建
	server := grpc.NewService()
	server.Init()

	// decode the incoming request as json 接收post发送过来的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*
		for s, i := range request {
			fmt.Println("PostRet session request",s, i)
		}

	*/

	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// call the backend service 调用请求
	PostRETClient := PostRET.NewPostRetService("go.micro.srv.PostRet", server.Client())
	rsp, err := PostRETClient.PostRet(context.TODO(), &PostRET.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 读取cookie 统一cookie userlogin
	// func (r *Request) Cookie(name string) (*Cookie, error) {
	cookie, err := r.Cookie("userlogin")
	if err != nil || "" == cookie.Value {
		// 创建一个cookie对象
		cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 3600}
		// 对浏览器的cookie进行设置
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response 准备回传数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json 发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// GetSession 获取session信息
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取session信息 GetSession api/v1.0/session")

	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 直接返回说明用户为登陆
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		// 回传数据的时候是直接发送过去的 并没有设置数据格式 所以需要设置   设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json 将数据回发给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	GetSessionClient := GetSESSION.NewGetSessionService("go.micro.srv.GetSession", server.Client())
	rsp, err := GetSessionClient.GetSession(context.TODO(), &GetSESSION.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]string)
	data["name"] = rsp.UserName

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	// 回传数据的时候是直接发送过去的 并没有设置数据格式 所以需要设置
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// PostLogin 登陆
func PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// decode the incoming request as json

	fmt.Println("登陆 PostLogin /api/v1.0/sessions")

	// 接收前端发送过来的json数据进行解码
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if request["mobile"].(string) == "" || request["password"].(string) == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service 调用服务
	PostLoginClient := POSTLogin.NewPostLoginService("go.micro.srv.PostLogin", server.Client())
	rsp, err := PostLoginClient.PostLogin(context.TODO(), &POSTLogin.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{
			Name:   "userlogin",
			Value:  rsp.SessionId,
			Path:   "/",
			MaxAge: 600,
		}
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// DeleteSession 退出登陆
func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println(" DeleteSession 退出登陆 /api/v1.0/session ")

	// 创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	DeleteSessionClient := DELETESession.NewDeleteSessionService("go.micro.srv.DeleteSession", server.Client())
	fmt.Println("11111")
	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	fmt.Println(cookie)
	fmt.Println("2222")

	rsp, err := DeleteSessionClient.DeleteSession(context.TODO(), &DELETESession.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("3333")
	// 删除sessionId
	cookie, err = r.Cookie("userlogin")
	if cookie.Value != "" || err == nil {
		cookie := http.Cookie{Name: "userlogin", Path: "/", MaxAge: -1, Value: ""}
		http.SetCookie(w, &cookie)
	}
	fmt.Println("4444")
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// GetUserinfo 获取用户信息
func GetUserinfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println(" GetUserinfo 获取用户信息 /api/v1.0/user ")

	server := grpc.NewService()
	server.Init()

	// call the backend service
	GetUserinfoClient := GetUserInfo.NewGetUserinfoService("go.micro.srv.GetUserinfo", server.Client())

	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 远程调用函数
	rsp, err := GetUserinfoClient.GetUserinfo(context.TODO(), &GetUserInfo.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]interface{})
	/*
		"user_id": 1,
		"name": "Panda",
		"mobile": "110",
		"real_name": "熊猫",
		"id_card": "210112244556677",
		"avatar_url":
	*/
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// PostAvatar 上传头像
func PostAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println(" PostAvatar 上传头像 api/v1.0/user/avatar ")

	// 获取前端发送的文件信息
	// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	File, FileHeader, err := r.FormFile("avatar")
	if err != nil {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	fmt.Println("文件大小", FileHeader.Size)
	fmt.Println("文件名", FileHeader.Filename)

	// 创建一个文件大小的切片
	fileBuffrt := make([]byte, FileHeader.Size)

	// 将file的数据读到fileBuffrt里
	_, err = File.Read(fileBuffrt)
	if err != nil {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 连接服务
	client := grpc.NewService()
	client.Init()

	// call the backend service
	PostAvatarClient := POSTAvatar.NewPostAvatarService("go.micro.srv.PostAvatar", client.Client())
	rsp, err := PostAvatarClient.PostAvatar(context.TODO(), &POSTAvatar.Request{
		SessionId: cookie.Value,
		Fileext:   FileHeader.Filename,
		Filesize:  FileHeader.Size,
		Avatar:    fileBuffrt,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]string)
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetUserAuth 用户信息检查
func GetUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println(" GetUserAuth 用户信息检查 api/v1.0/user/auth ")

	server := grpc.NewService()
	server.Init()

	// call the backend service
	GetUserinfoClient := GetUserInfo.NewGetUserinfoService("go.micro.srv.GetUserinfo", server.Client())

	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 远程调用函数
	rsp, err := GetUserinfoClient.GetUserinfo(context.TODO(), &GetUserInfo.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]interface{})
	/*
		"user_id": 1,
		"name": "Panda",
		"mobile": "110",
		"real_name": "熊猫",
		"id_card": "210112244556677",
		"avatar_url":
	*/
	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// PostUserAuth 实名认证
func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("PostUserAuth 实名认证 /api/v1.0/user/auth ")

	// decode the incoming request as json 接收前端发送过来的数据解码到request
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	client := grpc.NewService()
	client.Init()

	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// call the backend service
	PostUserAuth := POSTUserAuth.NewPostUserAuthService("go.micro.srv.PostUserAuth", client.Client())
	rsp, err := PostUserAuth.PostUserAuth(context.TODO(), &POSTUserAuth.Request{

		SessionId: cookie.Value,
		IdCard:    request["id_card"].(string),
		RealName:  request["real_name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetUserHouses 获取用户已发布的房源
func GetUserHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("获取用户已发布的房源 GetUserHouses api/v1.0/user/houses ")

	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	client := grpc.NewService()
	client.Init()

	// call the backend service
	GetUserHousesClient := GETUserHouses.NewGetUserHousesService("go.micro.srv.GetUserHouses", client.Client())
	rsp, err := GetUserHousesClient.GetUserHouses(context.TODO(), &GETUserHouses.Request{
		SessionId: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	housesList := []models.House{}
	json.Unmarshal(rsp.Mix, &housesList)

	var houses []interface{}
	// 遍历返回的完整房屋信息
	for _, value := range housesList {
		// 获取到有用的添加到切片当中
		houses = append(houses, value.To_house_info())
	}

	// 创建一个data的map
	data := make(map[string]interface{})
	data["houses"] = houses

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// PostHouses 发布房源信息
func PostHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("PostHouses 发布房源信息 /api/v1.0/houses ")

	// decode the incoming request as json 将前端发送过来的数据整体读取
	// func ReadAll(r io.Reader) ([]byte, error)
	body, _ := ioutil.ReadAll(r.Body) // body就是一个json的二进制流

	// 获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// we want to augment the response 准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// 设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json 发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	client := grpc.NewService()
	client.Init()

	// call the backend service
	PostHousesClient := POSTHouses.NewPostHousesService("go.micro.srv.PostHouses", client.Client())
	rsp, err := PostHousesClient.PostHouses(context.TODO(), &POSTHouses.Request{
		SessionId: cookie.Value,
		Body:      body,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]interface{})
	data["house_id"] = rsp.HousesId

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	// 设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// PostHousesImage 上传房屋图片流程
func PostHousesImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	PostHousesImageClient := POSTHousesImage.NewIhomeWebService("go.micro.srv.IhomeWeb", client.DefaultClient)
	rsp, err := PostHousesImageClient.PostHousesImage(context.TODO(), &POSTHousesImage.Request{
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
	// 设置返回数据的格式
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
