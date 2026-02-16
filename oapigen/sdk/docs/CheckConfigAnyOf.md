# CheckConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** |  | 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**CheckUser** | Pointer to **string** |  | [optional] 
**TextCheckOn** | Pointer to **string** |  | [optional] 
**CheckPassword** | Pointer to **string** |  | [optional] 
**WebTextMatch** | Pointer to **string** |  | [optional] 
**CheckPasswordHash** | Pointer to **string** |  | [optional] 
**TunnelOn** | Pointer to **string** | Set to on to turn on tunneling | [optional] [default to "off"]
**SshHost** | Pointer to **string** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | Pointer to **int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | Pointer to **string** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | Pointer to **string** | Password for user, if not using key based authentication | [optional] 

## Methods

### NewCheckConfigAnyOf

`func NewCheckConfigAnyOf(containerName string, ) *CheckConfigAnyOf`

NewCheckConfigAnyOf instantiates a new CheckConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckConfigAnyOfWithDefaults

`func NewCheckConfigAnyOfWithDefaults() *CheckConfigAnyOf`

NewCheckConfigAnyOfWithDefaults instantiates a new CheckConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *CheckConfigAnyOf) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CheckConfigAnyOf) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CheckConfigAnyOf) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetExternalId

`func (o *CheckConfigAnyOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CheckConfigAnyOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CheckConfigAnyOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CheckConfigAnyOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CheckConfigAnyOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CheckConfigAnyOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCheckUser

`func (o *CheckConfigAnyOf) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *CheckConfigAnyOf) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *CheckConfigAnyOf) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *CheckConfigAnyOf) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *CheckConfigAnyOf) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *CheckConfigAnyOf) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *CheckConfigAnyOf) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *CheckConfigAnyOf) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *CheckConfigAnyOf) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *CheckConfigAnyOf) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *CheckConfigAnyOf) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *CheckConfigAnyOf) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *CheckConfigAnyOf) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *CheckConfigAnyOf) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *CheckConfigAnyOf) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *CheckConfigAnyOf) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *CheckConfigAnyOf) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *CheckConfigAnyOf) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *CheckConfigAnyOf) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *CheckConfigAnyOf) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *CheckConfigAnyOf) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *CheckConfigAnyOf) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *CheckConfigAnyOf) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *CheckConfigAnyOf) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *CheckConfigAnyOf) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *CheckConfigAnyOf) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *CheckConfigAnyOf) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *CheckConfigAnyOf) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *CheckConfigAnyOf) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *CheckConfigAnyOf) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *CheckConfigAnyOf) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *CheckConfigAnyOf) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *CheckConfigAnyOf) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *CheckConfigAnyOf) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *CheckConfigAnyOf) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *CheckConfigAnyOf) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *CheckConfigAnyOf) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *CheckConfigAnyOf) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *CheckConfigAnyOf) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *CheckConfigAnyOf) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


