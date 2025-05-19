# GetAlerts200ResponseAllOfChecksInnerConfig

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
**TunnelOn** | Pointer to **string** | Set to on to turn on tunneling | [optional] [default to "off"]
**SshHost** | Pointer to **string** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | Pointer to **int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | Pointer to **string** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | Pointer to **string** | Password for user, if not using key based authentication | [optional] 
**DbHost** | **string** | Hostname or IP address of the database | 
**DbPort** | **string** | Database Port (defaults to default port of DB type selected) | 
**DbUser** | **string** | Database username | 
**DbPassword** | **string** | Database password, (all check data is encrypted inside the database) | 
**DbPasswordHash** | Pointer to **string** | Database password hash | [optional] 
**DbName** | **string** | Database name you would like to connect to | 
**DbQuery** | **string** | Query to test | 
**CheckOperator** | Pointer to **string** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional] 
**CheckResult** | Pointer to **float32** |  | [optional] 
**CheckPasswordHash** | Pointer to **string** |  | [optional] 
**EsHost** | **string** | Hostname or IP address of the Elasticsearch server | 
**EsPort** | **string** | Port to connect to the HTTP service | 
**Host** | **string** | Hostname or IP address of the socket server | 
**Port** | **string** | TCP port | 
**Send** | **string** | Connection string you might want to send to the service | 
**ResponseMatch** | **string** | Response from the service to match against | 
**ContainerName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAlerts200ResponseAllOfChecksInnerConfig

`func NewGetAlerts200ResponseAllOfChecksInnerConfig(webMethod string, webUrl string, dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, esHost string, esPort string, host string, port string, send string, responseMatch string, containerName string, ) *GetAlerts200ResponseAllOfChecksInnerConfig`

NewGetAlerts200ResponseAllOfChecksInnerConfig instantiates a new GetAlerts200ResponseAllOfChecksInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfChecksInnerConfigWithDefaults

`func NewGetAlerts200ResponseAllOfChecksInnerConfigWithDefaults() *GetAlerts200ResponseAllOfChecksInnerConfig`

NewGetAlerts200ResponseAllOfChecksInnerConfigWithDefaults instantiates a new GetAlerts200ResponseAllOfChecksInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebMethod

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.


### GetWebUrl

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetIgnoreSSL

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetIgnoreSSL() bool`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetIgnoreSSLOk() (*bool, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetIgnoreSSL(v bool)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetDbHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetDbName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbQuery

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckOperator

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResult

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckResult() float32`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckResultOk() (*float32, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetCheckResult(v float32)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetEsHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetEsHost() string`

GetEsHost returns the EsHost field if non-nil, zero value otherwise.

### GetEsHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetEsHostOk() (*string, bool)`

GetEsHostOk returns a tuple with the EsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetEsHost(v string)`

SetEsHost sets EsHost field to given value.


### GetEsPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetEsPort() string`

GetEsPort returns the EsPort field if non-nil, zero value otherwise.

### GetEsPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetEsPortOk() (*string, bool)`

GetEsPortOk returns a tuple with the EsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetEsPort(v string)`

SetEsPort sets EsPort field to given value.


### GetHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetPort(v string)`

SetPort sets Port field to given value.


### GetSend

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetSend(v string)`

SetSend sets Send field to given value.


### GetResponseMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.


### GetContainerName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetExternalId

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetAlerts200ResponseAllOfChecksInnerConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


