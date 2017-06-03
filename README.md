# gotrends
Unofficial wrapper for Google Trends using golang

## Using
### go get
You need `go` installed and `GOPATH` in your `PATH` , then run :
```shell
$ go get github.com/kahalKevin/gotrends
```

## Usage (Library)
```go
package main

import (
	"fmt"
	
	gotrends "github.com/kahalKevin/gotrends"
)


func main(){
	fmt.Println(gotrends.SearchWithKeyword("nike"))
}

```

## Output
Output will a string and int
```console
sepatu nike 100

```
