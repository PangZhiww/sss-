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
	"path"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	POSTAvatar "sss/PostAvatar/proto/PostAvatar"
)

type PostAvatar struct{}

// PostAvatar is a single request handler called via client.Call or the generated client code
func (e *PostAvatar) PostAvatar(ctx context.Context, req *POSTAvatar.Request, rsp *POSTAvatar.Response) error {

	fmt.Println(" PostAvatar 上传头像 api/v1.0/user/avatar ")

	/*初始化返回值*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	size := len(req.Avatar)

	/*图片数据验证*/
	if req.Filesize != int64(size) {
		fmt.Println("传输数据丢失")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*获取文件的后缀名*/
	// func Ext(path string) string
	ext := path.Ext(req.Fileext)

	/*调用fdfs函数上传到图片服务器*/
	fileId, err := utils.UploadByButter(req.Avatar, ext[1:])
	if err != nil {
		fmt.Println("上传失败：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*得到fileId*/
	fmt.Println("fileId:", fileId)

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
	}

	/*拼接key获取当前用户的user_id*/
	sessionUserId := sessionId + "user_id"
	userId := bm.Get(sessionUserId)
	userIdStr, _ := redis.String(userId, nil)
	id, _ := strconv.Atoi(userIdStr)

	/*将图片的存储地址（fileId）更新到user表中*/
	// 创建user表对象
	user := models.User{Id: id, Avatar_url: fileId}
	// 连接数据库
	o := orm.NewOrm()
	_, err = o.Update(&user, "avatar_url")
	if err != nil {
		fmt.Println("数据库表单更新失败：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*回传fileId*/
	rsp.AvatarUrl = fileId

	return nil
}
