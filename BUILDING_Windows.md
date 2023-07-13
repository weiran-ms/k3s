# Building for Windows

K3S for Windows cannot be built with Golang cross-platform compilation on Linux or Windows WSL, becuase some of its dependencies, e.g. [go-sqlite3](https://github.com/mattn/go-sqlite3), requires [CGO](https://pkg.go.dev/cmd/cgo), which must be built from the target platform.

Instead, K3S for Windows can be built with MSYS2 MINGW64 on Windows.
## Set up build environment for Windows
1. Install `MSYS2` (https://www.msys2.org) and open `MSYS2 MINGW64 Terminal`.
2. Install `Git`
```bash
pacman -S git
```
3. Install `Golang`
```bash
pacman -S mingw-w64-x86_64-go
```
4. Install `GCC` toolchain for [CGO](https://pkg.go.dev/cmd/cgo)
```bash
pacman -S mingw-w64-x86_64-gcc
```
## Build k3s from source

Before getting started, bear in mind that this repository includes all of Kubernetes history, so consider shallow cloning with (`--depth 1`) to speed up the process.

```bash
git clone --depth 1 -b <branch> https://github.com/k3s-io/k3s.git
```

Build the K3S executable and its dependencies:

```bash
 ./scripts/build_windows
```
Package all K3S binaries into a single executable, which will create `./dist/artifacts/k3s.exe`.

```bash
./scripts/package-cli_windows
```