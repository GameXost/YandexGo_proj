# \LocationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDriverLocation**](LocationAPI.md#GetDriverLocation) | **Get** /driver/{id}/location | getting driver location lat, lon



## GetDriverLocation

> ClientServiceLocation GetDriverLocation(ctx, id).Execute()

getting driver location lat, lon

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
	id := "id_example" // string | Driver unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationAPI.GetDriverLocation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationAPI.GetDriverLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriverLocation`: ClientServiceLocation
	fmt.Fprintf(os.Stdout, "Response from `LocationAPI.GetDriverLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Driver unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriverLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientServiceLocation**](ClientServiceLocation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

