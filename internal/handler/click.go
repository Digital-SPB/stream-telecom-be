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

func (h *Handler) clientReactionSpeed(c *gin.Context) {
	// Получаем и валидируем campaign_id
	campaignID, err := strconv.ParseInt(c.Query("campaign_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid campaign_id",
		})
		return
	}

	// Получаем метрики времени реакции
	metrics, err := h.services.GetCustomerReactionTime(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, metrics)
}

func (h *Handler) predictedBestTime(c *gin.Context) {}
