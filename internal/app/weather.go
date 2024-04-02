package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCurrentWeather(ctx *gin.Context) {
	currTemp, err := h.service.GetTemp()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"current_temp": currTemp,
	})
}
