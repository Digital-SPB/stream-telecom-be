package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) clickDynamic(c *gin.Context) {
	campaignId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	res, err := h.services.GetClickDynamic(int64(campaignId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) clientReactionSpeed(c *gin.Context) {}

func (h *Handler) predictedBestTime(c *gin.Context) {}
