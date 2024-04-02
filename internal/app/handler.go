package app

import (
	"github.com/gin-gonic/gin"
	"go-sample/internal/domain"
)

type Handler struct {
	service *domain.Service
}

func NewHandler(service *domain.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	weather := router.Group("/weather")
	weather.GET("/", h.GetCurrentWeather)

	user := router.Group("/user")
	user.POST("/", h.SetUser)
	user.GET("/:id", h.GetUserById)

	return router
}
