# AddCredentials200ResponseAllOfCredential

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

### NewAddCredentials200ResponseAllOfCredential

`func NewAddCredentials200ResponseAllOfCredential() *AddCredentials200ResponseAllOfCredential`

NewAddCredentials200ResponseAllOfCredential instantiates a new AddCredentials200ResponseAllOfCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCredentials200ResponseAllOfCredentialWithDefaults

`func NewAddCredentials200ResponseAllOfCredentialWithDefaults() *AddCredentials200ResponseAllOfCredential`

NewAddCredentials200ResponseAllOfCredentialWithDefaults instantiates a new AddCredentials200ResponseAllOfCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCredentials200ResponseAllOfCredential) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCredentials200ResponseAllOfCredential) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCredentials200ResponseAllOfCredential) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCredentials200ResponseAllOfCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddCredentials200ResponseAllOfCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCredentials200ResponseAllOfCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCredentials200ResponseAllOfCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCredentials200ResponseAllOfCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddCredentials200ResponseAllOfCredential) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCredentials200ResponseAllOfCredential) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCredentials200ResponseAllOfCredential) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *AddCredentials200ResponseAllOfCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *AddCredentials200ResponseAllOfCredential) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddCredentials200ResponseAllOfCredential) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddCredentials200ResponseAllOfCredential) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AddCredentials200ResponseAllOfCredential) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetDescription

`func (o *AddCredentials200ResponseAllOfCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCredentials200ResponseAllOfCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCredentials200ResponseAllOfCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCredentials200ResponseAllOfCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsername

`func (o *AddCredentials200ResponseAllOfCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddCredentials200ResponseAllOfCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddCredentials200ResponseAllOfCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddCredentials200ResponseAllOfCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddCredentials200ResponseAllOfCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddCredentials200ResponseAllOfCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddCredentials200ResponseAllOfCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddCredentials200ResponseAllOfCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AddCredentials200ResponseAllOfCredential) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AddCredentials200ResponseAllOfCredential) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AddCredentials200ResponseAllOfCredential) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AddCredentials200ResponseAllOfCredential) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetAuthKey

`func (o *AddCredentials200ResponseAllOfCredential) GetAuthKey() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *AddCredentials200ResponseAllOfCredential) GetAuthKeyOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *AddCredentials200ResponseAllOfCredential) SetAuthKey(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *AddCredentials200ResponseAllOfCredential) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPath

`func (o *AddCredentials200ResponseAllOfCredential) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *AddCredentials200ResponseAllOfCredential) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *AddCredentials200ResponseAllOfCredential) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.

### HasAuthPath

`func (o *AddCredentials200ResponseAllOfCredential) HasAuthPath() bool`

HasAuthPath returns a boolean if a field has been set.

### GetExternalId

`func (o *AddCredentials200ResponseAllOfCredential) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddCredentials200ResponseAllOfCredential) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddCredentials200ResponseAllOfCredential) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddCredentials200ResponseAllOfCredential) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRefType

`func (o *AddCredentials200ResponseAllOfCredential) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddCredentials200ResponseAllOfCredential) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddCredentials200ResponseAllOfCredential) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddCredentials200ResponseAllOfCredential) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AddCredentials200ResponseAllOfCredential) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddCredentials200ResponseAllOfCredential) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddCredentials200ResponseAllOfCredential) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddCredentials200ResponseAllOfCredential) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *AddCredentials200ResponseAllOfCredential) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddCredentials200ResponseAllOfCredential) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddCredentials200ResponseAllOfCredential) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddCredentials200ResponseAllOfCredential) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScope

`func (o *AddCredentials200ResponseAllOfCredential) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddCredentials200ResponseAllOfCredential) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddCredentials200ResponseAllOfCredential) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AddCredentials200ResponseAllOfCredential) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *AddCredentials200ResponseAllOfCredential) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddCredentials200ResponseAllOfCredential) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddCredentials200ResponseAllOfCredential) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddCredentials200ResponseAllOfCredential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddCredentials200ResponseAllOfCredential) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddCredentials200ResponseAllOfCredential) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddCredentials200ResponseAllOfCredential) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddCredentials200ResponseAllOfCredential) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *AddCredentials200ResponseAllOfCredential) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddCredentials200ResponseAllOfCredential) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddCredentials200ResponseAllOfCredential) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddCredentials200ResponseAllOfCredential) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCredentials200ResponseAllOfCredential) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCredentials200ResponseAllOfCredential) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCredentials200ResponseAllOfCredential) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCredentials200ResponseAllOfCredential) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *AddCredentials200ResponseAllOfCredential) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddCredentials200ResponseAllOfCredential) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddCredentials200ResponseAllOfCredential) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddCredentials200ResponseAllOfCredential) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetUser

`func (o *AddCredentials200ResponseAllOfCredential) GetUser() AddCredentials200ResponseAllOfCredentialUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddCredentials200ResponseAllOfCredential) GetUserOk() (*AddCredentials200ResponseAllOfCredentialUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddCredentials200ResponseAllOfCredential) SetUser(v AddCredentials200ResponseAllOfCredentialUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AddCredentials200ResponseAllOfCredential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddCredentials200ResponseAllOfCredential) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddCredentials200ResponseAllOfCredential) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddCredentials200ResponseAllOfCredential) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddCredentials200ResponseAllOfCredential) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddCredentials200ResponseAllOfCredential) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddCredentials200ResponseAllOfCredential) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddCredentials200ResponseAllOfCredential) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddCredentials200ResponseAllOfCredential) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *AddCredentials200ResponseAllOfCredential) GetConfig() AddCredentials200ResponseAllOfCredentialConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddCredentials200ResponseAllOfCredential) GetConfigOk() (*AddCredentials200ResponseAllOfCredentialConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddCredentials200ResponseAllOfCredential) SetConfig(v AddCredentials200ResponseAllOfCredentialConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddCredentials200ResponseAllOfCredential) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


