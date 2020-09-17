# \FotaApi

All URIs are relative to *https://api.lab5e.com/span*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearFirmwareError**](FotaApi.md#ClearFirmwareError) | **Delete** /collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
[**CreateFirmware**](FotaApi.md#CreateFirmware) | **Post** /collections/{collectionId}/firmware | Create firmware
[**DeleteFirmware**](FotaApi.md#DeleteFirmware) | **Delete** /collections/{collectionId}/firmware/{imageId} | Delete firmware
[**FirmwareUsage**](FotaApi.md#FirmwareUsage) | **Patch** /collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
[**ListFirmware**](FotaApi.md#ListFirmware) | **Get** /collections/{collectionId}/firmware | List firmware
[**RetrieveFirmware**](FotaApi.md#RetrieveFirmware) | **Get** /collections/{collectionId}/firmware/{imageId} | Retrieve firmware
[**UpdateFirmware**](FotaApi.md#UpdateFirmware) | **Patch** /collections/{collectionId}/firmware/{imageId} | Update firmware



## ClearFirmwareError

> map[string]interface{} ClearFirmwareError(ctx, collectionId, deviceId)

Clear FOTA error

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**deviceId** | **string**|  | 

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


## CreateFirmware

> Firmware CreateFirmware(ctx, collectionId, body)

Create firmware

Create a new firmware image. This is also invoked by the custom HTTP uploader if the POST uses multipart/formdata for the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**body** | [**CreateFirmwareRequest**](CreateFirmwareRequest.md)|  | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirmware

> Firmware DeleteFirmware(ctx, collectionId, imageId)

Delete firmware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**imageId** | **string**|  | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FirmwareUsage

> FirmwareUsageResponse FirmwareUsage(ctx, collectionId, imageId)

Firmware usage

Shows the firmware usage for devices in the collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**imageId** | **string**|  | 

### Return type

[**FirmwareUsageResponse**](FirmwareUsageResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirmware

> ListFirmwareResponse ListFirmware(ctx, collectionId)

List firmware

Lists available firmware images in collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 

### Return type

[**ListFirmwareResponse**](ListFirmwareResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirmware

> Firmware RetrieveFirmware(ctx, collectionId, imageId)

Retrieve firmware

Retrieve information on a firmware image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**|  | 
**imageId** | **string**|  | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirmware

> Firmware UpdateFirmware(ctx, collectionId, imageId, body)

Update firmware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| Collection ID  Collection ID for the collection owning the firmware image. | 
**imageId** | **string**| Firmware image ID | 
**body** | [**Firmware**](Firmware.md)|  | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

