/*
 * Driver Service API
 *
 * API for managing drivers and ride operations
 *
 * API version: 2.0
 * Contact: sergejs.dyldin@yandex.ru
 */

package server

// NewDriverAPI creates a new DriverAPI instance
func NewDriverAPI() *DriverAPI {
	return &DriverAPI{}
}

// NewLocationAPI creates a new LocationAPI instance
func NewLocationAPI() *LocationAPI {
	return &LocationAPI{}
}

// NewPassengersAPI creates a new PassengersAPI instance
func NewPassengersAPI() *PassengersAPI {
	return &PassengersAPI{}
}

// NewRidesAPI creates a new RidesAPI instance
func NewRidesAPI() *RidesAPI {
	return &RidesAPI{}
}
