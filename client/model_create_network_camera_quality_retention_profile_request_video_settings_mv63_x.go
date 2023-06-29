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

// checks if the CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X{}

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X Quality and resolution for MV63X camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
	Resolution string `json:"resolution"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63XWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63XWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quality"] = o.Quality
	toSerialize["resolution"] = o.Resolution
	return toSerialize, nil
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63X) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

