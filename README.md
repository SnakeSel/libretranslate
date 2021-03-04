# LibreTranslate in golang
[LibreTranslate](https://libretranslate.com) is an Open Source Machine Translation  
[API Docs](https://libretranslate.com/docs) | [Self-Hosted](https://github.com/uav4geo/LibreTranslate)

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

    // Detect the language of the text
    const textDe:="Nächster Stil"
    conf,lang, err := tr.Detect(textDe)
    if err == nil {
        fmt.Printf("%s (%f)", lang, conf)
    } else {
        fmt.Println(err.Error())
    }
}
```



