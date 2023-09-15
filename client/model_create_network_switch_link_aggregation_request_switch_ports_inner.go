/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner{}

// CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner struct for CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner
type CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner struct {
	// Serial number of the switch.
	Serial string `json:"serial"`
	// Port identifier of switch port. For modules, the identifier is \"SlotNumber_ModuleType_PortNumber\" (Ex: \"1_8X10G_1\"), otherwise it is just the port number (Ex: \"8\").
	PortId string `json:"portId"`
}

// NewCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner instantiates a new CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner(serial string, portId string) *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner {
	this := CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner{}
	this.Serial = serial
	this.PortId = portId
	return &this
}

// NewCreateNetworkSwitchLinkAggregationRequestSwitchPortsInnerWithDefaults instantiates a new CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchLinkAggregationRequestSwitchPortsInnerWithDefaults() *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner {
	this := CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner{}
	return &this
}

// GetSerial returns the Serial field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) SetSerial(v string) {
	o.Serial = v
}

// GetPortId returns the PortId field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) GetPortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) GetPortIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortId, true
}

// SetPortId sets field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) SetPortId(v string) {
	o.PortId = v
}

func (o CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serial"] = o.Serial
	toSerialize["portId"] = o.PortId
	return toSerialize, nil
}

type NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner struct {
	value *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner
	isSet bool
}

func (v NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) Get() *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner {
	return v.value
}

func (v *NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) Set(val *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner(val *CreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) *NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner {
	return &NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchLinkAggregationRequestSwitchPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


