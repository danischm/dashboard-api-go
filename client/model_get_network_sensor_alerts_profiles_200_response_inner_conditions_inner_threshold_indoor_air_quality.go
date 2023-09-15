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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality Indoor air quality score threshold. One of 'score' or 'quality' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality struct {
	// Alerting threshold as indoor air quality score.
	Score *int32 `json:"score,omitempty"`
	// Alerting threshold as a qualitative indoor air quality level.
	Quality *string `json:"quality,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQualityWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQualityWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) SetScore(v int32) {
	o.Score = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) SetQuality(v string) {
	o.Quality = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdIndoorAirQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


