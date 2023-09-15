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


// ChannelUtilizationHistoryApiService ChannelUtilizationHistoryApi service
type ChannelUtilizationHistoryApiService service

type ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest struct {
	ctx context.Context
	ApiService *ChannelUtilizationHistoryApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	autoResolution *bool
	clientId *string
	deviceSerial *string
	apTag *string
	band *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) T0(t0 string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) T1(t1 string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) Timespan(timespan float32) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) Resolution(resolution int32) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.resolution = &resolution
	return r
}

// Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) AutoResolution(autoResolution bool) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.autoResolution = &autoResolution
	return r
}

// Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client&#39;s connection history.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) ClientId(clientId string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.clientId = &clientId
	return r
}

// Filter results by device to return AP channel utilization metrics for the queried device; either :band or :clientId must be jointly specified.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) DeviceSerial(deviceSerial string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.deviceSerial = &deviceSerial
	return r
}

// Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified.
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) ApTag(apTag string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.apTag = &apTag
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;).
func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) Band(band string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	r.band = &band
	return r
}

func (r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) Execute() ([]GetNetworkWirelessChannelUtilizationHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessChannelUtilizationHistoryExecute(r)
}

/*
GetNetworkWirelessChannelUtilizationHistory Return AP channel utilization over time for a device or network client

Return AP channel utilization over time for a device or network client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest
*/
func (a *ChannelUtilizationHistoryApiService) GetNetworkWirelessChannelUtilizationHistory(ctx context.Context, networkId string) ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest {
	return ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkWirelessChannelUtilizationHistory200ResponseInner
func (a *ChannelUtilizationHistoryApiService) GetNetworkWirelessChannelUtilizationHistoryExecute(r ChannelUtilizationHistoryApiGetNetworkWirelessChannelUtilizationHistoryRequest) ([]GetNetworkWirelessChannelUtilizationHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWirelessChannelUtilizationHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelUtilizationHistoryApiService.GetNetworkWirelessChannelUtilizationHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/channelUtilizationHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.resolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resolution", r.resolution, "")
	}
	if r.autoResolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoResolution", r.autoResolution, "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "")
	}
	if r.deviceSerial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceSerial", r.deviceSerial, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
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
