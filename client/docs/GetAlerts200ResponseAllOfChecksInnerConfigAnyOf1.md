# GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1

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

### NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1

`func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, ) *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1`

NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1WithDefaults

`func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1WithDefaults() *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1`

NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1WithDefaults instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.

### GetDbName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbQuery

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckOperator

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResult

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckResult() float32`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckResultOk() (*float32, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckResult(v float32)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.

### GetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTunnelOn() string`

GetTunnelOn returns the TunnelOn field if non-nil, zero value otherwise.

### GetTunnelOnOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTunnelOnOk() (*string, bool)`

GetTunnelOnOk returns a tuple with the TunnelOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetTunnelOn(v string)`

SetTunnelOn sets TunnelOn field to given value.

### HasTunnelOn

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasTunnelOn() bool`

HasTunnelOn returns a boolean if a field has been set.

### GetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


