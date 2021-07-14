package route

import (
	"filestore-server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router()*gin.Engine{

	//1 定义包含中间件的路由器
	  router := gin.Default()


	  // 2 处理静态资源
	  router.Static("/static/","./static")

	  // 3 不需要验证就可以访问的接口
	//http.HandleFunc("/user/signup", handler.SignupHandler)
	router.GET("/user/signup", handler.SignupHandler)
	//http.HandleFunc("/user/signin", handler.SignInHandler)
	router.POST("/user/signup", handler.DoSignupHandler)
	//http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))

	router.GET("user/signup",handler.SignInHandler)
	  router.GET("user/signin",handler.DoSignHandler)

	  //加入中间件 【用于校验token的拦截器】
	  router.Use(handler.HTTPInterceptor())
	// 在该拦截器下的所有handler都会经过token校验

}
