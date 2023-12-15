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

// checks if the UpdateOrganizationInsightMonitoredMediaServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationInsightMonitoredMediaServerRequest{}

// UpdateOrganizationInsightMonitoredMediaServerRequest struct for UpdateOrganizationInsightMonitoredMediaServerRequest
type UpdateOrganizationInsightMonitoredMediaServerRequest struct {
	// The name of the VoIP provider
	Name *string `json:"name,omitempty"`
	// The IP address (IPv4 only) or hostname of the media server to monitor
	Address *string `json:"address,omitempty"`
	// Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
	BestEffortMonitoringEnabled *bool `json:"bestEffortMonitoringEnabled,omitempty"`
}

// NewUpdateOrganizationInsightMonitoredMediaServerRequest instantiates a new UpdateOrganizationInsightMonitoredMediaServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationInsightMonitoredMediaServerRequest() *UpdateOrganizationInsightMonitoredMediaServerRequest {
	this := UpdateOrganizationInsightMonitoredMediaServerRequest{}
	return &this
}

// NewUpdateOrganizationInsightMonitoredMediaServerRequestWithDefaults instantiates a new UpdateOrganizationInsightMonitoredMediaServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationInsightMonitoredMediaServerRequestWithDefaults() *UpdateOrganizationInsightMonitoredMediaServerRequest {
	this := UpdateOrganizationInsightMonitoredMediaServerRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) SetAddress(v string) {
	o.Address = &v
}

// GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) GetBestEffortMonitoringEnabled() bool {
	if o == nil || IsNil(o.BestEffortMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.BestEffortMonitoringEnabled
}

// GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) GetBestEffortMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BestEffortMonitoringEnabled) {
		return nil, false
	}
	return o.BestEffortMonitoringEnabled, true
}

// HasBestEffortMonitoringEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) HasBestEffortMonitoringEnabled() bool {
	if o != nil && !IsNil(o.BestEffortMonitoringEnabled) {
		return true
	}

	return false
}

// SetBestEffortMonitoringEnabled gets a reference to the given bool and assigns it to the BestEffortMonitoringEnabled field.
func (o *UpdateOrganizationInsightMonitoredMediaServerRequest) SetBestEffortMonitoringEnabled(v bool) {
	o.BestEffortMonitoringEnabled = &v
}

func (o UpdateOrganizationInsightMonitoredMediaServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationInsightMonitoredMediaServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BestEffortMonitoringEnabled) {
		toSerialize["bestEffortMonitoringEnabled"] = o.BestEffortMonitoringEnabled
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationInsightMonitoredMediaServerRequest struct {
	value *UpdateOrganizationInsightMonitoredMediaServerRequest
	isSet bool
}

func (v NullableUpdateOrganizationInsightMonitoredMediaServerRequest) Get() *UpdateOrganizationInsightMonitoredMediaServerRequest {
	return v.value
}

func (v *NullableUpdateOrganizationInsightMonitoredMediaServerRequest) Set(val *UpdateOrganizationInsightMonitoredMediaServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationInsightMonitoredMediaServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationInsightMonitoredMediaServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationInsightMonitoredMediaServerRequest(val *UpdateOrganizationInsightMonitoredMediaServerRequest) *NullableUpdateOrganizationInsightMonitoredMediaServerRequest {
	return &NullableUpdateOrganizationInsightMonitoredMediaServerRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationInsightMonitoredMediaServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationInsightMonitoredMediaServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


