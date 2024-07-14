package download

import (
	"fmt"
	"io"
	"net/http"
)

type RaederProgress struct {
	io.Reader
	current float64
	total   float64
}

// 总大小转换单位
func sizeFormat(n int) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	switch {
	case n > GB:
		return fmt.Sprintf("%.2fGB", float64(n/GB))
	case n > MB:
		return fmt.Sprintf("%.2fMB", float64(n/MB))
	case n > KB:
		return fmt.Sprintf("%.2fKB", float64(n/KB))
	default:
		return fmt.Sprintf("%.2fB", float64(n))
	}

}
func (rp *RaederProgress) Read(p []byte) (n int, err error) {
	n, err = rp.Reader.Read(p)
	rp.current += float64(n)
	fmt.Printf("\r当前进度为%.2f%%=============文件大小写%s", rp.current*100/rp.total, sizeFormat(int(rp.total)))
	return

}
func progress(dst io.Writer, r *http.Response) {
	progress := &RaederProgress{
		Reader: r.Body,
		total:  float64(r.ContentLength),
	}
	_, err := io.Copy(dst, progress)
	if err != nil {
		panic(err)
	}
	if progress.current == progress.total {
		fmt.Println("\n下载完成")
	}

}
