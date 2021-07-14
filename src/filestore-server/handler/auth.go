package handler

import (
	"filestore-server/common"
	"filestore-server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HTTPInterceptor : http请求拦截器
//func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
//	return http.HandlerFunc(
//		func(w http.ResponseWriter, r *http.Request) {
//			r.ParseForm()
//			username := r.Form.Get("username")
//			token := r.Form.Get("token")
//
//			//验证登录token是否有效
//			if len(username) < 3 || !IsTokenValid(token, username) {
//				// token校验失败则直接返回失败提示
//				resp := util.NewRespMsg(
//					int(common.StatusInvalidToken),
//					"token无效",
//					nil,
//				)
//				w.Write(resp.JSONBytes())
//				return
//			}
//			h(w, r)
//		})
//}

//拦截器改造
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
			username := c.Request.FormValue("username")
			token := c.Request.FormValue("token")

			//验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token, username) {
				// token校验失败则直接返回失败提示
				c.Abort()  //表明token请求失败后，后面的handler链路就不再执行了
				resp := util.NewRespMsg(
					int(common.StatusInvalidToken),
					"token无效",
					nil,
				)
				// 将write方法改成 c.json
				//w.Write(resp.JSONBytes())
				c.JSON(http.StatusOK,resp)
				return
			}
			// 将当前请求转到下一个中间件或者是handler
			c.Next()
		}
}