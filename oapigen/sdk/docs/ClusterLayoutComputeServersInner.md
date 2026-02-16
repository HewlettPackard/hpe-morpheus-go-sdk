# ClusterLayoutComputeServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PriorityOrder** | Pointer to **int64** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**MinNodeCount** | Pointer to **int64** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableString** |  | [optional] 
**DynamicCount** | Pointer to **bool** |  | [optional] 
**InstallContainerRuntime** | Pointer to **bool** |  | [optional] 
**InstallStorageRuntime** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **NullableString** |  | [optional] 
**ContainerType** | Pointer to [**ClusterLayoutComputeServersInnerContainerType**](ClusterLayoutComputeServersInnerContainerType.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ClusterLayoutComputeServersInnerComputeServerType**](ClusterLayoutComputeServersInnerComputeServerType.md) |  | [optional] 
**ProvisionService** | Pointer to **NullableString** |  | [optional] 
**PlanCategory** | Pointer to **NullableString** |  | [optional] 
**NamePrefix** | Pointer to **NullableString** |  | [optional] 
**NameSuffix** | Pointer to **NullableString** |  | [optional] 
**ForceNameIndex** | Pointer to **bool** |  | [optional] 
**LoadBalance** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterLayoutComputeServersInner

`func NewClusterLayoutComputeServersInner() *ClusterLayoutComputeServersInner`

NewClusterLayoutComputeServersInner instantiates a new ClusterLayoutComputeServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutComputeServersInnerWithDefaults

`func NewClusterLayoutComputeServersInnerWithDefaults() *ClusterLayoutComputeServersInner`

NewClusterLayoutComputeServersInnerWithDefaults instantiates a new ClusterLayoutComputeServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterLayoutComputeServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterLayoutComputeServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterLayoutComputeServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterLayoutComputeServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriorityOrder

`func (o *ClusterLayoutComputeServersInner) GetPriorityOrder() int64`

GetPriorityOrder returns the PriorityOrder field if non-nil, zero value otherwise.

### GetPriorityOrderOk

`func (o *ClusterLayoutComputeServersInner) GetPriorityOrderOk() (*int64, bool)`

GetPriorityOrderOk returns a tuple with the PriorityOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityOrder

`func (o *ClusterLayoutComputeServersInner) SetPriorityOrder(v int64)`

SetPriorityOrder sets PriorityOrder field to given value.

### HasPriorityOrder

`func (o *ClusterLayoutComputeServersInner) HasPriorityOrder() bool`

HasPriorityOrder returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterLayoutComputeServersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterLayoutComputeServersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterLayoutComputeServersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterLayoutComputeServersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeType

`func (o *ClusterLayoutComputeServersInner) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ClusterLayoutComputeServersInner) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ClusterLayoutComputeServersInner) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ClusterLayoutComputeServersInner) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetMinNodeCount

`func (o *ClusterLayoutComputeServersInner) GetMinNodeCount() int64`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *ClusterLayoutComputeServersInner) GetMinNodeCountOk() (*int64, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *ClusterLayoutComputeServersInner) SetMinNodeCount(v int64)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *ClusterLayoutComputeServersInner) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### GetMaxNodeCount

`func (o *ClusterLayoutComputeServersInner) GetMaxNodeCount() string`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *ClusterLayoutComputeServersInner) GetMaxNodeCountOk() (*string, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *ClusterLayoutComputeServersInner) SetMaxNodeCount(v string)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *ClusterLayoutComputeServersInner) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *ClusterLayoutComputeServersInner) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *ClusterLayoutComputeServersInner) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetDynamicCount

`func (o *ClusterLayoutComputeServersInner) GetDynamicCount() bool`

GetDynamicCount returns the DynamicCount field if non-nil, zero value otherwise.

### GetDynamicCountOk

`func (o *ClusterLayoutComputeServersInner) GetDynamicCountOk() (*bool, bool)`

GetDynamicCountOk returns a tuple with the DynamicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCount

`func (o *ClusterLayoutComputeServersInner) SetDynamicCount(v bool)`

SetDynamicCount sets DynamicCount field to given value.

### HasDynamicCount

`func (o *ClusterLayoutComputeServersInner) HasDynamicCount() bool`

HasDynamicCount returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *ClusterLayoutComputeServersInner) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *ClusterLayoutComputeServersInner) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *ClusterLayoutComputeServersInner) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *ClusterLayoutComputeServersInner) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetInstallStorageRuntime

`func (o *ClusterLayoutComputeServersInner) GetInstallStorageRuntime() bool`

GetInstallStorageRuntime returns the InstallStorageRuntime field if non-nil, zero value otherwise.

### GetInstallStorageRuntimeOk

`func (o *ClusterLayoutComputeServersInner) GetInstallStorageRuntimeOk() (*bool, bool)`

GetInstallStorageRuntimeOk returns a tuple with the InstallStorageRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStorageRuntime

`func (o *ClusterLayoutComputeServersInner) SetInstallStorageRuntime(v bool)`

SetInstallStorageRuntime sets InstallStorageRuntime field to given value.

### HasInstallStorageRuntime

`func (o *ClusterLayoutComputeServersInner) HasInstallStorageRuntime() bool`

HasInstallStorageRuntime returns a boolean if a field has been set.

### GetName

`func (o *ClusterLayoutComputeServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutComputeServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutComputeServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLayoutComputeServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ClusterLayoutComputeServersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterLayoutComputeServersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterLayoutComputeServersInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterLayoutComputeServersInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *ClusterLayoutComputeServersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterLayoutComputeServersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterLayoutComputeServersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterLayoutComputeServersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterLayoutComputeServersInner) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterLayoutComputeServersInner) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterLayoutComputeServersInner) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterLayoutComputeServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ClusterLayoutComputeServersInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ClusterLayoutComputeServersInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerType

