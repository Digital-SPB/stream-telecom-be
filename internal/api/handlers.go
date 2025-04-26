package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/greenblat17/stream-telecom/internal/service"
)

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
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	campaignIDStr := r.URL.Query().Get("campaign_id")
	if campaignIDStr == "" {
		writeJSONError(w, http.StatusBadRequest, "campaign_id is required")
		return
	}

	campaignID, err := strconv.ParseInt(campaignIDStr, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid campaign_id")
		return
	}

	countHoursStr := r.URL.Query().Get("count_hours")
	if countHoursStr == "" {
		writeJSONError(w, http.StatusBadRequest, "count_hours is required")
		return
	}

	countHours, err := strconv.ParseInt(countHoursStr, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid count_hours")
		return
	}

	metrics, err := h.activityService.GetCampaignActivity(campaignID, countHours)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if metrics == nil {
		writeJSONError(w, http.StatusNotFound, "Campaign not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) GetAllCampaignsActivity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	metrics, err := h.activityService.GetAllCampaignsActivity()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) GetCustomerReactionTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	campaignIDStr := r.URL.Query().Get("campaign_id")
	if campaignIDStr == "" {
		writeJSONError(w, http.StatusBadRequest, "campaign_id is required")
		return
	}

	campaignID, err := strconv.ParseInt(campaignIDStr, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid campaign_id")
		return
	}

	metrics, err := h.activityService.GetCustomerReactionTime(campaignID)
	if err != nil {
		if noClicksErr, ok := err.(*service.NoClicksFoundError); ok {
			writeJSONError(w, http.StatusNotFound, noClicksErr.Error())
			return
		}
		writeJSONError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) GetAllCampaigns(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
        return
    }

    // Получаем параметры пагинации из query parameters
    pageStr := r.URL.Query().Get("page")
    perPageStr := r.URL.Query().Get("per_page")

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

    campaigns := h.activityService.GetAllCampaigns(page, perPage)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(campaigns)
}