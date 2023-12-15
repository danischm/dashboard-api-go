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

// checks if the UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail{}

// UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail Email alert settings for DHCP servers
type UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail struct {
	// When enabled, send an email if a new DHCP server is seen. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail() *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail {
	this := UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail{}
	return &this
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmailWithDefaults instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmailWithDefaults() *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail {
	this := UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail struct {
	value *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail
	isSet bool
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) Get() *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) Set(val *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail(val *UpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail {
	return &NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestAlertsEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


