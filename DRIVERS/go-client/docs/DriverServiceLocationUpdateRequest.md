# DriverServiceLocationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverId** | Pointer to **string** | Driver ID | [optional] 
**Location** | Pointer to [**DriverServiceLocation**](DriverServiceLocation.md) |  | [optional] 

## Methods

### NewDriverServiceLocationUpdateRequest

`func NewDriverServiceLocationUpdateRequest() *DriverServiceLocationUpdateRequest`

NewDriverServiceLocationUpdateRequest instantiates a new DriverServiceLocationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverServiceLocationUpdateRequestWithDefaults

`func NewDriverServiceLocationUpdateRequestWithDefaults() *DriverServiceLocationUpdateRequest`

NewDriverServiceLocationUpdateRequestWithDefaults instantiates a new DriverServiceLocationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverId

`func (o *DriverServiceLocationUpdateRequest) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *DriverServiceLocationUpdateRequest) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *DriverServiceLocationUpdateRequest) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *DriverServiceLocationUpdateRequest) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetLocation

`func (o *DriverServiceLocationUpdateRequest) GetLocation() DriverServiceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DriverServiceLocationUpdateRequest) GetLocationOk() (*DriverServiceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DriverServiceLocationUpdateRequest) SetLocation(v DriverServiceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DriverServiceLocationUpdateRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


