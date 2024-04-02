package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SetUser(ctx *gin.Context) {
	var input UserAPIRequest

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	if err := h.service.SetUser(input.Id, input.Login); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	login, err := h.service.GetUser(int64(id))
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"login": login,
	})
}
