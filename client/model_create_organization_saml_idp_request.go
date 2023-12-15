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

// checks if the CreateOrganizationSamlIdpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationSamlIdpRequest{}

// CreateOrganizationSamlIdpRequest struct for CreateOrganizationSamlIdpRequest
type CreateOrganizationSamlIdpRequest struct {
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint string `json:"x509certSha1Fingerprint"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewCreateOrganizationSamlIdpRequest instantiates a new CreateOrganizationSamlIdpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationSamlIdpRequest(x509certSha1Fingerprint string) *CreateOrganizationSamlIdpRequest {
	this := CreateOrganizationSamlIdpRequest{}
	this.X509certSha1Fingerprint = x509certSha1Fingerprint
	return &this
}

// NewCreateOrganizationSamlIdpRequestWithDefaults instantiates a new CreateOrganizationSamlIdpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationSamlIdpRequestWithDefaults() *CreateOrganizationSamlIdpRequest {
	this := CreateOrganizationSamlIdpRequest{}
	return &this
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value
func (o *CreateOrganizationSamlIdpRequest) GetX509certSha1Fingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlIdpRequest) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X509certSha1Fingerprint, true
}

// SetX509certSha1Fingerprint sets field value
func (o *CreateOrganizationSamlIdpRequest) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *CreateOrganizationSamlIdpRequest) GetSloLogoutUrl() string {
	if o == nil || IsNil(o.SloLogoutUrl) {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlIdpRequest) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SloLogoutUrl) {
		return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *CreateOrganizationSamlIdpRequest) HasSloLogoutUrl() bool {
	if o != nil && !IsNil(o.SloLogoutUrl) {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *CreateOrganizationSamlIdpRequest) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o CreateOrganizationSamlIdpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationSamlIdpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	if !IsNil(o.SloLogoutUrl) {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return toSerialize, nil
}

type NullableCreateOrganizationSamlIdpRequest struct {
	value *CreateOrganizationSamlIdpRequest
	isSet bool
}

func (v NullableCreateOrganizationSamlIdpRequest) Get() *CreateOrganizationSamlIdpRequest {
	return v.value
}

func (v *NullableCreateOrganizationSamlIdpRequest) Set(val *CreateOrganizationSamlIdpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationSamlIdpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationSamlIdpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationSamlIdpRequest(val *CreateOrganizationSamlIdpRequest) *NullableCreateOrganizationSamlIdpRequest {
	return &NullableCreateOrganizationSamlIdpRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationSamlIdpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationSamlIdpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


