/*
USERS Service API

API for managing users and ride operations

API version: 2.0
Contact: sergejs.dyldin@yandex.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package go_client

import (
	"encoding/json"
)

// checks if the ClientServiceUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientServiceUser{}

// ClientServiceUser struct for ClientServiceUser
type ClientServiceUser struct {
	// uniq user identificator
	Id *string `json:"id,omitempty"`
	// user's firstname
	Username *string `json:"username,omitempty"`
	// user's email
	Email *string `json:"email,omitempty"`
	// user's phone number
	Phone *string `json:"phone,omitempty"`
}

// NewClientServiceUser instantiates a new ClientServiceUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientServiceUser() *ClientServiceUser {
	this := ClientServiceUser{}
	return &this
}

// NewClientServiceUserWithDefaults instantiates a new ClientServiceUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientServiceUserWithDefaults() *ClientServiceUser {
	this := ClientServiceUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientServiceUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientServiceUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientServiceUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientServiceUser) SetId(v string) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ClientServiceUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientServiceUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ClientServiceUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ClientServiceUser) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ClientServiceUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientServiceUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ClientServiceUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ClientServiceUser) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ClientServiceUser) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientServiceUser) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ClientServiceUser) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ClientServiceUser) SetPhone(v string) {
	o.Phone = &v
}

func (o ClientServiceUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientServiceUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

type NullableClientServiceUser struct {
	value *ClientServiceUser
	isSet bool
}

func (v NullableClientServiceUser) Get() *ClientServiceUser {
	return v.value
}

func (v *NullableClientServiceUser) Set(val *ClientServiceUser) {
	v.value = val
	v.isSet = true
}

func (v NullableClientServiceUser) IsSet() bool {
	return v.isSet
}

func (v *NullableClientServiceUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientServiceUser(val *ClientServiceUser) *NullableClientServiceUser {
	return &NullableClientServiceUser{value: val, isSet: true}
}

func (v NullableClientServiceUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientServiceUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


