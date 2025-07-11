/*
 * DRIVER Service API
 *
 * API for managing drivers and ride operations
 *
 * API version: 2.0
 * Contact: sergejs.dyldin@yandex.ru
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the DriverAPI part of the API
	DriverAPI DriverAPI
	// Routes for the LocationAPI part of the API
	LocationAPI LocationAPI
	// Routes for the PassengersAPI part of the API
	PassengersAPI PassengersAPI
	// Routes for the RidesAPI part of the API
	RidesAPI RidesAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"GetDriverProfile",
			http.MethodGet,
			"/driver/profile",
			handleFunctions.DriverAPI.GetDriverProfile,
		},
		{
			"UpdateDriverProfile",
			http.MethodPut,
			"/driver/profile",
			handleFunctions.DriverAPI.UpdateDriverProfile,
		},
		{
			"GetNearbyRequests",
			http.MethodGet,
			"/driver/nearby_req",
			handleFunctions.LocationAPI.GetNearbyRequests,
		},
		{
			"UpdateLocation",
			http.MethodPost,
			"/driver/location",
			handleFunctions.LocationAPI.UpdateLocation,
		},
		{
			"GetPassengerInfo",
			http.MethodGet,
			"/user/:id",
			handleFunctions.PassengersAPI.GetPassengerInfo,
		},
		{
			"AcceptRide",
			http.MethodPost,
			"/ride/:id/accept",
			handleFunctions.RidesAPI.AcceptRide,
		},
		{
			"CancelRide",
			http.MethodPost,
			"/ride/:id/cancel",
			handleFunctions.RidesAPI.CancelRide,
		},
		{
			"CompleteRide",
			http.MethodPost,
			"/ride/:id/complete",
			handleFunctions.RidesAPI.CompleteRide,
		},
		{
			"GetCurrentRide",
			http.MethodGet,
			"/driver/current_ride/:id",
			handleFunctions.RidesAPI.GetCurrentRide,
		},
		{
			"GetRideHistory",
			http.MethodGet,
			"/driver/:id/rides",
			handleFunctions.RidesAPI.GetRideHistory,
		},
	}
}
