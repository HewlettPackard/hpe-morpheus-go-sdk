# PolicyCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**CreateBackupType** | Pointer to **string** |  | [optional] 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**BackupStorageIds** | Pointer to **[]int64** |  | [optional] 
**MaxPrice** | Pointer to **float32** |  | [optional] 
**MaxPriceCurrency** | Pointer to **string** |  | [optional] 
**MaxPriceUnit** | Pointer to **string** |  | [optional] 
**ServerNamingType** | Pointer to **string** |  | [optional] 
**ServerNamingPattern** | Pointer to **string** |  | [optional] 
**ServerNamingConflict** | Pointer to **bool** |  | [optional] 
**KeyPattern** | Pointer to **string** |  | [optional] 
**Read** | Pointer to **bool** |  | [optional] 
**Write** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **bool** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**List** | Pointer to **bool** |  | [optional] 
**RemovalAge** | Pointer to **string** |  | [optional] 
**LifecycleType** | Pointer to **string** |  | [optional] 
**LifecycleAge** | Pointer to **string** |  | [optional] 
**LifecycleRenewal** | Pointer to **string** |  | [optional] 
**LifecycleNotify** | Pointer to **string** |  | [optional] 
**LifecycleMessage** | Pointer to **string** |  | [optional] 
**LifecycleAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**LifecycleHideFixed** | Pointer to **bool** |  | [optional] 
**LifecycleWorkflowId** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **string** |  | [optional] 
**HostNamingType** | Pointer to **string** |  | [optional] 
**HostNamingPattern** | Pointer to **string** |  | [optional] 
**NamingType** | Pointer to **string** |  | [optional] 
**NamingPattern** | Pointer to **string** |  | [optional] 
**NamingConflict** | Pointer to **bool** |  | [optional] 
**MaxContainers** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **string** |  | [optional] 
**ExcludeContainers** | Pointer to **string** |  | [optional] [default to "off"]
**MaxHosts** | Pointer to **string** |  | [optional] 
**MaxPools** | Pointer to **string** |  | [optional] 
**MaxMemory** | [**MaxMemoryPolicyTypeConfiguration7MaxMemory**](MaxMemoryPolicyTypeConfiguration7MaxMemory.md) |  | 
**MaxPoolMembers** | Pointer to **string** |  | [optional] 
**MaxSnapshots** | Pointer to **string** |  | [optional] 
**MaxVirtualServers** | Pointer to **string** |  | [optional] 
**MaxVms** | Pointer to **string** |  | [optional] 
**MotdTitle** | Pointer to **string** |  | [optional] 
**Motd** | Pointer to [**MessageOfTheDayPolicyTypeConfiguration7Motd**](MessageOfTheDayPolicyTypeConfiguration7Motd.md) |  | [optional] 
**MotdMessage** | Pointer to **string** |  | [optional] 
**MotdType** | Pointer to **string** |  | [optional] 
**MotdFullPage** | Pointer to **bool** |  | [optional] 
**MotdDate** | Pointer to **string** |  | [optional] 
**MaxNetworks** | Pointer to **string** |  | [optional] 
**PowerScheduleType** | Pointer to **string** |  | [optional] 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**PowerScheduleHideFixed** | Pointer to **bool** |  | [optional] 
**MaxRouters** | Pointer to **string** |  | [optional] 
**RequiredNetworks** | Pointer to **[]int64** | Array of network IDs that are required | [optional] 
**ShutdownType** | Pointer to **string** |  | [optional] 
**ShutdownAge** | Pointer to **string** |  | [optional] 
**ShutdownRenewal** | Pointer to **string** |  | [optional] 
**ShutdownNotify** | Pointer to **string** |  | [optional] 
**ShutdownMessage** | Pointer to **string** |  | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**ShutdownHideFixed** | Pointer to **bool** |  | [optional] 
**ShutdownWorkflowId** | Pointer to **string** |  | [optional] 
**StorageServerId** | Pointer to **string** |  | [optional] 
**Strict** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**ValueListId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CreateUserType** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyCreateConfig

