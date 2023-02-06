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

// InlineObject170 struct for InlineObject170
type InlineObject170 struct {
	// Name of the adaptive policy ACL
	Name string `json:"name"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// An ordered array of the adaptive policy ACL rules.
	Rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules `json:"rules"`
	// IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	IpVersion string `json:"ipVersion"`
}

// NewInlineObject170 instantiates a new InlineObject170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject170(name string, rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules, ipVersion string) *InlineObject170 {
	this := InlineObject170{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Rules = rules
	this.IpVersion = ipVersion
	return &this
}

// NewInlineObject170WithDefaults instantiates a new InlineObject170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject170WithDefaults() *InlineObject170 {
	this := InlineObject170{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *InlineObject170) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject170) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject170) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject170) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject170) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value
func (o *InlineObject170) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules {
	if o == nil {
		var ret []OrganizationsOrganizationIdAdaptivePolicyAclsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetRulesOk() ([]OrganizationsOrganizationIdAdaptivePolicyAclsRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject170) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules) {
	o.Rules = v
}

// GetIpVersion returns the IpVersion field value
func (o *InlineObject170) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetIpVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *InlineObject170) SetIpVersion(v string) {
	o.IpVersion = v
}

func (o InlineObject170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["ipVersion"] = o.IpVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject170 struct {
	value *InlineObject170
	isSet bool
}

func (v NullableInlineObject170) Get() *InlineObject170 {
	return v.value
}

func (v *NullableInlineObject170) Set(val *InlineObject170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject170(val *InlineObject170) *NullableInlineObject170 {
	return &NullableInlineObject170{value: val, isSet: true}
}

func (v NullableInlineObject170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

