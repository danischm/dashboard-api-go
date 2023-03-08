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

// InlineResponse200110 struct for InlineResponse200110
type InlineResponse200110 struct {
	// Resulting licenses from the move
	ResultingLicenses []InlineResponse200109 `json:"resultingLicenses,omitempty"`
}

// NewInlineResponse200110 instantiates a new InlineResponse200110 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110() *InlineResponse200110 {
	this := InlineResponse200110{}
	return &this
}

// NewInlineResponse200110WithDefaults instantiates a new InlineResponse200110 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110WithDefaults() *InlineResponse200110 {
	this := InlineResponse200110{}
	return &this
}

// GetResultingLicenses returns the ResultingLicenses field value if set, zero value otherwise.
func (o *InlineResponse200110) GetResultingLicenses() []InlineResponse200109 {
	if o == nil || isNil(o.ResultingLicenses) {
		var ret []InlineResponse200109
		return ret
	}
	return o.ResultingLicenses
}

// GetResultingLicensesOk returns a tuple with the ResultingLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetResultingLicensesOk() ([]InlineResponse200109, bool) {
	if o == nil || isNil(o.ResultingLicenses) {
    return nil, false
	}
	return o.ResultingLicenses, true
}

// HasResultingLicenses returns a boolean if a field has been set.
func (o *InlineResponse200110) HasResultingLicenses() bool {
	if o != nil && !isNil(o.ResultingLicenses) {
		return true
	}

	return false
}

// SetResultingLicenses gets a reference to the given []InlineResponse200109 and assigns it to the ResultingLicenses field.
func (o *InlineResponse200110) SetResultingLicenses(v []InlineResponse200109) {
	o.ResultingLicenses = v
}

func (o InlineResponse200110) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingLicenses) {
		toSerialize["resultingLicenses"] = o.ResultingLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110 struct {
	value *InlineResponse200110
	isSet bool
}

func (v NullableInlineResponse200110) Get() *InlineResponse200110 {
	return v.value
}

func (v *NullableInlineResponse200110) Set(val *InlineResponse200110) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110(val *InlineResponse200110) *NullableInlineResponse200110 {
	return &NullableInlineResponse200110{value: val, isSet: true}
}

func (v NullableInlineResponse200110) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