`func NewPolicyCreateConfig(maxMemory MaxMemoryPolicyTypeConfiguration7MaxMemory, ) *PolicyCreateConfig`

NewPolicyCreateConfig instantiates a new PolicyCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCreateConfigWithDefaults

`func NewPolicyCreateConfigWithDefaults() *PolicyCreateConfig`

NewPolicyCreateConfigWithDefaults instantiates a new PolicyCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *PolicyCreateConfig) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *PolicyCreateConfig) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *PolicyCreateConfig) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *PolicyCreateConfig) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *PolicyCreateConfig) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *PolicyCreateConfig) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *PolicyCreateConfig) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *PolicyCreateConfig) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetFlowId

`func (o *PolicyCreateConfig) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *PolicyCreateConfig) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *PolicyCreateConfig) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *PolicyCreateConfig) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *PolicyCreateConfig) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *PolicyCreateConfig) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *PolicyCreateConfig) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *PolicyCreateConfig) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetCreateBackupType

`func (o *PolicyCreateConfig) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *PolicyCreateConfig) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *PolicyCreateConfig) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.

### HasCreateBackupType

`func (o *PolicyCreateConfig) HasCreateBackupType() bool`

HasCreateBackupType returns a boolean if a field has been set.

### GetCreateBackup

`func (o *PolicyCreateConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *PolicyCreateConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *PolicyCreateConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *PolicyCreateConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackupStorageIds

`func (o *PolicyCreateConfig) GetBackupStorageIds() []int64`

GetBackupStorageIds returns the BackupStorageIds field if non-nil, zero value otherwise.

### GetBackupStorageIdsOk

`func (o *PolicyCreateConfig) GetBackupStorageIdsOk() (*[]int64, bool)`

GetBackupStorageIdsOk returns a tuple with the BackupStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageIds

`func (o *PolicyCreateConfig) SetBackupStorageIds(v []int64)`

SetBackupStorageIds sets BackupStorageIds field to given value.

### HasBackupStorageIds

`func (o *PolicyCreateConfig) HasBackupStorageIds() bool`

HasBackupStorageIds returns a boolean if a field has been set.

### GetMaxPrice

`func (o *PolicyCreateConfig) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *PolicyCreateConfig) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *PolicyCreateConfig) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *PolicyCreateConfig) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetMaxPriceCurrency

`func (o *PolicyCreateConfig) GetMaxPriceCurrency() string`

GetMaxPriceCurrency returns the MaxPriceCurrency field if non-nil, zero value otherwise.

### GetMaxPriceCurrencyOk

`func (o *PolicyCreateConfig) GetMaxPriceCurrencyOk() (*string, bool)`

GetMaxPriceCurrencyOk returns a tuple with the MaxPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceCurrency

`func (o *PolicyCreateConfig) SetMaxPriceCurrency(v string)`

SetMaxPriceCurrency sets MaxPriceCurrency field to given value.

### HasMaxPriceCurrency

`func (o *PolicyCreateConfig) HasMaxPriceCurrency() bool`

HasMaxPriceCurrency returns a boolean if a field has been set.

### GetMaxPriceUnit

`func (o *PolicyCreateConfig) GetMaxPriceUnit() string`

GetMaxPriceUnit returns the MaxPriceUnit field if non-nil, zero value otherwise.

### GetMaxPriceUnitOk

`func (o *PolicyCreateConfig) GetMaxPriceUnitOk() (*string, bool)`

GetMaxPriceUnitOk returns a tuple with the MaxPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceUnit

`func (o *PolicyCreateConfig) SetMaxPriceUnit(v string)`

SetMaxPriceUnit sets MaxPriceUnit field to given value.

### HasMaxPriceUnit

`func (o *PolicyCreateConfig) HasMaxPriceUnit() bool`

HasMaxPriceUnit returns a boolean if a field has been set.

### GetServerNamingType

