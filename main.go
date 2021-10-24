package main

import (
	"gitlab.com/Dank-del/SibylAPI-Go/server"
	"log"
)

func main() {
	r := server.SibylServer()
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
