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

// InlineResponse20054 struct for InlineResponse20054
type InlineResponse20054 struct {
	// The Meraki managed application Id for this record on a particular device.
	AppId *string `json:"appId,omitempty"`
	// The size of the software bundle.
	BundleSize *int32 `json:"bundleSize,omitempty"`
	// When the Meraki record for the software was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The Meraki managed device Id.
	DeviceId *string `json:"deviceId,omitempty"`
	// The size of the data stored in the application.
	DynamicSize *int32 `json:"dynamicSize,omitempty"`
	// The Meraki software Id.
	Id *string `json:"id,omitempty"`
	// Software bundle identifier.
	Identifier *string `json:"identifier,omitempty"`
	// When the Software was installed on the device.
	InstalledAt *string `json:"installedAt,omitempty"`
	// A boolean indicating this software record should be installed on the associated device.
	ToInstall *bool `json:"toInstall,omitempty"`
	// A boolean indicating whether or not an iOS redemption code was used.
	IosRedemptionCode *bool `json:"iosRedemptionCode,omitempty"`
	// A boolean indicating whether or not the software is managed by Meraki.
	IsManaged *bool `json:"isManaged,omitempty"`
	// The itunes numerical identifier.
	ItunesId *string `json:"itunesId,omitempty"`
	// The license key associated with this software installation.
	LicenseKey *string `json:"licenseKey,omitempty"`
	// The name of the software.
	Name *string `json:"name,omitempty"`
	// The path on the device where the software record is located.
	Path *string `json:"path,omitempty"`
	// The redemption code used for this software.
	RedemptionCode *int32 `json:"redemptionCode,omitempty"`
	// Short version notation for the software.
	ShortVersion *string `json:"shortVersion,omitempty"`
	// The management status of the software.
	Status *string `json:"status,omitempty"`
	// A boolean indicating this software record should be uninstalled on the associated device.
	ToUninstall *bool `json:"toUninstall,omitempty"`
	// When the record was uninstalled from the device.
	UninstalledAt *string `json:"uninstalledAt,omitempty"`
	// When the record was last updated by Meraki.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The vendor of the software.
	Vendor *string `json:"vendor,omitempty"`
	// Full version notation for the software.
	Version *string `json:"version,omitempty"`
}

// NewInlineResponse20054 instantiates a new InlineResponse20054 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20054() *InlineResponse20054 {
	this := InlineResponse20054{}
	return &this
}

// NewInlineResponse20054WithDefaults instantiates a new InlineResponse20054 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20054WithDefaults() *InlineResponse20054 {
	this := InlineResponse20054{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *InlineResponse20054) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *InlineResponse20054) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *InlineResponse20054) SetAppId(v string) {
	o.AppId = &v
}

// GetBundleSize returns the BundleSize field value if set, zero value otherwise.
func (o *InlineResponse20054) GetBundleSize() int32 {
	if o == nil || isNil(o.BundleSize) {
		var ret int32
		return ret
	}
	return *o.BundleSize
}

// GetBundleSizeOk returns a tuple with the BundleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetBundleSizeOk() (*int32, bool) {
	if o == nil || isNil(o.BundleSize) {
    return nil, false
	}
	return o.BundleSize, true
}

// HasBundleSize returns a boolean if a field has been set.
func (o *InlineResponse20054) HasBundleSize() bool {
	if o != nil && !isNil(o.BundleSize) {
		return true
	}

	return false
}

// SetBundleSize gets a reference to the given int32 and assigns it to the BundleSize field.
func (o *InlineResponse20054) SetBundleSize(v int32) {
	o.BundleSize = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse20054) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse20054) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse20054) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20054) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20054) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20054) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDynamicSize returns the DynamicSize field value if set, zero value otherwise.
func (o *InlineResponse20054) GetDynamicSize() int32 {
	if o == nil || isNil(o.DynamicSize) {
		var ret int32
		return ret
	}
	return *o.DynamicSize
}

// GetDynamicSizeOk returns a tuple with the DynamicSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetDynamicSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DynamicSize) {
    return nil, false
	}
	return o.DynamicSize, true
}

// HasDynamicSize returns a boolean if a field has been set.
func (o *InlineResponse20054) HasDynamicSize() bool {
	if o != nil && !isNil(o.DynamicSize) {
		return true
	}

	return false
}

// SetDynamicSize gets a reference to the given int32 and assigns it to the DynamicSize field.
func (o *InlineResponse20054) SetDynamicSize(v int32) {
	o.DynamicSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20054) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20054) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20054) SetId(v string) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *InlineResponse20054) GetIdentifier() string {
	if o == nil || isNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.Identifier) {
    return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *InlineResponse20054) HasIdentifier() bool {
	if o != nil && !isNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *InlineResponse20054) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetInstalledAt returns the InstalledAt field value if set, zero value otherwise.
func (o *InlineResponse20054) GetInstalledAt() string {
	if o == nil || isNil(o.InstalledAt) {
		var ret string
		return ret
	}
	return *o.InstalledAt
}

