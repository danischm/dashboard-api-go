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

// checks if the UpdateNetworkSwitchStackRoutingInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchStackRoutingInterfaceRequest{}

// UpdateNetworkSwitchStackRoutingInterfaceRequest struct for UpdateNetworkSwitchStackRoutingInterfaceRequest
type UpdateNetworkSwitchStackRoutingInterfaceRequest struct {
	// A friendly name or description for the interface or VLAN.
	Name *string `json:"name,omitempty"`
	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet *string `json:"subnet,omitempty"`
	// The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'.
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	VlanId *int32 `json:"vlanId,omitempty"`
	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *UpdateNetworkSwitchStackRoutingInterfaceRequestOspfSettings `json:"ospfSettings,omitempty"`
	Ipv6 *UpdateNetworkSwitchStackRoutingInterfaceRequestIpv6 `json:"ipv6,omitempty"`
}

// NewUpdateNetworkSwitchStackRoutingInterfaceRequest instantiates a new UpdateNetworkSwitchStackRoutingInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchStackRoutingInterfaceRequest() *UpdateNetworkSwitchStackRoutingInterfaceRequest {
	this := UpdateNetworkSwitchStackRoutingInterfaceRequest{}
	return &this
}

// NewUpdateNetworkSwitchStackRoutingInterfaceRequestWithDefaults instantiates a new UpdateNetworkSwitchStackRoutingInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchStackRoutingInterfaceRequestWithDefaults() *UpdateNetworkSwitchStackRoutingInterfaceRequest {
	this := UpdateNetworkSwitchStackRoutingInterfaceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetInterfaceIp() string {
	if o == nil || IsNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetInterfaceIpOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceIp) {
		return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasInterfaceIp() bool {
	if o != nil && !IsNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetMulticastRouting() string {
	if o == nil || IsNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.MulticastRouting) {
		return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasMulticastRouting() bool {
	if o != nil && !IsNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetDefaultGateway() string {
	if o == nil || IsNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultGateway) {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasDefaultGateway() bool {
	if o != nil && !IsNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetOspfSettings() UpdateNetworkSwitchStackRoutingInterfaceRequestOspfSettings {
	if o == nil || IsNil(o.OspfSettings) {
		var ret UpdateNetworkSwitchStackRoutingInterfaceRequestOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetOspfSettingsOk() (*UpdateNetworkSwitchStackRoutingInterfaceRequestOspfSettings, bool) {
	if o == nil || IsNil(o.OspfSettings) {
		return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasOspfSettings() bool {
	if o != nil && !IsNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given UpdateNetworkSwitchStackRoutingInterfaceRequestOspfSettings and assigns it to the OspfSettings field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetOspfSettings(v UpdateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) {
	o.OspfSettings = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetIpv6() UpdateNetworkSwitchStackRoutingInterfaceRequestIpv6 {
	if o == nil || IsNil(o.Ipv6) {
		var ret UpdateNetworkSwitchStackRoutingInterfaceRequestIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) GetIpv6Ok() (*UpdateNetworkSwitchStackRoutingInterfaceRequestIpv6, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given UpdateNetworkSwitchStackRoutingInterfaceRequestIpv6 and assigns it to the Ipv6 field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceRequest) SetIpv6(v UpdateNetworkSwitchStackRoutingInterfaceRequestIpv6) {
	o.Ipv6 = &v
}

func (o UpdateNetworkSwitchStackRoutingInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchStackRoutingInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.InterfaceIp) {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if !IsNil(o.MulticastRouting) {
		toSerialize["multicastRouting"] = o.MulticastRouting
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.DefaultGateway) {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if !IsNil(o.OspfSettings) {
		toSerialize["ospfSettings"] = o.OspfSettings
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchStackRoutingInterfaceRequest struct {
	value *UpdateNetworkSwitchStackRoutingInterfaceRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchStackRoutingInterfaceRequest) Get() *UpdateNetworkSwitchStackRoutingInterfaceRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchStackRoutingInterfaceRequest) Set(val *UpdateNetworkSwitchStackRoutingInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchStackRoutingInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchStackRoutingInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchStackRoutingInterfaceRequest(val *UpdateNetworkSwitchStackRoutingInterfaceRequest) *NullableUpdateNetworkSwitchStackRoutingInterfaceRequest {
	return &NullableUpdateNetworkSwitchStackRoutingInterfaceRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchStackRoutingInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchStackRoutingInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

