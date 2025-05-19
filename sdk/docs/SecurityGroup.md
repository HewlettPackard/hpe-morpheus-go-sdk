# SecurityGroup

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

## Methods

### NewSecurityGroup

`func NewSecurityGroup() *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountId

`func (o *SecurityGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecurityGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SecurityGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SecurityGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGroupSource

`func (o *SecurityGroup) GetGroupSource() string`

GetGroupSource returns the GroupSource field if non-nil, zero value otherwise.

### GetGroupSourceOk

`func (o *SecurityGroup) GetGroupSourceOk() (*string, bool)`

GetGroupSourceOk returns a tuple with the GroupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSource

`func (o *SecurityGroup) SetGroupSource(v string)`

SetGroupSource sets GroupSource field to given value.

### HasGroupSource

`func (o *SecurityGroup) HasGroupSource() bool`

HasGroupSource returns a boolean if a field has been set.

### GetExternalId

`func (o *SecurityGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SecurityGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SecurityGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SecurityGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *SecurityGroup) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityGroup) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityGroup) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSyncSource

`func (o *SecurityGroup) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *SecurityGroup) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *SecurityGroup) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *SecurityGroup) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetVisibility

`func (o *SecurityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SecurityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SecurityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SecurityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *SecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetZone

`func (o *SecurityGroup) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SecurityGroup) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SecurityGroup) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SecurityGroup) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetLocations

`func (o *SecurityGroup) GetLocations() []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SecurityGroup) GetLocationsOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SecurityGroup) SetLocations(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SecurityGroup) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRules

`func (o *SecurityGroup) GetRules() []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroup) GetRulesOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroup) SetRules(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTenants

`func (o *SecurityGroup) GetTenants() []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SecurityGroup) GetTenantsOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SecurityGroup) SetTenants(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SecurityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *SecurityGroup) GetResourcePermission() ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *SecurityGroup) GetResourcePermissionOk() (*ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *SecurityGroup) SetResourcePermission(v ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *SecurityGroup) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


