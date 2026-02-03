# TaskPuppetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**PuppetEnvironment** | Pointer to **NullableString** |  | [optional] 
**PuppetNodeName** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskPuppetConfig

`func NewTaskPuppetConfig() *TaskPuppetConfig`

NewTaskPuppetConfig instantiates a new TaskPuppetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPuppetConfigWithDefaults

`func NewTaskPuppetConfigWithDefaults() *TaskPuppetConfig`

NewTaskPuppetConfigWithDefaults instantiates a new TaskPuppetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *TaskPuppetConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskPuppetConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskPuppetConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskPuppetConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskPuppetConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskPuppetConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHost

`func (o *TaskPuppetConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskPuppetConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskPuppetConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskPuppetConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskPuppetConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskPuppetConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *TaskPuppetConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskPuppetConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskPuppetConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskPuppetConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskPuppetConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskPuppetConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPuppetEnvironment

`func (o *TaskPuppetConfig) GetPuppetEnvironment() string`

GetPuppetEnvironment returns the PuppetEnvironment field if non-nil, zero value otherwise.

### GetPuppetEnvironmentOk

`func (o *TaskPuppetConfig) GetPuppetEnvironmentOk() (*string, bool)`

GetPuppetEnvironmentOk returns a tuple with the PuppetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetEnvironment

`func (o *TaskPuppetConfig) SetPuppetEnvironment(v string)`

SetPuppetEnvironment sets PuppetEnvironment field to given value.

### HasPuppetEnvironment

`func (o *TaskPuppetConfig) HasPuppetEnvironment() bool`

HasPuppetEnvironment returns a boolean if a field has been set.

### SetPuppetEnvironmentNil

`func (o *TaskPuppetConfig) SetPuppetEnvironmentNil(b bool)`

 SetPuppetEnvironmentNil sets the value for PuppetEnvironment to be an explicit nil

### UnsetPuppetEnvironment
`func (o *TaskPuppetConfig) UnsetPuppetEnvironment()`

UnsetPuppetEnvironment ensures that no value is present for PuppetEnvironment, not even an explicit nil
### GetPuppetNodeName

`func (o *TaskPuppetConfig) GetPuppetNodeName() string`

GetPuppetNodeName returns the PuppetNodeName field if non-nil, zero value otherwise.

### GetPuppetNodeNameOk

`func (o *TaskPuppetConfig) GetPuppetNodeNameOk() (*string, bool)`

GetPuppetNodeNameOk returns a tuple with the PuppetNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetNodeName

`func (o *TaskPuppetConfig) SetPuppetNodeName(v string)`

SetPuppetNodeName sets PuppetNodeName field to given value.

### HasPuppetNodeName

`func (o *TaskPuppetConfig) HasPuppetNodeName() bool`

HasPuppetNodeName returns a boolean if a field has been set.

### SetPuppetNodeNameNil

`func (o *TaskPuppetConfig) SetPuppetNodeNameNil(b bool)`

 SetPuppetNodeNameNil sets the value for PuppetNodeName to be an explicit nil

### UnsetPuppetNodeName
`func (o *TaskPuppetConfig) UnsetPuppetNodeName()`

UnsetPuppetNodeName ensures that no value is present for PuppetNodeName, not even an explicit nil
### GetSshKey

`func (o *TaskPuppetConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskPuppetConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskPuppetConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskPuppetConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskPuppetConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskPuppetConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskPuppetConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskPuppetConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskPuppetConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskPuppetConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskPuppetConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskPuppetConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetLocalScriptGitRef

`func (o *TaskPuppetConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskPuppetConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskPuppetConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskPuppetConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskPuppetConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskPuppetConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetPassword

`func (o *TaskPuppetConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskPuppetConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskPuppetConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskPuppetConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskPuppetConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskPuppetConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskPuppetConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskPuppetConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskPuppetConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskPuppetConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskPuppetConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskPuppetConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


