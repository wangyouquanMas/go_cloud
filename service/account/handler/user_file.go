package handler

import (
	"context"
	"encoding/json"
	proto "filestore-server/service/account/proto"

	"filestore-server/common"
	dblayer "filestore-server/db"
)

// UserFiles : 获取用户文件列表
func (u *User) UserFiles(ctx context.Context, req *proto.ReqUserFile, res *proto.RespUserFile) error {
	userFiles, err := dblayer.QueryUserFileMetas(req.Username, int(req.Limit))
	if err != nil {
		res.Code = common.StatusServerError
		return err
	}

	data, err := json.Marshal(userFiles)
	if err != nil {
		res.Code = common.StatusServerError
		return nil
	}

	res.FileData = data
	return nil
}
