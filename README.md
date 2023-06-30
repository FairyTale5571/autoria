# Go AutoRIA

[![Go Reference](https://pkg.go.dev/badge/github.com/fairytale5571/autoria.svg)](https://pkg.go.dev/github.com/fairytale5571/autoria)
[![Go Report Card](https://goreportcard.com/badge/github.com/fairytale5571/autoria)](https://goreportcard.com/report/github.com/fairytale5571/autoria)


## Бібліотека для використання api auto.ria.com

### Для використання вам потрібен токен, який можна отримати на сайті [RIA Developers](https://developers.ria.com/)

### Використання

```go
package main

import (
	"fmt"
	"time"

	"github.com/fairytale5571/autoria"
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

### TODO:

 - [x] Типи транспорту
 - [x] Типи кузова
 - [x] Марки
 - [x] Моделі
 - [x] Покоління
 - [x] Модифікації
 - [x] Області
 - [x] Міста
 - [x] Типи приводу
 - [x] Типи палива
 - [x] Коробки передач
 - [x] Опції
 - [x] Кольори
 - [x] Країна виробник
 - [ ] Пошук оголошень
 - [ ] Підрахунок середньої ціни
 - [ ] Інформація по id оголошення