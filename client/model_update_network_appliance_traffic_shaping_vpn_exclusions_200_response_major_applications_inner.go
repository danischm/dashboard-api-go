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

// checks if the UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner{}

// UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner struct for UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner
type UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner struct {
	// Application's Meraki ID.
	Id string `json:"id"`
	// Application's name.
	Name string `json:"name"`
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner(id string, name string) *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) SetName(v string) {
	o.Name = v
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner struct {
	value *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) Get() *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) Set(val *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner(val *UpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner {
	return &NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusions200ResponseMajorApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


