# SqlCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check | 
**Description** | Pointer to **NullableString** | Optional description field | [optional] 
**CheckInterval** | Pointer to **int32** | Number of milliseconds you want between check executions (minimum is 1 minute, depending on your subscription plan) | [optional] [default to 300000]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**CheckType** | Pointer to [**SqlCheckAllOfCheckType**](SqlCheckAllOfCheckType.md) |  | [optional] 
**Config** | Pointer to [**SqlCheckAllOfConfig**](SqlCheckAllOfConfig.md) |  | [optional] 

## Methods

### NewSqlCheck

`func NewSqlCheck(name string, ) *SqlCheck`

NewSqlCheck instantiates a new SqlCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlCheckWithDefaults

`func NewSqlCheckWithDefaults() *SqlCheck`

NewSqlCheckWithDefaults instantiates a new SqlCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlCheck) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SqlCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SqlCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SqlCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SqlCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SqlCheck) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SqlCheck) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCheckInterval

`func (o *SqlCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *SqlCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *SqlCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *SqlCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *SqlCheck) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *SqlCheck) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *SqlCheck) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *SqlCheck) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *SqlCheck) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SqlCheck) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SqlCheck) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SqlCheck) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *SqlCheck) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SqlCheck) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SqlCheck) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SqlCheck) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCheckType

`func (o *SqlCheck) GetCheckType() SqlCheckAllOfCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *SqlCheck) GetCheckTypeOk() (*SqlCheckAllOfCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *SqlCheck) SetCheckType(v SqlCheckAllOfCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *SqlCheck) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *SqlCheck) GetConfig() SqlCheckAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SqlCheck) GetConfigOk() (*SqlCheckAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SqlCheck) SetConfig(v SqlCheckAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SqlCheck) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