`func (o *PolicyCreateConfig) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *PolicyCreateConfig) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *PolicyCreateConfig) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.

### HasServerNamingType

`func (o *PolicyCreateConfig) HasServerNamingType() bool`

HasServerNamingType returns a boolean if a field has been set.

### GetServerNamingPattern

`func (o *PolicyCreateConfig) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *PolicyCreateConfig) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *PolicyCreateConfig) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *PolicyCreateConfig) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *PolicyCreateConfig) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *PolicyCreateConfig) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *PolicyCreateConfig) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *PolicyCreateConfig) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.

### GetKeyPattern

`func (o *PolicyCreateConfig) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *PolicyCreateConfig) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *PolicyCreateConfig) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.

### HasKeyPattern

`func (o *PolicyCreateConfig) HasKeyPattern() bool`

HasKeyPattern returns a boolean if a field has been set.

### GetRead

`func (o *PolicyCreateConfig) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *PolicyCreateConfig) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *PolicyCreateConfig) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *PolicyCreateConfig) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *PolicyCreateConfig) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *PolicyCreateConfig) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *PolicyCreateConfig) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *PolicyCreateConfig) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *PolicyCreateConfig) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PolicyCreateConfig) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PolicyCreateConfig) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PolicyCreateConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *PolicyCreateConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *PolicyCreateConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *PolicyCreateConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *PolicyCreateConfig) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *PolicyCreateConfig) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *PolicyCreateConfig) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *PolicyCreateConfig) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *PolicyCreateConfig) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRemovalAge

`func (o *PolicyCreateConfig) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *PolicyCreateConfig) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *PolicyCreateConfig) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.

### HasRemovalAge

`func (o *PolicyCreateConfig) HasRemovalAge() bool`

HasRemovalAge returns a boolean if a field has been set.

### GetLifecycleType

`func (o *PolicyCreateConfig) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *PolicyCreateConfig) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *PolicyCreateConfig) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.

### HasLifecycleType

`func (o *PolicyCreateConfig) HasLifecycleType() bool`

HasLifecycleType returns a boolean if a field has been set.

### GetLifecycleAge

`func (o *PolicyCreateConfig) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *PolicyCreateConfig) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *PolicyCreateConfig) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *PolicyCreateConfig) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *PolicyCreateConfig) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *PolicyCreateConfig) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *PolicyCreateConfig) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *PolicyCreateConfig) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *PolicyCreateConfig) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *PolicyCreateConfig) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *PolicyCreateConfig) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *PolicyCreateConfig) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *PolicyCreateConfig) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *PolicyCreateConfig) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *PolicyCreateConfig) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *PolicyCreateConfig) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *PolicyCreateConfig) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *PolicyCreateConfig) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *PolicyCreateConfig) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *PolicyCreateConfig) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *PolicyCreateConfig) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *PolicyCreateConfig) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *PolicyCreateConfig) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *PolicyCreateConfig) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *PolicyCreateConfig) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *PolicyCreateConfig) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *PolicyCreateConfig) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *PolicyCreateConfig) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *PolicyCreateConfig) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *PolicyCreateConfig) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *PolicyCreateConfig) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *PolicyCreateConfig) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.

### GetLifecycleWorkflowId

`func (o *PolicyCreateConfig) GetLifecycleWorkflowId() string`

GetLifecycleWorkflowId returns the LifecycleWorkflowId field if non-nil, zero value otherwise.

### GetLifecycleWorkflowIdOk

`func (o *PolicyCreateConfig) GetLifecycleWorkflowIdOk() (*string, bool)`

GetLifecycleWorkflowIdOk returns a tuple with the LifecycleWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleWorkflowId

`func (o *PolicyCreateConfig) SetLifecycleWorkflowId(v string)`

SetLifecycleWorkflowId sets LifecycleWorkflowId field to given value.

### HasLifecycleWorkflowId

`func (o *PolicyCreateConfig) HasLifecycleWorkflowId() bool`

HasLifecycleWorkflowId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *PolicyCreateConfig) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *PolicyCreateConfig) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *PolicyCreateConfig) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *PolicyCreateConfig) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetHostNamingType

`func (o *PolicyCreateConfig) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *PolicyCreateConfig) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *PolicyCreateConfig) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.

