# AddSecurityGroups200ResponseSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**GroupSource** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **string** |  | [optional] 
**SyncSource** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Locations** | Pointer to [**[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner**](ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner.md) |  | [optional] 
**Rules** | Pointer to [**[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner**](ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner**](ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions**](ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddSecurityGroups200ResponseSecurityGroup

`func NewAddSecurityGroups200ResponseSecurityGroup() *AddSecurityGroups200ResponseSecurityGroup`

NewAddSecurityGroups200ResponseSecurityGroup instantiates a new AddSecurityGroups200ResponseSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroups200ResponseSecurityGroupWithDefaults

`func NewAddSecurityGroups200ResponseSecurityGroupWithDefaults() *AddSecurityGroups200ResponseSecurityGroup`

NewAddSecurityGroups200ResponseSecurityGroupWithDefaults instantiates a new AddSecurityGroups200ResponseSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountId

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGroupSource

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetGroupSource() string`

GetGroupSource returns the GroupSource field if non-nil, zero value otherwise.

### GetGroupSourceOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetGroupSourceOk() (*string, bool)`

GetGroupSourceOk returns a tuple with the GroupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSource

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetGroupSource(v string)`

SetGroupSource sets GroupSource field to given value.

### HasGroupSource

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasGroupSource() bool`

HasGroupSource returns a boolean if a field has been set.

### GetExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSyncSource

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetVisibility

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetZone

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetLocations

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetLocations() []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetLocationsOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetLocations(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRules

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetRules() []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetRulesOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetRules(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTenants

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetTenants() []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetTenantsOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetTenants(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetResourcePermission() ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetResourcePermissionOk() (*ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetResourcePermission(v ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetSuccess

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddSecurityGroups200ResponseSecurityGroup) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddSecurityGroups200ResponseSecurityGroup) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddSecurityGroups200ResponseSecurityGroup) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


