package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) heatMap(c *gin.Context) {
	// Получаем опциональные параметры периода
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	var start, end time.Time
	var err error

	// Парсим start_date если указан
	if startDate != "" {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "invalid start_date format, use YYYY-MM-DD",
			})
			return
		}
	}

	// Парсим end_date если указан
	if endDate != "" {
		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "invalid end_date format, use YYYY-MM-DD",
			})
			return
		}

		// Устанавливаем конец дня для end_date
		end = end.Add(24 * time.Hour).Add(-time.Second)
	}

	// Проверяем корректность периода
	if !start.IsZero() && !end.IsZero() && end.Before(start) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "end_date must be after start_date",
		})
		return
	}

	heatMap := h.services.Regions.GetMembersHeatMap(start, end)

	c.JSON(http.StatusOK, heatMap)
}

func (h *Handler) clientHotPoint(c *gin.Context) {
	// Получаем опциональные параметры периода
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	var start, end time.Time
	var err error

	// Парсим start_date если указан
	if startDate != "" {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "invalid start_date format, use YYYY-MM-DD",
			})
			return
		}
	}

	// Парсим end_date если указан
	if endDate != "" {
		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "invalid end_date format, use YYYY-MM-DD",
			})
			return
		}

		// Устанавливаем конец дня для end_date
		end = end.Add(24 * time.Hour).Add(-time.Second)
	}

	// Проверяем корректность периода
	if !start.IsZero() && !end.IsZero() && end.Before(start) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "end_date must be after start_date",
		})
		return
	}

	heatMap := h.services.Regions.GetCountClick(start, end)

	c.JSON(http.StatusOK, heatMap)
}

func (h *Handler) activityTime(c *gin.Context) {}
