package stream

import (
	"os"
	"path/filepath"
)

// CleanStream 删除临时提取出的 stream 文件
func CleanStream(tempFile string) error {
	if tempFile == "" {
		return nil // 不需要清理
	}
	// 删除整个临时目录
	return os.RemoveAll(filepath.Dir(tempFile))
}