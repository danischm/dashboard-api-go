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

// InlineResponse200119Networks struct for InlineResponse200119Networks
type InlineResponse200119Networks struct {
	// The network ID
	Id *string `json:"id,omitempty"`
	// The privilege of the SAML administrator on the network
	Access *string `json:"access,omitempty"`
}

// NewInlineResponse200119Networks instantiates a new InlineResponse200119Networks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200119Networks() *InlineResponse200119Networks {
	this := InlineResponse200119Networks{}
	return &this
}

// NewInlineResponse200119NetworksWithDefaults instantiates a new InlineResponse200119Networks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200119NetworksWithDefaults() *InlineResponse200119Networks {
	this := InlineResponse200119Networks{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200119Networks) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119Networks) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200119Networks) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200119Networks) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse200119Networks) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119Networks) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse200119Networks) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse200119Networks) SetAccess(v string) {
	o.Access = &v
}

func (o InlineResponse200119Networks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200119Networks struct {
	value *InlineResponse200119Networks
	isSet bool
}

func (v NullableInlineResponse200119Networks) Get() *InlineResponse200119Networks {
	return v.value
}

func (v *NullableInlineResponse200119Networks) Set(val *InlineResponse200119Networks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200119Networks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200119Networks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200119Networks(val *InlineResponse200119Networks) *NullableInlineResponse200119Networks {
	return &NullableInlineResponse200119Networks{value: val, isSet: true}
}

func (v NullableInlineResponse200119Networks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200119Networks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


