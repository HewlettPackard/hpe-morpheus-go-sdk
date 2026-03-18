# GetIdentitySources200ResponseUserSourceAnyOf

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
**Account** | Pointer to [**GetIdentitySources200ResponseUserSourceAnyOfAccount**](GetIdentitySources200ResponseUserSourceAnyOfAccount.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**GetIdentitySources200ResponseUserSourceAnyOfDefaultAccountRole**](GetIdentitySources200ResponseUserSourceAnyOfDefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**GetIdentitySources200ResponseUserSourceAnyOfConfig**](GetIdentitySources200ResponseUserSourceAnyOfConfig.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]GetIdentitySources200ResponseUserSourceAnyOfRoleMappingsInner**](GetIdentitySources200ResponseUserSourceAnyOfRoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetIdentitySources200ResponseUserSourceAnyOf

`func NewGetIdentitySources200ResponseUserSourceAnyOf() *GetIdentitySources200ResponseUserSourceAnyOf`

NewGetIdentitySources200ResponseUserSourceAnyOf instantiates a new GetIdentitySources200ResponseUserSourceAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentitySources200ResponseUserSourceAnyOfWithDefaults

`func NewGetIdentitySources200ResponseUserSourceAnyOfWithDefaults() *GetIdentitySources200ResponseUserSourceAnyOf`

NewGetIdentitySources200ResponseUserSourceAnyOfWithDefaults instantiates a new GetIdentitySources200ResponseUserSourceAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetAccount() GetIdentitySources200ResponseUserSourceAnyOfAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetAccountOk() (*GetIdentitySources200ResponseUserSourceAnyOfAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetAccount(v GetIdentitySources200ResponseUserSourceAnyOfAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDefaultAccountRole() GetIdentitySources200ResponseUserSourceAnyOfDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDefaultAccountRoleOk() (*GetIdentitySources200ResponseUserSourceAnyOfDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetDefaultAccountRole(v GetIdentitySources200ResponseUserSourceAnyOfDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetConfig() GetIdentitySources200ResponseUserSourceAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetConfigOk() (*GetIdentitySources200ResponseUserSourceAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetConfig(v GetIdentitySources200ResponseUserSourceAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetRoleMappings() []GetIdentitySources200ResponseUserSourceAnyOfRoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetRoleMappingsOk() (*[]GetIdentitySources200ResponseUserSourceAnyOfRoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetRoleMappings(v []GetIdentitySources200ResponseUserSourceAnyOfRoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetIdentitySources200ResponseUserSourceAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


