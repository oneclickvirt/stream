# stream

一个嵌入stream依赖的golang库(A golang library with embedded stream dependencies)

## 关于 About

这个库提供了对STREAM内存带宽基准测试工具的Go语言封装。STREAM基准测试是一个简单的合成基准程序，用于测量可持续的内存带宽（MB/s）和简单向量内核的相应计算速率。

This library provides a Go wrapper for the STREAM memory bandwidth benchmark tool. The STREAM benchmark is a simple synthetic benchmark program that measures sustainable memory bandwidth (in MB/s) and the corresponding computation rate for simple vector kernels.

## 特性 Features

- 支持多平台：Linux (amd64, arm64, 386, arm等), macOS (arm64), Windows (amd64, 386, arm64)
- 自动检测并使用系统安装的stream，或使用嵌入的二进制文件
- 自动清理临时文件
- Multi-platform support: Linux (amd64, arm64, 386, arm, etc.), macOS (arm64), Windows (amd64, 386, arm64)
- Automatically detects and uses system-installed stream, or uses embedded binaries
- Automatic cleanup of temporary files

## 安装 Installation

```bash
go get github.com/oneclickvirt/stream@v0.0.1-20250924124003
```

## 使用方法 Usage

```go
package main

import (
    "log"
    "github.com/oneclickvirt/stream"
)

func main() {
    // 获取stream命令路径
    // Get stream command path
    streamCmd, tempFile, err := stream.GetStream()
    if err != nil {
        log.Fatalf("Failed to get stream: %v", err)
    }

    // 如果使用了临时文件，确保清理
    // Clean up temporary files if used
    if tempFile != "" {
        defer stream.CleanStream(tempFile)
    }

    // 执行stream基准测试
    // Execute stream benchmark
    err = stream.ExecuteStream(streamCmd, []string{})
    if err != nil {
        log.Fatalf("Failed to execute stream: %v", err)
    }
}
```

## API文档 API Documentation

### 函数 Functions

#### `GetStream() (string, string, error)`

获取可用的stream命令路径。

Get the available stream command path.

**返回值 Returns:**
- `string`: stream命令的路径 (Path to the stream command)
- `string`: 临时文件路径（如果使用了嵌入的二进制文件）(Temporary file path if using embedded binary)
- `error`: 错误信息 (Error information)

#### `ExecuteStream(streamPath string, args []string) error`

执行stream命令。

Execute the stream command.

**参数 Parameters:**
- `streamPath`: stream命令的路径 (Path to the stream command)
- `args`: 传递给stream的参数 (Arguments to pass to stream)

#### `CleanStream(tempFile string) error`

清理临时文件。

Clean up temporary files.

**参数 Parameters:**
- `tempFile`: 需要清理的临时文件路径 (Path to temporary file to clean up)

## 平台支持 Platform Support

库包含了以下平台的预编译二进制文件：

The library includes precompiled binaries for the following platforms:

### Linux
- amd64 (x86_64)
- 386 (x86 32-bit) 
- arm64 (ARMv8)
- arm (ARMv7)
- armv6 (ARMv6)
- riscv64 (RISC-V 64-bit)
- ppc64le (PowerPC64 little-endian)
- ppc64 (PowerPC64 big-endian)
- mips64le (MIPS64 little-endian)
- mips64 (MIPS64 big-endian)
- mipsle (MIPS little-endian)
- mips (MIPS big-endian)
- s390x (IBM System z)

### macOS
- arm64 (Apple Silicon)

### Windows
- amd64 (x86_64)
- 386 (x86 32-bit)
- arm64

## 许可证 License

MIT License 
