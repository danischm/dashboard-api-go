/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// PrefixesApiService PrefixesApi service
type PrefixesApiService service

type PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	networkId string
	createNetworkAppliancePrefixesDelegatedStatic *CreateNetworkAppliancePrefixesDelegatedStaticRequest
}

func (r PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest) CreateNetworkAppliancePrefixesDelegatedStatic(createNetworkAppliancePrefixesDelegatedStatic CreateNetworkAppliancePrefixesDelegatedStaticRequest) PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest {
	r.createNetworkAppliancePrefixesDelegatedStatic = &createNetworkAppliancePrefixesDelegatedStatic
	return r
}

func (r PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
CreateNetworkAppliancePrefixesDelegatedStatic Add a static delegated prefix from a network

Add a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *PrefixesApiService) CreateNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string) PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest {
	return PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PrefixesApiService) CreateNetworkAppliancePrefixesDelegatedStaticExecute(r PrefixesApiCreateNetworkAppliancePrefixesDelegatedStaticRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.CreateNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkAppliancePrefixesDelegatedStatic == nil {
		return localVarReturnValue, nil, reportError("createNetworkAppliancePrefixesDelegatedStatic is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createNetworkAppliancePrefixesDelegatedStatic
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PrefixesApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	networkId string
	staticDelegatedPrefixId string
}

func (r PrefixesApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
DeleteNetworkAppliancePrefixesDelegatedStatic Delete a static delegated prefix from a network

Delete a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param staticDelegatedPrefixId
 @return PrefixesApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *PrefixesApiService) DeleteNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) PrefixesApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest {
	return PrefixesApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
func (a *PrefixesApiService) DeleteNetworkAppliancePrefixesDelegatedStaticExecute(r PrefixesApiDeleteNetworkAppliancePrefixesDelegatedStaticRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.DeleteNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterToString(r.staticDelegatedPrefixId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PrefixesApiGetDeviceAppliancePrefixesDelegatedRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	serial string
}

func (r PrefixesApiGetDeviceAppliancePrefixesDelegatedRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceAppliancePrefixesDelegatedExecute(r)
}

/*
GetDeviceAppliancePrefixesDelegated Return current delegated IPv6 prefixes on an appliance.

Return current delegated IPv6 prefixes on an appliance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return PrefixesApiGetDeviceAppliancePrefixesDelegatedRequest
*/
func (a *PrefixesApiService) GetDeviceAppliancePrefixesDelegated(ctx context.Context, serial string) PrefixesApiGetDeviceAppliancePrefixesDelegatedRequest {
	return PrefixesApiGetDeviceAppliancePrefixesDelegatedRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PrefixesApiService) GetDeviceAppliancePrefixesDelegatedExecute(r PrefixesApiGetDeviceAppliancePrefixesDelegatedRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.GetDeviceAppliancePrefixesDelegated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/prefixes/delegated"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PrefixesApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	serial string
}

func (r PrefixesApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceAppliancePrefixesDelegatedVlanAssignmentsExecute(r)
}

/*
GetDeviceAppliancePrefixesDelegatedVlanAssignments Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return PrefixesApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest
*/
func (a *PrefixesApiService) GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx context.Context, serial string) PrefixesApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest {
	return PrefixesApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PrefixesApiService) GetDeviceAppliancePrefixesDelegatedVlanAssignmentsExecute(r PrefixesApiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.GetDeviceAppliancePrefixesDelegatedVlanAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/prefixes/delegated/vlanAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	networkId string
	staticDelegatedPrefixId string
}

func (r PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
GetNetworkAppliancePrefixesDelegatedStatic Return a static delegated prefix from a network

Return a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param staticDelegatedPrefixId
 @return PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *PrefixesApiService) GetNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticRequest {
	return PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
//  @return GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
func (a *PrefixesApiService) GetNetworkAppliancePrefixesDelegatedStaticExecute(r PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticRequest) (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.GetNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterToString(r.staticDelegatedPrefixId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	networkId string
}

func (r PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticsRequest) Execute() ([]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAppliancePrefixesDelegatedStaticsExecute(r)
}

/*
GetNetworkAppliancePrefixesDelegatedStatics List static delegated prefixes for a network

List static delegated prefixes for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticsRequest
*/
func (a *PrefixesApiService) GetNetworkAppliancePrefixesDelegatedStatics(ctx context.Context, networkId string) PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticsRequest {
	return PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
func (a *PrefixesApiService) GetNetworkAppliancePrefixesDelegatedStaticsExecute(r PrefixesApiGetNetworkAppliancePrefixesDelegatedStaticsRequest) ([]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.GetNetworkAppliancePrefixesDelegatedStatics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *PrefixesApiService
	networkId string
	staticDelegatedPrefixId string
	updateNetworkAppliancePrefixesDelegatedStatic *UpdateNetworkAppliancePrefixesDelegatedStaticRequest
}

func (r PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest) UpdateNetworkAppliancePrefixesDelegatedStatic(updateNetworkAppliancePrefixesDelegatedStatic UpdateNetworkAppliancePrefixesDelegatedStaticRequest) PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	r.updateNetworkAppliancePrefixesDelegatedStatic = &updateNetworkAppliancePrefixesDelegatedStatic
	return r
}

func (r PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
UpdateNetworkAppliancePrefixesDelegatedStatic Update a static delegated prefix from a network

Update a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param staticDelegatedPrefixId
 @return PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *PrefixesApiService) UpdateNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	return PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PrefixesApiService) UpdateNetworkAppliancePrefixesDelegatedStaticExecute(r PrefixesApiUpdateNetworkAppliancePrefixesDelegatedStaticRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrefixesApiService.UpdateNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterToString(r.staticDelegatedPrefixId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateNetworkAppliancePrefixesDelegatedStatic
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}