### HasHostNamingType

`func (o *PolicyCreateConfig) HasHostNamingType() bool`

HasHostNamingType returns a boolean if a field has been set.

### GetHostNamingPattern

`func (o *PolicyCreateConfig) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *PolicyCreateConfig) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *PolicyCreateConfig) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *PolicyCreateConfig) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.

### GetNamingType

`func (o *PolicyCreateConfig) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *PolicyCreateConfig) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *PolicyCreateConfig) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.

### HasNamingType

`func (o *PolicyCreateConfig) HasNamingType() bool`

HasNamingType returns a boolean if a field has been set.

### GetNamingPattern

`func (o *PolicyCreateConfig) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *PolicyCreateConfig) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *PolicyCreateConfig) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *PolicyCreateConfig) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *PolicyCreateConfig) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *PolicyCreateConfig) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *PolicyCreateConfig) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *PolicyCreateConfig) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.

### GetMaxContainers

`func (o *PolicyCreateConfig) GetMaxContainers() string`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *PolicyCreateConfig) GetMaxContainersOk() (*string, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *PolicyCreateConfig) SetMaxContainers(v string)`

SetMaxContainers sets MaxContainers field to given value.

### HasMaxContainers

`func (o *PolicyCreateConfig) HasMaxContainers() bool`

HasMaxContainers returns a boolean if a field has been set.

### GetMaxCores

`func (o *PolicyCreateConfig) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *PolicyCreateConfig) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *PolicyCreateConfig) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *PolicyCreateConfig) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetExcludeContainers

`func (o *PolicyCreateConfig) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *PolicyCreateConfig) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *PolicyCreateConfig) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *PolicyCreateConfig) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *PolicyCreateConfig) GetMaxHosts() string`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *PolicyCreateConfig) GetMaxHostsOk() (*string, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *PolicyCreateConfig) SetMaxHosts(v string)`

SetMaxHosts sets MaxHosts field to given value.

### HasMaxHosts

`func (o *PolicyCreateConfig) HasMaxHosts() bool`

HasMaxHosts returns a boolean if a field has been set.

### GetMaxPools

`func (o *PolicyCreateConfig) GetMaxPools() string`

GetMaxPools returns the MaxPools field if non-nil, zero value otherwise.

### GetMaxPoolsOk

`func (o *PolicyCreateConfig) GetMaxPoolsOk() (*string, bool)`

GetMaxPoolsOk returns a tuple with the MaxPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPools

`func (o *PolicyCreateConfig) SetMaxPools(v string)`

SetMaxPools sets MaxPools field to given value.

### HasMaxPools

`func (o *PolicyCreateConfig) HasMaxPools() bool`

HasMaxPools returns a boolean if a field has been set.

### GetMaxMemory

`func (o *PolicyCreateConfig) GetMaxMemory() MaxMemoryPolicyTypeConfiguration7MaxMemory`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *PolicyCreateConfig) GetMaxMemoryOk() (*MaxMemoryPolicyTypeConfiguration7MaxMemory, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *PolicyCreateConfig) SetMaxMemory(v MaxMemoryPolicyTypeConfiguration7MaxMemory)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxPoolMembers

`func (o *PolicyCreateConfig) GetMaxPoolMembers() string`

GetMaxPoolMembers returns the MaxPoolMembers field if non-nil, zero value otherwise.

### GetMaxPoolMembersOk

`func (o *PolicyCreateConfig) GetMaxPoolMembersOk() (*string, bool)`

GetMaxPoolMembersOk returns a tuple with the MaxPoolMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolMembers

`func (o *PolicyCreateConfig) SetMaxPoolMembers(v string)`

SetMaxPoolMembers sets MaxPoolMembers field to given value.

### HasMaxPoolMembers

