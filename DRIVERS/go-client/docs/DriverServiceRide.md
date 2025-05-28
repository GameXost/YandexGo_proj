# DriverServiceRide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | unic id of ride | [optional] 
**UserId** | Pointer to **string** | unic user&#39;s id | [optional] 
**DriverId** | Pointer to **string** | unic drivers id | [optional] 
**StartLocation** | Pointer to [**DriverServiceLocation**](DriverServiceLocation.md) |  | [optional] 
**EndLocation** | Pointer to [**DriverServiceLocation**](DriverServiceLocation.md) |  | [optional] 
**Status** | Pointer to **string** | ride status | [optional] 
**Timestamp** | Pointer to **string** | time mark for start point | [optional] 

## Methods

### NewDriverServiceRide

`func NewDriverServiceRide() *DriverServiceRide`

NewDriverServiceRide instantiates a new DriverServiceRide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverServiceRideWithDefaults

`func NewDriverServiceRideWithDefaults() *DriverServiceRide`

NewDriverServiceRideWithDefaults instantiates a new DriverServiceRide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverServiceRide) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverServiceRide) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverServiceRide) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriverServiceRide) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *DriverServiceRide) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DriverServiceRide) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DriverServiceRide) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DriverServiceRide) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDriverId

`func (o *DriverServiceRide) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *DriverServiceRide) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *DriverServiceRide) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *DriverServiceRide) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetStartLocation

`func (o *DriverServiceRide) GetStartLocation() DriverServiceLocation`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *DriverServiceRide) GetStartLocationOk() (*DriverServiceLocation, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *DriverServiceRide) SetStartLocation(v DriverServiceLocation)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *DriverServiceRide) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### GetEndLocation

`func (o *DriverServiceRide) GetEndLocation() DriverServiceLocation`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *DriverServiceRide) GetEndLocationOk() (*DriverServiceLocation, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *DriverServiceRide) SetEndLocation(v DriverServiceLocation)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *DriverServiceRide) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.

### GetStatus

`func (o *DriverServiceRide) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DriverServiceRide) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DriverServiceRide) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DriverServiceRide) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *DriverServiceRide) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DriverServiceRide) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DriverServiceRide) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DriverServiceRide) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


