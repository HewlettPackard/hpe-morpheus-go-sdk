# TaskAnsibleTowerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**AnsibleTowerGitRef** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**AnsibleGroup** | Pointer to **NullableString** |  | [optional] 
**AnsibleTowerExecuteMode** | Pointer to **string** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskAnsibleTowerConfig

`func NewTaskAnsibleTowerConfig() *TaskAnsibleTowerConfig`

NewTaskAnsibleTowerConfig instantiates a new TaskAnsibleTowerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAnsibleTowerConfigWithDefaults

`func NewTaskAnsibleTowerConfigWithDefaults() *TaskAnsibleTowerConfig`

NewTaskAnsibleTowerConfigWithDefaults instantiates a new TaskAnsibleTowerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *TaskAnsibleTowerConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskAnsibleTowerConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskAnsibleTowerConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskAnsibleTowerConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskAnsibleTowerConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskAnsibleTowerConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskAnsibleTowerConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskAnsibleTowerConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskAnsibleTowerConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskAnsibleTowerConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskAnsibleTowerConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskAnsibleTowerConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetAnsibleTowerGitRef

`func (o *TaskAnsibleTowerConfig) GetAnsibleTowerGitRef() string`

GetAnsibleTowerGitRef returns the AnsibleTowerGitRef field if non-nil, zero value otherwise.

### GetAnsibleTowerGitRefOk

`func (o *TaskAnsibleTowerConfig) GetAnsibleTowerGitRefOk() (*string, bool)`

GetAnsibleTowerGitRefOk returns a tuple with the AnsibleTowerGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerGitRef

`func (o *TaskAnsibleTowerConfig) SetAnsibleTowerGitRef(v string)`

SetAnsibleTowerGitRef sets AnsibleTowerGitRef field to given value.

### HasAnsibleTowerGitRef

`func (o *TaskAnsibleTowerConfig) HasAnsibleTowerGitRef() bool`

HasAnsibleTowerGitRef returns a boolean if a field has been set.

### SetAnsibleTowerGitRefNil

`func (o *TaskAnsibleTowerConfig) SetAnsibleTowerGitRefNil(b bool)`

 SetAnsibleTowerGitRefNil sets the value for AnsibleTowerGitRef to be an explicit nil

### UnsetAnsibleTowerGitRef
`func (o *TaskAnsibleTowerConfig) UnsetAnsibleTowerGitRef()`

UnsetAnsibleTowerGitRef ensures that no value is present for AnsibleTowerGitRef, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskAnsibleTowerConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskAnsibleTowerConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskAnsibleTowerConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskAnsibleTowerConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskAnsibleTowerConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskAnsibleTowerConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetHost

`func (o *TaskAnsibleTowerConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskAnsibleTowerConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskAnsibleTowerConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskAnsibleTowerConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskAnsibleTowerConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskAnsibleTowerConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *TaskAnsibleTowerConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskAnsibleTowerConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskAnsibleTowerConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskAnsibleTowerConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskAnsibleTowerConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskAnsibleTowerConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSshKey

`func (o *TaskAnsibleTowerConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskAnsibleTowerConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskAnsibleTowerConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskAnsibleTowerConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskAnsibleTowerConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskAnsibleTowerConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetAnsibleGroup

`func (o *TaskAnsibleTowerConfig) GetAnsibleGroup() string`

GetAnsibleGroup returns the AnsibleGroup field if non-nil, zero value otherwise.

### GetAnsibleGroupOk

`func (o *TaskAnsibleTowerConfig) GetAnsibleGroupOk() (*string, bool)`

GetAnsibleGroupOk returns a tuple with the AnsibleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroup

`func (o *TaskAnsibleTowerConfig) SetAnsibleGroup(v string)`

SetAnsibleGroup sets AnsibleGroup field to given value.

### HasAnsibleGroup

`func (o *TaskAnsibleTowerConfig) HasAnsibleGroup() bool`

HasAnsibleGroup returns a boolean if a field has been set.

### SetAnsibleGroupNil

`func (o *TaskAnsibleTowerConfig) SetAnsibleGroupNil(b bool)`

 SetAnsibleGroupNil sets the value for AnsibleGroup to be an explicit nil

### UnsetAnsibleGroup
`func (o *TaskAnsibleTowerConfig) UnsetAnsibleGroup()`

UnsetAnsibleGroup ensures that no value is present for AnsibleGroup, not even an explicit nil
### GetAnsibleTowerExecuteMode

`func (o *TaskAnsibleTowerConfig) GetAnsibleTowerExecuteMode() string`

GetAnsibleTowerExecuteMode returns the AnsibleTowerExecuteMode field if non-nil, zero value otherwise.

### GetAnsibleTowerExecuteModeOk

`func (o *TaskAnsibleTowerConfig) GetAnsibleTowerExecuteModeOk() (*string, bool)`

GetAnsibleTowerExecuteModeOk returns a tuple with the AnsibleTowerExecuteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerExecuteMode

`func (o *TaskAnsibleTowerConfig) SetAnsibleTowerExecuteMode(v string)`

SetAnsibleTowerExecuteMode sets AnsibleTowerExecuteMode field to given value.

### HasAnsibleTowerExecuteMode

`func (o *TaskAnsibleTowerConfig) HasAnsibleTowerExecuteMode() bool`

HasAnsibleTowerExecuteMode returns a boolean if a field has been set.

### GetLocalScriptGitRef

`func (o *TaskAnsibleTowerConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskAnsibleTowerConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskAnsibleTowerConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskAnsibleTowerConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskAnsibleTowerConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskAnsibleTowerConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetPort

`func (o *TaskAnsibleTowerConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskAnsibleTowerConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskAnsibleTowerConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskAnsibleTowerConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskAnsibleTowerConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskAnsibleTowerConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


