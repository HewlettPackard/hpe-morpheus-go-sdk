# UpdateCredentials200ResponseAllOfCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UpdateCredentials200ResponseAllOfCredentialType**](UpdateCredentials200ResponseAllOfCredentialType.md) |  | [optional] 
**Integration** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**AuthKey** | Pointer to [**UpdateCredentials200ResponseAllOfCredentialAuthKey**](UpdateCredentials200ResponseAllOfCredentialAuthKey.md) |  | [optional] 
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
**Account** | Pointer to [**NullableUpdateCredentials200ResponseAllOfCredentialAccount**](UpdateCredentials200ResponseAllOfCredentialAccount.md) |  | [optional] 
**User** | Pointer to [**UpdateCredentials200ResponseAllOfCredentialUser**](UpdateCredentials200ResponseAllOfCredentialUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to [**UpdateCredentials200ResponseAllOfCredentialConfig**](UpdateCredentials200ResponseAllOfCredentialConfig.md) |  | [optional] 

## Methods

### NewUpdateCredentials200ResponseAllOfCredential

`func NewUpdateCredentials200ResponseAllOfCredential() *UpdateCredentials200ResponseAllOfCredential`

NewUpdateCredentials200ResponseAllOfCredential instantiates a new UpdateCredentials200ResponseAllOfCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentials200ResponseAllOfCredentialWithDefaults

`func NewUpdateCredentials200ResponseAllOfCredentialWithDefaults() *UpdateCredentials200ResponseAllOfCredential`

NewUpdateCredentials200ResponseAllOfCredentialWithDefaults instantiates a new UpdateCredentials200ResponseAllOfCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCredentials200ResponseAllOfCredential) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCredentials200ResponseAllOfCredential) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCredentials200ResponseAllOfCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCredentials200ResponseAllOfCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCredentials200ResponseAllOfCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCredentials200ResponseAllOfCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UpdateCredentials200ResponseAllOfCredential) GetType() UpdateCredentials200ResponseAllOfCredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetTypeOk() (*UpdateCredentials200ResponseAllOfCredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCredentials200ResponseAllOfCredential) SetType(v UpdateCredentials200ResponseAllOfCredentialType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCredentials200ResponseAllOfCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateCredentials200ResponseAllOfCredential) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateCredentials200ResponseAllOfCredential) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateCredentials200ResponseAllOfCredential) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetDescription

`func (o *UpdateCredentials200ResponseAllOfCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCredentials200ResponseAllOfCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCredentials200ResponseAllOfCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUsername

`func (o *UpdateCredentials200ResponseAllOfCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCredentials200ResponseAllOfCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateCredentials200ResponseAllOfCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateCredentials200ResponseAllOfCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCredentials200ResponseAllOfCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateCredentials200ResponseAllOfCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *UpdateCredentials200ResponseAllOfCredential) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *UpdateCredentials200ResponseAllOfCredential) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *UpdateCredentials200ResponseAllOfCredential) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetAuthKey

`func (o *UpdateCredentials200ResponseAllOfCredential) GetAuthKey() UpdateCredentials200ResponseAllOfCredentialAuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetAuthKeyOk() (*UpdateCredentials200ResponseAllOfCredentialAuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *UpdateCredentials200ResponseAllOfCredential) SetAuthKey(v UpdateCredentials200ResponseAllOfCredentialAuthKey)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *UpdateCredentials200ResponseAllOfCredential) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthPath

`func (o *UpdateCredentials200ResponseAllOfCredential) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *UpdateCredentials200ResponseAllOfCredential) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.

### HasAuthPath

`func (o *UpdateCredentials200ResponseAllOfCredential) HasAuthPath() bool`

HasAuthPath returns a boolean if a field has been set.

### SetAuthPathNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetAuthPathNil(b bool)`

 SetAuthPathNil sets the value for AuthPath to be an explicit nil

### UnsetAuthPath
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetAuthPath()`

UnsetAuthPath ensures that no value is present for AuthPath, not even an explicit nil
### GetExternalId

`func (o *UpdateCredentials200ResponseAllOfCredential) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateCredentials200ResponseAllOfCredential) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateCredentials200ResponseAllOfCredential) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetRefType

`func (o *UpdateCredentials200ResponseAllOfCredential) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateCredentials200ResponseAllOfCredential) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateCredentials200ResponseAllOfCredential) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *UpdateCredentials200ResponseAllOfCredential) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateCredentials200ResponseAllOfCredential) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateCredentials200ResponseAllOfCredential) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetCategory

`func (o *UpdateCredentials200ResponseAllOfCredential) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCredentials200ResponseAllOfCredential) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCredentials200ResponseAllOfCredential) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetScope

`func (o *UpdateCredentials200ResponseAllOfCredential) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateCredentials200ResponseAllOfCredential) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateCredentials200ResponseAllOfCredential) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateCredentials200ResponseAllOfCredential) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCredentials200ResponseAllOfCredential) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateCredentials200ResponseAllOfCredential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateCredentials200ResponseAllOfCredential) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateCredentials200ResponseAllOfCredential) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateCredentials200ResponseAllOfCredential) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateCredentials200ResponseAllOfCredential) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateCredentials200ResponseAllOfCredential) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateCredentials200ResponseAllOfCredential) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetEnabled

`func (o *UpdateCredentials200ResponseAllOfCredential) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCredentials200ResponseAllOfCredential) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCredentials200ResponseAllOfCredential) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateCredentials200ResponseAllOfCredential) GetAccount() UpdateCredentials200ResponseAllOfCredentialAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetAccountOk() (*UpdateCredentials200ResponseAllOfCredentialAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateCredentials200ResponseAllOfCredential) SetAccount(v UpdateCredentials200ResponseAllOfCredentialAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateCredentials200ResponseAllOfCredential) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *UpdateCredentials200ResponseAllOfCredential) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *UpdateCredentials200ResponseAllOfCredential) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetUser

`func (o *UpdateCredentials200ResponseAllOfCredential) GetUser() UpdateCredentials200ResponseAllOfCredentialUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetUserOk() (*UpdateCredentials200ResponseAllOfCredentialUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateCredentials200ResponseAllOfCredential) SetUser(v UpdateCredentials200ResponseAllOfCredentialUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateCredentials200ResponseAllOfCredential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateCredentials200ResponseAllOfCredential) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateCredentials200ResponseAllOfCredential) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateCredentials200ResponseAllOfCredential) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateCredentials200ResponseAllOfCredential) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateCredentials200ResponseAllOfCredential) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateCredentials200ResponseAllOfCredential) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateCredentials200ResponseAllOfCredential) GetConfig() UpdateCredentials200ResponseAllOfCredentialConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCredentials200ResponseAllOfCredential) GetConfigOk() (*UpdateCredentials200ResponseAllOfCredentialConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCredentials200ResponseAllOfCredential) SetConfig(v UpdateCredentials200ResponseAllOfCredentialConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateCredentials200ResponseAllOfCredential) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


