# GetAlerts200ResponseAllOfChecksInnerConfigAnyOf

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

### NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf

`func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf(containerName string, ) *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf`

NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOfWithDefaults

`func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOfWithDefaults() *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf`

NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOfWithDefaults instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetExternalId

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


