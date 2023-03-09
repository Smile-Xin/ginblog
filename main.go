package main

import (
	"ginbblog/dao"
	"ginbblog/routes"
)

func main() {

	// 引用数据库
	dao.InitDb()

	routes.InitRouter()
}
