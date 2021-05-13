# LibreTranslate in golang
[LibreTranslate](https://libretranslate.com) is an Open Source Machine Translation  
[API Docs](https://libretranslate.com/docs) | [Self-Hosted](https://github.com/uav4geo/LibreTranslate)

### Install:
```
go get -u github.com/snakesel/libretranslate
```

### Example usage:

```go
package main

import (
    "fmt"
    tr "github.com/snakesel/libretranslate"
)

func main() {
    translate := tr.New(tr.Config{
        Url:   "https://libretranslate.com",
        Key:   "XXX",
    })

    // you can use "auto" for source language
    // so, translator will detect language
    trtext, err := translate.Translate("Hello, World!", "auto", "ru")
    if err == nil {
        fmt.Println(trtext)
    } else {
        fmt.Println(err.Error())
    }

    // Detect the language of the text
    conf, lang, err := translate.Detect("Nächster Stil")
    if err == nil {
        fmt.Printf("%s (%f)\n", lang, conf)
    } else {
        fmt.Println(err.Error())
    }
}
```



