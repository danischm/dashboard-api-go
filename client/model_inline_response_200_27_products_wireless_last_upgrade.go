/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20027ProductsWirelessLastUpgrade Details of the last firmware upgrade on the device
type InlineResponse20027ProductsWirelessLastUpgrade struct {
	// Timestamp of the last successful firmware upgrade
	Time *time.Time `json:"time,omitempty"`
	FromVersion *InlineResponse20027ProductsWirelessLastUpgradeFromVersion `json:"fromVersion,omitempty"`
	ToVersion *InlineResponse20027ProductsWirelessLastUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewInlineResponse20027ProductsWirelessLastUpgrade instantiates a new InlineResponse20027ProductsWirelessLastUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027ProductsWirelessLastUpgrade() *InlineResponse20027ProductsWirelessLastUpgrade {
	this := InlineResponse20027ProductsWirelessLastUpgrade{}
	return &this
}

// NewInlineResponse20027ProductsWirelessLastUpgradeWithDefaults instantiates a new InlineResponse20027ProductsWirelessLastUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027ProductsWirelessLastUpgradeWithDefaults() *InlineResponse20027ProductsWirelessLastUpgrade {
	this := InlineResponse20027ProductsWirelessLastUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) SetTime(v time.Time) {
	o.Time = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) GetFromVersion() InlineResponse20027ProductsWirelessLastUpgradeFromVersion {
	if o == nil || isNil(o.FromVersion) {
		var ret InlineResponse20027ProductsWirelessLastUpgradeFromVersion
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) GetFromVersionOk() (*InlineResponse20027ProductsWirelessLastUpgradeFromVersion, bool) {
	if o == nil || isNil(o.FromVersion) {
    return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) HasFromVersion() bool {
	if o != nil && !isNil(o.FromVersion) {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given InlineResponse20027ProductsWirelessLastUpgradeFromVersion and assigns it to the FromVersion field.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) SetFromVersion(v InlineResponse20027ProductsWirelessLastUpgradeFromVersion) {
	o.FromVersion = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) GetToVersion() InlineResponse20027ProductsWirelessLastUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret InlineResponse20027ProductsWirelessLastUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) GetToVersionOk() (*InlineResponse20027ProductsWirelessLastUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given InlineResponse20027ProductsWirelessLastUpgradeToVersion and assigns it to the ToVersion field.
func (o *InlineResponse20027ProductsWirelessLastUpgrade) SetToVersion(v InlineResponse20027ProductsWirelessLastUpgradeToVersion) {
	o.ToVersion = &v
}

func (o InlineResponse20027ProductsWirelessLastUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.FromVersion) {
		toSerialize["fromVersion"] = o.FromVersion
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027ProductsWirelessLastUpgrade struct {
	value *InlineResponse20027ProductsWirelessLastUpgrade
	isSet bool
}

func (v NullableInlineResponse20027ProductsWirelessLastUpgrade) Get() *InlineResponse20027ProductsWirelessLastUpgrade {
	return v.value
}

func (v *NullableInlineResponse20027ProductsWirelessLastUpgrade) Set(val *InlineResponse20027ProductsWirelessLastUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027ProductsWirelessLastUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027ProductsWirelessLastUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027ProductsWirelessLastUpgrade(val *InlineResponse20027ProductsWirelessLastUpgrade) *NullableInlineResponse20027ProductsWirelessLastUpgrade {
	return &NullableInlineResponse20027ProductsWirelessLastUpgrade{value: val, isSet: true}
}

func (v NullableInlineResponse20027ProductsWirelessLastUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027ProductsWirelessLastUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

