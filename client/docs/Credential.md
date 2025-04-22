# Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Integration** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**AuthKey** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**AuthPath** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**User** | Pointer to [**AddCredentials200ResponseAllOfCredentialUser**](AddCredentials200ResponseAllOfCredentialUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to [**AddCredentials200ResponseAllOfCredentialConfig**](AddCredentials200ResponseAllOfCredentialConfig.md) |  | [optional] 

## Methods

### NewCredential

`func NewCredential() *Credential`

NewCredential instantiates a new Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialWithDefaults

`func NewCredentialWithDefaults() *Credential`

NewCredentialWithDefaults instantiates a new Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Credential) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credential) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credential) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Credential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Credential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Credential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Credential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Credential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Credential) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Credential) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Credential) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *Credential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *Credential) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *Credential) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *Credential) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *Credential) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetDescription

`func (o *Credential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Credential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Credential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Credential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsername

`func (o *Credential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Credential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Credential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Credential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *Credential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Credential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Credential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Credential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *Credential) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *Credential) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *Credential) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *Credential) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetAuthKey

`func (o *Credential) GetAuthKey() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *Credential) GetAuthKeyOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *Credential) SetAuthKey(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *Credential) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPath

`func (o *Credential) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *Credential) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *Credential) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.

### HasAuthPath

`func (o *Credential) HasAuthPath() bool`

HasAuthPath returns a boolean if a field has been set.

### GetExternalId

`func (o *Credential) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Credential) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Credential) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Credential) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRefType

`func (o *Credential) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Credential) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Credential) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Credential) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *Credential) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Credential) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Credential) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Credential) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *Credential) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Credential) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Credential) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Credential) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScope

`func (o *Credential) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Credential) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Credential) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Credential) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *Credential) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Credential) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Credential) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Credential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Credential) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Credential) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Credential) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Credential) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *Credential) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *Credential) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *Credential) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *Credential) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetEnabled

`func (o *Credential) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Credential) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Credential) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Credential) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *Credential) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Credential) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Credential) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Credential) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetUser

`func (o *Credential) GetUser() AddCredentials200ResponseAllOfCredentialUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Credential) GetUserOk() (*AddCredentials200ResponseAllOfCredentialUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Credential) SetUser(v AddCredentials200ResponseAllOfCredentialUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Credential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDateCreated

`func (o *Credential) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Credential) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Credential) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Credential) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Credential) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Credential) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Credential) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Credential) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *Credential) GetConfig() AddCredentials200ResponseAllOfCredentialConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Credential) GetConfigOk() (*AddCredentials200ResponseAllOfCredentialConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Credential) SetConfig(v AddCredentials200ResponseAllOfCredentialConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Credential) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


