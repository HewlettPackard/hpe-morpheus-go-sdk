# UpdateChecksRequestCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**CheckType** | [**AddChecksRequestCheckOneOf4CheckType**](AddChecksRequestCheckOneOf4CheckType.md) |  | 
**CheckInterval** | Pointer to **int32** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Config** | Pointer to [**AddChecksRequestCheckOneOf3Config**](AddChecksRequestCheckOneOf3Config.md) |  | [optional] 

## Methods

### NewUpdateChecksRequestCheck

`func NewUpdateChecksRequestCheck(name string, checkType AddChecksRequestCheckOneOf4CheckType, ) *UpdateChecksRequestCheck`

NewUpdateChecksRequestCheck instantiates a new UpdateChecksRequestCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChecksRequestCheckWithDefaults

`func NewUpdateChecksRequestCheckWithDefaults() *UpdateChecksRequestCheck`

NewUpdateChecksRequestCheckWithDefaults instantiates a new UpdateChecksRequestCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateChecksRequestCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateChecksRequestCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateChecksRequestCheck) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateChecksRequestCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateChecksRequestCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateChecksRequestCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateChecksRequestCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCheckType

`func (o *UpdateChecksRequestCheck) GetCheckType() AddChecksRequestCheckOneOf4CheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *UpdateChecksRequestCheck) GetCheckTypeOk() (*AddChecksRequestCheckOneOf4CheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *UpdateChecksRequestCheck) SetCheckType(v AddChecksRequestCheckOneOf4CheckType)`

SetCheckType sets CheckType field to given value.


### GetCheckInterval

`func (o *UpdateChecksRequestCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *UpdateChecksRequestCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *UpdateChecksRequestCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *UpdateChecksRequestCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *UpdateChecksRequestCheck) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateChecksRequestCheck) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateChecksRequestCheck) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateChecksRequestCheck) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *UpdateChecksRequestCheck) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateChecksRequestCheck) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateChecksRequestCheck) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateChecksRequestCheck) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *UpdateChecksRequestCheck) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateChecksRequestCheck) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateChecksRequestCheck) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateChecksRequestCheck) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateChecksRequestCheck) GetConfig() AddChecksRequestCheckOneOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateChecksRequestCheck) GetConfigOk() (*AddChecksRequestCheckOneOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateChecksRequestCheck) SetConfig(v AddChecksRequestCheckOneOf3Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateChecksRequestCheck) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


