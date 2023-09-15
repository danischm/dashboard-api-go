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

// checks if the GetNetworkClients200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkClients200Response{}

// GetNetworkClients200Response struct for GetNetworkClients200Response
type GetNetworkClients200Response struct {
	// The ID of the client
	Id *string `json:"id,omitempty"`
	// The MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// The IP address of the client
	Ip *string `json:"ip,omitempty"`
	// The IPv6 address of the client
	Ip6 *string `json:"ip6,omitempty"`
	// Short description of the client
	Description *string `json:"description,omitempty"`
	// Timestamp client was first seen in the network
	FirstSeen *int32 `json:"firstSeen,omitempty"`
	// Timestamp client was last seen in the network
	LastSeen *int32 `json:"lastSeen,omitempty"`
	// Manufacturer of the client
	Manufacturer *string `json:"manufacturer,omitempty"`
	// The operating system of the client
	Os *string `json:"os,omitempty"`
	// The username of the user of the client
	User *string `json:"user,omitempty"`
	// The name of the VLAN that the client is connected to
	Vlan *string `json:"vlan,omitempty"`
	// The name of the SSID that the client is connected to
	Ssid *string `json:"ssid,omitempty"`
	// The switch port that the client is connected to
	Switchport *string `json:"switchport,omitempty"`
	// Wireless capabilities of the client
	WirelessCapabilities *string `json:"wirelessCapabilities,omitempty"`
	// Status of SM for the client
	SmInstalled *bool `json:"smInstalled,omitempty"`
	// The MAC address of the node that the device was last connected to
	RecentDeviceMac *string `json:"recentDeviceMac,omitempty"`
	// The connection status of the client
	Status *string `json:"status,omitempty"`
	Usage *GetNetworkClients200ResponseUsage `json:"usage,omitempty"`
	// Named VLAN of the client
	NamedVlan *string `json:"namedVlan,omitempty"`
	// The adaptive policy group of the client
	AdaptivePolicyGroup *string `json:"adaptivePolicyGroup,omitempty"`
	// Prediction of the client's device type
	DeviceTypePrediction *string `json:"deviceTypePrediction,omitempty"`
	// The serial of the node the device was last connected to
	RecentDeviceSerial *string `json:"recentDeviceSerial,omitempty"`
	// The name of the node the device was last connected to
	RecentDeviceName *string `json:"recentDeviceName,omitempty"`
	// Client's most recent connection type
	RecentDeviceConnection *string `json:"recentDeviceConnection,omitempty"`
	// Notes on the client
	Notes *string `json:"notes,omitempty"`
	// Local IPv6 address of the client
	Ip6Local *string `json:"ip6Local,omitempty"`
	// 802.1x group policy of the client
	GroupPolicy8021x *string `json:"groupPolicy8021x,omitempty"`
	// iPSK name of the client
	PskGroup *string `json:"pskGroup,omitempty"`
}

// NewGetNetworkClients200Response instantiates a new GetNetworkClients200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkClients200Response() *GetNetworkClients200Response {
	this := GetNetworkClients200Response{}
	return &this
}

// NewGetNetworkClients200ResponseWithDefaults instantiates a new GetNetworkClients200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkClients200ResponseWithDefaults() *GetNetworkClients200Response {
	this := GetNetworkClients200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkClients200Response) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkClients200Response) SetMac(v string) {
	o.Mac = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetNetworkClients200Response) SetIp(v string) {
	o.Ip = &v
}

// GetIp6 returns the Ip6 field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetIp6() string {
	if o == nil || IsNil(o.Ip6) {
		var ret string
		return ret
	}
	return *o.Ip6
}

// GetIp6Ok returns a tuple with the Ip6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetIp6Ok() (*string, bool) {
	if o == nil || IsNil(o.Ip6) {
		return nil, false
	}
	return o.Ip6, true
}

