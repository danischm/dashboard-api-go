/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner{}

// GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner struct for GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner
type GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner struct {
	// The device model
	Model *string `json:"model,omitempty"`
	// Total number of devices per model
	Count *int32 `json:"count,omitempty"`
	Usage *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage `json:"usage,omitempty"`
}

// NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner instantiates a new GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner() *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner {
	this := GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerWithDefaults() *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner {
	this := GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) SetCount(v int32) {
	o.Count = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) GetUsage() GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage {
	if o == nil || IsNil(o.Usage) {
		var ret GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) GetUsageOk() (*GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage and assigns it to the Usage field.
func (o *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) SetUsage(v GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInnerUsage) {
	o.Usage = &v
}

func (o GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner struct {
	value *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) Get() *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) Set(val *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner(val *GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner {
	return &NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

