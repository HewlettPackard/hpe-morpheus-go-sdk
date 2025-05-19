# CheckWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the check | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**CheckType** | Pointer to [**AddChecksRequestCheckOneOfCheckType**](AddChecksRequestCheckOneOfCheckType.md) |  | [optional] 
**CheckInterval** | Pointer to **int32** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Config** | Pointer to [**AddChecksRequestCheckOneOfConfig**](AddChecksRequestCheckOneOfConfig.md) |  | [optional] 

## Methods

### NewCheckWeb

`func NewCheckWeb() *CheckWeb`

NewCheckWeb instantiates a new CheckWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWebWithDefaults

`func NewCheckWebWithDefaults() *CheckWeb`

NewCheckWebWithDefaults instantiates a new CheckWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckWeb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckWeb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckWeb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckWeb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CheckWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCheckType

`func (o *CheckWeb) GetCheckType() AddChecksRequestCheckOneOfCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *CheckWeb) GetCheckTypeOk() (*AddChecksRequestCheckOneOfCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *CheckWeb) SetCheckType(v AddChecksRequestCheckOneOfCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *CheckWeb) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetCheckInterval

`func (o *CheckWeb) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *CheckWeb) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *CheckWeb) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *CheckWeb) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *CheckWeb) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *CheckWeb) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *CheckWeb) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *CheckWeb) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *CheckWeb) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckWeb) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckWeb) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckWeb) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *CheckWeb) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CheckWeb) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CheckWeb) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CheckWeb) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfig

`func (o *CheckWeb) GetConfig() AddChecksRequestCheckOneOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CheckWeb) GetConfigOk() (*AddChecksRequestCheckOneOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CheckWeb) SetConfig(v AddChecksRequestCheckOneOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CheckWeb) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


