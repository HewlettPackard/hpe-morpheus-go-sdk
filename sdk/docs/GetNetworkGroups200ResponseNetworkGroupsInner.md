# GetNetworkGroups200ResponseNetworkGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Networks** | Pointer to **[]int64** |  | [optional] 
**Subnets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewGetNetworkGroups200ResponseNetworkGroupsInner

`func NewGetNetworkGroups200ResponseNetworkGroupsInner() *GetNetworkGroups200ResponseNetworkGroupsInner`

NewGetNetworkGroups200ResponseNetworkGroupsInner instantiates a new GetNetworkGroups200ResponseNetworkGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroups200ResponseNetworkGroupsInnerWithDefaults

`func NewGetNetworkGroups200ResponseNetworkGroupsInnerWithDefaults() *GetNetworkGroups200ResponseNetworkGroupsInner`

NewGetNetworkGroups200ResponseNetworkGroupsInnerWithDefaults instantiates a new GetNetworkGroups200ResponseNetworkGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworks

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetNetworks() []int64`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetNetworksOk() (*[]int64, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetNetworks(v []int64)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetworkGroups200ResponseNetworkGroupsInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


