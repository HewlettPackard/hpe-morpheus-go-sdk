# CheckMysqlConfig

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

## Methods

### NewCheckMysqlConfig

`func NewCheckMysqlConfig(dbPort string, dbName string, dbUser string, dbHost string, dbQuery string, dbPassword string, ) *CheckMysqlConfig`

NewCheckMysqlConfig instantiates a new CheckMysqlConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckMysqlConfigWithDefaults

`func NewCheckMysqlConfigWithDefaults() *CheckMysqlConfig`

NewCheckMysqlConfigWithDefaults instantiates a new CheckMysqlConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbPort

`func (o *CheckMysqlConfig) GetDbPort() string`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *CheckMysqlConfig) GetDbPortOk() (*string, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *CheckMysqlConfig) SetDbPort(v string)`

SetDbPort sets DbPort field to given value.


### GetDbName

`func (o *CheckMysqlConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *CheckMysqlConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *CheckMysqlConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbUser

`func (o *CheckMysqlConfig) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *CheckMysqlConfig) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *CheckMysqlConfig) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbHost

`func (o *CheckMysqlConfig) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *CheckMysqlConfig) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *CheckMysqlConfig) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetCheckOperator

`func (o *CheckMysqlConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *CheckMysqlConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *CheckMysqlConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *CheckMysqlConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetDbQuery

`func (o *CheckMysqlConfig) GetDbQuery() string`

GetDbQuery returns the DbQuery field if non-nil, zero value otherwise.

### GetDbQueryOk

`func (o *CheckMysqlConfig) GetDbQueryOk() (*string, bool)`

GetDbQueryOk returns a tuple with the DbQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQuery

`func (o *CheckMysqlConfig) SetDbQuery(v string)`

SetDbQuery sets DbQuery field to given value.


### GetCheckResult

`func (o *CheckMysqlConfig) GetCheckResult() int64`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *CheckMysqlConfig) GetCheckResultOk() (*int64, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *CheckMysqlConfig) SetCheckResult(v int64)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *CheckMysqlConfig) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetDbPassword

`func (o *CheckMysqlConfig) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *CheckMysqlConfig) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *CheckMysqlConfig) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.


### GetDbPasswordHash

`func (o *CheckMysqlConfig) GetDbPasswordHash() string`

GetDbPasswordHash returns the DbPasswordHash field if non-nil, zero value otherwise.

### GetDbPasswordHashOk

`func (o *CheckMysqlConfig) GetDbPasswordHashOk() (*string, bool)`

GetDbPasswordHashOk returns a tuple with the DbPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPasswordHash

`func (o *CheckMysqlConfig) SetDbPasswordHash(v string)`

SetDbPasswordHash sets DbPasswordHash field to given value.

### HasDbPasswordHash

`func (o *CheckMysqlConfig) HasDbPasswordHash() bool`

HasDbPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


