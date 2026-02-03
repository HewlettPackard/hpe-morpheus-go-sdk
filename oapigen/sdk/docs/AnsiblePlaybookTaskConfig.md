# AnsiblePlaybookTaskConfig

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

### NewAnsiblePlaybookTaskConfig

`func NewAnsiblePlaybookTaskConfig() *AnsiblePlaybookTaskConfig`

NewAnsiblePlaybookTaskConfig instantiates a new AnsiblePlaybookTaskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsiblePlaybookTaskConfigWithDefaults

`func NewAnsiblePlaybookTaskConfigWithDefaults() *AnsiblePlaybookTaskConfig`

NewAnsiblePlaybookTaskConfigWithDefaults instantiates a new AnsiblePlaybookTaskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleOptions

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleOptions() string`

GetAnsibleOptions returns the AnsibleOptions field if non-nil, zero value otherwise.

### GetAnsibleOptionsOk

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleOptionsOk() (*string, bool)`

GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleOptions

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleOptions(v string)`

SetAnsibleOptions sets AnsibleOptions field to given value.

### HasAnsibleOptions

`func (o *AnsiblePlaybookTaskConfig) HasAnsibleOptions() bool`

HasAnsibleOptions returns a boolean if a field has been set.

### SetAnsibleOptionsNil

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleOptionsNil(b bool)`

 SetAnsibleOptionsNil sets the value for AnsibleOptions to be an explicit nil

### UnsetAnsibleOptions
`func (o *AnsiblePlaybookTaskConfig) UnsetAnsibleOptions()`

UnsetAnsibleOptions ensures that no value is present for AnsibleOptions, not even an explicit nil
### GetAnsiblePlaybook

`func (o *AnsiblePlaybookTaskConfig) GetAnsiblePlaybook() string`

GetAnsiblePlaybook returns the AnsiblePlaybook field if non-nil, zero value otherwise.

### GetAnsiblePlaybookOk

`func (o *AnsiblePlaybookTaskConfig) GetAnsiblePlaybookOk() (*string, bool)`

GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybook

`func (o *AnsiblePlaybookTaskConfig) SetAnsiblePlaybook(v string)`

SetAnsiblePlaybook sets AnsiblePlaybook field to given value.

### HasAnsiblePlaybook

`func (o *AnsiblePlaybookTaskConfig) HasAnsiblePlaybook() bool`

HasAnsiblePlaybook returns a boolean if a field has been set.

### SetAnsiblePlaybookNil

`func (o *AnsiblePlaybookTaskConfig) SetAnsiblePlaybookNil(b bool)`

 SetAnsiblePlaybookNil sets the value for AnsiblePlaybook to be an explicit nil

### UnsetAnsiblePlaybook
`func (o *AnsiblePlaybookTaskConfig) UnsetAnsiblePlaybook()`

UnsetAnsiblePlaybook ensures that no value is present for AnsiblePlaybook, not even an explicit nil
### GetSshKey

`func (o *AnsiblePlaybookTaskConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *AnsiblePlaybookTaskConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *AnsiblePlaybookTaskConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *AnsiblePlaybookTaskConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *AnsiblePlaybookTaskConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *AnsiblePlaybookTaskConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetPort

`func (o *AnsiblePlaybookTaskConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AnsiblePlaybookTaskConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AnsiblePlaybookTaskConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *AnsiblePlaybookTaskConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AnsiblePlaybookTaskConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AnsiblePlaybookTaskConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetLocalScriptGitRef

`func (o *AnsiblePlaybookTaskConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *AnsiblePlaybookTaskConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *AnsiblePlaybookTaskConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *AnsiblePlaybookTaskConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *AnsiblePlaybookTaskConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *AnsiblePlaybookTaskConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetPassword

`func (o *AnsiblePlaybookTaskConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AnsiblePlaybookTaskConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AnsiblePlaybookTaskConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AnsiblePlaybookTaskConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AnsiblePlaybookTaskConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AnsiblePlaybookTaskConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *AnsiblePlaybookTaskConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AnsiblePlaybookTaskConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AnsiblePlaybookTaskConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AnsiblePlaybookTaskConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *AnsiblePlaybookTaskConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *AnsiblePlaybookTaskConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetLocalScriptGitId

`func (o *AnsiblePlaybookTaskConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *AnsiblePlaybookTaskConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *AnsiblePlaybookTaskConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *AnsiblePlaybookTaskConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *AnsiblePlaybookTaskConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *AnsiblePlaybookTaskConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetAnsibleGitId

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleGitId() string`

GetAnsibleGitId returns the AnsibleGitId field if non-nil, zero value otherwise.

### GetAnsibleGitIdOk

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleGitIdOk() (*string, bool)`

GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitId

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleGitId(v string)`

SetAnsibleGitId sets AnsibleGitId field to given value.

### HasAnsibleGitId

`func (o *AnsiblePlaybookTaskConfig) HasAnsibleGitId() bool`

HasAnsibleGitId returns a boolean if a field has been set.

### GetHost

`func (o *AnsiblePlaybookTaskConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AnsiblePlaybookTaskConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AnsiblePlaybookTaskConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AnsiblePlaybookTaskConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AnsiblePlaybookTaskConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AnsiblePlaybookTaskConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetAnsibleSkipTags

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleSkipTags() string`

GetAnsibleSkipTags returns the AnsibleSkipTags field if non-nil, zero value otherwise.

### GetAnsibleSkipTagsOk

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleSkipTagsOk() (*string, bool)`

GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleSkipTags

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleSkipTags(v string)`

SetAnsibleSkipTags sets AnsibleSkipTags field to given value.

### HasAnsibleSkipTags

`func (o *AnsiblePlaybookTaskConfig) HasAnsibleSkipTags() bool`

HasAnsibleSkipTags returns a boolean if a field has been set.

### SetAnsibleSkipTagsNil

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleSkipTagsNil(b bool)`

 SetAnsibleSkipTagsNil sets the value for AnsibleSkipTags to be an explicit nil

### UnsetAnsibleSkipTags
`func (o *AnsiblePlaybookTaskConfig) UnsetAnsibleSkipTags()`

UnsetAnsibleSkipTags ensures that no value is present for AnsibleSkipTags, not even an explicit nil
### GetAnsibleTags

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleTags() string`

GetAnsibleTags returns the AnsibleTags field if non-nil, zero value otherwise.

### GetAnsibleTagsOk

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleTagsOk() (*string, bool)`

GetAnsibleTagsOk returns a tuple with the AnsibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTags

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleTags(v string)`

SetAnsibleTags sets AnsibleTags field to given value.

### HasAnsibleTags

`func (o *AnsiblePlaybookTaskConfig) HasAnsibleTags() bool`

HasAnsibleTags returns a boolean if a field has been set.

### SetAnsibleTagsNil

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleTagsNil(b bool)`

 SetAnsibleTagsNil sets the value for AnsibleTags to be an explicit nil

### UnsetAnsibleTags
`func (o *AnsiblePlaybookTaskConfig) UnsetAnsibleTags()`

UnsetAnsibleTags ensures that no value is present for AnsibleTags, not even an explicit nil
### GetUsername

`func (o *AnsiblePlaybookTaskConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnsiblePlaybookTaskConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnsiblePlaybookTaskConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AnsiblePlaybookTaskConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AnsiblePlaybookTaskConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AnsiblePlaybookTaskConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAnsibleGitRef

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleGitRef() string`

GetAnsibleGitRef returns the AnsibleGitRef field if non-nil, zero value otherwise.

### GetAnsibleGitRefOk

`func (o *AnsiblePlaybookTaskConfig) GetAnsibleGitRefOk() (*string, bool)`

GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitRef

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleGitRef(v string)`

SetAnsibleGitRef sets AnsibleGitRef field to given value.

### HasAnsibleGitRef

`func (o *AnsiblePlaybookTaskConfig) HasAnsibleGitRef() bool`

HasAnsibleGitRef returns a boolean if a field has been set.

### SetAnsibleGitRefNil

`func (o *AnsiblePlaybookTaskConfig) SetAnsibleGitRefNil(b bool)`

 SetAnsibleGitRefNil sets the value for AnsibleGitRef to be an explicit nil

### UnsetAnsibleGitRef
`func (o *AnsiblePlaybookTaskConfig) UnsetAnsibleGitRef()`

UnsetAnsibleGitRef ensures that no value is present for AnsibleGitRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