`func (o *PolicyCreateConfig) HasMaxPoolMembers() bool`

HasMaxPoolMembers returns a boolean if a field has been set.

### GetMaxSnapshots

`func (o *PolicyCreateConfig) GetMaxSnapshots() string`

GetMaxSnapshots returns the MaxSnapshots field if non-nil, zero value otherwise.

### GetMaxSnapshotsOk

`func (o *PolicyCreateConfig) GetMaxSnapshotsOk() (*string, bool)`

GetMaxSnapshotsOk returns a tuple with the MaxSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshots

`func (o *PolicyCreateConfig) SetMaxSnapshots(v string)`

SetMaxSnapshots sets MaxSnapshots field to given value.

### HasMaxSnapshots

`func (o *PolicyCreateConfig) HasMaxSnapshots() bool`

HasMaxSnapshots returns a boolean if a field has been set.

### GetMaxVirtualServers

`func (o *PolicyCreateConfig) GetMaxVirtualServers() string`

GetMaxVirtualServers returns the MaxVirtualServers field if non-nil, zero value otherwise.

### GetMaxVirtualServersOk

`func (o *PolicyCreateConfig) GetMaxVirtualServersOk() (*string, bool)`

GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualServers

`func (o *PolicyCreateConfig) SetMaxVirtualServers(v string)`

SetMaxVirtualServers sets MaxVirtualServers field to given value.

### HasMaxVirtualServers

`func (o *PolicyCreateConfig) HasMaxVirtualServers() bool`

HasMaxVirtualServers returns a boolean if a field has been set.

### GetMaxVms

`func (o *PolicyCreateConfig) GetMaxVms() string`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *PolicyCreateConfig) GetMaxVmsOk() (*string, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *PolicyCreateConfig) SetMaxVms(v string)`

SetMaxVms sets MaxVms field to given value.

### HasMaxVms

`func (o *PolicyCreateConfig) HasMaxVms() bool`

HasMaxVms returns a boolean if a field has been set.

### GetMotdTitle

`func (o *PolicyCreateConfig) GetMotdTitle() string`

GetMotdTitle returns the MotdTitle field if non-nil, zero value otherwise.

### GetMotdTitleOk

`func (o *PolicyCreateConfig) GetMotdTitleOk() (*string, bool)`

GetMotdTitleOk returns a tuple with the MotdTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdTitle

`func (o *PolicyCreateConfig) SetMotdTitle(v string)`

SetMotdTitle sets MotdTitle field to given value.

### HasMotdTitle

`func (o *PolicyCreateConfig) HasMotdTitle() bool`

HasMotdTitle returns a boolean if a field has been set.

### GetMotd

`func (o *PolicyCreateConfig) GetMotd() MessageOfTheDayPolicyTypeConfiguration7Motd`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *PolicyCreateConfig) GetMotdOk() (*MessageOfTheDayPolicyTypeConfiguration7Motd, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *PolicyCreateConfig) SetMotd(v MessageOfTheDayPolicyTypeConfiguration7Motd)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *PolicyCreateConfig) HasMotd() bool`

HasMotd returns a boolean if a field has been set.

### GetMotdMessage

`func (o *PolicyCreateConfig) GetMotdMessage() string`

GetMotdMessage returns the MotdMessage field if non-nil, zero value otherwise.

### GetMotdMessageOk

`func (o *PolicyCreateConfig) GetMotdMessageOk() (*string, bool)`

GetMotdMessageOk returns a tuple with the MotdMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdMessage

`func (o *PolicyCreateConfig) SetMotdMessage(v string)`

SetMotdMessage sets MotdMessage field to given value.

### HasMotdMessage

`func (o *PolicyCreateConfig) HasMotdMessage() bool`

HasMotdMessage returns a boolean if a field has been set.

### GetMotdType

`func (o *PolicyCreateConfig) GetMotdType() string`

GetMotdType returns the MotdType field if non-nil, zero value otherwise.

### GetMotdTypeOk

