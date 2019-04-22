#!/bin/bash

check_process() {
    sleep 1
    res=`ps aux | grep -v grep | grep $1`
    if [[ res != '' ]]; then
        echo '服务:' $1 ' 已启动'
        return 1
    else
        echo '服务:' $1 ' 启动失败'
        return 0
    fi
}

# 创建运行日志目录
logpath=/data/log/filestore-server
mkdir -p $logpath

# 切换到工程根目录
# cd $GOPATH/filestore-server
cd /data/go/work/src/filestore-server

# 微服务可以用supervisor做进程管理工具；
# 或者也可以通过docker/k8s进行部署

# 启动dbproxy service
nohup go run service/dbproxy/main.go --registry=consul >> $logpath/dbproxy.log 2>&1 &
check_process 'dbproxy'

# 启动上传service
nohup go run service/upload/main.go --registry=consul >> $logpath/upload.log 2>&1 &
check_process 'upload'

# 启动下载service
nohup go run service/download/main.go --registry=consul >> $logpath/download.log 2>&1 &
check_process 'download'

# 启动文件转移service
nohup go run service/transfer/main.go --registry=consul >> $logpath/transfer.log 2>&1 &
check_process 'transfer'

# 启动用户系统service
nohup go run service/account/main.go --registry=consul >> $logpath/account.log 2>&1 &
check_process 'account'

# 启动apigw service
nohup go run service/apigw/main.go --registry=consul >> $logpath/apigw.log 2>&1 &
check_process 'apigw'

echo '微服务启动完毕.'

