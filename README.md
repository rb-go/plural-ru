# plural-ru

[![Go Reference](https://pkg.go.dev/badge/github.com/rb-go/plural-ru.svg)](https://pkg.go.dev/github.com/rb-go/plural-ru)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/rb-go/plural-ru)](https://github.com/rb-go/plural-ru/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rb-go/plural-ru)](https://github.com/rb-go/plural-ru)
[![Go Report Card](https://goreportcard.com/badge/github.com/rb-go/plural-ru)](https://goreportcard.com/report/github.com/rb-go/plural-ru)
[![Coverage Status](https://coveralls.io/repos/github/rb-go/plural-ru/badge.svg)](https://coveralls.io/github/rb-go/plural-ru)
[![license](https://img.shields.io/github/license/rb-go/plural-ru.svg)](LICENSE)

Package Pluralization of russian words for golang

## Usage

```go
package main

import (
    "log"

    "github.com/rb-go/plural-ru"
)

func main() {
    log.Println(1, plural.Noun(1, "день", "дня", "дней"))
    log.Println(2, plural.Noun(2, "день", "дня", "дней"))
    log.Println(3, plural.Noun(3, "день", "дня", "дней"))
    log.Println(4, plural.Noun(4, "день", "дня", "дней"))
    log.Println(5, plural.Noun(5, "день", "дня", "дней"))
    log.Println(8, plural.Noun(8, "день", "дня", "дней"))
    log.Println(11, plural.Noun(11, "день", "дня", "дней"))
    log.Println(1111, plural.Noun(1111, "день", "дня", "дней"))
    log.Println(121, plural.Noun(121, "день", "дня", "дней"))

    cnt := 100500
    log.Println(plural.Verb(cnt, "нашелся", "нашлись", "нашлось"), cnt, plural.Noun(cnt, "день", "дня", "дней"))
}
```

will print in console:

```
2022/10/28 17:27:09 1 день
2022/10/28 17:27:09 2 дня
2022/10/28 17:27:09 3 дня
2022/10/28 17:27:09 4 дня
2022/10/28 17:27:09 5 дней
2022/10/28 17:27:09 8 дней
2022/10/28 17:27:09 11 дней
2022/10/28 17:27:09 1111 дней
2022/10/28 17:27:09 121 день
2022/10/28 17:27:09 нашлось 100500 идей
```
