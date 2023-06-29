# Go AutoRIA

[![Go Reference](https://pkg.go.dev/badge/github.com/bwmarrin/discordgo.svg)](https://pkg.go.dev/github.com/bwmarrin/discordgo) 
[![Go Report Card](https://goreportcard.com/badge/github.com/bwmarrin/discordgo)](https://goreportcard.com/report/github.com/bwmarrin/discordgo) 
[![CI](https://github.com/bwmarrin/discordgo/actions/workflows/ci.yml/badge.svg)](https://github.com/bwmarrin/discordgo/actions/workflows/ci.yml)


## Бібліотека для використання api auto.ria.com

### Для використання вам потрібен токен, який можна отримати на сайті [RIA Developers](https://developers.ria.com/)

### Використання

```go
package main

import (
	"fmt"
	"time"

	"github.com/fairytale5571/go-autoria"
)

const (
	APIKey = "123"
)

func main() {
	ria := autoria.New(autoria.Opts{
		APIKey:  APIKey,
		Timeout: 10 * time.Second,
	})
	cars, err := ria.GetMarksByCategory(autoria.CategoryCars)
	if err != nil {
		panic(err)
	}
	for _, car := range cars {
		fmt.Printf("%s: %d\n", car.Name, car.Value)
	}
}

```