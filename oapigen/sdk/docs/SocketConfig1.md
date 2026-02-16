# SocketConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname or IP address of the socket server | 
**Port** | **string** | TCP port | 
**Send** | **string** | Connection string you might want to send to the service | 
**ResponseMatch** | **string** | Response from the service to match against | 
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

### NewSocketConfig1

`func NewSocketConfig1(host string, port string, send string, responseMatch string, ) *SocketConfig1`

NewSocketConfig1 instantiates a new SocketConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocketConfig1WithDefaults

`func NewSocketConfig1WithDefaults() *SocketConfig1`

NewSocketConfig1WithDefaults instantiates a new SocketConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SocketConfig1) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SocketConfig1) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SocketConfig1) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SocketConfig1) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SocketConfig1) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SocketConfig1) SetPort(v string)`

SetPort sets Port field to given value.


### GetSend

`func (o *SocketConfig1) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *SocketConfig1) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *SocketConfig1) SetSend(v string)`

SetSend sets Send field to given value.


### GetResponseMatch

`func (o *SocketConfig1) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *SocketConfig1) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *SocketConfig1) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.


### GetCheckUser

`func (o *SocketConfig1) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *SocketConfig1) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *SocketConfig1) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *SocketConfig1) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *SocketConfig1) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *SocketConfig1) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *SocketConfig1) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *SocketConfig1) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *SocketConfig1) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *SocketConfig1) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *SocketConfig1) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *SocketConfig1) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *SocketConfig1) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *SocketConfig1) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *SocketConfig1) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *SocketConfig1) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *SocketConfig1) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *SocketConfig1) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *SocketConfig1) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *SocketConfig1) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *SocketConfig1) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *SocketConfig1) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *SocketConfig1) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *SocketConfig1) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *SocketConfig1) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *SocketConfig1) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *SocketConfig1) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *SocketConfig1) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *SocketConfig1) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *SocketConfig1) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *SocketConfig1) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *SocketConfig1) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *SocketConfig1) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *SocketConfig1) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *SocketConfig1) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *SocketConfig1) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *SocketConfig1) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *SocketConfig1) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *SocketConfig1) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *SocketConfig1) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


