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

// InlineObject24 struct for InlineObject24
type InlineObject24 struct {
	// Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Uuid *string `json:"uuid,omitempty"`
	// Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Major *int32 `json:"major,omitempty"`
	// Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor *int32 `json:"minor,omitempty"`
}

// NewInlineObject24 instantiates a new InlineObject24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject24() *InlineObject24 {
	this := InlineObject24{}
	return &this
}

// NewInlineObject24WithDefaults instantiates a new InlineObject24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject24WithDefaults() *InlineObject24 {
	this := InlineObject24{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InlineObject24) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InlineObject24) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InlineObject24) SetUuid(v string) {
	o.Uuid = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *InlineObject24) GetMajor() int32 {
	if o == nil || isNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetMajorOk() (*int32, bool) {
	if o == nil || isNil(o.Major) {
    return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *InlineObject24) HasMajor() bool {
	if o != nil && !isNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *InlineObject24) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *InlineObject24) GetMinor() int32 {
	if o == nil || isNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject24) GetMinorOk() (*int32, bool) {
	if o == nil || isNil(o.Minor) {
    return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *InlineObject24) HasMinor() bool {
	if o != nil && !isNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *InlineObject24) SetMinor(v int32) {
	o.Minor = &v
}

func (o InlineObject24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !isNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject24 struct {
	value *InlineObject24
	isSet bool
}

func (v NullableInlineObject24) Get() *InlineObject24 {
	return v.value
}

func (v *NullableInlineObject24) Set(val *InlineObject24) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject24) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject24(val *InlineObject24) *NullableInlineObject24 {
	return &NullableInlineObject24{value: val, isSet: true}
}

func (v NullableInlineObject24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


