package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Perfect29/proxy-server/models"
	"github.com/Perfect29/proxy-server/storage"
	"github.com/google/uuid"
)

func HandleProxyRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.ProxyRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	proxyReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for key, value := range req.Headers {
		proxyReq.Header.Set(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Request to external server failed: "+err.Error(), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	status := resp.StatusCode
	headers := resp.Header
	length := int64(len(body))

	response := models.ProxyResponse{
		ID:      uuid.New().String(),
		Status:  status,
		Headers: headers,
		Length:  length,
	}

	storage.SaveLog(response.ID, models.ProxyLog{
		Request:  req,
		Response: response,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleGetLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[2] == "" {
		http.Error(w, "Missing or Invalid log ID", http.StatusBadRequest)
		return
	}

	id := parts[2]

	log, ok := storage.GetLog(id)

	if !ok {
		http.Error(w, "Log not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(log)
}
