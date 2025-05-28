# ClientServiceDriver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | current driver&#39;s uniq identificator | [optional] 
**Username** | Pointer to **string** | current driver&#39;s firstname | [optional] 
**Phone** | Pointer to **string** | your driver&#39;s phone number | [optional] 
**CarModel** | Pointer to **string** | driver&#39;s car model | [optional] 
**Location** | Pointer to [**ClientServiceLocation**](ClientServiceLocation.md) |  | [optional] 
**CarMake** | Pointer to **string** | car make | [optional] 
**CarNumber** | Pointer to **string** | plate number | [optional] 
**CarColor** | Pointer to **string** | car color | [optional] 

## Methods

### NewClientServiceDriver

`func NewClientServiceDriver() *ClientServiceDriver`

NewClientServiceDriver instantiates a new ClientServiceDriver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientServiceDriverWithDefaults

`func NewClientServiceDriverWithDefaults() *ClientServiceDriver`

NewClientServiceDriverWithDefaults instantiates a new ClientServiceDriver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientServiceDriver) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientServiceDriver) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientServiceDriver) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientServiceDriver) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ClientServiceDriver) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientServiceDriver) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientServiceDriver) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClientServiceDriver) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhone

`func (o *ClientServiceDriver) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClientServiceDriver) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClientServiceDriver) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClientServiceDriver) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCarModel

`func (o *ClientServiceDriver) GetCarModel() string`

GetCarModel returns the CarModel field if non-nil, zero value otherwise.

### GetCarModelOk

`func (o *ClientServiceDriver) GetCarModelOk() (*string, bool)`

GetCarModelOk returns a tuple with the CarModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarModel

`func (o *ClientServiceDriver) SetCarModel(v string)`

SetCarModel sets CarModel field to given value.

### HasCarModel

`func (o *ClientServiceDriver) HasCarModel() bool`

HasCarModel returns a boolean if a field has been set.

### GetLocation

`func (o *ClientServiceDriver) GetLocation() ClientServiceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClientServiceDriver) GetLocationOk() (*ClientServiceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClientServiceDriver) SetLocation(v ClientServiceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClientServiceDriver) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCarMake

`func (o *ClientServiceDriver) GetCarMake() string`

GetCarMake returns the CarMake field if non-nil, zero value otherwise.

### GetCarMakeOk

`func (o *ClientServiceDriver) GetCarMakeOk() (*string, bool)`

GetCarMakeOk returns a tuple with the CarMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarMake

`func (o *ClientServiceDriver) SetCarMake(v string)`

SetCarMake sets CarMake field to given value.

### HasCarMake

`func (o *ClientServiceDriver) HasCarMake() bool`

HasCarMake returns a boolean if a field has been set.

### GetCarNumber

`func (o *ClientServiceDriver) GetCarNumber() string`

GetCarNumber returns the CarNumber field if non-nil, zero value otherwise.

### GetCarNumberOk

`func (o *ClientServiceDriver) GetCarNumberOk() (*string, bool)`

GetCarNumberOk returns a tuple with the CarNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarNumber

`func (o *ClientServiceDriver) SetCarNumber(v string)`

SetCarNumber sets CarNumber field to given value.

### HasCarNumber

`func (o *ClientServiceDriver) HasCarNumber() bool`

HasCarNumber returns a boolean if a field has been set.

### GetCarColor

`func (o *ClientServiceDriver) GetCarColor() string`

GetCarColor returns the CarColor field if non-nil, zero value otherwise.

### GetCarColorOk

`func (o *ClientServiceDriver) GetCarColorOk() (*string, bool)`

GetCarColorOk returns a tuple with the CarColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarColor

`func (o *ClientServiceDriver) SetCarColor(v string)`

SetCarColor sets CarColor field to given value.

### HasCarColor

`func (o *ClientServiceDriver) HasCarColor() bool`

HasCarColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