// GetInstalledAtOk returns a tuple with the InstalledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetInstalledAtOk() (*string, bool) {
	if o == nil || isNil(o.InstalledAt) {
    return nil, false
	}
	return o.InstalledAt, true
}

// HasInstalledAt returns a boolean if a field has been set.
func (o *InlineResponse20054) HasInstalledAt() bool {
	if o != nil && !isNil(o.InstalledAt) {
		return true
	}

	return false
}

// SetInstalledAt gets a reference to the given string and assigns it to the InstalledAt field.
func (o *InlineResponse20054) SetInstalledAt(v string) {
	o.InstalledAt = &v
}

// GetToInstall returns the ToInstall field value if set, zero value otherwise.
func (o *InlineResponse20054) GetToInstall() bool {
	if o == nil || isNil(o.ToInstall) {
		var ret bool
		return ret
	}
	return *o.ToInstall
}

// GetToInstallOk returns a tuple with the ToInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetToInstallOk() (*bool, bool) {
	if o == nil || isNil(o.ToInstall) {
    return nil, false
	}
	return o.ToInstall, true
}

// HasToInstall returns a boolean if a field has been set.
func (o *InlineResponse20054) HasToInstall() bool {
	if o != nil && !isNil(o.ToInstall) {
		return true
	}

	return false
}

// SetToInstall gets a reference to the given bool and assigns it to the ToInstall field.
func (o *InlineResponse20054) SetToInstall(v bool) {
	o.ToInstall = &v
}

// GetIosRedemptionCode returns the IosRedemptionCode field value if set, zero value otherwise.
func (o *InlineResponse20054) GetIosRedemptionCode() bool {
	if o == nil || isNil(o.IosRedemptionCode) {
		var ret bool
		return ret
	}
	return *o.IosRedemptionCode
}

// GetIosRedemptionCodeOk returns a tuple with the IosRedemptionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetIosRedemptionCodeOk() (*bool, bool) {
	if o == nil || isNil(o.IosRedemptionCode) {
    return nil, false
	}
	return o.IosRedemptionCode, true
}

// HasIosRedemptionCode returns a boolean if a field has been set.
func (o *InlineResponse20054) HasIosRedemptionCode() bool {
	if o != nil && !isNil(o.IosRedemptionCode) {
		return true
	}

	return false
}

// SetIosRedemptionCode gets a reference to the given bool and assigns it to the IosRedemptionCode field.
func (o *InlineResponse20054) SetIosRedemptionCode(v bool) {
	o.IosRedemptionCode = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *InlineResponse20054) GetIsManaged() bool {
	if o == nil || isNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetIsManagedOk() (*bool, bool) {
	if o == nil || isNil(o.IsManaged) {
    return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *InlineResponse20054) HasIsManaged() bool {
	if o != nil && !isNil(o.IsManaged) {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *InlineResponse20054) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetItunesId returns the ItunesId field value if set, zero value otherwise.
func (o *InlineResponse20054) GetItunesId() string {
	if o == nil || isNil(o.ItunesId) {
		var ret string
		return ret
	}
	return *o.ItunesId
}

// GetItunesIdOk returns a tuple with the ItunesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetItunesIdOk() (*string, bool) {
	if o == nil || isNil(o.ItunesId) {
    return nil, false
	}
	return o.ItunesId, true
}

// HasItunesId returns a boolean if a field has been set.
func (o *InlineResponse20054) HasItunesId() bool {
	if o != nil && !isNil(o.ItunesId) {
		return true
	}

	return false
}

// SetItunesId gets a reference to the given string and assigns it to the ItunesId field.
func (o *InlineResponse20054) SetItunesId(v string) {
	o.ItunesId = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise.
func (o *InlineResponse20054) GetLicenseKey() string {
	if o == nil || isNil(o.LicenseKey) {
		var ret string
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetLicenseKeyOk() (*string, bool) {
	if o == nil || isNil(o.LicenseKey) {
    return nil, false
	}
	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *InlineResponse20054) HasLicenseKey() bool {
	if o != nil && !isNil(o.LicenseKey) {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given string and assigns it to the LicenseKey field.
func (o *InlineResponse20054) SetLicenseKey(v string) {
	o.LicenseKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20054) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20054) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20054) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *InlineResponse20054) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *InlineResponse20054) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *InlineResponse20054) SetPath(v string) {
	o.Path = &v
}

// GetRedemptionCode returns the RedemptionCode field value if set, zero value otherwise.
func (o *InlineResponse20054) GetRedemptionCode() int32 {
	if o == nil || isNil(o.RedemptionCode) {
		var ret int32
		return ret
	}
	return *o.RedemptionCode
}

// GetRedemptionCodeOk returns a tuple with the RedemptionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetRedemptionCodeOk() (*int32, bool) {
	if o == nil || isNil(o.RedemptionCode) {
    return nil, false
	}
	return o.RedemptionCode, true
}

// HasRedemptionCode returns a boolean if a field has been set.
func (o *InlineResponse20054) HasRedemptionCode() bool {
	if o != nil && !isNil(o.RedemptionCode) {
		return true
	}

	return false
}

// SetRedemptionCode gets a reference to the given int32 and assigns it to the RedemptionCode field.
func (o *InlineResponse20054) SetRedemptionCode(v int32) {
	o.RedemptionCode = &v
}

// GetShortVersion returns the ShortVersion field value if set, zero value otherwise.
func (o *InlineResponse20054) GetShortVersion() string {
	if o == nil || isNil(o.ShortVersion) {
		var ret string
		return ret
	}
	return *o.ShortVersion
}

// GetShortVersionOk returns a tuple with the ShortVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetShortVersionOk() (*string, bool) {
	if o == nil || isNil(o.ShortVersion) {
    return nil, false
	}
	return o.ShortVersion, true
}

