// server/custom_location_handlers.go
package server

import (
<<<<<<< HEAD
	"context"
	"net/http"
	"strconv"

	client "github.com/GameXost/YandexGo_proj/DRIVERS/go-client"
=======
	"net/http"
	"strconv"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
	"github.com/gin-gonic/gin"
)

type LocationCustomHandler struct {
	*LocationAPI
<<<<<<< HEAD
	Client *client.APIClient
}

func NewLocationCustomHandler(cli *client.APIClient) *LocationCustomHandler {
=======
	Client pb.DriversClient
}

func NewLocationCustomHandler(cli pb.DriversClient) *LocationCustomHandler {
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
	return &LocationCustomHandler{
		LocationAPI: &LocationAPI{},
		Client:      cli,
	}
}

func (h *LocationCustomHandler) GetNearbyRequests(c *gin.Context) {
	latStr := c.Query("lat")
	lonStr := c.Query("lon")

<<<<<<< HEAD
	// Check for required parameters
=======
	// Проверка наличия параметров
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
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

<<<<<<< HEAD
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
=======
	req := &pb.Location{
		Latitude:  lat,
		Longitude: lon,
	}

	resp, err := h.Client.GetNearbyRequests(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *LocationCustomHandler) UpdateLocation(c *gin.Context) {
	var updates []pb.LocationUpdateRequest
	if err := c.BindJSON(&updates); err != nil {
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<< HEAD
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
=======
	stream, err := h.Client.UpdateLocation(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, u := range updates {
		if err := stream.Send(&u); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reply)
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
}
