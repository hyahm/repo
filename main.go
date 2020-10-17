package main

import (
	"repo/handle"

	"github.com/hyahm/xmux"
)

// 用来上传下载包
func main() {

	router := xmux.NewRouter()
	router.Get("/list", handle.ListPackage)
	router.Post("/search/{name}", handle.SearchPackage)
	router.Get("/install/{name}", handle.InstallPackage)
	router.Run()
}
