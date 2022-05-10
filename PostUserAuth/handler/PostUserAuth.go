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
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	POSTUserAuth "sss/PostUserAuth/proto/PostUserAuth"
)

type PostUserAuth struct{}

// PostUserAuth is a single request handler called via client.Call or the generated client code
func (e *PostUserAuth) PostUserAuth(ctx context.Context, req *POSTUserAuth.Request, rsp *POSTUserAuth.Response) error {

	fmt.Println("PostUserAuth 实名认证 /api/v1.0/user/auth ")

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

	/*通过sessionId拼接Key 查询user_id*/
	sessionUserId := sessionId + "user_id"

	userId := bm.Get(sessionUserId)
	fmt.Println("userId:", userId)
	userIdStr, _ := redis.String(userId, nil)
	fmt.Println("userIdStr:", userIdStr)
	id, _ := strconv.Atoi(userIdStr)
	fmt.Println("id:", id)

	/*通过user_id 更新表 将身份证号和姓名更新到数据库表单中*/
	// 创建user表单对象
	user := models.User{Id: id, Id_card: req.IdCard, Real_name: req.RealName}

	o := orm.NewOrm()
	_, err = o.Update(&user, "real_name", "id_card")
	if err != nil {
		beego.Info("身份信息更新失败", err)
		// 初始化 错误码
		rsp.Errno = utils.RECODE_DBERR
		// 错误信息
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*刷新下session时间*/
	bm.Put(sessionUserId, userIdStr, time.Second*600)

	return nil
}
