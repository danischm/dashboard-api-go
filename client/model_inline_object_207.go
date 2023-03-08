/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject207 struct for InlineObject207
type InlineObject207 struct {
	Destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination `json:"destination"`
	// The list of licenses to move
	Licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses `json:"licenses"`
}

// NewInlineObject207 instantiates a new InlineObject207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject207(destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) *InlineObject207 {
	this := InlineObject207{}
	this.Destination = destination
	this.Licenses = licenses
	return &this
}

// NewInlineObject207WithDefaults instantiates a new InlineObject207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject207WithDefaults() *InlineObject207 {
	this := InlineObject207{}
	return &this
}

// GetDestination returns the Destination field value
func (o *InlineObject207) GetDestination() OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	if o == nil {
		var ret OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineObject207) GetDestinationOk() (*OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineObject207) SetDestination(v OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) {
	o.Destination = v
}

// GetLicenses returns the Licenses field value
func (o *InlineObject207) GetLicenses() []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	if o == nil {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *InlineObject207) GetLicensesOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses, bool) {
	if o == nil {
    return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *InlineObject207) SetLicenses(v []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) {
	o.Licenses = v
}

func (o InlineObject207) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject207 struct {
	value *InlineObject207
	isSet bool
}

func (v NullableInlineObject207) Get() *InlineObject207 {
	return v.value
}

func (v *NullableInlineObject207) Set(val *InlineObject207) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject207) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject207) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject207(val *InlineObject207) *NullableInlineObject207 {
	return &NullableInlineObject207{value: val, isSet: true}
}

func (v NullableInlineObject207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


