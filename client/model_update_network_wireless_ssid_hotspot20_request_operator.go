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

// checks if the UpdateNetworkWirelessSsidHotspot20RequestOperator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidHotspot20RequestOperator{}

// UpdateNetworkWirelessSsidHotspot20RequestOperator Operator settings for this SSID
type UpdateNetworkWirelessSsidHotspot20RequestOperator struct {
	// Operator name
	Name *string `json:"name,omitempty"`
}

// NewUpdateNetworkWirelessSsidHotspot20RequestOperator instantiates a new UpdateNetworkWirelessSsidHotspot20RequestOperator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidHotspot20RequestOperator() *UpdateNetworkWirelessSsidHotspot20RequestOperator {
	this := UpdateNetworkWirelessSsidHotspot20RequestOperator{}
	return &this
}

// NewUpdateNetworkWirelessSsidHotspot20RequestOperatorWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20RequestOperator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidHotspot20RequestOperatorWithDefaults() *UpdateNetworkWirelessSsidHotspot20RequestOperator {
	this := UpdateNetworkWirelessSsidHotspot20RequestOperator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestOperator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestOperator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestOperator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestOperator) SetName(v string) {
	o.Name = &v
}

func (o UpdateNetworkWirelessSsidHotspot20RequestOperator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidHotspot20RequestOperator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidHotspot20RequestOperator struct {
	value *UpdateNetworkWirelessSsidHotspot20RequestOperator
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestOperator) Get() *UpdateNetworkWirelessSsidHotspot20RequestOperator {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestOperator) Set(val *UpdateNetworkWirelessSsidHotspot20RequestOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidHotspot20RequestOperator(val *UpdateNetworkWirelessSsidHotspot20RequestOperator) *NullableUpdateNetworkWirelessSsidHotspot20RequestOperator {
	return &NullableUpdateNetworkWirelessSsidHotspot20RequestOperator{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


