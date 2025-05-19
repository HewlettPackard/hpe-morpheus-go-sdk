# AddChecksRequestCheckOneOfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebMethod** | **string** | HTTP method to use for testing | 
**WebUrl** | **string** | Web URL you wish to use to run a check on | 
**IgnoreSSL** | Pointer to **bool** | Ignore SSL Errors | [optional] [default to false]
**CheckUser** | Pointer to **string** |  | [optional] 
**CheckPassword** | Pointer to **string** |  | [optional] 
**TextCheckOn** | Pointer to **string** |  | [optional] 
**WebTextMatch** | Pointer to **string** |  | [optional] 
**TunnelOn** | Pointer to **string** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** | Password for user, if not using key based authentication | [optional] 
**CheckPasswordHash** | Pointer to **string** |  | [optional] 

## Methods

### NewAddChecksRequestCheckOneOfConfig

`func NewAddChecksRequestCheckOneOfConfig(webMethod string, webUrl string, ) *AddChecksRequestCheckOneOfConfig`

NewAddChecksRequestCheckOneOfConfig instantiates a new AddChecksRequestCheckOneOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChecksRequestCheckOneOfConfigWithDefaults

`func NewAddChecksRequestCheckOneOfConfigWithDefaults() *AddChecksRequestCheckOneOfConfig`

NewAddChecksRequestCheckOneOfConfigWithDefaults instantiates a new AddChecksRequestCheckOneOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebMethod

`func (o *AddChecksRequestCheckOneOfConfig) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *AddChecksRequestCheckOneOfConfig) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *AddChecksRequestCheckOneOfConfig) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.


### GetWebUrl

`func (o *AddChecksRequestCheckOneOfConfig) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *AddChecksRequestCheckOneOfConfig) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *AddChecksRequestCheckOneOfConfig) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetIgnoreSSL

`func (o *AddChecksRequestCheckOneOfConfig) GetIgnoreSSL() bool`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *AddChecksRequestCheckOneOfConfig) GetIgnoreSSLOk() (*bool, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *AddChecksRequestCheckOneOfConfig) SetIgnoreSSL(v bool)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *AddChecksRequestCheckOneOfConfig) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetCheckUser

`func (o *AddChecksRequestCheckOneOfConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *AddChecksRequestCheckOneOfConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *AddChecksRequestCheckOneOfConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *AddChecksRequestCheckOneOfConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetCheckPassword

`func (o *AddChecksRequestCheckOneOfConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *AddChecksRequestCheckOneOfConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *AddChecksRequestCheckOneOfConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *AddChecksRequestCheckOneOfConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *AddChecksRequestCheckOneOfConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *AddChecksRequestCheckOneOfConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *AddChecksRequestCheckOneOfConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *AddChecksRequestCheckOneOfConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *AddChecksRequestCheckOneOfConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *AddChecksRequestCheckOneOfConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *AddChecksRequestCheckOneOfConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *AddChecksRequestCheckOneOfConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetTunnelOn

`func (o *AddChecksRequestCheckOneOfConfig) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *AddChecksRequestCheckOneOfConfig) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *AddChecksRequestCheckOneOfConfig) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *AddChecksRequestCheckOneOfConfig) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *AddChecksRequestCheckOneOfConfig) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *AddChecksRequestCheckOneOfConfig) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *AddChecksRequestCheckOneOfConfig) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *AddChecksRequestCheckOneOfConfig) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *AddChecksRequestCheckOneOfConfig) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *AddChecksRequestCheckOneOfConfig) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *AddChecksRequestCheckOneOfConfig) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *AddChecksRequestCheckOneOfConfig) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *AddChecksRequestCheckOneOfConfig) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *AddChecksRequestCheckOneOfConfig) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *AddChecksRequestCheckOneOfConfig) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *AddChecksRequestCheckOneOfConfig) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *AddChecksRequestCheckOneOfConfig) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddChecksRequestCheckOneOfConfig) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddChecksRequestCheckOneOfConfig) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddChecksRequestCheckOneOfConfig) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *AddChecksRequestCheckOneOfConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *AddChecksRequestCheckOneOfConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *AddChecksRequestCheckOneOfConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *AddChecksRequestCheckOneOfConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


