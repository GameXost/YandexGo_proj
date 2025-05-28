# ClientServiceRide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | uniq id of the ride | [optional] 
**UserId** | Pointer to **string** | user&#39;s id | [optional] 
**DriverId** | Pointer to **string** | driver&#39;s id | [optional] 
**StartLocation** | Pointer to [**ClientServiceLocation**](ClientServiceLocation.md) |  | [optional] 
**EndLocation** | Pointer to [**ClientServiceLocation**](ClientServiceLocation.md) |  | [optional] 
**Status** | Pointer to **string** | ride status | [optional] 
**Timestamp** | Pointer to **string** | starting time point | [optional] 

## Methods

### NewClientServiceRide

`func NewClientServiceRide() *ClientServiceRide`

NewClientServiceRide instantiates a new ClientServiceRide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientServiceRideWithDefaults

`func NewClientServiceRideWithDefaults() *ClientServiceRide`

NewClientServiceRideWithDefaults instantiates a new ClientServiceRide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientServiceRide) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientServiceRide) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientServiceRide) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientServiceRide) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *ClientServiceRide) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ClientServiceRide) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ClientServiceRide) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ClientServiceRide) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDriverId

`func (o *ClientServiceRide) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *ClientServiceRide) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *ClientServiceRide) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *ClientServiceRide) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetStartLocation

`func (o *ClientServiceRide) GetStartLocation() ClientServiceLocation`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *ClientServiceRide) GetStartLocationOk() (*ClientServiceLocation, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *ClientServiceRide) SetStartLocation(v ClientServiceLocation)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *ClientServiceRide) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### GetEndLocation

`func (o *ClientServiceRide) GetEndLocation() ClientServiceLocation`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *ClientServiceRide) GetEndLocationOk() (*ClientServiceLocation, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *ClientServiceRide) SetEndLocation(v ClientServiceLocation)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *ClientServiceRide) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.

### GetStatus

`func (o *ClientServiceRide) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientServiceRide) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientServiceRide) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientServiceRide) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ClientServiceRide) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClientServiceRide) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClientServiceRide) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ClientServiceRide) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


