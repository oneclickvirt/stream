//go:build linux && !386 && !amd64 && !arm && !arm64 && !armv6 && !mips && !mipsle && !mips64 && !mips64le && !ppc64 && !ppc64le && !riscv64 && !s390x
// +build linux,!386,!amd64,!arm,!arm64,!armv6,!mips,!mipsle,!mips64,!mips64le,!ppc64,!ppc64le,!riscv64,!s390x

package stream

import "fmt"

func GetStream() (streamCmd string, tempFile string, err error) {
	return "", "", fmt.Errorf("stream binary not available for this platform")
}

func ExecuteStream(streamPath string, args []string) error {
	return fmt.Errorf("stream binary not available for this platform")
}