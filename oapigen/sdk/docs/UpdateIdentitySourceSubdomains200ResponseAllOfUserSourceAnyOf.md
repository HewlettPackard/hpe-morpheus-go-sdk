# UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf

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
**Account** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfAccount**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfAccount.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfDefaultAccountRole**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfDefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfConfig**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfConfig.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfRoleMappingsInner**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfRoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf

`func NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf() *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf`

NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf instantiates a new UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetAccount() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetAccountOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetAccount(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDefaultAccountRole() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDefaultAccountRoleOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetDefaultAccountRole(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetConfig() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetConfigOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetConfig(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetRoleMappings() []UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfRoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetRoleMappingsOk() (*[]UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfRoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetRoleMappings(v []UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOfRoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


