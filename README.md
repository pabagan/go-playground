# Golang Playground

This [Git](https://github.com/pabagan/go-playground) is a project to work on demo code while learning Go.

## Workflow

Develop new recipes at `main.go`. Once achieved move the file to `/recipes` folder. Export a function `Test*` from each recipe.

## Setting project

### Create Go Module
```sh
# create go module for https://github.com/pabagan/go-playground
go mod init github.com/pabagan/go-playground
```

### Create 1st recipe

```go
package main

import (
	"log"
	"recipes"
)

func main() {
	recipes.Print("hola")
}
```

##### Run
```sh
go run main.go
```

## Execute script

Run: 
```sh
# development
run go main.go

# recipes
run go recipes/xxx.go
```

