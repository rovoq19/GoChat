package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rovoq19/GoChat/pkg/models"
	"net/http"
)

func (h *Handler) SendMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.SendMessage(message); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &errorResponse{
			"failed to broadcast message",
		})
		return
	}

	c.JSON(http.StatusOK, &statusResponse{
		"ok",
	})
}
