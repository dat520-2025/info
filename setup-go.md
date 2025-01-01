# Installing Go

The most recent version of Go is 1.23.4, while Go 1.24 is expected in early February 2025.
If your system already comes with a version of Go, you may check your version like this:

```sh
go version
```

If not, please see install instructions on the [Go website](https://go.dev/doc/install).

## Special Instructions for 2025

Since the course starts before the release of Go 1.24, which we will depend on, you need will to install a pre-release version of Go.

### Option 1: Installing Go 1.24 release candidate

If you have Go installed already, an easy way to try go1.24rc1 is by using the go command:

```sh
go install golang.org/dl/go1.24rc1@latest
go1.24rc1 download
```

Alternatively, you can download binary and source distributions from the usual place:
[https://go.dev/dl/#go1.24rc1](https://go.dev/dl/#go1.24rc1)

To find out what has changed in Go 1.24, read the [draft release notes](https://tip.golang.org/doc/go1.24).

Check your Go version again:

```sh
go version
```

### Option 2: Installing Go tip

Alternatively, you can install the latest development version of Go, which include a few more changes than the release candidate.
You can install the latest development version, by using the `gotip` command:

```sh
go install golang.org/dl/gotip@latest
gotip download
```

Check your Go version again:

```sh
% go version
go version devel go1.24-485ed2fa Tue Dec 3 00:06:52 2024 +0000 darwin/arm64
```

### Troubleshooting: Finding the Right go Command After Installation

After installing Go 1.24 development version or release candidate, you may need to adjust your `PATH` environment variable to use the new version of the `go` command.

On macOS, I have the following in my `~/.zshrc` file:

```sh
# Make gotip the default Go command;
# gotip is installed in ~/go/bin and ~/sdk/gotip/bin contains the gotip distribution.
[[ -x $(command -v gotip) ]] && export PATH=$(gotip env GOROOT)/bin:$PATH
```

For me this results in:

```sh
% env | grep PATH
PATH=/Users/meling/sdk/gotip/bin:/opt/homebrew/bin:...
```

If the above instructions do not work for you, please ask the TAs for help.

## Getting Started

To get started with Go, you should read the [Getting Started](https://go.dev/doc/tutorial/getting-started) tutorial on the Go website.
