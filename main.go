package main

import (
	"repo/handle"

	"github.com/hyahm/xmux"
)

// 用来上传下载包
func main() {
	router := xmux.NewRouter()
	router.Get("/list", handle.ListPackage)
	router.Post("/search/{derivative}/{name}", handle.SearchPackage)
	router.Get("/install/{derivative}/{name}/{file}", handle.InstallPackage)
	// 上传文件
	router.Post("/upload", handle.UploadPackage)
	router.Run()
}
