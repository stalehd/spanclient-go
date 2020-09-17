# \OutputsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOutput**](OutputsApi.md#CreateOutput) | **Post** /collections/{collectionId}/outputs | Create output
[**DeleteOutput**](OutputsApi.md#DeleteOutput) | **Delete** /collections/{collectionId}/outputs/{outputId} | Delete output
[**ListOutputs**](OutputsApi.md#ListOutputs) | **Get** /collections/{collectionId}/outputs | List outputs
[**Logs**](OutputsApi.md#Logs) | **Get** /collections/{collectionId}/outputs/{outputId}/logs | Output logs
[**RetrieveOutput**](OutputsApi.md#RetrieveOutput) | **Get** /collections/{collectionId}/outputs/{outputId} | Retrieve output
[**Status**](OutputsApi.md#Status) | **Get** /collections/{collectionId}/outputs/{outputId}/status | Output status
[**UpdateOutput**](OutputsApi.md#UpdateOutput) | **Patch** /collections/{collectionId}/outputs/{outputId} | Update output



## CreateOutput

> Output CreateOutput(ctx, collectionId, body)

Create output

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**body** | [**Output**](Output.md)|  | 

### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutput

> Output DeleteOutput(ctx, collectionId, outputId)

Delete output

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**outputId** | **string**|  | 

### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutputs

> ListOutputResponse ListOutputs(ctx, collectionId)

List outputs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 

### Return type

[**ListOutputResponse**](ListOutputResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logs

> OutputLogs Logs(ctx, collectionId, outputId)

Output logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**outputId** | **string**|  | 

### Return type

[**OutputLogs**](OutputLogs.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOutput

> Output RetrieveOutput(ctx, collectionId, outputId)

Retrieve output

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**outputId** | **string**|  | 

### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> OutputStatus Status(ctx, collectionId, outputId)

Output status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**outputId** | **string**|  | 

### Return type

[**OutputStatus**](OutputStatus.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutput

> Output UpdateOutput(ctx, collectionId, outputId, body)

Update output

Running outputs will be restarted if required. Note that the collection ID can't be changed on an existing output.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**outputId** | **string**|  | 
**body** | [**Output**](Output.md)|  | 

### Return type

[**Output**](Output.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

