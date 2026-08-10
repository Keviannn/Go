## Code organization
### Packages
Go programs are organized into packages which are a collection of source files in the same directory that are compiled together.

### Repositories
A module is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. 

A file named go.mod there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its go.mod file as well as subdirectories of that directory, up to the next subdirectory containing another go.mod file (if any).

## Creating a module
Generating a new module: `go mod init example/user/mod`

Building and installing: `go install example/user/mod` (`go install .` or `go install` also work if in the same directory).

## Environment
The install directory is controlled by the GOPATH and GOBIN environment variables. 

If GOBIN is set, binaries are installed to that directory. 

If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. 

Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).

Change the GOBIN: `go env -w GOBIN=/somewhere/else/bin`

Unset changes: `go env -u GOBIN`

## Executing
``` shell
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ mod
Hello, world.
```

## Using git
```shell
$ git init
Initialized empty Git repository in /home/user/hello/.git/
$ git add go.mod hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 go.mod hello.go
$
```
The go command locates the repository containing a given module path by requesting a corresponding HTTPS URL and reading metadata embedded in the HTML response (see go help importpath). 

Many hosting services already provide that metadata for repositories containing Go code, so the easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.

> In order first git init, then go mod with the path.

## Importing packages from your module
Creating a new directory, adding different functions such as morestrings.

Execute `go build`.

Then it can be imported inside as in the example: `	"example/user/mod/morestrings"`

Then it can be used as: `fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))`

## Importing packages from remote modules
```go 
import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)
```

Now that you have a dependency on an external module, you need to download that module and record its version in your go.mod file. 

The go mod tidy command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore: `go mod tidy`

Example:
```shell
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
-     "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```
    
To remove all downloaded modules, you can pass the -modcache flag to go clean: `go clean -modcache`

## Testing
Go has a lightweight test framework composed of the go test command and the testing package.

You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). 

The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.

Then executing the `go test` command:
```shell
$ go test
PASS
ok  	example/user/hello/morestrings 0.165s
$
```

