package handler

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"time"

	"github.com/micro/go-micro/util/log"

	GetSmsCD "sss/GetSmscd/proto/GetSmscd"

	"encoding/json"
	// redis缓存操作与支持驱动
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

type GetSmscd struct{}

// GetSmscd is a single request handler called via client.Call or the generated client code
func (e *GetSmscd) GetSmscd(ctx context.Context, req *GetSmsCD.Request, rsp *GetSmsCD.Response) error {

	fmt.Println(" GetSmscd 获取短信验证码 /api/v1.0/smscode/:mobile ")

	// 初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	/*验证手机号是否存在*/
	// 创建数据库orm句柄
	o := orm.NewOrm()
	// 使用手机号作为查询条件
	user := models.User{Mobile: req.Mobile}
	err := o.Read(&user)
	// fmt.Println("err",err)
	// 如果不报错就说明查到了  查找到就说明手机号存在
	if err == nil {
		fmt.Println("用户已存在")
		rsp.Error = utils.RECODE_MOBILEUSERERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	/*验证图片验证码是否正确*/
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
		rsp.Error = utils.RECODE_DBERR
		// 错误信息
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 通过uuid查找图片验证码的值进行对比
	value := bm.Get(req.Uuid)
	if value == nil {
		beego.Info("redis获取失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//fmt.Println("打印value格式:", reflect.TypeOf(value), value) // reflect.TypeOf(value) 会返还的那个前数据的变量类型

	// 格式转换
	valueStr, _ := redis.String(value, nil)
	if valueStr != req.Imagestr {
		fmt.Println("数据不匹配 图片验证码值错误")
		rsp.Error = utils.RECODE_IMAGEERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	/*调用 短信接口发送短信*/
	// 创建随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(9999) + 1001
	fmt.Println("短信验证码：", size)
	/*将短信验证码存入redis缓存数据库*/
	err = bm.Put(req.Mobile, size, time.Second*300)
	if err != nil {
		beego.Info("redis创建失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetSmscd) Stream(ctx context.Context, req *GetSmsCD.StreamingRequest, stream GetSmsCD.GetSmscd_StreamStream) error {
	log.Logf("Received GetSmscd.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GetSmsCD.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetSmscd) PingPong(ctx context.Context, stream GetSmsCD.GetSmscd_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GetSmsCD.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
