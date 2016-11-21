package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func main() {
	res := Response{
		Code:   200,
		Result: "Hello Charlie",
	}
	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
