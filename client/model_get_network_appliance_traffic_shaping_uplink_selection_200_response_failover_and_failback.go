/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback{}

// GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback WAN failover and failback
type GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback struct {
	Immediate *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback{}
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) GetImmediate() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate {
	if o == nil || IsNil(o.Immediate) {
		var ret GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) GetImmediateOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate, bool) {
	if o == nil || IsNil(o.Immediate) {
		return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) HasImmediate() bool {
	if o != nil && !IsNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) SetImmediate(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


