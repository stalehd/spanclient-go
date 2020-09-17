# \TokensApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensApi.md#CreateToken) | **Post** /tokens | Create token
[**DeleteToken**](TokensApi.md#DeleteToken) | **Delete** /tokens/{token} | Remove token
[**ListTokens**](TokensApi.md#ListTokens) | **Get** /tokens | List tokens
[**RetrieveToken**](TokensApi.md#RetrieveToken) | **Get** /tokens/{token} | Retrieve token
[**UpdateToken**](TokensApi.md#UpdateToken) | **Patch** /tokens/{token} | Update token



## CreateToken

> Token CreateToken(ctx, body)

Create token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Token**](Token.md)|  | 

### Return type

[**Token**](Token.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> map[string]interface{} DeleteToken(ctx, token)

Remove token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| The token to delete. | 

### Return type

**map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> TokenList ListTokens(ctx, )

List tokens

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TokenList**](TokenList.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveToken

> Token RetrieveToken(ctx, token)

Retrieve token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**|  | 

### Return type

[**Token**](Token.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> Token UpdateToken(ctx, token, body)

Update token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| The token  Use this in the &#x60;X-API-Token&#x60; header when using the API. | 
**body** | [**Token**](Token.md)|  | 

### Return type

[**Token**](Token.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

