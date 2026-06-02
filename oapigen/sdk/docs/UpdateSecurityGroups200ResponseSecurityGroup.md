# UpdateSecurityGroups200ResponseSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**GroupSource** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 
**SyncSource** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to [**UpdateSecurityGroups200ResponseSecurityGroupAllOfZone**](UpdateSecurityGroups200ResponseSecurityGroupAllOfZone.md) |  | [optional] 
**Locations** | Pointer to [**[]UpdateSecurityGroups200ResponseSecurityGroupAllOfLocationsInner**](UpdateSecurityGroups200ResponseSecurityGroupAllOfLocationsInner.md) |  | [optional] 
**Rules** | Pointer to [**[]UpdateSecurityGroups200ResponseSecurityGroupAllOfRulesInner**](UpdateSecurityGroups200ResponseSecurityGroupAllOfRulesInner.md) |  | [optional] 
**Tenants** | Pointer to [**[]UpdateSecurityGroups200ResponseSecurityGroupAllOfTenantsInner**](UpdateSecurityGroups200ResponseSecurityGroupAllOfTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**UpdateSecurityGroups200ResponseSecurityGroupAllOfResourcePermission**](UpdateSecurityGroups200ResponseSecurityGroupAllOfResourcePermission.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateSecurityGroups200ResponseSecurityGroup

`func NewUpdateSecurityGroups200ResponseSecurityGroup() *UpdateSecurityGroups200ResponseSecurityGroup`

NewUpdateSecurityGroups200ResponseSecurityGroup instantiates a new UpdateSecurityGroups200ResponseSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateSecurityGroups200ResponseSecurityGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccountId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGroupSource

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetGroupSource() string`

GetGroupSource returns the GroupSource field if non-nil, zero value otherwise.

### GetGroupSourceOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetGroupSourceOk() (*string, bool)`

GetGroupSourceOk returns a tuple with the GroupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSource

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetGroupSource(v string)`

SetGroupSource sets GroupSource field to given value.

### HasGroupSource

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasGroupSource() bool`

HasGroupSource returns a boolean if a field has been set.

### SetGroupSourceNil

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetGroupSourceNil(b bool)`

 SetGroupSourceNil sets the value for GroupSource to be an explicit nil

### UnsetGroupSource
`func (o *UpdateSecurityGroups200ResponseSecurityGroup) UnsetGroupSource()`

UnsetGroupSource ensures that no value is present for GroupSource, not even an explicit nil
### GetExternalId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateSecurityGroups200ResponseSecurityGroup) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateSecurityGroups200ResponseSecurityGroup) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetSyncSource

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetZone

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetZone() UpdateSecurityGroups200ResponseSecurityGroupAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetZoneOk() (*UpdateSecurityGroups200ResponseSecurityGroupAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetZone(v UpdateSecurityGroups200ResponseSecurityGroupAllOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetLocations

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetLocations() []UpdateSecurityGroups200ResponseSecurityGroupAllOfLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetLocationsOk() (*[]UpdateSecurityGroups200ResponseSecurityGroupAllOfLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetLocations(v []UpdateSecurityGroups200ResponseSecurityGroupAllOfLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRules

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetRules() []UpdateSecurityGroups200ResponseSecurityGroupAllOfRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetRulesOk() (*[]UpdateSecurityGroups200ResponseSecurityGroupAllOfRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetRules(v []UpdateSecurityGroups200ResponseSecurityGroupAllOfRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetTenants() []UpdateSecurityGroups200ResponseSecurityGroupAllOfTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetTenantsOk() (*[]UpdateSecurityGroups200ResponseSecurityGroupAllOfTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetTenants(v []UpdateSecurityGroups200ResponseSecurityGroupAllOfTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetResourcePermission() UpdateSecurityGroups200ResponseSecurityGroupAllOfResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetResourcePermissionOk() (*UpdateSecurityGroups200ResponseSecurityGroupAllOfResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetResourcePermission(v UpdateSecurityGroups200ResponseSecurityGroupAllOfResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateSecurityGroups200ResponseSecurityGroup) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


