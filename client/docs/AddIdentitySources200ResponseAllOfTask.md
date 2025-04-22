# AddIdentitySources200ResponseAllOfTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**AutoSyncOnLogin** | Pointer to **bool** |  | [optional] 
**ExternalLogin** | Pointer to **bool** |  | [optional] 
**AllowCustomMappings** | Pointer to **bool** |  | [optional] 
**ManualRoleAssignment** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf7Config**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf7Config.md) |  | [optional] 
**RoleMappings** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddIdentitySources200ResponseAllOfTask

`func NewAddIdentitySources200ResponseAllOfTask() *AddIdentitySources200ResponseAllOfTask`

NewAddIdentitySources200ResponseAllOfTask instantiates a new AddIdentitySources200ResponseAllOfTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdentitySources200ResponseAllOfTaskWithDefaults

`func NewAddIdentitySources200ResponseAllOfTaskWithDefaults() *AddIdentitySources200ResponseAllOfTask`

NewAddIdentitySources200ResponseAllOfTaskWithDefaults instantiates a new AddIdentitySources200ResponseAllOfTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIdentitySources200ResponseAllOfTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIdentitySources200ResponseAllOfTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIdentitySources200ResponseAllOfTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddIdentitySources200ResponseAllOfTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIdentitySources200ResponseAllOfTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIdentitySources200ResponseAllOfTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddIdentitySources200ResponseAllOfTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIdentitySources200ResponseAllOfTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIdentitySources200ResponseAllOfTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *AddIdentitySources200ResponseAllOfTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddIdentitySources200ResponseAllOfTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddIdentitySources200ResponseAllOfTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *AddIdentitySources200ResponseAllOfTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIdentitySources200ResponseAllOfTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddIdentitySources200ResponseAllOfTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *AddIdentitySources200ResponseAllOfTask) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddIdentitySources200ResponseAllOfTask) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddIdentitySources200ResponseAllOfTask) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *AddIdentitySources200ResponseAllOfTask) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *AddIdentitySources200ResponseAllOfTask) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *AddIdentitySources200ResponseAllOfTask) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *AddIdentitySources200ResponseAllOfTask) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *AddIdentitySources200ResponseAllOfTask) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *AddIdentitySources200ResponseAllOfTask) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *AddIdentitySources200ResponseAllOfTask) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *AddIdentitySources200ResponseAllOfTask) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *AddIdentitySources200ResponseAllOfTask) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *AddIdentitySources200ResponseAllOfTask) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *AddIdentitySources200ResponseAllOfTask) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *AddIdentitySources200ResponseAllOfTask) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *AddIdentitySources200ResponseAllOfTask) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *AddIdentitySources200ResponseAllOfTask) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *AddIdentitySources200ResponseAllOfTask) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *AddIdentitySources200ResponseAllOfTask) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddIdentitySources200ResponseAllOfTask) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddIdentitySources200ResponseAllOfTask) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *AddIdentitySources200ResponseAllOfTask) GetDefaultAccountRole() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetDefaultAccountRoleOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *AddIdentitySources200ResponseAllOfTask) SetDefaultAccountRole(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *AddIdentitySources200ResponseAllOfTask) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *AddIdentitySources200ResponseAllOfTask) GetConfig() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf7Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetConfigOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf7Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIdentitySources200ResponseAllOfTask) SetConfig(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf7Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIdentitySources200ResponseAllOfTask) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *AddIdentitySources200ResponseAllOfTask) GetRoleMappings() []map[string]interface{}`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetRoleMappingsOk() (*[]map[string]interface{}, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *AddIdentitySources200ResponseAllOfTask) SetRoleMappings(v []map[string]interface{})`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *AddIdentitySources200ResponseAllOfTask) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *AddIdentitySources200ResponseAllOfTask) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *AddIdentitySources200ResponseAllOfTask) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *AddIdentitySources200ResponseAllOfTask) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *AddIdentitySources200ResponseAllOfTask) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *AddIdentitySources200ResponseAllOfTask) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *AddIdentitySources200ResponseAllOfTask) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *AddIdentitySources200ResponseAllOfTask) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *AddIdentitySources200ResponseAllOfTask) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *AddIdentitySources200ResponseAllOfTask) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddIdentitySources200ResponseAllOfTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddIdentitySources200ResponseAllOfTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddIdentitySources200ResponseAllOfTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddIdentitySources200ResponseAllOfTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddIdentitySources200ResponseAllOfTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddIdentitySources200ResponseAllOfTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddIdentitySources200ResponseAllOfTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


