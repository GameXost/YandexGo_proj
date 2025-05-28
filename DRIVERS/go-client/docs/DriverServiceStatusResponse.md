# DriverServiceStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | ride status | [optional] 
**Message** | Pointer to **string** | details | [optional] 

## Methods

### NewDriverServiceStatusResponse

`func NewDriverServiceStatusResponse() *DriverServiceStatusResponse`

NewDriverServiceStatusResponse instantiates a new DriverServiceStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverServiceStatusResponseWithDefaults

`func NewDriverServiceStatusResponseWithDefaults() *DriverServiceStatusResponse`

NewDriverServiceStatusResponseWithDefaults instantiates a new DriverServiceStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DriverServiceStatusResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DriverServiceStatusResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DriverServiceStatusResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DriverServiceStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DriverServiceStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DriverServiceStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DriverServiceStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DriverServiceStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


