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
