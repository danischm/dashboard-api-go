/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// BluetoothClientsApiService BluetoothClientsApi service
type BluetoothClientsApiService service

type BluetoothClientsApiGetNetworkBluetoothClientRequest struct {
	ctx context.Context
	ApiService *BluetoothClientsApiService
	networkId string
	bluetoothClientId string
	includeConnectivityHistory *bool
	connectivityHistoryTimespan *int32
}

// Include the connectivity history for this client
func (r BluetoothClientsApiGetNetworkBluetoothClientRequest) IncludeConnectivityHistory(includeConnectivityHistory bool) BluetoothClientsApiGetNetworkBluetoothClientRequest {
	r.includeConnectivityHistory = &includeConnectivityHistory
	return r
}

// The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used.
func (r BluetoothClientsApiGetNetworkBluetoothClientRequest) ConnectivityHistoryTimespan(connectivityHistoryTimespan int32) BluetoothClientsApiGetNetworkBluetoothClientRequest {
	r.connectivityHistoryTimespan = &connectivityHistoryTimespan
	return r
}

func (r BluetoothClientsApiGetNetworkBluetoothClientRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkBluetoothClientExecute(r)
}

/*
GetNetworkBluetoothClient Return a Bluetooth client

Return a Bluetooth client. Bluetooth clients can be identified by their ID or their MAC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param bluetoothClientId Bluetooth client ID
 @return BluetoothClientsApiGetNetworkBluetoothClientRequest
*/
func (a *BluetoothClientsApiService) GetNetworkBluetoothClient(ctx context.Context, networkId string, bluetoothClientId string) BluetoothClientsApiGetNetworkBluetoothClientRequest {
	return BluetoothClientsApiGetNetworkBluetoothClientRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		bluetoothClientId: bluetoothClientId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *BluetoothClientsApiService) GetNetworkBluetoothClientExecute(r BluetoothClientsApiGetNetworkBluetoothClientRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BluetoothClientsApiService.GetNetworkBluetoothClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/bluetoothClients/{bluetoothClientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bluetoothClientId"+"}", url.PathEscape(parameterValueToString(r.bluetoothClientId, "bluetoothClientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeConnectivityHistory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeConnectivityHistory", r.includeConnectivityHistory, "")
	}
	if r.connectivityHistoryTimespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "connectivityHistoryTimespan", r.connectivityHistoryTimespan, "")
	}
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type BluetoothClientsApiGetNetworkBluetoothClientsRequest struct {
	ctx context.Context
	ApiService *BluetoothClientsApiService
	networkId string
	t0 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	includeConnectivityHistory *bool
}

// The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) T0(t0 string) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	r.t0 = &t0
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day.
func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) Timespan(timespan float32) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10.
func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) PerPage(perPage int32) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) StartingAfter(startingAfter string) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) EndingBefore(endingBefore string) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	r.endingBefore = &endingBefore
	return r
}

// Include the connectivity history for this client
func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) IncludeConnectivityHistory(includeConnectivityHistory bool) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	r.includeConnectivityHistory = &includeConnectivityHistory
	return r
}

func (r BluetoothClientsApiGetNetworkBluetoothClientsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkBluetoothClientsExecute(r)
}

/*
GetNetworkBluetoothClients List the Bluetooth clients seen by APs in this network

List the Bluetooth clients seen by APs in this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return BluetoothClientsApiGetNetworkBluetoothClientsRequest
*/
func (a *BluetoothClientsApiService) GetNetworkBluetoothClients(ctx context.Context, networkId string) BluetoothClientsApiGetNetworkBluetoothClientsRequest {
	return BluetoothClientsApiGetNetworkBluetoothClientsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *BluetoothClientsApiService) GetNetworkBluetoothClientsExecute(r BluetoothClientsApiGetNetworkBluetoothClientsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BluetoothClientsApiService.GetNetworkBluetoothClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/bluetoothClients"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.includeConnectivityHistory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeConnectivityHistory", r.includeConnectivityHistory, "")
	}
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
