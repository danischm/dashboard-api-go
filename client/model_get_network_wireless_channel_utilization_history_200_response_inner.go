/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessChannelUtilizationHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessChannelUtilizationHistory200ResponseInner{}

// GetNetworkWirelessChannelUtilizationHistory200ResponseInner struct for GetNetworkWirelessChannelUtilizationHistory200ResponseInner
type GetNetworkWirelessChannelUtilizationHistory200ResponseInner struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Total channel utilization
	UtilizationTotal *float32 `json:"utilizationTotal,omitempty"`
	// Average wifi utilization
	Utilization80211 *float32 `json:"utilization80211,omitempty"`
	// Average signal interference
	UtilizationNon80211 *float32 `json:"utilizationNon80211,omitempty"`
}

// NewGetNetworkWirelessChannelUtilizationHistory200ResponseInner instantiates a new GetNetworkWirelessChannelUtilizationHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessChannelUtilizationHistory200ResponseInner() *GetNetworkWirelessChannelUtilizationHistory200ResponseInner {
	this := GetNetworkWirelessChannelUtilizationHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessChannelUtilizationHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessChannelUtilizationHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessChannelUtilizationHistory200ResponseInnerWithDefaults() *GetNetworkWirelessChannelUtilizationHistory200ResponseInner {
	this := GetNetworkWirelessChannelUtilizationHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetUtilizationTotal returns the UtilizationTotal field value if set, zero value otherwise.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetUtilizationTotal() float32 {
	if o == nil || IsNil(o.UtilizationTotal) {
		var ret float32
		return ret
	}
	return *o.UtilizationTotal
}

// GetUtilizationTotalOk returns a tuple with the UtilizationTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetUtilizationTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.UtilizationTotal) {
		return nil, false
	}
	return o.UtilizationTotal, true
}

// HasUtilizationTotal returns a boolean if a field has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) HasUtilizationTotal() bool {
	if o != nil && !IsNil(o.UtilizationTotal) {
		return true
	}

	return false
}

// SetUtilizationTotal gets a reference to the given float32 and assigns it to the UtilizationTotal field.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) SetUtilizationTotal(v float32) {
	o.UtilizationTotal = &v
}

// GetUtilization80211 returns the Utilization80211 field value if set, zero value otherwise.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetUtilization80211() float32 {
	if o == nil || IsNil(o.Utilization80211) {
		var ret float32
		return ret
	}
	return *o.Utilization80211
}

// GetUtilization80211Ok returns a tuple with the Utilization80211 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetUtilization80211Ok() (*float32, bool) {
	if o == nil || IsNil(o.Utilization80211) {
		return nil, false
	}
	return o.Utilization80211, true
}

// HasUtilization80211 returns a boolean if a field has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) HasUtilization80211() bool {
	if o != nil && !IsNil(o.Utilization80211) {
		return true
	}

	return false
}

// SetUtilization80211 gets a reference to the given float32 and assigns it to the Utilization80211 field.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) SetUtilization80211(v float32) {
	o.Utilization80211 = &v
}

// GetUtilizationNon80211 returns the UtilizationNon80211 field value if set, zero value otherwise.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetUtilizationNon80211() float32 {
	if o == nil || IsNil(o.UtilizationNon80211) {
		var ret float32
		return ret
	}
	return *o.UtilizationNon80211
}

// GetUtilizationNon80211Ok returns a tuple with the UtilizationNon80211 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) GetUtilizationNon80211Ok() (*float32, bool) {
	if o == nil || IsNil(o.UtilizationNon80211) {
		return nil, false
	}
	return o.UtilizationNon80211, true
}

// HasUtilizationNon80211 returns a boolean if a field has been set.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) HasUtilizationNon80211() bool {
	if o != nil && !IsNil(o.UtilizationNon80211) {
		return true
	}

	return false
}

// SetUtilizationNon80211 gets a reference to the given float32 and assigns it to the UtilizationNon80211 field.
func (o *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) SetUtilizationNon80211(v float32) {
	o.UtilizationNon80211 = &v
}

func (o GetNetworkWirelessChannelUtilizationHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessChannelUtilizationHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.UtilizationTotal) {
		toSerialize["utilizationTotal"] = o.UtilizationTotal
	}
	if !IsNil(o.Utilization80211) {
		toSerialize["utilization80211"] = o.Utilization80211
	}
	if !IsNil(o.UtilizationNon80211) {
		toSerialize["utilizationNon80211"] = o.UtilizationNon80211
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner struct {
	value *GetNetworkWirelessChannelUtilizationHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner) Get() *GetNetworkWirelessChannelUtilizationHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner) Set(val *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner(val *GetNetworkWirelessChannelUtilizationHistory200ResponseInner) *NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner {
	return &NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessChannelUtilizationHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

