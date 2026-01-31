package main

import (
	"log"
	"time"

	"github.com/angelperez0709/markdown-app/internal/app"
	"github.com/angelperez0709/markdown-app/internal/config"
)

func main(){
	cfg := config.Config{
		Addr: ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
	}
	app := app.Application{
		Config: cfg,
	}
	handler := app.Mount()
	if err := app.Run(handler); err != nil {
		log.Printf("An error ocurred %s",err.Error())
		panic(err)
	}

}