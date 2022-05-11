package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	GETUserHouses "sss/GetUserHouses/proto/GetUserHouses"
)

type GetUserHouses struct{}

// GetUserHouses is a single request handler called via client.Call or the generated client code
func (e *GetUserHouses) GetUserHouses(ctx context.Context, req *GETUserHouses.Request, rsp *GETUserHouses.Response) error {
	fmt.Println("获取用户已发布的房源 GetUserHouses api/v1.0/user/houses ")

	/*初始化 返回值*/
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
		return nil
	}

	/*拼接key*/
	sessionIdUserId := sessionId + "user_id"

	/*查询对应的user_id*/
	userId := bm.Get(sessionIdUserId)

	/*转换格式*/
	userIdStr, _ := redis.String(userId, nil)
	id, _ := strconv.Atoi(userIdStr)

	/*查询数据库*/
	o := orm.NewOrm()
	qs := o.QueryTable("house")

	houseList := []models.House{}

	/*获得当前用户房屋信息*/
	_, err = qs.Filter("user_id", id).All(&houseList)
	if err != nil {
		beego.Info("查询房屋数据库失败", err)
		// 初始化 错误码
		rsp.Errno = utils.RECODE_DBERR
		// 错误信息
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*json编码成为二进制 返回*/
	house, _ := json.Marshal(houseList)
	/*返回二进制数据*/
	rsp.Mix = house

	return nil
}
