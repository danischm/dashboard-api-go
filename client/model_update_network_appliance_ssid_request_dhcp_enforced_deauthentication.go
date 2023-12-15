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

// checks if the UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication{}

// UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication DHCP Enforced Deauthentication enables the disassociation of wireless clients in addition to Mandatory DHCP. This param is only valid on firmware versions >= MX 17.0 where the associated LAN has Mandatory DHCP Enabled 
type UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication struct {
	// Enable DCHP Enforced Deauthentication on the SSID.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication instantiates a new UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication() *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication {
	this := UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication{}
	return &this
}

// NewUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthenticationWithDefaults instantiates a new UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthenticationWithDefaults() *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication {
	this := UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication struct {
	value *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication
	isSet bool
}

func (v NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) Get() *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) Set(val *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication(val *UpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) *NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication {
	return &NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSsidRequestDhcpEnforcedDeauthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


