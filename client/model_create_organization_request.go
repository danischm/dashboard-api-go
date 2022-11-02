/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateOrganizationRequest struct for CreateOrganizationRequest
type CreateOrganizationRequest struct {
	// The name of the organization
	Name string `json:"name"`
}

// NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationRequest(name string) *CreateOrganizationRequest {
	this := CreateOrganizationRequest{}
	this.Name = name
	return &this
}

// NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest {
	this := CreateOrganizationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationRequest) SetName(v string) {
	o.Name = v
}

func (o CreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationRequest struct {
	value *CreateOrganizationRequest
	isSet bool
}

func (v NullableCreateOrganizationRequest) Get() *CreateOrganizationRequest {
	return v.value
}

func (v *NullableCreateOrganizationRequest) Set(val *CreateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationRequest(val *CreateOrganizationRequest) *NullableCreateOrganizationRequest {
	return &NullableCreateOrganizationRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

