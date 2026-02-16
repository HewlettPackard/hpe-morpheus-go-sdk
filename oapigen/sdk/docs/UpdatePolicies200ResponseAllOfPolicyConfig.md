# UpdatePolicies200ResponseAllOfPolicyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | **string** |  | 
**WorkflowId** | **string** |  | 
**FlowId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**CreateBackupType** | **string** |  | 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**BackupStorageIds** | **[]string** |  | 
**MaxPrice** | **float32** |  | 
**MaxPriceCurrency** | Pointer to **string** |  | [optional] 
**MaxPriceUnit** | Pointer to **string** |  | [optional] 
**ServerNamingType** | **string** |  | 
**ServerNamingPattern** | Pointer to **string** |  | [optional] 
**ServerNamingConflict** | Pointer to **bool** |  | [optional] 
**KeyPattern** | **string** |  | 
**Read** | Pointer to **bool** |  | [optional] 
**Write** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **bool** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**List** | Pointer to **bool** |  | [optional] 
**RemovalAge** | **string** |  | 
**LifecycleType** | **string** |  | 
**LifecycleAge** | Pointer to **string** |  | [optional] 
**LifecycleRenewal** | Pointer to **string** |  | [optional] 
**LifecycleNotify** | Pointer to **string** |  | [optional] 
**LifecycleMessage** | Pointer to **string** |  | [optional] 
**LifecycleAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**LifecycleWorkflowId** | Pointer to **string** |  | [optional] 
**LifecycleHideFixed** | Pointer to **bool** |  | [optional] 
**HostNamingType** | **string** |  | 
**HostNamingPattern** | Pointer to **string** |  | [optional] 
**NamingType** | **string** |  | 
**NamingPattern** | Pointer to **string** |  | [optional] 
**NamingConflict** | Pointer to **bool** |  | [optional] 
**MaxContainers** | **string** |  | 
**MaxCores** | **string** |  | 
**ExcludeContainers** | Pointer to **string** |  | [optional] [default to "off"]
**MaxHosts** | **string** |  | 
**MaxPools** | **string** |  | 
**MaxMemory** | **string** |  | 
**MaxPoolMembers** | **string** |  | 
**MaxSnapshots** | **string** |  | 
**MaxStorage** | **string** |  | 
**MaxVirtualServers** | **string** |  | 
**MaxVms** | **string** |  | 
**MotdTitle** | Pointer to **string** |  | [optional] 
**Motd** | [**MessageOfTheDayPolicyTypeConfiguration5Motd**](MessageOfTheDayPolicyTypeConfiguration5Motd.md) |  | 
**MotdMessage** | Pointer to **string** |  | [optional] 
**MotdType** | Pointer to **string** |  | [optional] 
**MotdFullPage** | Pointer to [**MessageOfTheDayPolicyTypeConfiguration5MotdFullPage**](MessageOfTheDayPolicyTypeConfiguration5MotdFullPage.md) |  | [optional] 
**MotdDate** | Pointer to **string** |  | [optional] 
**MaxNetworks** | **string** |  | 
**PowerScheduleType** | **string** |  | 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**PowerScheduleHideFixed** | Pointer to **bool** |  | [optional] 
**MaxRouters** | **string** |  | 
**RequiredNetworks** | **[]int64** |  | 
**ShutdownType** | **string** |  | 
**ShutdownAge** | Pointer to **string** |  | [optional] 
**ShutdownRenewal** | Pointer to **string** |  | [optional] 
**ShutdownNotify** | Pointer to **string** |  | [optional] 
**ShutdownMessage** | Pointer to **string** |  | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**ShutdownHideFixed** | Pointer to **bool** |  | [optional] 
**ShutdownWorkflowId** | Pointer to **string** |  | [optional] 
**StorageServerId** | **string** |  | 
**Strict** | **bool** |  | 
**Key** | Pointer to **string** |  | [optional] 
**ValueListId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CreateUserType** | **string** |  | 
**CreateUser** | Pointer to **bool** |  | [optional] 
**UserGroup** | **string** |  | 

## Methods

### NewUpdatePolicies200ResponseAllOfPolicyConfig

`func NewUpdatePolicies200ResponseAllOfPolicyConfig(accountIntegrationId string, workflowId string, createBackupType string, backupStorageIds []string, maxPrice float32, serverNamingType string, keyPattern string, removalAge string, lifecycleType string, hostNamingType string, namingType string, maxContainers string, maxCores string, maxHosts string, maxPools string, maxMemory string, maxPoolMembers string, maxSnapshots string, maxStorage string, maxVirtualServers string, maxVms string, motd MessageOfTheDayPolicyTypeConfiguration5Motd, maxNetworks string, powerScheduleType string, maxRouters string, requiredNetworks []int64, shutdownType string, storageServerId string, strict bool, createUserType string, userGroup string, ) *UpdatePolicies200ResponseAllOfPolicyConfig`

NewUpdatePolicies200ResponseAllOfPolicyConfig instantiates a new UpdatePolicies200ResponseAllOfPolicyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePolicies200ResponseAllOfPolicyConfigWithDefaults

