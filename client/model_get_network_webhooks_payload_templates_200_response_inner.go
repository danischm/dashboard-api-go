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

// checks if the GetNetworkWebhooksPayloadTemplates200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWebhooksPayloadTemplates200ResponseInner{}

// GetNetworkWebhooksPayloadTemplates200ResponseInner struct for GetNetworkWebhooksPayloadTemplates200ResponseInner
type GetNetworkWebhooksPayloadTemplates200ResponseInner struct {
	// Webhook payload template Id
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The type of the payload template
	Type *string `json:"type,omitempty"`
	// The name of the payload template
	Name *string `json:"name,omitempty"`
	// The payload template headers, will be rendered as a key-value pair in the webhook.
	Headers []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner `json:"headers,omitempty"`
	// The body of the payload template, in liquid template
	Body *string `json:"body,omitempty"`
	Sharing *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing `json:"sharing,omitempty"`
}

// NewGetNetworkWebhooksPayloadTemplates200ResponseInner instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWebhooksPayloadTemplates200ResponseInner() *GetNetworkWebhooksPayloadTemplates200ResponseInner {
	this := GetNetworkWebhooksPayloadTemplates200ResponseInner{}
	return &this
}

// NewGetNetworkWebhooksPayloadTemplates200ResponseInnerWithDefaults instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWebhooksPayloadTemplates200ResponseInnerWithDefaults() *GetNetworkWebhooksPayloadTemplates200ResponseInner {
	this := GetNetworkWebhooksPayloadTemplates200ResponseInner{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetPayloadTemplateId() string {
	if o == nil || IsNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadTemplateId) {
		return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasPayloadTemplateId() bool {
	if o != nil && !IsNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetHeaders() []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner {
	if o == nil || IsNil(o.Headers) {
		var ret []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetHeadersOk() ([]GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner and assigns it to the Headers field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetHeaders(v []GetNetworkWebhooksPayloadTemplates200ResponseInnerHeadersInner) {
	o.Headers = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetBody(v string) {
	o.Body = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetSharing() GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing {
	if o == nil || IsNil(o.Sharing) {
		var ret GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) GetSharingOk() (*GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing and assigns it to the Sharing field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInner) SetSharing(v GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) {
	o.Sharing = &v
}

func (o GetNetworkWebhooksPayloadTemplates200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWebhooksPayloadTemplates200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	return toSerialize, nil
}

type NullableGetNetworkWebhooksPayloadTemplates200ResponseInner struct {
	value *GetNetworkWebhooksPayloadTemplates200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInner) Get() *GetNetworkWebhooksPayloadTemplates200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInner) Set(val *GetNetworkWebhooksPayloadTemplates200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWebhooksPayloadTemplates200ResponseInner(val *GetNetworkWebhooksPayloadTemplates200ResponseInner) *NullableGetNetworkWebhooksPayloadTemplates200ResponseInner {
	return &NullableGetNetworkWebhooksPayloadTemplates200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


