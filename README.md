# Learning GO Lang

Download Go binary and extract it somewhere.
Set the following environment variables in your .bashrc or .zshrc file.

```bash
export PROGRAMS_PATH="/Users/joaopedroschmitt/programs"
export GOROOT="$PROGRAMS_PATH/golang/go-1.21.3" # Go installation home dir
export GOPATH="$PROGRAMS_PATH/golang/workspace" # 3rd-party dependencies home dir
export PATH="$PATH:$GOROOT/bin"
```

To init a module (usually the root of your git repository) run:

```bash
go mod init schmittjoaopedro/fundamentals
```

Run module:
    
```bash
go run .
```

Find Third-party dependencies/packages at: https://pkg.go.dev
Add new module requirements and sums

```bash
go mod tidy
```

Referencing a local module (not published in GitHub). Use the following syntax:

```bash
go mod edit -replace example.com/greetings=../greetings
go mod tidy
```

To run tests:

```bash
go test
go test -v # verbose
```
