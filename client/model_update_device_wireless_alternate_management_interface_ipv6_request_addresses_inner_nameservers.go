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

// checks if the UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers{}

// UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers The DNS servers settings for this address.
type UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers struct {
	// Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
	Addresses []string `json:"addresses,omitempty"`
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers{}
	return &this
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameserversWithDefaults instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameserversWithDefaults() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) GetAddresses() []string {
	if o == nil || IsNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) GetAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) SetAddresses(v []string) {
	o.Addresses = v
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return toSerialize, nil
}

type NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers struct {
	value *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers
	isSet bool
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) Get() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers {
	return v.value
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) Set(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers {
	return &NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


