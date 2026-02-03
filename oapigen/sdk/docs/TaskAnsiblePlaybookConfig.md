# TaskAnsiblePlaybookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnsibleOptions** | Pointer to **NullableString** |  | [optional] 
**AnsiblePlaybook** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**AnsibleGitId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**AnsibleSkipTags** | Pointer to **NullableString** |  | [optional] 
**AnsibleTags** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**AnsibleGitRef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskAnsiblePlaybookConfig

`func NewTaskAnsiblePlaybookConfig() *TaskAnsiblePlaybookConfig`

NewTaskAnsiblePlaybookConfig instantiates a new TaskAnsiblePlaybookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAnsiblePlaybookConfigWithDefaults

`func NewTaskAnsiblePlaybookConfigWithDefaults() *TaskAnsiblePlaybookConfig`

NewTaskAnsiblePlaybookConfigWithDefaults instantiates a new TaskAnsiblePlaybookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleOptions

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleOptions() string`

GetAnsibleOptions returns the AnsibleOptions field if non-nil, zero value otherwise.

### GetAnsibleOptionsOk

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleOptionsOk() (*string, bool)`

GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleOptions

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleOptions(v string)`

SetAnsibleOptions sets AnsibleOptions field to given value.

### HasAnsibleOptions

`func (o *TaskAnsiblePlaybookConfig) HasAnsibleOptions() bool`

HasAnsibleOptions returns a boolean if a field has been set.

### SetAnsibleOptionsNil

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleOptionsNil(b bool)`

 SetAnsibleOptionsNil sets the value for AnsibleOptions to be an explicit nil

### UnsetAnsibleOptions
`func (o *TaskAnsiblePlaybookConfig) UnsetAnsibleOptions()`

UnsetAnsibleOptions ensures that no value is present for AnsibleOptions, not even an explicit nil
### GetAnsiblePlaybook

`func (o *TaskAnsiblePlaybookConfig) GetAnsiblePlaybook() string`

GetAnsiblePlaybook returns the AnsiblePlaybook field if non-nil, zero value otherwise.

### GetAnsiblePlaybookOk

`func (o *TaskAnsiblePlaybookConfig) GetAnsiblePlaybookOk() (*string, bool)`

GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybook

`func (o *TaskAnsiblePlaybookConfig) SetAnsiblePlaybook(v string)`

SetAnsiblePlaybook sets AnsiblePlaybook field to given value.

### HasAnsiblePlaybook

`func (o *TaskAnsiblePlaybookConfig) HasAnsiblePlaybook() bool`

HasAnsiblePlaybook returns a boolean if a field has been set.

### SetAnsiblePlaybookNil

`func (o *TaskAnsiblePlaybookConfig) SetAnsiblePlaybookNil(b bool)`

 SetAnsiblePlaybookNil sets the value for AnsiblePlaybook to be an explicit nil

### UnsetAnsiblePlaybook
`func (o *TaskAnsiblePlaybookConfig) UnsetAnsiblePlaybook()`

UnsetAnsiblePlaybook ensures that no value is present for AnsiblePlaybook, not even an explicit nil
### GetSshKey

`func (o *TaskAnsiblePlaybookConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskAnsiblePlaybookConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskAnsiblePlaybookConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskAnsiblePlaybookConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskAnsiblePlaybookConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskAnsiblePlaybookConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetPort

`func (o *TaskAnsiblePlaybookConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskAnsiblePlaybookConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskAnsiblePlaybookConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskAnsiblePlaybookConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskAnsiblePlaybookConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskAnsiblePlaybookConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetLocalScriptGitRef

`func (o *TaskAnsiblePlaybookConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskAnsiblePlaybookConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskAnsiblePlaybookConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskAnsiblePlaybookConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskAnsiblePlaybookConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskAnsiblePlaybookConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetPassword

`func (o *TaskAnsiblePlaybookConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskAnsiblePlaybookConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskAnsiblePlaybookConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskAnsiblePlaybookConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskAnsiblePlaybookConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskAnsiblePlaybookConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskAnsiblePlaybookConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskAnsiblePlaybookConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskAnsiblePlaybookConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskAnsiblePlaybookConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskAnsiblePlaybookConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskAnsiblePlaybookConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskAnsiblePlaybookConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskAnsiblePlaybookConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskAnsiblePlaybookConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskAnsiblePlaybookConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskAnsiblePlaybookConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskAnsiblePlaybookConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetAnsibleGitId

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleGitId() string`

GetAnsibleGitId returns the AnsibleGitId field if non-nil, zero value otherwise.

### GetAnsibleGitIdOk

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleGitIdOk() (*string, bool)`

GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitId

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleGitId(v string)`

SetAnsibleGitId sets AnsibleGitId field to given value.

### HasAnsibleGitId

`func (o *TaskAnsiblePlaybookConfig) HasAnsibleGitId() bool`

HasAnsibleGitId returns a boolean if a field has been set.

### GetHost

`func (o *TaskAnsiblePlaybookConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskAnsiblePlaybookConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskAnsiblePlaybookConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskAnsiblePlaybookConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskAnsiblePlaybookConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskAnsiblePlaybookConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetAnsibleSkipTags

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleSkipTags() string`

GetAnsibleSkipTags returns the AnsibleSkipTags field if non-nil, zero value otherwise.

### GetAnsibleSkipTagsOk

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleSkipTagsOk() (*string, bool)`

GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleSkipTags

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleSkipTags(v string)`

SetAnsibleSkipTags sets AnsibleSkipTags field to given value.

### HasAnsibleSkipTags

`func (o *TaskAnsiblePlaybookConfig) HasAnsibleSkipTags() bool`

HasAnsibleSkipTags returns a boolean if a field has been set.

### SetAnsibleSkipTagsNil

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleSkipTagsNil(b bool)`

 SetAnsibleSkipTagsNil sets the value for AnsibleSkipTags to be an explicit nil

### UnsetAnsibleSkipTags
`func (o *TaskAnsiblePlaybookConfig) UnsetAnsibleSkipTags()`

UnsetAnsibleSkipTags ensures that no value is present for AnsibleSkipTags, not even an explicit nil
### GetAnsibleTags

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleTags() string`

GetAnsibleTags returns the AnsibleTags field if non-nil, zero value otherwise.

### GetAnsibleTagsOk

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleTagsOk() (*string, bool)`

GetAnsibleTagsOk returns a tuple with the AnsibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTags

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleTags(v string)`

SetAnsibleTags sets AnsibleTags field to given value.

### HasAnsibleTags

`func (o *TaskAnsiblePlaybookConfig) HasAnsibleTags() bool`

HasAnsibleTags returns a boolean if a field has been set.

### SetAnsibleTagsNil

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleTagsNil(b bool)`

 SetAnsibleTagsNil sets the value for AnsibleTags to be an explicit nil

### UnsetAnsibleTags
`func (o *TaskAnsiblePlaybookConfig) UnsetAnsibleTags()`

UnsetAnsibleTags ensures that no value is present for AnsibleTags, not even an explicit nil
### GetUsername

`func (o *TaskAnsiblePlaybookConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskAnsiblePlaybookConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskAnsiblePlaybookConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskAnsiblePlaybookConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskAnsiblePlaybookConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskAnsiblePlaybookConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAnsibleGitRef

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleGitRef() string`

GetAnsibleGitRef returns the AnsibleGitRef field if non-nil, zero value otherwise.

### GetAnsibleGitRefOk

`func (o *TaskAnsiblePlaybookConfig) GetAnsibleGitRefOk() (*string, bool)`

GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitRef

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleGitRef(v string)`

SetAnsibleGitRef sets AnsibleGitRef field to given value.

### HasAnsibleGitRef

`func (o *TaskAnsiblePlaybookConfig) HasAnsibleGitRef() bool`

HasAnsibleGitRef returns a boolean if a field has been set.

### SetAnsibleGitRefNil

`func (o *TaskAnsiblePlaybookConfig) SetAnsibleGitRefNil(b bool)`

 SetAnsibleGitRefNil sets the value for AnsibleGitRef to be an explicit nil

### UnsetAnsibleGitRef
`func (o *TaskAnsiblePlaybookConfig) UnsetAnsibleGitRef()`

UnsetAnsibleGitRef ensures that no value is present for AnsibleGitRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


