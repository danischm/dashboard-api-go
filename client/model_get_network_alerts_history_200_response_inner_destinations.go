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

// checks if the GetNetworkAlertsHistory200ResponseInnerDestinations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAlertsHistory200ResponseInnerDestinations{}

// GetNetworkAlertsHistory200ResponseInnerDestinations the destinations this alert is configured to be delivered to
type GetNetworkAlertsHistory200ResponseInnerDestinations struct {
	Email *GetNetworkAlertsHistory200ResponseInnerDestinationsEmail `json:"email,omitempty"`
	Push *GetNetworkAlertsHistory200ResponseInnerDestinationsPush `json:"push,omitempty"`
	Sms *GetNetworkAlertsHistory200ResponseInnerDestinationsSms `json:"sms,omitempty"`
	Webhook *GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook `json:"webhook,omitempty"`
}

// NewGetNetworkAlertsHistory200ResponseInnerDestinations instantiates a new GetNetworkAlertsHistory200ResponseInnerDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAlertsHistory200ResponseInnerDestinations() *GetNetworkAlertsHistory200ResponseInnerDestinations {
	this := GetNetworkAlertsHistory200ResponseInnerDestinations{}
	return &this
}

// NewGetNetworkAlertsHistory200ResponseInnerDestinationsWithDefaults instantiates a new GetNetworkAlertsHistory200ResponseInnerDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAlertsHistory200ResponseInnerDestinationsWithDefaults() *GetNetworkAlertsHistory200ResponseInnerDestinations {
	this := GetNetworkAlertsHistory200ResponseInnerDestinations{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetEmail() GetNetworkAlertsHistory200ResponseInnerDestinationsEmail {
	if o == nil || IsNil(o.Email) {
		var ret GetNetworkAlertsHistory200ResponseInnerDestinationsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetEmailOk() (*GetNetworkAlertsHistory200ResponseInnerDestinationsEmail, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given GetNetworkAlertsHistory200ResponseInnerDestinationsEmail and assigns it to the Email field.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) SetEmail(v GetNetworkAlertsHistory200ResponseInnerDestinationsEmail) {
	o.Email = &v
}

// GetPush returns the Push field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetPush() GetNetworkAlertsHistory200ResponseInnerDestinationsPush {
	if o == nil || IsNil(o.Push) {
		var ret GetNetworkAlertsHistory200ResponseInnerDestinationsPush
		return ret
	}
	return *o.Push
}

// GetPushOk returns a tuple with the Push field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetPushOk() (*GetNetworkAlertsHistory200ResponseInnerDestinationsPush, bool) {
	if o == nil || IsNil(o.Push) {
		return nil, false
	}
	return o.Push, true
}

// HasPush returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) HasPush() bool {
	if o != nil && !IsNil(o.Push) {
		return true
	}

	return false
}

// SetPush gets a reference to the given GetNetworkAlertsHistory200ResponseInnerDestinationsPush and assigns it to the Push field.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) SetPush(v GetNetworkAlertsHistory200ResponseInnerDestinationsPush) {
	o.Push = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetSms() GetNetworkAlertsHistory200ResponseInnerDestinationsSms {
	if o == nil || IsNil(o.Sms) {
		var ret GetNetworkAlertsHistory200ResponseInnerDestinationsSms
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetSmsOk() (*GetNetworkAlertsHistory200ResponseInnerDestinationsSms, bool) {
	if o == nil || IsNil(o.Sms) {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) HasSms() bool {
	if o != nil && !IsNil(o.Sms) {
		return true
	}

	return false
}

// SetSms gets a reference to the given GetNetworkAlertsHistory200ResponseInnerDestinationsSms and assigns it to the Sms field.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) SetSms(v GetNetworkAlertsHistory200ResponseInnerDestinationsSms) {
	o.Sms = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetWebhook() GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook {
	if o == nil || IsNil(o.Webhook) {
		var ret GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) GetWebhookOk() (*GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook and assigns it to the Webhook field.
func (o *GetNetworkAlertsHistory200ResponseInnerDestinations) SetWebhook(v GetNetworkAlertsHistory200ResponseInnerDestinationsWebhook) {
	o.Webhook = &v
}

func (o GetNetworkAlertsHistory200ResponseInnerDestinations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAlertsHistory200ResponseInnerDestinations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Push) {
		toSerialize["push"] = o.Push
	}
	if !IsNil(o.Sms) {
		toSerialize["sms"] = o.Sms
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableGetNetworkAlertsHistory200ResponseInnerDestinations struct {
	value *GetNetworkAlertsHistory200ResponseInnerDestinations
	isSet bool
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinations) Get() *GetNetworkAlertsHistory200ResponseInnerDestinations {
	return v.value
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinations) Set(val *GetNetworkAlertsHistory200ResponseInnerDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAlertsHistory200ResponseInnerDestinations(val *GetNetworkAlertsHistory200ResponseInnerDestinations) *NullableGetNetworkAlertsHistory200ResponseInnerDestinations {
	return &NullableGetNetworkAlertsHistory200ResponseInnerDestinations{value: val, isSet: true}
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