// HasIp6 returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasIp6() bool {
	if o != nil && !IsNil(o.Ip6) {
		return true
	}

	return false
}

// SetIp6 gets a reference to the given string and assigns it to the Ip6 field.
func (o *GetNetworkClients200Response) SetIp6(v string) {
	o.Ip6 = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkClients200Response) SetDescription(v string) {
	o.Description = &v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetFirstSeen() int32 {
	if o == nil || IsNil(o.FirstSeen) {
		var ret int32
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetFirstSeenOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstSeen) {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasFirstSeen() bool {
	if o != nil && !IsNil(o.FirstSeen) {
		return true
	}

	return false
}

// SetFirstSeen gets a reference to the given int32 and assigns it to the FirstSeen field.
func (o *GetNetworkClients200Response) SetFirstSeen(v int32) {
	o.FirstSeen = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetLastSeen() int32 {
	if o == nil || IsNil(o.LastSeen) {
		var ret int32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetLastSeenOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int32 and assigns it to the LastSeen field.
func (o *GetNetworkClients200Response) SetLastSeen(v int32) {
	o.LastSeen = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *GetNetworkClients200Response) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *GetNetworkClients200Response) SetOs(v string) {
	o.Os = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *GetNetworkClients200Response) SetUser(v string) {
	o.User = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetVlan() string {
	if o == nil || IsNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetVlanOk() (*string, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *GetNetworkClients200Response) SetVlan(v string) {
	o.Vlan = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *GetNetworkClients200Response) SetSsid(v string) {
	o.Ssid = &v
}

// GetSwitchport returns the Switchport field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetSwitchport() string {
	if o == nil || IsNil(o.Switchport) {
		var ret string
		return ret
	}
	return *o.Switchport
}

// GetSwitchportOk returns a tuple with the Switchport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetSwitchportOk() (*string, bool) {
	if o == nil || IsNil(o.Switchport) {
		return nil, false
	}
	return o.Switchport, true
}

// HasSwitchport returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasSwitchport() bool {
	if o != nil && !IsNil(o.Switchport) {
		return true
	}

	return false
}

// SetSwitchport gets a reference to the given string and assigns it to the Switchport field.
func (o *GetNetworkClients200Response) SetSwitchport(v string) {
	o.Switchport = &v
}

// GetWirelessCapabilities returns the WirelessCapabilities field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetWirelessCapabilities() string {
	if o == nil || IsNil(o.WirelessCapabilities) {
		var ret string
		return ret
	}
	return *o.WirelessCapabilities
}

// GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetWirelessCapabilitiesOk() (*string, bool) {
	if o == nil || IsNil(o.WirelessCapabilities) {
		return nil, false
	}
	return o.WirelessCapabilities, true
}

// HasWirelessCapabilities returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasWirelessCapabilities() bool {
	if o != nil && !IsNil(o.WirelessCapabilities) {
		return true
	}

	return false
}

// SetWirelessCapabilities gets a reference to the given string and assigns it to the WirelessCapabilities field.
func (o *GetNetworkClients200Response) SetWirelessCapabilities(v string) {
	o.WirelessCapabilities = &v
}

// GetSmInstalled returns the SmInstalled field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetSmInstalled() bool {
	if o == nil || IsNil(o.SmInstalled) {
		var ret bool
		return ret
	}
	return *o.SmInstalled
}

// GetSmInstalledOk returns a tuple with the SmInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetSmInstalledOk() (*bool, bool) {
	if o == nil || IsNil(o.SmInstalled) {
		return nil, false
	}
	return o.SmInstalled, true
}

// HasSmInstalled returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasSmInstalled() bool {
	if o != nil && !IsNil(o.SmInstalled) {
		return true
	}

	return false
}

// SetSmInstalled gets a reference to the given bool and assigns it to the SmInstalled field.
func (o *GetNetworkClients200Response) SetSmInstalled(v bool) {
	o.SmInstalled = &v
}

