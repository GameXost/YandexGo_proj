package server

import (
	"context"
	"net/http"

	client "github.com/GameXost/YandexGo_proj/DRIVERS/go-client"
	"github.com/gin-gonic/gin"
)

type PassengerCustomHandler struct {
	*PassengersAPI
	Client *client.APIClient
}

func NewPassengerCustomHandler(cli *client.APIClient) *PassengerCustomHandler {
	return &PassengerCustomHandler{
		PassengersAPI: &PassengersAPI{},
		Client:        cli,
	}
}

func (h *PassengerCustomHandler) GetPassengerInfo(c *gin.Context) {
	userID := c.Param("id")

	// Pass the token if available
	ctx := context.WithValue(c.Request.Context(), client.ContextAccessToken, c.GetHeader("Authorization"))

	// Use the go-client to make the API call
	resp, httpResp, err := h.Client.PassengersAPI.GetPassengerInfo(ctx, userID).Execute()

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
