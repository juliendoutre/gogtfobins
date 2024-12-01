# Go GTFO Bins

https://gtfobins.github.io/ lists Unix binaries that can be abused if misconfigured.

This project is a Golang CLI meant to be used as a discovery tool to evaluate what potential binaries could be abused in a Unix system to escalate privileges, open a reverse shell, _etc._

## Getting started

Checkout out the project releases on https://github.com/juliendoutre/gogtfobins/releases!

## Usage

```shell
# List all available binaries allowing for opening a reverse shell on the current host.
gogtfobins list --function reverse-shell
# Print possible exploits for the docker binary as JSON.
gogtfobins describe docker --format json
```

## Development

### Lint the code

```shell
brew install golangci-lint
golangci-lint run
```

### Release a new version

```shell
brew install goreleaser
goreleaser check syft
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
rm -rf ./dist
GITHUB_TOKEN=$(gh auth token) goreleaser release
```
