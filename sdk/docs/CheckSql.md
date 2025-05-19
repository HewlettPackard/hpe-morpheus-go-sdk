# CheckSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the check | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**CheckType** | Pointer to [**AddChecksRequestCheckOneOf1CheckType**](AddChecksRequestCheckOneOf1CheckType.md) |  | [optional] 
**CheckInterval** | Pointer to **int32** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Config** | Pointer to [**AddChecksRequestCheckOneOf1Config**](AddChecksRequestCheckOneOf1Config.md) |  | [optional] 

## Methods

### NewCheckSql

`func NewCheckSql() *CheckSql`

NewCheckSql instantiates a new CheckSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckSqlWithDefaults

`func NewCheckSqlWithDefaults() *CheckSql`

NewCheckSqlWithDefaults instantiates a new CheckSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckSql) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckSql) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckSql) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckSql) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CheckSql) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckSql) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckSql) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckSql) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCheckType

`func (o *CheckSql) GetCheckType() AddChecksRequestCheckOneOf1CheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *CheckSql) GetCheckTypeOk() (*AddChecksRequestCheckOneOf1CheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *CheckSql) SetCheckType(v AddChecksRequestCheckOneOf1CheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *CheckSql) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetCheckInterval

`func (o *CheckSql) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *CheckSql) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *CheckSql) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *CheckSql) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *CheckSql) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *CheckSql) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *CheckSql) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *CheckSql) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *CheckSql) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckSql) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckSql) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckSql) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *CheckSql) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CheckSql) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CheckSql) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CheckSql) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfig

`func (o *CheckSql) GetConfig() AddChecksRequestCheckOneOf1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CheckSql) GetConfigOk() (*AddChecksRequestCheckOneOf1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CheckSql) SetConfig(v AddChecksRequestCheckOneOf1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CheckSql) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


