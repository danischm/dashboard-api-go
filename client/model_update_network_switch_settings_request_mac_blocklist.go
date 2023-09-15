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

// checks if the UpdateNetworkSwitchSettingsRequestMacBlocklist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchSettingsRequestMacBlocklist{}

// UpdateNetworkSwitchSettingsRequestMacBlocklist MAC blocklist
type UpdateNetworkSwitchSettingsRequestMacBlocklist struct {
	// Enable MAC blocklist
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkSwitchSettingsRequestMacBlocklist instantiates a new UpdateNetworkSwitchSettingsRequestMacBlocklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchSettingsRequestMacBlocklist() *UpdateNetworkSwitchSettingsRequestMacBlocklist {
	this := UpdateNetworkSwitchSettingsRequestMacBlocklist{}
	return &this
}

// NewUpdateNetworkSwitchSettingsRequestMacBlocklistWithDefaults instantiates a new UpdateNetworkSwitchSettingsRequestMacBlocklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchSettingsRequestMacBlocklistWithDefaults() *UpdateNetworkSwitchSettingsRequestMacBlocklist {
	this := UpdateNetworkSwitchSettingsRequestMacBlocklist{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchSettingsRequestMacBlocklist) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequestMacBlocklist) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchSettingsRequestMacBlocklist) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchSettingsRequestMacBlocklist) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkSwitchSettingsRequestMacBlocklist) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchSettingsRequestMacBlocklist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchSettingsRequestMacBlocklist struct {
	value *UpdateNetworkSwitchSettingsRequestMacBlocklist
	isSet bool
}

func (v NullableUpdateNetworkSwitchSettingsRequestMacBlocklist) Get() *UpdateNetworkSwitchSettingsRequestMacBlocklist {
	return v.value
}

func (v *NullableUpdateNetworkSwitchSettingsRequestMacBlocklist) Set(val *UpdateNetworkSwitchSettingsRequestMacBlocklist) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchSettingsRequestMacBlocklist) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchSettingsRequestMacBlocklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchSettingsRequestMacBlocklist(val *UpdateNetworkSwitchSettingsRequestMacBlocklist) *NullableUpdateNetworkSwitchSettingsRequestMacBlocklist {
	return &NullableUpdateNetworkSwitchSettingsRequestMacBlocklist{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchSettingsRequestMacBlocklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchSettingsRequestMacBlocklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


