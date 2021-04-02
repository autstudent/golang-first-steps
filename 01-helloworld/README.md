# Golang Essentials

## Commands

- go build -> Just compile

```$bash
go build ./main.go
```

- go run -> Compile and Run a go file

```$bash
go run main.go
```

- go fmt -> Format all the code

- go install -> Compile and install a package

- go get -> Downloads the raw source code of someone else's packages

- go test -> Run and execute test files

## Packages

Package === Project === Workspace

- Package -> It can have multiple go files and all files have package imported at the top of the file

```$bash
package main
...
```

In Golang exists two king of packages:

- **Executable** generates a file that we can run (main package always is executable)
- **Reusable** contains code used as *helpers* or stuff to help us with some code (Any other package is reusable)

## Imports

It allow import specific code from other reusable packages, named normally libraries (golang.org/pkg). Some important packages:

- fmt -> Standard lib
- io
- cryto
- debug
- enconding
- math

```$bash
...
import "fmt"
...
```

## Functions

The following list includes the parts of a function definition:

- func -> Declare a function
- main -> Name of the function
- () -> List of arguments to pass the function
- { xx } -> Function body

```$bash
func main() {
    fmt.Println("Hola")
}
```

## Go file organization

- package declaration
- import other packages
- declare functions
