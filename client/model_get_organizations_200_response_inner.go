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

// checks if the GetOrganizations200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizations200ResponseInner{}

// GetOrganizations200ResponseInner struct for GetOrganizations200ResponseInner
type GetOrganizations200ResponseInner struct {
	// Organization ID
	Id *string `json:"id,omitempty"`
	// Organization name
	Name *string `json:"name,omitempty"`
	// Organization URL
	Url *string `json:"url,omitempty"`
	Api *GetOrganizations200ResponseInnerApi `json:"api,omitempty"`
	Licensing *GetOrganizations200ResponseInnerLicensing `json:"licensing,omitempty"`
	Cloud *GetOrganizations200ResponseInnerCloud `json:"cloud,omitempty"`
	Management *GetOrganizations200ResponseInnerManagement `json:"management,omitempty"`
}

// NewGetOrganizations200ResponseInner instantiates a new GetOrganizations200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizations200ResponseInner() *GetOrganizations200ResponseInner {
	this := GetOrganizations200ResponseInner{}
	return &this
}

// NewGetOrganizations200ResponseInnerWithDefaults instantiates a new GetOrganizations200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizations200ResponseInnerWithDefaults() *GetOrganizations200ResponseInner {
	this := GetOrganizations200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizations200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizations200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizations200ResponseInner) SetUrl(v string) {
	o.Url = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetApi() GetOrganizations200ResponseInnerApi {
	if o == nil || IsNil(o.Api) {
		var ret GetOrganizations200ResponseInnerApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetApiOk() (*GetOrganizations200ResponseInnerApi, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given GetOrganizations200ResponseInnerApi and assigns it to the Api field.
func (o *GetOrganizations200ResponseInner) SetApi(v GetOrganizations200ResponseInnerApi) {
	o.Api = &v
}

// GetLicensing returns the Licensing field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetLicensing() GetOrganizations200ResponseInnerLicensing {
	if o == nil || IsNil(o.Licensing) {
		var ret GetOrganizations200ResponseInnerLicensing
		return ret
	}
	return *o.Licensing
}

// GetLicensingOk returns a tuple with the Licensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetLicensingOk() (*GetOrganizations200ResponseInnerLicensing, bool) {
	if o == nil || IsNil(o.Licensing) {
		return nil, false
	}
	return o.Licensing, true
}

// HasLicensing returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasLicensing() bool {
	if o != nil && !IsNil(o.Licensing) {
		return true
	}

	return false
}

// SetLicensing gets a reference to the given GetOrganizations200ResponseInnerLicensing and assigns it to the Licensing field.
func (o *GetOrganizations200ResponseInner) SetLicensing(v GetOrganizations200ResponseInnerLicensing) {
	o.Licensing = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetCloud() GetOrganizations200ResponseInnerCloud {
	if o == nil || IsNil(o.Cloud) {
		var ret GetOrganizations200ResponseInnerCloud
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetCloudOk() (*GetOrganizations200ResponseInnerCloud, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given GetOrganizations200ResponseInnerCloud and assigns it to the Cloud field.
func (o *GetOrganizations200ResponseInner) SetCloud(v GetOrganizations200ResponseInnerCloud) {
	o.Cloud = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInner) GetManagement() GetOrganizations200ResponseInnerManagement {
	if o == nil || IsNil(o.Management) {
		var ret GetOrganizations200ResponseInnerManagement
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInner) GetManagementOk() (*GetOrganizations200ResponseInnerManagement, bool) {
	if o == nil || IsNil(o.Management) {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInner) HasManagement() bool {
	if o != nil && !IsNil(o.Management) {
		return true
	}

	return false
}

// SetManagement gets a reference to the given GetOrganizations200ResponseInnerManagement and assigns it to the Management field.
func (o *GetOrganizations200ResponseInner) SetManagement(v GetOrganizations200ResponseInnerManagement) {
	o.Management = &v
}

func (o GetOrganizations200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizations200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.Licensing) {
		toSerialize["licensing"] = o.Licensing
	}
	if !IsNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !IsNil(o.Management) {
		toSerialize["management"] = o.Management
	}
	return toSerialize, nil
}

type NullableGetOrganizations200ResponseInner struct {
	value *GetOrganizations200ResponseInner
	isSet bool
}

func (v NullableGetOrganizations200ResponseInner) Get() *GetOrganizations200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizations200ResponseInner) Set(val *GetOrganizations200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizations200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizations200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizations200ResponseInner(val *GetOrganizations200ResponseInner) *NullableGetOrganizations200ResponseInner {
	return &NullableGetOrganizations200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizations200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizations200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

