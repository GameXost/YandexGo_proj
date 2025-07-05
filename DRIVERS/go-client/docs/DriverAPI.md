# \DriverAPI

All URIs are relative to *http://localhost:9096*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDriverProfile**](DriverAPI.md#GetDriverProfile) | **Get** /driver/profile | Get driver profile
[**UpdateDriverProfile**](DriverAPI.md#UpdateDriverProfile) | **Put** /driver/profile | Update driver profile



## GetDriverProfile

> DriverServiceDriver GetDriverProfile(ctx).Execute()

Get driver profile

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriverAPI.GetDriverProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriverAPI.GetDriverProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriverProfile`: DriverServiceDriver
	fmt.Fprintf(os.Stdout, "Response from `DriverAPI.GetDriverProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriverProfileRequest struct via the builder pattern


### Return type

[**DriverServiceDriver**](DriverServiceDriver.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDriverProfile

> DriverServiceDriver UpdateDriverProfile(ctx).Body(body).Execute()

Update driver profile

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
	body := *openapiclient.NewDriverServiceUpdateDriverProfileRequest() // DriverServiceUpdateDriverProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriverAPI.UpdateDriverProfile(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriverAPI.UpdateDriverProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDriverProfile`: DriverServiceDriver
	fmt.Fprintf(os.Stdout, "Response from `DriverAPI.UpdateDriverProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriverProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DriverServiceUpdateDriverProfileRequest**](DriverServiceUpdateDriverProfileRequest.md) |  | 

### Return type

[**DriverServiceDriver**](DriverServiceDriver.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