// GetRecentDeviceMac returns the RecentDeviceMac field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetRecentDeviceMac() string {
	if o == nil || IsNil(o.RecentDeviceMac) {
		var ret string
		return ret
	}
	return *o.RecentDeviceMac
}

// GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetRecentDeviceMacOk() (*string, bool) {
	if o == nil || IsNil(o.RecentDeviceMac) {
		return nil, false
	}
	return o.RecentDeviceMac, true
}

// HasRecentDeviceMac returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasRecentDeviceMac() bool {
	if o != nil && !IsNil(o.RecentDeviceMac) {
		return true
	}

	return false
}

// SetRecentDeviceMac gets a reference to the given string and assigns it to the RecentDeviceMac field.
func (o *GetNetworkClients200Response) SetRecentDeviceMac(v string) {
	o.RecentDeviceMac = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetNetworkClients200Response) SetStatus(v string) {
	o.Status = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetUsage() GetNetworkClients200ResponseUsage {
	if o == nil || IsNil(o.Usage) {
		var ret GetNetworkClients200ResponseUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetUsageOk() (*GetNetworkClients200ResponseUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetNetworkClients200ResponseUsage and assigns it to the Usage field.
func (o *GetNetworkClients200Response) SetUsage(v GetNetworkClients200ResponseUsage) {
	o.Usage = &v
}

// GetNamedVlan returns the NamedVlan field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetNamedVlan() string {
	if o == nil || IsNil(o.NamedVlan) {
		var ret string
		return ret
	}
	return *o.NamedVlan
}

// GetNamedVlanOk returns a tuple with the NamedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetNamedVlanOk() (*string, bool) {
	if o == nil || IsNil(o.NamedVlan) {
		return nil, false
	}
	return o.NamedVlan, true
}

// HasNamedVlan returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasNamedVlan() bool {
	if o != nil && !IsNil(o.NamedVlan) {
		return true
	}

	return false
}

// SetNamedVlan gets a reference to the given string and assigns it to the NamedVlan field.
func (o *GetNetworkClients200Response) SetNamedVlan(v string) {
	o.NamedVlan = &v
}

// GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetAdaptivePolicyGroup() string {
	if o == nil || IsNil(o.AdaptivePolicyGroup) {
		var ret string
		return ret
	}
	return *o.AdaptivePolicyGroup
}

// GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetAdaptivePolicyGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AdaptivePolicyGroup) {
		return nil, false
	}
	return o.AdaptivePolicyGroup, true
}

// HasAdaptivePolicyGroup returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasAdaptivePolicyGroup() bool {
	if o != nil && !IsNil(o.AdaptivePolicyGroup) {
		return true
	}

	return false
}

// SetAdaptivePolicyGroup gets a reference to the given string and assigns it to the AdaptivePolicyGroup field.
func (o *GetNetworkClients200Response) SetAdaptivePolicyGroup(v string) {
	o.AdaptivePolicyGroup = &v
}

// GetDeviceTypePrediction returns the DeviceTypePrediction field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetDeviceTypePrediction() string {
	if o == nil || IsNil(o.DeviceTypePrediction) {
		var ret string
		return ret
	}
	return *o.DeviceTypePrediction
}

// GetDeviceTypePredictionOk returns a tuple with the DeviceTypePrediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetDeviceTypePredictionOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceTypePrediction) {
		return nil, false
	}
	return o.DeviceTypePrediction, true
}

// HasDeviceTypePrediction returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasDeviceTypePrediction() bool {
	if o != nil && !IsNil(o.DeviceTypePrediction) {
		return true
	}

	return false
}

// SetDeviceTypePrediction gets a reference to the given string and assigns it to the DeviceTypePrediction field.
func (o *GetNetworkClients200Response) SetDeviceTypePrediction(v string) {
	o.DeviceTypePrediction = &v
}

