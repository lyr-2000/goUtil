package util

import (
	"log"
	"os"
	"strings"
)

//获取当前路径，比如：E:/abc/data/test

func GetCurrentProjectPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}