`func (o *ClusterLayoutComputeServersInner) GetContainerType() ClusterLayoutComputeServersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ClusterLayoutComputeServersInner) GetContainerTypeOk() (*ClusterLayoutComputeServersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ClusterLayoutComputeServersInner) SetContainerType(v ClusterLayoutComputeServersInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ClusterLayoutComputeServersInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ClusterLayoutComputeServersInner) GetComputeServerType() ClusterLayoutComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ClusterLayoutComputeServersInner) GetComputeServerTypeOk() (*ClusterLayoutComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ClusterLayoutComputeServersInner) SetComputeServerType(v ClusterLayoutComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ClusterLayoutComputeServersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetProvisionService

`func (o *ClusterLayoutComputeServersInner) GetProvisionService() string`

GetProvisionService returns the ProvisionService field if non-nil, zero value otherwise.

### GetProvisionServiceOk

`func (o *ClusterLayoutComputeServersInner) GetProvisionServiceOk() (*string, bool)`

GetProvisionServiceOk returns a tuple with the ProvisionService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionService

`func (o *ClusterLayoutComputeServersInner) SetProvisionService(v string)`

SetProvisionService sets ProvisionService field to given value.

### HasProvisionService

`func (o *ClusterLayoutComputeServersInner) HasProvisionService() bool`

HasProvisionService returns a boolean if a field has been set.

### SetProvisionServiceNil

`func (o *ClusterLayoutComputeServersInner) SetProvisionServiceNil(b bool)`

 SetProvisionServiceNil sets the value for ProvisionService to be an explicit nil

### UnsetProvisionService
`func (o *ClusterLayoutComputeServersInner) UnsetProvisionService()`

UnsetProvisionService ensures that no value is present for ProvisionService, not even an explicit nil
### GetPlanCategory

`func (o *ClusterLayoutComputeServersInner) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *ClusterLayoutComputeServersInner) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *ClusterLayoutComputeServersInner) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *ClusterLayoutComputeServersInner) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### SetPlanCategoryNil

`func (o *ClusterLayoutComputeServersInner) SetPlanCategoryNil(b bool)`

 SetPlanCategoryNil sets the value for PlanCategory to be an explicit nil

### UnsetPlanCategory
`func (o *ClusterLayoutComputeServersInner) UnsetPlanCategory()`

UnsetPlanCategory ensures that no value is present for PlanCategory, not even an explicit nil
### GetNamePrefix

`func (o *ClusterLayoutComputeServersInner) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *ClusterLayoutComputeServersInner) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *ClusterLayoutComputeServersInner) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *ClusterLayoutComputeServersInner) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### SetNamePrefixNil

`func (o *ClusterLayoutComputeServersInner) SetNamePrefixNil(b bool)`

 SetNamePrefixNil sets the value for NamePrefix to be an explicit nil

### UnsetNamePrefix
`func (o *ClusterLayoutComputeServersInner) UnsetNamePrefix()`

UnsetNamePrefix ensures that no value is present for NamePrefix, not even an explicit nil
### GetNameSuffix

`func (o *ClusterLayoutComputeServersInner) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *ClusterLayoutComputeServersInner) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *ClusterLayoutComputeServersInner) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *ClusterLayoutComputeServersInner) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### SetNameSuffixNil

`func (o *ClusterLayoutComputeServersInner) SetNameSuffixNil(b bool)`

 SetNameSuffixNil sets the value for NameSuffix to be an explicit nil

### UnsetNameSuffix
`func (o *ClusterLayoutComputeServersInner) UnsetNameSuffix()`

UnsetNameSuffix ensures that no value is present for NameSuffix, not even an explicit nil
### GetForceNameIndex

`func (o *ClusterLayoutComputeServersInner) GetForceNameIndex() bool`

GetForceNameIndex returns the ForceNameIndex field if non-nil, zero value otherwise.

### GetForceNameIndexOk

`func (o *ClusterLayoutComputeServersInner) GetForceNameIndexOk() (*bool, bool)`

GetForceNameIndexOk returns a tuple with the ForceNameIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNameIndex

`func (o *ClusterLayoutComputeServersInner) SetForceNameIndex(v bool)`

SetForceNameIndex sets ForceNameIndex field to given value.

### HasForceNameIndex

`func (o *ClusterLayoutComputeServersInner) HasForceNameIndex() bool`

HasForceNameIndex returns a boolean if a field has been set.

### GetLoadBalance

`func (o *ClusterLayoutComputeServersInner) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *ClusterLayoutComputeServersInner) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *ClusterLayoutComputeServersInner) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *ClusterLayoutComputeServersInner) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


