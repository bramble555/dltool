package download

import (
	"dltool/parameter"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Do() {
	url, dir := parameter.Parse()
	if string(dir[len(dir)-1]) != `/` {
		dir = dir + `/`
	}
	// 创建目录
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Println("创建目录失败")
				os.Exit(0)
			}
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// 最后关闭响应体
	defer resp.Body.Close()
	// 获取下载到本地文件的名字
	fileName := filepath.Base(url) + ".png"
	file := dir + fileName
	// 创建文件，并且这个文件可读可写
	filePath, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	// 复制文件
	_, err = io.Copy(filePath, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("下载完成")

}
