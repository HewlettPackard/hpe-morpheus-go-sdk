# ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf

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
**Account** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfAccount**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfAccount.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfConfig**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfConfig.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner**](ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf

`func NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf() *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf`

NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf instantiates a new ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfWithDefaults

`func NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfWithDefaults() *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf`

NewListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfWithDefaults instantiates a new ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetAccount() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetAccountOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetAccount(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDefaultAccountRole() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDefaultAccountRoleOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetDefaultAccountRole(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetConfig() ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetConfigOk() (*ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetConfig(v ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetRoleMappings() []ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetRoleMappingsOk() (*[]ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetRoleMappings(v []ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTenantSubtenantIdentitySources200ResponseAllOfUserSourcesInnerAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


