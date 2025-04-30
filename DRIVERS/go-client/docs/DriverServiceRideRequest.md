# DriverServiceRideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User ID | [optional] 
**StartLocation** | Pointer to [**DriverServiceLocation**](DriverServiceLocation.md) |  | [optional] 
**EndLocation** | Pointer to [**DriverServiceLocation**](DriverServiceLocation.md) |  | [optional] 

## Methods

### NewDriverServiceRideRequest

`func NewDriverServiceRideRequest() *DriverServiceRideRequest`

NewDriverServiceRideRequest instantiates a new DriverServiceRideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverServiceRideRequestWithDefaults

`func NewDriverServiceRideRequestWithDefaults() *DriverServiceRideRequest`

NewDriverServiceRideRequestWithDefaults instantiates a new DriverServiceRideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *DriverServiceRideRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DriverServiceRideRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DriverServiceRideRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DriverServiceRideRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStartLocation

`func (o *DriverServiceRideRequest) GetStartLocation() DriverServiceLocation`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *DriverServiceRideRequest) GetStartLocationOk() (*DriverServiceLocation, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *DriverServiceRideRequest) SetStartLocation(v DriverServiceLocation)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *DriverServiceRideRequest) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### GetEndLocation

`func (o *DriverServiceRideRequest) GetEndLocation() DriverServiceLocation`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *DriverServiceRideRequest) GetEndLocationOk() (*DriverServiceLocation, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *DriverServiceRideRequest) SetEndLocation(v DriverServiceLocation)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *DriverServiceRideRequest) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


