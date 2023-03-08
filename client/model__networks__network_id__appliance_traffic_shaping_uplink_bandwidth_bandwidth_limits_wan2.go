/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 The bandwidth settings for the 'wan2' uplink
type NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2() *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2WithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2WithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2(val *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


