# \Ipv4Api

All URIs are relative to *https://api.maltiverse.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIPv4**](Ipv4Api.md#AddIPv4) | **Put** /ip/{ipv4} | adds an ipv4 item
[**DeleteIP**](Ipv4Api.md#DeleteIP) | **Delete** /ip/{ipv4} | Deletes an IP indicator. Only an admin can delete indicators.
[**GetIP**](Ipv4Api.md#GetIP) | **Get** /ip/{ipv4} | retrieves an IP indicator


# **AddIPv4**
> AddIPv4(ctx, ipv4, body)
adds an ipv4 item

Adds an ipv4 to the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipv4** | **string**| ID of IPv4 address to return | 
  **body** | [**IPItem**](IPItem.md)| ipv4 object that needs to be added | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIP**
> IPItem DeleteIP(ctx, ipv4)
Deletes an IP indicator. Only an admin can delete indicators.

Deletes a specified IP address indicator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipv4** | **string**| ID of IPv4 address to return | 

### Return type

[**IPItem**](IPItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIP**
> IPItem GetIP(ctx, ipv4)
retrieves an IP indicator

Retrieves the specified ip if exists 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipv4** | **string**| ID of IPv4 address to return | 

### Return type

[**IPItem**](IPItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

