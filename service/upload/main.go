package main

import (
	cfg "filestore-server/config"
	"filestore-server/route"
)

func main() {
	// gin framework
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}
