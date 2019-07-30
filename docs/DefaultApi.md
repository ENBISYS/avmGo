# \DefaultApi

All URIs are relative to *https://avm.enbisys.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFastValuation**](DefaultApi.md#GetFastValuation) | **Post** /properties/getFastValuation | 
[**GetValuation**](DefaultApi.md#GetValuation) | **Post** /properties/getValuation | 


## GetFastValuation

> int32 GetFastValuation(ctx, propertyFeatures)


Get only property price valuation without confidence estimation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyFeatures** | [**PropertyFeatures**](PropertyFeatures.md)| Property features that describe property | 

### Return type

**int32**

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValuation

> Valuation GetValuation(ctx, propertyFeatures)


Get property price valuation with confidence estimation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyFeatures** | [**PropertyFeatures**](PropertyFeatures.md)| Property features to valuate | 

### Return type

[**Valuation**](Valuation.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

