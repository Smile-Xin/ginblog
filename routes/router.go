package routes

import (
	v1 "ginbblog/api/v1"
	"ginbblog/middleware"
	"ginbblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	front := r.Group("api/v1")
	{
		front.POST("login", v1.Login)
		front.POST("loginfront", v1.FrontLogin)
		// user业务
		front.POST("user/add", v1.AddUser)
		front.GET("user/get", v1.GetUser)

		// article业务
		front.GET("article/query", v1.QueryArticle)
		front.GET("article/list", v1.QueryArtList)
		front.GET("article/get", v1.GetArticle)
		// category业务
		front.GET("category/get", v1.GetCategory)
		front.GET("category/query", v1.QueryCategory)
		// profile 业务
		front.GET("profile/:id", v1.GetProfile)

		// 评论模块
		front.POST("addcomment", v1.AddComment)
		front.GET("comment/info/:id", v1.GetComment)
		front.GET("commentfront/:id", v1.GetCommentListFront)
		front.GET("commentcount/:id", v1.GetCommentCount)

	}
	auth := r.Group("api/v1", middleware.JwtToken())
	{
		// user业务
		//auth.GET("user/query", v1.QueryUser)
		//auth.GET("user/exist", v1.ExistUser)
		//auth.GET("user/get", v1.GetUser)
		//auth.POST("user/add", v1.AddUser)
		auth.POST("user/edit", v1.EditUser)
		auth.DELETE("user/delete", v1.DeleteUser)

		// article业务
		//auth.GET("article/query", v1.QueryArticle)
		//auth.GET("article/list", v1.QueryArtList)
		//auth.GET("article/get", v1.GetArticle)
		auth.POST("article/add", v1.AddArticle)
		auth.POST("article/edit", v1.EditArticle)
		auth.DELETE("article/delete", v1.DeleteArticle)

		// category业务
		//auth.GET("category/get", v1.GetCategory)
		//auth.GET("category/query", v1.QueryCategory)
		auth.POST("category/add", v1.AddCategory)
		auth.POST("category/edit", v1.EditCategory)
		auth.DELETE("category/delete", v1.DeleteCategory)

		// profile 业务
		//auth.GET("profile/:id", v1.GetProfile)
		auth.POST("profile/edit/:id", v1.EditProfile)
		auth.GET("admin/profile/:id", v1.GetProfile)

		// 评论模块
		auth.GET("comment/get", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)

		// Pic业务
		//auth.GET("pic/query", v1.QueryPic)
		//auth.POST("pic/add", v1.AddPic)

		// 阿里业务
		// 上传文件
		auth.POST("ali/upload", v1.UploadAli)

		// 七牛业务
		// 上传文件
		auth.POST("qiniu/upload", v1.Upload)
	}

	r.Run(utils.HttpPort)
}
