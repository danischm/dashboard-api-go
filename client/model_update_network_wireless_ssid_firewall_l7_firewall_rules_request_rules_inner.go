/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner struct for UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner struct {
	// 'Deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Type *string `json:"type,omitempty"`
	// The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.
	Value *string `json:"value,omitempty"`
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner {
	this := UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner{}
	return &this
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInnerWithDefaults() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner {
	this := UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) SetValue(v string) {
	o.Value = &v
}

func (o UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner struct {
	value *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) Get() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) Set(val *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner(val *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner {
	return &NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

