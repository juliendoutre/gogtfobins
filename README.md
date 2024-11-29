# Go GTFO Bins

https://gtfobins.github.io/ lists Unix binaries that can be abused thanks to misconfigurations.

This project is a Golang CLI meant to be used as a discovery tool to evaluate what potential binaries could be abused in a Unix system to escalate privileges, open a reverse shell, _etc._

## Usage

```shell
# List all available binaries allowing for opening a reverse shell on the current host.
gogtfobins list --function reverse-shell
# Print details about the docker binary.
gogtfobins describe docker
```

## Development

### Lint the code

```shell
brew install golangci-lint
golangci-lint run
```
