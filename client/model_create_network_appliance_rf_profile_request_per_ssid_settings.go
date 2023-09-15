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

// checks if the CreateNetworkApplianceRfProfileRequestPerSsidSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkApplianceRfProfileRequestPerSsidSettings{}

// CreateNetworkApplianceRfProfileRequestPerSsidSettings Per-SSID radio settings by number.
type CreateNetworkApplianceRfProfileRequestPerSsidSettings struct {
	Var1 *CreateNetworkApplianceRfProfileRequestPerSsidSettings1 `json:"1,omitempty"`
	Var2 *CreateNetworkApplianceRfProfileRequestPerSsidSettings2 `json:"2,omitempty"`
	Var3 *CreateNetworkApplianceRfProfileRequestPerSsidSettings3 `json:"3,omitempty"`
	Var4 *CreateNetworkApplianceRfProfileRequestPerSsidSettings4 `json:"4,omitempty"`
}

// NewCreateNetworkApplianceRfProfileRequestPerSsidSettings instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings() *CreateNetworkApplianceRfProfileRequestPerSsidSettings {
	this := CreateNetworkApplianceRfProfileRequestPerSsidSettings{}
	return &this
}

// NewCreateNetworkApplianceRfProfileRequestPerSsidSettingsWithDefaults instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceRfProfileRequestPerSsidSettingsWithDefaults() *CreateNetworkApplianceRfProfileRequestPerSsidSettings {
	this := CreateNetworkApplianceRfProfileRequestPerSsidSettings{}
	return &this
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar1() CreateNetworkApplianceRfProfileRequestPerSsidSettings1 {
	if o == nil || IsNil(o.Var1) {
		var ret CreateNetworkApplianceRfProfileRequestPerSsidSettings1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar1Ok() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings1, bool) {
	if o == nil || IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) HasVar1() bool {
	if o != nil && !IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given CreateNetworkApplianceRfProfileRequestPerSsidSettings1 and assigns it to the Var1 field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) SetVar1(v CreateNetworkApplianceRfProfileRequestPerSsidSettings1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar2() CreateNetworkApplianceRfProfileRequestPerSsidSettings2 {
	if o == nil || IsNil(o.Var2) {
		var ret CreateNetworkApplianceRfProfileRequestPerSsidSettings2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar2Ok() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings2, bool) {
	if o == nil || IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) HasVar2() bool {
	if o != nil && !IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given CreateNetworkApplianceRfProfileRequestPerSsidSettings2 and assigns it to the Var2 field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) SetVar2(v CreateNetworkApplianceRfProfileRequestPerSsidSettings2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar3() CreateNetworkApplianceRfProfileRequestPerSsidSettings3 {
	if o == nil || IsNil(o.Var3) {
		var ret CreateNetworkApplianceRfProfileRequestPerSsidSettings3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar3Ok() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings3, bool) {
	if o == nil || IsNil(o.Var3) {
		return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) HasVar3() bool {
	if o != nil && !IsNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given CreateNetworkApplianceRfProfileRequestPerSsidSettings3 and assigns it to the Var3 field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) SetVar3(v CreateNetworkApplianceRfProfileRequestPerSsidSettings3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar4() CreateNetworkApplianceRfProfileRequestPerSsidSettings4 {
	if o == nil || IsNil(o.Var4) {
		var ret CreateNetworkApplianceRfProfileRequestPerSsidSettings4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) GetVar4Ok() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings4, bool) {
	if o == nil || IsNil(o.Var4) {
		return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) HasVar4() bool {
	if o != nil && !IsNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given CreateNetworkApplianceRfProfileRequestPerSsidSettings4 and assigns it to the Var4 field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings) SetVar4(v CreateNetworkApplianceRfProfileRequestPerSsidSettings4) {
	o.Var4 = &v
}

func (o CreateNetworkApplianceRfProfileRequestPerSsidSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkApplianceRfProfileRequestPerSsidSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !IsNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}
	if !IsNil(o.Var3) {
		toSerialize["3"] = o.Var3
	}
	if !IsNil(o.Var4) {
		toSerialize["4"] = o.Var4
	}
	return toSerialize, nil
}

type NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings struct {
	value *CreateNetworkApplianceRfProfileRequestPerSsidSettings
	isSet bool
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings) Get() *CreateNetworkApplianceRfProfileRequestPerSsidSettings {
	return v.value
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings) Set(val *CreateNetworkApplianceRfProfileRequestPerSsidSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceRfProfileRequestPerSsidSettings(val *CreateNetworkApplianceRfProfileRequestPerSsidSettings) *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings {
	return &NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


