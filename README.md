# Golang-learning

Just a repo for my personal GoLang learning and progression :)

- [Install Go here](https://go.dev/doc/install)

## Some useful GO CLI commands

```sh
- go version (prints current running Go version)

- go env (prints env vars)

- go fmt (formats code of all files in current directory)

- go help (list of all commands)

- go mod init example/test (creates a new module, initialising go.mod)

- go run . (compiles and executes your code)

- go build (compiles a bunch of go source code files)

- go install (compiles the program & "installs" a package)

- go mod tidy (remove unused dependencies)

- go get (downloads raw source code of a 3rd party package)

- go test (runs any tests associated with the current project)

- go list (lists all currnet module dependencies)
```

## managing & upgrading dependencies (example) 

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