# \DevicesApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /collections/{collectionId}/devices | Create device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /collections/{collectionId}/devices/{deviceId} | Remove device
[**ListDeviceData**](DevicesApi.md#ListDeviceData) | **Get** /collections/{collectionId}/devices/{deviceId}/data | Get payloads
[**ListDevices**](DevicesApi.md#ListDevices) | **Get** /collections/{collectionId}/devices | List devices
[**RetrieveDevice**](DevicesApi.md#RetrieveDevice) | **Get** /collections/{collectionId}/devices/{deviceId} | Retrieve device
[**SendMessage**](DevicesApi.md#SendMessage) | **Post** /collections/{collectionId}/devices/{deviceId}/to | Send message
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Patch** /collections/{existingCollectionId}/devices/{deviceId} | Update device



## CreateDevice

> Device CreateDevice(ctx, collectionId, body)

Create device

Create a new device. This will add a device to the collection. You must have write access to the collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| This is the containing collection | 
**body** | [**Device**](Device.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> Device DeleteDevice(ctx, collectionId, deviceId)

Remove device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**deviceId** | **string**|  | 

### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceData

> ListDataResponse ListDeviceData(ctx, collectionId, deviceId, optional)

Get payloads

List the data received from the device. Use the query parameters to control what data you retrieve.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| The collection ID. This is included in the request path. | 
**deviceId** | **string**| The device ID. This is included in the request path. | 
 **optional** | ***ListDeviceDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeviceDataOpts struct


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


## ListDevices

> ListDevicesResponse ListDevices(ctx, collectionId)

List devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 

### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDevice

> Device RetrieveDevice(ctx, collectionId, deviceId)

Retrieve device

Retrieve a single device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**deviceId** | **string**|  | 

### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> map[string]interface{} SendMessage(ctx, collectionId, deviceId, body)

Send message

Send a message to the device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**deviceId** | **string**|  | 
**body** | [**SendMessageRequest**](SendMessageRequest.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, existingCollectionId, deviceId, body)

Update device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**existingCollectionId** | **string**|  | 
**deviceId** | **string**|  | 
**body** | [**UpdateDeviceRequest**](UpdateDeviceRequest.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

