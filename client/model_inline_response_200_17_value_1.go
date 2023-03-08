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

// InlineResponse20017Value1 Value of traffic filter
type InlineResponse20017Value1 struct {
	// ID of 'applicationCategory' or 'application' type traffic filter
	Id *string `json:"id,omitempty"`
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source *InlineResponse20017Value1Source `json:"source,omitempty"`
	Destination *InlineResponse20017Value1Destination `json:"destination,omitempty"`
}

// NewInlineResponse20017Value1 instantiates a new InlineResponse20017Value1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017Value1() *InlineResponse20017Value1 {
	this := InlineResponse20017Value1{}
	return &this
}

// NewInlineResponse20017Value1WithDefaults instantiates a new InlineResponse20017Value1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017Value1WithDefaults() *InlineResponse20017Value1 {
	this := InlineResponse20017Value1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20017Value1) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Value1) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20017Value1) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20017Value1) SetId(v string) {
	o.Id = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20017Value1) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Value1) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20017Value1) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20017Value1) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InlineResponse20017Value1) GetSource() InlineResponse20017Value1Source {
	if o == nil || isNil(o.Source) {
		var ret InlineResponse20017Value1Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Value1) GetSourceOk() (*InlineResponse20017Value1Source, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InlineResponse20017Value1) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given InlineResponse20017Value1Source and assigns it to the Source field.
func (o *InlineResponse20017Value1) SetSource(v InlineResponse20017Value1Source) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *InlineResponse20017Value1) GetDestination() InlineResponse20017Value1Destination {
	if o == nil || isNil(o.Destination) {
		var ret InlineResponse20017Value1Destination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017Value1) GetDestinationOk() (*InlineResponse20017Value1Destination, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *InlineResponse20017Value1) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given InlineResponse20017Value1Destination and assigns it to the Destination field.
func (o *InlineResponse20017Value1) SetDestination(v InlineResponse20017Value1Destination) {
	o.Destination = &v
}

func (o InlineResponse20017Value1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017Value1 struct {
	value *InlineResponse20017Value1
	isSet bool
}

func (v NullableInlineResponse20017Value1) Get() *InlineResponse20017Value1 {
	return v.value
}

func (v *NullableInlineResponse20017Value1) Set(val *InlineResponse20017Value1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017Value1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017Value1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017Value1(val *InlineResponse20017Value1) *NullableInlineResponse20017Value1 {
	return &NullableInlineResponse20017Value1{value: val, isSet: true}
}

func (v NullableInlineResponse20017Value1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017Value1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


