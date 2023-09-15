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

// checks if the UpdateNetworkApplianceSingleLanRequestMandatoryDhcp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSingleLanRequestMandatoryDhcp{}

// UpdateNetworkApplianceSingleLanRequestMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
type UpdateNetworkApplianceSingleLanRequestMandatoryDhcp struct {
	// Enable Mandatory DHCP on LAN.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkApplianceSingleLanRequestMandatoryDhcp instantiates a new UpdateNetworkApplianceSingleLanRequestMandatoryDhcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSingleLanRequestMandatoryDhcp() *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp {
	this := UpdateNetworkApplianceSingleLanRequestMandatoryDhcp{}
	return &this
}

// NewUpdateNetworkApplianceSingleLanRequestMandatoryDhcpWithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequestMandatoryDhcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSingleLanRequestMandatoryDhcpWithDefaults() *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp {
	this := UpdateNetworkApplianceSingleLanRequestMandatoryDhcp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp struct {
	value *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp
	isSet bool
}

func (v NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp) Get() *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp) Set(val *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp(val *UpdateNetworkApplianceSingleLanRequestMandatoryDhcp) *NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp {
	return &NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestMandatoryDhcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


