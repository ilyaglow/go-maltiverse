# \HostnameApi

All URIs are relative to *https://api.maltiverse.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHostname**](HostnameApi.md#AddHostname) | **Put** /hostname/{hostname} | adds a hostname item
[**DeleteHostname**](HostnameApi.md#DeleteHostname) | **Delete** /hostname/{hostname} | Deletes a hostname indicator. Only an admin can delete indicators.
[**GetHostname**](HostnameApi.md#GetHostname) | **Get** /hostname/{hostname} | retrieves an Honstame indicator


# **AddHostname**
> AddHostname(ctx, hostname, body)
adds a hostname item

Adds a hostname to the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostname** | **string**| ID of hotname address to return | 
  **body** | [**HostnameItem**](HostnameItem.md)| hostname object that needs to be added | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostname**
> HostnameItem DeleteHostname(ctx, hostname)
Deletes a hostname indicator. Only an admin can delete indicators.

Deletes a specified Hostname indicator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostname** | **string**| ID of hostname to return | 

### Return type

[**HostnameItem**](HostnameItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostname**
> HostnameItem GetHostname(ctx, hostname)
retrieves an Honstame indicator

Retrieves the specified hostname if exists 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostname** | **string**| ID of Hostname to return | 

### Return type

[**HostnameItem**](HostnameItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

