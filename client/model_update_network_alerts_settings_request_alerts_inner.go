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

// UpdateNetworkAlertsSettingsRequestAlertsInner struct for UpdateNetworkAlertsSettingsRequestAlertsInner
type UpdateNetworkAlertsSettingsRequestAlertsInner struct {
	// The type of alert
	Type string `json:"type"`
	// A boolean depicting if the alert is turned on or off
	Enabled *bool `json:"enabled,omitempty"`
	AlertDestinations *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations `json:"alertDestinations,omitempty"`
	// A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// NewUpdateNetworkAlertsSettingsRequestAlertsInner instantiates a new UpdateNetworkAlertsSettingsRequestAlertsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAlertsSettingsRequestAlertsInner(type_ string) *UpdateNetworkAlertsSettingsRequestAlertsInner {
	this := UpdateNetworkAlertsSettingsRequestAlertsInner{}
	this.Type = type_
	return &this
}

// NewUpdateNetworkAlertsSettingsRequestAlertsInnerWithDefaults instantiates a new UpdateNetworkAlertsSettingsRequestAlertsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAlertsSettingsRequestAlertsInnerWithDefaults() *UpdateNetworkAlertsSettingsRequestAlertsInner {
	this := UpdateNetworkAlertsSettingsRequestAlertsInner{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAlertDestinations returns the AlertDestinations field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetAlertDestinations() UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations {
	if o == nil || isNil(o.AlertDestinations) {
		var ret UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations
		return ret
	}
	return *o.AlertDestinations
}

// GetAlertDestinationsOk returns a tuple with the AlertDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetAlertDestinationsOk() (*UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations, bool) {
	if o == nil || isNil(o.AlertDestinations) {
    return nil, false
	}
	return o.AlertDestinations, true
}

// HasAlertDestinations returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) HasAlertDestinations() bool {
	if o != nil && !isNil(o.AlertDestinations) {
		return true
	}

	return false
}

// SetAlertDestinations gets a reference to the given UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations and assigns it to the AlertDestinations field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) SetAlertDestinations(v UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) {
	o.AlertDestinations = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetFilters() map[string]interface{} {
	if o == nil || isNil(o.Filters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) GetFiltersOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Filters) {
    return map[string]interface{}{}, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string]interface{} and assigns it to the Filters field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInner) SetFilters(v map[string]interface{}) {
	o.Filters = v
}

func (o UpdateNetworkAlertsSettingsRequestAlertsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AlertDestinations) {
		toSerialize["alertDestinations"] = o.AlertDestinations
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkAlertsSettingsRequestAlertsInner struct {
	value *UpdateNetworkAlertsSettingsRequestAlertsInner
	isSet bool
}

func (v NullableUpdateNetworkAlertsSettingsRequestAlertsInner) Get() *UpdateNetworkAlertsSettingsRequestAlertsInner {
	return v.value
}

func (v *NullableUpdateNetworkAlertsSettingsRequestAlertsInner) Set(val *UpdateNetworkAlertsSettingsRequestAlertsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAlertsSettingsRequestAlertsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAlertsSettingsRequestAlertsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAlertsSettingsRequestAlertsInner(val *UpdateNetworkAlertsSettingsRequestAlertsInner) *NullableUpdateNetworkAlertsSettingsRequestAlertsInner {
	return &NullableUpdateNetworkAlertsSettingsRequestAlertsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkAlertsSettingsRequestAlertsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAlertsSettingsRequestAlertsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

