# GetApp200ResponseApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Blueprint** | Pointer to [**ListApps200ResponseAllOfAppsInnerBlueprint**](ListApps200ResponseAllOfAppsInnerBlueprint.md) |  | [optional] 
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
**Instances** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Stats** | Pointer to [**ListApps200ResponseAllOfAppsInnerStats**](ListApps200ResponseAllOfAppsInnerStats.md) |  | [optional] 

## Methods

### NewGetApp200ResponseApp

`func NewGetApp200ResponseApp() *GetApp200ResponseApp`

NewGetApp200ResponseApp instantiates a new GetApp200ResponseApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApp200ResponseAppWithDefaults

`func NewGetApp200ResponseAppWithDefaults() *GetApp200ResponseApp`

NewGetApp200ResponseAppWithDefaults instantiates a new GetApp200ResponseApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApp200ResponseApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApp200ResponseApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApp200ResponseApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetApp200ResponseApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetApp200ResponseApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApp200ResponseApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApp200ResponseApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetApp200ResponseApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetApp200ResponseApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetApp200ResponseApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetApp200ResponseApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetApp200ResponseApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *GetApp200ResponseApp) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetApp200ResponseApp) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetApp200ResponseApp) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetApp200ResponseApp) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironment

`func (o *GetApp200ResponseApp) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetApp200ResponseApp) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetApp200ResponseApp) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetApp200ResponseApp) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAccountId

`func (o *GetApp200ResponseApp) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetApp200ResponseApp) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetApp200ResponseApp) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetApp200ResponseApp) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *GetApp200ResponseApp) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetApp200ResponseApp) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetApp200ResponseApp) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetApp200ResponseApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *GetApp200ResponseApp) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetApp200ResponseApp) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetApp200ResponseApp) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetApp200ResponseApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSiteId

`func (o *GetApp200ResponseApp) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetApp200ResponseApp) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetApp200ResponseApp) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetApp200ResponseApp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetGroup

`func (o *GetApp200ResponseApp) GetGroup() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetApp200ResponseApp) GetGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetApp200ResponseApp) SetGroup(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetApp200ResponseApp) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBlueprint

`func (o *GetApp200ResponseApp) GetBlueprint() ListApps200ResponseAllOfAppsInnerBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *GetApp200ResponseApp) GetBlueprintOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *GetApp200ResponseApp) SetBlueprint(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *GetApp200ResponseApp) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetType

`func (o *GetApp200ResponseApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetApp200ResponseApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetApp200ResponseApp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetApp200ResponseApp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetApp200ResponseApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetApp200ResponseApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetApp200ResponseApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetApp200ResponseApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetApp200ResponseApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetApp200ResponseApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetApp200ResponseApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetApp200ResponseApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRemovalDate

`func (o *GetApp200ResponseApp) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *GetApp200ResponseApp) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *GetApp200ResponseApp) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *GetApp200ResponseApp) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### SetRemovalDateNil

`func (o *GetApp200ResponseApp) SetRemovalDateNil(b bool)`

 SetRemovalDateNil sets the value for RemovalDate to be an explicit nil

### UnsetRemovalDate
`func (o *GetApp200ResponseApp) UnsetRemovalDate()`

UnsetRemovalDate ensures that no value is present for RemovalDate, not even an explicit nil
### GetAppContext

`func (o *GetApp200ResponseApp) GetAppContext() string`

GetAppContext returns the AppContext field if non-nil, zero value otherwise.

### GetAppContextOk

`func (o *GetApp200ResponseApp) GetAppContextOk() (*string, bool)`

GetAppContextOk returns a tuple with the AppContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContext

`func (o *GetApp200ResponseApp) SetAppContext(v string)`

SetAppContext sets AppContext field to given value.

### HasAppContext

`func (o *GetApp200ResponseApp) HasAppContext() bool`

HasAppContext returns a boolean if a field has been set.

### GetStatus

`func (o *GetApp200ResponseApp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApp200ResponseApp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApp200ResponseApp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApp200ResponseApp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppStatus

`func (o *GetApp200ResponseApp) GetAppStatus() string`

GetAppStatus returns the AppStatus field if non-nil, zero value otherwise.

### GetAppStatusOk

`func (o *GetApp200ResponseApp) GetAppStatusOk() (*string, bool)`

GetAppStatusOk returns a tuple with the AppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatus

`func (o *GetApp200ResponseApp) SetAppStatus(v string)`

SetAppStatus sets AppStatus field to given value.

### HasAppStatus

`func (o *GetApp200ResponseApp) HasAppStatus() bool`

HasAppStatus returns a boolean if a field has been set.

### GetInstanceCount

`func (o *GetApp200ResponseApp) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *GetApp200ResponseApp) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *GetApp200ResponseApp) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *GetApp200ResponseApp) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetContainerCount

`func (o *GetApp200ResponseApp) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *GetApp200ResponseApp) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *GetApp200ResponseApp) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *GetApp200ResponseApp) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetAppTiers

`func (o *GetApp200ResponseApp) GetAppTiers() []map[string]interface{}`

GetAppTiers returns the AppTiers field if non-nil, zero value otherwise.

### GetAppTiersOk

`func (o *GetApp200ResponseApp) GetAppTiersOk() (*[]map[string]interface{}, bool)`

GetAppTiersOk returns a tuple with the AppTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTiers

`func (o *GetApp200ResponseApp) SetAppTiers(v []map[string]interface{})`

SetAppTiers sets AppTiers field to given value.

### HasAppTiers

`func (o *GetApp200ResponseApp) HasAppTiers() bool`

HasAppTiers returns a boolean if a field has been set.

### GetInstances

`func (o *GetApp200ResponseApp) GetInstances() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetApp200ResponseApp) GetInstancesOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetApp200ResponseApp) SetInstances(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetApp200ResponseApp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetStats

`func (o *GetApp200ResponseApp) GetStats() ListApps200ResponseAllOfAppsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetApp200ResponseApp) GetStatsOk() (*ListApps200ResponseAllOfAppsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetApp200ResponseApp) SetStats(v ListApps200ResponseAllOfAppsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetApp200ResponseApp) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


