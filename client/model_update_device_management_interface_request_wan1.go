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

// checks if the UpdateDeviceManagementInterfaceRequestWan1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceManagementInterfaceRequestWan1{}

// UpdateDeviceManagementInterfaceRequestWan1 WAN 1 settings
type UpdateDeviceManagementInterfaceRequestWan1 struct {
	// Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
	WanEnabled *string `json:"wanEnabled,omitempty"`
	// Configure the interface to have static IP settings or use DHCP.
	UsingStaticIp *bool `json:"usingStaticIp,omitempty"`
	// The IP the device should use on the WAN.
	StaticIp *string `json:"staticIp,omitempty"`
	// The IP of the gateway on the WAN.
	StaticGatewayIp *string `json:"staticGatewayIp,omitempty"`
	// The subnet mask for the WAN.
	StaticSubnetMask *string `json:"staticSubnetMask,omitempty"`
	// Up to two DNS IPs.
	StaticDns []string `json:"staticDns,omitempty"`
	// The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.
	Vlan *int32 `json:"vlan,omitempty"`
}

// NewUpdateDeviceManagementInterfaceRequestWan1 instantiates a new UpdateDeviceManagementInterfaceRequestWan1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceManagementInterfaceRequestWan1() *UpdateDeviceManagementInterfaceRequestWan1 {
	this := UpdateDeviceManagementInterfaceRequestWan1{}
	return &this
}

// NewUpdateDeviceManagementInterfaceRequestWan1WithDefaults instantiates a new UpdateDeviceManagementInterfaceRequestWan1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceManagementInterfaceRequestWan1WithDefaults() *UpdateDeviceManagementInterfaceRequestWan1 {
	this := UpdateDeviceManagementInterfaceRequestWan1{}
	return &this
}

// GetWanEnabled returns the WanEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetWanEnabled() string {
	if o == nil || IsNil(o.WanEnabled) {
		var ret string
		return ret
	}
	return *o.WanEnabled
}

// GetWanEnabledOk returns a tuple with the WanEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetWanEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.WanEnabled) {
		return nil, false
	}
	return o.WanEnabled, true
}

// HasWanEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasWanEnabled() bool {
	if o != nil && !IsNil(o.WanEnabled) {
		return true
	}

	return false
}

// SetWanEnabled gets a reference to the given string and assigns it to the WanEnabled field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetWanEnabled(v string) {
	o.WanEnabled = &v
}

// GetUsingStaticIp returns the UsingStaticIp field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetUsingStaticIp() bool {
	if o == nil || IsNil(o.UsingStaticIp) {
		var ret bool
		return ret
	}
	return *o.UsingStaticIp
}

// GetUsingStaticIpOk returns a tuple with the UsingStaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetUsingStaticIpOk() (*bool, bool) {
	if o == nil || IsNil(o.UsingStaticIp) {
		return nil, false
	}
	return o.UsingStaticIp, true
}

// HasUsingStaticIp returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasUsingStaticIp() bool {
	if o != nil && !IsNil(o.UsingStaticIp) {
		return true
	}

	return false
}

// SetUsingStaticIp gets a reference to the given bool and assigns it to the UsingStaticIp field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetUsingStaticIp(v bool) {
	o.UsingStaticIp = &v
}

// GetStaticIp returns the StaticIp field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticIp() string {
	if o == nil || IsNil(o.StaticIp) {
		var ret string
		return ret
	}
	return *o.StaticIp
}

// GetStaticIpOk returns a tuple with the StaticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticIpOk() (*string, bool) {
	if o == nil || IsNil(o.StaticIp) {
		return nil, false
	}
	return o.StaticIp, true
}

// HasStaticIp returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasStaticIp() bool {
	if o != nil && !IsNil(o.StaticIp) {
		return true
	}

	return false
}

// SetStaticIp gets a reference to the given string and assigns it to the StaticIp field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetStaticIp(v string) {
	o.StaticIp = &v
}

// GetStaticGatewayIp returns the StaticGatewayIp field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticGatewayIp() string {
	if o == nil || IsNil(o.StaticGatewayIp) {
		var ret string
		return ret
	}
	return *o.StaticGatewayIp
}

