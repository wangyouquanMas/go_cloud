package main

import (
	"filestore-server/service/dbproxy/mapper"
	"log"
)

func main() {
	res, err := mapper.FuncCall("/user/UserExist", []interface{}{"haha"}...)
	log.Printf("error: %+v\n", err)
	log.Printf("result: %+v\n", res[0].Interface())

	res, err = mapper.FuncCall("/user/UserExist", []interface{}{"admin"}...)
	log.Printf("error: %+v\n", err)
	log.Printf("result: %+v\n", res[0].Interface())
}
