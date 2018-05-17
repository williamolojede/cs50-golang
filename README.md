# cs50-golang [WIP]

An implentation of harvard's cs50 library in go

## Installation
``` sh
  go get github.com/williamolojede/cs50-golang
```

## Usage
``` go
package main

import (
	"github.com/williamolojede/cs50-golang"
)

func main(){
  c := cs50.GetChar("Prompt: ")
  s := cs50.GetString("Prompt: ")
  i := cs50.GetInt("Prompt: ")
  f := cs50.GetFloat("Prompt: ")
}
```

## Todo
- [] Write tests
- [] Write Documentation
- [] Add default parameters to functions
