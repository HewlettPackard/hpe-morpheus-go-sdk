# AddPoliciesCloudRequestPolicyPolicyTypeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | Pointer to **string** | ID of your ServiceNow or approval integration | [optional] 
**WorkflowId** | Pointer to **string** | ID of the workflow to execute | [optional] 
**FlowId** | Pointer to **string** | ID of ServiceNow Flow (set if workflowType is &#39;flow&#39;) | [optional] 
**WorkflowType** | Pointer to **string** | Options: \&quot;workflow\&quot; (legacy workflow), \&quot;flow\&quot; (ServiceNow Flow) | [optional] 
**CreateBackupType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**CreateBackup** | Pointer to **bool** | Enforce backup creation | [optional] 
**BackupStorageIds** | Pointer to **[]int64** | Array of backup storage IDs to restrict available backup targets | [optional] 
**MaxPrice** | Pointer to **float32** | Maximum price limit | [optional] 
**MaxPriceCurrency** | Pointer to **string** | Currency code (e.g., USD) | [optional] 
**MaxPriceUnit** | Pointer to **string** | Options: \&quot;hour\&quot;, \&quot;month\&quot; | [optional] 
**ServerNamingType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**ServerNamingPattern** | Pointer to **string** | Name pattern uses ${variable} string interpolation. Available variables are: groupName, groupCode, cloudName, cloudCode, type, accountId, account, accountType, platform, username, userId, userInitials, provisionType | [optional] 
**ServerNamingConflict** | Pointer to **bool** | Auto-resolve conflicts | [optional] 
**KeyPattern** | Pointer to **string** | Pattern to match Cypher keys (e.g., \&quot;secret/_*\&quot;, \&quot;password/_*\&quot;) | [optional] 
**Read** | Pointer to **bool** | Allow read access | [optional] 
**Write** | Pointer to **bool** | Allow write access | [optional] 
**Update** | Pointer to **bool** | Allow update access | [optional] 
**Delete** | Pointer to **bool** | Allow delete access | [optional] 
**List** | Pointer to **bool** | Allow list access | [optional] 
**RemovalAge** | Pointer to **string** | Number of days to delay deletion | [optional] 
**LifecycleType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (fixed expiration) | [optional] 
**LifecycleAge** | Pointer to **string** | Days until expiration | [optional] 
**LifecycleRenewal** | Pointer to **string** | Days for renewal window | [optional] 
**LifecycleNotify** | Pointer to **string** | Days before expiration to notify | [optional] 
**LifecycleMessage** | Pointer to **string** | Notification message | [optional] 
**LifecycleAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleExtensionsBeforeApproval** | Pointer to **string** | Number of extensions before requiring approval | [optional] 
**LifecycleHideFixed** | Pointer to **bool** | Hide fixed expiration from users | [optional] 
**LifecycleWorkflowId** | Pointer to **string** | ID of legacy ServiceNow workflow (set if workflowType is &#39;workflow&#39;) | [optional] 
**MaxStorage** | Pointer to **string** | Max Storage (GB) | [optional] 
**HostNamingType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**HostNamingPattern** | Pointer to **string** | Name pattern uses ${variable} string interpolation. Available variables are: groupName, groupCode, cloudName, cloudCode, type, accountId, account, accountType, platform, username, userId, userInitials, provisionType | [optional] 
**NamingType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**NamingPattern** | Pointer to **string** | Name pattern uses ${variable} string interpolation. Available variables are: groupName, groupCode, cloudName, cloudCode, type, accountId, account, accountType, platform, username, userId, userInitials, provisionType | [optional] 
**NamingConflict** | Pointer to **bool** | Auto-resolve conflicts | [optional] 
**MaxContainers** | Pointer to **string** | Max Containers | [optional] 
**MaxCores** | Pointer to **string** | Max Cores | [optional] 
**ExcludeContainers** | Pointer to **string** |  | [optional] [default to "off"]
**MaxHosts** | Pointer to **string** | Max Hosts | [optional] 
**MaxPools** | Pointer to **string** | Max Pools | [optional] 
**MaxMemory** | Pointer to **string** | Max Memory (GB) | [optional] 
**MaxPoolMembers** | Pointer to **string** | Max Pool Members | [optional] 
**MaxSnapshots** | Pointer to **string** | Max Snapshots | [optional] 
**MaxVirtualServers** | Pointer to **string** | Max Virtual Servers | [optional] 
**MaxVms** | Pointer to **string** | Max VMs | [optional] 
**MotdTitle** | Pointer to **string** | Message title | [optional] 
**Motd** | Pointer to [**MessageOfTheDayPolicyTypeConfiguration1Motd**](MessageOfTheDayPolicyTypeConfiguration1Motd.md) |  | [optional] 
**MotdMessage** | Pointer to **string** | Message content | [optional] 
**MotdType** | Pointer to **string** | Options: \&quot;info\&quot;, \&quot;warning\&quot;, \&quot;critical\&quot; | [optional] 
**MotdFullPage** | Pointer to **bool** | Display full page | [optional] 
**MotdDate** | Pointer to **string** | Display date for message | [optional] 
**MaxNetworks** | Pointer to **string** | Max Networks | [optional] 
**PowerScheduleType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict schedule) | [optional] 
**PowerSchedule** | Pointer to **string** | ID of the power schedule | [optional] 
**PowerScheduleHideFixed** | Pointer to **bool** |  | [optional] 
**MaxRouters** | Pointer to **string** | Max Routers | [optional] 
**RequiredNetworks** | Pointer to **[]int64** | Array of required network IDs | [optional] 
**ShutdownType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict shutdown) | [optional] 
**ShutdownAge** | Pointer to **string** | Shutdown Age (days) | [optional] 
**ShutdownRenewal** | Pointer to **string** | If the instance is renewed, this is the number of day increments the shutdown date is increased by | [optional] 
**ShutdownNotify** | Pointer to **string** | Days before shutdown to notify via email | [optional] 
**ShutdownMessage** | Pointer to **string** | Notification message | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** | Number of extensions before requiring approval | [optional] 
**ShutdownHideFixed** | Pointer to **bool** | Hide fixed shutdown from users | [optional] 
**ShutdownWorkflowId** | Pointer to **string** | ID of legacy ServiceNow workflow (set if workflowType is &#39;workflow&#39;) | [optional] 
**StorageServerId** | Pointer to **string** | ID of the storage server | [optional] 
**Strict** | Pointer to **bool** | Strict enforcement | [optional] 
**Key** | Pointer to **string** | Tag key to enforce | [optional] 
**ValueListId** | Pointer to **string** | ID of value from value list (optional) | [optional] 
**Value** | Pointer to **string** | Tag value (optional, can be left empty for any value) | [optional] 
**CreateUserType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; | [optional] 
**CreateUser** | Pointer to **bool** | Enforce user creation | [optional] 
**UserGroup** | Pointer to **string** | ID of the user group to assign | [optional] 

## Methods

### NewAddPoliciesCloudRequestPolicyPolicyTypeConfig

`func NewAddPoliciesCloudRequestPolicyPolicyTypeConfig() *AddPoliciesCloudRequestPolicyPolicyTypeConfig`

NewAddPoliciesCloudRequestPolicyPolicyTypeConfig instantiates a new AddPoliciesCloudRequestPolicyPolicyTypeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPoliciesCloudRequestPolicyPolicyTypeConfigWithDefaults

`func NewAddPoliciesCloudRequestPolicyPolicyTypeConfigWithDefaults() *AddPoliciesCloudRequestPolicyPolicyTypeConfig`

NewAddPoliciesCloudRequestPolicyPolicyTypeConfigWithDefaults instantiates a new AddPoliciesCloudRequestPolicyPolicyTypeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetFlowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetCreateBackupType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.

### HasCreateBackupType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasCreateBackupType() bool`

HasCreateBackupType returns a boolean if a field has been set.

### GetCreateBackup

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackupStorageIds

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetBackupStorageIds() []int64`

GetBackupStorageIds returns the BackupStorageIds field if non-nil, zero value otherwise.

### GetBackupStorageIdsOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetBackupStorageIdsOk() (*[]int64, bool)`

GetBackupStorageIdsOk returns a tuple with the BackupStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStorageIds

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetBackupStorageIds(v []int64)`

SetBackupStorageIds sets BackupStorageIds field to given value.

### HasBackupStorageIds

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasBackupStorageIds() bool`

HasBackupStorageIds returns a boolean if a field has been set.

### GetMaxPrice

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetMaxPriceCurrency

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPriceCurrency() string`

GetMaxPriceCurrency returns the MaxPriceCurrency field if non-nil, zero value otherwise.

### GetMaxPriceCurrencyOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPriceCurrencyOk() (*string, bool)`

GetMaxPriceCurrencyOk returns a tuple with the MaxPriceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceCurrency

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxPriceCurrency(v string)`

SetMaxPriceCurrency sets MaxPriceCurrency field to given value.

### HasMaxPriceCurrency

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxPriceCurrency() bool`

HasMaxPriceCurrency returns a boolean if a field has been set.

### GetMaxPriceUnit

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPriceUnit() string`

GetMaxPriceUnit returns the MaxPriceUnit field if non-nil, zero value otherwise.

### GetMaxPriceUnitOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPriceUnitOk() (*string, bool)`

GetMaxPriceUnitOk returns a tuple with the MaxPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriceUnit

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxPriceUnit(v string)`

SetMaxPriceUnit sets MaxPriceUnit field to given value.

### HasMaxPriceUnit

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxPriceUnit() bool`

HasMaxPriceUnit returns a boolean if a field has been set.

### GetServerNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.

### HasServerNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasServerNamingType() bool`

HasServerNamingType returns a boolean if a field has been set.

### GetServerNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.

### GetKeyPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.

### HasKeyPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasKeyPattern() bool`

HasKeyPattern returns a boolean if a field has been set.

### GetRead

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRemovalAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.

### HasRemovalAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasRemovalAge() bool`

HasRemovalAge returns a boolean if a field has been set.

### GetLifecycleType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.

### HasLifecycleType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleType() bool`

HasLifecycleType returns a boolean if a field has been set.

### GetLifecycleAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.

### GetLifecycleWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleWorkflowId() string`

GetLifecycleWorkflowId returns the LifecycleWorkflowId field if non-nil, zero value otherwise.

### GetLifecycleWorkflowIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetLifecycleWorkflowIdOk() (*string, bool)`

GetLifecycleWorkflowIdOk returns a tuple with the LifecycleWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetLifecycleWorkflowId(v string)`

SetLifecycleWorkflowId sets LifecycleWorkflowId field to given value.

### HasLifecycleWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasLifecycleWorkflowId() bool`

HasLifecycleWorkflowId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetHostNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.

### HasHostNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasHostNamingType() bool`

HasHostNamingType returns a boolean if a field has been set.

### GetHostNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.

### GetNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.

### HasNamingType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasNamingType() bool`

HasNamingType returns a boolean if a field has been set.

### GetNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.

### GetMaxContainers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxContainers() string`

GetMaxContainers returns the MaxContainers field if non-nil, zero value otherwise.

### GetMaxContainersOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxContainersOk() (*string, bool)`

GetMaxContainersOk returns a tuple with the MaxContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxContainers(v string)`

SetMaxContainers sets MaxContainers field to given value.

### HasMaxContainers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxContainers() bool`

HasMaxContainers returns a boolean if a field has been set.

### GetMaxCores

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetExcludeContainers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxHosts() string`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxHostsOk() (*string, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxHosts(v string)`

SetMaxHosts sets MaxHosts field to given value.

### HasMaxHosts

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxHosts() bool`

HasMaxHosts returns a boolean if a field has been set.

### GetMaxPools

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPools() string`

GetMaxPools returns the MaxPools field if non-nil, zero value otherwise.

### GetMaxPoolsOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPoolsOk() (*string, bool)`

GetMaxPoolsOk returns a tuple with the MaxPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPools

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxPools(v string)`

SetMaxPools sets MaxPools field to given value.

### HasMaxPools

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxPools() bool`

HasMaxPools returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxPoolMembers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPoolMembers() string`

GetMaxPoolMembers returns the MaxPoolMembers field if non-nil, zero value otherwise.

### GetMaxPoolMembersOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxPoolMembersOk() (*string, bool)`

GetMaxPoolMembersOk returns a tuple with the MaxPoolMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolMembers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxPoolMembers(v string)`

SetMaxPoolMembers sets MaxPoolMembers field to given value.

### HasMaxPoolMembers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxPoolMembers() bool`

HasMaxPoolMembers returns a boolean if a field has been set.

### GetMaxSnapshots

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxSnapshots() string`

GetMaxSnapshots returns the MaxSnapshots field if non-nil, zero value otherwise.

### GetMaxSnapshotsOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxSnapshotsOk() (*string, bool)`

GetMaxSnapshotsOk returns a tuple with the MaxSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshots

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxSnapshots(v string)`

SetMaxSnapshots sets MaxSnapshots field to given value.

### HasMaxSnapshots

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxSnapshots() bool`

HasMaxSnapshots returns a boolean if a field has been set.

### GetMaxVirtualServers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxVirtualServers() string`

GetMaxVirtualServers returns the MaxVirtualServers field if non-nil, zero value otherwise.

### GetMaxVirtualServersOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxVirtualServersOk() (*string, bool)`

GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVirtualServers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxVirtualServers(v string)`

SetMaxVirtualServers sets MaxVirtualServers field to given value.

### HasMaxVirtualServers

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxVirtualServers() bool`

HasMaxVirtualServers returns a boolean if a field has been set.

### GetMaxVms

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxVms() string`

GetMaxVms returns the MaxVms field if non-nil, zero value otherwise.

### GetMaxVmsOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxVmsOk() (*string, bool)`

GetMaxVmsOk returns a tuple with the MaxVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVms

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxVms(v string)`

SetMaxVms sets MaxVms field to given value.

### HasMaxVms

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxVms() bool`

HasMaxVms returns a boolean if a field has been set.

### GetMotdTitle

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdTitle() string`

GetMotdTitle returns the MotdTitle field if non-nil, zero value otherwise.

### GetMotdTitleOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdTitleOk() (*string, bool)`

GetMotdTitleOk returns a tuple with the MotdTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdTitle

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMotdTitle(v string)`

SetMotdTitle sets MotdTitle field to given value.

### HasMotdTitle

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMotdTitle() bool`

HasMotdTitle returns a boolean if a field has been set.

### GetMotd

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotd() MessageOfTheDayPolicyTypeConfiguration1Motd`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdOk() (*MessageOfTheDayPolicyTypeConfiguration1Motd, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMotd(v MessageOfTheDayPolicyTypeConfiguration1Motd)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMotd() bool`

HasMotd returns a boolean if a field has been set.

### GetMotdMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdMessage() string`

GetMotdMessage returns the MotdMessage field if non-nil, zero value otherwise.

### GetMotdMessageOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdMessageOk() (*string, bool)`

GetMotdMessageOk returns a tuple with the MotdMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMotdMessage(v string)`

SetMotdMessage sets MotdMessage field to given value.

### HasMotdMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMotdMessage() bool`

HasMotdMessage returns a boolean if a field has been set.

### GetMotdType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdType() string`

GetMotdType returns the MotdType field if non-nil, zero value otherwise.

### GetMotdTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdTypeOk() (*string, bool)`

GetMotdTypeOk returns a tuple with the MotdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMotdType(v string)`

SetMotdType sets MotdType field to given value.

### HasMotdType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMotdType() bool`

HasMotdType returns a boolean if a field has been set.

### GetMotdFullPage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdFullPage() bool`

GetMotdFullPage returns the MotdFullPage field if non-nil, zero value otherwise.

### GetMotdFullPageOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdFullPageOk() (*bool, bool)`

GetMotdFullPageOk returns a tuple with the MotdFullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdFullPage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMotdFullPage(v bool)`

SetMotdFullPage sets MotdFullPage field to given value.

### HasMotdFullPage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMotdFullPage() bool`

HasMotdFullPage returns a boolean if a field has been set.

### GetMotdDate

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdDate() string`

GetMotdDate returns the MotdDate field if non-nil, zero value otherwise.

### GetMotdDateOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMotdDateOk() (*string, bool)`

GetMotdDateOk returns a tuple with the MotdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotdDate

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMotdDate(v string)`

SetMotdDate sets MotdDate field to given value.

### HasMotdDate

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMotdDate() bool`

HasMotdDate returns a boolean if a field has been set.

### GetMaxNetworks

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxNetworks() string`

GetMaxNetworks returns the MaxNetworks field if non-nil, zero value otherwise.

### GetMaxNetworksOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxNetworksOk() (*string, bool)`

GetMaxNetworksOk returns a tuple with the MaxNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworks

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxNetworks(v string)`

SetMaxNetworks sets MaxNetworks field to given value.

### HasMaxNetworks

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxNetworks() bool`

HasMaxNetworks returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.

### GetMaxRouters

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxRouters() string`

GetMaxRouters returns the MaxRouters field if non-nil, zero value otherwise.

### GetMaxRoutersOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetMaxRoutersOk() (*string, bool)`

GetMaxRoutersOk returns a tuple with the MaxRouters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRouters

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetMaxRouters(v string)`

SetMaxRouters sets MaxRouters field to given value.

### HasMaxRouters

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasMaxRouters() bool`

HasMaxRouters returns a boolean if a field has been set.

### GetRequiredNetworks

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetRequiredNetworks() []int64`

GetRequiredNetworks returns the RequiredNetworks field if non-nil, zero value otherwise.

### GetRequiredNetworksOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetRequiredNetworksOk() (*[]int64, bool)`

GetRequiredNetworksOk returns a tuple with the RequiredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNetworks

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetRequiredNetworks(v []int64)`

SetRequiredNetworks sets RequiredNetworks field to given value.

### HasRequiredNetworks

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasRequiredNetworks() bool`

HasRequiredNetworks returns a boolean if a field has been set.

### GetShutdownType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.

### HasShutdownType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownType() bool`

HasShutdownType returns a boolean if a field has been set.

### GetShutdownAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetShutdownWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownWorkflowId() string`

GetShutdownWorkflowId returns the ShutdownWorkflowId field if non-nil, zero value otherwise.

### GetShutdownWorkflowIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetShutdownWorkflowIdOk() (*string, bool)`

GetShutdownWorkflowIdOk returns a tuple with the ShutdownWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetShutdownWorkflowId(v string)`

SetShutdownWorkflowId sets ShutdownWorkflowId field to given value.

### HasShutdownWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasShutdownWorkflowId() bool`

HasShutdownWorkflowId returns a boolean if a field has been set.

### GetStorageServerId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetStorageServerId() string`

GetStorageServerId returns the StorageServerId field if non-nil, zero value otherwise.

### GetStorageServerIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetStorageServerIdOk() (*string, bool)`

GetStorageServerIdOk returns a tuple with the StorageServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetStorageServerId(v string)`

SetStorageServerId sets StorageServerId field to given value.

### HasStorageServerId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasStorageServerId() bool`

HasStorageServerId returns a boolean if a field has been set.

### GetStrict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetKey

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreateUserType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateUserType() string`

GetCreateUserType returns the CreateUserType field if non-nil, zero value otherwise.

### GetCreateUserTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateUserTypeOk() (*string, bool)`

GetCreateUserTypeOk returns a tuple with the CreateUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetCreateUserType(v string)`

SetCreateUserType sets CreateUserType field to given value.

### HasCreateUserType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasCreateUserType() bool`

HasCreateUserType returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


