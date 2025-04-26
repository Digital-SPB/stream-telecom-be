package main

import (
	"log"
	"net/http"

	"github.com/greenblat17/stream-telecom/internal/api"
	"github.com/greenblat17/stream-telecom/internal/service"
)

func main() {
	// Initialize services
	activityService := service.NewActivityService()
	if err := activityService.LoadData(); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// Initialize handlers
	handler := api.NewHandler(activityService)

	// Set up routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/campaigns", handler.GetAllCampaigns)
	mux.HandleFunc("/api/v1/campaigns/activity", handler.GetAllCampaignsActivity)
	mux.HandleFunc("/api/v1/campaigns/activity/single", handler.GetCampaignActivity)
	mux.HandleFunc("/api/v1/campaigns/reaction-time", handler.GetCustomerReactionTime)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
} 