package main

import (
	"log"

	"github.com/nikitasmall/gonews/router"
)

func main() {
	r := router.New()

	log.Println("The server listen on :3000 port!")
	r.Run(":3000")
}
