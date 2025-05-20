# ListApps200ResponseAllOfAppsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Blueprint** | Pointer to [**ListApps200ResponseAllOfAppsInnerBlueprint**](ListApps200ResponseAllOfAppsInnerBlueprint.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RemovalDate** | Pointer to **time.Time** |  | [optional] 
**AppContext** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AppStatus** | Pointer to **string** |  | [optional] 
**InstanceCount** | Pointer to **int64** |  | [optional] 
**ContainerCount** | Pointer to **int64** |  | [optional] 
**AppTiers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Instances** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Stats** | Pointer to [**ListApps200ResponseAllOfAppsInnerStats**](ListApps200ResponseAllOfAppsInnerStats.md) |  | [optional] 

## Methods

### NewListApps200ResponseAllOfAppsInner

`func NewListApps200ResponseAllOfAppsInner() *ListApps200ResponseAllOfAppsInner`

NewListApps200ResponseAllOfAppsInner instantiates a new ListApps200ResponseAllOfAppsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApps200ResponseAllOfAppsInnerWithDefaults

`func NewListApps200ResponseAllOfAppsInnerWithDefaults() *ListApps200ResponseAllOfAppsInner`

NewListApps200ResponseAllOfAppsInnerWithDefaults instantiates a new ListApps200ResponseAllOfAppsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListApps200ResponseAllOfAppsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApps200ResponseAllOfAppsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApps200ResponseAllOfAppsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListApps200ResponseAllOfAppsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListApps200ResponseAllOfAppsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApps200ResponseAllOfAppsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApps200ResponseAllOfAppsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListApps200ResponseAllOfAppsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListApps200ResponseAllOfAppsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListApps200ResponseAllOfAppsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListApps200ResponseAllOfAppsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListApps200ResponseAllOfAppsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ListApps200ResponseAllOfAppsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListApps200ResponseAllOfAppsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListApps200ResponseAllOfAppsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListApps200ResponseAllOfAppsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironment

`func (o *ListApps200ResponseAllOfAppsInner) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ListApps200ResponseAllOfAppsInner) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ListApps200ResponseAllOfAppsInner) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ListApps200ResponseAllOfAppsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAccountId

`func (o *ListApps200ResponseAllOfAppsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListApps200ResponseAllOfAppsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListApps200ResponseAllOfAppsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListApps200ResponseAllOfAppsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ListApps200ResponseAllOfAppsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListApps200ResponseAllOfAppsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListApps200ResponseAllOfAppsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListApps200ResponseAllOfAppsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListApps200ResponseAllOfAppsInner) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListApps200ResponseAllOfAppsInner) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListApps200ResponseAllOfAppsInner) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListApps200ResponseAllOfAppsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSiteId

`func (o *ListApps200ResponseAllOfAppsInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListApps200ResponseAllOfAppsInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListApps200ResponseAllOfAppsInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListApps200ResponseAllOfAppsInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetGroup

`func (o *ListApps200ResponseAllOfAppsInner) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListApps200ResponseAllOfAppsInner) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListApps200ResponseAllOfAppsInner) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListApps200ResponseAllOfAppsInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBlueprint

`func (o *ListApps200ResponseAllOfAppsInner) GetBlueprint() ListApps200ResponseAllOfAppsInnerBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *ListApps200ResponseAllOfAppsInner) GetBlueprintOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *ListApps200ResponseAllOfAppsInner) SetBlueprint(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *ListApps200ResponseAllOfAppsInner) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetType

`func (o *ListApps200ResponseAllOfAppsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListApps200ResponseAllOfAppsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListApps200ResponseAllOfAppsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListApps200ResponseAllOfAppsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListApps200ResponseAllOfAppsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListApps200ResponseAllOfAppsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListApps200ResponseAllOfAppsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListApps200ResponseAllOfAppsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListApps200ResponseAllOfAppsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListApps200ResponseAllOfAppsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListApps200ResponseAllOfAppsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListApps200ResponseAllOfAppsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRemovalDate

`func (o *ListApps200ResponseAllOfAppsInner) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *ListApps200ResponseAllOfAppsInner) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *ListApps200ResponseAllOfAppsInner) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *ListApps200ResponseAllOfAppsInner) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### GetAppContext

`func (o *ListApps200ResponseAllOfAppsInner) GetAppContext() string`

GetAppContext returns the AppContext field if non-nil, zero value otherwise.

### GetAppContextOk

`func (o *ListApps200ResponseAllOfAppsInner) GetAppContextOk() (*string, bool)`

GetAppContextOk returns a tuple with the AppContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContext

`func (o *ListApps200ResponseAllOfAppsInner) SetAppContext(v string)`

SetAppContext sets AppContext field to given value.

### HasAppContext

`func (o *ListApps200ResponseAllOfAppsInner) HasAppContext() bool`

HasAppContext returns a boolean if a field has been set.

### GetStatus

`func (o *ListApps200ResponseAllOfAppsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListApps200ResponseAllOfAppsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListApps200ResponseAllOfAppsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListApps200ResponseAllOfAppsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppStatus

`func (o *ListApps200ResponseAllOfAppsInner) GetAppStatus() string`

GetAppStatus returns the AppStatus field if non-nil, zero value otherwise.

### GetAppStatusOk

`func (o *ListApps200ResponseAllOfAppsInner) GetAppStatusOk() (*string, bool)`

GetAppStatusOk returns a tuple with the AppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatus

`func (o *ListApps200ResponseAllOfAppsInner) SetAppStatus(v string)`

SetAppStatus sets AppStatus field to given value.

### HasAppStatus

`func (o *ListApps200ResponseAllOfAppsInner) HasAppStatus() bool`

HasAppStatus returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ListApps200ResponseAllOfAppsInner) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ListApps200ResponseAllOfAppsInner) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ListApps200ResponseAllOfAppsInner) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ListApps200ResponseAllOfAppsInner) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetContainerCount

`func (o *ListApps200ResponseAllOfAppsInner) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *ListApps200ResponseAllOfAppsInner) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *ListApps200ResponseAllOfAppsInner) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *ListApps200ResponseAllOfAppsInner) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetAppTiers

`func (o *ListApps200ResponseAllOfAppsInner) GetAppTiers() []map[string]interface{}`

GetAppTiers returns the AppTiers field if non-nil, zero value otherwise.

### GetAppTiersOk

`func (o *ListApps200ResponseAllOfAppsInner) GetAppTiersOk() (*[]map[string]interface{}, bool)`

GetAppTiersOk returns a tuple with the AppTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTiers

`func (o *ListApps200ResponseAllOfAppsInner) SetAppTiers(v []map[string]interface{})`

SetAppTiers sets AppTiers field to given value.

### HasAppTiers

`func (o *ListApps200ResponseAllOfAppsInner) HasAppTiers() bool`

HasAppTiers returns a boolean if a field has been set.

### GetInstances

`func (o *ListApps200ResponseAllOfAppsInner) GetInstances() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListApps200ResponseAllOfAppsInner) GetInstancesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListApps200ResponseAllOfAppsInner) SetInstances(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListApps200ResponseAllOfAppsInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetStats

`func (o *ListApps200ResponseAllOfAppsInner) GetStats() ListApps200ResponseAllOfAppsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListApps200ResponseAllOfAppsInner) GetStatsOk() (*ListApps200ResponseAllOfAppsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListApps200ResponseAllOfAppsInner) SetStats(v ListApps200ResponseAllOfAppsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListApps200ResponseAllOfAppsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


