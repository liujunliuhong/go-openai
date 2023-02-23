package main

import (
	"fmt"
	"models"
	"request"
)

func main() {
	fmt.Println("Hello, World")

	c := request.NewConfiguration(Key)

	list, err := models.List(c)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(list.Models)

}
