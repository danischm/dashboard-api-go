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

// checks if the UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication{}

// UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication Settings for PPPoE Authentication.
type UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication struct {
	// Whether PPPoE authentication is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Username for PPPoE authentication.
	Username *string `json:"username,omitempty"`
	// Password for PPPoE authentication. This parameter is not returned.
	Password *string `json:"password,omitempty"`
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication{}
	return &this
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthenticationWithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthenticationWithDefaults() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication struct {
	value *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication
	isSet bool
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) Get() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication {
	return v.value
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) Set(val *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication(val *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication {
	return &NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


