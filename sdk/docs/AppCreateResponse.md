# AppCreateResponse

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

### NewAppCreateResponse

`func NewAppCreateResponse() *AppCreateResponse`

NewAppCreateResponse instantiates a new AppCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCreateResponseWithDefaults

`func NewAppCreateResponseWithDefaults() *AppCreateResponse`

NewAppCreateResponseWithDefaults instantiates a new AppCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppCreateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppCreateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppCreateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AppCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AppCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AppCreateResponse) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AppCreateResponse) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AppCreateResponse) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AppCreateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironment

`func (o *AppCreateResponse) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AppCreateResponse) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AppCreateResponse) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AppCreateResponse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAccountId

`func (o *AppCreateResponse) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AppCreateResponse) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AppCreateResponse) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AppCreateResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *AppCreateResponse) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AppCreateResponse) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AppCreateResponse) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AppCreateResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *AppCreateResponse) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AppCreateResponse) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AppCreateResponse) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AppCreateResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSiteId

`func (o *AppCreateResponse) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AppCreateResponse) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AppCreateResponse) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AppCreateResponse) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetGroup

`func (o *AppCreateResponse) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AppCreateResponse) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AppCreateResponse) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AppCreateResponse) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetBlueprint

`func (o *AppCreateResponse) GetBlueprint() ListApps200ResponseAllOfAppsInnerBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *AppCreateResponse) GetBlueprintOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *AppCreateResponse) SetBlueprint(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *AppCreateResponse) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetType

`func (o *AppCreateResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppCreateResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppCreateResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppCreateResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreated

`func (o *AppCreateResponse) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AppCreateResponse) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AppCreateResponse) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AppCreateResponse) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AppCreateResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AppCreateResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AppCreateResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AppCreateResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRemovalDate

`func (o *AppCreateResponse) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *AppCreateResponse) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *AppCreateResponse) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *AppCreateResponse) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### GetAppContext

`func (o *AppCreateResponse) GetAppContext() string`

GetAppContext returns the AppContext field if non-nil, zero value otherwise.

### GetAppContextOk

`func (o *AppCreateResponse) GetAppContextOk() (*string, bool)`

GetAppContextOk returns a tuple with the AppContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContext

`func (o *AppCreateResponse) SetAppContext(v string)`

SetAppContext sets AppContext field to given value.

### HasAppContext

`func (o *AppCreateResponse) HasAppContext() bool`

HasAppContext returns a boolean if a field has been set.

### GetStatus

`func (o *AppCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppCreateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppStatus

`func (o *AppCreateResponse) GetAppStatus() string`

GetAppStatus returns the AppStatus field if non-nil, zero value otherwise.

### GetAppStatusOk

`func (o *AppCreateResponse) GetAppStatusOk() (*string, bool)`

GetAppStatusOk returns a tuple with the AppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatus

`func (o *AppCreateResponse) SetAppStatus(v string)`

SetAppStatus sets AppStatus field to given value.

### HasAppStatus

`func (o *AppCreateResponse) HasAppStatus() bool`

HasAppStatus returns a boolean if a field has been set.

### GetInstanceCount

`func (o *AppCreateResponse) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AppCreateResponse) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AppCreateResponse) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *AppCreateResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetContainerCount

`func (o *AppCreateResponse) GetContainerCount() int64`

GetContainerCount returns the ContainerCount field if non-nil, zero value otherwise.

### GetContainerCountOk

`func (o *AppCreateResponse) GetContainerCountOk() (*int64, bool)`

GetContainerCountOk returns a tuple with the ContainerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCount

`func (o *AppCreateResponse) SetContainerCount(v int64)`

SetContainerCount sets ContainerCount field to given value.

### HasContainerCount

`func (o *AppCreateResponse) HasContainerCount() bool`

HasContainerCount returns a boolean if a field has been set.

### GetAppTiers

`func (o *AppCreateResponse) GetAppTiers() []map[string]interface{}`

GetAppTiers returns the AppTiers field if non-nil, zero value otherwise.

### GetAppTiersOk

`func (o *AppCreateResponse) GetAppTiersOk() (*[]map[string]interface{}, bool)`

GetAppTiersOk returns a tuple with the AppTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTiers

`func (o *AppCreateResponse) SetAppTiers(v []map[string]interface{})`

SetAppTiers sets AppTiers field to given value.

### HasAppTiers

`func (o *AppCreateResponse) HasAppTiers() bool`

HasAppTiers returns a boolean if a field has been set.

### GetInstances

`func (o *AppCreateResponse) GetInstances() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *AppCreateResponse) GetInstancesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *AppCreateResponse) SetInstances(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *AppCreateResponse) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


