// server/custom_location_handlers.go
package server

import (
	"net/http"
	"strconv"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/gin-gonic/gin"
)

type LocationCustomHandler struct {
	*LocationAPI
	Client pb.DriversClient
}

func NewLocationCustomHandler(cli pb.DriversClient) *LocationCustomHandler {
	return &LocationCustomHandler{
		LocationAPI: &LocationAPI{},
		Client:      cli,
	}
}

func (h *LocationCustomHandler) GetNearbyRequests(c *gin.Context) {
	latStr := c.Query("lat")
	lonStr := c.Query("lon")

	// Проверка наличия параметров
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
}
