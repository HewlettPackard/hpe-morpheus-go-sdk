# AddApps200ResponseApp

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

## Methods

### NewAddApps200ResponseApp

`func NewAddApps200ResponseApp() *AddApps200ResponseApp`

NewAddApps200ResponseApp instantiates a new AddApps200ResponseApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddApps200ResponseAppWithDefaults

`func NewAddApps200ResponseAppWithDefaults() *AddApps200ResponseApp`

NewAddApps200ResponseAppWithDefaults instantiates a new AddApps200ResponseApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddApps200ResponseApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddApps200ResponseApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddApps200ResponseApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddApps200ResponseApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddApps200ResponseApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddApps200ResponseApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddApps200ResponseApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddApps200ResponseApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddApps200ResponseApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddApps200ResponseApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddApps200ResponseApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddApps200ResponseApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddApps200ResponseApp) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddApps200ResponseApp) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddApps200ResponseApp) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddApps200ResponseApp) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironment

`func (o *AddApps200ResponseApp) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AddApps200ResponseApp) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AddApps200ResponseApp) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AddApps200ResponseApp) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAccountId

`func (o *AddApps200ResponseApp) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddApps200ResponseApp) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddApps200ResponseApp) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddApps200ResponseApp) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *AddApps200ResponseApp) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddApps200ResponseApp) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddApps200ResponseApp) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddApps200ResponseApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *AddApps200ResponseApp) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddApps200ResponseApp) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddApps200ResponseApp) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddApps200ResponseApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSiteId

`func (o *AddApps200ResponseApp) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AddApps200ResponseApp) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AddApps200ResponseApp) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AddApps200ResponseApp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetGroup

`func (o *AddApps200ResponseApp) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddApps200ResponseApp) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddApps200ResponseApp) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddApps200ResponseApp) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBlueprint

`func (o *AddApps200ResponseApp) GetBlueprint() ListApps200ResponseAllOfAppsInnerBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *AddApps200ResponseApp) GetBlueprintOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *AddApps200ResponseApp) SetBlueprint(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *AddApps200ResponseApp) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetType

`func (o *AddApps200ResponseApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddApps200ResponseApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddApps200ResponseApp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddApps200ResponseApp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddApps200ResponseApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddApps200ResponseApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddApps200ResponseApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddApps200ResponseApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddApps200ResponseApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddApps200ResponseApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddApps200ResponseApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddApps200ResponseApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRemovalDate

`func (o *AddApps200ResponseApp) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *AddApps200ResponseApp) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *AddApps200ResponseApp) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *AddApps200ResponseApp) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### GetAppContext

`func (o *AddApps200ResponseApp) GetAppContext() string`

GetAppContext returns the AppContext field if non-nil, zero value otherwise.

### GetAppContextOk

`func (o *AddApps200ResponseApp) GetAppContextOk() (*string, bool)`

GetAppContextOk returns a tuple with the AppContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContext

`func (o *AddApps200ResponseApp) SetAppContext(v string)`

SetAppContext sets AppContext field to given value.

### HasAppContext

`func (o *AddApps200ResponseApp) HasAppContext() bool`

HasAppContext returns a boolean if a field has been set.

### GetStatus

`func (o *AddApps200ResponseApp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddApps200ResponseApp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddApps200ResponseApp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddApps200ResponseApp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppStatus

`func (o *AddApps200ResponseApp) GetAppStatus() string`

GetAppStatus returns the AppStatus field if non-nil, zero value otherwise.

### GetAppStatusOk

`func (o *AddApps200ResponseApp) GetAppStatusOk() (*string, bool)`

GetAppStatusOk returns a tuple with the AppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatus

`func (o *AddApps200ResponseApp) SetAppStatus(v string)`

SetAppStatus sets AppStatus field to given value.

### HasAppStatus

`func (o *AddApps200ResponseApp) HasAppStatus() bool`

HasAppStatus returns a boolean if a field has been set.

### GetInstanceCount

`func (o *AddApps200ResponseApp) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AddApps200ResponseApp) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AddApps200ResponseApp) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *AddApps200ResponseApp) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetContainerCount

`func (o *AddApps200ResponseApp) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *AddApps200ResponseApp) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *AddApps200ResponseApp) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *AddApps200ResponseApp) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetAppTiers

`func (o *AddApps200ResponseApp) GetAppTiers() []map[string]interface{}`

GetAppTiers returns the AppTiers field if non-nil, zero value otherwise.

### GetAppTiersOk

`func (o *AddApps200ResponseApp) GetAppTiersOk() (*[]map[string]interface{}, bool)`

GetAppTiersOk returns a tuple with the AppTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTiers

`func (o *AddApps200ResponseApp) SetAppTiers(v []map[string]interface{})`

SetAppTiers sets AppTiers field to given value.

### HasAppTiers

`func (o *AddApps200ResponseApp) HasAppTiers() bool`

HasAppTiers returns a boolean if a field has been set.

### GetInstances

`func (o *AddApps200ResponseApp) GetInstances() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *AddApps200ResponseApp) GetInstancesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *AddApps200ResponseApp) SetInstances(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *AddApps200ResponseApp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


