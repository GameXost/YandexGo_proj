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

type DriverServiceRideRequest struct {

	// User ID
	UserId string `json:"userId,omitempty"`

	StartLocation DriverServiceLocation `json:"startLocation,omitempty"`

	EndLocation DriverServiceLocation `json:"endLocation,omitempty"`
}
