package server

import (
<<<<<<< HEAD
	"context"
	"net/http"

	client "github.com/GameXost/YandexGo_proj/DRIVERS/go-client"
	"github.com/gin-gonic/gin"
)

// DriverCustomHandler embeds the generated DriverAPI and stores a go-client
type DriverCustomHandler struct {
	*DriverAPI
	Client *client.APIClient
}

func NewDriverCustomHandler(cli *client.APIClient) *DriverCustomHandler {
=======
	"net/http"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/gin-gonic/gin"
)

// DriverCustomHandler встраивает сгенерированный DriverAPI и хранит gRPC-клиент
type DriverCustomHandler struct {
	*DriverAPI
	Client pb.DriversClient
}

func NewDriverCustomHandler(cli pb.DriversClient) *DriverCustomHandler {
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
	return &DriverCustomHandler{
		DriverAPI: &DriverAPI{},
		Client:    cli,
	}
}

<<<<<<< HEAD
// Overriding GetDriverProfile to call the actual go-client
func (h *DriverCustomHandler) GetDriverProfile(c *gin.Context) {
	// Pass the auth token to the client
	token := c.GetHeader("Authorization")
	ctx := context.WithValue(c.Request.Context(), client.ContextAccessToken, token)

	// Use the go-client to make the API call - token will be set from the context
	resp, httpResp, err := h.Client.DriverAPI.GetDriverProfile(ctx).Execute()

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

// Overriding UpdateDriverProfile to call the actual go-client
func (h *DriverCustomHandler) UpdateDriverProfile(c *gin.Context) {
	var req client.DriverServiceUpdateDriverProfileRequest
=======
// Переопределяем GetDriverProfile — вызываем реальный gRPC
func (h *DriverCustomHandler) GetDriverProfile(c *gin.Context) {
	token := c.GetHeader("Authorization")
	resp, err := h.Client.GetDriverProfile(c, &pb.AuthToken{Token: token})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Переопределяем UpdateDriverProfile — вызываем реальный gRPC
func (h *DriverCustomHandler) UpdateDriverProfile(c *gin.Context) {
	var req pb.UpdateDriverProfileRequest
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
<<<<<<< HEAD

	// Pass the token if available
	ctx := context.WithValue(c.Request.Context(), client.ContextAccessToken, c.GetHeader("Authorization"))

	// Use the go-client to make the API call
	resp, httpResp, err := h.Client.DriverAPI.UpdateDriverProfile(ctx).Body(req).Execute()

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
	resp, err := h.Client.UpdateDriverProfile(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
>>>>>>> 555ea6aa6e96e61c690234e3c5f1c16a72265729
}
