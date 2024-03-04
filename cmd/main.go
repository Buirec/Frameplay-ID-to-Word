package main

import (
	"fmt"

	"github.com/buirec/Frameplay-ID-to-Word/server"
)

func main() {
	id := server.GetId("data.json")

	if id != "" {
		fmt.Println(server.PostId(id))
	}
}