// GetRecentDeviceSerial returns the RecentDeviceSerial field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetRecentDeviceSerial() string {
	if o == nil || IsNil(o.RecentDeviceSerial) {
		var ret string
		return ret
	}
	return *o.RecentDeviceSerial
}

// GetRecentDeviceSerialOk returns a tuple with the RecentDeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetRecentDeviceSerialOk() (*string, bool) {
	if o == nil || IsNil(o.RecentDeviceSerial) {
		return nil, false
	}
	return o.RecentDeviceSerial, true
}

// HasRecentDeviceSerial returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasRecentDeviceSerial() bool {
	if o != nil && !IsNil(o.RecentDeviceSerial) {
		return true
	}

	return false
}

// SetRecentDeviceSerial gets a reference to the given string and assigns it to the RecentDeviceSerial field.
func (o *GetNetworkClients200Response) SetRecentDeviceSerial(v string) {
	o.RecentDeviceSerial = &v
}

// GetRecentDeviceName returns the RecentDeviceName field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetRecentDeviceName() string {
	if o == nil || IsNil(o.RecentDeviceName) {
		var ret string
		return ret
	}
	return *o.RecentDeviceName
}

// GetRecentDeviceNameOk returns a tuple with the RecentDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetRecentDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.RecentDeviceName) {
		return nil, false
	}
	return o.RecentDeviceName, true
}

// HasRecentDeviceName returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasRecentDeviceName() bool {
	if o != nil && !IsNil(o.RecentDeviceName) {
		return true
	}

	return false
}

// SetRecentDeviceName gets a reference to the given string and assigns it to the RecentDeviceName field.
func (o *GetNetworkClients200Response) SetRecentDeviceName(v string) {
	o.RecentDeviceName = &v
}

// GetRecentDeviceConnection returns the RecentDeviceConnection field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetRecentDeviceConnection() string {
	if o == nil || IsNil(o.RecentDeviceConnection) {
		var ret string
		return ret
	}
	return *o.RecentDeviceConnection
}

// GetRecentDeviceConnectionOk returns a tuple with the RecentDeviceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetRecentDeviceConnectionOk() (*string, bool) {
	if o == nil || IsNil(o.RecentDeviceConnection) {
		return nil, false
	}
	return o.RecentDeviceConnection, true
}

// HasRecentDeviceConnection returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasRecentDeviceConnection() bool {
	if o != nil && !IsNil(o.RecentDeviceConnection) {
		return true
	}

	return false
}

// SetRecentDeviceConnection gets a reference to the given string and assigns it to the RecentDeviceConnection field.
func (o *GetNetworkClients200Response) SetRecentDeviceConnection(v string) {
	o.RecentDeviceConnection = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GetNetworkClients200Response) SetNotes(v string) {
	o.Notes = &v
}

// GetIp6Local returns the Ip6Local field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetIp6Local() string {
	if o == nil || IsNil(o.Ip6Local) {
		var ret string
		return ret
	}
	return *o.Ip6Local
}

// GetIp6LocalOk returns a tuple with the Ip6Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetIp6LocalOk() (*string, bool) {
	if o == nil || IsNil(o.Ip6Local) {
		return nil, false
	}
	return o.Ip6Local, true
}

// HasIp6Local returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasIp6Local() bool {
	if o != nil && !IsNil(o.Ip6Local) {
		return true
	}

	return false
}

// SetIp6Local gets a reference to the given string and assigns it to the Ip6Local field.
func (o *GetNetworkClients200Response) SetIp6Local(v string) {
	o.Ip6Local = &v
}

// GetGroupPolicy8021x returns the GroupPolicy8021x field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetGroupPolicy8021x() string {
	if o == nil || IsNil(o.GroupPolicy8021x) {
		var ret string
		return ret
	}
	return *o.GroupPolicy8021x
}

