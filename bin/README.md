# STREAM Static Binaries

This directory contains statically compiled binaries of [STREAM](https://github.com/jeffhammond/STREAM) memory bandwidth benchmark for various platforms and architectures.

## About STREAM

The STREAM benchmark is a simple synthetic benchmark program that measures sustainable memory bandwidth (in MB/s) and the corresponding computation rate for simple vector kernels.

## Available Binaries

### Linux (Complete Architecture Support)
- `stream-linux-amd64` - Linux x86_64 (optimized)
- `stream-linux-amd64-compat` - Linux x86_64 (maximum compatibility)
- `stream-linux-386` - Linux x86 (32-bit)
- `stream-linux-arm64` - Linux ARM64 (ARMv8)
- `stream-linux-armv7` - Linux ARMv7
- `stream-linux-armv6` - Linux ARMv6
- `stream-linux-riscv64` - Linux RISC-V 64-bit
- `stream-linux-ppc64le` - Linux PowerPC64 (little-endian)
- `stream-linux-ppc64` - Linux PowerPC64 (big-endian)
- `stream-linux-s390x` - Linux IBM System z (s390x)
- `stream-linux-mips` - Linux MIPS (32-bit, big-endian)
- `stream-linux-mipsle` - Linux MIPS (32-bit, little-endian)
- `stream-linux-mips64` - Linux MIPS64 (big-endian)
- `stream-linux-mips64le` - Linux MIPS64 (little-endian)

### macOS
- `stream-darwin-amd64` - macOS x86_64 (macOS 10.12+)
- `stream-darwin-arm64` - macOS ARM64 (Apple Silicon, macOS 11.0+)

### Windows
- `stream-windows-amd64.exe` - Windows x86_64
- `stream-windows-386.exe` - Windows x86 (32-bit)
- `stream-windows-arm64.exe` - Windows ARM64

## Architecture Coverage

This build now supports **ALL** Go-supported Linux architectures:
- **x86**: amd64, 386
- **ARM**: arm64, armv7, armv6
- **RISC-V**: riscv64
- **PowerPC**: ppc64, ppc64le
- **IBM Z**: s390x
- **MIPS**: mips, mipsle, mips64, mips64le

## Binary Variants

Some platforms may include both C and Fortran versions:
- `stream-*_c` - C version
- `stream-*_f` - Fortran version (if available)
- `stream-*` - Default version (usually C)

## Compatibility Notes

- **Linux binaries**: Statically linked for maximum portability (no dependencies required)
- **Linux amd64-compat**: Built without advanced CPU instructions for older processors
- **Exotic architectures** (MIPS, s390x, etc.): May have reduced optimization due to cross-compilation constraints
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

## Architecture-Specific Notes

### Mainstream Architectures
- **amd64/386**: Fully optimized with SSE/AVX support (amd64) or i686 compatibility (386)
- **arm64**: ARMv8-A optimizations enabled
- **armv7/armv6**: Hardware floating-point support where available

### Exotic Architectures
- **MIPS variants**: May require QEMU emulation during build, basic optimization only
- **PowerPC**: Both big-endian (ppc64) and little-endian (ppc64le) support
- **s390x**: IBM Z mainframe architecture support
- **RISC-V**: Modern RISC-V 64-bit support

### Build Quality
- **Tier 1**: amd64, 386, arm64, armv7 - Fully tested and optimized
- **Tier 2**: ppc64le, riscv64, armv6 - Good support, less testing
- **Tier 3**: mips*, ppc64, s390x - Basic support, may have compilation issues

## Troubleshooting

### If a binary doesn't work:
1. Check architecture: `uname -m` on Linux/macOS
2. Try the compatibility version: `stream-linux-amd64-compat`
3. Check for missing dependencies: `ldd stream-linux-*` (should show "statically linked")
4. Verify executable permissions: `chmod +x stream-linux-*`

### Performance Notes:
- Results vary significantly between architectures
- Memory bandwidth is often the limiting factor, not CPU performance
- Static linking may slightly reduce performance vs. dynamic linking
- Some exotic architectures may show lower performance due to cross-compilation limitations

## Build Information

These binaries are automatically built using GitHub Actions from the [STREAM source repository](https://github.com/jeffhammond/STREAM).
All binaries are optimized for their respective platforms with fallback compilation strategies for exotic architectures.

**Cross-compilation toolchains used:**
- GNU GCC cross-compilers for most Linux architectures
- LLVM-MinGW for Windows ARM64
- Native Clang/GCC for macOS
- QEMU emulation for MIPS architectures during build

Built in: oneclickvirt/stream
Covers all Go-supported Linux architectures as of 2024.
