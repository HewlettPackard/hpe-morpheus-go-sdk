# ElasticSearchConfig1

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

### NewElasticSearchConfig1

`func NewElasticSearchConfig1(esHost string, esPort string, ) *ElasticSearchConfig1`

NewElasticSearchConfig1 instantiates a new ElasticSearchConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticSearchConfig1WithDefaults

`func NewElasticSearchConfig1WithDefaults() *ElasticSearchConfig1`

NewElasticSearchConfig1WithDefaults instantiates a new ElasticSearchConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsHost

`func (o *ElasticSearchConfig1) GetEsHost() string`

GetEsHost returns the EsHost field if non-nil, zero value otherwise.

### GetEsHostOk

`func (o *ElasticSearchConfig1) GetEsHostOk() (*string, bool)`

GetEsHostOk returns a tuple with the EsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsHost

`func (o *ElasticSearchConfig1) SetEsHost(v string)`

SetEsHost sets EsHost field to given value.


### GetEsPort

`func (o *ElasticSearchConfig1) GetEsPort() string`

GetEsPort returns the EsPort field if non-nil, zero value otherwise.

### GetEsPortOk

`func (o *ElasticSearchConfig1) GetEsPortOk() (*string, bool)`

GetEsPortOk returns a tuple with the EsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsPort

`func (o *ElasticSearchConfig1) SetEsPort(v string)`

SetEsPort sets EsPort field to given value.


### GetCheckUser

`func (o *ElasticSearchConfig1) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *ElasticSearchConfig1) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *ElasticSearchConfig1) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *ElasticSearchConfig1) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *ElasticSearchConfig1) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *ElasticSearchConfig1) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *ElasticSearchConfig1) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *ElasticSearchConfig1) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *ElasticSearchConfig1) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *ElasticSearchConfig1) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *ElasticSearchConfig1) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *ElasticSearchConfig1) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *ElasticSearchConfig1) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *ElasticSearchConfig1) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *ElasticSearchConfig1) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *ElasticSearchConfig1) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *ElasticSearchConfig1) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *ElasticSearchConfig1) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *ElasticSearchConfig1) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *ElasticSearchConfig1) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *ElasticSearchConfig1) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *ElasticSearchConfig1) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *ElasticSearchConfig1) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *ElasticSearchConfig1) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *ElasticSearchConfig1) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ElasticSearchConfig1) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ElasticSearchConfig1) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ElasticSearchConfig1) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *ElasticSearchConfig1) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ElasticSearchConfig1) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ElasticSearchConfig1) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ElasticSearchConfig1) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *ElasticSearchConfig1) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *ElasticSearchConfig1) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *ElasticSearchConfig1) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *ElasticSearchConfig1) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *ElasticSearchConfig1) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ElasticSearchConfig1) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ElasticSearchConfig1) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ElasticSearchConfig1) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


