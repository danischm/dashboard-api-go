/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// [optional] If set to \"true\" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.
	Fullframe *bool `json:"fullframe,omitempty"`
}

// NewInlineObject4 instantiates a new InlineObject4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject4() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// NewInlineObject4WithDefaults instantiates a new InlineObject4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject4WithDefaults() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *InlineObject4) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InlineObject4) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *InlineObject4) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetFullframe returns the Fullframe field value if set, zero value otherwise.
func (o *InlineObject4) GetFullframe() bool {
	if o == nil || isNil(o.Fullframe) {
		var ret bool
		return ret
	}
	return *o.Fullframe
}

// GetFullframeOk returns a tuple with the Fullframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetFullframeOk() (*bool, bool) {
	if o == nil || isNil(o.Fullframe) {
    return nil, false
	}
	return o.Fullframe, true
}

// HasFullframe returns a boolean if a field has been set.
func (o *InlineObject4) HasFullframe() bool {
	if o != nil && !isNil(o.Fullframe) {
		return true
	}

	return false
}

// SetFullframe gets a reference to the given bool and assigns it to the Fullframe field.
func (o *InlineObject4) SetFullframe(v bool) {
	o.Fullframe = &v
}

func (o InlineObject4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Fullframe) {
		toSerialize["fullframe"] = o.Fullframe
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject4 struct {
	value *InlineObject4
	isSet bool
}

func (v NullableInlineObject4) Get() *InlineObject4 {
	return v.value
}

func (v *NullableInlineObject4) Set(val *InlineObject4) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject4) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject4(val *InlineObject4) *NullableInlineObject4 {
	return &NullableInlineObject4{value: val, isSet: true}
}

func (v NullableInlineObject4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


