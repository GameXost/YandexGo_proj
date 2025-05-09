// server/custom_location_handlers.go
package server

import (
	"context"
	"net/http"
	"strconv"

	client "github.com/GameXost/YandexGo_proj/DRIVERS/go-client"
	"github.com/gin-gonic/gin"
)

type LocationCustomHandler struct {
	*LocationAPI
	Client *client.APIClient
}

func NewLocationCustomHandler(cli *client.APIClient) *LocationCustomHandler {
	return &LocationCustomHandler{
		LocationAPI: &LocationAPI{},
		Client:      cli,
	}
}

func (h *LocationCustomHandler) GetNearbyRequests(c *gin.Context) {
	latStr := c.Query("lat")
	lonStr := c.Query("lon")

	// Check for required parameters
	if latStr == "" || lonStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lat and lon are required"})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lat format"})
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lon format"})
		return
	}

	// Pass the token if available
	ctx := context.WithValue(c.Request.Context(), client.ContextAccessToken, c.GetHeader("Authorization"))

	// Use the go-client to make the API call
	req := h.Client.LocationAPI.GetNearbyRequests(ctx)
	req = req.Latitude(lat)
	req = req.Longitude(lon)

	resp, httpResp, err := req.Execute()

	// Handle errors
	if err != nil {
		// If we got an HTTP response, use its status code
		statusCode := http.StatusInternalServerError
		if httpResp != nil {
			statusCode = httpResp.StatusCode
		}

		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpResp.StatusCode, resp)
}

func (h *LocationCustomHandler) UpdateLocation(c *gin.Context) {
	var locationUpdate client.DriverServiceLocationUpdateRequest
	if err := c.BindJSON(&locationUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pass the token if available
	ctx := context.WithValue(c.Request.Context(), client.ContextAccessToken, c.GetHeader("Authorization"))

	// Use the go-client to make the API call
	resp, httpResp, err := h.Client.LocationAPI.UpdateLocation(ctx).Body(locationUpdate).Execute()

	// Handle errors
	if err != nil {
		// If we got an HTTP response, use its status code
		statusCode := http.StatusInternalServerError
		if httpResp != nil {
			statusCode = httpResp.StatusCode
		}

		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpResp.StatusCode, resp)
}
