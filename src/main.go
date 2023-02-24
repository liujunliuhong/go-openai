package main

import (
	"configuration"
	"fmt"
	"models"
)

func main() {
	fmt.Println("Hello, World")

	c := configuration.NewConfiguration("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")

	// list, err := models.List(c)
	model, err := models.Retrieve(c, "babbage")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(model)

}
