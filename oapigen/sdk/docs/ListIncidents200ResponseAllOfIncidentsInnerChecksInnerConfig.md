# ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig

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
**DbHost** | **string** | Hostname or IP address of the database | 
**DbPort** | **string** | Database Port (defaults to default port of DB type selected) | 
**DbUser** | **string** | Database username | 
**DbPassword** | **string** | Database password, (all check data is encrypted inside the database) | 
**DbPasswordHash** | Pointer to **string** | Database password hash | [optional] 
**DbName** | **string** | Database name you would like to connect to | 
**DbQuery** | **string** | Query to test | 
**CheckOperator** | Pointer to **string** | Operator to use when comparing returned value with the expected check response value. One of &#39;lt&#39;, &#39;equal&#39;, or &#39;gt&#39;. | [optional] 
**CheckResult** | Pointer to **float32** |  | [optional] 
**CheckPasswordHash** | Pointer to **string** |  | [optional] 
**EsHost** | **string** | Hostname or IP address of the Elasticsearch server | 
**EsPort** | **string** | Port to connect to the HTTP service | 
**Host** | **string** | Hostname or IP address of the SNMP network entity | 
**Port** | **string** | Port to connect to the SNMP entity | 
**Send** | **string** | Connection string you might want to send to the service | 
**ResponseMatch** | **string** | Response from the service to match against | 
**ContainerName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Oid** | Pointer to **string** | Object ID to get from the network entity | [optional] 
**CheckResponse** | Pointer to **string** | Value to use with the check operator | [optional] 
**Version** | Pointer to **string** | SNMP Version 1/2c/3 of snmp get command | [optional] 
**Community** | Pointer to **string** | Community string acts as simple user or ID password (only valid for v1/2c) | [optional] 
**Username** | Pointer to **string** | Username used with SNMPv3 auth/privacy protocol passwords | [optional] 
**SecurityLevel** | Pointer to **string** | Level of security for authentication and privacty | [optional] 
**Auth** | Pointer to **string** | Authentication protocol | [optional] 
**Authpassword** | Pointer to **string** |  | [optional] 
**Priv** | Pointer to **string** | Privacy protocol | [optional] 
**Privpassword** | Pointer to **string** |  | [optional] 

## Methods

### NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig

`func NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig(webMethod string, webUrl string, dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, esHost string, esPort string, host string, port string, send string, responseMatch string, containerName string, ) *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig`

NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig instantiates a new ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfigWithDefaults

`func NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfigWithDefaults() *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig`

NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfigWithDefaults instantiates a new ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebMethod

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.


### GetWebUrl

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetIgnoreSSL

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetIgnoreSSL() bool`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetIgnoreSSLOk() (*bool, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetIgnoreSSL(v bool)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetCheckUser

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetCheckPassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetDbHost

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbUser

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbPassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetDbName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbQuery

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckOperator

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResult

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckResult() float32`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckResultOk() (*float32, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCheckResult(v float32)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetEsHost

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetEsHost() string`

GetEsHost returns the EsHost field if non-nil, zero value otherwise.

### GetEsHostOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetEsHostOk() (*string, bool)`

GetEsHostOk returns a tuple with the EsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsHost

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetEsHost(v string)`

SetEsHost sets EsHost field to given value.


### GetEsPort

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetEsPort() string`

GetEsPort returns the EsPort field if non-nil, zero value otherwise.

### GetEsPortOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetEsPortOk() (*string, bool)`

GetEsPortOk returns a tuple with the EsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsPort

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetEsPort(v string)`

SetEsPort sets EsPort field to given value.


### GetHost

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetPort(v string)`

SetPort sets Port field to given value.


### GetSend

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetSend(v string)`

SetSend sets Send field to given value.


### GetResponseMatch

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.


### GetContainerName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetExternalId

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetOid

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetCheckResponse

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckResponse() string`

GetCheckResponse returns the CheckResponse field if non-nil, zero value otherwise.

### GetCheckResponseOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCheckResponseOk() (*string, bool)`

GetCheckResponseOk returns a tuple with the CheckResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponse

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCheckResponse(v string)`

SetCheckResponse sets CheckResponse field to given value.

### HasCheckResponse

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCheckResponse() bool`

HasCheckResponse returns a boolean if a field has been set.

### GetVersion

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommunity

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetUsername

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetAuth

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthpassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetAuthpassword() string`

GetAuthpassword returns the Authpassword field if non-nil, zero value otherwise.

### GetAuthpasswordOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetAuthpasswordOk() (*string, bool)`

GetAuthpasswordOk returns a tuple with the Authpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthpassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetAuthpassword(v string)`

SetAuthpassword sets Authpassword field to given value.

### HasAuthpassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasAuthpassword() bool`

HasAuthpassword returns a boolean if a field has been set.

### GetPriv

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetPriv() string`

GetPriv returns the Priv field if non-nil, zero value otherwise.

### GetPrivOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetPrivOk() (*string, bool)`

GetPrivOk returns a tuple with the Priv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriv

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetPriv(v string)`

SetPriv sets Priv field to given value.

### HasPriv

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasPriv() bool`

HasPriv returns a boolean if a field has been set.

### GetPrivpassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetPrivpassword() string`

GetPrivpassword returns the Privpassword field if non-nil, zero value otherwise.

### GetPrivpasswordOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) GetPrivpasswordOk() (*string, bool)`

GetPrivpasswordOk returns a tuple with the Privpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivpassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) SetPrivpassword(v string)`

SetPrivpassword sets Privpassword field to given value.

### HasPrivpassword

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig) HasPrivpassword() bool`

HasPrivpassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


