# InlineObject136

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name or description for layer 3 static route | [optional] 
**Subnet** | Pointer to **string** | The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24) | [optional] 
**NextHopIp** | Pointer to **string** | IP address of the next hop device to which the device sends its traffic for the subnet | [optional] 
**AdvertiseViaOspfEnabled** | Pointer to **bool** | Option to advertise static route via OSPF | [optional] 
**PreferOverOspfRoutesEnabled** | Pointer to **bool** | Option to prefer static route over OSPF routes | [optional] 

## Methods

### NewInlineObject136

`func NewInlineObject136() *InlineObject136`

NewInlineObject136 instantiates a new InlineObject136 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject136WithDefaults

`func NewInlineObject136WithDefaults() *InlineObject136`

NewInlineObject136WithDefaults instantiates a new InlineObject136 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject136) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject136) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject136) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject136) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject136) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject136) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject136) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject136) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNextHopIp

`func (o *InlineObject136) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineObject136) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineObject136) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *InlineObject136) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetAdvertiseViaOspfEnabled

`func (o *InlineObject136) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *InlineObject136) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *InlineObject136) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *InlineObject136) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *InlineObject136) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *InlineObject136) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *InlineObject136) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *InlineObject136) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

