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

// InlineResponse20038NamedVlans A hash of Named VLANs options applied to the Network.
type InlineResponse20038NamedVlans struct {
	// Enables / disables Named VLANs on the Network.
	Enabled bool `json:"enabled"`
}

// NewInlineResponse20038NamedVlans instantiates a new InlineResponse20038NamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20038NamedVlans(enabled bool) *InlineResponse20038NamedVlans {
	this := InlineResponse20038NamedVlans{}
	this.Enabled = enabled
	return &this
}

// NewInlineResponse20038NamedVlansWithDefaults instantiates a new InlineResponse20038NamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20038NamedVlansWithDefaults() *InlineResponse20038NamedVlans {
	this := InlineResponse20038NamedVlans{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineResponse20038NamedVlans) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20038NamedVlans) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineResponse20038NamedVlans) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineResponse20038NamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20038NamedVlans struct {
	value *InlineResponse20038NamedVlans
	isSet bool
}

func (v NullableInlineResponse20038NamedVlans) Get() *InlineResponse20038NamedVlans {
	return v.value
}

func (v *NullableInlineResponse20038NamedVlans) Set(val *InlineResponse20038NamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20038NamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20038NamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20038NamedVlans(val *InlineResponse20038NamedVlans) *NullableInlineResponse20038NamedVlans {
	return &NullableInlineResponse20038NamedVlans{value: val, isSet: true}
}

func (v NullableInlineResponse20038NamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20038NamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


