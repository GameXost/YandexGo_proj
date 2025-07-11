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

// checks if the DriverServiceRideRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriverServiceRideRequest{}

// DriverServiceRideRequest struct for DriverServiceRideRequest
type DriverServiceRideRequest struct {
	// User ID
	UserId *string `json:"userId,omitempty"`
	StartLocation *DriverServiceLocation `json:"startLocation,omitempty"`
	EndLocation *DriverServiceLocation `json:"endLocation,omitempty"`
}

// NewDriverServiceRideRequest instantiates a new DriverServiceRideRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriverServiceRideRequest() *DriverServiceRideRequest {
	this := DriverServiceRideRequest{}
	return &this
}

// NewDriverServiceRideRequestWithDefaults instantiates a new DriverServiceRideRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriverServiceRideRequestWithDefaults() *DriverServiceRideRequest {
	this := DriverServiceRideRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DriverServiceRideRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverServiceRideRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DriverServiceRideRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DriverServiceRideRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetStartLocation returns the StartLocation field value if set, zero value otherwise.
func (o *DriverServiceRideRequest) GetStartLocation() DriverServiceLocation {
	if o == nil || IsNil(o.StartLocation) {
		var ret DriverServiceLocation
		return ret
	}
	return *o.StartLocation
}

// GetStartLocationOk returns a tuple with the StartLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverServiceRideRequest) GetStartLocationOk() (*DriverServiceLocation, bool) {
	if o == nil || IsNil(o.StartLocation) {
		return nil, false
	}
	return o.StartLocation, true
}

// HasStartLocation returns a boolean if a field has been set.
func (o *DriverServiceRideRequest) HasStartLocation() bool {
	if o != nil && !IsNil(o.StartLocation) {
		return true
	}

	return false
}

// SetStartLocation gets a reference to the given DriverServiceLocation and assigns it to the StartLocation field.
func (o *DriverServiceRideRequest) SetStartLocation(v DriverServiceLocation) {
	o.StartLocation = &v
}

// GetEndLocation returns the EndLocation field value if set, zero value otherwise.
func (o *DriverServiceRideRequest) GetEndLocation() DriverServiceLocation {
	if o == nil || IsNil(o.EndLocation) {
		var ret DriverServiceLocation
		return ret
	}
	return *o.EndLocation
}

// GetEndLocationOk returns a tuple with the EndLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverServiceRideRequest) GetEndLocationOk() (*DriverServiceLocation, bool) {
	if o == nil || IsNil(o.EndLocation) {
		return nil, false
	}
	return o.EndLocation, true
}

// HasEndLocation returns a boolean if a field has been set.
func (o *DriverServiceRideRequest) HasEndLocation() bool {
	if o != nil && !IsNil(o.EndLocation) {
		return true
	}

	return false
}

// SetEndLocation gets a reference to the given DriverServiceLocation and assigns it to the EndLocation field.
func (o *DriverServiceRideRequest) SetEndLocation(v DriverServiceLocation) {
	o.EndLocation = &v
}

func (o DriverServiceRideRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriverServiceRideRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.StartLocation) {
		toSerialize["startLocation"] = o.StartLocation
	}
	if !IsNil(o.EndLocation) {
		toSerialize["endLocation"] = o.EndLocation
	}
	return toSerialize, nil
}

type NullableDriverServiceRideRequest struct {
	value *DriverServiceRideRequest
	isSet bool
}

func (v NullableDriverServiceRideRequest) Get() *DriverServiceRideRequest {
	return v.value
}

func (v *NullableDriverServiceRideRequest) Set(val *DriverServiceRideRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDriverServiceRideRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDriverServiceRideRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriverServiceRideRequest(val *DriverServiceRideRequest) *NullableDriverServiceRideRequest {
	return &NullableDriverServiceRideRequest{value: val, isSet: true}
}

func (v NullableDriverServiceRideRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriverServiceRideRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


