package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Update(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	// reset cursor
	file.Seek(0, 0)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	// reset cursor
	file.Seek(0, 0)
	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// Hex2Dec: 十六进制转十进制
func Hex2Dec(val string) int64 {
	n, err := strconv.ParseInt(val, 16, 0)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

// Contain : 判断某个元素是否在 slice,array ,map中
func Contain(target interface{}, obj interface{}) (bool, error) {
	targetVal := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		// 是否在slice/array中
		for i := 0; i < targetVal.Len(); i++ {
			if targetVal.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		// 是否在map key中
		if targetVal.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	default:
		fmt.Println(reflect.TypeOf(target).Kind())
	}

	return false, errors.New("not in this array/slice/map")
}
