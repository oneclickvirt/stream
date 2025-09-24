# STREAM Static Binaries

This directory contains statically compiled binaries of [STREAM](https://github.com/jeffhammond/STREAM) memory bandwidth benchmark for various platforms and architectures.

## About STREAM

The STREAM benchmark is a simple synthetic benchmark program that measures sustainable memory bandwidth (in MB/s) and the corresponding computation rate for simple vector kernels.

## Available Binaries

### Linux
- `stream-linux-amd64` - Linux x86_64 (optimized)
- `stream-linux-amd64-compat` - Linux x86_64 (maximum compatibility)
- `stream-linux-386` - Linux x86 (32-bit)
- `stream-linux-arm64` - Linux ARM64
- `stream-linux-armv7` - Linux ARMv7
- `stream-linux-armv6` - Linux ARMv6
- `stream-linux-riscv64` - Linux RISC-V 64-bit
- `stream-linux-ppc64le` - Linux PowerPC64 (little-endian)

### macOS
- `stream-darwin-amd64` - macOS x86_64 (macOS 10.12+)
- `stream-darwin-arm64` - macOS ARM64 (Apple Silicon, macOS 11.0+)

### Windows
- `stream-windows-amd64.exe` - Windows x86_64
- `stream-windows-386.exe` - Windows x86 (32-bit)
- `stream-windows-arm64.exe` - Windows ARM64

## Binary Variants

Some platforms may include both C and Fortran versions:
- `stream-*_c` - C version
- `stream-*_f` - Fortran version (if available)
- `stream-*` - Default version (usually C)

## Compatibility Notes

- **Linux binaries**: Statically linked for maximum portability (no dependencies required)
- **Linux amd64-compat**: Built without advanced CPU instructions for older processors
- **Windows binaries**: Statically linked, no additional runtime dependencies required
- **macOS binaries**: Dynamically linked with OpenMP support for better performance

## Usage

### Linux/macOS
```bash
# Make executable (if needed)
chmod +x stream-linux-amd64

# Run memory bandwidth benchmark
./stream-linux-amd64
```

### Windows
```cmd
# Run memory bandwidth benchmark
stream-windows-amd64.exe
```

## Understanding STREAM Results

STREAM measures four different vector operations:
- **Copy**: `a[i] = b[i]`
- **Scale**: `a[i] = q*b[i]`
- **Add**: `a[i] = b[i] + c[i]`
- **Triad**: `a[i] = b[i] + q*c[i]`

The results show the sustainable memory bandwidth in MB/s for each operation.

## Build Information

These binaries are automatically built using GitHub Actions from the [STREAM source repository](https://github.com/jeffhammond/STREAM).
All binaries are optimized for their respective platforms.

Built in: oneclickvirt/stream
