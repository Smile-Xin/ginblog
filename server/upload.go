package server

import (
	"context"
	"fmt"
	"ginbblog/utils"
	"ginbblog/utils/errmsg"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"os"
)

var (
	Accesskey  = utils.AccessKey
	SecretKey  = utils.SecretKey
	Bucket     = utils.Bucket
	QiniuSever = utils.QiniuSever
	// AccessKeyId 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	AccessKeyId     = utils.AccessKeyId
	AccessKeySecret = utils.AccessKeySecret
	// BucketName yourBucketName填写Bucket名称。
	BucketName = utils.BucketName
	// Endpoint yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	Endpoint = utils.Endpoint
)

//var SecretKey = utils.SecretKey
//var Bucket = utils.Bucket
//var QiniuSever = utils.QiniuSever
//
//// AccessKeyId 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
//var AccessKeyId = utils.AccessKeyId
//var AccessKeySecret = utils.AccessKeySecret
//
//// BucketName yourBucketName填写Bucket名称。
//var BucketName = utils.BucketName
//
//// Endpoint yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
//var Endpoint = utils.Endpoint

func Upload(file multipart.File, fileSzie int64) (string, uint) {
	// 获取uptoken 可以上传的凭证
	policy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(Accesskey, SecretKey)
	upToken := policy.UploadToken(mac)

	// 获取formuploader
	config := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	formUpLoader := storage.NewFormUploader(&config)

	// 上传文件
	extra := storage.PutExtra{}
	ret := storage.PutRet{}
	err := formUpLoader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSzie, &extra)
	if err != nil {
		fmt.Printf("七牛上传文件失败:%s", err)
		return "", errmsg.QN_UPLOAD_ERROR
	}

	url := QiniuSever + ret.Key
	return url, errmsg.SUCCESS
}

func UploadAli(file multipart.File, fileName string) (code uint, url string) {
	// 创建OSSClient实例。
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
	}
	// 创建存储空间。
	err = client.CreateBucket(BucketName)
	if err != nil {
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 将字符串"Hello OSS"上传至exampledir目录下的exampleobject.txt文件。
	fileName = "pic/" + fileName
	err = bucket.PutObject(fileName, file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	code = 200
	url = "https://blog-database-ali.oss-cn-beijing.aliyuncs.com/" + fileName
	fmt.Println(url)
	return
}
