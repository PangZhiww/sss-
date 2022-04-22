package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/util/log"
	"sss/IhomeWeb/utils"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	GETSession "sss/GetSession/proto/GetSession"
)

type GetSession struct{}

// GetSession is a single request handler called via client.Call or the generated client code
func (e *GetSession) GetSession(ctx context.Context, req *GETSession.Request, rsp *GETSession.Response) error {

	beego.Info("获取session信息 GetSession api/v1.0/session")

	// 初始化返回值
	rsp.Errno = utils.RECODE_OK
	// 错误信息
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*获取username*/

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
	username := bm.Get(req.SessionId + "name")

	/*没有 就返回失败*/
	if username == nil {
		beego.Info("获取数据并不存在", err)
		rsp.Errno = utils.RECODE_USERERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*有 就返回成功*/
	rsp.UserName, _ = redis.String(username, nil)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetSession) Stream(ctx context.Context, req *GETSession.StreamingRequest, stream GETSession.GetSession_StreamStream) error {
	log.Logf("Received GetSession.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GETSession.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetSession) PingPong(ctx context.Context, stream GETSession.GetSession_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GETSession.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
