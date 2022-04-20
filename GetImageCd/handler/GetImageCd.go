package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/micro/go-micro/util/log"
	"image/color"
	"sss/IhomeWeb/utils"
	"time"

	GetImageCD "sss/GetImageCd/proto/GetImageCd"

	_ "github.com/astaxie/beego/cache/redis"

	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

type GetImageCd struct{}

// GetImageCd is a single request handler called via client.Call or the generated client code
func (e *GetImageCd) GetImageCd(ctx context.Context, req *GetImageCD.Request, rsp *GetImageCD.Response) error {

	fmt.Println("获取验证码图片 GetImageCd api/v1.0/imagecode/:uuid")

	/*
		生成验证码图片
	*/
	cap := captcha.New() // 创建图片句柄

	// 设置字体
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(90, 41)                                                                                 // 设置图片大小
	cap.SetDisturbance(captcha.MEDIUM)                                                                  // 设置干扰强度
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})                                                   // 设置前景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255}) // 设置背景色

	// 生成随机的验证码图片
	img, str := cap.Create(4, captcha.NUM) // 生成设置款式的图片
	fmt.Println("验证码:", str)

	// 将uuid和随即验证码进行缓存 配置缓存参数
	redisConf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port, // 127.0.0.1:6379
		"dbNum": utils.G_redis_dbnum,
	}
	fmt.Println("redis_conf: ", redisConf)

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
	}
	// 把验证码与uuid进行缓存
	bm.Put(req.Uuid, str, time.Second*300)

	// 图片解引用
	img1 := *img
	img2 := *img1.RGBA

	// 返回错误信息
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	// 返回图片拆分
	rsp.Pix = []byte(img2.Pix)
	rsp.Stride = int64(img2.Stride)

	rsp.Max = &GetImageCD.Response_Point{
		X: int64(img2.Rect.Max.X),
		Y: int64(img2.Rect.Max.Y),
	}
	rsp.Min = &GetImageCD.Response_Point{
		X: int64(img2.Rect.Min.X),
		Y: int64(img2.Rect.Min.Y),
	}

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetImageCd) Stream(ctx context.Context, req *GetImageCD.StreamingRequest, stream GetImageCD.GetImageCd_StreamStream) error {
	log.Logf("Received GetImageCd.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GetImageCD.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetImageCd) PingPong(ctx context.Context, stream GetImageCD.GetImageCd_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GetImageCD.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
