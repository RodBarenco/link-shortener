package main

import (
	"github.com/RodBarenco/link-shortener/model"
	"github.com/RodBarenco/link-shortener/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
