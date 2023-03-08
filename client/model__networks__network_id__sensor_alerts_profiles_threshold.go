/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThreshold Threshold for sensor readings that will cause an alert to be sent. This object should contain a single property key matching the condition's 'metric' value.
type NetworksNetworkIdSensorAlertsProfilesThreshold struct {
	Temperature *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature `json:"temperature,omitempty"`
	Humidity *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity `json:"humidity,omitempty"`
	Water *NetworksNetworkIdSensorAlertsProfilesThresholdWater `json:"water,omitempty"`
	Door *NetworksNetworkIdSensorAlertsProfilesThresholdDoor `json:"door,omitempty"`
	Tvoc *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc `json:"tvoc,omitempty"`
	Pm25 *NetworksNetworkIdSensorAlertsProfilesThresholdPm25 `json:"pm25,omitempty"`
	Noise *NetworksNetworkIdSensorAlertsProfilesThresholdNoise `json:"noise,omitempty"`
	IndoorAirQuality *NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality `json:"indoorAirQuality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThreshold instantiates a new NetworksNetworkIdSensorAlertsProfilesThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThreshold() *NetworksNetworkIdSensorAlertsProfilesThreshold {
	this := NetworksNetworkIdSensorAlertsProfilesThreshold{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThreshold {
	this := NetworksNetworkIdSensorAlertsProfilesThreshold{}
	return &this
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetTemperature() NetworksNetworkIdSensorAlertsProfilesThresholdTemperature {
	if o == nil || isNil(o.Temperature) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdTemperature
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetTemperatureOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdTemperature, bool) {
	if o == nil || isNil(o.Temperature) {
    return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasTemperature() bool {
	if o != nil && !isNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdTemperature and assigns it to the Temperature field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetTemperature(v NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) {
	o.Temperature = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetHumidity() NetworksNetworkIdSensorAlertsProfilesThresholdHumidity {
	if o == nil || isNil(o.Humidity) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdHumidity
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetHumidityOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdHumidity, bool) {
	if o == nil || isNil(o.Humidity) {
    return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasHumidity() bool {
	if o != nil && !isNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdHumidity and assigns it to the Humidity field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetHumidity(v NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) {
	o.Humidity = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetWater() NetworksNetworkIdSensorAlertsProfilesThresholdWater {
	if o == nil || isNil(o.Water) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdWater
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetWaterOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdWater, bool) {
	if o == nil || isNil(o.Water) {
    return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasWater() bool {
	if o != nil && !isNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdWater and assigns it to the Water field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetWater(v NetworksNetworkIdSensorAlertsProfilesThresholdWater) {
	o.Water = &v
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetDoor() NetworksNetworkIdSensorAlertsProfilesThresholdDoor {
	if o == nil || isNil(o.Door) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdDoor
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetDoorOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdDoor, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdDoor and assigns it to the Door field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetDoor(v NetworksNetworkIdSensorAlertsProfilesThresholdDoor) {
	o.Door = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetTvoc() NetworksNetworkIdSensorAlertsProfilesThresholdTvoc {
	if o == nil || isNil(o.Tvoc) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdTvoc
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetTvocOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdTvoc, bool) {
	if o == nil || isNil(o.Tvoc) {
    return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasTvoc() bool {
	if o != nil && !isNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdTvoc and assigns it to the Tvoc field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetTvoc(v NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) {
	o.Tvoc = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetPm25() NetworksNetworkIdSensorAlertsProfilesThresholdPm25 {
	if o == nil || isNil(o.Pm25) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdPm25
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetPm25Ok() (*NetworksNetworkIdSensorAlertsProfilesThresholdPm25, bool) {
	if o == nil || isNil(o.Pm25) {
    return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasPm25() bool {
	if o != nil && !isNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdPm25 and assigns it to the Pm25 field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetPm25(v NetworksNetworkIdSensorAlertsProfilesThresholdPm25) {
	o.Pm25 = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetNoise() NetworksNetworkIdSensorAlertsProfilesThresholdNoise {
	if o == nil || isNil(o.Noise) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetNoiseOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdNoise, bool) {
	if o == nil || isNil(o.Noise) {
    return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasNoise() bool {
	if o != nil && !isNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdNoise and assigns it to the Noise field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetNoise(v NetworksNetworkIdSensorAlertsProfilesThresholdNoise) {
	o.Noise = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetIndoorAirQuality() NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality {
	if o == nil || isNil(o.IndoorAirQuality) {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) GetIndoorAirQualityOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality, bool) {
	if o == nil || isNil(o.IndoorAirQuality) {
    return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) HasIndoorAirQuality() bool {
	if o != nil && !isNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality and assigns it to the IndoorAirQuality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThreshold) SetIndoorAirQuality(v NetworksNetworkIdSensorAlertsProfilesThresholdIndoorAirQuality) {
	o.IndoorAirQuality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !isNil(o.Humidity) {
		toSerialize["humidity"] = o.Humidity
	}
	if !isNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	if !isNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !isNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !isNil(o.Pm25) {
		toSerialize["pm25"] = o.Pm25
	}
	if !isNil(o.Noise) {
		toSerialize["noise"] = o.Noise
	}
	if !isNil(o.IndoorAirQuality) {
		toSerialize["indoorAirQuality"] = o.IndoorAirQuality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThreshold struct {
	value *NetworksNetworkIdSensorAlertsProfilesThreshold
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThreshold) Get() *NetworksNetworkIdSensorAlertsProfilesThreshold {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThreshold) Set(val *NetworksNetworkIdSensorAlertsProfilesThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThreshold(val *NetworksNetworkIdSensorAlertsProfilesThreshold) *NullableNetworksNetworkIdSensorAlertsProfilesThreshold {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThreshold{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


