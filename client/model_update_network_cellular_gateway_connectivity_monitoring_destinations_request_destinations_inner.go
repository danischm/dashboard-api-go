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

// checks if the UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner{}

// UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner struct for UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner
type UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner struct {
	// The IP address to test connectivity with
	Ip string `json:"ip"`
	// Description of the testing destination. Optional, defaults to an empty string
	Description *string `json:"description,omitempty"`
	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default *bool `json:"default,omitempty"`
}

// NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner instantiates a new UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner(ip string) *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner {
	this := UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner{}
	this.Ip = ip
	var description string = ""
	this.Description = &description
	var default_ bool = false
	this.Default = &default_
	return &this
}

// NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInnerWithDefaults instantiates a new UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInnerWithDefaults() *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner {
	this := UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner{}
	var description string = ""
	this.Description = &description
	var default_ bool = false
	this.Default = &default_
	return &this
}

// GetIp returns the Ip field value
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) SetIp(v string) {
	o.Ip = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) SetDescription(v string) {
	o.Description = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) SetDefault(v bool) {
	o.Default = &v
}

func (o UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner struct {
	value *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner
	isSet bool
}

func (v NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) Get() *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner {
	return v.value
}

func (v *NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) Set(val *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner(val *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) *NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner {
	return &NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequestDestinationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


