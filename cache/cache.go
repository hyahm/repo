package cache

import (
	"os"
	"path/filepath"
)

var PackageRoot = "pkg"
var OSS = []string{"ubuntu", "centos"}

// 假如先建立2个目录， ubuntu， centos
func init() {
	// 自动生成目录， 如果不存在就自动生成
	for _, path := range OSS {
		pr := filepath.Join(PackageRoot, path)
		err := os.MkdirAll(pr, 0755)
		if err == os.ErrPermission {
			panic(err)
		}
	}

}
