# SqlCheck1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check | 
**Description** | Pointer to **NullableString** | Optional description field | [optional] 
**CheckInterval** | Pointer to **int32** | Number of milliseconds you want between check executions (minimum is 1 minute, depending on your subscription plan) | [optional] [default to 300000]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**CheckType** | Pointer to [**SqlCheck1AllOfCheckType**](SqlCheck1AllOfCheckType.md) |  | [optional] 
**Config** | Pointer to [**SqlCheck1AllOfConfig**](SqlCheck1AllOfConfig.md) |  | [optional] 

## Methods

### NewSqlCheck1

`func NewSqlCheck1(name string, ) *SqlCheck1`

NewSqlCheck1 instantiates a new SqlCheck1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlCheck1WithDefaults

`func NewSqlCheck1WithDefaults() *SqlCheck1`

NewSqlCheck1WithDefaults instantiates a new SqlCheck1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlCheck1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlCheck1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlCheck1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SqlCheck1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SqlCheck1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SqlCheck1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SqlCheck1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SqlCheck1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SqlCheck1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCheckInterval

`func (o *SqlCheck1) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *SqlCheck1) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *SqlCheck1) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *SqlCheck1) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *SqlCheck1) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *SqlCheck1) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *SqlCheck1) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *SqlCheck1) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *SqlCheck1) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SqlCheck1) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SqlCheck1) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SqlCheck1) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *SqlCheck1) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SqlCheck1) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SqlCheck1) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SqlCheck1) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCheckType

`func (o *SqlCheck1) GetCheckType() SqlCheck1AllOfCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *SqlCheck1) GetCheckTypeOk() (*SqlCheck1AllOfCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *SqlCheck1) SetCheckType(v SqlCheck1AllOfCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *SqlCheck1) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *SqlCheck1) GetConfig() SqlCheck1AllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SqlCheck1) GetConfigOk() (*SqlCheck1AllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SqlCheck1) SetConfig(v SqlCheck1AllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SqlCheck1) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


