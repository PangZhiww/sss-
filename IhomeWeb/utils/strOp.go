package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ljl8086/go_fdfs_client"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求

func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}

// Md5String 加密函数
func Md5String(s string) string {
	// 创建一个md5对象
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// UploadByButter 上传二进制文件到fdfs中的操作
func UploadByButter(filebuffer []byte, fileExt string) (fileId string, err error) {
	fdClient, err := fdfs_client.NewFdfsClient("/home/microv1/go/src/sss/IhomeWeb/conf/client.conf")
	if err != nil {
		fmt.Println("创建上传图片句柄失败", err)
		fileId = ""
		return
	}
	fdRsp, err := fdClient.UploadByBuffer(filebuffer, fileExt)
	if err != nil {
		fmt.Println("上传图片失败", err)
		fileId = ""
		return
	}
	fmt.Println("fdRsp.GroupName", fdRsp.GroupName)
	fmt.Println("fdRsp.RemoteFileId", fdRsp.RemoteFileId)

	fileId = fdRsp.RemoteFileId

	return fileId, nil

}
