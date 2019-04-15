package handler

import (
	"context"
	"filestore-server/common"
	"filestore-server/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"

	cmn "filestore-server/common"
	userProto "filestore-server/service/account/proto"
	upProto "filestore-server/service/upload/proto"
)

var (
	userCli userProto.UserService
	upCli   upProto.UploadService
)

func init() {
	service := micro.NewService()
	// 初始化， 解析命令行参数等
	service.Init()

	// 初始化一个account服务的客户端
	userCli = userProto.NewUserService("go.micro.service.user", service.Client())
	// 初始化一个upload服务的客户端
	upCli = upProto.NewUploadService("go.micro.service.upload", service.Client())
}

// SignupHandler : 响应注册页面
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// DoSignupHandler : 处理注册post请求
func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := userCli.Signup(context.TODO(), &userProto.ReqSignup{
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
		"msg":  resp.Message,
	})
}

// SigninHandler : 响应登录页面
func SigninHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// DoSigninHandler : 处理登录post请求
func DoSigninHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	rpcResp, err := userCli.Signin(context.TODO(), &userProto.ReqSignin{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if rpcResp.Code != cmn.StatusOK {
		c.JSON(200, gin.H{
			"msg":  "登录失败",
			"code": rpcResp.Code,
		})
		return
	}

	entryResp, err := upCli.UploadEntry(context.TODO(), &upProto.ReqEntry{})
	if err != nil {
		log.Println(err.Error())
	} else if entryResp.Code != cmn.StatusOK {
		log.Println(entryResp.Message)
	}

	// 3. 登录成功，返回用户信息
	cliResp := util.RespMsg{
		Code: int(common.StatusOK),
		Msg:  "登录成功",
		Data: struct {
			Location    string
			Username    string
			Token       string
			UploadEntry string
		}{
			Location:    "/static/view/home.html",
			Username:    username,
			Token:       rpcResp.Token,
			UploadEntry: entryResp.Entry,
		},
	}
	c.Data(http.StatusOK, "application/json", cliResp.JSONBytes())
}

// UserInfoHandler ： 查询用户信息
func UserInfoHandler(c *gin.Context) {
	// 1. 解析请求参数
	username := c.Request.FormValue("username")

	resp, err := userCli.UserInfo(context.TODO(), &userProto.ReqUserInfo{
		Username: username,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// 3. 组装并且响应用户数据
	cliResp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: gin.H{
			"Username": username,
			"SignupAt": resp.SignupAt,
			// TODO: 完善其他字段信息
			"LastActive": resp.LastActiveAt,
		},
	}
	c.Data(http.StatusOK, "application/json", cliResp.JSONBytes())
}
