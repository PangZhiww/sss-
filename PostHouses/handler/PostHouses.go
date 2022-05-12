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
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"

	POSTHouses "sss/PostHouses/proto/PostHouses"
)

type PostHouses struct{}

// PostHouses is a single request handler called via client.Call or the generated client code
func (e *PostHouses) PostHouses(ctx context.Context, req *POSTHouses.Request, rsp *POSTHouses.Response) error {

	fmt.Println("PostHouses 发布房源信息 /api/v1.0/houses ")

	/*初始化返回值*/
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
	sessionUserId := sessionId + "user_id"

	/*查询user_id*/
	userId := bm.Get(sessionUserId)

	/*转换user_id类型*/
	userIdStr, _ := redis.String(userId, nil)
	id, _ := strconv.Atoi(userIdStr)

	/*解析对端发送过来的body*/
	var request = make(map[string]interface{})
	json.Unmarshal(req.Body, &request)

	/*准备插入数据库的对象*/
	house := models.House{}
	house.Title = request["title"].(string)
	price, _ := strconv.Atoi(request["price"].(string))
	house.Price = price * 100
	house.Address = request["address"].(string)
	house.Room_count, _ = strconv.Atoi(request["room_count"].(string))
	house.Acreage, _ = strconv.Atoi(request["acreage"].(string))
	house.Unit = request["unit"].(string)
	house.Capacity, _ = strconv.Atoi(request["capacity"].(string))
	house.Beds = request["beds"].(string)
	deposit, _ := strconv.Atoi(request["deposit"].(string))
	house.Deposit = deposit * 100
	house.Min_days, _ = strconv.Atoi(request["min_days"].(string))
	house.Max_days, _ = strconv.Atoi(request["max_days"].(string))

	area_id, _ := strconv.Atoi(request["area_id"].(string))
	area := models.Area{Id: area_id}
	house.Area = &area

	facility := []*models.Facility{}
	for _, value := range request["facility"].([]interface{}) {
		fid, _ := strconv.Atoi(value.(string)) // 将设施编号转换成为对应的类型
		ftmp := &models.Facility{Id: fid}      // 创建临时变量，使用设施编号创建的设施表对象的指针
		facility = append(facility, ftmp)
	}

	/*数据库插入操作*/
	user := models.User{Id: id}
	house.User = &user

	// 创建orm句柄
	o := orm.NewOrm()
	houseId, err := o.Insert(&house)
	if err != nil {
		beego.Info("数据库插入失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*多对多插入*/
	m2m := o.QueryM2M(&house, "Facilities")
	_, err = m2m.Add(facility)
	if err != nil {
		beego.Info("房屋设施数据库多对多插入失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*返回houses_id*/
	rsp.HousesId = strconv.Itoa(int(houseId))

	return nil
}
