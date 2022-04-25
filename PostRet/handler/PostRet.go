package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"time"

	"github.com/micro/go-micro/util/log"

	PostRET "sss/PostRet/proto/PostRet"

	_ "github.com/astaxie/beego/cache/redis"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

type PostRet struct{}

// Md5String 加密函数
func Md5String(s string) string {
	// 创建一个md5对象
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// PostRet is a single request handler called via client.Call or the generated client code
func (e *PostRet) PostRet(ctx context.Context, req *PostRET.Request, rsp *PostRET.Response) error {

	fmt.Println("PostRet 注册请求 api/v1.0/users")
	// 初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*验证短信验证码*/
	// redis操作
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
	// 通过手机号获取到短信验证码
	smsCode := bm.Get(req.Mobile)
	if smsCode == nil {
		beego.Info("获取redis数据失败", err)
		// 初始化 错误码
		rsp.Errno = utils.RECODE_DBERR
		// 错误信息
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// 进行短信验证码对比
	smsCodeStr, _ := redis.String(smsCode, nil)
	if smsCodeStr != req.SmsCode {
		beego.Info("短信验证码错误", err)
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*将数据存入数据库*/
	o := orm.NewOrm()
	user := models.User{Mobile: req.Mobile, Password_hash: Md5String(req.Password), Name: req.Mobile}
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println("注册数据失败 数据库插入失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	fmt.Println("userId:", id)

	/*创建 sessionId （唯一随机码）*/
	sessionId := Md5String(req.Mobile + req.Password)
	rsp.SessionId = sessionId

	/*以 sessionId 为 key 的一部分创建 session*/
	// name 名字暂时使用手机号
	bm.Put(sessionId+"name", user.Mobile, time.Second*3600)
	// user_id
	bm.Put(sessionId+"user_id", id, time.Second*3600)
	// 手机号
	bm.Put(sessionId+"mobile", user.Mobile, time.Second*3600)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *PostRet) Stream(ctx context.Context, req *PostRET.StreamingRequest, stream PostRET.PostRet_StreamStream) error {
	log.Logf("Received PostRet.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&PostRET.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *PostRet) PingPong(ctx context.Context, stream PostRET.PostRet_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&PostRET.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
