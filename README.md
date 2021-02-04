# LibreTranslate in golang

### Install:
```
go get github.com/SnakeSel/libretranslate
```

### Example usage:

```go
package main

import (
    "fmt"
    tr "github.com/snakesel/libretranslate"
)

func main() {
    const text string = `Hello, World!`

    // you can use "auto" for source language
    // so, translator will detect language
    trtext, err := tr.Translate(text, "auto", "ru")
    if err == nil {
        fmt.Println(trtext)
    } else {
        fmt.Println(err.Error())
    }
}
```



