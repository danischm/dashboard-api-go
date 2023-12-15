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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser User parameters
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser struct {
	// The name of the device user for Meraki monitoring
	Username *string `json:"username,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUserWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUserWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) SetUsername(v string) {
	o.Username = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInnerUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


