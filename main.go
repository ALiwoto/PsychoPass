package main

import (
	"gitlab.com/Dank-del/SibylAPI-Go/config"
	"gitlab.com/Dank-del/SibylAPI-Go/routes"
	"gitlab.com/Dank-del/SibylAPI-Go/server"
	"log"
)

func main() {
	c := new(config.ServerConfig)
	r := server.SibylServer(c)
	r.GET("createToken", routes.CreateToken)
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
