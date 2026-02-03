# HTTPTaskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebPassword** | Pointer to **NullableString** |  | [optional] 
**WebPasswordHash** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**WebUser** | Pointer to **NullableString** |  | [optional] 
**WebBody** | Pointer to **NullableString** |  | [optional] 
**WebHeaders** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**IgnoreSSL** | Pointer to **NullableString** |  | [optional] 
**WebMethod** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHTTPTaskConfig

`func NewHTTPTaskConfig() *HTTPTaskConfig`

NewHTTPTaskConfig instantiates a new HTTPTaskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPTaskConfigWithDefaults

`func NewHTTPTaskConfigWithDefaults() *HTTPTaskConfig`

NewHTTPTaskConfigWithDefaults instantiates a new HTTPTaskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebPassword

`func (o *HTTPTaskConfig) GetWebPassword() string`

GetWebPassword returns the WebPassword field if non-nil, zero value otherwise.

### GetWebPasswordOk

`func (o *HTTPTaskConfig) GetWebPasswordOk() (*string, bool)`

GetWebPasswordOk returns a tuple with the WebPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPassword

`func (o *HTTPTaskConfig) SetWebPassword(v string)`

SetWebPassword sets WebPassword field to given value.

### HasWebPassword

`func (o *HTTPTaskConfig) HasWebPassword() bool`

HasWebPassword returns a boolean if a field has been set.

### SetWebPasswordNil

`func (o *HTTPTaskConfig) SetWebPasswordNil(b bool)`

 SetWebPasswordNil sets the value for WebPassword to be an explicit nil

### UnsetWebPassword
`func (o *HTTPTaskConfig) UnsetWebPassword()`

UnsetWebPassword ensures that no value is present for WebPassword, not even an explicit nil
### GetWebPasswordHash

`func (o *HTTPTaskConfig) GetWebPasswordHash() string`

GetWebPasswordHash returns the WebPasswordHash field if non-nil, zero value otherwise.

### GetWebPasswordHashOk

`func (o *HTTPTaskConfig) GetWebPasswordHashOk() (*string, bool)`

GetWebPasswordHashOk returns a tuple with the WebPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPasswordHash

`func (o *HTTPTaskConfig) SetWebPasswordHash(v string)`

SetWebPasswordHash sets WebPasswordHash field to given value.

### HasWebPasswordHash

`func (o *HTTPTaskConfig) HasWebPasswordHash() bool`

HasWebPasswordHash returns a boolean if a field has been set.

### SetWebPasswordHashNil

`func (o *HTTPTaskConfig) SetWebPasswordHashNil(b bool)`

 SetWebPasswordHashNil sets the value for WebPasswordHash to be an explicit nil

### UnsetWebPasswordHash
`func (o *HTTPTaskConfig) UnsetWebPasswordHash()`

UnsetWebPasswordHash ensures that no value is present for WebPasswordHash, not even an explicit nil
### GetLocalScriptGitId

`func (o *HTTPTaskConfig) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *HTTPTaskConfig) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *HTTPTaskConfig) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *HTTPTaskConfig) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *HTTPTaskConfig) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *HTTPTaskConfig) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetLocalScriptGitRef

`func (o *HTTPTaskConfig) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *HTTPTaskConfig) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *HTTPTaskConfig) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *HTTPTaskConfig) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *HTTPTaskConfig) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *HTTPTaskConfig) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetWebUser

`func (o *HTTPTaskConfig) GetWebUser() string`

GetWebUser returns the WebUser field if non-nil, zero value otherwise.

### GetWebUserOk

`func (o *HTTPTaskConfig) GetWebUserOk() (*string, bool)`

GetWebUserOk returns a tuple with the WebUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUser

`func (o *HTTPTaskConfig) SetWebUser(v string)`

SetWebUser sets WebUser field to given value.

### HasWebUser

`func (o *HTTPTaskConfig) HasWebUser() bool`

HasWebUser returns a boolean if a field has been set.

### SetWebUserNil

`func (o *HTTPTaskConfig) SetWebUserNil(b bool)`

 SetWebUserNil sets the value for WebUser to be an explicit nil

### UnsetWebUser
`func (o *HTTPTaskConfig) UnsetWebUser()`

UnsetWebUser ensures that no value is present for WebUser, not even an explicit nil
### GetWebBody

