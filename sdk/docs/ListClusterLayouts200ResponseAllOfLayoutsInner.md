# ListClusterLayouts200ResponseAllOfLayoutsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**HasAutoScale** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **int64** |  | [optional] 
**ClusterVersion** | Pointer to **string** |  | [optional] 
**ComputeVersion** | Pointer to **string** |  | [optional] 
**HasSettings** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**HasConfig** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GroupType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ComputeServers** | Pointer to [**[]ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner.md) |  | [optional] 
**InstallContainerRuntime** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**SpecTemplates** | Pointer to [**[]ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner**](ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 

## Methods

### NewListClusterLayouts200ResponseAllOfLayoutsInner

`func NewListClusterLayouts200ResponseAllOfLayoutsInner() *ListClusterLayouts200ResponseAllOfLayoutsInner`

NewListClusterLayouts200ResponseAllOfLayoutsInner instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLayouts200ResponseAllOfLayoutsInnerWithDefaults

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerWithDefaults() *ListClusterLayouts200ResponseAllOfLayoutsInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerWithDefaults instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetServerCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetClusterVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.

### GetComputeVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetComputeVersion() string`

GetComputeVersion returns the ComputeVersion field if non-nil, zero value otherwise.

### GetComputeVersionOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetComputeVersionOk() (*string, bool)`

GetComputeVersionOk returns a tuple with the ComputeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetComputeVersion(v string)`

SetComputeVersion sets ComputeVersion field to given value.

### HasComputeVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasComputeVersion() bool`

HasComputeVersion returns a boolean if a field has been set.

### GetHasSettings

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetHasSettings() bool`

GetHasSettings returns the HasSettings field if non-nil, zero value otherwise.

### GetHasSettingsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetHasSettingsOk() (*bool, bool)`

GetHasSettingsOk returns a tuple with the HasSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSettings

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetHasSettings(v bool)`

SetHasSettings sets HasSettings field to given value.

### HasHasSettings

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasHasSettings() bool`

HasHasSettings returns a boolean if a field has been set.

### GetSortOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetHasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetHasConfig() bool`

GetHasConfig returns the HasConfig field if non-nil, zero value otherwise.

### GetHasConfigOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetHasConfigOk() (*bool, bool)`

GetHasConfigOk returns a tuple with the HasConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetHasConfig(v bool)`

SetHasConfig sets HasConfig field to given value.

### HasHasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasHasConfig() bool`

HasHasConfig returns a boolean if a field has been set.

### GetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetEnabled

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetGroupType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetGroupTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetGroupType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetActions

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetComputeServers

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetComputeServers() []ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetComputeServersOk() (*[]ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetComputeServers(v []ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetProvisionType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetProvisionTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetProvisionType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetSpecTemplates() []ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetSpecTemplatesOk() (*[]ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetSpecTemplates(v []ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### GetTaskSets

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


