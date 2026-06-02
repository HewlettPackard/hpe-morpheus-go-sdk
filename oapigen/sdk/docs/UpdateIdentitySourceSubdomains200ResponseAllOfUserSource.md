# UpdateIdentitySourceSubdomains200ResponseAllOfUserSource

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
**Account** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Account**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Account.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7DefaultAccountRole**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7DefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Config**](UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Config.md) |  | [optional] 
**RoleMappings** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSource

`func NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSource() *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource`

NewUpdateIdentitySourceSubdomains200ResponseAllOfUserSource instantiates a new UpdateIdentitySourceSubdomains200ResponseAllOfUserSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetAccount() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetAccountOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetAccount(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDefaultAccountRole() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7DefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDefaultAccountRoleOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7DefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetDefaultAccountRole(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7DefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetConfig() UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetConfigOk() (*UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetConfig(v UpdateIdentitySourceSubdomains200ResponseAllOfUserSourceAnyOf7Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetRoleMappings() []map[string]interface{}`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetRoleMappingsOk() (*[]map[string]interface{}, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetRoleMappings(v []map[string]interface{})`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateIdentitySourceSubdomains200ResponseAllOfUserSource) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


