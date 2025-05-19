# GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsHost** | **string** | Hostname or IP address of the Elasticsearch server | 
**EsPort** | **string** | Port to connect to the HTTP service | 
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

### NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf2

`func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf2(esHost string, esPort string, ) *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2`

NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf2 instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf2WithDefaults

`func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf2WithDefaults() *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2`

NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf2WithDefaults instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetEsHost() string`

GetEsHost returns the EsHost field if non-nil, zero value otherwise.

### GetEsHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetEsHostOk() (*string, bool)`

GetEsHostOk returns a tuple with the EsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetEsHost(v string)`

SetEsHost sets EsHost field to given value.


### GetEsPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetEsPort() string`

GetEsPort returns the EsPort field if non-nil, zero value otherwise.

### GetEsPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetEsPortOk() (*string, bool)`

GetEsPortOk returns a tuple with the EsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetEsPort(v string)`

SetEsPort sets EsPort field to given value.


### GetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf2) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