`func (o *PolicyCreateConfig) GetMotdTypeOk() (*string, bool)`

GetMotdTypeOk returns a tuple with the MotdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdType

`func (o *PolicyCreateConfig) SetMotdType(v string)`

SetMotdType sets MotdType field to given value.

### HasMotdType

`func (o *PolicyCreateConfig) HasMotdType() bool`

HasMotdType returns a boolean if a field has been set.

### GetMotdFullPage

`func (o *PolicyCreateConfig) GetMotdFullPage() bool`

GetMotdFullPage returns the MotdFullPage field if non-nil, zero value otherwise.

### GetMotdFullPageOk

`func (o *PolicyCreateConfig) GetMotdFullPageOk() (*bool, bool)`

GetMotdFullPageOk returns a tuple with the MotdFullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdFullPage

`func (o *PolicyCreateConfig) SetMotdFullPage(v bool)`

SetMotdFullPage sets MotdFullPage field to given value.

### HasMotdFullPage

`func (o *PolicyCreateConfig) HasMotdFullPage() bool`

HasMotdFullPage returns a boolean if a field has been set.

### GetMotdDate

`func (o *PolicyCreateConfig) GetMotdDate() string`

GetMotdDate returns the MotdDate field if non-nil, zero value otherwise.

### GetMotdDateOk

`func (o *PolicyCreateConfig) GetMotdDateOk() (*string, bool)`

GetMotdDateOk returns a tuple with the MotdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdDate

`func (o *PolicyCreateConfig) SetMotdDate(v string)`

SetMotdDate sets MotdDate field to given value.

### HasMotdDate

`func (o *PolicyCreateConfig) HasMotdDate() bool`

HasMotdDate returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *PolicyCreateConfig) GetMaxNetworks() string`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *PolicyCreateConfig) GetMaxNetworksOk() (*string, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *PolicyCreateConfig) SetMaxNetworks(v string)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *PolicyCreateConfig) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *PolicyCreateConfig) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *PolicyCreateConfig) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *PolicyCreateConfig) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *PolicyCreateConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *PolicyCreateConfig) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *PolicyCreateConfig) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *PolicyCreateConfig) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *PolicyCreateConfig) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *PolicyCreateConfig) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *PolicyCreateConfig) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *PolicyCreateConfig) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *PolicyCreateConfig) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.

### GetMaxRouters

`func (o *PolicyCreateConfig) GetMaxRouters() string`

GetMaxRouters returns the MaxRouters field if non-nil, zero value otherwise.

### GetMaxRoutersOk

`func (o *PolicyCreateConfig) GetMaxRoutersOk() (*string, bool)`

GetMaxRoutersOk returns a tuple with the MaxRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRouters

`func (o *PolicyCreateConfig) SetMaxRouters(v string)`

SetMaxRouters sets MaxRouters field to given value.

### HasMaxRouters

`func (o *PolicyCreateConfig) HasMaxRouters() bool`

HasMaxRouters returns a boolean if a field has been set.

### GetRequiredNetworks

`func (o *PolicyCreateConfig) GetRequiredNetworks() []int64`

GetRequiredNetworks returns the RequiredNetworks field if non-nil, zero value otherwise.

### GetRequiredNetworksOk

`func (o *PolicyCreateConfig) GetRequiredNetworksOk() (*[]int64, bool)`

GetRequiredNetworksOk returns a tuple with the RequiredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNetworks

`func (o *PolicyCreateConfig) SetRequiredNetworks(v []int64)`

SetRequiredNetworks sets RequiredNetworks field to given value.

### HasRequiredNetworks

`func (o *PolicyCreateConfig) HasRequiredNetworks() bool`

HasRequiredNetworks returns a boolean if a field has been set.

### GetShutdownType

`func (o *PolicyCreateConfig) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *PolicyCreateConfig) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *PolicyCreateConfig) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.

### HasShutdownType

`func (o *PolicyCreateConfig) HasShutdownType() bool`

HasShutdownType returns a boolean if a field has been set.

