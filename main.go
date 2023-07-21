package main

import (
	"fmt"

	"github.com/haziqam/tarjans-algo/packages/webapi"
)

func main() {
	err := webapi.StartServer("5000")
	if (err != nil) {
		fmt.Println(err)
	}
}