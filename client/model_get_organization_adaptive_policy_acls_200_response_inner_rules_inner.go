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

// checks if the GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner{}

// GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner struct for GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner
type GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner struct {
	// 'allow' or 'deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// The type of protocol
	Protocol *string `json:"protocol,omitempty"`
	// Source port
	SrcPort *string `json:"srcPort,omitempty"`
	// Destination port
	DstPort *string `json:"dstPort,omitempty"`
}

// NewGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner instantiates a new GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner() *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner {
	this := GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner{}
	return &this
}

// NewGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInnerWithDefaults() *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner {
	this := GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetSrcPort() string {
	if o == nil || IsNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetSrcPortOk() (*string, bool) {
	if o == nil || IsNil(o.SrcPort) {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) HasSrcPort() bool {
	if o != nil && !IsNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetDstPort() string {
	if o == nil || IsNil(o.DstPort) {
		var ret string
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) GetDstPortOk() (*string, bool) {
	if o == nil || IsNil(o.DstPort) {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) HasDstPort() bool {
	if o != nil && !IsNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given string and assigns it to the DstPort field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) SetDstPort(v string) {
	o.DstPort = &v
}

func (o GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !IsNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner struct {
	value *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) Get() *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) Set(val *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner(val *GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) *NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner {
	return &NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


