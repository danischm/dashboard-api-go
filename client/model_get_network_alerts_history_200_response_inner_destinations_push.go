/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkAlertsHistory200ResponseInnerDestinationsPush push destinations for this alert
type GetNetworkAlertsHistory200ResponseInnerDestinationsPush struct {
	// time when the alert was sent to the user(s) for this channel
	SentAt *string `json:"sentAt,omitempty"`
}

// NewGetNetworkAlertsHistory200ResponseInnerDestinationsPush instantiates a new GetNetworkAlertsHistory200ResponseInnerDestinationsPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAlertsHistory200ResponseInnerDestinationsPush() *GetNetworkAlertsHistory200ResponseInnerDestinationsPush {
	this := GetNetworkAlertsHistory200ResponseInnerDestinationsPush{}
	return &this
}

// NewGetNetworkAlertsHistory200ResponseInnerDestinationsPushWithDefaults instantiates a new GetNetworkAlertsHistory200ResponseInnerDestinationsPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAlertsHistory200ResponseInnerDestinationsPushWithDefaults() *GetNetworkAlertsHistory200ResponseInnerDestinationsPush {
	this := GetNetworkAlertsHistory200ResponseInnerDestinationsPush{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsPush) GetSentAt() string {
	if o == nil || isNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsPush) GetSentAtOk() (*string, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsPush) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinationsPush) SetSentAt(v string) {
	o.SentAt = &v
}

func (o GetNetworkAlertsHistory200ResponseInnerDestinationsPush) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush struct {
	value *GetNetworkAlertsHistory200ResponseInnerDestinationsPush
	isSet bool
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush) Get() *GetNetworkAlertsHistory200ResponseInnerDestinationsPush {
	return v.value
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush) Set(val *GetNetworkAlertsHistory200ResponseInnerDestinationsPush) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush(val *GetNetworkAlertsHistory200ResponseInnerDestinationsPush) *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush {
	return &NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush{value: val, isSet: true}
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinationsPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

