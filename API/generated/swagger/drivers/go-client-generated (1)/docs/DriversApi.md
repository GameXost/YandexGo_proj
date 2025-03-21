# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriversAcceptRide**](DriversApi.md#DriversAcceptRide) | **Post** /ride/{id}/accept | операции с заказами
[**DriversCancelRide**](DriversApi.md#DriversCancelRide) | **Post** /ride/{id}/cancel | 
[**DriversCompleteRide**](DriversApi.md#DriversCompleteRide) | **Post** /ride/{id}/complete | 
[**DriversGetCurrentRide**](DriversApi.md#DriversGetCurrentRide) | **Get** /driver/current_ride/{id} | 
[**DriversGetDriverProfile**](DriversApi.md#DriversGetDriverProfile) | **Get** /driver/profile | операции с профилем водителя
[**DriversGetNearbyRequests**](DriversApi.md#DriversGetNearbyRequests) | **Get** /driver/nearby_req | 
[**DriversGetPassengerInfo**](DriversApi.md#DriversGetPassengerInfo) | **Get** /user/{id} | сведения о пассажире
[**DriversUpdateDriverProfile**](DriversApi.md#DriversUpdateDriverProfile) | **Put** /driver/profile | 
[**DriversUpdateLocation**](DriversApi.md#DriversUpdateLocation) | **Post** /driver/location | операции с местоположением

# **DriversAcceptRide**
> DriverServiceStatusResponse DriversAcceptRide(ctx, id)
операции с заказами

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DriverServiceStatusResponse**](driver_serviceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversCancelRide**
> DriverServiceStatusResponse DriversCancelRide(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DriverServiceStatusResponse**](driver_serviceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversCompleteRide**
> DriverServiceStatusResponse DriversCompleteRide(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DriverServiceStatusResponse**](driver_serviceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversGetCurrentRide**
> DriverServiceRide DriversGetCurrentRide(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DriverServiceRide**](driver_serviceRide.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversGetDriverProfile**
> DriverServiceDriver DriversGetDriverProfile(ctx, optional)
операции с профилем водителя

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriversApiDriversGetDriverProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DriversApiDriversGetDriverProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**|  | 

### Return type

[**DriverServiceDriver**](driver_serviceDriver.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversGetNearbyRequests**
> DriverServiceRideRequestsResponse DriversGetNearbyRequests(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DriversApiDriversGetNearbyRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DriversApiDriversGetNearbyRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **latitude** | **optional.Float64**|  | 
 **longitude** | **optional.Float64**|  | 

### Return type

[**DriverServiceRideRequestsResponse**](driver_serviceRideRequestsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversGetPassengerInfo**
> DriverServiceUser DriversGetPassengerInfo(ctx, id)
сведения о пассажире

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DriverServiceUser**](driver_serviceUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversUpdateDriverProfile**
> DriverServiceDriver DriversUpdateDriverProfile(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DriverServiceUpdateDriverProfileRequest**](DriverServiceUpdateDriverProfileRequest.md)|  | 

### Return type

[**DriverServiceDriver**](driver_serviceDriver.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DriversUpdateLocation**
> DriverServiceStatusResponse DriversUpdateLocation(ctx, body)
операции с местоположением

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DriverServiceLocationUpdateRequest**](DriverServiceLocationUpdateRequest.md)|  (streaming inputs) | 

### Return type

[**DriverServiceStatusResponse**](driver_serviceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

