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

// InlineResponse2001InterfacesWan1SvisIpv4Nameservers The nameserver settings for this SVI.
type InlineResponse2001InterfacesWan1SvisIpv4Nameservers struct {
	// Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
	Addresses []string `json:"addresses,omitempty"`
}

// NewInlineResponse2001InterfacesWan1SvisIpv4Nameservers instantiates a new InlineResponse2001InterfacesWan1SvisIpv4Nameservers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001InterfacesWan1SvisIpv4Nameservers() *InlineResponse2001InterfacesWan1SvisIpv4Nameservers {
	this := InlineResponse2001InterfacesWan1SvisIpv4Nameservers{}
	return &this
}

// NewInlineResponse2001InterfacesWan1SvisIpv4NameserversWithDefaults instantiates a new InlineResponse2001InterfacesWan1SvisIpv4Nameservers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001InterfacesWan1SvisIpv4NameserversWithDefaults() *InlineResponse2001InterfacesWan1SvisIpv4Nameservers {
	this := InlineResponse2001InterfacesWan1SvisIpv4Nameservers{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InlineResponse2001InterfacesWan1SvisIpv4Nameservers) GetAddresses() []string {
	if o == nil || isNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001InterfacesWan1SvisIpv4Nameservers) GetAddressesOk() ([]string, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InlineResponse2001InterfacesWan1SvisIpv4Nameservers) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *InlineResponse2001InterfacesWan1SvisIpv4Nameservers) SetAddresses(v []string) {
	o.Addresses = v
}

func (o InlineResponse2001InterfacesWan1SvisIpv4Nameservers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers struct {
	value *InlineResponse2001InterfacesWan1SvisIpv4Nameservers
	isSet bool
}

func (v NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers) Get() *InlineResponse2001InterfacesWan1SvisIpv4Nameservers {
	return v.value
}

func (v *NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers) Set(val *InlineResponse2001InterfacesWan1SvisIpv4Nameservers) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers(val *InlineResponse2001InterfacesWan1SvisIpv4Nameservers) *NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers {
	return &NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers{value: val, isSet: true}
}

func (v NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001InterfacesWan1SvisIpv4Nameservers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


