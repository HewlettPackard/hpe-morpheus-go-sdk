# ZoneNetworkOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]ListOptionNetworkOptions200ResponseAllOfNetworksInner**](ListOptionNetworkOptions200ResponseAllOfNetworksInner.md) |  | [optional] 
**NetworkGroups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NetworkTypes** | Pointer to [**[]ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner**](ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner.md) |  | [optional] 
**NetworkSubnets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**MaxNetworks** | Pointer to **int64** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **string** |  | [optional] 
**SupportsNetworkSelection** | Pointer to **bool** |  | [optional] 

## Methods

### NewZoneNetworkOptions

`func NewZoneNetworkOptions() *ZoneNetworkOptions`

NewZoneNetworkOptions instantiates a new ZoneNetworkOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneNetworkOptionsWithDefaults

`func NewZoneNetworkOptionsWithDefaults() *ZoneNetworkOptions`

NewZoneNetworkOptionsWithDefaults instantiates a new ZoneNetworkOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *ZoneNetworkOptions) GetNetworks() []ListOptionNetworkOptions200ResponseAllOfNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ZoneNetworkOptions) GetNetworksOk() (*[]ListOptionNetworkOptions200ResponseAllOfNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ZoneNetworkOptions) SetNetworks(v []ListOptionNetworkOptions200ResponseAllOfNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ZoneNetworkOptions) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNetworkGroups

`func (o *ZoneNetworkOptions) GetNetworkGroups() []map[string]interface{}`

GetNetworkGroups returns the NetworkGroups field if non-nil, zero value otherwise.

### GetNetworkGroupsOk

`func (o *ZoneNetworkOptions) GetNetworkGroupsOk() (*[]map[string]interface{}, bool)`

GetNetworkGroupsOk returns a tuple with the NetworkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroups

`func (o *ZoneNetworkOptions) SetNetworkGroups(v []map[string]interface{})`

SetNetworkGroups sets NetworkGroups field to given value.

### HasNetworkGroups

`func (o *ZoneNetworkOptions) HasNetworkGroups() bool`

HasNetworkGroups returns a boolean if a field has been set.

### GetNetworkTypes

`func (o *ZoneNetworkOptions) GetNetworkTypes() []ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *ZoneNetworkOptions) GetNetworkTypesOk() (*[]ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *ZoneNetworkOptions) SetNetworkTypes(v []ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *ZoneNetworkOptions) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### GetNetworkSubnets

`func (o *ZoneNetworkOptions) GetNetworkSubnets() []map[string]interface{}`

GetNetworkSubnets returns the NetworkSubnets field if non-nil, zero value otherwise.

### GetNetworkSubnetsOk

`func (o *ZoneNetworkOptions) GetNetworkSubnetsOk() (*[]map[string]interface{}, bool)`

GetNetworkSubnetsOk returns a tuple with the NetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSubnets

`func (o *ZoneNetworkOptions) SetNetworkSubnets(v []map[string]interface{})`

SetNetworkSubnets sets NetworkSubnets field to given value.

### HasNetworkSubnets

`func (o *ZoneNetworkOptions) HasNetworkSubnets() bool`

HasNetworkSubnets returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ZoneNetworkOptions) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ZoneNetworkOptions) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ZoneNetworkOptions) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ZoneNetworkOptions) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *ZoneNetworkOptions) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *ZoneNetworkOptions) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *ZoneNetworkOptions) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *ZoneNetworkOptions) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ZoneNetworkOptions) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneNetworkOptions) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneNetworkOptions) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneNetworkOptions) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetSupportsNetworkSelection

`func (o *ZoneNetworkOptions) GetSupportsNetworkSelection() bool`

GetSupportsNetworkSelection returns the SupportsNetworkSelection field if non-nil, zero value otherwise.

### GetSupportsNetworkSelectionOk

`func (o *ZoneNetworkOptions) GetSupportsNetworkSelectionOk() (*bool, bool)`

GetSupportsNetworkSelectionOk returns a tuple with the SupportsNetworkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNetworkSelection

`func (o *ZoneNetworkOptions) SetSupportsNetworkSelection(v bool)`

SetSupportsNetworkSelection sets SupportsNetworkSelection field to given value.

### HasSupportsNetworkSelection

`func (o *ZoneNetworkOptions) HasSupportsNetworkSelection() bool`

HasSupportsNetworkSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


