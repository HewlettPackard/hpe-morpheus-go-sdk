# UpdateIntegrationInventoryRequestInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant accounts that will use this inventory as Default. Used by jobs set to &#39;Use Tenant Default&#39; | [optional] 

## Methods

### NewUpdateIntegrationInventoryRequestInventory

`func NewUpdateIntegrationInventoryRequestInventory() *UpdateIntegrationInventoryRequestInventory`

NewUpdateIntegrationInventoryRequestInventory instantiates a new UpdateIntegrationInventoryRequestInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationInventoryRequestInventoryWithDefaults

`func NewUpdateIntegrationInventoryRequestInventoryWithDefaults() *UpdateIntegrationInventoryRequestInventory`

NewUpdateIntegrationInventoryRequestInventoryWithDefaults instantiates a new UpdateIntegrationInventoryRequestInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *UpdateIntegrationInventoryRequestInventory) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateIntegrationInventoryRequestInventory) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateIntegrationInventoryRequestInventory) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateIntegrationInventoryRequestInventory) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


