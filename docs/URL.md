# \URLApi

All URIs are relative to *https://api.maltiverse.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddURL**](URLApi.md#AddURL) | **Put** /url/{urlchecksum} | adds a url item
[**DeleteURL**](URLApi.md#DeleteURL) | **Delete** /url/{urlchecksum} | Deletes a url indicator. Only an admin can delete indicators.
[**GetURL**](URLApi.md#GetURL) | **Get** /url/{urlchecksum} | retrieves an URL indicator


# **AddURL**
> AddURL(ctx, urlchecksum, body)
adds a url item

Adds a url to the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urlchecksum** | **string**| ID of URL to return. Must be calculated by the sha256 hash of the effective URL. | 
  **body** | [**URLItem**](URLItem.md)| URL object that needs to be added | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteURL**
> URLItem DeleteURL(ctx, urlchecksum)
Deletes a url indicator. Only an admin can delete indicators.

Deletes a specified URL indicator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urlchecksum** | **string**| ID of URL to delete. Must be calculated by the sha256 hash of the effective URL. | 

### Return type

[**URLItem**](URLItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetURL**
> URLItem GetURL(ctx, urlchecksum)
retrieves an URL indicator

Retrieves the specified url if exists fetched by its sha256 integrity checksum 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **urlchecksum** | **string**| ID of URL to return. Must be calculated by the sha256 hash of the effective URL | 

### Return type

[**URLItem**](URLItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

