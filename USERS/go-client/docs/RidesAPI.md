# \RidesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRide**](RidesAPI.md#CancelRide) | **Post** /ride/{id}/cancel | Cancel ride
[**GetCurrentRide**](RidesAPI.md#GetCurrentRide) | **Get** /ride/{id} | get ride information
[**GetRidesHistory**](RidesAPI.md#GetRidesHistory) | **Get** /ride/history | Get history of last rides
[**RequestRide**](RidesAPI.md#RequestRide) | **Post** /ride/request | request a ride



## CancelRide

> ClientServiceStatusResponse CancelRide(ctx, id).Execute()

Cancel ride

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GameXost/YandexGo_proj"
)

func main() {
	id := "id_example" // string | Ride unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.CancelRide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.CancelRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRide`: ClientServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.CancelRide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Ride unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientServiceStatusResponse**](ClientServiceStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentRide

> ClientServiceRide GetCurrentRide(ctx, id).Execute()

get ride information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GameXost/YandexGo_proj"
)

func main() {
	id := "id_example" // string | User unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.GetCurrentRide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.GetCurrentRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentRide`: ClientServiceRide
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.GetCurrentRide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientServiceRide**](ClientServiceRide.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRidesHistory

> ClientServiceRideHistoryResponse GetRidesHistory(ctx).Id(id).Execute()

Get history of last rides

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GameXost/YandexGo_proj"
)

func main() {
	id := "id_example" // string | User unique identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.GetRidesHistory(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.GetRidesHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRidesHistory`: ClientServiceRideHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.GetRidesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRidesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | User unique identifier | 

### Return type

[**ClientServiceRideHistoryResponse**](ClientServiceRideHistoryResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestRide

> ClientServiceRide RequestRide(ctx).Body(body).Execute()

request a ride

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GameXost/YandexGo_proj"
)

func main() {
	body := *openapiclient.NewClientServiceRideRequest() // ClientServiceRideRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.RequestRide(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.RequestRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestRide`: ClientServiceRide
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.RequestRide`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClientServiceRideRequest**](ClientServiceRideRequest.md) |  | 

### Return type

[**ClientServiceRide**](ClientServiceRide.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

