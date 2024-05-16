package handlers

type Handler struct {
	Ch *Chat
}

func NewHandler() *Handler {
	return &Handler{Ch: NewChat()}
}
