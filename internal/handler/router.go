package handlers

import (
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
		apiv1.GET("/activity/:id", h.campainActivity)
		apiv1.GET("/click-dynamic/:id", h.clickDynamic)
		apiv1.GET("/create-campain-dynamic", h.campainCreateDynamic)
		apiv1.GET("/client-reaction-speed", h.clientReactionSpeed)
		apiv1.GET("/heat-map", h.heatMap)
		apiv1.GET("/client-hot-point", h.clientHotPoint)
		apiv1.GET("/activity-time", h.activityTime)
		apiv1.GET("predict-best-time", h.predictedBestTime)
	}

	return router
}
