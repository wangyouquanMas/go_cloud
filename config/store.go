package config

import (
	cmn "filestore-server/common"
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "/data/fileserver/"
	// CephRootDir : Ceph的存储路径prefix
	CephRootDir = "/ceph"
	// OSSRootDir : OSS的存储路径prefix
	OSSRootDir = "oss/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = cmn.StoreLocal
)
