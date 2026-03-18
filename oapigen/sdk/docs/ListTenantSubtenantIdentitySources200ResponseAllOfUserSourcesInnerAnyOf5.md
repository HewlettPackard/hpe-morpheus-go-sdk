# ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5

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
**Account** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Account**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Account.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5DefaultAccountRole**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5DefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Config**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Config.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5

`func NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5() *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5`

NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5 instantiates a new ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5WithDefaults

`func NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5WithDefaults() *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5`

NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5WithDefaults instantiates a new ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetAccount() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetAccountOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetAccount(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDefaultAccountRole() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5DefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDefaultAccountRoleOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5DefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetDefaultAccountRole(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5DefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetConfig() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetConfigOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetConfig(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetRoleMappings() []ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetRoleMappingsOk() (*[]ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetRoleMappings(v []ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetProviderSettings() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetProviderSettingsOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetProviderSettings(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


