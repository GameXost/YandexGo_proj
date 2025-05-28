# DriverServiceUpdateDriverProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Driver ID | [optional] 
**Driver** | Pointer to [**DriverServiceDriver**](DriverServiceDriver.md) |  | [optional] 

## Methods

### NewDriverServiceUpdateDriverProfileRequest

`func NewDriverServiceUpdateDriverProfileRequest() *DriverServiceUpdateDriverProfileRequest`

NewDriverServiceUpdateDriverProfileRequest instantiates a new DriverServiceUpdateDriverProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverServiceUpdateDriverProfileRequestWithDefaults

`func NewDriverServiceUpdateDriverProfileRequestWithDefaults() *DriverServiceUpdateDriverProfileRequest`

NewDriverServiceUpdateDriverProfileRequestWithDefaults instantiates a new DriverServiceUpdateDriverProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverServiceUpdateDriverProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverServiceUpdateDriverProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverServiceUpdateDriverProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriverServiceUpdateDriverProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDriver

`func (o *DriverServiceUpdateDriverProfileRequest) GetDriver() DriverServiceDriver`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *DriverServiceUpdateDriverProfileRequest) GetDriverOk() (*DriverServiceDriver, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *DriverServiceUpdateDriverProfileRequest) SetDriver(v DriverServiceDriver)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *DriverServiceUpdateDriverProfileRequest) HasDriver() bool`

HasDriver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


