# ListCredentials200ResponseAllOfCredentialsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Integration** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**AuthKey** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**AuthPath** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 
**User** | Pointer to [**ListCredentials200ResponseAllOfCredentialsInnerUser**](ListCredentials200ResponseAllOfCredentialsInnerUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to [**ListCredentials200ResponseAllOfCredentialsInnerConfig**](ListCredentials200ResponseAllOfCredentialsInnerConfig.md) |  | [optional] 

## Methods

### NewListCredentials200ResponseAllOfCredentialsInner

`func NewListCredentials200ResponseAllOfCredentialsInner() *ListCredentials200ResponseAllOfCredentialsInner`

NewListCredentials200ResponseAllOfCredentialsInner instantiates a new ListCredentials200ResponseAllOfCredentialsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentials200ResponseAllOfCredentialsInnerWithDefaults

`func NewListCredentials200ResponseAllOfCredentialsInnerWithDefaults() *ListCredentials200ResponseAllOfCredentialsInner`

NewListCredentials200ResponseAllOfCredentialsInnerWithDefaults instantiates a new ListCredentials200ResponseAllOfCredentialsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetDescription

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUsername

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetAuthKey

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetAuthKey() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetAuthKeyOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetAuthKey(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPath

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.

### HasAuthPath

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasAuthPath() bool`

HasAuthPath returns a boolean if a field has been set.

### SetAuthPathNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetAuthPathNil(b bool)`

 SetAuthPathNil sets the value for AuthPath to be an explicit nil

### UnsetAuthPath
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetAuthPath()`

UnsetAuthPath ensures that no value is present for AuthPath, not even an explicit nil
### GetExternalId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetRefType

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetCategory

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetScope

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetEnabled

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetAccount() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetAccountOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetAccount(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ListCredentials200ResponseAllOfCredentialsInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetUser

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetUser() ListCredentials200ResponseAllOfCredentialsInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetUserOk() (*ListCredentials200ResponseAllOfCredentialsInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetUser(v ListCredentials200ResponseAllOfCredentialsInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetConfig() ListCredentials200ResponseAllOfCredentialsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListCredentials200ResponseAllOfCredentialsInner) GetConfigOk() (*ListCredentials200ResponseAllOfCredentialsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListCredentials200ResponseAllOfCredentialsInner) SetConfig(v ListCredentials200ResponseAllOfCredentialsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListCredentials200ResponseAllOfCredentialsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


