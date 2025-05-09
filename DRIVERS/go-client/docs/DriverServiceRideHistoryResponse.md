# DriverServiceRideHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rides** | Pointer to [**[]DriverServiceRide**](DriverServiceRide.md) | previous rides | [optional] 

## Methods

### NewDriverServiceRideHistoryResponse

`func NewDriverServiceRideHistoryResponse() *DriverServiceRideHistoryResponse`

NewDriverServiceRideHistoryResponse instantiates a new DriverServiceRideHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverServiceRideHistoryResponseWithDefaults

`func NewDriverServiceRideHistoryResponseWithDefaults() *DriverServiceRideHistoryResponse`

NewDriverServiceRideHistoryResponseWithDefaults instantiates a new DriverServiceRideHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRides

`func (o *DriverServiceRideHistoryResponse) GetRides() []DriverServiceRide`

GetRides returns the Rides field if non-nil, zero value otherwise.

### GetRidesOk

`func (o *DriverServiceRideHistoryResponse) GetRidesOk() (*[]DriverServiceRide, bool)`

GetRidesOk returns a tuple with the Rides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRides

`func (o *DriverServiceRideHistoryResponse) SetRides(v []DriverServiceRide)`

SetRides sets Rides field to given value.

### HasRides

`func (o *DriverServiceRideHistoryResponse) HasRides() bool`

HasRides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


