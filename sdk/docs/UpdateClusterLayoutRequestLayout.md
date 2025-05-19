# UpdateClusterLayoutRequestLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster layout name | [optional] 
**Description** | Pointer to **string** | Cluster layout description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**ComputeVersion** | Pointer to **string** | Version of the cluster layout | [optional] 
**Creatable** | Pointer to **bool** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to true]
**HasAutoScale** | Pointer to **bool** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**InstallContainerRuntime** | Pointer to **bool** | Install Docker (container runtime). | [optional] [default to false]
**MemoryRequirement** | Pointer to **int64** | Memory requirement in bytes | [optional] 
**GroupType** | Pointer to [**AddClusterLayoutsRequestLayoutGroupType**](AddClusterLayoutsRequestLayoutGroupType.md) |  | [optional] 
**ProvisionType** | Pointer to [**AddClusterLayoutsRequestLayoutProvisionType**](AddClusterLayoutsRequestLayoutProvisionType.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of cluster layout option types | [optional] 
**TaskSets** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of cluster layout task sets | [optional] 
**EnvironmentVariables** | Pointer to [**[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner**](AddClusterLayoutsRequestLayoutEnvironmentVariablesInner.md) | Array of cluster layout env variables | [optional] 
**Masters** | Pointer to [**[]AddClusterLayoutsRequestLayoutMastersInner**](AddClusterLayoutsRequestLayoutMastersInner.md) | Array of cluster layout master nodes | [optional] 
**Workers** | Pointer to [**[]AddClusterLayoutsRequestLayoutMastersInner**](AddClusterLayoutsRequestLayoutMastersInner.md) | Array of cluster layout worker nodes | [optional] 

## Methods

### NewUpdateClusterLayoutRequestLayout

`func NewUpdateClusterLayoutRequestLayout() *UpdateClusterLayoutRequestLayout`

NewUpdateClusterLayoutRequestLayout instantiates a new UpdateClusterLayoutRequestLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterLayoutRequestLayoutWithDefaults

`func NewUpdateClusterLayoutRequestLayoutWithDefaults() *UpdateClusterLayoutRequestLayout`

NewUpdateClusterLayoutRequestLayoutWithDefaults instantiates a new UpdateClusterLayoutRequestLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterLayoutRequestLayout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterLayoutRequestLayout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterLayoutRequestLayout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterLayoutRequestLayout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateClusterLayoutRequestLayout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClusterLayoutRequestLayout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClusterLayoutRequestLayout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClusterLayoutRequestLayout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateClusterLayoutRequestLayout) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateClusterLayoutRequestLayout) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateClusterLayoutRequestLayout) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateClusterLayoutRequestLayout) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetComputeVersion

`func (o *UpdateClusterLayoutRequestLayout) GetComputeVersion() string`

GetComputeVersion returns the ComputeVersion field if non-nil, zero value otherwise.

### GetComputeVersionOk

`func (o *UpdateClusterLayoutRequestLayout) GetComputeVersionOk() (*string, bool)`

GetComputeVersionOk returns a tuple with the ComputeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVersion

`func (o *UpdateClusterLayoutRequestLayout) SetComputeVersion(v string)`

SetComputeVersion sets ComputeVersion field to given value.

### HasComputeVersion

`func (o *UpdateClusterLayoutRequestLayout) HasComputeVersion() bool`

HasComputeVersion returns a boolean if a field has been set.

### GetCreatable

`func (o *UpdateClusterLayoutRequestLayout) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *UpdateClusterLayoutRequestLayout) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *UpdateClusterLayoutRequestLayout) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *UpdateClusterLayoutRequestLayout) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *UpdateClusterLayoutRequestLayout) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *UpdateClusterLayoutRequestLayout) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *UpdateClusterLayoutRequestLayout) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *UpdateClusterLayoutRequestLayout) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *UpdateClusterLayoutRequestLayout) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *UpdateClusterLayoutRequestLayout) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *UpdateClusterLayoutRequestLayout) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *UpdateClusterLayoutRequestLayout) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *UpdateClusterLayoutRequestLayout) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *UpdateClusterLayoutRequestLayout) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *UpdateClusterLayoutRequestLayout) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *UpdateClusterLayoutRequestLayout) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetGroupType

`func (o *UpdateClusterLayoutRequestLayout) GetGroupType() AddClusterLayoutsRequestLayoutGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *UpdateClusterLayoutRequestLayout) GetGroupTypeOk() (*AddClusterLayoutsRequestLayoutGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *UpdateClusterLayoutRequestLayout) SetGroupType(v AddClusterLayoutsRequestLayoutGroupType)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *UpdateClusterLayoutRequestLayout) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetProvisionType

`func (o *UpdateClusterLayoutRequestLayout) GetProvisionType() AddClusterLayoutsRequestLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *UpdateClusterLayoutRequestLayout) GetProvisionTypeOk() (*AddClusterLayoutsRequestLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *UpdateClusterLayoutRequestLayout) SetProvisionType(v AddClusterLayoutsRequestLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *UpdateClusterLayoutRequestLayout) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *UpdateClusterLayoutRequestLayout) GetOptionTypes() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *UpdateClusterLayoutRequestLayout) GetOptionTypesOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *UpdateClusterLayoutRequestLayout) SetOptionTypes(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *UpdateClusterLayoutRequestLayout) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTaskSets

`func (o *UpdateClusterLayoutRequestLayout) GetTaskSets() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *UpdateClusterLayoutRequestLayout) GetTaskSetsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *UpdateClusterLayoutRequestLayout) SetTaskSets(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *UpdateClusterLayoutRequestLayout) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *UpdateClusterLayoutRequestLayout) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpdateClusterLayoutRequestLayout) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpdateClusterLayoutRequestLayout) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpdateClusterLayoutRequestLayout) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetMasters

`func (o *UpdateClusterLayoutRequestLayout) GetMasters() []AddClusterLayoutsRequestLayoutMastersInner`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *UpdateClusterLayoutRequestLayout) GetMastersOk() (*[]AddClusterLayoutsRequestLayoutMastersInner, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *UpdateClusterLayoutRequestLayout) SetMasters(v []AddClusterLayoutsRequestLayoutMastersInner)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *UpdateClusterLayoutRequestLayout) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetWorkers

`func (o *UpdateClusterLayoutRequestLayout) GetWorkers() []AddClusterLayoutsRequestLayoutMastersInner`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *UpdateClusterLayoutRequestLayout) GetWorkersOk() (*[]AddClusterLayoutsRequestLayoutMastersInner, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *UpdateClusterLayoutRequestLayout) SetWorkers(v []AddClusterLayoutsRequestLayoutMastersInner)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *UpdateClusterLayoutRequestLayout) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


