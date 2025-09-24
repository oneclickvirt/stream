//go:build windows && !arm64 && !amd64 && !386
// +build windows,!arm64,!amd64,!386

package stream

import "fmt"

func GetStream() (streamCmd string, tempFile string, err error) {
	return "", "", fmt.Errorf("stream binary not available for this platform")
}

func ExecuteStream(streamPath string, args []string) error {
	return fmt.Errorf("stream binary not available for this platform")
}