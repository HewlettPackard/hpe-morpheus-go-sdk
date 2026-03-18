# UpdateChecks200ResponseAllOfCheckConfig

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

### NewUpdateChecks200ResponseAllOfCheckConfig

`func NewUpdateChecks200ResponseAllOfCheckConfig(webMethod string, webUrl string, dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, esHost string, esPort string, host string, port string, send string, responseMatch string, containerName string, ) *UpdateChecks200ResponseAllOfCheckConfig`

NewUpdateChecks200ResponseAllOfCheckConfig instantiates a new UpdateChecks200ResponseAllOfCheckConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChecks200ResponseAllOfCheckConfigWithDefaults

`func NewUpdateChecks200ResponseAllOfCheckConfigWithDefaults() *UpdateChecks200ResponseAllOfCheckConfig`

NewUpdateChecks200ResponseAllOfCheckConfigWithDefaults instantiates a new UpdateChecks200ResponseAllOfCheckConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebMethod

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.


### GetWebUrl

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetIgnoreSSL

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetIgnoreSSL() bool`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetIgnoreSSLOk() (*bool, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetIgnoreSSL(v bool)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetCheckUser

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetCheckPassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetDbHost

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbUser

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbPassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetDbName

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbQuery

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckOperator

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResult

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckResult() float32`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckResultOk() (*float32, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCheckResult(v float32)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetEsHost

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetEsHost() string`

GetEsHost returns the EsHost field if non-nil, zero value otherwise.

### GetEsHostOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetEsHostOk() (*string, bool)`

GetEsHostOk returns a tuple with the EsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsHost

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetEsHost(v string)`

SetEsHost sets EsHost field to given value.


### GetEsPort

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetEsPort() string`

GetEsPort returns the EsPort field if non-nil, zero value otherwise.

### GetEsPortOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetEsPortOk() (*string, bool)`

GetEsPortOk returns a tuple with the EsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsPort

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetEsPort(v string)`

SetEsPort sets EsPort field to given value.


### GetHost

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetPort(v string)`

SetPort sets Port field to given value.


### GetSend

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetSend(v string)`

SetSend sets Send field to given value.


### GetResponseMatch

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetResponseMatch() string`

GetResponseMatch returns the ResponseMatch field if non-nil, zero value otherwise.

### GetResponseMatchOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetResponseMatchOk() (*string, bool)`

GetResponseMatchOk returns a tuple with the ResponseMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMatch

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetResponseMatch(v string)`

SetResponseMatch sets ResponseMatch field to given value.


### GetContainerName

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetExternalId

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetOid

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetCheckResponse

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckResponse() string`

GetCheckResponse returns the CheckResponse field if non-nil, zero value otherwise.

### GetCheckResponseOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCheckResponseOk() (*string, bool)`

GetCheckResponseOk returns a tuple with the CheckResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponse

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCheckResponse(v string)`

SetCheckResponse sets CheckResponse field to given value.

### HasCheckResponse

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCheckResponse() bool`

HasCheckResponse returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommunity

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetAuth

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthpassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetAuthpassword() string`

GetAuthpassword returns the Authpassword field if non-nil, zero value otherwise.

### GetAuthpasswordOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetAuthpasswordOk() (*string, bool)`

GetAuthpasswordOk returns a tuple with the Authpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthpassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetAuthpassword(v string)`

SetAuthpassword sets Authpassword field to given value.

### HasAuthpassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasAuthpassword() bool`

HasAuthpassword returns a boolean if a field has been set.

### GetPriv

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetPriv() string`

GetPriv returns the Priv field if non-nil, zero value otherwise.

### GetPrivOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetPrivOk() (*string, bool)`

GetPrivOk returns a tuple with the Priv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriv

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetPriv(v string)`

SetPriv sets Priv field to given value.

### HasPriv

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasPriv() bool`

HasPriv returns a boolean if a field has been set.

### GetPrivpassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetPrivpassword() string`

GetPrivpassword returns the Privpassword field if non-nil, zero value otherwise.

### GetPrivpasswordOk

`func (o *UpdateChecks200ResponseAllOfCheckConfig) GetPrivpasswordOk() (*string, bool)`

GetPrivpasswordOk returns a tuple with the Privpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivpassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) SetPrivpassword(v string)`

SetPrivpassword sets Privpassword field to given value.

### HasPrivpassword

`func (o *UpdateChecks200ResponseAllOfCheckConfig) HasPrivpassword() bool`

HasPrivpassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


