package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PassengerCustomHandler struct {
	*PassengersAPI
}

	return &PassengerCustomHandler{
		PassengersAPI: &PassengersAPI{},
	}
}

}
