package server

import (
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
	return &DriverCustomHandler{
		DriverAPI: &DriverAPI{},
		Client:    cli,
	}
}

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
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Client.UpdateDriverProfile(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
