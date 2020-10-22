package handle

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"repo/cache"

	"github.com/hyahm/xmux"
)

func InstallPackage(w http.ResponseWriter, r *http.Request) {
	// bw := xmux.NewWebsocket(w, r)
	// 获取目录， 其实就是获取文件包
	name := xmux.Var(r)["name"]
	dv := xmux.Var(r)["derivative"]
	file := xmux.Var(r)["file"]

	installPath := filepath.Join(cache.PackageRoot, dv, name, file)
	installByte, err := ioutil.ReadFile(installPath)
	if err != nil {
		w.Write([]byte("not found this package"))
		return
	}

	w.Write(installByte)
	return
}
