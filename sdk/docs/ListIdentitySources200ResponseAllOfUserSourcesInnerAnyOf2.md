# ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2

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
**Config** | Pointer to [**ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2Config**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2Config.md) |  | [optional] 
**RoleMappings** | Pointer to [**[]ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf1RoleMappingsInner**](ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf1RoleMappingsInner.md) |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginURL** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2

`func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2`

NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2 instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2WithDefaults

`func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2WithDefaults() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2`

NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2WithDefaults instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDeleted

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetAutoSyncOnLogin

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetAutoSyncOnLogin() bool`

GetAutoSyncOnLogin returns the AutoSyncOnLogin field if non-nil, zero value otherwise.

### GetAutoSyncOnLoginOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetAutoSyncOnLoginOk() (*bool, bool)`

GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncOnLogin

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetAutoSyncOnLogin(v bool)`

SetAutoSyncOnLogin sets AutoSyncOnLogin field to given value.

### HasAutoSyncOnLogin

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasAutoSyncOnLogin() bool`

HasAutoSyncOnLogin returns a boolean if a field has been set.

### GetExternalLogin

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetExternalLogin() bool`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetExternalLoginOk() (*bool, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetExternalLogin(v bool)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetAccount

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDefaultAccountRole() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDefaultAccountRoleOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetDefaultAccountRole(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetConfig

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetConfig() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetConfigOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetConfig(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRoleMappings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetRoleMappings() []ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf1RoleMappingsInner`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetRoleMappingsOk() (*[]ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf1RoleMappingsInner, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetRoleMappings(v []ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf1RoleMappingsInner)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetSubdomain

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginURL

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetLoginURL() string`

GetLoginURL returns the LoginURL field if non-nil, zero value otherwise.

### GetLoginURLOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetLoginURLOk() (*string, bool)`

GetLoginURLOk returns a tuple with the LoginURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginURL

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetLoginURL(v string)`

SetLoginURL sets LoginURL field to given value.

### HasLoginURL

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasLoginURL() bool`

HasLoginURL returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf2) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


