package route

import (
 hdl	"filestore-server/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
		// gin framework
		router := gin.Default()

		// 静态资源处理
		router.Static("/static/", "./static")
	
		// 文件存取接口
		router.POST("/file/upload", hdl.HTTPInterceptor(hdl.UploadHandler))
		router.GET("/file/upload/suc", hdl.HTTPInterceptor(hdl.UploadSucHandler))
		router.GET("/file/meta", hdl.HTTPInterceptor(hdl.GetFileMetaHandler))
		router.POST("/file/query", hdl.HTTPInterceptor(hdl.FileQueryHandler))
		router.POST("/file/download", hdl.HTTPInterceptor(hdl.DownloadHandler))
		router.POST("/file/update", hdl.HTTPInterceptor(hdl.FileMetaUpdateHandler))
		router.POST("/file/delete", hdl.HTTPInterceptor(hdl.FileDeleteHandler))
		// 秒传接口
		router.POST("/file/fastupload", hdl.HTTPInterceptor(
			hdl.TryFastUploadHandler))
	
		router.POST("/file/downloadurl", hdl.HTTPInterceptor(
			hdl.DownloadURLHandler))
	
		// 分块上传接口
		router.POST("/file/mpupload/init",
			hdl.HTTPInterceptor(hdl.InitialMultipartUploadHandler))
		router.POST("/file/mpupload/uppart",
			hdl.HTTPInterceptor(hdl.UploadPartHandler))
		router.POST("/file/mpupload/complete",
			hdl.HTTPInterceptor(hdl.CompleteUploadHandler))
	
		// 用户相关接口
		router.GET("/user/signup", hdl.SignupHandler)
		router.GET("/user/signin", hdl.SigninHandler)
		router.POST("/user/signup", hdl.DoSignupHandler)
		router.POST("/user/signin", hdl.DoSignInHandler)
		router.POST("/user/info", hdl.HTTPInterceptor(hdl.UserInfoHandler))
	
		router.GET("/user/exists", hdl.UserExistsHandler)
		return router
}