`func NewUpdatePolicies200ResponseAllOfPolicyConfigWithDefaults() *UpdatePolicies200ResponseAllOfPolicyConfig`

NewUpdatePolicies200ResponseAllOfPolicyConfigWithDefaults instantiates a new UpdatePolicies200ResponseAllOfPolicyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.


### GetWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetFlowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetCreateBackupType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.


### GetCreateBackup

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackupStorageIds

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetBackupStorageIds() []string`

GetBackupStorageIds returns the BackupStorageIds field if non-nil, zero value otherwise.

### GetBackupStorageIdsOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetBackupStorageIdsOk() (*[]string, bool)`

GetBackupStorageIdsOk returns a tuple with the BackupStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageIds

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetBackupStorageIds(v []string)`

SetBackupStorageIds sets BackupStorageIds field to given value.


### GetMaxPrice

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.


### GetMaxPriceCurrency

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPriceCurrency() string`

GetMaxPriceCurrency returns the MaxPriceCurrency field if non-nil, zero value otherwise.

### GetMaxPriceCurrencyOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPriceCurrencyOk() (*string, bool)`

GetMaxPriceCurrencyOk returns a tuple with the MaxPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceCurrency

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxPriceCurrency(v string)`

SetMaxPriceCurrency sets MaxPriceCurrency field to given value.

### HasMaxPriceCurrency

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMaxPriceCurrency() bool`

HasMaxPriceCurrency returns a boolean if a field has been set.

### GetMaxPriceUnit

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPriceUnit() string`

GetMaxPriceUnit returns the MaxPriceUnit field if non-nil, zero value otherwise.

### GetMaxPriceUnitOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPriceUnitOk() (*string, bool)`

GetMaxPriceUnitOk returns a tuple with the MaxPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceUnit

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxPriceUnit(v string)`

SetMaxPriceUnit sets MaxPriceUnit field to given value.

### HasMaxPriceUnit

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMaxPriceUnit() bool`

HasMaxPriceUnit returns a boolean if a field has been set.

### GetServerNamingType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.


### GetServerNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.

### GetKeyPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.


### GetRead

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRemovalAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.


### GetLifecycleType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.


### GetLifecycleAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetLifecycleWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleWorkflowId() string`

GetLifecycleWorkflowId returns the LifecycleWorkflowId field if non-nil, zero value otherwise.

### GetLifecycleWorkflowIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleWorkflowIdOk() (*string, bool)`

GetLifecycleWorkflowIdOk returns a tuple with the LifecycleWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleWorkflowId(v string)`

SetLifecycleWorkflowId sets LifecycleWorkflowId field to given value.

### HasLifecycleWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleWorkflowId() bool`

HasLifecycleWorkflowId returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.

### GetHostNamingType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.


### GetHostNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.

### GetNamingType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.


### GetNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.

### GetMaxContainers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxContainers() string`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxContainersOk() (*string, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxContainers(v string)`

SetMaxContainers sets MaxContainers field to given value.


### GetMaxCores

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.


### GetExcludeContainers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxHosts() string`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxHostsOk() (*string, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxHosts(v string)`

SetMaxHosts sets MaxHosts field to given value.


### GetMaxPools

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPools() string`

GetMaxPools returns the MaxPools field if non-nil, zero value otherwise.

### GetMaxPoolsOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPoolsOk() (*string, bool)`

GetMaxPoolsOk returns a tuple with the MaxPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPools

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxPools(v string)`

SetMaxPools sets MaxPools field to given value.


### GetMaxMemory

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxPoolMembers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPoolMembers() string`

GetMaxPoolMembers returns the MaxPoolMembers field if non-nil, zero value otherwise.

### GetMaxPoolMembersOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxPoolMembersOk() (*string, bool)`

GetMaxPoolMembersOk returns a tuple with the MaxPoolMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolMembers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxPoolMembers(v string)`

SetMaxPoolMembers sets MaxPoolMembers field to given value.


### GetMaxSnapshots

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxSnapshots() string`

GetMaxSnapshots returns the MaxSnapshots field if non-nil, zero value otherwise.

### GetMaxSnapshotsOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxSnapshotsOk() (*string, bool)`

GetMaxSnapshotsOk returns a tuple with the MaxSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshots

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxSnapshots(v string)`

SetMaxSnapshots sets MaxSnapshots field to given value.


### GetMaxStorage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxVirtualServers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxVirtualServers() string`

GetMaxVirtualServers returns the MaxVirtualServers field if non-nil, zero value otherwise.

### GetMaxVirtualServersOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxVirtualServersOk() (*string, bool)`

GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualServers

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxVirtualServers(v string)`

SetMaxVirtualServers sets MaxVirtualServers field to given value.


### GetMaxVms

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxVms() string`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxVmsOk() (*string, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxVms(v string)`

SetMaxVms sets MaxVms field to given value.


### GetMotdTitle

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdTitle() string`

GetMotdTitle returns the MotdTitle field if non-nil, zero value otherwise.

### GetMotdTitleOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdTitleOk() (*string, bool)`

GetMotdTitleOk returns a tuple with the MotdTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdTitle

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMotdTitle(v string)`

SetMotdTitle sets MotdTitle field to given value.

### HasMotdTitle

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMotdTitle() bool`

HasMotdTitle returns a boolean if a field has been set.

### GetMotd

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotd() MessageOfTheDayPolicyTypeConfiguration5Motd`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdOk() (*MessageOfTheDayPolicyTypeConfiguration5Motd, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMotd(v MessageOfTheDayPolicyTypeConfiguration5Motd)`

SetMotd sets Motd field to given value.


### GetMotdMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdMessage() string`

GetMotdMessage returns the MotdMessage field if non-nil, zero value otherwise.

### GetMotdMessageOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdMessageOk() (*string, bool)`

GetMotdMessageOk returns a tuple with the MotdMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMotdMessage(v string)`

SetMotdMessage sets MotdMessage field to given value.

### HasMotdMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMotdMessage() bool`

HasMotdMessage returns a boolean if a field has been set.

### GetMotdType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdType() string`

GetMotdType returns the MotdType field if non-nil, zero value otherwise.

### GetMotdTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdTypeOk() (*string, bool)`

GetMotdTypeOk returns a tuple with the MotdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMotdType(v string)`

SetMotdType sets MotdType field to given value.

### HasMotdType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMotdType() bool`

HasMotdType returns a boolean if a field has been set.

### GetMotdFullPage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdFullPage() MessageOfTheDayPolicyTypeConfiguration5MotdFullPage`

GetMotdFullPage returns the MotdFullPage field if non-nil, zero value otherwise.

### GetMotdFullPageOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdFullPageOk() (*MessageOfTheDayPolicyTypeConfiguration5MotdFullPage, bool)`

GetMotdFullPageOk returns a tuple with the MotdFullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdFullPage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMotdFullPage(v MessageOfTheDayPolicyTypeConfiguration5MotdFullPage)`

SetMotdFullPage sets MotdFullPage field to given value.

### HasMotdFullPage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMotdFullPage() bool`

HasMotdFullPage returns a boolean if a field has been set.

### GetMotdDate

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdDate() string`

GetMotdDate returns the MotdDate field if non-nil, zero value otherwise.

### GetMotdDateOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMotdDateOk() (*string, bool)`

GetMotdDateOk returns a tuple with the MotdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdDate

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMotdDate(v string)`

SetMotdDate sets MotdDate field to given value.

### HasMotdDate

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasMotdDate() bool`

HasMotdDate returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxNetworks() string`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxNetworksOk() (*string, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxNetworks(v string)`

SetMaxNetworks sets MaxNetworks field to given value.


### GetPowerScheduleType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.


### GetPowerSchedule

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.

### GetMaxRouters

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxRouters() string`

GetMaxRouters returns the MaxRouters field if non-nil, zero value otherwise.

### GetMaxRoutersOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetMaxRoutersOk() (*string, bool)`

GetMaxRoutersOk returns a tuple with the MaxRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRouters

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetMaxRouters(v string)`

SetMaxRouters sets MaxRouters field to given value.


### GetRequiredNetworks

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetRequiredNetworks() []int64`

GetRequiredNetworks returns the RequiredNetworks field if non-nil, zero value otherwise.

### GetRequiredNetworksOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetRequiredNetworksOk() (*[]int64, bool)`

GetRequiredNetworksOk returns a tuple with the RequiredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNetworks

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetRequiredNetworks(v []int64)`

SetRequiredNetworks sets RequiredNetworks field to given value.


### GetShutdownType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.


### GetShutdownAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetShutdownWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownWorkflowId() string`

GetShutdownWorkflowId returns the ShutdownWorkflowId field if non-nil, zero value otherwise.

### GetShutdownWorkflowIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetShutdownWorkflowIdOk() (*string, bool)`

GetShutdownWorkflowIdOk returns a tuple with the ShutdownWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetShutdownWorkflowId(v string)`

SetShutdownWorkflowId sets ShutdownWorkflowId field to given value.

### HasShutdownWorkflowId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasShutdownWorkflowId() bool`

HasShutdownWorkflowId returns a boolean if a field has been set.

### GetStorageServerId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetStorageServerId() string`

GetStorageServerId returns the StorageServerId field if non-nil, zero value otherwise.

### GetStorageServerIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetStorageServerIdOk() (*string, bool)`

GetStorageServerIdOk returns a tuple with the StorageServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetStorageServerId(v string)`

SetStorageServerId sets StorageServerId field to given value.


### GetStrict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetStrict(v bool)`

SetStrict sets Strict field to given value.


### GetKey

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreateUserType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateUserType() string`

GetCreateUserType returns the CreateUserType field if non-nil, zero value otherwise.

### GetCreateUserTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateUserTypeOk() (*string, bool)`

GetCreateUserTypeOk returns a tuple with the CreateUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserType

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetCreateUserType(v string)`

SetCreateUserType sets CreateUserType field to given value.


### GetCreateUser

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UpdatePolicies200ResponseAllOfPolicyConfig) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


