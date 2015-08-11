package main

import (
	"fmt"

	"github.com/gjohnson/haikunator"
)

func main() {
	name, err := haikunator.Haikunate()
	if err != nil {
		panic(err)
	}
	fmt.Printf(name)
}
