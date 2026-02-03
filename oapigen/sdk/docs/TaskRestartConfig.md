# TaskRestartConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskRestartConfig

`func NewTaskRestartConfig() *TaskRestartConfig`

NewTaskRestartConfig instantiates a new TaskRestartConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRestartConfigWithDefaults

`func NewTaskRestartConfigWithDefaults() *TaskRestartConfig`

NewTaskRestartConfigWithDefaults instantiates a new TaskRestartConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *TaskRestartConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskRestartConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskRestartConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskRestartConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskRestartConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskRestartConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskRestartConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskRestartConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskRestartConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskRestartConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskRestartConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskRestartConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetLocalScriptGitRef

`func (o *TaskRestartConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskRestartConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskRestartConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskRestartConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskRestartConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskRestartConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetHost

`func (o *TaskRestartConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskRestartConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskRestartConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskRestartConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskRestartConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskRestartConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *TaskRestartConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskRestartConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskRestartConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskRestartConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskRestartConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskRestartConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskRestartConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskRestartConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskRestartConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskRestartConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskRestartConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskRestartConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetSshKey

`func (o *TaskRestartConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskRestartConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskRestartConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskRestartConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskRestartConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskRestartConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetPort

`func (o *TaskRestartConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskRestartConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskRestartConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskRestartConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskRestartConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskRestartConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


