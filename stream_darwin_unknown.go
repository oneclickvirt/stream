//go:build darwin && !arm64
// +build darwin,!arm64

package stream

import "fmt"

func GetStream() (streamCmd string, tempFile string, err error) {
	return "", "", fmt.Errorf("stream binary not available for this platform")
}

func ExecuteStream(streamPath string, args []string) error {
	return fmt.Errorf("stream binary not available for this platform")
}