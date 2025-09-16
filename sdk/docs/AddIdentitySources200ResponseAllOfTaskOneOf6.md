# AddIdentitySources200ResponseAllOfTaskOneOf6

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
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DefaultAccountRole** | Pointer to [**ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole.md) |  | [optional] 
**Config** | Pointer to [**ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to [**ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddIdentitySources200ResponseAllOfTaskOneOf6

`func NewAddIdentitySources200ResponseAllOfTaskOneOf6() *AddIdentitySources200ResponseAllOfTaskOneOf6`

NewAddIdentitySources200ResponseAllOfTaskOneOf6 instantiates a new AddIdentitySources200ResponseAllOfTaskOneOf6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdentitySources200ResponseAllOfTaskOneOf6WithDefaults

`func NewAddIdentitySources200ResponseAllOfTaskOneOf6WithDefaults() *AddIdentitySources200ResponseAllOfTaskOneOf6`

NewAddIdentitySources200ResponseAllOfTaskOneOf6WithDefaults instantiates a new AddIdentitySources200ResponseAllOfTaskOneOf6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDefaultAccountRole() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDefaultAccountRoleOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetDefaultAccountRole(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetConfig() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetConfigOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetConfig(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetRoleMappings() []ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetRoleMappingsOk() (*[]ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetRoleMappings(v []ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5RoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetProviderSettings() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetProviderSettingsOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetProviderSettings(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddIdentitySources200ResponseAllOfTaskOneOf6) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


