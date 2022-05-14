package handler

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"path"
	models "sss/IhomeWeb/model"
	"sss/IhomeWeb/utils"
	"strconv"

	POSTHousesImage "sss/PostHousesImage/proto/PostHousesImage"
)

type PostHousesImage struct{}

// PostHousesImage is a single request handler called via client.Call or the generated client code
func (e *PostHousesImage) PostHousesImage(ctx context.Context, req *POSTHousesImage.Request, rsp *POSTHousesImage.Response) error {

	fmt.Println("PostHousesImage 上传房屋图片流程 /api/v1.0/houses/:id/images ")

	/*初始化返回值*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*获取文件的后缀名*/
	fmt.Println("文件后缀名:", path.Ext(req.Filename))
	// .jpg
	fileext := path.Ext(req.Filename)

	/*将获取到的图片数据转成为二进制信息存入fastdfs*/
	fileId, err := utils.UploadByButter(req.Image, fileext)
	if err != nil {
		fmt.Println("房屋图片上传失败：", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	fmt.Println("fileId:", fileId)

	/*从请求url中得到我们的house_id*/
	houseId, _ := strconv.Atoi(req.Id)

	/*创建house 对象*/
	house := models.House{Id: houseId}

	/*创建数据库句柄*/
	o := orm.NewOrm()
	err = o.Read(&house)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*判断Index_image_url是否为空*/
	if house.Index_image_url == "" {
		// 如果是空，就把这张图片设置为主图片
		house.Index_image_url = fileId
	}

	/*将该图片添加到 house 的全部图片中*/
	houseImage := models.HouseImage{House: &house, Url: fileId}
	house.Images = append(house.Images, &houseImage)

	/*将图片对象插入表单中*/
	_, err = o.Insert(&houseImage)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*对house表进行更新*/
	_, err = o.Update(&house)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*返回正确的数据回显给前端*/
	rsp.Url = fileId

	return nil
}
