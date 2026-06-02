# GetCredentials200ResponseCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetCredentials200ResponseCredentialType**](GetCredentials200ResponseCredentialType.md) |  | [optional] 
**Integration** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**AuthKey** | Pointer to [**GetCredentials200ResponseCredentialAuthKey**](GetCredentials200ResponseCredentialAuthKey.md) |  | [optional] 
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
**Account** | Pointer to [**NullableGetCredentials200ResponseCredentialAccount**](GetCredentials200ResponseCredentialAccount.md) |  | [optional] 
**User** | Pointer to [**GetCredentials200ResponseCredentialUser**](GetCredentials200ResponseCredentialUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to [**GetCredentials200ResponseCredentialConfig**](GetCredentials200ResponseCredentialConfig.md) |  | [optional] 

## Methods

### NewGetCredentials200ResponseCredential

`func NewGetCredentials200ResponseCredential() *GetCredentials200ResponseCredential`

NewGetCredentials200ResponseCredential instantiates a new GetCredentials200ResponseCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCredentials200ResponseCredential) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCredentials200ResponseCredential) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCredentials200ResponseCredential) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCredentials200ResponseCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCredentials200ResponseCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCredentials200ResponseCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCredentials200ResponseCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCredentials200ResponseCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetCredentials200ResponseCredential) GetType() GetCredentials200ResponseCredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCredentials200ResponseCredential) GetTypeOk() (*GetCredentials200ResponseCredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCredentials200ResponseCredential) SetType(v GetCredentials200ResponseCredentialType)`

SetType sets Type field to given value.

### HasType

`func (o *GetCredentials200ResponseCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *GetCredentials200ResponseCredential) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *GetCredentials200ResponseCredential) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *GetCredentials200ResponseCredential) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *GetCredentials200ResponseCredential) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *GetCredentials200ResponseCredential) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *GetCredentials200ResponseCredential) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetDescription

`func (o *GetCredentials200ResponseCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCredentials200ResponseCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCredentials200ResponseCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCredentials200ResponseCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCredentials200ResponseCredential) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCredentials200ResponseCredential) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUsername

`func (o *GetCredentials200ResponseCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetCredentials200ResponseCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetCredentials200ResponseCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetCredentials200ResponseCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetCredentials200ResponseCredential) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetCredentials200ResponseCredential) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *GetCredentials200ResponseCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetCredentials200ResponseCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetCredentials200ResponseCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetCredentials200ResponseCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GetCredentials200ResponseCredential) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GetCredentials200ResponseCredential) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *GetCredentials200ResponseCredential) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *GetCredentials200ResponseCredential) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *GetCredentials200ResponseCredential) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *GetCredentials200ResponseCredential) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *GetCredentials200ResponseCredential) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *GetCredentials200ResponseCredential) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetAuthKey

`func (o *GetCredentials200ResponseCredential) GetAuthKey() GetCredentials200ResponseCredentialAuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *GetCredentials200ResponseCredential) GetAuthKeyOk() (*GetCredentials200ResponseCredentialAuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *GetCredentials200ResponseCredential) SetAuthKey(v GetCredentials200ResponseCredentialAuthKey)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *GetCredentials200ResponseCredential) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPath

`func (o *GetCredentials200ResponseCredential) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *GetCredentials200ResponseCredential) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *GetCredentials200ResponseCredential) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.

### HasAuthPath

`func (o *GetCredentials200ResponseCredential) HasAuthPath() bool`

HasAuthPath returns a boolean if a field has been set.

### SetAuthPathNil

`func (o *GetCredentials200ResponseCredential) SetAuthPathNil(b bool)`

 SetAuthPathNil sets the value for AuthPath to be an explicit nil

### UnsetAuthPath
`func (o *GetCredentials200ResponseCredential) UnsetAuthPath()`

UnsetAuthPath ensures that no value is present for AuthPath, not even an explicit nil
### GetExternalId

`func (o *GetCredentials200ResponseCredential) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetCredentials200ResponseCredential) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetCredentials200ResponseCredential) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetCredentials200ResponseCredential) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetCredentials200ResponseCredential) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetCredentials200ResponseCredential) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetRefType

`func (o *GetCredentials200ResponseCredential) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetCredentials200ResponseCredential) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetCredentials200ResponseCredential) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetCredentials200ResponseCredential) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetCredentials200ResponseCredential) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetCredentials200ResponseCredential) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetCredentials200ResponseCredential) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetCredentials200ResponseCredential) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetCredentials200ResponseCredential) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetCredentials200ResponseCredential) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetCredentials200ResponseCredential) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetCredentials200ResponseCredential) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetCategory

`func (o *GetCredentials200ResponseCredential) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetCredentials200ResponseCredential) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetCredentials200ResponseCredential) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetCredentials200ResponseCredential) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetCredentials200ResponseCredential) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetCredentials200ResponseCredential) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetScope

`func (o *GetCredentials200ResponseCredential) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetCredentials200ResponseCredential) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetCredentials200ResponseCredential) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetCredentials200ResponseCredential) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *GetCredentials200ResponseCredential) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCredentials200ResponseCredential) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCredentials200ResponseCredential) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCredentials200ResponseCredential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetCredentials200ResponseCredential) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetCredentials200ResponseCredential) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetCredentials200ResponseCredential) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetCredentials200ResponseCredential) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetCredentials200ResponseCredential) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetCredentials200ResponseCredential) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetCredentials200ResponseCredential) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetCredentials200ResponseCredential) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetCredentials200ResponseCredential) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetCredentials200ResponseCredential) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetCredentials200ResponseCredential) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetCredentials200ResponseCredential) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetEnabled

`func (o *GetCredentials200ResponseCredential) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCredentials200ResponseCredential) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCredentials200ResponseCredential) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCredentials200ResponseCredential) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *GetCredentials200ResponseCredential) GetAccount() GetCredentials200ResponseCredentialAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCredentials200ResponseCredential) GetAccountOk() (*GetCredentials200ResponseCredentialAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCredentials200ResponseCredential) SetAccount(v GetCredentials200ResponseCredentialAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCredentials200ResponseCredential) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetCredentials200ResponseCredential) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetCredentials200ResponseCredential) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetUser

`func (o *GetCredentials200ResponseCredential) GetUser() GetCredentials200ResponseCredentialUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetCredentials200ResponseCredential) GetUserOk() (*GetCredentials200ResponseCredentialUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetCredentials200ResponseCredential) SetUser(v GetCredentials200ResponseCredentialUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetCredentials200ResponseCredential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCredentials200ResponseCredential) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCredentials200ResponseCredential) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCredentials200ResponseCredential) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCredentials200ResponseCredential) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCredentials200ResponseCredential) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCredentials200ResponseCredential) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCredentials200ResponseCredential) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCredentials200ResponseCredential) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *GetCredentials200ResponseCredential) GetConfig() GetCredentials200ResponseCredentialConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetCredentials200ResponseCredential) GetConfigOk() (*GetCredentials200ResponseCredentialConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetCredentials200ResponseCredential) SetConfig(v GetCredentials200ResponseCredentialConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetCredentials200ResponseCredential) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


