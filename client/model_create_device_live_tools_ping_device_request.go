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

// checks if the CreateDeviceLiveToolsPingDeviceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsPingDeviceRequest{}

// CreateDeviceLiveToolsPingDeviceRequest struct for CreateDeviceLiveToolsPingDeviceRequest
type CreateDeviceLiveToolsPingDeviceRequest struct {
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
	Callback *CreateDeviceLiveToolsPingRequestCallback `json:"callback,omitempty"`
}

// NewCreateDeviceLiveToolsPingDeviceRequest instantiates a new CreateDeviceLiveToolsPingDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsPingDeviceRequest() *CreateDeviceLiveToolsPingDeviceRequest {
	this := CreateDeviceLiveToolsPingDeviceRequest{}
	return &this
}

// NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults instantiates a new CreateDeviceLiveToolsPingDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults() *CreateDeviceLiveToolsPingDeviceRequest {
	this := CreateDeviceLiveToolsPingDeviceRequest{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingDeviceRequest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CreateDeviceLiveToolsPingDeviceRequest) SetCount(v int32) {
	o.Count = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCallback() CreateDeviceLiveToolsPingRequestCallback {
	if o == nil || IsNil(o.Callback) {
		var ret CreateDeviceLiveToolsPingRequestCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCallbackOk() (*CreateDeviceLiveToolsPingRequestCallback, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingDeviceRequest) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given CreateDeviceLiveToolsPingRequestCallback and assigns it to the Callback field.
func (o *CreateDeviceLiveToolsPingDeviceRequest) SetCallback(v CreateDeviceLiveToolsPingRequestCallback) {
	o.Callback = &v
}

func (o CreateDeviceLiveToolsPingDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsPingDeviceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsPingDeviceRequest struct {
	value *CreateDeviceLiveToolsPingDeviceRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsPingDeviceRequest) Get() *CreateDeviceLiveToolsPingDeviceRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsPingDeviceRequest) Set(val *CreateDeviceLiveToolsPingDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsPingDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsPingDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsPingDeviceRequest(val *CreateDeviceLiveToolsPingDeviceRequest) *NullableCreateDeviceLiveToolsPingDeviceRequest {
	return &NullableCreateDeviceLiveToolsPingDeviceRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsPingDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsPingDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


