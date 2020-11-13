# \CollectionsApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastMessage**](CollectionsApi.md#BroadcastMessage) | **Post** /collections/{collectionId}/to | Broadcast message
[**CreateCollection**](CollectionsApi.md#CreateCollection) | **Post** /collections | Create collection
[**DeleteCollection**](CollectionsApi.md#DeleteCollection) | **Delete** /collections/{collectionId} | Delete collection
[**ListCollectionData**](CollectionsApi.md#ListCollectionData) | **Get** /collections/{collectionId}/data | Get payloads
[**ListCollections**](CollectionsApi.md#ListCollections) | **Get** /collections | List collections
[**RetrieveCollection**](CollectionsApi.md#RetrieveCollection) | **Get** /collections/{collectionId} | Retrieve collection
[**UpdateCollection**](CollectionsApi.md#UpdateCollection) | **Patch** /collections/{collectionId} | Update collection



## BroadcastMessage

> MultiSendMessageResponse BroadcastMessage(ctx, collectionId, body)

Broadcast message

Broadcast a message to all devices in the collection. This request will always succeed if the collection exists, even if there are one or more send errors. Individual errors are returned as an array of error messages in the response. Use equivalent to resource for devices to send a message to single device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**body** | [**SendMessageRequest**](SendMessageRequest.md)|  | 

### Return type

[**MultiSendMessageResponse**](MultiSendMessageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> Collection CreateCollection(ctx, body)

Create collection

The returned collection is the collection stored in the backend. Defaults have been set. There are no required fields in a collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Collection**](Collection.md)|  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> Collection DeleteCollection(ctx, collectionId)

Delete collection

You must have write access to the collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| The ID of the collection you want to delete | 

### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionData

> ListDataResponse ListCollectionData(ctx, collectionId, optional)

Get payloads

List the data received from all the devices in the collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| The collection ID requested. This is included in the request path. | 
 **optional** | ***ListCollectionDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCollectionDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limit the number of payloads to return. The default is 512. | 
 **start** | **optional.String**| Start of time range. The default is 24 hours ago. Value is in milliseconds since epoch. | 
 **end** | **optional.String**| End of time range. The default is the current time stamp. Value is in milliseconds since epoch. | 

### Return type

[**ListDataResponse**](ListDataResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> ListCollectionResponse ListCollections(ctx, )

List collections

Lists all the collections that one of your teams owns.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ListCollectionResponse**](ListCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCollection

> Collection RetrieveCollection(ctx, collectionId)

Retrieve collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| The collection ID of the collection you are requesting | 

### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, collectionId, body)

Update collection

You must have write access to the collection, ie. you must administer it

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| The ID of the collection. This is assigned by the backend. | 
**body** | [**Collection**](Collection.md)|  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

