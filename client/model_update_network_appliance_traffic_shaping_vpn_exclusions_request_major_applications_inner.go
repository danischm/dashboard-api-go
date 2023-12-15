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

// checks if the UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner{}

// UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner struct for UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner
type UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner struct {
	// Application's Meraki ID.
	Id string `json:"id"`
	// Application's name.
	Name *string `json:"name,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner(id string) *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner{}
	this.Id = id
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) SetName(v string) {
	o.Name = &v
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner struct {
	value *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) Get() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) Set(val *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner(val *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner {
	return &NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


