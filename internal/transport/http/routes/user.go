package routes

import (
	"github.com/gin-gonic/gin"
	"hw_3/internal/service/domain"
	handlers "hw_3/internal/transport/http/haldlers"
)

func InitUserRoutes(router *gin.RouterGroup, us domain.UserService) {
	h := handlers.NewUserHandler(us)

	balance := router.Group("/balance")
	{
		balance.POST("/transfer", h.TransferBalance)
		balance.POST("/add", h.AddToUserBalance)
	}
	user := router.Group("/user")
	{
		user.GET("/:id", h.GetUser)
		user.POST("/create", h.CreateUser)
	}
}
