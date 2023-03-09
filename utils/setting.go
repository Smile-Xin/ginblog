package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbHost     string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string

	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	Endpoint        string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadServer(file)
	LoadData(file)
	LoadQiuniu(file)
	LoadAli(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("qwera")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("236753")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiuniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}

func LoadAli(file *ini.File) {
	AccessKeyId = file.Section("ali").Key("AccessKeyId").String()
	AccessKeySecret = file.Section("ali").Key("AccessKeySecret").String()
	BucketName = file.Section("ali").Key("BucketName").String()
	Endpoint = file.Section("ali").Key("Endpoint").String()
}
