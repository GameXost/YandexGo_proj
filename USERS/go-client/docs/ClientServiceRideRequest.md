# ClientServiceRideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User&#39;s unique identifier | [optional] 
**StartLocation** | Pointer to [**ClientServiceLocation**](ClientServiceLocation.md) |  | [optional] 
**EndLocation** | Pointer to [**ClientServiceLocation**](ClientServiceLocation.md) |  | [optional] 

## Methods

### NewClientServiceRideRequest

`func NewClientServiceRideRequest() *ClientServiceRideRequest`

NewClientServiceRideRequest instantiates a new ClientServiceRideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientServiceRideRequestWithDefaults

`func NewClientServiceRideRequestWithDefaults() *ClientServiceRideRequest`

NewClientServiceRideRequestWithDefaults instantiates a new ClientServiceRideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ClientServiceRideRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ClientServiceRideRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ClientServiceRideRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ClientServiceRideRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStartLocation

`func (o *ClientServiceRideRequest) GetStartLocation() ClientServiceLocation`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *ClientServiceRideRequest) GetStartLocationOk() (*ClientServiceLocation, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *ClientServiceRideRequest) SetStartLocation(v ClientServiceLocation)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *ClientServiceRideRequest) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### GetEndLocation

`func (o *ClientServiceRideRequest) GetEndLocation() ClientServiceLocation`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *ClientServiceRideRequest) GetEndLocationOk() (*ClientServiceLocation, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *ClientServiceRideRequest) SetEndLocation(v ClientServiceLocation)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *ClientServiceRideRequest) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


