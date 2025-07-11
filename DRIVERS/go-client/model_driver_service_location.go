/*
DRIVER Service API

API for managing drivers and ride operations

API version: 2.0
Contact: sergejs.dyldin@yandex.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go_client

import (
	"encoding/json"
)

// checks if the DriverServiceLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriverServiceLocation{}

// DriverServiceLocation struct for DriverServiceLocation
type DriverServiceLocation struct {
	// Latitude
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// NewDriverServiceLocation instantiates a new DriverServiceLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriverServiceLocation() *DriverServiceLocation {
	this := DriverServiceLocation{}
	return &this
}

// NewDriverServiceLocationWithDefaults instantiates a new DriverServiceLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriverServiceLocationWithDefaults() *DriverServiceLocation {
	this := DriverServiceLocation{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *DriverServiceLocation) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverServiceLocation) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *DriverServiceLocation) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *DriverServiceLocation) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *DriverServiceLocation) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverServiceLocation) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *DriverServiceLocation) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *DriverServiceLocation) SetLongitude(v float64) {
	o.Longitude = &v
}

func (o DriverServiceLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriverServiceLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	return toSerialize, nil
}

type NullableDriverServiceLocation struct {
	value *DriverServiceLocation
	isSet bool
}

func (v NullableDriverServiceLocation) Get() *DriverServiceLocation {
	return v.value
}

func (v *NullableDriverServiceLocation) Set(val *DriverServiceLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDriverServiceLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDriverServiceLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriverServiceLocation(val *DriverServiceLocation) *NullableDriverServiceLocation {
	return &NullableDriverServiceLocation{value: val, isSet: true}
}

func (v NullableDriverServiceLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriverServiceLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


