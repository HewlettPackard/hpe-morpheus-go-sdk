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
**TunnelOn** | Pointer to **string** | Turn &#x60;on&#x60; to enable checks through a proxy host | [optional] [default to "off"]
**SshHost** | Pointer to **string** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | Pointer to **int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | Pointer to **string** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | Pointer to **string** | Password for user, if not using key based authentication | [optional] 

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

### GetSshPassword

`func (o *MySqlCheckAllOfConfig) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *MySqlCheckAllOfConfig) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *MySqlCheckAllOfConfig) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *MySqlCheckAllOfConfig) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


