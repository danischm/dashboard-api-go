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

// checks if the UpdateNetworkWirelessSsidBonjourForwardingRequestException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidBonjourForwardingRequestException{}

// UpdateNetworkWirelessSsidBonjourForwardingRequestException Bonjour forwarding exception
type UpdateNetworkWirelessSsidBonjourForwardingRequestException struct {
	// If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkWirelessSsidBonjourForwardingRequestException instantiates a new UpdateNetworkWirelessSsidBonjourForwardingRequestException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidBonjourForwardingRequestException() *UpdateNetworkWirelessSsidBonjourForwardingRequestException {
	this := UpdateNetworkWirelessSsidBonjourForwardingRequestException{}
	return &this
}

// NewUpdateNetworkWirelessSsidBonjourForwardingRequestExceptionWithDefaults instantiates a new UpdateNetworkWirelessSsidBonjourForwardingRequestException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidBonjourForwardingRequestExceptionWithDefaults() *UpdateNetworkWirelessSsidBonjourForwardingRequestException {
	this := UpdateNetworkWirelessSsidBonjourForwardingRequestException{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequestException) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequestException) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequestException) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequestException) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkWirelessSsidBonjourForwardingRequestException) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidBonjourForwardingRequestException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException struct {
	value *UpdateNetworkWirelessSsidBonjourForwardingRequestException
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException) Get() *UpdateNetworkWirelessSsidBonjourForwardingRequestException {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException) Set(val *UpdateNetworkWirelessSsidBonjourForwardingRequestException) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidBonjourForwardingRequestException(val *UpdateNetworkWirelessSsidBonjourForwardingRequestException) *NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException {
	return &NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidBonjourForwardingRequestException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


