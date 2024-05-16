package app

import (
	"fmt"
	"github.com/MaksKazantsev/Chattery/internal/config"
	"github.com/MaksKazantsev/Chattery/internal/handlers"
	"github.com/MaksKazantsev/Chattery/internal/log"
	"github.com/MaksKazantsev/Chattery/internal/routes"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(config *config.Config) {
	// Init logger
	l := log.NewLogger(config.Env)

	// New app
	app := fiber.New()

	// Init handler
	h := handlers.NewHandler()

	// Init routes
	routes.InitRoutes(h, app)

	// Running server
	shutdown(func() {
		l.Log.Info("Starting server")
		run(config.Port, app)
	})
	l.Log.Info("Server stopped")
}

func run(port string, app *fiber.App) {
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic("failed to listen to tcp " + err.Error())
	}
}
func shutdown(fn func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT|syscall.SIGTERM)
	go fn()
	<-stopCh
}
