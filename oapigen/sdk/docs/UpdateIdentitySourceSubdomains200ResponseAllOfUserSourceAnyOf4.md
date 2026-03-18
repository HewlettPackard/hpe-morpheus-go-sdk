# UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4

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
**Account** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Account**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Account.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4DefaultAccountRole**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4DefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Config**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Config.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4RoleMappingsInner**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4RoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4

`func NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4() *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4`

NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4 instantiates a new UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4WithDefaults

`func NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4WithDefaults() *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4`

NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4WithDefaults instantiates a new UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetAccount() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetAccountOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetAccount(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDefaultAccountRole() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4DefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDefaultAccountRoleOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4DefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetDefaultAccountRole(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4DefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetConfig() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetConfigOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetConfig(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetRoleMappings() []UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4RoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetRoleMappingsOk() (*[]UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4RoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetRoleMappings(v []UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4RoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf4) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


