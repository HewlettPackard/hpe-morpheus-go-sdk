# GetClusterLayout200ResponseLayoutComputeServersInner

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
**ContainerType** | Pointer to [**GetClusterLayout200ResponseLayoutComputeServersInnerContainerType**](GetClusterLayout200ResponseLayoutComputeServersInnerContainerType.md) |  | [optional] 
**ComputeServerType** | Pointer to [**GetClusterLayout200ResponseLayoutComputeServersInnerComputeServerType**](GetClusterLayout200ResponseLayoutComputeServersInnerComputeServerType.md) |  | [optional] 
**ProvisionService** | Pointer to **NullableString** |  | [optional] 
**PlanCategory** | Pointer to **NullableString** |  | [optional] 
**NamePrefix** | Pointer to **NullableString** |  | [optional] 
**NameSuffix** | Pointer to **NullableString** |  | [optional] 
**ForceNameIndex** | Pointer to **bool** |  | [optional] 
**LoadBalance** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetClusterLayout200ResponseLayoutComputeServersInner

`func NewGetClusterLayout200ResponseLayoutComputeServersInner() *GetClusterLayout200ResponseLayoutComputeServersInner`

NewGetClusterLayout200ResponseLayoutComputeServersInner instantiates a new GetClusterLayout200ResponseLayoutComputeServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterLayout200ResponseLayoutComputeServersInnerWithDefaults

`func NewGetClusterLayout200ResponseLayoutComputeServersInnerWithDefaults() *GetClusterLayout200ResponseLayoutComputeServersInner`

NewGetClusterLayout200ResponseLayoutComputeServersInnerWithDefaults instantiates a new GetClusterLayout200ResponseLayoutComputeServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriorityOrder

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetPriorityOrder() int64`

GetPriorityOrder returns the PriorityOrder field if non-nil, zero value otherwise.

### GetPriorityOrderOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetPriorityOrderOk() (*int64, bool)`

GetPriorityOrderOk returns a tuple with the PriorityOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityOrder

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetPriorityOrder(v int64)`

SetPriorityOrder sets PriorityOrder field to given value.

### HasPriorityOrder

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasPriorityOrder() bool`

HasPriorityOrder returns a boolean if a field has been set.

### GetNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetMinNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetMinNodeCount() int64`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetMinNodeCountOk() (*int64, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetMinNodeCount(v int64)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### GetMaxNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetMaxNodeCount() string`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetMaxNodeCountOk() (*string, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetMaxNodeCount(v string)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetDynamicCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetDynamicCount() bool`

GetDynamicCount returns the DynamicCount field if non-nil, zero value otherwise.

### GetDynamicCountOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetDynamicCountOk() (*bool, bool)`

GetDynamicCountOk returns a tuple with the DynamicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetDynamicCount(v bool)`

SetDynamicCount sets DynamicCount field to given value.

### HasDynamicCount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasDynamicCount() bool`

HasDynamicCount returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetInstallStorageRuntime

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetInstallStorageRuntime() bool`

GetInstallStorageRuntime returns the InstallStorageRuntime field if non-nil, zero value otherwise.

### GetInstallStorageRuntimeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetInstallStorageRuntimeOk() (*bool, bool)`

GetInstallStorageRuntimeOk returns a tuple with the InstallStorageRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStorageRuntime

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetInstallStorageRuntime(v bool)`

SetInstallStorageRuntime sets InstallStorageRuntime field to given value.

### HasInstallStorageRuntime

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasInstallStorageRuntime() bool`

HasInstallStorageRuntime returns a boolean if a field has been set.

### GetName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetContainerType() GetClusterLayout200ResponseLayoutComputeServersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetContainerTypeOk() (*GetClusterLayout200ResponseLayoutComputeServersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetContainerType(v GetClusterLayout200ResponseLayoutComputeServersInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetComputeServerType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetComputeServerType() GetClusterLayout200ResponseLayoutComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetComputeServerTypeOk() (*GetClusterLayout200ResponseLayoutComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetComputeServerType(v GetClusterLayout200ResponseLayoutComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetProvisionService

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetProvisionService() string`

GetProvisionService returns the ProvisionService field if non-nil, zero value otherwise.

### GetProvisionServiceOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetProvisionServiceOk() (*string, bool)`

GetProvisionServiceOk returns a tuple with the ProvisionService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionService

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetProvisionService(v string)`

SetProvisionService sets ProvisionService field to given value.

### HasProvisionService

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasProvisionService() bool`

HasProvisionService returns a boolean if a field has been set.

### SetProvisionServiceNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetProvisionServiceNil(b bool)`

 SetProvisionServiceNil sets the value for ProvisionService to be an explicit nil

### UnsetProvisionService
`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) UnsetProvisionService()`

UnsetProvisionService ensures that no value is present for ProvisionService, not even an explicit nil
### GetPlanCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### SetPlanCategoryNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetPlanCategoryNil(b bool)`

 SetPlanCategoryNil sets the value for PlanCategory to be an explicit nil

### UnsetPlanCategory
`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) UnsetPlanCategory()`

UnsetPlanCategory ensures that no value is present for PlanCategory, not even an explicit nil
### GetNamePrefix

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### SetNamePrefixNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetNamePrefixNil(b bool)`

 SetNamePrefixNil sets the value for NamePrefix to be an explicit nil

### UnsetNamePrefix
`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) UnsetNamePrefix()`

UnsetNamePrefix ensures that no value is present for NamePrefix, not even an explicit nil
### GetNameSuffix

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### SetNameSuffixNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetNameSuffixNil(b bool)`

 SetNameSuffixNil sets the value for NameSuffix to be an explicit nil

### UnsetNameSuffix
`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) UnsetNameSuffix()`

UnsetNameSuffix ensures that no value is present for NameSuffix, not even an explicit nil
### GetForceNameIndex

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetForceNameIndex() bool`

GetForceNameIndex returns the ForceNameIndex field if non-nil, zero value otherwise.

### GetForceNameIndexOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetForceNameIndexOk() (*bool, bool)`

GetForceNameIndexOk returns a tuple with the ForceNameIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNameIndex

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetForceNameIndex(v bool)`

SetForceNameIndex sets ForceNameIndex field to given value.

### HasForceNameIndex

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasForceNameIndex() bool`

HasForceNameIndex returns a boolean if a field has been set.

### GetLoadBalance

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *GetClusterLayout200ResponseLayoutComputeServersInner) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


