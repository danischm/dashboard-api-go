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


// DhcpApiService DhcpApi service
type DhcpApiService service

type DhcpApiGetDeviceApplianceDhcpSubnetsRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	serial string
}

func (r DhcpApiGetDeviceApplianceDhcpSubnetsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceApplianceDhcpSubnetsExecute(r)
}

/*
GetDeviceApplianceDhcpSubnets Return the DHCP subnet information for an appliance

Return the DHCP subnet information for an appliance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return DhcpApiGetDeviceApplianceDhcpSubnetsRequest
*/
func (a *DhcpApiService) GetDeviceApplianceDhcpSubnets(ctx context.Context, serial string) DhcpApiGetDeviceApplianceDhcpSubnetsRequest {
	return DhcpApiGetDeviceApplianceDhcpSubnetsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DhcpApiService) GetDeviceApplianceDhcpSubnetsExecute(r DhcpApiGetDeviceApplianceDhcpSubnetsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetDeviceApplianceDhcpSubnets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/dhcp/subnets"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

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

type DhcpApiGetDeviceSwitchRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	serial string
	interfaceId string
}

func (r DhcpApiGetDeviceSwitchRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceSwitchRoutingInterfaceDhcpExecute(r)
}

/*
GetDeviceSwitchRoutingInterfaceDhcp Return a layer 3 interface DHCP configuration for a switch

Return a layer 3 interface DHCP configuration for a switch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param interfaceId Interface ID
 @return DhcpApiGetDeviceSwitchRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) GetDeviceSwitchRoutingInterfaceDhcp(ctx context.Context, serial string, interfaceId string) DhcpApiGetDeviceSwitchRoutingInterfaceDhcpRequest {
	return DhcpApiGetDeviceSwitchRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) GetDeviceSwitchRoutingInterfaceDhcpExecute(r DhcpApiGetDeviceSwitchRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetDeviceSwitchRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterValueToString(r.interfaceId, "interfaceId")), -1)

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

type DhcpApiGetNetworkCellularGatewayDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
}

func (r DhcpApiGetNetworkCellularGatewayDhcpRequest) Execute() (*GetNetworkCellularGatewayDhcp200Response, *http.Response, error) {
	return r.ApiService.GetNetworkCellularGatewayDhcpExecute(r)
}

/*
GetNetworkCellularGatewayDhcp List common DHCP settings of MGs

List common DHCP settings of MGs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpApiGetNetworkCellularGatewayDhcpRequest
*/
func (a *DhcpApiService) GetNetworkCellularGatewayDhcp(ctx context.Context, networkId string) DhcpApiGetNetworkCellularGatewayDhcpRequest {
	return DhcpApiGetNetworkCellularGatewayDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkCellularGatewayDhcp200Response
func (a *DhcpApiService) GetNetworkCellularGatewayDhcpExecute(r DhcpApiGetNetworkCellularGatewayDhcpRequest) (*GetNetworkCellularGatewayDhcp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkCellularGatewayDhcp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetNetworkCellularGatewayDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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

type DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	t0 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) T0(t0 string) DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.t0 = &t0
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) Timespan(timespan float32) DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) PerPage(perPage int32) DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) StartingAfter(startingAfter string) DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) EndingBefore(endingBefore string) DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) Execute() ([]GetNetworkSwitchDhcpV4ServersSeen200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpV4ServersSeenExecute(r)
}