// HasShortVersion returns a boolean if a field has been set.
func (o *InlineResponse20054) HasShortVersion() bool {
	if o != nil && !isNil(o.ShortVersion) {
		return true
	}

	return false
}

// SetShortVersion gets a reference to the given string and assigns it to the ShortVersion field.
func (o *InlineResponse20054) SetShortVersion(v string) {
	o.ShortVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20054) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20054) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20054) SetStatus(v string) {
	o.Status = &v
}

// GetToUninstall returns the ToUninstall field value if set, zero value otherwise.
func (o *InlineResponse20054) GetToUninstall() bool {
	if o == nil || isNil(o.ToUninstall) {
		var ret bool
		return ret
	}
	return *o.ToUninstall
}

// GetToUninstallOk returns a tuple with the ToUninstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetToUninstallOk() (*bool, bool) {
	if o == nil || isNil(o.ToUninstall) {
    return nil, false
	}
	return o.ToUninstall, true
}

// HasToUninstall returns a boolean if a field has been set.
func (o *InlineResponse20054) HasToUninstall() bool {
	if o != nil && !isNil(o.ToUninstall) {
		return true
	}

	return false
}

// SetToUninstall gets a reference to the given bool and assigns it to the ToUninstall field.
func (o *InlineResponse20054) SetToUninstall(v bool) {
	o.ToUninstall = &v
}

// GetUninstalledAt returns the UninstalledAt field value if set, zero value otherwise.
func (o *InlineResponse20054) GetUninstalledAt() string {
	if o == nil || isNil(o.UninstalledAt) {
		var ret string
		return ret
	}
	return *o.UninstalledAt
}

// GetUninstalledAtOk returns a tuple with the UninstalledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetUninstalledAtOk() (*string, bool) {
	if o == nil || isNil(o.UninstalledAt) {
    return nil, false
	}
	return o.UninstalledAt, true
}

// HasUninstalledAt returns a boolean if a field has been set.
func (o *InlineResponse20054) HasUninstalledAt() bool {
	if o != nil && !isNil(o.UninstalledAt) {
		return true
	}

	return false
}

// SetUninstalledAt gets a reference to the given string and assigns it to the UninstalledAt field.
func (o *InlineResponse20054) SetUninstalledAt(v string) {
	o.UninstalledAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse20054) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse20054) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *InlineResponse20054) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *InlineResponse20054) GetVendor() string {
	if o == nil || isNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetVendorOk() (*string, bool) {
	if o == nil || isNil(o.Vendor) {
    return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *InlineResponse20054) HasVendor() bool {
	if o != nil && !isNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *InlineResponse20054) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20054) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20054) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse20054) SetVersion(v string) {
	o.Version = &v
}

func (o InlineResponse20054) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.BundleSize) {
		toSerialize["bundleSize"] = o.BundleSize
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.DynamicSize) {
		toSerialize["dynamicSize"] = o.DynamicSize
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !isNil(o.InstalledAt) {
		toSerialize["installedAt"] = o.InstalledAt
	}
	if !isNil(o.ToInstall) {
		toSerialize["toInstall"] = o.ToInstall
	}
	if !isNil(o.IosRedemptionCode) {
		toSerialize["iosRedemptionCode"] = o.IosRedemptionCode
	}
	if !isNil(o.IsManaged) {
		toSerialize["isManaged"] = o.IsManaged
	}
	if !isNil(o.ItunesId) {
		toSerialize["itunesId"] = o.ItunesId
	}
	if !isNil(o.LicenseKey) {
		toSerialize["licenseKey"] = o.LicenseKey
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.RedemptionCode) {
		toSerialize["redemptionCode"] = o.RedemptionCode
	}
	if !isNil(o.ShortVersion) {
		toSerialize["shortVersion"] = o.ShortVersion
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ToUninstall) {
		toSerialize["toUninstall"] = o.ToUninstall
	}
	if !isNil(o.UninstalledAt) {
		toSerialize["uninstalledAt"] = o.UninstalledAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20054 struct {
	value *InlineResponse20054
	isSet bool
}

func (v NullableInlineResponse20054) Get() *InlineResponse20054 {
	return v.value
}

func (v *NullableInlineResponse20054) Set(val *InlineResponse20054) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20054) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20054) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20054(val *InlineResponse20054) *NullableInlineResponse20054 {
	return &NullableInlineResponse20054{value: val, isSet: true}
}

func (v NullableInlineResponse20054) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20054) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


