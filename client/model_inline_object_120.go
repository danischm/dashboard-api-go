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

// InlineObject120 struct for InlineObject120
type InlineObject120 struct {
	// The name for your port schedule.
	Name *string `json:"name,omitempty"`
	PortSchedule *NetworksNetworkIdSwitchPortSchedulesPortSchedule `json:"portSchedule,omitempty"`
}

// NewInlineObject120 instantiates a new InlineObject120 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject120() *InlineObject120 {
	this := InlineObject120{}
	return &this
}

// NewInlineObject120WithDefaults instantiates a new InlineObject120 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject120WithDefaults() *InlineObject120 {
	this := InlineObject120{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject120) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject120) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject120) SetName(v string) {
	o.Name = &v
}

// GetPortSchedule returns the PortSchedule field value if set, zero value otherwise.
func (o *InlineObject120) GetPortSchedule() NetworksNetworkIdSwitchPortSchedulesPortSchedule {
	if o == nil || isNil(o.PortSchedule) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortSchedule
		return ret
	}
	return *o.PortSchedule
}

// GetPortScheduleOk returns a tuple with the PortSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetPortScheduleOk() (*NetworksNetworkIdSwitchPortSchedulesPortSchedule, bool) {
	if o == nil || isNil(o.PortSchedule) {
    return nil, false
	}
	return o.PortSchedule, true
}

// HasPortSchedule returns a boolean if a field has been set.
func (o *InlineObject120) HasPortSchedule() bool {
	if o != nil && !isNil(o.PortSchedule) {
		return true
	}

	return false
}

// SetPortSchedule gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortSchedule and assigns it to the PortSchedule field.
func (o *InlineObject120) SetPortSchedule(v NetworksNetworkIdSwitchPortSchedulesPortSchedule) {
	o.PortSchedule = &v
}

func (o InlineObject120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PortSchedule) {
		toSerialize["portSchedule"] = o.PortSchedule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject120 struct {
	value *InlineObject120
	isSet bool
}

func (v NullableInlineObject120) Get() *InlineObject120 {
	return v.value
}

func (v *NullableInlineObject120) Set(val *InlineObject120) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject120) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject120) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject120(val *InlineObject120) *NullableInlineObject120 {
	return &NullableInlineObject120{value: val, isSet: true}
}

func (v NullableInlineObject120) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject120) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


