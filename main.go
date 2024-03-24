package main

import (
	"practice/router"
)

func main() {
	// config.InitConfig()
	s := router.NewServer()
	s.Run()
}