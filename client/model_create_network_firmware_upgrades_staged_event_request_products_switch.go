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

// CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch The network device to be updated
type CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch struct {
	NextUpgrade *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchWithDefaults() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch {
	this := CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) GetNextUpgrade() CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) GetNextUpgradeOk() (*CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade and assigns it to the NextUpgrade field.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) SetNextUpgrade(v CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitchNextUpgrade) {
	o.NextUpgrade = &v
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch struct {
	value *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) Get() *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) Set(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch(val *CreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch {
	return &NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequestProductsSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

