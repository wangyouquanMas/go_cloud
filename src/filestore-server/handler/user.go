package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"

	// "io/ioutil"
	"net/http"
	"time"

	dblayer "filestore-server/db"
	"filestore-server/util"
)

const (
	// 用于加密的盐值(自定义)
	pwdSalt = "*#890"
)

// SignupHandler : 处理用户注册请求
//func SignupHandler(w http.ResponseWriter, r *http.Request) {
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound,"/static/view/signup.html")
    return
}


//gin框架处理注册post请求
func DoSignupHandler(c *gin.Context){

	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if len(username) < 3 || len(passwd) < 5 {
		//统一的json格式，有利于业务上的统一
		c.JSON(http.StatusOK,gin.H{
			"msg":"Invalid parameter",
			"code":-1,   //定义一个枚举，将所有可能发生的状态码写出进行具体分析
		})
		return
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	// 将用户信息注册到用户表中
	suc := dblayer.UserSignup(username, encPasswd)
	if suc {
		c.JSON(http.StatusOK,gin.H{
			"msg":"Login Succeeded",
			"code":0,
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"msg":"Login failed",
			"code":-2,
		})
	}
}



// SignInHandler : 响应登陆页面
func SignInHandler(c *gin.Context) {
	//if r.Method == http.MethodGet {
	//	// data, err := ioutil.ReadFile("./static/view/signin.html")
	//	// if err != nil {
	//	// 	w.WriteHeader(http.StatusInternalServerError)
	//	// 	return
	//	// }
	//	// w.Write(data)
	//	http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
	//	return
	//}

	c.Redirect(http.StatusFound,"/static/view/signin.html")
}


// 处理登陆Post请求
func DoSignHandler(c *gin.Context){
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	encPasswd := util.Sha1([]byte(password + pwdSalt))

	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		//w.Write([]byte("FAILED"))
		//
		//return
		c.JSON(http.StatusOK,gin.H{
			"msg":"Login failed",
			"code":-1,
		})
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		//w.Write([]byte("FAILED"))
		//return
		c.JSON(http.StatusOK,gin.H{
			"msg":"Login failed",
			"code":-2,
		})
	}

	// 3. 登录成功后重定向到首页
	//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			//Location: "http://" + r.Host + "/static/view/home.html",
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":"Login succeeded",
		"code":0,
	})

	//响应的 response 以  json格式返回给客户端
	c.Data(http.StatusOK,"application/json",resp.JSONBytes())

}


// UserInfoHandler ： 查询用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	//	token := r.Form.Get("token")

	// // 2. 验证token是否有效
	// isValidToken := IsTokenValid(token)
	// if !isValidToken {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }

	// 3. 查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 4. 组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// IsTokenValid : token是否有效
func IsTokenValid(token string, username string) bool {
	if len(token) != 40 {
		fmt.Println("token invalid: " + token)
		return false
	}
	// example，假设token的有效期为1天   (根据同学们反馈完善, 相对于视频有所更新)
	tokenTS := token[32:40]
	if util.Hex2Dec(tokenTS) < time.Now().Unix()-86400 {
		fmt.Println("token expired: " + token)
		return false
	}
	// example, IsTokenValid方法增加传入参数username
	if dblayer.GetUserToken(username) != token {
		return false
	}

	return true
}
