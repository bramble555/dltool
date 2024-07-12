package parameter

import (
	"flag"
	"fmt"
	"os"
)

// 返回url和dir路径
func Parse() (string, string) {
	// argsWithProg := os.Args

	// 输入的第一个参数是main.exe,第二个参数为ditool，说明用了此工具,但是如果第一个参数没有-，那么就失效了
	urlPtr := flag.String("url", "", "下载文件的地址")
	dirPtr := flag.String("dir", "./file/", "文件下载到本地的地址")
	_, err := os.Stat(*dirPtr)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(*dirPtr, 0755)
			if err != nil {
				fmt.Println("创建目录失败")
				os.Exit(0)
			}
		} else {
			panic(err)
		}
	}
	// 进行解析
	flag.Parse()
	return *urlPtr, *dirPtr

}