`func (o *HTTPTaskConfig) GetWebBody() string`

GetWebBody returns the WebBody field if non-nil, zero value otherwise.

### GetWebBodyOk

`func (o *HTTPTaskConfig) GetWebBodyOk() (*string, bool)`

GetWebBodyOk returns a tuple with the WebBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBody

`func (o *HTTPTaskConfig) SetWebBody(v string)`

SetWebBody sets WebBody field to given value.

### HasWebBody

`func (o *HTTPTaskConfig) HasWebBody() bool`

HasWebBody returns a boolean if a field has been set.

### SetWebBodyNil

`func (o *HTTPTaskConfig) SetWebBodyNil(b bool)`

 SetWebBodyNil sets the value for WebBody to be an explicit nil

### UnsetWebBody
`func (o *HTTPTaskConfig) UnsetWebBody()`

UnsetWebBody ensures that no value is present for WebBody, not even an explicit nil
### GetWebHeaders

`func (o *HTTPTaskConfig) GetWebHeaders() string`

GetWebHeaders returns the WebHeaders field if non-nil, zero value otherwise.

### GetWebHeadersOk

`func (o *HTTPTaskConfig) GetWebHeadersOk() (*string, bool)`

GetWebHeadersOk returns a tuple with the WebHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHeaders

`func (o *HTTPTaskConfig) SetWebHeaders(v string)`

SetWebHeaders sets WebHeaders field to given value.

### HasWebHeaders

`func (o *HTTPTaskConfig) HasWebHeaders() bool`

HasWebHeaders returns a boolean if a field has been set.

### GetPassword

`func (o *HTTPTaskConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HTTPTaskConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HTTPTaskConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HTTPTaskConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *HTTPTaskConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *HTTPTaskConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *HTTPTaskConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *HTTPTaskConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *HTTPTaskConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *HTTPTaskConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *HTTPTaskConfig) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *HTTPTaskConfig) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetUsername

`func (o *HTTPTaskConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HTTPTaskConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HTTPTaskConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HTTPTaskConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *HTTPTaskConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *HTTPTaskConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetIgnoreSSL

`func (o *HTTPTaskConfig) GetIgnoreSSL() string`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *HTTPTaskConfig) GetIgnoreSSLOk() (*string, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *HTTPTaskConfig) SetIgnoreSSL(v string)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *HTTPTaskConfig) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### SetIgnoreSSLNil

`func (o *HTTPTaskConfig) SetIgnoreSSLNil(b bool)`

 SetIgnoreSSLNil sets the value for IgnoreSSL to be an explicit nil

### UnsetIgnoreSSL
`func (o *HTTPTaskConfig) UnsetIgnoreSSL()`

UnsetIgnoreSSL ensures that no value is present for IgnoreSSL, not even an explicit nil
### GetWebMethod

`func (o *HTTPTaskConfig) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *HTTPTaskConfig) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *HTTPTaskConfig) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.

### HasWebMethod

`func (o *HTTPTaskConfig) HasWebMethod() bool`

HasWebMethod returns a boolean if a field has been set.

### SetWebMethodNil

`func (o *HTTPTaskConfig) SetWebMethodNil(b bool)`

 SetWebMethodNil sets the value for WebMethod to be an explicit nil

### UnsetWebMethod
`func (o *HTTPTaskConfig) UnsetWebMethod()`

UnsetWebMethod ensures that no value is present for WebMethod, not even an explicit nil
### GetWebUrl

`func (o *HTTPTaskConfig) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *HTTPTaskConfig) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *HTTPTaskConfig) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *HTTPTaskConfig) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *HTTPTaskConfig) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *HTTPTaskConfig) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetHost

`func (o *HTTPTaskConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HTTPTaskConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HTTPTaskConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HTTPTaskConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HTTPTaskConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HTTPTaskConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HTTPTaskConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HTTPTaskConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HTTPTaskConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *HTTPTaskConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HTTPTaskConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HTTPTaskConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSshKey

`func (o *HTTPTaskConfig) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *HTTPTaskConfig) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *HTTPTaskConfig) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *HTTPTaskConfig) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *HTTPTaskConfig) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *HTTPTaskConfig) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


