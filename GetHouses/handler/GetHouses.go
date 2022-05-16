package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"

	"github.com/micro/go-micro/util/log"

	GETHouses "sss/GetHouses/proto/GetHouses"
)

type GetHouses struct{}

// GetHouses is a single request handler called via client.Call or the generated client code
func (e *GetHouses) GetHouses(ctx context.Context, req *GETHouses.Request, rsp *GETHouses.Response) error {

	fmt.Println("GetHouses 搜索房源 /api/v1.0/houses ")

	/*创建返回空间*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*获取url上的参数信息*/
	var aid int
	aid, _ = strconv.Atoi(req.Aid)
	var sd string
	sd = req.Sd
	var ed string
	ed = req.Ed
	var sk string
	sk = req.Sk
	var page int
	page, _ = strconv.Atoi(req.P)
	fmt.Println(aid, sd, ed, sk, page)

	/*返回json*/
	houses := []models.House{}
	// 创建orm句柄
	o := orm.NewOrm()
	// 设置查找的表
	qs := o.QueryTable("house")
	// 根据查询条件 来查找内容
	// 查找传入地区的所有房屋
	num, err := qs.Filter("area_id", aid).All(&houses)
	if err != nil {
		rsp.Errno = utils.RECODE_PARAMERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	// 计算以下所有房屋数量 除以 一页显示的数量
	totalPage := int(num)/models.HOUSE_LIST_PAGE_CAPACITY + 1
	housePage := 1

	houseList := []interface{}{}
	for _, house := range houses {
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Images")
		o.LoadRelated(&house, "Facilities")
		houseList = append(houseList, house.To_house_info())
	}

	rsp.TotalPage = int64(totalPage)
	rsp.CurrentPage = int64(housePage)
	rsp.Houses, _ = json.Marshal(houseList)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetHouses) Stream(ctx context.Context, req *GETHouses.StreamingRequest, stream GETHouses.GetHouses_StreamStream) error {
	log.Logf("Received GetHouses.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GETHouses.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetHouses) PingPong(ctx context.Context, stream GETHouses.GetHouses_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GETHouses.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
