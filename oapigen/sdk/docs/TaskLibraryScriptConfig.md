# TaskLibraryScriptConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**ContainerScriptId** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**ContainerScript** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskLibraryScriptConfig

`func NewTaskLibraryScriptConfig() *TaskLibraryScriptConfig`

NewTaskLibraryScriptConfig instantiates a new TaskLibraryScriptConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLibraryScriptConfigWithDefaults

`func NewTaskLibraryScriptConfigWithDefaults() *TaskLibraryScriptConfig`

NewTaskLibraryScriptConfigWithDefaults instantiates a new TaskLibraryScriptConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TaskLibraryScriptConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskLibraryScriptConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskLibraryScriptConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskLibraryScriptConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskLibraryScriptConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskLibraryScriptConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *TaskLibraryScriptConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskLibraryScriptConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskLibraryScriptConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskLibraryScriptConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskLibraryScriptConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskLibraryScriptConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *TaskLibraryScriptConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskLibraryScriptConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskLibraryScriptConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskLibraryScriptConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskLibraryScriptConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskLibraryScriptConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskLibraryScriptConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskLibraryScriptConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskLibraryScriptConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskLibraryScriptConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskLibraryScriptConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskLibraryScriptConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetLocalScriptGitRef

`func (o *TaskLibraryScriptConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskLibraryScriptConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskLibraryScriptConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskLibraryScriptConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskLibraryScriptConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskLibraryScriptConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskLibraryScriptConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskLibraryScriptConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskLibraryScriptConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskLibraryScriptConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskLibraryScriptConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskLibraryScriptConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetContainerScriptId

`func (o *TaskLibraryScriptConfig) GetContainerScriptId() string`

GetContainerScriptId returns the ContainerScriptId field if non-nil, zero value otherwise.

### GetContainerScriptIdOk

`func (o *TaskLibraryScriptConfig) GetContainerScriptIdOk() (*string, bool)`

GetContainerScriptIdOk returns a tuple with the ContainerScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScriptId

`func (o *TaskLibraryScriptConfig) SetContainerScriptId(v string)`

SetContainerScriptId sets ContainerScriptId field to given value.

### HasContainerScriptId

`func (o *TaskLibraryScriptConfig) HasContainerScriptId() bool`

HasContainerScriptId returns a boolean if a field has been set.

### SetContainerScriptIdNil

`func (o *TaskLibraryScriptConfig) SetContainerScriptIdNil(b bool)`

 SetContainerScriptIdNil sets the value for ContainerScriptId to be an explicit nil

### UnsetContainerScriptId
`func (o *TaskLibraryScriptConfig) UnsetContainerScriptId()`

UnsetContainerScriptId ensures that no value is present for ContainerScriptId, not even an explicit nil
### GetSshKey

`func (o *TaskLibraryScriptConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskLibraryScriptConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskLibraryScriptConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskLibraryScriptConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskLibraryScriptConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskLibraryScriptConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetContainerScript

`func (o *TaskLibraryScriptConfig) GetContainerScript() string`

GetContainerScript returns the ContainerScript field if non-nil, zero value otherwise.

### GetContainerScriptOk

`func (o *TaskLibraryScriptConfig) GetContainerScriptOk() (*string, bool)`

GetContainerScriptOk returns a tuple with the ContainerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScript

`func (o *TaskLibraryScriptConfig) SetContainerScript(v string)`

SetContainerScript sets ContainerScript field to given value.

### HasContainerScript

`func (o *TaskLibraryScriptConfig) HasContainerScript() bool`

HasContainerScript returns a boolean if a field has been set.

### SetContainerScriptNil

`func (o *TaskLibraryScriptConfig) SetContainerScriptNil(b bool)`

 SetContainerScriptNil sets the value for ContainerScript to be an explicit nil

### UnsetContainerScript
`func (o *TaskLibraryScriptConfig) UnsetContainerScript()`

UnsetContainerScript ensures that no value is present for ContainerScript, not even an explicit nil
### GetPort

`func (o *TaskLibraryScriptConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskLibraryScriptConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskLibraryScriptConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskLibraryScriptConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskLibraryScriptConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskLibraryScriptConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


