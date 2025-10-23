# ZoneNetworkOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]ZoneNetworkOptionsResponseNetworksInner**](ZoneNetworkOptionsResponseNetworksInner.md) |  | [optional] 
**NetworkGroups** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**NetworkTypes** | Pointer to [**[]ZoneNetworkOptionsResponseNetworkTypesInner**](ZoneNetworkOptionsResponseNetworkTypesInner.md) |  | [optional] 
**NetworkSubnets** | Pointer to [**[]ZoneNetworkOptionsResponseNetworkSubnetsInner**](ZoneNetworkOptionsResponseNetworkSubnetsInner.md) |  | [optional] 
**HasNetworks** | Pointer to **NullableBool** |  | [optional] 
**MaxNetworks** | Pointer to **NullableInt64** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **NullableString** |  | [optional] 
**SupportsNetworkSelection** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewZoneNetworkOptionsResponse

`func NewZoneNetworkOptionsResponse() *ZoneNetworkOptionsResponse`

NewZoneNetworkOptionsResponse instantiates a new ZoneNetworkOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneNetworkOptionsResponseWithDefaults

`func NewZoneNetworkOptionsResponseWithDefaults() *ZoneNetworkOptionsResponse`

NewZoneNetworkOptionsResponseWithDefaults instantiates a new ZoneNetworkOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *ZoneNetworkOptionsResponse) GetNetworks() []ZoneNetworkOptionsResponseNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ZoneNetworkOptionsResponse) GetNetworksOk() (*[]ZoneNetworkOptionsResponseNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ZoneNetworkOptionsResponse) SetNetworks(v []ZoneNetworkOptionsResponseNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ZoneNetworkOptionsResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNetworkGroups

`func (o *ZoneNetworkOptionsResponse) GetNetworkGroups() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetNetworkGroups returns the NetworkGroups field if non-nil, zero value otherwise.

### GetNetworkGroupsOk

`func (o *ZoneNetworkOptionsResponse) GetNetworkGroupsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetNetworkGroupsOk returns a tuple with the NetworkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroups

`func (o *ZoneNetworkOptionsResponse) SetNetworkGroups(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetNetworkGroups sets NetworkGroups field to given value.

### HasNetworkGroups

`func (o *ZoneNetworkOptionsResponse) HasNetworkGroups() bool`

HasNetworkGroups returns a boolean if a field has been set.

### GetNetworkTypes

`func (o *ZoneNetworkOptionsResponse) GetNetworkTypes() []ZoneNetworkOptionsResponseNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *ZoneNetworkOptionsResponse) GetNetworkTypesOk() (*[]ZoneNetworkOptionsResponseNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *ZoneNetworkOptionsResponse) SetNetworkTypes(v []ZoneNetworkOptionsResponseNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *ZoneNetworkOptionsResponse) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### GetNetworkSubnets

`func (o *ZoneNetworkOptionsResponse) GetNetworkSubnets() []ZoneNetworkOptionsResponseNetworkSubnetsInner`

GetNetworkSubnets returns the NetworkSubnets field if non-nil, zero value otherwise.

### GetNetworkSubnetsOk

`func (o *ZoneNetworkOptionsResponse) GetNetworkSubnetsOk() (*[]ZoneNetworkOptionsResponseNetworkSubnetsInner, bool)`

GetNetworkSubnetsOk returns a tuple with the NetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSubnets

`func (o *ZoneNetworkOptionsResponse) SetNetworkSubnets(v []ZoneNetworkOptionsResponseNetworkSubnetsInner)`

SetNetworkSubnets sets NetworkSubnets field to given value.

### HasNetworkSubnets

`func (o *ZoneNetworkOptionsResponse) HasNetworkSubnets() bool`

HasNetworkSubnets returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ZoneNetworkOptionsResponse) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ZoneNetworkOptionsResponse) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ZoneNetworkOptionsResponse) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ZoneNetworkOptionsResponse) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### SetHasNetworksNil

`func (o *ZoneNetworkOptionsResponse) SetHasNetworksNil(b bool)`

 SetHasNetworksNil sets the value for HasNetworks to be an explicit nil

### UnsetHasNetworks
`func (o *ZoneNetworkOptionsResponse) UnsetHasNetworks()`

UnsetHasNetworks ensures that no value is present for HasNetworks, not even an explicit nil
### GetMaxNetworks

`func (o *ZoneNetworkOptionsResponse) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *ZoneNetworkOptionsResponse) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *ZoneNetworkOptionsResponse) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *ZoneNetworkOptionsResponse) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### SetMaxNetworksNil

`func (o *ZoneNetworkOptionsResponse) SetMaxNetworksNil(b bool)`

 SetMaxNetworksNil sets the value for MaxNetworks to be an explicit nil

### UnsetMaxNetworks
`func (o *ZoneNetworkOptionsResponse) UnsetMaxNetworks()`

UnsetMaxNetworks ensures that no value is present for MaxNetworks, not even an explicit nil
### GetEnableNetworkTypeSelection

`func (o *ZoneNetworkOptionsResponse) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneNetworkOptionsResponse) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneNetworkOptionsResponse) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneNetworkOptionsResponse) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *ZoneNetworkOptionsResponse) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *ZoneNetworkOptionsResponse) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil
### GetSupportsNetworkSelection

`func (o *ZoneNetworkOptionsResponse) GetSupportsNetworkSelection() bool`

GetSupportsNetworkSelection returns the SupportsNetworkSelection field if non-nil, zero value otherwise.

### GetSupportsNetworkSelectionOk

`func (o *ZoneNetworkOptionsResponse) GetSupportsNetworkSelectionOk() (*bool, bool)`

GetSupportsNetworkSelectionOk returns a tuple with the SupportsNetworkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNetworkSelection

`func (o *ZoneNetworkOptionsResponse) SetSupportsNetworkSelection(v bool)`

SetSupportsNetworkSelection sets SupportsNetworkSelection field to given value.

### HasSupportsNetworkSelection

`func (o *ZoneNetworkOptionsResponse) HasSupportsNetworkSelection() bool`

HasSupportsNetworkSelection returns a boolean if a field has been set.

### SetSupportsNetworkSelectionNil

`func (o *ZoneNetworkOptionsResponse) SetSupportsNetworkSelectionNil(b bool)`

 SetSupportsNetworkSelectionNil sets the value for SupportsNetworkSelection to be an explicit nil

### UnsetSupportsNetworkSelection
`func (o *ZoneNetworkOptionsResponse) UnsetSupportsNetworkSelection()`

UnsetSupportsNetworkSelection ensures that no value is present for SupportsNetworkSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


