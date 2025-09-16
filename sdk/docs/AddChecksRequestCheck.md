# AddChecksRequestCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the check | 
**CheckType** | **map[string]interface{}** | Check type you want to create | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**CheckInterval** | Pointer to **int32** | Number of milliseconds you want between check executions (minimum is 1 minute, depending on your subscription plan) | [optional] [default to 300000]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Config** | Pointer to [**SNMPCheckAllOfConfig**](SNMPCheckAllOfConfig.md) |  | [optional] 

## Methods

### NewAddChecksRequestCheck

`func NewAddChecksRequestCheck(name string, checkType map[string]interface{}, ) *AddChecksRequestCheck`

NewAddChecksRequestCheck instantiates a new AddChecksRequestCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChecksRequestCheckWithDefaults

`func NewAddChecksRequestCheckWithDefaults() *AddChecksRequestCheck`

NewAddChecksRequestCheckWithDefaults instantiates a new AddChecksRequestCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddChecksRequestCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddChecksRequestCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddChecksRequestCheck) SetName(v string)`

SetName sets Name field to given value.


### GetCheckType

`func (o *AddChecksRequestCheck) GetCheckType() map[string]interface{}`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *AddChecksRequestCheck) GetCheckTypeOk() (*map[string]interface{}, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *AddChecksRequestCheck) SetCheckType(v map[string]interface{})`

SetCheckType sets CheckType field to given value.


### GetDescription

`func (o *AddChecksRequestCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddChecksRequestCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddChecksRequestCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddChecksRequestCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCheckInterval

`func (o *AddChecksRequestCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *AddChecksRequestCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *AddChecksRequestCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *AddChecksRequestCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *AddChecksRequestCheck) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddChecksRequestCheck) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddChecksRequestCheck) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddChecksRequestCheck) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *AddChecksRequestCheck) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddChecksRequestCheck) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddChecksRequestCheck) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddChecksRequestCheck) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *AddChecksRequestCheck) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddChecksRequestCheck) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddChecksRequestCheck) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddChecksRequestCheck) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfig

`func (o *AddChecksRequestCheck) GetConfig() SNMPCheckAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddChecksRequestCheck) GetConfigOk() (*SNMPCheckAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddChecksRequestCheck) SetConfig(v SNMPCheckAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddChecksRequestCheck) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


