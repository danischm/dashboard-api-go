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

// InlineResponse20027Products The network devices to be updated
type InlineResponse20027Products struct {
	Switch *InlineResponse20027ProductsSwitch `json:"switch,omitempty"`
}

// NewInlineResponse20027Products instantiates a new InlineResponse20027Products object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027Products() *InlineResponse20027Products {
	this := InlineResponse20027Products{}
	return &this
}

// NewInlineResponse20027ProductsWithDefaults instantiates a new InlineResponse20027Products object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027ProductsWithDefaults() *InlineResponse20027Products {
	this := InlineResponse20027Products{}
	return &this
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *InlineResponse20027Products) GetSwitch() InlineResponse20027ProductsSwitch {
	if o == nil || isNil(o.Switch) {
		var ret InlineResponse20027ProductsSwitch
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027Products) GetSwitchOk() (*InlineResponse20027ProductsSwitch, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *InlineResponse20027Products) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given InlineResponse20027ProductsSwitch and assigns it to the Switch field.
func (o *InlineResponse20027Products) SetSwitch(v InlineResponse20027ProductsSwitch) {
	o.Switch = &v
}

func (o InlineResponse20027Products) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027Products struct {
	value *InlineResponse20027Products
	isSet bool
}

func (v NullableInlineResponse20027Products) Get() *InlineResponse20027Products {
	return v.value
}

func (v *NullableInlineResponse20027Products) Set(val *InlineResponse20027Products) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027Products) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027Products) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027Products(val *InlineResponse20027Products) *NullableInlineResponse20027Products {
	return &NullableInlineResponse20027Products{value: val, isSet: true}
}

func (v NullableInlineResponse20027Products) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027Products) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