### GetShutdownAge

`func (o *PolicyCreateConfig) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *PolicyCreateConfig) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *PolicyCreateConfig) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *PolicyCreateConfig) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *PolicyCreateConfig) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *PolicyCreateConfig) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *PolicyCreateConfig) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *PolicyCreateConfig) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *PolicyCreateConfig) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *PolicyCreateConfig) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *PolicyCreateConfig) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *PolicyCreateConfig) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *PolicyCreateConfig) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *PolicyCreateConfig) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *PolicyCreateConfig) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *PolicyCreateConfig) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *PolicyCreateConfig) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *PolicyCreateConfig) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *PolicyCreateConfig) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *PolicyCreateConfig) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *PolicyCreateConfig) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *PolicyCreateConfig) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *PolicyCreateConfig) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *PolicyCreateConfig) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *PolicyCreateConfig) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *PolicyCreateConfig) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *PolicyCreateConfig) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *PolicyCreateConfig) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *PolicyCreateConfig) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *PolicyCreateConfig) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *PolicyCreateConfig) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *PolicyCreateConfig) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetShutdownWorkflowId

`func (o *PolicyCreateConfig) GetShutdownWorkflowId() string`

GetShutdownWorkflowId returns the ShutdownWorkflowId field if non-nil, zero value otherwise.

### GetShutdownWorkflowIdOk

`func (o *PolicyCreateConfig) GetShutdownWorkflowIdOk() (*string, bool)`

GetShutdownWorkflowIdOk returns a tuple with the ShutdownWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWorkflowId

`func (o *PolicyCreateConfig) SetShutdownWorkflowId(v string)`

SetShutdownWorkflowId sets ShutdownWorkflowId field to given value.

### HasShutdownWorkflowId

`func (o *PolicyCreateConfig) HasShutdownWorkflowId() bool`

HasShutdownWorkflowId returns a boolean if a field has been set.

### GetStorageServerId

`func (o *PolicyCreateConfig) GetStorageServerId() string`

GetStorageServerId returns the StorageServerId field if non-nil, zero value otherwise.

### GetStorageServerIdOk

`func (o *PolicyCreateConfig) GetStorageServerIdOk() (*string, bool)`

GetStorageServerIdOk returns a tuple with the StorageServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerId

`func (o *PolicyCreateConfig) SetStorageServerId(v string)`

SetStorageServerId sets StorageServerId field to given value.

### HasStorageServerId

`func (o *PolicyCreateConfig) HasStorageServerId() bool`

HasStorageServerId returns a boolean if a field has been set.

### GetStrict

`func (o *PolicyCreateConfig) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *PolicyCreateConfig) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *PolicyCreateConfig) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *PolicyCreateConfig) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetKey

`func (o *PolicyCreateConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PolicyCreateConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PolicyCreateConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PolicyCreateConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *PolicyCreateConfig) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *PolicyCreateConfig) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *PolicyCreateConfig) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *PolicyCreateConfig) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *PolicyCreateConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PolicyCreateConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PolicyCreateConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PolicyCreateConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreateUserType

`func (o *PolicyCreateConfig) GetCreateUserType() string`

GetCreateUserType returns the CreateUserType field if non-nil, zero value otherwise.

### GetCreateUserTypeOk

`func (o *PolicyCreateConfig) GetCreateUserTypeOk() (*string, bool)`

GetCreateUserTypeOk returns a tuple with the CreateUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserType

`func (o *PolicyCreateConfig) SetCreateUserType(v string)`

SetCreateUserType sets CreateUserType field to given value.

### HasCreateUserType

`func (o *PolicyCreateConfig) HasCreateUserType() bool`

HasCreateUserType returns a boolean if a field has been set.

### GetCreateUser

`func (o *PolicyCreateConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *PolicyCreateConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *PolicyCreateConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *PolicyCreateConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *PolicyCreateConfig) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *PolicyCreateConfig) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *PolicyCreateConfig) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *PolicyCreateConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


