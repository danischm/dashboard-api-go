/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceSingleLan200ResponseMandatoryDhcp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSingleLan200ResponseMandatoryDhcp{}

// GetNetworkApplianceSingleLan200ResponseMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
type GetNetworkApplianceSingleLan200ResponseMandatoryDhcp struct {
	// Enable Mandatory DHCP on single LAN.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkApplianceSingleLan200ResponseMandatoryDhcp instantiates a new GetNetworkApplianceSingleLan200ResponseMandatoryDhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSingleLan200ResponseMandatoryDhcp() *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp {
	this := GetNetworkApplianceSingleLan200ResponseMandatoryDhcp{}
	return &this
}

// NewGetNetworkApplianceSingleLan200ResponseMandatoryDhcpWithDefaults instantiates a new GetNetworkApplianceSingleLan200ResponseMandatoryDhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSingleLan200ResponseMandatoryDhcpWithDefaults() *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp {
	this := GetNetworkApplianceSingleLan200ResponseMandatoryDhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp struct {
	value *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp
	isSet bool
}

func (v NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp) Get() *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp {
	return v.value
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp) Set(val *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp(val *GetNetworkApplianceSingleLan200ResponseMandatoryDhcp) *NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp {
	return &NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseMandatoryDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


