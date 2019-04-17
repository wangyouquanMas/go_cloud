package rpc

import (
	"context"
	"encoding/json"

	"filestore-server/service/dbproxy/mapper"
	"filestore-server/service/dbproxy/orm"
	dbProxy "filestore-server/service/dbproxy/proto"
)

// DBProxy : DBProxy结构体
type DBProxy struct{}

// ExecuteAction : 请求执行sql函数
func (db *DBProxy) ExecuteAction(ctx context.Context, req *dbProxy.ReqExec, res *dbProxy.RespExec) error {
	resList := make([]orm.ExecResult, len(req.Action))

	// TODO: 检查	req.Sequence req.Transaction两个参数，执行不同的流程
	for idx, singleAction := range req.Action {
		params := []interface{}{}
		if err := json.Unmarshal(singleAction.Params, &params); err != nil {
			resList[idx] = orm.ExecResult{
				Suc: false,
				Msg: "请求参数有误",
			}
			continue
		}

		// 默认串行执行sql函数
		execRes, err := mapper.FuncCall(singleAction.Name, params...)
		if err != nil {
			resList[idx] = orm.ExecResult{
				Suc: false,
				Msg: "函数调用有误",
			}
			continue
		}
		resList[idx] = execRes[0].Interface().(orm.ExecResult)
	}

	// TODO: 处理异常
	res.Data, _ = json.Marshal(resList)
	return nil
}
