package routes

import (
	"github.com/MaksKazantsev/Chattery/internal/handlers"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(h *handlers.Handler, a *fiber.App) {
	ch := a.Group("/chats")
	ch.Get("/join", websocket.New(h.Ch.Join))
}
