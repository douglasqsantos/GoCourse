# Go - Course

- [Go PlayGround](https://play.golang.org/)

## Go CLI

- go build
  - Compiles a bunch of go source code files
- go run
  - Compiles and executes one or two files
- go fmt
  - Formats all the code in each file in the current directory
- go install
  - Compiles and "installs" a package.
- go get
  - Download the raw source code of someone else's package
- go test
  - Runs any tests associates with the current package

## Basic Go Types

- bool
  - true, false
- string
  - "Hi!", "Hows it going?"
- int
  - 0, -1000, 9999
- float64
  - 10.00001, 0.00009, -100.002

## Pointer Operations

- &variable
  - Give me the memory address of the value this variable is pointing at
- *pointer
  - Give me the value this memory address is pointing at
- *person
  - This is a type description - it means we're working with a pointer to a person
- *pointerToPerson
  - This is an operator - it means we want to manipulate the value the pointer is referencing
- Rule
  - Turn **address** into **value** with *address
  - Turn **value** into **address** with &value

## Reference vs Value Types

- Value Types: Use pointers to change these things in a function
  - int
  - float
  - string
  - bool
  - structs
- Reference Types: Don't worry about pointers with these
  - slices
  - maps
  - channels
  - pointers
  - functions

## Initialize a new module

- https://linguinecode.com/post/how-to-import-local-files-packages-in-golang

```bash
cd project
go mod init github.com/douglasqantos/goproject
```

## Kubernetes

- https://github.com/kubernetes/client-go
- https://github.com/kubernetes/client-go/blob/master/INSTALL.md
- https://gist.github.com/ks888/0a0e0fbf4694d7955999a6f59aa2766d
- https://miminar.fedorapeople.org/_preview/openshift-enterprise/registry-redeploy/go_client/getting_started.html
- https://github.com/openshift/client-go
- https://github.com/openshift/api

## References

- https://medium.com/rungo/string-formatting-in-go-dd752ff7f60
- https://golangbyexample.com/remainder-modulus-go-golang/
- https://zetcode.com/all/#go
- github.com/spf13/cobra
