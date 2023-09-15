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

// checks if the CombineOrganizationNetworks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CombineOrganizationNetworks200Response{}

// CombineOrganizationNetworks200Response struct for CombineOrganizationNetworks200Response
type CombineOrganizationNetworks200Response struct {
	ResultingNetwork *CombineOrganizationNetworks200ResponseResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewCombineOrganizationNetworks200Response instantiates a new CombineOrganizationNetworks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombineOrganizationNetworks200Response() *CombineOrganizationNetworks200Response {
	this := CombineOrganizationNetworks200Response{}
	return &this
}

// NewCombineOrganizationNetworks200ResponseWithDefaults instantiates a new CombineOrganizationNetworks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombineOrganizationNetworks200ResponseWithDefaults() *CombineOrganizationNetworks200Response {
	this := CombineOrganizationNetworks200Response{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *CombineOrganizationNetworks200Response) GetResultingNetwork() CombineOrganizationNetworks200ResponseResultingNetwork {
	if o == nil || IsNil(o.ResultingNetwork) {
		var ret CombineOrganizationNetworks200ResponseResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombineOrganizationNetworks200Response) GetResultingNetworkOk() (*CombineOrganizationNetworks200ResponseResultingNetwork, bool) {
	if o == nil || IsNil(o.ResultingNetwork) {
		return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *CombineOrganizationNetworks200Response) HasResultingNetwork() bool {
	if o != nil && !IsNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given CombineOrganizationNetworks200ResponseResultingNetwork and assigns it to the ResultingNetwork field.
func (o *CombineOrganizationNetworks200Response) SetResultingNetwork(v CombineOrganizationNetworks200ResponseResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o CombineOrganizationNetworks200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CombineOrganizationNetworks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return toSerialize, nil
}

type NullableCombineOrganizationNetworks200Response struct {
	value *CombineOrganizationNetworks200Response
	isSet bool
}

func (v NullableCombineOrganizationNetworks200Response) Get() *CombineOrganizationNetworks200Response {
	return v.value
}

func (v *NullableCombineOrganizationNetworks200Response) Set(val *CombineOrganizationNetworks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCombineOrganizationNetworks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCombineOrganizationNetworks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombineOrganizationNetworks200Response(val *CombineOrganizationNetworks200Response) *NullableCombineOrganizationNetworks200Response {
	return &NullableCombineOrganizationNetworks200Response{value: val, isSet: true}
}

func (v NullableCombineOrganizationNetworks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombineOrganizationNetworks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


