package main

import (
	"github.com/HunterGooD/Chat-Go/internal/app"
)

func main() {
	app := app.NewApp()
	app.Init()
	app.RunServer()
}
