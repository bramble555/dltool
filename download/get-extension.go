package download

func getExtension(contentType string) string {
	// MIME类型到文件扩展名的映射
	mimeName := map[string]string{
		"image/avif":                     ".avif",
		"text/css; charset=utf-8":        ".css",
		"image/gif":                      ".gif",
		"text/html; charset=utf-8":       ".htm",
		"text/html; charset=utf-8l":      ".htm",
		"image/jpeg":                     ".jpeg",
		"text/javascript; charset=utf-8": ".js",
		"application/json":               ".json",
		"application/pdf":                ".pdf",
		"image/png":                      ".png",
		"image/svg+xml":                  ".svg",
		"application/wasm":               ".wasm",
		"image/webp":                     ".webp",
		"text/xml; charset=utf-8":        ".xml",
		"text/plain":                     ".txt",
		"video/mp4":                      "mp4",
	}
	if v, ok := mimeName[contentType]; ok {
		return v
	}
	return ""
}
