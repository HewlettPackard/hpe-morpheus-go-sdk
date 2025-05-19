# ListPolicies200ResponseAllOfPoliciesInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | Pointer to **string** |  | [optional] 
**CreateBackupType** | **string** |  | 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**BackupStorageIds** | **[]string** |  | 
**MaxPrice** | **string** |  | 
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
**MaxStorage** | **string** |  | 
**MaxVirtualServers** | **string** |  | 
**MaxVms** | **string** |  | 
**MotdTitle** | Pointer to **string** |  | [optional] 
**Motd** | [**ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf19Motd**](ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf19Motd.md) |  | 
**MotdMessage** | Pointer to **string** |  | [optional] 
**MotdType** | Pointer to **string** |  | [optional] 
**MotdFullPage** | Pointer to **string** |  | [optional] 
**MotdDate** | Pointer to **time.Time** |  | [optional] 
**MaxNetworks** | **string** |  | 
**PowerScheduleType** | **string** |  | 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**PowerScheduleHideFixed** | Pointer to **bool** |  | [optional] 
**MaxRouters** | **string** |  | 
**ShutdownType** | **string** |  | 
**ShutdownAge** | Pointer to **string** |  | [optional] 
**ShutdownRenewal** | Pointer to **string** |  | [optional] 
**ShutdownNotify** | Pointer to **string** |  | [optional] 
**ShutdownMessage** | Pointer to **string** |  | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**ShutdownHideFixed** | Pointer to **bool** |  | [optional] 
**StorageServerId** | **string** |  | 
**Strict** | **bool** |  | 
**Key** | Pointer to **string** |  | [optional] 
**ValueListId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**CreateUserType** | **string** |  | 
**CreateUser** | Pointer to **bool** |  | [optional] 
**UserGroup** | **string** |  | 
**WorkflowId** | **string** |  | 

## Methods

### NewListPolicies200ResponseAllOfPoliciesInnerConfig

`func NewListPolicies200ResponseAllOfPoliciesInnerConfig(createBackupType string, backupStorageIds []string, maxPrice string, serverNamingType string, keyPattern string, removalAge string, lifecycleType string, hostNamingType string, namingType string, maxContainers string, maxCores string, maxHosts string, maxPools string, maxMemory string, maxPoolMembers string, maxStorage string, maxVirtualServers string, maxVms string, motd ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf19Motd, maxNetworks string, powerScheduleType string, maxRouters string, shutdownType string, storageServerId string, strict bool, createUserType string, userGroup string, workflowId string, ) *ListPolicies200ResponseAllOfPoliciesInnerConfig`

NewListPolicies200ResponseAllOfPoliciesInnerConfig instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicies200ResponseAllOfPoliciesInnerConfigWithDefaults

`func NewListPolicies200ResponseAllOfPoliciesInnerConfigWithDefaults() *ListPolicies200ResponseAllOfPoliciesInnerConfig`

NewListPolicies200ResponseAllOfPoliciesInnerConfigWithDefaults instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetCreateBackupType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.


### GetCreateBackup

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackupStorageIds

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetBackupStorageIds() []string`

GetBackupStorageIds returns the BackupStorageIds field if non-nil, zero value otherwise.

### GetBackupStorageIdsOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetBackupStorageIdsOk() (*[]string, bool)`

GetBackupStorageIdsOk returns a tuple with the BackupStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageIds

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetBackupStorageIds(v []string)`

SetBackupStorageIds sets BackupStorageIds field to given value.


### GetMaxPrice

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.


### GetMaxPriceCurrency

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPriceCurrency() string`

GetMaxPriceCurrency returns the MaxPriceCurrency field if non-nil, zero value otherwise.

### GetMaxPriceCurrencyOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPriceCurrencyOk() (*string, bool)`

GetMaxPriceCurrencyOk returns a tuple with the MaxPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceCurrency

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxPriceCurrency(v string)`

SetMaxPriceCurrency sets MaxPriceCurrency field to given value.

### HasMaxPriceCurrency

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMaxPriceCurrency() bool`

HasMaxPriceCurrency returns a boolean if a field has been set.

### GetMaxPriceUnit

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPriceUnit() string`

GetMaxPriceUnit returns the MaxPriceUnit field if non-nil, zero value otherwise.

### GetMaxPriceUnitOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPriceUnitOk() (*string, bool)`

GetMaxPriceUnitOk returns a tuple with the MaxPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceUnit

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxPriceUnit(v string)`

SetMaxPriceUnit sets MaxPriceUnit field to given value.

### HasMaxPriceUnit

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMaxPriceUnit() bool`

HasMaxPriceUnit returns a boolean if a field has been set.

### GetServerNamingType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.


### GetServerNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.

### GetKeyPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.


### GetRead

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRemovalAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.


### GetLifecycleType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.


### GetLifecycleAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.

### GetHostNamingType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.


### GetHostNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.

### GetNamingType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.


### GetNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.

### GetMaxContainers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxContainers() string`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxContainersOk() (*string, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxContainers(v string)`

SetMaxContainers sets MaxContainers field to given value.


### GetMaxCores

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.


### GetExcludeContainers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxHosts() string`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxHostsOk() (*string, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxHosts(v string)`

SetMaxHosts sets MaxHosts field to given value.


