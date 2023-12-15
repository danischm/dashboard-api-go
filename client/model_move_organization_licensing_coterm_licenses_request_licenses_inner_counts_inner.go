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

// checks if the MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner{}

// MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner struct for MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner
type MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner struct {
	// The license model type to move counts of
	Model string `json:"model"`
	// The number of counts to move
	Count int32 `json:"count"`
}

// NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner instantiates a new MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner(model string, count int32) *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner {
	this := MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner{}
	this.Model = model
	this.Count = count
	return &this
}

// NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInnerWithDefaults instantiates a new MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInnerWithDefaults() *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner {
	this := MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner{}
	return &this
}

// GetModel returns the Model field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) SetModel(v string) {
	o.Model = v
}

// GetCount returns the Count field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) SetCount(v int32) {
	o.Count = v
}

func (o MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

type NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner struct {
	value *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner
	isSet bool
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) Get() *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner {
	return v.value
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) Set(val *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner(val *MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner {
	return &NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


