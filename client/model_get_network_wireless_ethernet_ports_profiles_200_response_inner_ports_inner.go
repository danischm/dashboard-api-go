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

// checks if the GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner{}

// GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner struct for GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner
type GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Number
	Number *int32 `json:"number,omitempty"`
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Ssid number
	Ssid *int32 `json:"ssid,omitempty"`
	// PSK Group number
	PskGroupId *string `json:"pskGroupId,omitempty"`
}

// NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner instantiates a new GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner() *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner {
	this := GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner{}
	return &this
}

// NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInnerWithDefaults instantiates a new GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInnerWithDefaults() *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner {
	this := GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) SetNumber(v int32) {
	o.Number = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetSsid() int32 {
	if o == nil || IsNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetSsidOk() (*int32, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) SetSsid(v int32) {
	o.Ssid = &v
}

// GetPskGroupId returns the PskGroupId field value if set, zero value otherwise.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetPskGroupId() string {
	if o == nil || IsNil(o.PskGroupId) {
		var ret string
		return ret
	}
	return *o.PskGroupId
}

// GetPskGroupIdOk returns a tuple with the PskGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) GetPskGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PskGroupId) {
		return nil, false
	}
	return o.PskGroupId, true
}

// HasPskGroupId returns a boolean if a field has been set.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) HasPskGroupId() bool {
	if o != nil && !IsNil(o.PskGroupId) {
		return true
	}

	return false
}

// SetPskGroupId gets a reference to the given string and assigns it to the PskGroupId field.
func (o *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) SetPskGroupId(v string) {
	o.PskGroupId = &v
}

func (o GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.PskGroupId) {
		toSerialize["pskGroupId"] = o.PskGroupId
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner struct {
	value *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner
	isSet bool
}

func (v NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) Get() *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner {
	return v.value
}

func (v *NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) Set(val *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner(val *GetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) *NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner {
	return &NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessEthernetPortsProfiles200ResponseInnerPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


