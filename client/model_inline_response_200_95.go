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

// InlineResponse20095 struct for InlineResponse20095
type InlineResponse20095 struct {
	// Switch profile id
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// Switch profile name
	Name *string `json:"name,omitempty"`
	// Switch model
	Model *string `json:"model,omitempty"`
}

// NewInlineResponse20095 instantiates a new InlineResponse20095 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095() *InlineResponse20095 {
	this := InlineResponse20095{}
	return &this
}

// NewInlineResponse20095WithDefaults instantiates a new InlineResponse20095 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095WithDefaults() *InlineResponse20095 {
	this := InlineResponse20095{}
	return &this
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *InlineResponse20095) GetSwitchProfileId() string {
	if o == nil || isNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.SwitchProfileId) {
    return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *InlineResponse20095) HasSwitchProfileId() bool {
	if o != nil && !isNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *InlineResponse20095) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20095) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20095) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20095) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse20095) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse20095) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse20095) SetModel(v string) {
	o.Model = &v
}

func (o InlineResponse20095) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SwitchProfileId) {
		toSerialize["switchProfileId"] = o.SwitchProfileId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095 struct {
	value *InlineResponse20095
	isSet bool
}

func (v NullableInlineResponse20095) Get() *InlineResponse20095 {
	return v.value
}

func (v *NullableInlineResponse20095) Set(val *InlineResponse20095) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095(val *InlineResponse20095) *NullableInlineResponse20095 {
	return &NullableInlineResponse20095{value: val, isSet: true}
}

func (v NullableInlineResponse20095) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

