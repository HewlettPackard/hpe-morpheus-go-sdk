# MySqlCheckAllOfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbPort** | **string** |  | 
**DbName** | **string** |  | 
**DbUser** | **string** |  | 
**DbHost** | **string** |  | 
**CheckOperator** | Pointer to **string** |  | [optional] 
**DbQuery** | **string** |  | 
**CheckResult** | Pointer to **int64** |  | [optional] 
**DbPassword** | **string** |  | 
**DbPasswordHash** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**CheckUser** | Pointer to **string** |  | [optional] 
**TunnelOn** | Pointer to **string** |  | [optional] 
**TextCheckOn** | Pointer to **string** |  | [optional] 
**CheckPassword** | Pointer to **string** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**WebTextMatch** | Pointer to **string** |  | [optional] 
**CheckPasswordHash** | Pointer to **string** |  | [optional] 

## Methods

### NewMySqlCheckAllOfConfig

`func NewMySqlCheckAllOfConfig(dbPort string, dbName string, dbUser string, dbHost string, dbQuery string, dbPassword string, ) *MySqlCheckAllOfConfig`

NewMySqlCheckAllOfConfig instantiates a new MySqlCheckAllOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySqlCheckAllOfConfigWithDefaults

`func NewMySqlCheckAllOfConfigWithDefaults() *MySqlCheckAllOfConfig`

NewMySqlCheckAllOfConfigWithDefaults instantiates a new MySqlCheckAllOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbPort

`func (o *MySqlCheckAllOfConfig) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *MySqlCheckAllOfConfig) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *MySqlCheckAllOfConfig) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbName

`func (o *MySqlCheckAllOfConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MySqlCheckAllOfConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MySqlCheckAllOfConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbUser

`func (o *MySqlCheckAllOfConfig) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *MySqlCheckAllOfConfig) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *MySqlCheckAllOfConfig) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbHost

`func (o *MySqlCheckAllOfConfig) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *MySqlCheckAllOfConfig) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *MySqlCheckAllOfConfig) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetCheckOperator

`func (o *MySqlCheckAllOfConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *MySqlCheckAllOfConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *MySqlCheckAllOfConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *MySqlCheckAllOfConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetDbQuery

`func (o *MySqlCheckAllOfConfig) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *MySqlCheckAllOfConfig) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *MySqlCheckAllOfConfig) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckResult

`func (o *MySqlCheckAllOfConfig) GetCheckResult() int64`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *MySqlCheckAllOfConfig) GetCheckResultOk() (*int64, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *MySqlCheckAllOfConfig) SetCheckResult(v int64)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *MySqlCheckAllOfConfig) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetDbPassword

`func (o *MySqlCheckAllOfConfig) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *MySqlCheckAllOfConfig) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *MySqlCheckAllOfConfig) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *MySqlCheckAllOfConfig) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *MySqlCheckAllOfConfig) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *MySqlCheckAllOfConfig) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *MySqlCheckAllOfConfig) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetSshPort

`func (o *MySqlCheckAllOfConfig) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *MySqlCheckAllOfConfig) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *MySqlCheckAllOfConfig) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *MySqlCheckAllOfConfig) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetCheckUser

`func (o *MySqlCheckAllOfConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *MySqlCheckAllOfConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *MySqlCheckAllOfConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *MySqlCheckAllOfConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTunnelOn

`func (o *MySqlCheckAllOfConfig) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *MySqlCheckAllOfConfig) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *MySqlCheckAllOfConfig) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *MySqlCheckAllOfConfig) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *MySqlCheckAllOfConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *MySqlCheckAllOfConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *MySqlCheckAllOfConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *MySqlCheckAllOfConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *MySqlCheckAllOfConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *MySqlCheckAllOfConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *MySqlCheckAllOfConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *MySqlCheckAllOfConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetSshHost

`func (o *MySqlCheckAllOfConfig) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *MySqlCheckAllOfConfig) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *MySqlCheckAllOfConfig) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *MySqlCheckAllOfConfig) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshUser

`func (o *MySqlCheckAllOfConfig) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *MySqlCheckAllOfConfig) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *MySqlCheckAllOfConfig) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *MySqlCheckAllOfConfig) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *MySqlCheckAllOfConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *MySqlCheckAllOfConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *MySqlCheckAllOfConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *MySqlCheckAllOfConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *MySqlCheckAllOfConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *MySqlCheckAllOfConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *MySqlCheckAllOfConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *MySqlCheckAllOfConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


