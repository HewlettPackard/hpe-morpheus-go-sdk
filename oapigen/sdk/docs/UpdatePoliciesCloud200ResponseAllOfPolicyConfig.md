# UpdatePoliciesCloud200ResponseAllOfPolicyConfig

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
**Motd** | [**MessageOfTheDayPolicyTypeConfiguration13Motd**](MessageOfTheDayPolicyTypeConfiguration13Motd.md) |  | 
**MotdMessage** | Pointer to **string** |  | [optional] 
**MotdType** | Pointer to **string** |  | [optional] 
**MotdFullPage** | Pointer to [**MessageOfTheDayPolicyTypeConfiguration13MotdFullPage**](MessageOfTheDayPolicyTypeConfiguration13MotdFullPage.md) |  | [optional] 
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

### NewUpdatePoliciesCloud200ResponseAllOfPolicyConfig

`func NewUpdatePoliciesCloud200ResponseAllOfPolicyConfig(accountIntegrationId string, workflowId string, createBackupType string, backupStorageIds []string, maxPrice float32, serverNamingType string, keyPattern string, removalAge string, lifecycleType string, hostNamingType string, namingType string, maxContainers string, maxCores string, maxHosts string, maxPools string, maxMemory string, maxPoolMembers string, maxSnapshots string, maxStorage string, maxVirtualServers string, maxVms string, motd MessageOfTheDayPolicyTypeConfiguration13Motd, maxNetworks string, powerScheduleType string, maxRouters string, requiredNetworks []int64, shutdownType string, storageServerId string, strict bool, createUserType string, userGroup string, ) *UpdatePoliciesCloud200ResponseAllOfPolicyConfig`

NewUpdatePoliciesCloud200ResponseAllOfPolicyConfig instantiates a new UpdatePoliciesCloud200ResponseAllOfPolicyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePoliciesCloud200ResponseAllOfPolicyConfigWithDefaults

`func NewUpdatePoliciesCloud200ResponseAllOfPolicyConfigWithDefaults() *UpdatePoliciesCloud200ResponseAllOfPolicyConfig`

NewUpdatePoliciesCloud200ResponseAllOfPolicyConfigWithDefaults instantiates a new UpdatePoliciesCloud200ResponseAllOfPolicyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.


### GetWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetFlowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetCreateBackupType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.


### GetCreateBackup

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackupStorageIds

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetBackupStorageIds() []string`

GetBackupStorageIds returns the BackupStorageIds field if non-nil, zero value otherwise.

### GetBackupStorageIdsOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetBackupStorageIdsOk() (*[]string, bool)`

GetBackupStorageIdsOk returns a tuple with the BackupStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageIds

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetBackupStorageIds(v []string)`

SetBackupStorageIds sets BackupStorageIds field to given value.


### GetMaxPrice

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.


### GetMaxPriceCurrency

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPriceCurrency() string`

GetMaxPriceCurrency returns the MaxPriceCurrency field if non-nil, zero value otherwise.

### GetMaxPriceCurrencyOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPriceCurrencyOk() (*string, bool)`

GetMaxPriceCurrencyOk returns a tuple with the MaxPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceCurrency

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxPriceCurrency(v string)`

SetMaxPriceCurrency sets MaxPriceCurrency field to given value.

### HasMaxPriceCurrency

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMaxPriceCurrency() bool`

HasMaxPriceCurrency returns a boolean if a field has been set.

### GetMaxPriceUnit

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPriceUnit() string`

GetMaxPriceUnit returns the MaxPriceUnit field if non-nil, zero value otherwise.

### GetMaxPriceUnitOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPriceUnitOk() (*string, bool)`

GetMaxPriceUnitOk returns a tuple with the MaxPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceUnit

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxPriceUnit(v string)`

SetMaxPriceUnit sets MaxPriceUnit field to given value.

### HasMaxPriceUnit

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMaxPriceUnit() bool`

HasMaxPriceUnit returns a boolean if a field has been set.

### GetServerNamingType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.


### GetServerNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.

### GetKeyPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.


### GetRead

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRemovalAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.


### GetLifecycleType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.


### GetLifecycleAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetLifecycleWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleWorkflowId() string`

GetLifecycleWorkflowId returns the LifecycleWorkflowId field if non-nil, zero value otherwise.

### GetLifecycleWorkflowIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleWorkflowIdOk() (*string, bool)`

GetLifecycleWorkflowIdOk returns a tuple with the LifecycleWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleWorkflowId(v string)`

SetLifecycleWorkflowId sets LifecycleWorkflowId field to given value.

### HasLifecycleWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleWorkflowId() bool`

HasLifecycleWorkflowId returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.

### GetHostNamingType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.


### GetHostNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.

### GetNamingType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.


### GetNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.

### GetMaxContainers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxContainers() string`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxContainersOk() (*string, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxContainers(v string)`

SetMaxContainers sets MaxContainers field to given value.


### GetMaxCores

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.


### GetExcludeContainers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxHosts() string`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxHostsOk() (*string, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxHosts(v string)`

SetMaxHosts sets MaxHosts field to given value.


### GetMaxPools

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPools() string`

GetMaxPools returns the MaxPools field if non-nil, zero value otherwise.

### GetMaxPoolsOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPoolsOk() (*string, bool)`

GetMaxPoolsOk returns a tuple with the MaxPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPools

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxPools(v string)`

SetMaxPools sets MaxPools field to given value.


### GetMaxMemory

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxPoolMembers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPoolMembers() string`

GetMaxPoolMembers returns the MaxPoolMembers field if non-nil, zero value otherwise.

### GetMaxPoolMembersOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxPoolMembersOk() (*string, bool)`

GetMaxPoolMembersOk returns a tuple with the MaxPoolMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolMembers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxPoolMembers(v string)`

SetMaxPoolMembers sets MaxPoolMembers field to given value.


### GetMaxSnapshots

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxSnapshots() string`

GetMaxSnapshots returns the MaxSnapshots field if non-nil, zero value otherwise.

### GetMaxSnapshotsOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxSnapshotsOk() (*string, bool)`

GetMaxSnapshotsOk returns a tuple with the MaxSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshots

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxSnapshots(v string)`

SetMaxSnapshots sets MaxSnapshots field to given value.


### GetMaxStorage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxVirtualServers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxVirtualServers() string`

GetMaxVirtualServers returns the MaxVirtualServers field if non-nil, zero value otherwise.

### GetMaxVirtualServersOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxVirtualServersOk() (*string, bool)`

GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualServers

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxVirtualServers(v string)`

SetMaxVirtualServers sets MaxVirtualServers field to given value.


### GetMaxVms

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxVms() string`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxVmsOk() (*string, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxVms(v string)`

SetMaxVms sets MaxVms field to given value.


### GetMotdTitle

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdTitle() string`

GetMotdTitle returns the MotdTitle field if non-nil, zero value otherwise.

### GetMotdTitleOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdTitleOk() (*string, bool)`

GetMotdTitleOk returns a tuple with the MotdTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdTitle

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMotdTitle(v string)`

SetMotdTitle sets MotdTitle field to given value.

### HasMotdTitle

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMotdTitle() bool`

HasMotdTitle returns a boolean if a field has been set.

### GetMotd

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotd() MessageOfTheDayPolicyTypeConfiguration13Motd`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdOk() (*MessageOfTheDayPolicyTypeConfiguration13Motd, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMotd(v MessageOfTheDayPolicyTypeConfiguration13Motd)`

SetMotd sets Motd field to given value.


### GetMotdMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdMessage() string`

GetMotdMessage returns the MotdMessage field if non-nil, zero value otherwise.

### GetMotdMessageOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdMessageOk() (*string, bool)`

GetMotdMessageOk returns a tuple with the MotdMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMotdMessage(v string)`

SetMotdMessage sets MotdMessage field to given value.

### HasMotdMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMotdMessage() bool`

HasMotdMessage returns a boolean if a field has been set.

### GetMotdType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdType() string`

GetMotdType returns the MotdType field if non-nil, zero value otherwise.

### GetMotdTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdTypeOk() (*string, bool)`

GetMotdTypeOk returns a tuple with the MotdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMotdType(v string)`

SetMotdType sets MotdType field to given value.

### HasMotdType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMotdType() bool`

HasMotdType returns a boolean if a field has been set.

### GetMotdFullPage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdFullPage() MessageOfTheDayPolicyTypeConfiguration13MotdFullPage`

GetMotdFullPage returns the MotdFullPage field if non-nil, zero value otherwise.

### GetMotdFullPageOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdFullPageOk() (*MessageOfTheDayPolicyTypeConfiguration13MotdFullPage, bool)`

GetMotdFullPageOk returns a tuple with the MotdFullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdFullPage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMotdFullPage(v MessageOfTheDayPolicyTypeConfiguration13MotdFullPage)`

SetMotdFullPage sets MotdFullPage field to given value.

### HasMotdFullPage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMotdFullPage() bool`

HasMotdFullPage returns a boolean if a field has been set.

### GetMotdDate

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdDate() string`

GetMotdDate returns the MotdDate field if non-nil, zero value otherwise.

### GetMotdDateOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMotdDateOk() (*string, bool)`

GetMotdDateOk returns a tuple with the MotdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdDate

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMotdDate(v string)`

SetMotdDate sets MotdDate field to given value.

### HasMotdDate

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasMotdDate() bool`

HasMotdDate returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxNetworks() string`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxNetworksOk() (*string, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxNetworks(v string)`

SetMaxNetworks sets MaxNetworks field to given value.


### GetPowerScheduleType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.


### GetPowerSchedule

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.

### GetMaxRouters

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxRouters() string`

GetMaxRouters returns the MaxRouters field if non-nil, zero value otherwise.

### GetMaxRoutersOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetMaxRoutersOk() (*string, bool)`

GetMaxRoutersOk returns a tuple with the MaxRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRouters

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetMaxRouters(v string)`

SetMaxRouters sets MaxRouters field to given value.


### GetRequiredNetworks

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetRequiredNetworks() []int64`

GetRequiredNetworks returns the RequiredNetworks field if non-nil, zero value otherwise.

### GetRequiredNetworksOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetRequiredNetworksOk() (*[]int64, bool)`

GetRequiredNetworksOk returns a tuple with the RequiredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNetworks

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetRequiredNetworks(v []int64)`

SetRequiredNetworks sets RequiredNetworks field to given value.


### GetShutdownType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.


### GetShutdownAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetShutdownWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownWorkflowId() string`

GetShutdownWorkflowId returns the ShutdownWorkflowId field if non-nil, zero value otherwise.

### GetShutdownWorkflowIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetShutdownWorkflowIdOk() (*string, bool)`

GetShutdownWorkflowIdOk returns a tuple with the ShutdownWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetShutdownWorkflowId(v string)`

SetShutdownWorkflowId sets ShutdownWorkflowId field to given value.

### HasShutdownWorkflowId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasShutdownWorkflowId() bool`

HasShutdownWorkflowId returns a boolean if a field has been set.

### GetStorageServerId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetStorageServerId() string`

GetStorageServerId returns the StorageServerId field if non-nil, zero value otherwise.

### GetStorageServerIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetStorageServerIdOk() (*string, bool)`

GetStorageServerIdOk returns a tuple with the StorageServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetStorageServerId(v string)`

SetStorageServerId sets StorageServerId field to given value.


### GetStrict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetStrict(v bool)`

SetStrict sets Strict field to given value.


### GetKey

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreateUserType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateUserType() string`

GetCreateUserType returns the CreateUserType field if non-nil, zero value otherwise.

### GetCreateUserTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateUserTypeOk() (*string, bool)`

GetCreateUserTypeOk returns a tuple with the CreateUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetCreateUserType(v string)`

SetCreateUserType sets CreateUserType field to given value.


### GetCreateUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicyConfig) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


