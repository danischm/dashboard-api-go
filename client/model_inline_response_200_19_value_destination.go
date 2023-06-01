/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20019ValueDestination Destination of 'custom' type traffic filter
type InlineResponse20019ValueDestination struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address (e.g.\"192.168.10.1\", which is the same as \"192.168.10.1/32\"), or \"any\"
	Cidr *string `json:"cidr,omitempty"`
}

// NewInlineResponse20019ValueDestination instantiates a new InlineResponse20019ValueDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20019ValueDestination() *InlineResponse20019ValueDestination {
	this := InlineResponse20019ValueDestination{}
	return &this
}

// NewInlineResponse20019ValueDestinationWithDefaults instantiates a new InlineResponse20019ValueDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20019ValueDestinationWithDefaults() *InlineResponse20019ValueDestination {
	this := InlineResponse20019ValueDestination{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse20019ValueDestination) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20019ValueDestination) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse20019ValueDestination) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *InlineResponse20019ValueDestination) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse20019ValueDestination) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20019ValueDestination) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse20019ValueDestination) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse20019ValueDestination) SetCidr(v string) {
	o.Cidr = &v
}

func (o InlineResponse20019ValueDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20019ValueDestination struct {
	value *InlineResponse20019ValueDestination
	isSet bool
}

func (v NullableInlineResponse20019ValueDestination) Get() *InlineResponse20019ValueDestination {
	return v.value
}

func (v *NullableInlineResponse20019ValueDestination) Set(val *InlineResponse20019ValueDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20019ValueDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20019ValueDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20019ValueDestination(val *InlineResponse20019ValueDestination) *NullableInlineResponse20019ValueDestination {
	return &NullableInlineResponse20019ValueDestination{value: val, isSet: true}
}

func (v NullableInlineResponse20019ValueDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20019ValueDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

