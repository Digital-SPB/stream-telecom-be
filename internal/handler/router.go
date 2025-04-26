package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenblat17/stream-telecom/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	apiv1 := router.Group("api/v1")
	{
		apiv1.GET("activity", h.campainActivity)
		apiv1.GET("click-dynamic/:id", h.clickDynamic)
		apiv1.GET("create-campain-dynamic", h.campainCreateDynamic)
		apiv1.GET("reaction-time", h.clientReactionSpeed)
		apiv1.GET("heat-map", h.heatMap)
		apiv1.GET("client-hot-point", h.clientHotPoint)
		apiv1.GET("activity-time", h.activityTime)
		apiv1.GET("predict-best-time", h.predictedBestTime)
		apiv1.GET("campaigns", h.allCampaigns)
	}

	return router
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Status:  status,
		Message: message,
	})
}
