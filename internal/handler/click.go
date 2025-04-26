package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) clickDynamic(c *gin.Context) {}

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