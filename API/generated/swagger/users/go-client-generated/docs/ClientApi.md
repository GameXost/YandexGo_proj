# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientCancelRide**](ClientApi.md#ClientCancelRide) | **Post** /ride/{id}/cancel | 
[**ClientGetDriverInfo**](ClientApi.md#ClientGetDriverInfo) | **Get** /driver/{id} | 
[**ClientGetDriverLocation**](ClientApi.md#ClientGetDriverLocation) | **Get** /driver/{id}/location | сведения о водителе
[**ClientGetRideHistory**](ClientApi.md#ClientGetRideHistory) | **Get** /ride/history | 
[**ClientGetRideStatus**](ClientApi.md#ClientGetRideStatus) | **Get** /ride/{id} | 
[**ClientGetUserProfile**](ClientApi.md#ClientGetUserProfile) | **Get** /user/profile | операции с пользовательским профилем
[**ClientRequestRide**](ClientApi.md#ClientRequestRide) | **Post** /ride/request | операции с заказами
[**ClientUpdateUserProfile**](ClientApi.md#ClientUpdateUserProfile) | **Put** /user/profile | 

# **ClientCancelRide**
> ClientServiceStatusResponse ClientCancelRide(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ClientServiceStatusResponse**](client_serviceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientGetDriverInfo**
> ClientServiceDriver ClientGetDriverInfo(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ClientServiceDriver**](client_serviceDriver.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientGetDriverLocation**
> ClientServiceLocation ClientGetDriverLocation(ctx, id)
сведения о водителе

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ClientServiceLocation**](client_serviceLocation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientGetRideHistory**
> ClientServiceRideHistoryResponse ClientGetRideHistory(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientApiClientGetRideHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiClientGetRideHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 

### Return type

[**ClientServiceRideHistoryResponse**](client_serviceRideHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientGetRideStatus**
> ClientServiceRide ClientGetRideStatus(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ClientServiceRide**](client_serviceRide.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientGetUserProfile**
> ClientServiceUser ClientGetUserProfile(ctx, optional)
операции с пользовательским профилем

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientApiClientGetUserProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiClientGetUserProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**|  | 

### Return type

[**ClientServiceUser**](client_serviceUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientRequestRide**
> ClientServiceRide ClientRequestRide(ctx, body)
операции с заказами

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientServiceRideRequest**](ClientServiceRideRequest.md)|  | 

### Return type

[**ClientServiceRide**](client_serviceRide.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientUpdateUserProfile**
> ClientServiceUser ClientUpdateUserProfile(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientServiceUpdateProfileRequest**](ClientServiceUpdateProfileRequest.md)|  | 

### Return type

[**ClientServiceUser**](client_serviceUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

