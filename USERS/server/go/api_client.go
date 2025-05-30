/*
 * clients.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"github.com/gin-gonic/gin"
)

type ClientAPI struct {
}

// Post /ride/:id/cancel
func (api *ClientAPI) ClientCancelRide(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /driver/:id
func (api *ClientAPI) ClientGetDriverInfo(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /driver/:id/location
// сведения о водителе
func (api *ClientAPI) ClientGetDriverLocation(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /ride/history
func (api *ClientAPI) ClientGetRideHistory(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /ride/:id
func (api *ClientAPI) ClientGetRideStatus(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /user/profile
// операции с пользовательским профилем
func (api *ClientAPI) ClientGetUserProfile(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /ride/request
// операции с заказами
func (api *ClientAPI) ClientRequestRide(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /user/profile
func (api *ClientAPI) ClientUpdateUserProfile(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
