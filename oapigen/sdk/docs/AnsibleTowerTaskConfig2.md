# AnsibleTowerTaskConfig2

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

### NewAnsibleTowerTaskConfig2

`func NewAnsibleTowerTaskConfig2() *AnsibleTowerTaskConfig2`

NewAnsibleTowerTaskConfig2 instantiates a new AnsibleTowerTaskConfig2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleTowerTaskConfig2WithDefaults

`func NewAnsibleTowerTaskConfig2WithDefaults() *AnsibleTowerTaskConfig2`

NewAnsibleTowerTaskConfig2WithDefaults instantiates a new AnsibleTowerTaskConfig2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *AnsibleTowerTaskConfig2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AnsibleTowerTaskConfig2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AnsibleTowerTaskConfig2) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AnsibleTowerTaskConfig2) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AnsibleTowerTaskConfig2) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AnsibleTowerTaskConfig2) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *AnsibleTowerTaskConfig2) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AnsibleTowerTaskConfig2) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AnsibleTowerTaskConfig2) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AnsibleTowerTaskConfig2) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *AnsibleTowerTaskConfig2) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *AnsibleTowerTaskConfig2) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetAnsibleTowerGitRef

`func (o *AnsibleTowerTaskConfig2) GetAnsibleTowerGitRef() string`

GetAnsibleTowerGitRef returns the AnsibleTowerGitRef field if non-nil, zero value otherwise.

### GetAnsibleTowerGitRefOk

`func (o *AnsibleTowerTaskConfig2) GetAnsibleTowerGitRefOk() (*string, bool)`

GetAnsibleTowerGitRefOk returns a tuple with the AnsibleTowerGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerGitRef

`func (o *AnsibleTowerTaskConfig2) SetAnsibleTowerGitRef(v string)`

SetAnsibleTowerGitRef sets AnsibleTowerGitRef field to given value.

### HasAnsibleTowerGitRef

`func (o *AnsibleTowerTaskConfig2) HasAnsibleTowerGitRef() bool`

HasAnsibleTowerGitRef returns a boolean if a field has been set.

### SetAnsibleTowerGitRefNil

`func (o *AnsibleTowerTaskConfig2) SetAnsibleTowerGitRefNil(b bool)`

 SetAnsibleTowerGitRefNil sets the value for AnsibleTowerGitRef to be an explicit nil

### UnsetAnsibleTowerGitRef
`func (o *AnsibleTowerTaskConfig2) UnsetAnsibleTowerGitRef()`

UnsetAnsibleTowerGitRef ensures that no value is present for AnsibleTowerGitRef, not even an explicit nil
### GetLocalScriptGitId

`func (o *AnsibleTowerTaskConfig2) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *AnsibleTowerTaskConfig2) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *AnsibleTowerTaskConfig2) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *AnsibleTowerTaskConfig2) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *AnsibleTowerTaskConfig2) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *AnsibleTowerTaskConfig2) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetHost

`func (o *AnsibleTowerTaskConfig2) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AnsibleTowerTaskConfig2) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AnsibleTowerTaskConfig2) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AnsibleTowerTaskConfig2) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AnsibleTowerTaskConfig2) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AnsibleTowerTaskConfig2) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *AnsibleTowerTaskConfig2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnsibleTowerTaskConfig2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnsibleTowerTaskConfig2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AnsibleTowerTaskConfig2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AnsibleTowerTaskConfig2) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AnsibleTowerTaskConfig2) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSshKey

`func (o *AnsibleTowerTaskConfig2) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *AnsibleTowerTaskConfig2) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *AnsibleTowerTaskConfig2) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *AnsibleTowerTaskConfig2) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *AnsibleTowerTaskConfig2) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *AnsibleTowerTaskConfig2) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetAnsibleGroup

`func (o *AnsibleTowerTaskConfig2) GetAnsibleGroup() string`

GetAnsibleGroup returns the AnsibleGroup field if non-nil, zero value otherwise.

### GetAnsibleGroupOk

`func (o *AnsibleTowerTaskConfig2) GetAnsibleGroupOk() (*string, bool)`

GetAnsibleGroupOk returns a tuple with the AnsibleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroup

`func (o *AnsibleTowerTaskConfig2) SetAnsibleGroup(v string)`

SetAnsibleGroup sets AnsibleGroup field to given value.

### HasAnsibleGroup

`func (o *AnsibleTowerTaskConfig2) HasAnsibleGroup() bool`

HasAnsibleGroup returns a boolean if a field has been set.

### SetAnsibleGroupNil

`func (o *AnsibleTowerTaskConfig2) SetAnsibleGroupNil(b bool)`

 SetAnsibleGroupNil sets the value for AnsibleGroup to be an explicit nil

### UnsetAnsibleGroup
`func (o *AnsibleTowerTaskConfig2) UnsetAnsibleGroup()`

UnsetAnsibleGroup ensures that no value is present for AnsibleGroup, not even an explicit nil
### GetAnsibleTowerExecuteMode

`func (o *AnsibleTowerTaskConfig2) GetAnsibleTowerExecuteMode() string`

GetAnsibleTowerExecuteMode returns the AnsibleTowerExecuteMode field if non-nil, zero value otherwise.

### GetAnsibleTowerExecuteModeOk

`func (o *AnsibleTowerTaskConfig2) GetAnsibleTowerExecuteModeOk() (*string, bool)`

GetAnsibleTowerExecuteModeOk returns a tuple with the AnsibleTowerExecuteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerExecuteMode

`func (o *AnsibleTowerTaskConfig2) SetAnsibleTowerExecuteMode(v string)`

SetAnsibleTowerExecuteMode sets AnsibleTowerExecuteMode field to given value.

### HasAnsibleTowerExecuteMode

`func (o *AnsibleTowerTaskConfig2) HasAnsibleTowerExecuteMode() bool`

HasAnsibleTowerExecuteMode returns a boolean if a field has been set.

### GetLocalScriptGitRef

`func (o *AnsibleTowerTaskConfig2) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *AnsibleTowerTaskConfig2) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *AnsibleTowerTaskConfig2) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *AnsibleTowerTaskConfig2) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *AnsibleTowerTaskConfig2) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *AnsibleTowerTaskConfig2) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetPort

`func (o *AnsibleTowerTaskConfig2) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AnsibleTowerTaskConfig2) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AnsibleTowerTaskConfig2) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *AnsibleTowerTaskConfig2) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AnsibleTowerTaskConfig2) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AnsibleTowerTaskConfig2) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


