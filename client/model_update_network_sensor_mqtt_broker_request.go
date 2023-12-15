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

// checks if the UpdateNetworkSensorMqttBrokerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSensorMqttBrokerRequest{}

// UpdateNetworkSensorMqttBrokerRequest struct for UpdateNetworkSensorMqttBrokerRequest
type UpdateNetworkSensorMqttBrokerRequest struct {
	// Set to true to enable MQTT broker for sensor network
	Enabled bool `json:"enabled"`
}

// NewUpdateNetworkSensorMqttBrokerRequest instantiates a new UpdateNetworkSensorMqttBrokerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSensorMqttBrokerRequest(enabled bool) *UpdateNetworkSensorMqttBrokerRequest {
	this := UpdateNetworkSensorMqttBrokerRequest{}
	this.Enabled = enabled
	return &this
}

// NewUpdateNetworkSensorMqttBrokerRequestWithDefaults instantiates a new UpdateNetworkSensorMqttBrokerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSensorMqttBrokerRequestWithDefaults() *UpdateNetworkSensorMqttBrokerRequest {
	this := UpdateNetworkSensorMqttBrokerRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateNetworkSensorMqttBrokerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSensorMqttBrokerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateNetworkSensorMqttBrokerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UpdateNetworkSensorMqttBrokerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSensorMqttBrokerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableUpdateNetworkSensorMqttBrokerRequest struct {
	value *UpdateNetworkSensorMqttBrokerRequest
	isSet bool
}

func (v NullableUpdateNetworkSensorMqttBrokerRequest) Get() *UpdateNetworkSensorMqttBrokerRequest {
	return v.value
}

func (v *NullableUpdateNetworkSensorMqttBrokerRequest) Set(val *UpdateNetworkSensorMqttBrokerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSensorMqttBrokerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSensorMqttBrokerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSensorMqttBrokerRequest(val *UpdateNetworkSensorMqttBrokerRequest) *NullableUpdateNetworkSensorMqttBrokerRequest {
	return &NullableUpdateNetworkSensorMqttBrokerRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSensorMqttBrokerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSensorMqttBrokerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


