# AddPoliciesGroupRequestPolicyPolicyTypeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | Pointer to **string** |  | [optional] 
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
**MaxMemory** | Pointer to **string** |  | [optional] 
**MaxPoolMembers** | Pointer to **string** |  | [optional] 
**MaxVirtualServers** | Pointer to **string** |  | [optional] 
**MaxVms** | Pointer to **string** |  | [optional] 
**MotdTitle** | Pointer to **string** |  | [optional] 
**Motd** | Pointer to [**MessageOfTheDayPolicyTypeConfigurationMotd**](MessageOfTheDayPolicyTypeConfigurationMotd.md) |  | [optional] 
**MotdMessage** | Pointer to **string** |  | [optional] 
**MotdType** | Pointer to **string** |  | [optional] 
**MotdFullPage** | Pointer to **bool** |  | [optional] 
**MotdDate** | Pointer to **time.Time** |  | [optional] 
**MaxNetworks** | Pointer to **string** |  | [optional] 
**PowerScheduleType** | Pointer to **string** |  | [optional] 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**PowerScheduleHideFixed** | Pointer to **bool** |  | [optional] 
**MaxRouters** | Pointer to **string** |  | [optional] 
**ShutdownType** | Pointer to **string** |  | [optional] 
**ShutdownAge** | Pointer to **string** |  | [optional] 
**ShutdownRenewal** | Pointer to **string** |  | [optional] 
**ShutdownNotify** | Pointer to **string** |  | [optional] 
**ShutdownMessage** | Pointer to **string** |  | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**ShutdownHideFixed** | Pointer to **bool** |  | [optional] 
**StorageServerId** | Pointer to **string** |  | [optional] 
**Strict** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**ValueListId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CreateUserType** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewAddPoliciesGroupRequestPolicyPolicyTypeConfig

`func NewAddPoliciesGroupRequestPolicyPolicyTypeConfig() *AddPoliciesGroupRequestPolicyPolicyTypeConfig`

NewAddPoliciesGroupRequestPolicyPolicyTypeConfig instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPoliciesGroupRequestPolicyPolicyTypeConfigWithDefaults

`func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigWithDefaults() *AddPoliciesGroupRequestPolicyPolicyTypeConfig`

NewAddPoliciesGroupRequestPolicyPolicyTypeConfigWithDefaults instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetCreateBackupType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.

### HasCreateBackupType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasCreateBackupType() bool`

HasCreateBackupType returns a boolean if a field has been set.

### GetCreateBackup

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackupStorageIds

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetBackupStorageIds() []int64`

GetBackupStorageIds returns the BackupStorageIds field if non-nil, zero value otherwise.

### GetBackupStorageIdsOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetBackupStorageIdsOk() (*[]int64, bool)`

GetBackupStorageIdsOk returns a tuple with the BackupStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageIds

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetBackupStorageIds(v []int64)`

SetBackupStorageIds sets BackupStorageIds field to given value.

### HasBackupStorageIds

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasBackupStorageIds() bool`

HasBackupStorageIds returns a boolean if a field has been set.

### GetMaxPrice

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetMaxPriceCurrency

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPriceCurrency() string`

GetMaxPriceCurrency returns the MaxPriceCurrency field if non-nil, zero value otherwise.

### GetMaxPriceCurrencyOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPriceCurrencyOk() (*string, bool)`

GetMaxPriceCurrencyOk returns a tuple with the MaxPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceCurrency

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxPriceCurrency(v string)`

SetMaxPriceCurrency sets MaxPriceCurrency field to given value.

### HasMaxPriceCurrency

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxPriceCurrency() bool`

HasMaxPriceCurrency returns a boolean if a field has been set.

### GetMaxPriceUnit

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPriceUnit() string`

GetMaxPriceUnit returns the MaxPriceUnit field if non-nil, zero value otherwise.

### GetMaxPriceUnitOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPriceUnitOk() (*string, bool)`

GetMaxPriceUnitOk returns a tuple with the MaxPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceUnit

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxPriceUnit(v string)`

SetMaxPriceUnit sets MaxPriceUnit field to given value.

### HasMaxPriceUnit

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxPriceUnit() bool`

HasMaxPriceUnit returns a boolean if a field has been set.

### GetServerNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.

### HasServerNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasServerNamingType() bool`

HasServerNamingType returns a boolean if a field has been set.

### GetServerNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.

### GetKeyPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.

### HasKeyPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasKeyPattern() bool`

HasKeyPattern returns a boolean if a field has been set.

### GetRead

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRemovalAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.

### HasRemovalAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasRemovalAge() bool`

HasRemovalAge returns a boolean if a field has been set.

### GetLifecycleType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.

### HasLifecycleType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleType() bool`

HasLifecycleType returns a boolean if a field has been set.

### GetLifecycleAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetHostNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.

### HasHostNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasHostNamingType() bool`

HasHostNamingType returns a boolean if a field has been set.

### GetHostNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.

### GetNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.

### HasNamingType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasNamingType() bool`

HasNamingType returns a boolean if a field has been set.

### GetNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.

### GetMaxContainers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxContainers() string`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxContainersOk() (*string, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxContainers(v string)`

SetMaxContainers sets MaxContainers field to given value.

### HasMaxContainers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxContainers() bool`

HasMaxContainers returns a boolean if a field has been set.

### GetMaxCores

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetExcludeContainers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxHosts() string`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxHostsOk() (*string, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxHosts(v string)`

SetMaxHosts sets MaxHosts field to given value.

### HasMaxHosts

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxHosts() bool`

HasMaxHosts returns a boolean if a field has been set.

### GetMaxPools

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPools() string`

GetMaxPools returns the MaxPools field if non-nil, zero value otherwise.

### GetMaxPoolsOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPoolsOk() (*string, bool)`

GetMaxPoolsOk returns a tuple with the MaxPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPools

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxPools(v string)`

SetMaxPools sets MaxPools field to given value.

### HasMaxPools

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxPools() bool`

HasMaxPools returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxPoolMembers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPoolMembers() string`

GetMaxPoolMembers returns the MaxPoolMembers field if non-nil, zero value otherwise.

### GetMaxPoolMembersOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxPoolMembersOk() (*string, bool)`

GetMaxPoolMembersOk returns a tuple with the MaxPoolMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolMembers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxPoolMembers(v string)`

SetMaxPoolMembers sets MaxPoolMembers field to given value.

### HasMaxPoolMembers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxPoolMembers() bool`

HasMaxPoolMembers returns a boolean if a field has been set.

### GetMaxVirtualServers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxVirtualServers() string`

GetMaxVirtualServers returns the MaxVirtualServers field if non-nil, zero value otherwise.

### GetMaxVirtualServersOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxVirtualServersOk() (*string, bool)`

GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualServers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxVirtualServers(v string)`

SetMaxVirtualServers sets MaxVirtualServers field to given value.

### HasMaxVirtualServers

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxVirtualServers() bool`

HasMaxVirtualServers returns a boolean if a field has been set.

### GetMaxVms

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxVms() string`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxVmsOk() (*string, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxVms(v string)`

SetMaxVms sets MaxVms field to given value.

### HasMaxVms

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxVms() bool`

HasMaxVms returns a boolean if a field has been set.

### GetMotdTitle

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdTitle() string`

GetMotdTitle returns the MotdTitle field if non-nil, zero value otherwise.

### GetMotdTitleOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdTitleOk() (*string, bool)`

GetMotdTitleOk returns a tuple with the MotdTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdTitle

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMotdTitle(v string)`

SetMotdTitle sets MotdTitle field to given value.

### HasMotdTitle

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMotdTitle() bool`

HasMotdTitle returns a boolean if a field has been set.

### GetMotd

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotd() MessageOfTheDayPolicyTypeConfigurationMotd`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdOk() (*MessageOfTheDayPolicyTypeConfigurationMotd, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMotd(v MessageOfTheDayPolicyTypeConfigurationMotd)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMotd() bool`

HasMotd returns a boolean if a field has been set.

### GetMotdMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdMessage() string`

GetMotdMessage returns the MotdMessage field if non-nil, zero value otherwise.

### GetMotdMessageOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdMessageOk() (*string, bool)`

GetMotdMessageOk returns a tuple with the MotdMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMotdMessage(v string)`

SetMotdMessage sets MotdMessage field to given value.

### HasMotdMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMotdMessage() bool`

HasMotdMessage returns a boolean if a field has been set.

### GetMotdType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdType() string`

GetMotdType returns the MotdType field if non-nil, zero value otherwise.

### GetMotdTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdTypeOk() (*string, bool)`

GetMotdTypeOk returns a tuple with the MotdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMotdType(v string)`

SetMotdType sets MotdType field to given value.

### HasMotdType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMotdType() bool`

HasMotdType returns a boolean if a field has been set.

### GetMotdFullPage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdFullPage() bool`

GetMotdFullPage returns the MotdFullPage field if non-nil, zero value otherwise.

### GetMotdFullPageOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdFullPageOk() (*bool, bool)`

GetMotdFullPageOk returns a tuple with the MotdFullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdFullPage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMotdFullPage(v bool)`

SetMotdFullPage sets MotdFullPage field to given value.

### HasMotdFullPage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMotdFullPage() bool`

HasMotdFullPage returns a boolean if a field has been set.

### GetMotdDate

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdDate() time.Time`

GetMotdDate returns the MotdDate field if non-nil, zero value otherwise.

### GetMotdDateOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMotdDateOk() (*time.Time, bool)`

GetMotdDateOk returns a tuple with the MotdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdDate

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMotdDate(v time.Time)`

SetMotdDate sets MotdDate field to given value.

### HasMotdDate

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMotdDate() bool`

HasMotdDate returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxNetworks() string`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxNetworksOk() (*string, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxNetworks(v string)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.

### GetMaxRouters

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxRouters() string`

GetMaxRouters returns the MaxRouters field if non-nil, zero value otherwise.

### GetMaxRoutersOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetMaxRoutersOk() (*string, bool)`

GetMaxRoutersOk returns a tuple with the MaxRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRouters

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetMaxRouters(v string)`

SetMaxRouters sets MaxRouters field to given value.

### HasMaxRouters

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasMaxRouters() bool`

HasMaxRouters returns a boolean if a field has been set.

### GetShutdownType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.

### HasShutdownType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownType() bool`

HasShutdownType returns a boolean if a field has been set.

### GetShutdownAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetStorageServerId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetStorageServerId() string`

GetStorageServerId returns the StorageServerId field if non-nil, zero value otherwise.

### GetStorageServerIdOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetStorageServerIdOk() (*string, bool)`

GetStorageServerIdOk returns a tuple with the StorageServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetStorageServerId(v string)`

SetStorageServerId sets StorageServerId field to given value.

### HasStorageServerId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasStorageServerId() bool`

HasStorageServerId returns a boolean if a field has been set.

### GetStrict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetKey

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreateUserType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateUserType() string`

GetCreateUserType returns the CreateUserType field if non-nil, zero value otherwise.

### GetCreateUserTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateUserTypeOk() (*string, bool)`

GetCreateUserTypeOk returns a tuple with the CreateUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetCreateUserType(v string)`

SetCreateUserType sets CreateUserType field to given value.

### HasCreateUserType

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasCreateUserType() bool`

HasCreateUserType returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetWorkflowId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfig) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


