# UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5

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
**Account** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Account**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Account.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5DefaultAccountRole**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5DefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Config**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Config.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5RoleMappingsInner**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5RoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5ProviderSettings**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5ProviderSettings.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5

`func NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5() *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5`

NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5 instantiates a new UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5WithDefaults

`func NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5WithDefaults() *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5`

NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5WithDefaults instantiates a new UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetAccount() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetAccountOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetAccount(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDefaultAccountRole() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5DefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDefaultAccountRoleOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5DefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetDefaultAccountRole(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5DefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetConfig() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetConfigOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetConfig(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetRoleMappings() []UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5RoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetRoleMappingsOk() (*[]UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5RoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetRoleMappings(v []UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5RoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetProviderSettings() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5ProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetProviderSettingsOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5ProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetProviderSettings(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5ProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf5) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


