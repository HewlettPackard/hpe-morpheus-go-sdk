# IntegrationInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewIntegrationInventory

`func NewIntegrationInventory() *IntegrationInventory`

NewIntegrationInventory instantiates a new IntegrationInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationInventoryWithDefaults

`func NewIntegrationInventoryWithDefaults() *IntegrationInventory`

NewIntegrationInventoryWithDefaults instantiates a new IntegrationInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationInventory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationInventory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationInventory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationInventory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationInventory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationInventory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationInventory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationInventory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *IntegrationInventory) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IntegrationInventory) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IntegrationInventory) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IntegrationInventory) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetOwner

`func (o *IntegrationInventory) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IntegrationInventory) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IntegrationInventory) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IntegrationInventory) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *IntegrationInventory) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *IntegrationInventory) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *IntegrationInventory) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *IntegrationInventory) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


