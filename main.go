package main

import (
	"github.com/MaksKazantsev/Chattery/internal/app"
	"github.com/MaksKazantsev/Chattery/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
