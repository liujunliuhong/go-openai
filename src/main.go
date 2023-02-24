package main

import (
	"completions"
	"configuration"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	key := "sk-xxxx"
	c := configuration.NewConfiguration(key)

	p := completions.Param{
		Model:       "text-davinci-003",
		Prompt:      "今天是几号？",
		Temperature: 1,
		MaxTokens:   1024,
	}

	response, err := completions.Create(c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	marshal, err := json.Marshal(response)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	//
	//// list, err := models.List(c)
	//model, err := models.Retrieve(c, "babbage")
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(model)

}
