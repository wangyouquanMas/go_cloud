package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	userProto "filestore-server/service/account/proto"
)

// FileQueryHandler : 查询批量的文件元信息
func FileQueryHandler(c *gin.Context) {
	limitCnt, _ := strconv.Atoi(c.Request.FormValue("limit"))
	username := c.Request.FormValue("username")

	rpcResp, err := userCli.UserFiles(context.TODO(), &userProto.ReqUserFile{
		Username: username,
		Limit:    int32(limitCnt),
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json", rpcResp.FileData)
}
