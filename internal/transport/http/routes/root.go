package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hw_3/internal/transport/http/models"
	"net/http"
)

func root(c *gin.Context) {
	c.JSON(http.StatusOK, models.SemVerResponse{Ver: viper.GetString("apiVersion")})
}

func InitRootRoute(router *gin.RouterGroup) {
	router.GET("/", root)
}
