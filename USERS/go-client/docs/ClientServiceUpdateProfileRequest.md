# ClientServiceUpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | uniq id of user | [optional] 
**Username** | Pointer to **string** | user&#39;s firstname | [optional] 
**Phone** | Pointer to **string** | user&#39;s phone number | [optional] 
**Email** | Pointer to **string** | user&#39;s email address | [optional] 

## Methods

### NewClientServiceUpdateProfileRequest

`func NewClientServiceUpdateProfileRequest() *ClientServiceUpdateProfileRequest`

NewClientServiceUpdateProfileRequest instantiates a new ClientServiceUpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientServiceUpdateProfileRequestWithDefaults

`func NewClientServiceUpdateProfileRequestWithDefaults() *ClientServiceUpdateProfileRequest`

NewClientServiceUpdateProfileRequestWithDefaults instantiates a new ClientServiceUpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientServiceUpdateProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientServiceUpdateProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientServiceUpdateProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientServiceUpdateProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ClientServiceUpdateProfileRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientServiceUpdateProfileRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientServiceUpdateProfileRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClientServiceUpdateProfileRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhone

`func (o *ClientServiceUpdateProfileRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClientServiceUpdateProfileRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClientServiceUpdateProfileRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClientServiceUpdateProfileRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *ClientServiceUpdateProfileRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClientServiceUpdateProfileRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClientServiceUpdateProfileRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClientServiceUpdateProfileRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


