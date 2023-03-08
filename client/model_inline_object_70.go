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

// InlineObject70 struct for InlineObject70
type InlineObject70 struct {
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"bandwidthLimits,omitempty"`
}

// NewInlineObject70 instantiates a new InlineObject70 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject70() *InlineObject70 {
	this := InlineObject70{}
	return &this
}

// NewInlineObject70WithDefaults instantiates a new InlineObject70 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject70WithDefaults() *InlineObject70 {
	this := InlineObject70{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineObject70) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject70) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineObject70) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular and assigns it to the BandwidthLimits field.
func (o *InlineObject70) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular) {
	o.BandwidthLimits = &v
}

func (o InlineObject70) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject70 struct {
	value *InlineObject70
	isSet bool
}

func (v NullableInlineObject70) Get() *InlineObject70 {
	return v.value
}

func (v *NullableInlineObject70) Set(val *InlineObject70) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject70) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject70) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject70(val *InlineObject70) *NullableInlineObject70 {
	return &NullableInlineObject70{value: val, isSet: true}
}

func (v NullableInlineObject70) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject70) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


