package handler

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"

	"encoding/json"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	GETIndex "sss/GetIndex/proto/GetIndex"
)

type GetIndex struct{}

// GetIndex is a single request handler called via client.Call or the generated client code
func (e *GetIndex) GetIndex(ctx context.Context, req *GETIndex.Request, rsp *GETIndex.Response) error {

	beego.Info("获取首页轮播图信息 GetIndex api/v1.0/house/index")

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	data := []interface{}{}

	/*从缓存服务器中请求 home_page_data 字段 如果有就直接返回*/
	/*先从缓存种获取房屋数据 将缓存数据返回前端即可*/

	// 准备连接redis信息 {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port, // 127.0.0.1:6379
		"dbNum": utils.G_redis_dbnum,
	}
	beego.Info("redis_conf: ", redisConf)

	// 将map进行转化成为json

	redisConfJson, _ := json.Marshal(redisConf)
	fmt.Println(string(redisConfJson))

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
	housePageKey := "home_page_data"
	housePageValue := bm.Get(housePageKey)
	if housePageValue != nil {

		fmt.Println("=============== get house page info from CACHE!!!! ===============")
		// 直接将二进制发送给客户端
		rsp.Max = housePageValue.([]byte)
		return nil

	}

	houses := []models.House{}

	/*如果缓存没有 需要从数据库中查询到房屋列表*/
	o := orm.NewOrm()
	if _, err := o.QueryTable("house").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses); err == nil {
		for _, house := range houses {
			o.LoadRelated(&house, "Area")
			o.LoadRelated(&house, "User")
			o.LoadRelated(&house, "Images")
			o.LoadRelated(&house, "Facilities")
			data = append(data, house.To_house_info())
		}
	}
	fmt.Println("data:", data)
	fmt.Println("houses:", houses)
	/*将data存入缓存数据*/
	housePageValue, _ = json.Marshal(data)
	bm.Put(housePageKey, housePageValue, time.Second*3600)
	rsp.Max = housePageValue.([]byte)

	return nil
}
