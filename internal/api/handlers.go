package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/greenblat17/stream-telecom/internal/service"
)

type Handler struct {
	activityService *service.ActivityService
}

func NewHandler(activityService *service.ActivityService) *Handler {
	return &Handler{
		activityService: activityService,
	}
}

func (h *Handler) GetCampaignActivity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	campaignIDStr := r.URL.Query().Get("campaign_id")
	if campaignIDStr == "" {
		http.Error(w, "campaign_id is required", http.StatusBadRequest)
		return
	}

	campaignID, err := strconv.ParseInt(campaignIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid campaign_id", http.StatusBadRequest)
		return
	}

	metrics, err := h.activityService.GetCampaignActivity(campaignID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if metrics == nil {
		http.Error(w, "Campaign not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) GetAllCampaignsActivity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	metrics, err := h.activityService.GetAllCampaignsActivity()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
} 