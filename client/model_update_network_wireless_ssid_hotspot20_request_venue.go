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

// checks if the UpdateNetworkWirelessSsidHotspot20RequestVenue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidHotspot20RequestVenue{}

// UpdateNetworkWirelessSsidHotspot20RequestVenue Venue settings for this SSID
type UpdateNetworkWirelessSsidHotspot20RequestVenue struct {
	// Venue name
	Name *string `json:"name,omitempty"`
	// Venue type ('Unspecified', 'Unspecified Assembly', 'Arena', 'Stadium', 'Passenger Terminal', 'Amphitheater', 'Amusement Park', 'Place of Worship', 'Convention Center', 'Library', 'Museum', 'Restaurant', 'Theater', 'Bar', 'Coffee Shop', 'Zoo or Aquarium', 'Emergency Coordination Center', 'Unspecified Business', 'Doctor or Dentist office', 'Bank', 'Fire Station', 'Police Station', 'Post Office', 'Professional Office', 'Research and Development Facility', 'Attorney Office', 'Unspecified Educational', 'School, Primary', 'School, Secondary', 'University or College', 'Unspecified Factory and Industrial', 'Factory', 'Unspecified Institutional', 'Hospital', 'Long-Term Care Facility', 'Alcohol and Drug Rehabilitation Center', 'Group Home', 'Prison or Jail', 'Unspecified Mercantile', 'Retail Store', 'Grocery Market', 'Automotive Service Station', 'Shopping Mall', 'Gas Station', 'Unspecified Residential', 'Private Residence', 'Hotel or Motel', 'Dormitory', 'Boarding House', 'Unspecified Storage', 'Unspecified Utility and Miscellaneous', 'Unspecified Vehicular', 'Automobile or Truck', 'Airplane', 'Bus', 'Ferry', 'Ship or Boat', 'Train', 'Motor Bike', 'Unspecified Outdoor', 'Muni-mesh Network', 'City Park', 'Rest Area', 'Traffic Control', 'Bus Stop', 'Kiosk')
	Type *string `json:"type,omitempty"`
}

// NewUpdateNetworkWirelessSsidHotspot20RequestVenue instantiates a new UpdateNetworkWirelessSsidHotspot20RequestVenue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidHotspot20RequestVenue() *UpdateNetworkWirelessSsidHotspot20RequestVenue {
	this := UpdateNetworkWirelessSsidHotspot20RequestVenue{}
	return &this
}

// NewUpdateNetworkWirelessSsidHotspot20RequestVenueWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20RequestVenue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidHotspot20RequestVenueWithDefaults() *UpdateNetworkWirelessSsidHotspot20RequestVenue {
	this := UpdateNetworkWirelessSsidHotspot20RequestVenue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestVenue) SetType(v string) {
	o.Type = &v
}

func (o UpdateNetworkWirelessSsidHotspot20RequestVenue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidHotspot20RequestVenue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidHotspot20RequestVenue struct {
	value *UpdateNetworkWirelessSsidHotspot20RequestVenue
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestVenue) Get() *UpdateNetworkWirelessSsidHotspot20RequestVenue {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestVenue) Set(val *UpdateNetworkWirelessSsidHotspot20RequestVenue) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestVenue) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestVenue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidHotspot20RequestVenue(val *UpdateNetworkWirelessSsidHotspot20RequestVenue) *NullableUpdateNetworkWirelessSsidHotspot20RequestVenue {
	return &NullableUpdateNetworkWirelessSsidHotspot20RequestVenue{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestVenue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestVenue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


