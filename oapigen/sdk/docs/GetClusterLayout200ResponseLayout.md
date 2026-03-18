# GetClusterLayout200ResponseLayout

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
**GroupType** | Pointer to [**GetClusterLayout200ResponseLayoutGroupType**](GetClusterLayout200ResponseLayoutGroupType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to [**[]GetClusterLayout200ResponseLayoutOptionTypesInner**](GetClusterLayout200ResponseLayoutOptionTypesInner.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ComputeServers** | Pointer to [**[]GetClusterLayout200ResponseLayoutComputeServersInner**](GetClusterLayout200ResponseLayoutComputeServersInner.md) |  | [optional] 
**InstallContainerRuntime** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**GetClusterLayout200ResponseLayoutProvisionType**](GetClusterLayout200ResponseLayoutProvisionType.md) |  | [optional] 
**SpecTemplates** | Pointer to [**[]GetClusterLayout200ResponseLayoutSpecTemplatesInner**](GetClusterLayout200ResponseLayoutSpecTemplatesInner.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**GetClusterLayout200ResponseLayoutType**](GetClusterLayout200ResponseLayoutType.md) |  | [optional] 

## Methods

### NewGetClusterLayout200ResponseLayout

`func NewGetClusterLayout200ResponseLayout() *GetClusterLayout200ResponseLayout`

NewGetClusterLayout200ResponseLayout instantiates a new GetClusterLayout200ResponseLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterLayout200ResponseLayoutWithDefaults

`func NewGetClusterLayout200ResponseLayoutWithDefaults() *GetClusterLayout200ResponseLayout`

NewGetClusterLayout200ResponseLayoutWithDefaults instantiates a new GetClusterLayout200ResponseLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterLayout200ResponseLayout) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterLayout200ResponseLayout) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterLayout200ResponseLayout) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterLayout200ResponseLayout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *GetClusterLayout200ResponseLayout) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetClusterLayout200ResponseLayout) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetClusterLayout200ResponseLayout) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetClusterLayout200ResponseLayout) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetServerCount

`func (o *GetClusterLayout200ResponseLayout) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *GetClusterLayout200ResponseLayout) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *GetClusterLayout200ResponseLayout) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *GetClusterLayout200ResponseLayout) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetClusterLayout200ResponseLayout) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterLayout200ResponseLayout) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterLayout200ResponseLayout) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterLayout200ResponseLayout) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCode

`func (o *GetClusterLayout200ResponseLayout) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterLayout200ResponseLayout) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterLayout200ResponseLayout) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterLayout200ResponseLayout) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterLayout200ResponseLayout) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterLayout200ResponseLayout) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterLayout200ResponseLayout) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterLayout200ResponseLayout) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *GetClusterLayout200ResponseLayout) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *GetClusterLayout200ResponseLayout) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *GetClusterLayout200ResponseLayout) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *GetClusterLayout200ResponseLayout) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *GetClusterLayout200ResponseLayout) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *GetClusterLayout200ResponseLayout) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *GetClusterLayout200ResponseLayout) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *GetClusterLayout200ResponseLayout) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetClusterVersion

`func (o *GetClusterLayout200ResponseLayout) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *GetClusterLayout200ResponseLayout) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *GetClusterLayout200ResponseLayout) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *GetClusterLayout200ResponseLayout) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.

### GetComputeVersion

`func (o *GetClusterLayout200ResponseLayout) GetComputeVersion() string`

GetComputeVersion returns the ComputeVersion field if non-nil, zero value otherwise.

### GetComputeVersionOk

`func (o *GetClusterLayout200ResponseLayout) GetComputeVersionOk() (*string, bool)`

GetComputeVersionOk returns a tuple with the ComputeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVersion

`func (o *GetClusterLayout200ResponseLayout) SetComputeVersion(v string)`

SetComputeVersion sets ComputeVersion field to given value.

### HasComputeVersion

`func (o *GetClusterLayout200ResponseLayout) HasComputeVersion() bool`

HasComputeVersion returns a boolean if a field has been set.

### GetHasSettings

`func (o *GetClusterLayout200ResponseLayout) GetHasSettings() bool`

GetHasSettings returns the HasSettings field if non-nil, zero value otherwise.

### GetHasSettingsOk

`func (o *GetClusterLayout200ResponseLayout) GetHasSettingsOk() (*bool, bool)`

GetHasSettingsOk returns a tuple with the HasSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSettings

`func (o *GetClusterLayout200ResponseLayout) SetHasSettings(v bool)`

SetHasSettings sets HasSettings field to given value.

### HasHasSettings

`func (o *GetClusterLayout200ResponseLayout) HasHasSettings() bool`

HasHasSettings returns a boolean if a field has been set.

### GetSortOrder

`func (o *GetClusterLayout200ResponseLayout) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetClusterLayout200ResponseLayout) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetClusterLayout200ResponseLayout) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetClusterLayout200ResponseLayout) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetHasConfig

