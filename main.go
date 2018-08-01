package main

import (
	"fmt"

	"github.com/joy/server"
)

func main() {
	var app server.App
	fmt.Println(app.Run())
}
