# CheckSqlConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbHost** | **string** | Hostname or IP address of the database | 
**DbPort** | **string** | Database Port (defaults to default port of DB type selected) | 
**DbUser** | **string** | Database username | 
**DbPassword** | **string** | Database password, (all check data is encrypted inside the database) | 
**DbPasswordHash** | Pointer to **string** | Database password hash | [optional] 
**DbName** | **string** | Database name you would like to connect to | 
**DbQuery** | **string** | Query to test | 
**CheckOperator** | Pointer to **string** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional] 
**CheckResult** | Pointer to **float32** |  | [optional] 
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

### NewCheckSqlConfig

`func NewCheckSqlConfig(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, ) *CheckSqlConfig`

NewCheckSqlConfig instantiates a new CheckSqlConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckSqlConfigWithDefaults

`func NewCheckSqlConfigWithDefaults() *CheckSqlConfig`

NewCheckSqlConfigWithDefaults instantiates a new CheckSqlConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbHost

`func (o *CheckSqlConfig) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *CheckSqlConfig) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *CheckSqlConfig) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *CheckSqlConfig) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *CheckSqlConfig) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *CheckSqlConfig) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbUser

`func (o *CheckSqlConfig) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *CheckSqlConfig) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *CheckSqlConfig) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbPassword

`func (o *CheckSqlConfig) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *CheckSqlConfig) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *CheckSqlConfig) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *CheckSqlConfig) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *CheckSqlConfig) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *CheckSqlConfig) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *CheckSqlConfig) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetDbName

`func (o *CheckSqlConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *CheckSqlConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *CheckSqlConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbQuery

`func (o *CheckSqlConfig) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *CheckSqlConfig) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *CheckSqlConfig) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckOperator

`func (o *CheckSqlConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *CheckSqlConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *CheckSqlConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *CheckSqlConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResult

`func (o *CheckSqlConfig) GetCheckResult() float32`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *CheckSqlConfig) GetCheckResultOk() (*float32, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *CheckSqlConfig) SetCheckResult(v float32)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *CheckSqlConfig) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCheckUser

`func (o *CheckSqlConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *CheckSqlConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *CheckSqlConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *CheckSqlConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *CheckSqlConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *CheckSqlConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *CheckSqlConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *CheckSqlConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *CheckSqlConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *CheckSqlConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *CheckSqlConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *CheckSqlConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *CheckSqlConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *CheckSqlConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *CheckSqlConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *CheckSqlConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *CheckSqlConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *CheckSqlConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *CheckSqlConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *CheckSqlConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *CheckSqlConfig) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *CheckSqlConfig) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *CheckSqlConfig) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *CheckSqlConfig) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *CheckSqlConfig) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *CheckSqlConfig) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *CheckSqlConfig) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *CheckSqlConfig) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *CheckSqlConfig) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *CheckSqlConfig) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *CheckSqlConfig) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *CheckSqlConfig) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *CheckSqlConfig) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *CheckSqlConfig) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *CheckSqlConfig) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *CheckSqlConfig) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *CheckSqlConfig) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *CheckSqlConfig) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *CheckSqlConfig) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *CheckSqlConfig) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


