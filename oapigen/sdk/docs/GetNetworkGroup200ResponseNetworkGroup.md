# GetNetworkGroup200ResponseNetworkGroup

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
**Tenants** | Pointer to [**[]GetNetworkGroup200ResponseNetworkGroupTenantsInner**](GetNetworkGroup200ResponseNetworkGroupTenantsInner.md) |  | [optional] 

## Methods

### NewGetNetworkGroup200ResponseNetworkGroup

`func NewGetNetworkGroup200ResponseNetworkGroup() *GetNetworkGroup200ResponseNetworkGroup`

NewGetNetworkGroup200ResponseNetworkGroup instantiates a new GetNetworkGroup200ResponseNetworkGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworks

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetNetworks() []int64`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetNetworksOk() (*[]int64, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetNetworks(v []int64)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetTenants() []GetNetworkGroup200ResponseNetworkGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetworkGroup200ResponseNetworkGroup) GetTenantsOk() (*[]GetNetworkGroup200ResponseNetworkGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetworkGroup200ResponseNetworkGroup) SetTenants(v []GetNetworkGroup200ResponseNetworkGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetworkGroup200ResponseNetworkGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


