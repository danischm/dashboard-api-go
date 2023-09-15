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

// checks if the GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner{}

// GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner struct for GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner
type GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner struct {
	Network *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner `json:"byBand,omitempty"`
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner{}
	return &this
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	if o == nil || IsNil(o.ByBand) {
		var ret []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) GetByBandOk() ([]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool) {
	if o == nil || IsNil(o.ByBand) {
		return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) HasByBand() bool {
	if o != nil && !IsNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner and assigns it to the ByBand field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) {
	o.ByBand = v
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner struct {
	value *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) Get() *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) Set(val *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner(val *GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) *NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner {
	return &NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


