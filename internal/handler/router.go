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

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())


	apiv1 := router.Group("api/v1")
	{
		apiv1.GET("activity", h.campainActivity)
		apiv1.GET("click-dynamic/:id", h.clickDynamic)
		apiv1.GET("create-campaign-dynamic", h.campainCreateDynamic)
		apiv1.GET("reaction-time", h.clientReactionSpeed)
		apiv1.GET("heat-map", h.heatMap)
		apiv1.GET("client-hot-point", h.clientHotPoint)
		apiv1.GET("activity-time", h.activityTime)
		apiv1.GET("predict-best-time", h.predictedBestTime)
		apiv1.GET("campaigns", h.allCampaigns)
		apiv1.GET("regions-info", h.regionInfo)
	}

	return router
}
