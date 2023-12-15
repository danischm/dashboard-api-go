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

// checks if the MoveOrganizationLicensingCotermLicensesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveOrganizationLicensingCotermLicensesRequest{}

// MoveOrganizationLicensingCotermLicensesRequest struct for MoveOrganizationLicensingCotermLicensesRequest
type MoveOrganizationLicensingCotermLicensesRequest struct {
	Destination MoveOrganizationLicensingCotermLicensesRequestDestination `json:"destination"`
	// The list of licenses to move
	Licenses []MoveOrganizationLicensingCotermLicensesRequestLicensesInner `json:"licenses"`
}

// NewMoveOrganizationLicensingCotermLicensesRequest instantiates a new MoveOrganizationLicensingCotermLicensesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensingCotermLicensesRequest(destination MoveOrganizationLicensingCotermLicensesRequestDestination, licenses []MoveOrganizationLicensingCotermLicensesRequestLicensesInner) *MoveOrganizationLicensingCotermLicensesRequest {
	this := MoveOrganizationLicensingCotermLicensesRequest{}
	this.Destination = destination
	this.Licenses = licenses
	return &this
}

// NewMoveOrganizationLicensingCotermLicensesRequestWithDefaults instantiates a new MoveOrganizationLicensingCotermLicensesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensingCotermLicensesRequestWithDefaults() *MoveOrganizationLicensingCotermLicensesRequest {
	this := MoveOrganizationLicensingCotermLicensesRequest{}
	return &this
}

// GetDestination returns the Destination field value
func (o *MoveOrganizationLicensingCotermLicensesRequest) GetDestination() MoveOrganizationLicensingCotermLicensesRequestDestination {
	if o == nil {
		var ret MoveOrganizationLicensingCotermLicensesRequestDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequest) GetDestinationOk() (*MoveOrganizationLicensingCotermLicensesRequestDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *MoveOrganizationLicensingCotermLicensesRequest) SetDestination(v MoveOrganizationLicensingCotermLicensesRequestDestination) {
	o.Destination = v
}

// GetLicenses returns the Licenses field value
func (o *MoveOrganizationLicensingCotermLicensesRequest) GetLicenses() []MoveOrganizationLicensingCotermLicensesRequestLicensesInner {
	if o == nil {
		var ret []MoveOrganizationLicensingCotermLicensesRequestLicensesInner
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequest) GetLicensesOk() ([]MoveOrganizationLicensingCotermLicensesRequestLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *MoveOrganizationLicensingCotermLicensesRequest) SetLicenses(v []MoveOrganizationLicensingCotermLicensesRequestLicensesInner) {
	o.Licenses = v
}

func (o MoveOrganizationLicensingCotermLicensesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveOrganizationLicensingCotermLicensesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["licenses"] = o.Licenses
	return toSerialize, nil
}

type NullableMoveOrganizationLicensingCotermLicensesRequest struct {
	value *MoveOrganizationLicensingCotermLicensesRequest
	isSet bool
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequest) Get() *MoveOrganizationLicensingCotermLicensesRequest {
	return v.value
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequest) Set(val *MoveOrganizationLicensingCotermLicensesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensingCotermLicensesRequest(val *MoveOrganizationLicensingCotermLicensesRequest) *NullableMoveOrganizationLicensingCotermLicensesRequest {
	return &NullableMoveOrganizationLicensingCotermLicensesRequest{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


