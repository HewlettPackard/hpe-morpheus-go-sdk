# CheckPostgres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check | 
**Description** | Pointer to **NullableString** | Optional description field | [optional] 
**CheckInterval** | Pointer to **int32** | Number of milliseconds you want between check executions (minimum is 1 minute, depending on your subscription plan) | [optional] [default to 300000]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**CheckType** | Pointer to [**PostgresCheckAllOfCheckType**](PostgresCheckAllOfCheckType.md) |  | [optional] 
**Config** | Pointer to [**SqlCheckAllOfConfig**](SqlCheckAllOfConfig.md) |  | [optional] 

## Methods

### NewCheckPostgres

`func NewCheckPostgres(name string, ) *CheckPostgres`

NewCheckPostgres instantiates a new CheckPostgres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPostgresWithDefaults

`func NewCheckPostgresWithDefaults() *CheckPostgres`

NewCheckPostgresWithDefaults instantiates a new CheckPostgres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckPostgres) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckPostgres) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckPostgres) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CheckPostgres) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckPostgres) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckPostgres) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckPostgres) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CheckPostgres) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CheckPostgres) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCheckInterval

`func (o *CheckPostgres) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *CheckPostgres) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *CheckPostgres) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *CheckPostgres) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *CheckPostgres) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *CheckPostgres) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *CheckPostgres) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *CheckPostgres) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *CheckPostgres) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckPostgres) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckPostgres) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckPostgres) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *CheckPostgres) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CheckPostgres) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CheckPostgres) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CheckPostgres) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCheckType

`func (o *CheckPostgres) GetCheckType() PostgresCheckAllOfCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *CheckPostgres) GetCheckTypeOk() (*PostgresCheckAllOfCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *CheckPostgres) SetCheckType(v PostgresCheckAllOfCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *CheckPostgres) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *CheckPostgres) GetConfig() SqlCheckAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CheckPostgres) GetConfigOk() (*SqlCheckAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CheckPostgres) SetConfig(v SqlCheckAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CheckPostgres) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


