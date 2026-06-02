# ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4

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
**Account** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Account**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Account.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4DefaultAccountRole**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4DefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Config**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Config.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4RoleMappingsInner**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4RoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4

`func NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4() *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4`

NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4 instantiates a new ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetAccount() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetAccountOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetAccount(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDefaultAccountRole() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4DefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDefaultAccountRoleOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4DefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetDefaultAccountRole(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4DefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetConfig() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetConfigOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetConfig(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetRoleMappings() []ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4RoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetRoleMappingsOk() (*[]ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4RoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetRoleMappings(v []ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4RoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf4) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


