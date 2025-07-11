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

type DriverServiceDriver struct {

	// Unic driver identificator
	Id string `json:"id,omitempty"`

	// Driver's first_name
	Username string `json:"username,omitempty"`

	// contact number
	Phone string `json:"phone,omitempty"`

	// Car model
	CarModel string `json:"carModel,omitempty"`

	// Email address
	Email string `json:"email,omitempty"`

	// Car color
	CarColor string `json:"carColor,omitempty"`

	// Car manufacturer
	CarMark string `json:"carMark,omitempty"`

	// License plate number
	CarNumber string `json:"carNumber,omitempty"`
}
