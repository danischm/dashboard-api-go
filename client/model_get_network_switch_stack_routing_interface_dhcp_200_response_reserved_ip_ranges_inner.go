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

// checks if the GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner{}

// GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner struct for GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner
type GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner struct {
	// The starting IP address of the reserved IP range
	Start *string `json:"start,omitempty"`
	// The ending IP address of the reserved IP range
	End *string `json:"end,omitempty"`
	// The comment for the reserved IP range
	Comment *string `json:"comment,omitempty"`
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner() *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner {
	this := GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner{}
	return &this
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInnerWithDefaults instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInnerWithDefaults() *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner {
	this := GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) SetEnd(v string) {
	o.End = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) SetComment(v string) {
	o.Comment = &v
}

func (o GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner struct {
	value *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner
	isSet bool
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) Get() *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner {
	return v.value
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) Set(val *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner(val *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner {
	return &NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseReservedIpRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


