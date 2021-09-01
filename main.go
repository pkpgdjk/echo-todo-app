package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkpgdjk/echo-todo-app/router"
)

func main() {
	r := router.Init()
	data, err := json.MarshalIndent(r.Routes(), "", "  ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(data))
	//r.Run(fasthttp.New(":8888"))
	r.Logger.Fatal(r.Start("0.0.0.0:8080"))
}