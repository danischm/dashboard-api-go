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

// InlineResponse20046 struct for InlineResponse20046
type InlineResponse20046 struct {
	// The name of the certificate.
	Name *string `json:"name,omitempty"`
	// The date after which the certificate is no longer valid.
	NotValidAfter *string `json:"notValidAfter,omitempty"`
	// The date before which the certificate is not valid.
	NotValidBefore *string `json:"notValidBefore,omitempty"`
	// The PEM of the certificate.
	CertPem *string `json:"certPem,omitempty"`
	// The Meraki managed device Id.
	DeviceId *string `json:"deviceId,omitempty"`
	// The certificate issuer.
	Issuer *string `json:"issuer,omitempty"`
	// The subject of the certificate.
	Subject *string `json:"subject,omitempty"`
	// The Meraki Id of the certificate record.
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse20046 instantiates a new InlineResponse20046 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20046() *InlineResponse20046 {
	this := InlineResponse20046{}
	return &this
}

// NewInlineResponse20046WithDefaults instantiates a new InlineResponse20046 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20046WithDefaults() *InlineResponse20046 {
	this := InlineResponse20046{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20046) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20046) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20046) SetName(v string) {
	o.Name = &v
}

// GetNotValidAfter returns the NotValidAfter field value if set, zero value otherwise.
func (o *InlineResponse20046) GetNotValidAfter() string {
	if o == nil || isNil(o.NotValidAfter) {
		var ret string
		return ret
	}
	return *o.NotValidAfter
}

// GetNotValidAfterOk returns a tuple with the NotValidAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetNotValidAfterOk() (*string, bool) {
	if o == nil || isNil(o.NotValidAfter) {
    return nil, false
	}
	return o.NotValidAfter, true
}

// HasNotValidAfter returns a boolean if a field has been set.
func (o *InlineResponse20046) HasNotValidAfter() bool {
	if o != nil && !isNil(o.NotValidAfter) {
		return true
	}

	return false
}

// SetNotValidAfter gets a reference to the given string and assigns it to the NotValidAfter field.
func (o *InlineResponse20046) SetNotValidAfter(v string) {
	o.NotValidAfter = &v
}

// GetNotValidBefore returns the NotValidBefore field value if set, zero value otherwise.
func (o *InlineResponse20046) GetNotValidBefore() string {
	if o == nil || isNil(o.NotValidBefore) {
		var ret string
		return ret
	}
	return *o.NotValidBefore
}

// GetNotValidBeforeOk returns a tuple with the NotValidBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetNotValidBeforeOk() (*string, bool) {
	if o == nil || isNil(o.NotValidBefore) {
    return nil, false
	}
	return o.NotValidBefore, true
}

// HasNotValidBefore returns a boolean if a field has been set.
func (o *InlineResponse20046) HasNotValidBefore() bool {
	if o != nil && !isNil(o.NotValidBefore) {
		return true
	}

	return false
}

// SetNotValidBefore gets a reference to the given string and assigns it to the NotValidBefore field.
func (o *InlineResponse20046) SetNotValidBefore(v string) {
	o.NotValidBefore = &v
}

// GetCertPem returns the CertPem field value if set, zero value otherwise.
func (o *InlineResponse20046) GetCertPem() string {
	if o == nil || isNil(o.CertPem) {
		var ret string
		return ret
	}
	return *o.CertPem
}

// GetCertPemOk returns a tuple with the CertPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetCertPemOk() (*string, bool) {
	if o == nil || isNil(o.CertPem) {
    return nil, false
	}
	return o.CertPem, true
}

// HasCertPem returns a boolean if a field has been set.
func (o *InlineResponse20046) HasCertPem() bool {
	if o != nil && !isNil(o.CertPem) {
		return true
	}

	return false
}

// SetCertPem gets a reference to the given string and assigns it to the CertPem field.
func (o *InlineResponse20046) SetCertPem(v string) {
	o.CertPem = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20046) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20046) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20046) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *InlineResponse20046) GetIssuer() string {
	if o == nil || isNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetIssuerOk() (*string, bool) {
	if o == nil || isNil(o.Issuer) {
    return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *InlineResponse20046) HasIssuer() bool {
	if o != nil && !isNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *InlineResponse20046) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *InlineResponse20046) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
    return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *InlineResponse20046) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *InlineResponse20046) SetSubject(v string) {
	o.Subject = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20046) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20046) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20046) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse20046) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NotValidAfter) {
		toSerialize["notValidAfter"] = o.NotValidAfter
	}
	if !isNil(o.NotValidBefore) {
		toSerialize["notValidBefore"] = o.NotValidBefore
	}
	if !isNil(o.CertPem) {
		toSerialize["certPem"] = o.CertPem
	}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20046 struct {
	value *InlineResponse20046
	isSet bool
}

func (v NullableInlineResponse20046) Get() *InlineResponse20046 {
	return v.value
}

func (v *NullableInlineResponse20046) Set(val *InlineResponse20046) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20046) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20046) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20046(val *InlineResponse20046) *NullableInlineResponse20046 {
	return &NullableInlineResponse20046{value: val, isSet: true}
}

func (v NullableInlineResponse20046) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20046) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


