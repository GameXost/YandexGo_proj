# \DriverAPI

All URIs are relative to *http://localhost:9093*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDriverInfo**](DriverAPI.md#GetDriverInfo) | **Get** /driver/{id} | get info about your driver



## GetDriverInfo

> ClientServiceDriver GetDriverInfo(ctx, id).Execute()

get info about your driver

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
	resp, r, err := apiClient.DriverAPI.GetDriverInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriverAPI.GetDriverInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriverInfo`: ClientServiceDriver
	fmt.Fprintf(os.Stdout, "Response from `DriverAPI.GetDriverInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Driver unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriverInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientServiceDriver**](ClientServiceDriver.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

