package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"sss/IhomeWeb/utils"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	DELETESession "sss/DeleteSession/proto/DeleteSession"
)

type DeleteSession struct{}

// DeleteSession is a single request handler called via client.Call or the generated client code
func (e *DeleteSession) DeleteSession(ctx context.Context, req *DELETESession.Request, rsp *DELETESession.Response) error {

	fmt.Println(" DeleteSession 退出登陆 /api/v1.0/session ")

	// 返回值初始化
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*连接redis*/
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
	// 获取sessionId
	sessionId := req.SessionId

	/* 拼接key 删除session相关字段*/
	// user_id
	sessionuser_id := sessionId + "user_id"
	bm.Delete(sessionuser_id)
	// name
	sessionname := sessionId + "name"
	bm.Delete(sessionname)
	// mobile
	sessionmobile := sessionId + "mobile"
	bm.Delete(sessionmobile)

	return nil

}
