# ListOptionNetworkOptions200Response

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
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewListOptionNetworkOptions200Response

`func NewListOptionNetworkOptions200Response() *ListOptionNetworkOptions200Response`

NewListOptionNetworkOptions200Response instantiates a new ListOptionNetworkOptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionNetworkOptions200ResponseWithDefaults

`func NewListOptionNetworkOptions200ResponseWithDefaults() *ListOptionNetworkOptions200Response`

NewListOptionNetworkOptions200ResponseWithDefaults instantiates a new ListOptionNetworkOptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *ListOptionNetworkOptions200Response) GetNetworks() []ListOptionNetworkOptions200ResponseAllOfNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ListOptionNetworkOptions200Response) GetNetworksOk() (*[]ListOptionNetworkOptions200ResponseAllOfNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ListOptionNetworkOptions200Response) SetNetworks(v []ListOptionNetworkOptions200ResponseAllOfNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ListOptionNetworkOptions200Response) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNetworkGroups

`func (o *ListOptionNetworkOptions200Response) GetNetworkGroups() []map[string]interface{}`

GetNetworkGroups returns the NetworkGroups field if non-nil, zero value otherwise.

### GetNetworkGroupsOk

`func (o *ListOptionNetworkOptions200Response) GetNetworkGroupsOk() (*[]map[string]interface{}, bool)`

GetNetworkGroupsOk returns a tuple with the NetworkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroups

`func (o *ListOptionNetworkOptions200Response) SetNetworkGroups(v []map[string]interface{})`

SetNetworkGroups sets NetworkGroups field to given value.

### HasNetworkGroups

`func (o *ListOptionNetworkOptions200Response) HasNetworkGroups() bool`

HasNetworkGroups returns a boolean if a field has been set.

### GetNetworkTypes

`func (o *ListOptionNetworkOptions200Response) GetNetworkTypes() []ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *ListOptionNetworkOptions200Response) GetNetworkTypesOk() (*[]ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *ListOptionNetworkOptions200Response) SetNetworkTypes(v []ListOptionNetworkOptions200ResponseAllOfNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *ListOptionNetworkOptions200Response) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### GetNetworkSubnets

`func (o *ListOptionNetworkOptions200Response) GetNetworkSubnets() []map[string]interface{}`

GetNetworkSubnets returns the NetworkSubnets field if non-nil, zero value otherwise.

### GetNetworkSubnetsOk

`func (o *ListOptionNetworkOptions200Response) GetNetworkSubnetsOk() (*[]map[string]interface{}, bool)`

GetNetworkSubnetsOk returns a tuple with the NetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSubnets

`func (o *ListOptionNetworkOptions200Response) SetNetworkSubnets(v []map[string]interface{})`

SetNetworkSubnets sets NetworkSubnets field to given value.

### HasNetworkSubnets

`func (o *ListOptionNetworkOptions200Response) HasNetworkSubnets() bool`

HasNetworkSubnets returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ListOptionNetworkOptions200Response) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ListOptionNetworkOptions200Response) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ListOptionNetworkOptions200Response) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ListOptionNetworkOptions200Response) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *ListOptionNetworkOptions200Response) GetMaxNetworks() int64`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *ListOptionNetworkOptions200Response) GetMaxNetworksOk() (*int64, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *ListOptionNetworkOptions200Response) SetMaxNetworks(v int64)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *ListOptionNetworkOptions200Response) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ListOptionNetworkOptions200Response) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ListOptionNetworkOptions200Response) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ListOptionNetworkOptions200Response) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ListOptionNetworkOptions200Response) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetSupportsNetworkSelection

`func (o *ListOptionNetworkOptions200Response) GetSupportsNetworkSelection() bool`

GetSupportsNetworkSelection returns the SupportsNetworkSelection field if non-nil, zero value otherwise.

### GetSupportsNetworkSelectionOk

`func (o *ListOptionNetworkOptions200Response) GetSupportsNetworkSelectionOk() (*bool, bool)`

GetSupportsNetworkSelectionOk returns a tuple with the SupportsNetworkSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNetworkSelection

`func (o *ListOptionNetworkOptions200Response) SetSupportsNetworkSelection(v bool)`

SetSupportsNetworkSelection sets SupportsNetworkSelection field to given value.

### HasSupportsNetworkSelection

`func (o *ListOptionNetworkOptions200Response) HasSupportsNetworkSelection() bool`

HasSupportsNetworkSelection returns a boolean if a field has been set.

### GetSuccess

`func (o *ListOptionNetworkOptions200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListOptionNetworkOptions200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListOptionNetworkOptions200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListOptionNetworkOptions200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


