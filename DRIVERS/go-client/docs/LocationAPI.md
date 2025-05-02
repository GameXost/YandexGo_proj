# \LocationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNearbyRequests**](LocationAPI.md#GetNearbyRequests) | **Get** /driver/nearby_req | Get nearby ride requests
[**UpdateLocation**](LocationAPI.md#UpdateLocation) | **Post** /driver/location | Update driver location (streaming)



## GetNearbyRequests

> DriverServiceRideRequestsResponse GetNearbyRequests(ctx).Latitude(latitude).Longitude(longitude).Execute()

Get nearby ride requests

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
	latitude := float64(1.2) // float64 | Latitude (optional)
	longitude := float64(1.2) // float64 | Longitude (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationAPI.GetNearbyRequests(context.Background()).Latitude(latitude).Longitude(longitude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationAPI.GetNearbyRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNearbyRequests`: DriverServiceRideRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `LocationAPI.GetNearbyRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNearbyRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **latitude** | **float64** | Latitude | 
 **longitude** | **float64** | Longitude | 

### Return type

[**DriverServiceRideRequestsResponse**](DriverServiceRideRequestsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> DriverServiceStatusResponse UpdateLocation(ctx).Body(body).Execute()

Update driver location (streaming)

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
	body := *openapiclient.NewDriverServiceLocationUpdateRequest() // DriverServiceLocationUpdateRequest |  (streaming inputs)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationAPI.UpdateLocation(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationAPI.UpdateLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocation`: DriverServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `LocationAPI.UpdateLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DriverServiceLocationUpdateRequest**](DriverServiceLocationUpdateRequest.md) |  (streaming inputs) | 

### Return type

[**DriverServiceStatusResponse**](DriverServiceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

