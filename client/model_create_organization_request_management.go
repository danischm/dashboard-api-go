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

// checks if the CreateOrganizationRequestManagement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationRequestManagement{}

// CreateOrganizationRequestManagement Information about the organization's management system
type CreateOrganizationRequestManagement struct {
	// Details related to organization management, possibly empty
	Details []GetOrganizations200ResponseInnerManagementDetailsInner `json:"details,omitempty"`
}

// NewCreateOrganizationRequestManagement instantiates a new CreateOrganizationRequestManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationRequestManagement() *CreateOrganizationRequestManagement {
	this := CreateOrganizationRequestManagement{}
	return &this
}

// NewCreateOrganizationRequestManagementWithDefaults instantiates a new CreateOrganizationRequestManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationRequestManagementWithDefaults() *CreateOrganizationRequestManagement {
	this := CreateOrganizationRequestManagement{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CreateOrganizationRequestManagement) GetDetails() []GetOrganizations200ResponseInnerManagementDetailsInner {
	if o == nil || IsNil(o.Details) {
		var ret []GetOrganizations200ResponseInnerManagementDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequestManagement) GetDetailsOk() ([]GetOrganizations200ResponseInnerManagementDetailsInner, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CreateOrganizationRequestManagement) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GetOrganizations200ResponseInnerManagementDetailsInner and assigns it to the Details field.
func (o *CreateOrganizationRequestManagement) SetDetails(v []GetOrganizations200ResponseInnerManagementDetailsInner) {
	o.Details = v
}

func (o CreateOrganizationRequestManagement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationRequestManagement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableCreateOrganizationRequestManagement struct {
	value *CreateOrganizationRequestManagement
	isSet bool
}

func (v NullableCreateOrganizationRequestManagement) Get() *CreateOrganizationRequestManagement {
	return v.value
}

func (v *NullableCreateOrganizationRequestManagement) Set(val *CreateOrganizationRequestManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationRequestManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationRequestManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationRequestManagement(val *CreateOrganizationRequestManagement) *NullableCreateOrganizationRequestManagement {
	return &NullableCreateOrganizationRequestManagement{value: val, isSet: true}
}

func (v NullableCreateOrganizationRequestManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationRequestManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


