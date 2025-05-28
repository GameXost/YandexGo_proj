# \RidesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptRide**](RidesAPI.md#AcceptRide) | **Post** /ride/{id}/accept | Accept a ride
[**CancelRide**](RidesAPI.md#CancelRide) | **Post** /ride/{id}/cancel | Cancel a ride
[**CompleteRide**](RidesAPI.md#CompleteRide) | **Post** /ride/{id}/complete | Complete a ride
[**GetCurrentRide**](RidesAPI.md#GetCurrentRide) | **Get** /driver/current_ride/{id} | Get current ride information
[**GetRideHistory**](RidesAPI.md#GetRideHistory) | **Get** /driver/{id}/rides | Get driver&#39;s ride history



## AcceptRide

> DriverServiceStatusResponse AcceptRide(ctx, id).Execute()

Accept a ride

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
	id := "id_example" // string | Ride ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.AcceptRide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.AcceptRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptRide`: DriverServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.AcceptRide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Ride ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverServiceStatusResponse**](DriverServiceStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelRide

> DriverServiceStatusResponse CancelRide(ctx, id).Execute()

Cancel a ride

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
	id := "id_example" // string | Ride ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.CancelRide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.CancelRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRide`: DriverServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.CancelRide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Ride ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverServiceStatusResponse**](DriverServiceStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteRide

> DriverServiceStatusResponse CompleteRide(ctx, id).Execute()

Complete a ride

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
	id := "id_example" // string | Ride ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.CompleteRide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.CompleteRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteRide`: DriverServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.CompleteRide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Ride ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverServiceStatusResponse**](DriverServiceStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentRide

> DriverServiceRide GetCurrentRide(ctx, id).Execute()

Get current ride information

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
	id := "id_example" // string | driver's id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.GetCurrentRide(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.GetCurrentRide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentRide`: DriverServiceRide
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.GetCurrentRide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | driver&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverServiceRide**](DriverServiceRide.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRideHistory

> DriverServiceRideHistoryResponse GetRideHistory(ctx, id).Execute()

Get driver's ride history

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
	id := "id_example" // string | driver's id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RidesAPI.GetRideHistory(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RidesAPI.GetRideHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRideHistory`: DriverServiceRideHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RidesAPI.GetRideHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | driver&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRideHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverServiceRideHistoryResponse**](DriverServiceRideHistoryResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

