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

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	// Enable or disable warm spare for a switch
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare switch
	SpareSerial *string `json:"spareSerial,omitempty"`
}

// NewInlineObject23 instantiates a new InlineObject23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject23(enabled bool) *InlineObject23 {
	this := InlineObject23{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject23WithDefaults instantiates a new InlineObject23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject23WithDefaults() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject23) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject23) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineObject23) GetSpareSerial() string {
	if o == nil || isNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetSpareSerialOk() (*string, bool) {
	if o == nil || isNil(o.SpareSerial) {
    return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineObject23) HasSpareSerial() bool {
	if o != nil && !isNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineObject23) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

func (o InlineObject23) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject23 struct {
	value *InlineObject23
	isSet bool
}

func (v NullableInlineObject23) Get() *InlineObject23 {
	return v.value
}

func (v *NullableInlineObject23) Set(val *InlineObject23) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject23) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject23(val *InlineObject23) *NullableInlineObject23 {
	return &NullableInlineObject23{value: val, isSet: true}
}

func (v NullableInlineObject23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


