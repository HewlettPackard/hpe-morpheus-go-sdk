# GetIntegrationInventory200ResponseInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewGetIntegrationInventory200ResponseInventory

`func NewGetIntegrationInventory200ResponseInventory() *GetIntegrationInventory200ResponseInventory`

NewGetIntegrationInventory200ResponseInventory instantiates a new GetIntegrationInventory200ResponseInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrationInventory200ResponseInventoryWithDefaults

`func NewGetIntegrationInventory200ResponseInventoryWithDefaults() *GetIntegrationInventory200ResponseInventory`

NewGetIntegrationInventory200ResponseInventoryWithDefaults instantiates a new GetIntegrationInventory200ResponseInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIntegrationInventory200ResponseInventory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrationInventory200ResponseInventory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrationInventory200ResponseInventory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrationInventory200ResponseInventory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrationInventory200ResponseInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrationInventory200ResponseInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrationInventory200ResponseInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrationInventory200ResponseInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetIntegrationInventory200ResponseInventory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIntegrationInventory200ResponseInventory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIntegrationInventory200ResponseInventory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetIntegrationInventory200ResponseInventory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetIntegrationInventory200ResponseInventory) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetIntegrationInventory200ResponseInventory) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *GetIntegrationInventory200ResponseInventory) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetIntegrationInventory200ResponseInventory) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetIntegrationInventory200ResponseInventory) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetIntegrationInventory200ResponseInventory) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetIntegrationInventory200ResponseInventory) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetIntegrationInventory200ResponseInventory) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetOwner

`func (o *GetIntegrationInventory200ResponseInventory) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetIntegrationInventory200ResponseInventory) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetIntegrationInventory200ResponseInventory) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetIntegrationInventory200ResponseInventory) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *GetIntegrationInventory200ResponseInventory) GetTenants() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetIntegrationInventory200ResponseInventory) GetTenantsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetIntegrationInventory200ResponseInventory) SetTenants(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetIntegrationInventory200ResponseInventory) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


