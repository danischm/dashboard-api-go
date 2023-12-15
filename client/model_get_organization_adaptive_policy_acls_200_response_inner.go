/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationAdaptivePolicyAcls200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyAcls200ResponseInner{}

// GetOrganizationAdaptivePolicyAcls200ResponseInner struct for GetOrganizationAdaptivePolicyAcls200ResponseInner
type GetOrganizationAdaptivePolicyAcls200ResponseInner struct {
	// ID of the adaptive policy ACL
	AclId *string `json:"aclId,omitempty"`
	// Name of the adaptive policy ACL
	Name *string `json:"name,omitempty"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// IP version of adpative policy ACL
	IpVersion *string `json:"ipVersion,omitempty"`
	// An ordered array of the adaptive policy ACL rules
	Rules []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner `json:"rules,omitempty"`
	// When the adaptive policy ACL was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// When the adaptive policy ACL was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewGetOrganizationAdaptivePolicyAcls200ResponseInner instantiates a new GetOrganizationAdaptivePolicyAcls200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyAcls200ResponseInner() *GetOrganizationAdaptivePolicyAcls200ResponseInner {
	this := GetOrganizationAdaptivePolicyAcls200ResponseInner{}
	return &this
}

// NewGetOrganizationAdaptivePolicyAcls200ResponseInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyAcls200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyAcls200ResponseInnerWithDefaults() *GetOrganizationAdaptivePolicyAcls200ResponseInner {
	this := GetOrganizationAdaptivePolicyAcls200ResponseInner{}
	return &this
}

// GetAclId returns the AclId field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetAclId() string {
	if o == nil || IsNil(o.AclId) {
		var ret string
		return ret
	}
	return *o.AclId
}

// GetAclIdOk returns a tuple with the AclId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetAclIdOk() (*string, bool) {
	if o == nil || IsNil(o.AclId) {
		return nil, false
	}
	return o.AclId, true
}

// HasAclId returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasAclId() bool {
	if o != nil && !IsNil(o.AclId) {
		return true
	}

	return false
}

// SetAclId gets a reference to the given string and assigns it to the AclId field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetAclId(v string) {
	o.AclId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetIpVersion() string {
	if o == nil || IsNil(o.IpVersion) {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetIpVersionOk() (*string, bool) {
	if o == nil || IsNil(o.IpVersion) {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasIpVersion() bool {
	if o != nil && !IsNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetRules() []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetRulesOk() ([]GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner and assigns it to the Rules field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetRules(v []GetOrganizationAdaptivePolicyAcls200ResponseInnerRulesInner) {
	o.Rules = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GetOrganizationAdaptivePolicyAcls200ResponseInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GetOrganizationAdaptivePolicyAcls200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyAcls200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AclId) {
		toSerialize["aclId"] = o.AclId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyAcls200ResponseInner struct {
	value *GetOrganizationAdaptivePolicyAcls200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyAcls200ResponseInner) Get() *GetOrganizationAdaptivePolicyAcls200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyAcls200ResponseInner) Set(val *GetOrganizationAdaptivePolicyAcls200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyAcls200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyAcls200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyAcls200ResponseInner(val *GetOrganizationAdaptivePolicyAcls200ResponseInner) *NullableGetOrganizationAdaptivePolicyAcls200ResponseInner {
	return &NullableGetOrganizationAdaptivePolicyAcls200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyAcls200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyAcls200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


