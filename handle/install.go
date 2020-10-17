package handle

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"repo/cache"

	"repo/config"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
	"gopkg.in/yaml.v3"
)

func InstallPackage(w http.ResponseWriter, r *http.Request) {
	bw := xmux.NewWebsocket(w, r)
	// 获取目录， 其实就是获取文件包
	name := xmux.Var(r)["name"]
	dv := os.Getenv("DERIVATIVE_VERSION")
	golog.Info(dv)
	root := filepath.Join(cache.PackageRoot, dv)
	fs, err := ioutil.ReadDir(root)
	if err != nil {
		bw.SendMessage([]byte("not support this system"), xmux.TypeMsg)
		return
	}

	for _, fi := range fs {
		if fi.IsDir() && fi.Name() == name {
			installPath := filepath.Join(cache.PackageRoot, dv, fi.Name(), name+".yaml")
			config := &config.Config{}
			installByte, _ := ioutil.ReadFile(installPath)
			golog.Info(string(installByte))
			err := yaml.Unmarshal(installByte, config)
			if err != nil {
				golog.Error(err)
				return
			}
			golog.Info(config)
			// 读取到配置文件后， 第一步，先读取depend， 获取depend里面的环境变量
			// 第二部， 添加自己的环境变量
			// 第三步，执行install.sh 脚本， websocket， 失败就返回， 否则进行第四步
			// 第四步， 生成script， 如果name是空的就简单了， 不用生成script
			// 如果需要生成script， 那么需要把第一步和第二步的环境变量都添加进script的env中
			// 替换 dir, command, env的值

		}
	}
	return

}
