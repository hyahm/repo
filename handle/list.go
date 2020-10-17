package handle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"repo/cache"

	"github.com/hyahm/golog"
)

type PackageInfo struct {
	// 包信息
	Name string `json:"name"`
	Info string `json:"info"`
}

func ListPackage(w http.ResponseWriter, r *http.Request) {
	// 获取目录， 其实就是获取文件包
	// scs list
	dv := os.Getenv("DERIVATIVE_VERSION")
	golog.Info(dv)
	root := filepath.Join(cache.PackageRoot, dv)
	fs, err := ioutil.ReadDir(root)
	if err != nil {
		w.Write([]byte("not support this system"))
		return
	}
	packages := make([]*PackageInfo, 0)
	for _, fi := range fs {
		if fi.IsDir() {
			describePath := filepath.Join(cache.PackageRoot, dv, fi.Name(), "describe.txt")
			db, _ := ioutil.ReadFile(describePath)
			pi := &PackageInfo{
				Name: fi.Name(),
				Info: string(db),
			}
			packages = append(packages, pi)
		}
	}
	send, _ := json.Marshal(packages)
	w.Write(send)
	return
}
