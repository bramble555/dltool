package download

import (
	"dltool/parameter"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Do() {
	url, dir := parameter.Parse()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// 最后关闭响应体
	defer resp.Body.Close()
	fileName := filepath.Base(url) + ".png"
	file := dir + fileName
	filePath, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	// 复制文件
	_, err = io.Copy(filePath, resp.Body)
	if err != nil {
		panic(err)
	}

}
