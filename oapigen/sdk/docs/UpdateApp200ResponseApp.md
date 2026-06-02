# UpdateApp200ResponseApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**UpdateApp200ResponseAppAccount**](UpdateApp200ResponseAppAccount.md) |  | [optional] 
**Owner** | Pointer to [**UpdateApp200ResponseAppOwner**](UpdateApp200ResponseAppOwner.md) |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**UpdateApp200ResponseAppGroup**](UpdateApp200ResponseAppGroup.md) |  | [optional] 
**Blueprint** | Pointer to [**UpdateApp200ResponseAppBlueprint**](UpdateApp200ResponseAppBlueprint.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RemovalDate** | Pointer to **NullableTime** |  | [optional] 
**AppContext** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AppStatus** | Pointer to **string** |  | [optional] 
**InstanceCount** | Pointer to **int64** |  | [optional] 
**ContainerCount** | Pointer to **int64** |  | [optional] 
**AppTiers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Instances** | Pointer to [**[]UpdateApp200ResponseAppInstancesInner**](UpdateApp200ResponseAppInstancesInner.md) |  | [optional] 
**Stats** | Pointer to [**UpdateApp200ResponseAppStats**](UpdateApp200ResponseAppStats.md) |  | [optional] 

## Methods

### NewUpdateApp200ResponseApp

`func NewUpdateApp200ResponseApp() *UpdateApp200ResponseApp`

NewUpdateApp200ResponseApp instantiates a new UpdateApp200ResponseApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateApp200ResponseApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateApp200ResponseApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateApp200ResponseApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateApp200ResponseApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateApp200ResponseApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApp200ResponseApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApp200ResponseApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateApp200ResponseApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateApp200ResponseApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApp200ResponseApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApp200ResponseApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApp200ResponseApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateApp200ResponseApp) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateApp200ResponseApp) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateApp200ResponseApp) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateApp200ResponseApp) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironment

`func (o *UpdateApp200ResponseApp) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateApp200ResponseApp) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateApp200ResponseApp) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateApp200ResponseApp) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateApp200ResponseApp) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateApp200ResponseApp) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateApp200ResponseApp) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateApp200ResponseApp) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateApp200ResponseApp) GetAccount() UpdateApp200ResponseAppAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateApp200ResponseApp) GetAccountOk() (*UpdateApp200ResponseAppAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateApp200ResponseApp) SetAccount(v UpdateApp200ResponseAppAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateApp200ResponseApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateApp200ResponseApp) GetOwner() UpdateApp200ResponseAppOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateApp200ResponseApp) GetOwnerOk() (*UpdateApp200ResponseAppOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateApp200ResponseApp) SetOwner(v UpdateApp200ResponseAppOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateApp200ResponseApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSiteId

`func (o *UpdateApp200ResponseApp) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateApp200ResponseApp) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateApp200ResponseApp) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateApp200ResponseApp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateApp200ResponseApp) GetGroup() UpdateApp200ResponseAppGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateApp200ResponseApp) GetGroupOk() (*UpdateApp200ResponseAppGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateApp200ResponseApp) SetGroup(v UpdateApp200ResponseAppGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateApp200ResponseApp) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBlueprint

`func (o *UpdateApp200ResponseApp) GetBlueprint() UpdateApp200ResponseAppBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *UpdateApp200ResponseApp) GetBlueprintOk() (*UpdateApp200ResponseAppBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *UpdateApp200ResponseApp) SetBlueprint(v UpdateApp200ResponseAppBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *UpdateApp200ResponseApp) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetType

`func (o *UpdateApp200ResponseApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateApp200ResponseApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateApp200ResponseApp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateApp200ResponseApp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateApp200ResponseApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateApp200ResponseApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateApp200ResponseApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateApp200ResponseApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateApp200ResponseApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateApp200ResponseApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateApp200ResponseApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateApp200ResponseApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRemovalDate

`func (o *UpdateApp200ResponseApp) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *UpdateApp200ResponseApp) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *UpdateApp200ResponseApp) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *UpdateApp200ResponseApp) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### SetRemovalDateNil

`func (o *UpdateApp200ResponseApp) SetRemovalDateNil(b bool)`

 SetRemovalDateNil sets the value for RemovalDate to be an explicit nil

### UnsetRemovalDate
`func (o *UpdateApp200ResponseApp) UnsetRemovalDate()`

UnsetRemovalDate ensures that no value is present for RemovalDate, not even an explicit nil
### GetAppContext

`func (o *UpdateApp200ResponseApp) GetAppContext() string`

GetAppContext returns the AppContext field if non-nil, zero value otherwise.

### GetAppContextOk

`func (o *UpdateApp200ResponseApp) GetAppContextOk() (*string, bool)`

GetAppContextOk returns a tuple with the AppContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContext

`func (o *UpdateApp200ResponseApp) SetAppContext(v string)`

SetAppContext sets AppContext field to given value.

### HasAppContext

`func (o *UpdateApp200ResponseApp) HasAppContext() bool`

HasAppContext returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateApp200ResponseApp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateApp200ResponseApp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateApp200ResponseApp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateApp200ResponseApp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppStatus

`func (o *UpdateApp200ResponseApp) GetAppStatus() string`

GetAppStatus returns the AppStatus field if non-nil, zero value otherwise.

### GetAppStatusOk

`func (o *UpdateApp200ResponseApp) GetAppStatusOk() (*string, bool)`

GetAppStatusOk returns a tuple with the AppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatus

`func (o *UpdateApp200ResponseApp) SetAppStatus(v string)`

SetAppStatus sets AppStatus field to given value.

### HasAppStatus

`func (o *UpdateApp200ResponseApp) HasAppStatus() bool`

HasAppStatus returns a boolean if a field has been set.

### GetInstanceCount

`func (o *UpdateApp200ResponseApp) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *UpdateApp200ResponseApp) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *UpdateApp200ResponseApp) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *UpdateApp200ResponseApp) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetContainerCount

`func (o *UpdateApp200ResponseApp) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *UpdateApp200ResponseApp) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *UpdateApp200ResponseApp) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *UpdateApp200ResponseApp) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetAppTiers

`func (o *UpdateApp200ResponseApp) GetAppTiers() []map[string]interface{}`

GetAppTiers returns the AppTiers field if non-nil, zero value otherwise.

### GetAppTiersOk

`func (o *UpdateApp200ResponseApp) GetAppTiersOk() (*[]map[string]interface{}, bool)`

GetAppTiersOk returns a tuple with the AppTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTiers

`func (o *UpdateApp200ResponseApp) SetAppTiers(v []map[string]interface{})`

SetAppTiers sets AppTiers field to given value.

### HasAppTiers

`func (o *UpdateApp200ResponseApp) HasAppTiers() bool`

HasAppTiers returns a boolean if a field has been set.

### GetInstances

`func (o *UpdateApp200ResponseApp) GetInstances() []UpdateApp200ResponseAppInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *UpdateApp200ResponseApp) GetInstancesOk() (*[]UpdateApp200ResponseAppInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *UpdateApp200ResponseApp) SetInstances(v []UpdateApp200ResponseAppInstancesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *UpdateApp200ResponseApp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetStats

`func (o *UpdateApp200ResponseApp) GetStats() UpdateApp200ResponseAppStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateApp200ResponseApp) GetStatsOk() (*UpdateApp200ResponseAppStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateApp200ResponseApp) SetStats(v UpdateApp200ResponseAppStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateApp200ResponseApp) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


