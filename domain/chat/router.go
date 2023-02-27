package chat

import (
	"github.com/gin-gonic/gin"
)

type ChatRouter struct {
	handler *ChatHandler
}

func NewChatRouter(chatHandler *ChatHandler) *ChatRouter {
	return &ChatRouter{
		handler: chatHandler,
	}
}

func (cr *ChatRouter) Init(router *gin.Engine) {
	router.POST("/api/chat/v1", cr.handler.HandleChatV1)
	router.POST("/api/chat/v2", cr.handler.HandleChatV2)
}
