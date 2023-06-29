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

// checks if the CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72{}

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 Quality and resolution for MV12/MV22/MV72 camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
	Resolution string `json:"resolution"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72WithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72WithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quality"] = o.Quality
	toSerialize["resolution"] = o.Resolution
	return toSerialize, nil
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72 {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12MV22MV72) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

