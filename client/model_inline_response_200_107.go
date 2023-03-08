/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200107 struct for InlineResponse200107
type InlineResponse200107 struct {
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Model type of the device
	Model *string `json:"model,omitempty"`
	// Network Id of the device
	NetworkId *string `json:"networkId,omitempty"`
	// Order number of the device
	OrderNumber *string `json:"orderNumber,omitempty"`
	// Claimed time of the device
	ClaimedAt *time.Time `json:"claimedAt,omitempty"`
	// License expiration date of the device
	LicenseExpirationDate *time.Time `json:"licenseExpirationDate,omitempty"`
	// Device tags
	Tags []string `json:"tags,omitempty"`
	// Product type of the device
	ProductType *string `json:"productType,omitempty"`
}

// NewInlineResponse200107 instantiates a new InlineResponse200107 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200107() *InlineResponse200107 {
	this := InlineResponse200107{}
	return &this
}

// NewInlineResponse200107WithDefaults instantiates a new InlineResponse200107 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200107WithDefaults() *InlineResponse200107 {
	this := InlineResponse200107{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200107) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200107) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200107) SetMac(v string) {
	o.Mac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200107) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200107) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200107) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200107) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200107) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200107) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200107) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200107) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200107) SetModel(v string) {
	o.Model = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200107) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200107) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200107) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *InlineResponse200107) GetOrderNumber() string {
	if o == nil || isNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetOrderNumberOk() (*string, bool) {
	if o == nil || isNil(o.OrderNumber) {
    return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *InlineResponse200107) HasOrderNumber() bool {
	if o != nil && !isNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *InlineResponse200107) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetClaimedAt returns the ClaimedAt field value if set, zero value otherwise.
func (o *InlineResponse200107) GetClaimedAt() time.Time {
	if o == nil || isNil(o.ClaimedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClaimedAt
}

// GetClaimedAtOk returns a tuple with the ClaimedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetClaimedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ClaimedAt) {
    return nil, false
	}
	return o.ClaimedAt, true
}

// HasClaimedAt returns a boolean if a field has been set.
func (o *InlineResponse200107) HasClaimedAt() bool {
	if o != nil && !isNil(o.ClaimedAt) {
		return true
	}

	return false
}

// SetClaimedAt gets a reference to the given time.Time and assigns it to the ClaimedAt field.
func (o *InlineResponse200107) SetClaimedAt(v time.Time) {
	o.ClaimedAt = &v
}

// GetLicenseExpirationDate returns the LicenseExpirationDate field value if set, zero value otherwise.
func (o *InlineResponse200107) GetLicenseExpirationDate() time.Time {
	if o == nil || isNil(o.LicenseExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.LicenseExpirationDate
}

// GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetLicenseExpirationDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.LicenseExpirationDate) {
    return nil, false
	}
	return o.LicenseExpirationDate, true
}

// HasLicenseExpirationDate returns a boolean if a field has been set.
func (o *InlineResponse200107) HasLicenseExpirationDate() bool {
	if o != nil && !isNil(o.LicenseExpirationDate) {
		return true
	}

	return false
}

// SetLicenseExpirationDate gets a reference to the given time.Time and assigns it to the LicenseExpirationDate field.
func (o *InlineResponse200107) SetLicenseExpirationDate(v time.Time) {
	o.LicenseExpirationDate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200107) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200107) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200107) SetTags(v []string) {
	o.Tags = v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse200107) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse200107) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse200107) SetProductType(v string) {
	o.ProductType = &v
}

func (o InlineResponse200107) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !isNil(o.ClaimedAt) {
		toSerialize["claimedAt"] = o.ClaimedAt
	}
	if !isNil(o.LicenseExpirationDate) {
		toSerialize["licenseExpirationDate"] = o.LicenseExpirationDate
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200107 struct {
	value *InlineResponse200107
	isSet bool
}

func (v NullableInlineResponse200107) Get() *InlineResponse200107 {
	return v.value
}

func (v *NullableInlineResponse200107) Set(val *InlineResponse200107) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200107) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200107) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200107(val *InlineResponse200107) *NullableInlineResponse200107 {
	return &NullableInlineResponse200107{value: val, isSet: true}
}

func (v NullableInlineResponse200107) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200107) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


