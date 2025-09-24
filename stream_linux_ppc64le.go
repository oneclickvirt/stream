//go:build linux && ppc64le
// +build linux,ppc64le

package stream

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed bin/stream-linux-ppc64le
var binFiles embed.FS

func GetStream() (streamCmd string, tempFile string, err error) {
	var errors []string
	// 首先尝试系统中已安装的 stream
	if path, lookErr := exec.LookPath("stream"); lookErr == nil {
		output, runErr := exec.Command(path).CombinedOutput()
		if runErr == nil || strings.Contains(string(output), "STREAM") {
			return "stream", "", nil
		} else {
			errors = append(errors, fmt.Sprintf("stream 直接运行失败: %v\n输出: %s", runErr, string(output)))
		}
	} else {
		errors = append(errors, fmt.Sprintf("无法找到 stream: %v", lookErr))
	}
	
	// 创建临时目录来存放嵌入的二进制文件
	tempDir, tempErr := os.MkdirTemp("", "streamwrapper")
	if tempErr != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", tempErr)
	}
	
	binName := "stream-linux-ppc64le"
	binPath := filepath.Join("bin", binName)
	fileContent, readErr := binFiles.ReadFile(binPath)
	if readErr == nil {
		tempFile = filepath.Join(tempDir, binName)
		writeErr := os.WriteFile(tempFile, fileContent, 0755)
		if writeErr == nil {
			// 测试嵌入的二进制文件是否可运行
			output, runErr := exec.Command(tempFile).CombinedOutput()
			if runErr == nil || strings.Contains(string(output), "STREAM") {
				return tempFile, tempFile, nil
			} else {
				errors = append(errors, fmt.Sprintf("%s 运行失败: %v\n输出: %s", tempFile, runErr, string(output)))
			}
		} else {
			errors = append(errors, fmt.Sprintf("写入临时文件失败 (%s): %v", tempFile, writeErr))
		}
	} else {
		errors = append(errors, fmt.Sprintf("读取嵌入的 stream 版本失败: %v", readErr))
	}
	
	return "", "", fmt.Errorf("无法找到可用的 stream 命令:\n%s", strings.Join(errors, "\n"))
}

func ExecuteStream(streamPath string, args []string) error {
	var cmd *exec.Cmd
	if streamPath == "stream" {
		cmd = exec.Command(streamPath, args...)
	} else {
		allArgs := append([]string{streamPath}, args...)
		cmd = exec.Command("sh", "-c", strings.Join(allArgs, " "))
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}