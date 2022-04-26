package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/micro/go-micro/util/log"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	POSTLogin "sss/PostLogin/proto/PostLogin"
)

type PostLogin struct{}

// PostLogin is a single request handler called via client.Call or the generated client code
func (e *PostLogin) PostLogin(ctx context.Context, req *POSTLogin.Request, rsp *POSTLogin.Response) error {
	fmt.Println("登陆 PostLogin /api/v1.0/sessions")

	// 初始化返回值
	rsp.Errno = utils.RECODE_OK
	// 错误信息
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*查询数据*/
	// 创建数据库orm句柄
	o := orm.NewOrm()

	// 创建user对象
	var user models.User

	// 创建查询条件句柄
	qs := o.QueryTable("user")

	// 通过qs句柄进行查询
	err := qs.Filter("mobile", req.Mobile).One(&user)

	if err != nil {
		fmt.Println("查询数据失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*密码校验*/
	if utils.Md5String(req.Password) != user.Password_hash {
		fmt.Println("密码错误")
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*创建sessionId 顺便就把数据返回*/
	sessionId := utils.Md5String(req.Mobile + req.Password)
	rsp.SessionId = sessionId

	/*将登陆信息进行缓存*/

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

	/*拼接Key*/
	// user_id
	sessionuser_id := sessionId + "user_id"
	bm.Put(sessionuser_id, user.Id, time.Second*600)
	fmt.Println("user.Id:", user.Id)
	// name
	sessionname := sessionId + "name"
	bm.Put(sessionname, user.Name, time.Second*600)
	// mobile
	sessionmobile := sessionId + "mobile"
	bm.Put(sessionmobile, user.Mobile, time.Second*600)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *PostLogin) Stream(ctx context.Context, req *POSTLogin.StreamingRequest, stream POSTLogin.PostLogin_StreamStream) error {
	log.Logf("Received PostLogin.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&POSTLogin.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *PostLogin) PingPong(ctx context.Context, stream POSTLogin.PostLogin_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&POSTLogin.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