### GetMaxPools

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPools() string`

GetMaxPools returns the MaxPools field if non-nil, zero value otherwise.

### GetMaxPoolsOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPoolsOk() (*string, bool)`

GetMaxPoolsOk returns a tuple with the MaxPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPools

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxPools(v string)`

SetMaxPools sets MaxPools field to given value.


### GetMaxMemory

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxPoolMembers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPoolMembers() string`

GetMaxPoolMembers returns the MaxPoolMembers field if non-nil, zero value otherwise.

### GetMaxPoolMembersOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxPoolMembersOk() (*string, bool)`

GetMaxPoolMembersOk returns a tuple with the MaxPoolMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolMembers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxPoolMembers(v string)`

SetMaxPoolMembers sets MaxPoolMembers field to given value.


### GetMaxStorage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxVirtualServers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxVirtualServers() string`

GetMaxVirtualServers returns the MaxVirtualServers field if non-nil, zero value otherwise.

### GetMaxVirtualServersOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxVirtualServersOk() (*string, bool)`

GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualServers

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxVirtualServers(v string)`

SetMaxVirtualServers sets MaxVirtualServers field to given value.


### GetMaxVms

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxVms() string`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxVmsOk() (*string, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxVms(v string)`

SetMaxVms sets MaxVms field to given value.


### GetMotdTitle

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdTitle() string`

GetMotdTitle returns the MotdTitle field if non-nil, zero value otherwise.

### GetMotdTitleOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdTitleOk() (*string, bool)`

GetMotdTitleOk returns a tuple with the MotdTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdTitle

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMotdTitle(v string)`

SetMotdTitle sets MotdTitle field to given value.

### HasMotdTitle

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMotdTitle() bool`

HasMotdTitle returns a boolean if a field has been set.

### GetMotd

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotd() ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf19Motd`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdOk() (*ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf19Motd, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMotd(v ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf19Motd)`

SetMotd sets Motd field to given value.


### GetMotdMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdMessage() string`

GetMotdMessage returns the MotdMessage field if non-nil, zero value otherwise.

### GetMotdMessageOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdMessageOk() (*string, bool)`

GetMotdMessageOk returns a tuple with the MotdMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMotdMessage(v string)`

SetMotdMessage sets MotdMessage field to given value.

### HasMotdMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMotdMessage() bool`

HasMotdMessage returns a boolean if a field has been set.

### GetMotdType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdType() string`

GetMotdType returns the MotdType field if non-nil, zero value otherwise.

### GetMotdTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdTypeOk() (*string, bool)`

GetMotdTypeOk returns a tuple with the MotdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMotdType(v string)`

SetMotdType sets MotdType field to given value.

### HasMotdType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMotdType() bool`

HasMotdType returns a boolean if a field has been set.

### GetMotdFullPage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdFullPage() string`

GetMotdFullPage returns the MotdFullPage field if non-nil, zero value otherwise.

### GetMotdFullPageOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdFullPageOk() (*string, bool)`

GetMotdFullPageOk returns a tuple with the MotdFullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdFullPage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMotdFullPage(v string)`

SetMotdFullPage sets MotdFullPage field to given value.

### HasMotdFullPage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMotdFullPage() bool`

HasMotdFullPage returns a boolean if a field has been set.

### GetMotdDate

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdDate() time.Time`

GetMotdDate returns the MotdDate field if non-nil, zero value otherwise.

### GetMotdDateOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMotdDateOk() (*time.Time, bool)`

GetMotdDateOk returns a tuple with the MotdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdDate

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMotdDate(v time.Time)`

SetMotdDate sets MotdDate field to given value.

### HasMotdDate

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasMotdDate() bool`

HasMotdDate returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxNetworks() string`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxNetworksOk() (*string, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxNetworks(v string)`

SetMaxNetworks sets MaxNetworks field to given value.


### GetPowerScheduleType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.


### GetPowerSchedule

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.

### GetMaxRouters

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxRouters() string`

GetMaxRouters returns the MaxRouters field if non-nil, zero value otherwise.

### GetMaxRoutersOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetMaxRoutersOk() (*string, bool)`

GetMaxRoutersOk returns a tuple with the MaxRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRouters

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetMaxRouters(v string)`

SetMaxRouters sets MaxRouters field to given value.


### GetShutdownType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.


### GetShutdownAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetStorageServerId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetStorageServerId() string`

GetStorageServerId returns the StorageServerId field if non-nil, zero value otherwise.

### GetStorageServerIdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetStorageServerIdOk() (*string, bool)`

GetStorageServerIdOk returns a tuple with the StorageServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetStorageServerId(v string)`

SetStorageServerId sets StorageServerId field to given value.


### GetStrict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetStrict(v bool)`

SetStrict sets Strict field to given value.


### GetKey

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreateUserType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateUserType() string`

GetCreateUserType returns the CreateUserType field if non-nil, zero value otherwise.

### GetCreateUserTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateUserTypeOk() (*string, bool)`

GetCreateUserTypeOk returns a tuple with the CreateUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserType

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetCreateUserType(v string)`

SetCreateUserType sets CreateUserType field to given value.


### GetCreateUser

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.


### GetWorkflowId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ListPolicies200ResponseAllOfPoliciesInnerConfig) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


