package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	GetUserInfo "sss/GetUserinfo/proto/GetUserinfo"
)

type GetUserinfo struct{}

// GetUserinfo is a single request handler called via client.Call or the generated client code
func (e *GetUserinfo) GetUserinfo(ctx context.Context, req *GetUserInfo.Request, rsp *GetUserInfo.Response) error {

	fmt.Println(" GetUserinfo 获取用户信息 /api/v1.0/user ")

	/*初始化错误码*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*获取sessionId*/
	sessionId := req.SessionId

	/*连接redis*/
	// 准备连接redis信息 {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port, // 127.0.0.1:6379
		"dbNum": utils.G_redis_dbnum,
	}
	beego.Info("redis_conf: ", redisConf)

	// 将map进行转化成为json

	redisConfJson, _ := json.Marshal(redisConf)

	// 创建redis句柄
	bm, err := cache.NewCache("redis", string(redisConfJson))
	if err != nil {
		beego.Info("redis连接失败", err)
		// 初始化 错误码
		rsp.Errno = utils.RECODE_DBERR
		// 错误信息
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*拼接Key*/
	sessionuserId := sessionId + "user_id"

	/*通过key获取到user_id*/
	userId := bm.Get(sessionuserId)
	// fmt.Println(reflect.TypeOf(userId), "user_id:", userId)
	userIdStr, _ := redis.String(userId, nil)

	/*
		id := int(userId.([]uint8)[0]) 已废弃
		fmt.Println(reflect.TypeOf(id), "id:", id)
	*/

	// fmt.Println(reflect.TypeOf(userIdStr), "id:", userIdStr)
	id, _ := strconv.Atoi(userIdStr)

	/*通过user_id获取到用户表信息*/
	// 创建一个user对象
	user := models.User{Id: id}
	// 创建orm句柄
	o := orm.NewOrm()
	err = o.Read(&user)
	if err != nil {
		beego.Info("数据库获取失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*将信息返回*/
	rsp.UserId = strconv.Itoa(user.Id)
	rsp.Name = user.Name
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.Mobile = user.Mobile
	rsp.AvatarUrl = user.Avatar_url

	return nil
}
