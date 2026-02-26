# SocketConfig

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

## Methods

### NewSocketConfig

`func NewSocketConfig(host string, port string, send string, responseMatch string, ) *SocketConfig`

NewSocketConfig instantiates a new SocketConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocketConfigWithDefaults

`func NewSocketConfigWithDefaults() *SocketConfig`

NewSocketConfigWithDefaults instantiates a new SocketConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SocketConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SocketConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SocketConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SocketConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SocketConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SocketConfig) SetPort(v string)`

SetPort sets Port field to given value.


### GetSend

`func (o *SocketConfig) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *SocketConfig) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *SocketConfig) SetSend(v string)`

SetSend sets Send field to given value.


### GetResponseMatch

`func (o *SocketConfig) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *SocketConfig) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *SocketConfig) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.


### GetCheckUser

`func (o *SocketConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *SocketConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *SocketConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *SocketConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *SocketConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *SocketConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *SocketConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *SocketConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *SocketConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *SocketConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *SocketConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *SocketConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *SocketConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *SocketConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *SocketConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *SocketConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *SocketConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *SocketConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *SocketConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *SocketConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


