# \PassengersAPI

All URIs are relative to *http://localhost:9096*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPassengerInfo**](PassengersAPI.md#GetPassengerInfo) | **Get** /user/{id} | Get passenger information



## GetPassengerInfo

> DriverServiceUser GetPassengerInfo(ctx, id).Execute()

Get passenger information

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
	id := "id_example" // string | user's id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PassengersAPI.GetPassengerInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PassengersAPI.GetPassengerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPassengerInfo`: DriverServiceUser
	fmt.Fprintf(os.Stdout, "Response from `PassengersAPI.GetPassengerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | user&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassengerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriverServiceUser**](DriverServiceUser.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

