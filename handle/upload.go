package handle

import (
	"archive/tar"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"repo/cache"
	"strings"

	"github.com/hyahm/golog"
)

func UploadPackage(w http.ResponseWriter, r *http.Request) {
	// bw := xmux.NewWebsocket(w, r)
	// 获取目录， 其实就是获取文件包
	// 获取文件
	body, header, err := r.FormFile("file")
	if err != nil {
		golog.Error(err)
		return
	}
	defer body.Close()
	overwrite := r.FormValue("overwrite")
	golog.Info(overwrite)
	osversion := r.FormValue("osversion")
	osversionList := make([]string, 0)
	err = json.Unmarshal([]byte(osversion), &osversionList)
	if err != nil {
		golog.Error(err)
		return
	}

	suffixName := header.Filename[:strings.LastIndex(header.Filename, ".")]
	for _, osDir := range osversionList {
		// 判断文件是否存在
		packageName := filepath.Join(cache.PackageRoot, osDir, suffixName)
		_, err := os.Stat(packageName)
		if err != nil {
			if os.IsExist(err) {
				// 如果覆盖
				if overwrite == "false" {
					golog.Error(err)
					continue
				}
				os.RemoveAll(packageName)

			} else {
				os.MkdirAll(packageName, 0755)
			}

		}

		tr := tar.NewReader(body)
		for {
			hdr, err := tr.Next()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					golog.Error(err)
					return
				}
			}

			filename := filepath.Join(cache.PackageRoot, osDir, suffixName, hdr.Name)

			fi := hdr.FileInfo()
			if fi.IsDir() {
				os.MkdirAll(filename, 0755)
			} else {
				file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					golog.Error(err)
					return
				}
				io.Copy(file, tr)
				file.Close()
			}

		}
	}

	// w.Write(installByte)
	return
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
