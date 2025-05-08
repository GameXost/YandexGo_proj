package server

import (
	"net/http"

	pb "github.com/GameXost/YandexGo_proj/DRIVERS/API/generated/drivers"
	"github.com/gin-gonic/gin"
)

type PassengerCustomHandler struct {
	*PassengersAPI
	Client pb.DriversClient
}

func NewPassengerCustomHandler(client pb.DriversClient) *PassengerCustomHandler {
	return &PassengerCustomHandler{
		PassengersAPI: &PassengersAPI{},
		Client:        client,
	}
}

func GetPassengerDetails(c *gin.Context) {
	// Реализация метода
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
