# AddChecksRequestCheckOneOf1Config

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
**TunnelOn** | Pointer to **string** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** | Password for user, if not using key based authentication | [optional] 

## Methods

### NewAddChecksRequestCheckOneOf1Config

`func NewAddChecksRequestCheckOneOf1Config(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, ) *AddChecksRequestCheckOneOf1Config`

NewAddChecksRequestCheckOneOf1Config instantiates a new AddChecksRequestCheckOneOf1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChecksRequestCheckOneOf1ConfigWithDefaults

`func NewAddChecksRequestCheckOneOf1ConfigWithDefaults() *AddChecksRequestCheckOneOf1Config`

NewAddChecksRequestCheckOneOf1ConfigWithDefaults instantiates a new AddChecksRequestCheckOneOf1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbHost

`func (o *AddChecksRequestCheckOneOf1Config) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *AddChecksRequestCheckOneOf1Config) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *AddChecksRequestCheckOneOf1Config) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *AddChecksRequestCheckOneOf1Config) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbUser

`func (o *AddChecksRequestCheckOneOf1Config) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *AddChecksRequestCheckOneOf1Config) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbPassword

`func (o *AddChecksRequestCheckOneOf1Config) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *AddChecksRequestCheckOneOf1Config) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *AddChecksRequestCheckOneOf1Config) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *AddChecksRequestCheckOneOf1Config) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *AddChecksRequestCheckOneOf1Config) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetDbName

`func (o *AddChecksRequestCheckOneOf1Config) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *AddChecksRequestCheckOneOf1Config) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbQuery

`func (o *AddChecksRequestCheckOneOf1Config) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *AddChecksRequestCheckOneOf1Config) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *AddChecksRequestCheckOneOf1Config) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckOperator

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *AddChecksRequestCheckOneOf1Config) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *AddChecksRequestCheckOneOf1Config) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResult

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckResult() float32`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckResultOk() (*float32, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *AddChecksRequestCheckOneOf1Config) SetCheckResult(v float32)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *AddChecksRequestCheckOneOf1Config) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCheckUser

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *AddChecksRequestCheckOneOf1Config) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *AddChecksRequestCheckOneOf1Config) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *AddChecksRequestCheckOneOf1Config) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *AddChecksRequestCheckOneOf1Config) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *AddChecksRequestCheckOneOf1Config) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *AddChecksRequestCheckOneOf1Config) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *AddChecksRequestCheckOneOf1Config) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *AddChecksRequestCheckOneOf1Config) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *AddChecksRequestCheckOneOf1Config) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *AddChecksRequestCheckOneOf1Config) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *AddChecksRequestCheckOneOf1Config) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *AddChecksRequestCheckOneOf1Config) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *AddChecksRequestCheckOneOf1Config) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *AddChecksRequestCheckOneOf1Config) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *AddChecksRequestCheckOneOf1Config) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *AddChecksRequestCheckOneOf1Config) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *AddChecksRequestCheckOneOf1Config) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *AddChecksRequestCheckOneOf1Config) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *AddChecksRequestCheckOneOf1Config) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *AddChecksRequestCheckOneOf1Config) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *AddChecksRequestCheckOneOf1Config) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *AddChecksRequestCheckOneOf1Config) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *AddChecksRequestCheckOneOf1Config) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *AddChecksRequestCheckOneOf1Config) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *AddChecksRequestCheckOneOf1Config) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *AddChecksRequestCheckOneOf1Config) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *AddChecksRequestCheckOneOf1Config) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *AddChecksRequestCheckOneOf1Config) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *AddChecksRequestCheckOneOf1Config) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *AddChecksRequestCheckOneOf1Config) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *AddChecksRequestCheckOneOf1Config) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *AddChecksRequestCheckOneOf1Config) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddChecksRequestCheckOneOf1Config) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddChecksRequestCheckOneOf1Config) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddChecksRequestCheckOneOf1Config) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