// GetGroupPolicy8021xOk returns a tuple with the GroupPolicy8021x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetGroupPolicy8021xOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicy8021x) {
		return nil, false
	}
	return o.GroupPolicy8021x, true
}

// HasGroupPolicy8021x returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasGroupPolicy8021x() bool {
	if o != nil && !IsNil(o.GroupPolicy8021x) {
		return true
	}

	return false
}

// SetGroupPolicy8021x gets a reference to the given string and assigns it to the GroupPolicy8021x field.
func (o *GetNetworkClients200Response) SetGroupPolicy8021x(v string) {
	o.GroupPolicy8021x = &v
}

// GetPskGroup returns the PskGroup field value if set, zero value otherwise.
func (o *GetNetworkClients200Response) GetPskGroup() string {
	if o == nil || IsNil(o.PskGroup) {
		var ret string
		return ret
	}
	return *o.PskGroup
}

// GetPskGroupOk returns a tuple with the PskGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200Response) GetPskGroupOk() (*string, bool) {
	if o == nil || IsNil(o.PskGroup) {
		return nil, false
	}
	return o.PskGroup, true
}

// HasPskGroup returns a boolean if a field has been set.
func (o *GetNetworkClients200Response) HasPskGroup() bool {
	if o != nil && !IsNil(o.PskGroup) {
		return true
	}

	return false
}

// SetPskGroup gets a reference to the given string and assigns it to the PskGroup field.
func (o *GetNetworkClients200Response) SetPskGroup(v string) {
	o.PskGroup = &v
}

func (o GetNetworkClients200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkClients200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ip6) {
		toSerialize["ip6"] = o.Ip6
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FirstSeen) {
		toSerialize["firstSeen"] = o.FirstSeen
	}
	if !IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Switchport) {
		toSerialize["switchport"] = o.Switchport
	}
	if !IsNil(o.WirelessCapabilities) {
		toSerialize["wirelessCapabilities"] = o.WirelessCapabilities
	}
	if !IsNil(o.SmInstalled) {
		toSerialize["smInstalled"] = o.SmInstalled
	}
	if !IsNil(o.RecentDeviceMac) {
		toSerialize["recentDeviceMac"] = o.RecentDeviceMac
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.NamedVlan) {
		toSerialize["namedVlan"] = o.NamedVlan
	}
	if !IsNil(o.AdaptivePolicyGroup) {
		toSerialize["adaptivePolicyGroup"] = o.AdaptivePolicyGroup
	}
	if !IsNil(o.DeviceTypePrediction) {
		toSerialize["deviceTypePrediction"] = o.DeviceTypePrediction
	}
	if !IsNil(o.RecentDeviceSerial) {
		toSerialize["recentDeviceSerial"] = o.RecentDeviceSerial
	}
	if !IsNil(o.RecentDeviceName) {
		toSerialize["recentDeviceName"] = o.RecentDeviceName
	}
	if !IsNil(o.RecentDeviceConnection) {
		toSerialize["recentDeviceConnection"] = o.RecentDeviceConnection
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Ip6Local) {
		toSerialize["ip6Local"] = o.Ip6Local
	}
	if !IsNil(o.GroupPolicy8021x) {
		toSerialize["groupPolicy8021x"] = o.GroupPolicy8021x
	}
	if !IsNil(o.PskGroup) {
		toSerialize["pskGroup"] = o.PskGroup
	}
	return toSerialize, nil
}

type NullableGetNetworkClients200Response struct {
	value *GetNetworkClients200Response
	isSet bool
}

func (v NullableGetNetworkClients200Response) Get() *GetNetworkClients200Response {
	return v.value
}

func (v *NullableGetNetworkClients200Response) Set(val *GetNetworkClients200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkClients200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkClients200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkClients200Response(val *GetNetworkClients200Response) *NullableGetNetworkClients200Response {
	return &NullableGetNetworkClients200Response{value: val, isSet: true}
}

func (v NullableGetNetworkClients200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkClients200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


