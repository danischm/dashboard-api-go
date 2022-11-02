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

// UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner struct for UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner
type UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner struct {
	// Auto assign a /64 prefix from the origin to the VLAN
	Autonomous *bool `json:"autonomous,omitempty"`
	// Manual configuration of a /64 prefix on the VLAN
	StaticPrefix *string `json:"staticPrefix,omitempty"`
	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 *string `json:"staticApplianceIp6,omitempty"`
	Origin *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin `json:"origin,omitempty"`
}

// NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner() *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner {
	this := UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner{}
	return &this
}

// NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerWithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerWithDefaults() *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner {
	this := UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner{}
	return &this
}

// GetAutonomous returns the Autonomous field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetAutonomous() bool {
	if o == nil || isNil(o.Autonomous) {
		var ret bool
		return ret
	}
	return *o.Autonomous
}

// GetAutonomousOk returns a tuple with the Autonomous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetAutonomousOk() (*bool, bool) {
	if o == nil || isNil(o.Autonomous) {
    return nil, false
	}
	return o.Autonomous, true
}

// HasAutonomous returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) HasAutonomous() bool {
	if o != nil && !isNil(o.Autonomous) {
		return true
	}

	return false
}

// SetAutonomous gets a reference to the given bool and assigns it to the Autonomous field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) SetAutonomous(v bool) {
	o.Autonomous = &v
}

// GetStaticPrefix returns the StaticPrefix field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetStaticPrefix() string {
	if o == nil || isNil(o.StaticPrefix) {
		var ret string
		return ret
	}
	return *o.StaticPrefix
}

// GetStaticPrefixOk returns a tuple with the StaticPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetStaticPrefixOk() (*string, bool) {
	if o == nil || isNil(o.StaticPrefix) {
    return nil, false
	}
	return o.StaticPrefix, true
}

// HasStaticPrefix returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) HasStaticPrefix() bool {
	if o != nil && !isNil(o.StaticPrefix) {
		return true
	}

	return false
}

// SetStaticPrefix gets a reference to the given string and assigns it to the StaticPrefix field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) SetStaticPrefix(v string) {
	o.StaticPrefix = &v
}

// GetStaticApplianceIp6 returns the StaticApplianceIp6 field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetStaticApplianceIp6() string {
	if o == nil || isNil(o.StaticApplianceIp6) {
		var ret string
		return ret
	}
	return *o.StaticApplianceIp6
}

// GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetStaticApplianceIp6Ok() (*string, bool) {
	if o == nil || isNil(o.StaticApplianceIp6) {
    return nil, false
	}
	return o.StaticApplianceIp6, true
}

// HasStaticApplianceIp6 returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) HasStaticApplianceIp6() bool {
	if o != nil && !isNil(o.StaticApplianceIp6) {
		return true
	}

	return false
}

// SetStaticApplianceIp6 gets a reference to the given string and assigns it to the StaticApplianceIp6 field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) SetStaticApplianceIp6(v string) {
	o.StaticApplianceIp6 = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetOrigin() UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin {
	if o == nil || isNil(o.Origin) {
		var ret UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) GetOriginOk() (*UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin, bool) {
	if o == nil || isNil(o.Origin) {
    return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) HasOrigin() bool {
	if o != nil && !isNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin and assigns it to the Origin field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) SetOrigin(v UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) {
	o.Origin = &v
}

func (o UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Autonomous) {
		toSerialize["autonomous"] = o.Autonomous
	}
	if !isNil(o.StaticPrefix) {
		toSerialize["staticPrefix"] = o.StaticPrefix
	}
	if !isNil(o.StaticApplianceIp6) {
		toSerialize["staticApplianceIp6"] = o.StaticApplianceIp6
	}
	if !isNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner struct {
	value *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) Get() *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) Set(val *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner(val *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner {
	return &NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

