package handler

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"

	"encoding/json"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	GETHouseInfo "sss/GetHouseInfo/proto/GetHouseInfo"
)

type GetHouseInfo struct{}

// GetHouseInfo is a single request handler called via client.Call or the generated client code
func (e *GetHouseInfo) GetHouseInfo(ctx context.Context, req *GETHouseInfo.Request, rsp *GETHouseInfo.Response) error {

	fmt.Println("GetHouseInfo 获取房源详细信息 /api/v1.0/houses/:id ")

	/*创建返回空间（初始化返回值）*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*从session中获取我们的user_id的字段 得到当前用户Id*/
	/*通过session 获取我们当前登陆用户的user_id*/
	/*构建连接缓存的数据*/
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
	sessionIdUserId := req.SessionId + "user_id"

	valueId := bm.Get(sessionIdUserId)
	fmt.Println("valueId: ", valueId, reflect.TypeOf(valueId))
	id := int(valueId.([]uint8)[0])
	fmt.Println(id, reflect.TypeOf(id))

	/*从请求中的url获取房源Id*/
	houseId, _ := strconv.Atoi(req.Id)

	/*从缓存数据库中获取到当前房屋的数据*/
	houseInfoKey := fmt.Sprintf("house_info_%s", houseId)
	houseInfoValue := bm.Get(houseInfoKey)
	if houseInfoValue != nil {
		rsp.UserId = int64(id)
		rsp.HouseData = houseInfoValue.([]byte)
	}

	/*查询当前数据库得到当前house详细信息*/
	// 创建数据对象
	house := models.House{Id: houseId}

	// 创建数据库句柄
	o := orm.NewOrm()
	o.Read(&house)

	/*关联查询 area user images fac 等表*/
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "User")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")
	//o.LoadRelated(&house,"Orders")

	/*将查询到的结果存储到缓存当中*/
	houseMix, err := json.Marshal(house)
	bm.Put(houseInfoKey, houseMix, time.Second*3600)

	/*返回正确数据给前端*/
	rsp.UserId = int64(id)
	rsp.HouseData = houseMix

	return nil

}
