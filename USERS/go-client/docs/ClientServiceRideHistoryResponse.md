# ClientServiceRideHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rides** | Pointer to [**[]ClientServiceRide**](ClientServiceRide.md) | List of past rides | [optional] 

## Methods

### NewClientServiceRideHistoryResponse

`func NewClientServiceRideHistoryResponse() *ClientServiceRideHistoryResponse`

NewClientServiceRideHistoryResponse instantiates a new ClientServiceRideHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientServiceRideHistoryResponseWithDefaults

`func NewClientServiceRideHistoryResponseWithDefaults() *ClientServiceRideHistoryResponse`

NewClientServiceRideHistoryResponseWithDefaults instantiates a new ClientServiceRideHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRides

`func (o *ClientServiceRideHistoryResponse) GetRides() []ClientServiceRide`

GetRides returns the Rides field if non-nil, zero value otherwise.

### GetRidesOk

`func (o *ClientServiceRideHistoryResponse) GetRidesOk() (*[]ClientServiceRide, bool)`

GetRidesOk returns a tuple with the Rides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRides

`func (o *ClientServiceRideHistoryResponse) SetRides(v []ClientServiceRide)`

SetRides sets Rides field to given value.

### HasRides

`func (o *ClientServiceRideHistoryResponse) HasRides() bool`

HasRides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


