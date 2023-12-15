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

// checks if the GetNetworkPoliciesByClient200ResponseInnerAssignedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkPoliciesByClient200ResponseInnerAssignedInner{}

// GetNetworkPoliciesByClient200ResponseInnerAssignedInner struct for GetNetworkPoliciesByClient200ResponseInnerAssignedInner
type GetNetworkPoliciesByClient200ResponseInnerAssignedInner struct {
	// name of policy
	Name *string `json:"name,omitempty"`
	// type of policy
	Type *string `json:"type,omitempty"`
	// id of policy
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// ssid
	Ssid []GetNetworkPoliciesByClient200ResponseInnerAssignedInnerSsidInner `json:"ssid,omitempty"`
}

// NewGetNetworkPoliciesByClient200ResponseInnerAssignedInner instantiates a new GetNetworkPoliciesByClient200ResponseInnerAssignedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkPoliciesByClient200ResponseInnerAssignedInner() *GetNetworkPoliciesByClient200ResponseInnerAssignedInner {
	this := GetNetworkPoliciesByClient200ResponseInnerAssignedInner{}
	return &this
}

// NewGetNetworkPoliciesByClient200ResponseInnerAssignedInnerWithDefaults instantiates a new GetNetworkPoliciesByClient200ResponseInnerAssignedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkPoliciesByClient200ResponseInnerAssignedInnerWithDefaults() *GetNetworkPoliciesByClient200ResponseInnerAssignedInner {
	this := GetNetworkPoliciesByClient200ResponseInnerAssignedInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) SetType(v string) {
	o.Type = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetSsid() []GetNetworkPoliciesByClient200ResponseInnerAssignedInnerSsidInner {
	if o == nil || IsNil(o.Ssid) {
		var ret []GetNetworkPoliciesByClient200ResponseInnerAssignedInnerSsidInner
		return ret
	}
	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) GetSsidOk() ([]GetNetworkPoliciesByClient200ResponseInnerAssignedInnerSsidInner, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given []GetNetworkPoliciesByClient200ResponseInnerAssignedInnerSsidInner and assigns it to the Ssid field.
func (o *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) SetSsid(v []GetNetworkPoliciesByClient200ResponseInnerAssignedInnerSsidInner) {
	o.Ssid = v
}

func (o GetNetworkPoliciesByClient200ResponseInnerAssignedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkPoliciesByClient200ResponseInnerAssignedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	return toSerialize, nil
}

type NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner struct {
	value *GetNetworkPoliciesByClient200ResponseInnerAssignedInner
	isSet bool
}

func (v NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner) Get() *GetNetworkPoliciesByClient200ResponseInnerAssignedInner {
	return v.value
}

func (v *NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner) Set(val *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner(val *GetNetworkPoliciesByClient200ResponseInnerAssignedInner) *NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner {
	return &NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner{value: val, isSet: true}
}

func (v NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkPoliciesByClient200ResponseInnerAssignedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