/*
GetNetworkSwitchDhcpV4ServersSeen Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)

Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest
*/
func (a *DhcpApiService) GetNetworkSwitchDhcpV4ServersSeen(ctx context.Context, networkId string) DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest {
	return DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSwitchDhcpV4ServersSeen200ResponseInner
func (a *DhcpApiService) GetNetworkSwitchDhcpV4ServersSeenExecute(r DhcpApiGetNetworkSwitchDhcpV4ServersSeenRequest) ([]GetNetworkSwitchDhcpV4ServersSeen200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSwitchDhcpV4ServersSeen200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetNetworkSwitchDhcpV4ServersSeen")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcp/v4/servers/seen"
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

type DhcpApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	switchStackId string
	interfaceId string
}

func (r DhcpApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest) Execute() (*GetNetworkSwitchStackRoutingInterfaceDhcp200Response, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchStackRoutingInterfaceDhcpExecute(r)
}

/*
GetNetworkSwitchStackRoutingInterfaceDhcp Return a layer 3 interface DHCP configuration for a switch stack

Return a layer 3 interface DHCP configuration for a switch stack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param switchStackId Switch stack ID
 @param interfaceId Interface ID
 @return DhcpApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) GetNetworkSwitchStackRoutingInterfaceDhcp(ctx context.Context, networkId string, switchStackId string, interfaceId string) DhcpApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest {
	return DhcpApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		switchStackId: switchStackId,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchStackRoutingInterfaceDhcp200Response
func (a *DhcpApiService) GetNetworkSwitchStackRoutingInterfaceDhcpExecute(r DhcpApiGetNetworkSwitchStackRoutingInterfaceDhcpRequest) (*GetNetworkSwitchStackRoutingInterfaceDhcp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchStackRoutingInterfaceDhcp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.GetNetworkSwitchStackRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"switchStackId"+"}", url.PathEscape(parameterValueToString(r.switchStackId, "switchStackId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterValueToString(r.interfaceId, "interfaceId")), -1)

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

type DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	serial string
	interfaceId string
	updateDeviceSwitchRoutingInterfaceDhcpRequest *UpdateDeviceSwitchRoutingInterfaceDhcpRequest
}

func (r DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest) UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest UpdateDeviceSwitchRoutingInterfaceDhcpRequest) DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest {
	r.updateDeviceSwitchRoutingInterfaceDhcpRequest = &updateDeviceSwitchRoutingInterfaceDhcpRequest
	return r
}

func (r DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateDeviceSwitchRoutingInterfaceDhcpExecute(r)
}

/*
UpdateDeviceSwitchRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch

Update a layer 3 interface DHCP configuration for a switch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param interfaceId Interface ID
 @return DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) UpdateDeviceSwitchRoutingInterfaceDhcp(ctx context.Context, serial string, interfaceId string) DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest {
	return DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) UpdateDeviceSwitchRoutingInterfaceDhcpExecute(r DhcpApiUpdateDeviceSwitchRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.UpdateDeviceSwitchRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterValueToString(r.interfaceId, "interfaceId")), -1)

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
	localVarPostBody = r.updateDeviceSwitchRoutingInterfaceDhcpRequest
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

type DhcpApiUpdateNetworkCellularGatewayDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	updateNetworkCellularGatewayDhcpRequest *UpdateNetworkCellularGatewayDhcpRequest
}

func (r DhcpApiUpdateNetworkCellularGatewayDhcpRequest) UpdateNetworkCellularGatewayDhcpRequest(updateNetworkCellularGatewayDhcpRequest UpdateNetworkCellularGatewayDhcpRequest) DhcpApiUpdateNetworkCellularGatewayDhcpRequest {
	r.updateNetworkCellularGatewayDhcpRequest = &updateNetworkCellularGatewayDhcpRequest
	return r
}

func (r DhcpApiUpdateNetworkCellularGatewayDhcpRequest) Execute() (*GetNetworkCellularGatewayDhcp200Response, *http.Response, error) {
	return r.ApiService.UpdateNetworkCellularGatewayDhcpExecute(r)
}

/*
UpdateNetworkCellularGatewayDhcp Update common DHCP settings of MGs

Update common DHCP settings of MGs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpApiUpdateNetworkCellularGatewayDhcpRequest
*/
func (a *DhcpApiService) UpdateNetworkCellularGatewayDhcp(ctx context.Context, networkId string) DhcpApiUpdateNetworkCellularGatewayDhcpRequest {
	return DhcpApiUpdateNetworkCellularGatewayDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkCellularGatewayDhcp200Response
func (a *DhcpApiService) UpdateNetworkCellularGatewayDhcpExecute(r DhcpApiUpdateNetworkCellularGatewayDhcpRequest) (*GetNetworkCellularGatewayDhcp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkCellularGatewayDhcp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.UpdateNetworkCellularGatewayDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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
	localVarPostBody = r.updateNetworkCellularGatewayDhcpRequest
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

type DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct {
	ctx context.Context
	ApiService *DhcpApiService
	networkId string
	switchStackId string
	interfaceId string
	updateNetworkSwitchStackRoutingInterfaceDhcpRequest *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest
}

func (r DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	r.updateNetworkSwitchStackRoutingInterfaceDhcpRequest = &updateNetworkSwitchStackRoutingInterfaceDhcpRequest
	return r
}

func (r DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchStackRoutingInterfaceDhcpExecute(r)
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcp Update a layer 3 interface DHCP configuration for a switch stack

Update a layer 3 interface DHCP configuration for a switch stack

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param switchStackId Switch stack ID
 @param interfaceId Interface ID
 @return DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest
*/
func (a *DhcpApiService) UpdateNetworkSwitchStackRoutingInterfaceDhcp(ctx context.Context, networkId string, switchStackId string, interfaceId string) DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	return DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		switchStackId: switchStackId,
		interfaceId: interfaceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpApiService) UpdateNetworkSwitchStackRoutingInterfaceDhcpExecute(r DhcpApiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpApiService.UpdateNetworkSwitchStackRoutingInterfaceDhcp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"switchStackId"+"}", url.PathEscape(parameterValueToString(r.switchStackId, "switchStackId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interfaceId"+"}", url.PathEscape(parameterValueToString(r.interfaceId, "interfaceId")), -1)

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
	localVarPostBody = r.updateNetworkSwitchStackRoutingInterfaceDhcpRequest
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