`func (o *GetClusterLayout200ResponseLayout) GetHasConfig() bool`

GetHasConfig returns the HasConfig field if non-nil, zero value otherwise.

### GetHasConfigOk

`func (o *GetClusterLayout200ResponseLayout) GetHasConfigOk() (*bool, bool)`

GetHasConfigOk returns a tuple with the HasConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfig

`func (o *GetClusterLayout200ResponseLayout) SetHasConfig(v bool)`

SetHasConfig sets HasConfig field to given value.

### HasHasConfig

`func (o *GetClusterLayout200ResponseLayout) HasHasConfig() bool`

HasHasConfig returns a boolean if a field has been set.

### GetName

`func (o *GetClusterLayout200ResponseLayout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterLayout200ResponseLayout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterLayout200ResponseLayout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterLayout200ResponseLayout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatable

`func (o *GetClusterLayout200ResponseLayout) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetClusterLayout200ResponseLayout) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetClusterLayout200ResponseLayout) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetClusterLayout200ResponseLayout) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetEnabled

`func (o *GetClusterLayout200ResponseLayout) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetClusterLayout200ResponseLayout) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetClusterLayout200ResponseLayout) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetClusterLayout200ResponseLayout) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterLayout200ResponseLayout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterLayout200ResponseLayout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterLayout200ResponseLayout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterLayout200ResponseLayout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupType

`func (o *GetClusterLayout200ResponseLayout) GetGroupType() GetClusterLayout200ResponseLayoutGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *GetClusterLayout200ResponseLayout) GetGroupTypeOk() (*GetClusterLayout200ResponseLayoutGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *GetClusterLayout200ResponseLayout) SetGroupType(v GetClusterLayout200ResponseLayoutGroupType)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *GetClusterLayout200ResponseLayout) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetLabels

`func (o *GetClusterLayout200ResponseLayout) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetClusterLayout200ResponseLayout) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetClusterLayout200ResponseLayout) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetClusterLayout200ResponseLayout) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *GetClusterLayout200ResponseLayout) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetClusterLayout200ResponseLayout) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetClusterLayout200ResponseLayout) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetClusterLayout200ResponseLayout) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetClusterLayout200ResponseLayout) GetOptionTypes() []GetClusterLayout200ResponseLayoutOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetClusterLayout200ResponseLayout) GetOptionTypesOk() (*[]GetClusterLayout200ResponseLayoutOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetClusterLayout200ResponseLayout) SetOptionTypes(v []GetClusterLayout200ResponseLayoutOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetClusterLayout200ResponseLayout) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetActions

`func (o *GetClusterLayout200ResponseLayout) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetClusterLayout200ResponseLayout) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetClusterLayout200ResponseLayout) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetClusterLayout200ResponseLayout) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetComputeServers

`func (o *GetClusterLayout200ResponseLayout) GetComputeServers() []GetClusterLayout200ResponseLayoutComputeServersInner`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *GetClusterLayout200ResponseLayout) GetComputeServersOk() (*[]GetClusterLayout200ResponseLayoutComputeServersInner, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *GetClusterLayout200ResponseLayout) SetComputeServers(v []GetClusterLayout200ResponseLayoutComputeServersInner)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *GetClusterLayout200ResponseLayout) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *GetClusterLayout200ResponseLayout) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *GetClusterLayout200ResponseLayout) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *GetClusterLayout200ResponseLayout) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *GetClusterLayout200ResponseLayout) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetClusterLayout200ResponseLayout) GetProvisionType() GetClusterLayout200ResponseLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetClusterLayout200ResponseLayout) GetProvisionTypeOk() (*GetClusterLayout200ResponseLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetClusterLayout200ResponseLayout) SetProvisionType(v GetClusterLayout200ResponseLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetClusterLayout200ResponseLayout) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *GetClusterLayout200ResponseLayout) GetSpecTemplates() []GetClusterLayout200ResponseLayoutSpecTemplatesInner`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *GetClusterLayout200ResponseLayout) GetSpecTemplatesOk() (*[]GetClusterLayout200ResponseLayoutSpecTemplatesInner, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *GetClusterLayout200ResponseLayout) SetSpecTemplates(v []GetClusterLayout200ResponseLayoutSpecTemplatesInner)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *GetClusterLayout200ResponseLayout) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### GetTaskSets

`func (o *GetClusterLayout200ResponseLayout) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *GetClusterLayout200ResponseLayout) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *GetClusterLayout200ResponseLayout) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *GetClusterLayout200ResponseLayout) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetType

`func (o *GetClusterLayout200ResponseLayout) GetType() GetClusterLayout200ResponseLayoutType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterLayout200ResponseLayout) GetTypeOk() (*GetClusterLayout200ResponseLayoutType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterLayout200ResponseLayout) SetType(v GetClusterLayout200ResponseLayoutType)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterLayout200ResponseLayout) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


