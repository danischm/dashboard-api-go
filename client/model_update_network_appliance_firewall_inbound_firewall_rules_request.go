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

// checks if the UpdateNetworkApplianceFirewallInboundFirewallRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallInboundFirewallRulesRequest{}

// UpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct for UpdateNetworkApplianceFirewallInboundFirewallRulesRequest
type UpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequest instantiates a new UpdateNetworkApplianceFirewallInboundFirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequest() *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest {
	this := UpdateNetworkApplianceFirewallInboundFirewallRulesRequest{}
	return &this
}

// NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallInboundFirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequestWithDefaults() *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest {
	this := UpdateNetworkApplianceFirewallInboundFirewallRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) GetRules() []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) GetRulesOk() ([]UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) SetRules(v []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) GetSyslogDefaultRule() bool {
	if o == nil || IsNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || IsNil(o.SyslogDefaultRule) {
		return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) HasSyslogDefaultRule() bool {
	if o != nil && !IsNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct {
	value *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) Get() *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) Set(val *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest(val *UpdateNetworkApplianceFirewallInboundFirewallRulesRequest) *NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest {
	return &NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallInboundFirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


