# NetworkPoolServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PoolService** | Pointer to **string** |  | [optional] 
**Selectable** | Pointer to **bool** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**IntegrationCode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewNetworkPoolServerType

`func NewNetworkPoolServerType() *NetworkPoolServerType`

NewNetworkPoolServerType instantiates a new NetworkPoolServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerTypeWithDefaults

`func NewNetworkPoolServerTypeWithDefaults() *NetworkPoolServerType`

NewNetworkPoolServerTypeWithDefaults instantiates a new NetworkPoolServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkPoolServerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkPoolServerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkPoolServerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkPoolServerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPoolService

`func (o *NetworkPoolServerType) GetPoolService() string`

GetPoolService returns the PoolService field if non-nil, zero value otherwise.

### GetPoolServiceOk

`func (o *NetworkPoolServerType) GetPoolServiceOk() (*string, bool)`

GetPoolServiceOk returns a tuple with the PoolService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolService

`func (o *NetworkPoolServerType) SetPoolService(v string)`

SetPoolService sets PoolService field to given value.

### HasPoolService

`func (o *NetworkPoolServerType) HasPoolService() bool`

HasPoolService returns a boolean if a field has been set.

### GetSelectable

`func (o *NetworkPoolServerType) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *NetworkPoolServerType) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *NetworkPoolServerType) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *NetworkPoolServerType) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetCode

`func (o *NetworkPoolServerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkPoolServerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkPoolServerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkPoolServerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIntegrationCode

`func (o *NetworkPoolServerType) GetIntegrationCode() string`

GetIntegrationCode returns the IntegrationCode field if non-nil, zero value otherwise.

### GetIntegrationCodeOk

`func (o *NetworkPoolServerType) GetIntegrationCodeOk() (*string, bool)`

GetIntegrationCodeOk returns a tuple with the IntegrationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationCode

`func (o *NetworkPoolServerType) SetIntegrationCode(v string)`

SetIntegrationCode sets IntegrationCode field to given value.

### HasIntegrationCode

`func (o *NetworkPoolServerType) HasIntegrationCode() bool`

HasIntegrationCode returns a boolean if a field has been set.

### GetName

`func (o *NetworkPoolServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPoolServerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkPoolServerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkPoolServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkPoolServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkPoolServerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkPoolServerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptionTypes

`func (o *NetworkPoolServerType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *NetworkPoolServerType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *NetworkPoolServerType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *NetworkPoolServerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


