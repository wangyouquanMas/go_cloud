package handler

import (
	"context"
	"filestore-server/service/account/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
)

var (
	userCli proto.UserService
)

func init() {
	service := micro.NewService()
	// 初始化， 解析命令行参数等
	service.Init()

	// 初始化一个rpcClient
	userCli = proto.NewUserService("go.micro.service.user", service.Client())
}

// SignupHandler : 响应注册页面
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// DoSignupHandler : 处理注册post请求
func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := userCli.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Mesaage,
	})
}
