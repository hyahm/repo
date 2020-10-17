package handle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"repo/cache"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchPackage(w http.ResponseWriter, r *http.Request) {
	// 获取目录， 其实就是获取文件包
	name := xmux.Var(r)["name"]
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
		if fi.IsDir() && strings.Contains(fi.Name(), name) {
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
