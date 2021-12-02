# Golang-learning

Just a repo for my personal GoLang learning and progression :)

- [Install Go here](https://go.dev/doc/install)


### some useful GO CLI commands:

```sh
- go version (current running Go version)

- go env (prints env vars)

- go fmt (formats code)

- go help (list of all commands)

- go mod init example/test (creates a new module, initialising go.mod)

- go run . (runs your code in current directory)

- go build (builds and creates an executable or package)

- go install (compiles the program & builds it)

- go mod tidy (remove unused dependencies)

- go get (changes required version of of a dependency or adds a new dependency)

- go test

- go list (lists all currnet module dependencies)
```
### managing & upgrading dependencies (example)

```sh
- go get goloang.org/x/text (example)
- go list -m all (list all dependencies)
- go get rsc.io/sampler
- go list -m -version rsc.io/sampler (search for versions)
- go get rsc.io/sampler@v1.3.1
```
## links & docs

- [golang.org](golang.org)

- [godoc.org](godoc.org)