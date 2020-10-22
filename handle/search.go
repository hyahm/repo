package handle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"repo/cache"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchPackage(w http.ResponseWriter, r *http.Request) {
	// 获取目录， 其实就是获取文件包
	name := xmux.Var(r)["name"]
	derivative := xmux.Var(r)["derivative"]
	golog.Info(derivative)
	root := filepath.Join(cache.PackageRoot, derivative)
	packages := make([]*PackageInfo, 0)
	fs, err := ioutil.ReadDir(root)
	if err == nil {
		for _, fi := range fs {
			if fi.IsDir() && strings.Contains(fi.Name(), name) {
				describePath := filepath.Join(cache.PackageRoot, derivative, fi.Name(), "describe.txt")
				db, _ := ioutil.ReadFile(describePath)
				pi := &PackageInfo{
					Name: fi.Name(),
					Info: string(db),
				}
				packages = append(packages, pi)
			}
		}
	}

	send, _ := json.Marshal(packages)
	golog.Info(string(send))
	w.Write(send)
	return
}