// GetStaticGatewayIpOk returns a tuple with the StaticGatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.StaticGatewayIp) {
		return nil, false
	}
	return o.StaticGatewayIp, true
}

// HasStaticGatewayIp returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasStaticGatewayIp() bool {
	if o != nil && !IsNil(o.StaticGatewayIp) {
		return true
	}

	return false
}

// SetStaticGatewayIp gets a reference to the given string and assigns it to the StaticGatewayIp field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetStaticGatewayIp(v string) {
	o.StaticGatewayIp = &v
}

// GetStaticSubnetMask returns the StaticSubnetMask field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticSubnetMask() string {
	if o == nil || IsNil(o.StaticSubnetMask) {
		var ret string
		return ret
	}
	return *o.StaticSubnetMask
}

// GetStaticSubnetMaskOk returns a tuple with the StaticSubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticSubnetMaskOk() (*string, bool) {
	if o == nil || IsNil(o.StaticSubnetMask) {
		return nil, false
	}
	return o.StaticSubnetMask, true
}

// HasStaticSubnetMask returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasStaticSubnetMask() bool {
	if o != nil && !IsNil(o.StaticSubnetMask) {
		return true
	}

	return false
}

// SetStaticSubnetMask gets a reference to the given string and assigns it to the StaticSubnetMask field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetStaticSubnetMask(v string) {
	o.StaticSubnetMask = &v
}

// GetStaticDns returns the StaticDns field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticDns() []string {
	if o == nil || IsNil(o.StaticDns) {
		var ret []string
		return ret
	}
	return o.StaticDns
}

// GetStaticDnsOk returns a tuple with the StaticDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetStaticDnsOk() ([]string, bool) {
	if o == nil || IsNil(o.StaticDns) {
		return nil, false
	}
	return o.StaticDns, true
}

// HasStaticDns returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasStaticDns() bool {
	if o != nil && !IsNil(o.StaticDns) {
		return true
	}

	return false
}

// SetStaticDns gets a reference to the given []string and assigns it to the StaticDns field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetStaticDns(v []string) {
	o.StaticDns = v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequestWan1) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *UpdateDeviceManagementInterfaceRequestWan1) SetVlan(v int32) {
	o.Vlan = &v
}

func (o UpdateDeviceManagementInterfaceRequestWan1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceManagementInterfaceRequestWan1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WanEnabled) {
		toSerialize["wanEnabled"] = o.WanEnabled
	}
	if !IsNil(o.UsingStaticIp) {
		toSerialize["usingStaticIp"] = o.UsingStaticIp
	}
	if !IsNil(o.StaticIp) {
		toSerialize["staticIp"] = o.StaticIp
	}
	if !IsNil(o.StaticGatewayIp) {
		toSerialize["staticGatewayIp"] = o.StaticGatewayIp
	}
	if !IsNil(o.StaticSubnetMask) {
		toSerialize["staticSubnetMask"] = o.StaticSubnetMask
	}
	if !IsNil(o.StaticDns) {
		toSerialize["staticDns"] = o.StaticDns
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	return toSerialize, nil
}

type NullableUpdateDeviceManagementInterfaceRequestWan1 struct {
	value *UpdateDeviceManagementInterfaceRequestWan1
	isSet bool
}

func (v NullableUpdateDeviceManagementInterfaceRequestWan1) Get() *UpdateDeviceManagementInterfaceRequestWan1 {
	return v.value
}

func (v *NullableUpdateDeviceManagementInterfaceRequestWan1) Set(val *UpdateDeviceManagementInterfaceRequestWan1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceManagementInterfaceRequestWan1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceManagementInterfaceRequestWan1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceManagementInterfaceRequestWan1(val *UpdateDeviceManagementInterfaceRequestWan1) *NullableUpdateDeviceManagementInterfaceRequestWan1 {
	return &NullableUpdateDeviceManagementInterfaceRequestWan1{value: val, isSet: true}
}

func (v NullableUpdateDeviceManagementInterfaceRequestWan1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceManagementInterfaceRequestWan1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


