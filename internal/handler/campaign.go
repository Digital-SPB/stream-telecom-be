package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) allCampaigns(c *gin.Context) {
	pageStr := c.Query("page")
	perPageStr := c.Query("per_page")

	page := 1
	perPage := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if perPageStr != "" {
		if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 {
			perPage = pp
		}
	}

	campaigns := h.services.GetAllCampaigns(page, perPage)

	c.JSON(http.StatusOK, campaigns)
}

func (h *Handler) campainActivity(c *gin.Context) {
	// Получаем и валидируем campaign_id
	campaignID, err := strconv.ParseInt(c.Query("campaign_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid campaign_id",
		})
		return
	}

	// Получаем и валидируем count_hours
	countHours, err := strconv.ParseInt(c.Query("count_hours"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid count_hours",
		})
		return
	}

	// Получаем метрики активности
	metrics, err := h.services.GetCampaignActivity(campaignID, countHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	// Проверяем, найдена ли кампания
	if metrics == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Campaign not found",
		})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, metrics)
}

func (h *Handler) campainCreateDynamic(c *gin.Context) {
	intervalType := c.Query("interval_type")
	startTimeString := c.Query("start_time")
	endTimeString := c.Query("end_time")

	starTime, err := time.Parse("2006-01-02", startTimeString)
	fmt.Println(startTimeString)
	fmt.Println(starTime)
	fmt.Println(err)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid start time")
	}
	endTime, err := time.Parse("2006-01-02", endTimeString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid end time")
		return
	}

	fmt.Println(intervalType)
	if intervalType != "day" && intervalType != "month" {
		newErrorResponse(c, http.StatusBadGateway, "invalid interval type")
		return
	}

	res, err := h.services.GetCreationDynamic(starTime, endTime, intervalType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, res)
}
