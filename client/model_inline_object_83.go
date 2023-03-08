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

// InlineObject83 struct for InlineObject83
type InlineObject83 struct {
	// Name of the Staged Upgrade Group. Length must be 1 to 255 characters
	Name string `json:"name"`
	// Description of the Staged Upgrade Group. Length must be 1 to 255 characters
	Description *string `json:"description,omitempty"`
	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault bool `json:"isDefault"`
	AssignedDevices *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 `json:"assignedDevices,omitempty"`
}

// NewInlineObject83 instantiates a new InlineObject83 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject83(name string, isDefault bool) *InlineObject83 {
	this := InlineObject83{}
	this.Name = name
	this.IsDefault = isDefault
	return &this
}

// NewInlineObject83WithDefaults instantiates a new InlineObject83 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject83WithDefaults() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject83) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject83) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject83) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject83) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject83) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value
func (o *InlineObject83) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *InlineObject83) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetAssignedDevices returns the AssignedDevices field value if set, zero value otherwise.
func (o *InlineObject83) GetAssignedDevices() NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 {
	if o == nil || isNil(o.AssignedDevices) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1
		return ret
	}
	return *o.AssignedDevices
}

// GetAssignedDevicesOk returns a tuple with the AssignedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetAssignedDevicesOk() (*NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1, bool) {
	if o == nil || isNil(o.AssignedDevices) {
    return nil, false
	}
	return o.AssignedDevices, true
}

// HasAssignedDevices returns a boolean if a field has been set.
func (o *InlineObject83) HasAssignedDevices() bool {
	if o != nil && !isNil(o.AssignedDevices) {
		return true
	}

	return false
}

// SetAssignedDevices gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 and assigns it to the AssignedDevices field.
func (o *InlineObject83) SetAssignedDevices(v NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) {
	o.AssignedDevices = &v
}

func (o InlineObject83) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.AssignedDevices) {
		toSerialize["assignedDevices"] = o.AssignedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject83 struct {
	value *InlineObject83
	isSet bool
}

func (v NullableInlineObject83) Get() *InlineObject83 {
	return v.value
}

func (v *NullableInlineObject83) Set(val *InlineObject83) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject83) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject83) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject83(val *InlineObject83) *NullableInlineObject83 {
	return &NullableInlineObject83{value: val, isSet: true}
}

func (v NullableInlineObject83) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject83) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


