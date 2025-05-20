# AddChecksRequestCheckOneOf

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

### NewAddChecksRequestCheckOneOf

`func NewAddChecksRequestCheckOneOf() *AddChecksRequestCheckOneOf`

NewAddChecksRequestCheckOneOf instantiates a new AddChecksRequestCheckOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChecksRequestCheckOneOfWithDefaults

`func NewAddChecksRequestCheckOneOfWithDefaults() *AddChecksRequestCheckOneOf`

NewAddChecksRequestCheckOneOfWithDefaults instantiates a new AddChecksRequestCheckOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddChecksRequestCheckOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddChecksRequestCheckOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddChecksRequestCheckOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddChecksRequestCheckOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddChecksRequestCheckOneOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddChecksRequestCheckOneOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddChecksRequestCheckOneOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddChecksRequestCheckOneOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCheckType

`func (o *AddChecksRequestCheckOneOf) GetCheckType() AddChecksRequestCheckOneOfCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *AddChecksRequestCheckOneOf) GetCheckTypeOk() (*AddChecksRequestCheckOneOfCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *AddChecksRequestCheckOneOf) SetCheckType(v AddChecksRequestCheckOneOfCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *AddChecksRequestCheckOneOf) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetCheckInterval

`func (o *AddChecksRequestCheckOneOf) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *AddChecksRequestCheckOneOf) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *AddChecksRequestCheckOneOf) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *AddChecksRequestCheckOneOf) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *AddChecksRequestCheckOneOf) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddChecksRequestCheckOneOf) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddChecksRequestCheckOneOf) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddChecksRequestCheckOneOf) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *AddChecksRequestCheckOneOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddChecksRequestCheckOneOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddChecksRequestCheckOneOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddChecksRequestCheckOneOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *AddChecksRequestCheckOneOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddChecksRequestCheckOneOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddChecksRequestCheckOneOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddChecksRequestCheckOneOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfig

`func (o *AddChecksRequestCheckOneOf) GetConfig() AddChecksRequestCheckOneOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddChecksRequestCheckOneOf) GetConfigOk() (*AddChecksRequestCheckOneOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddChecksRequestCheckOneOf) SetConfig(v AddChecksRequestCheckOneOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddChecksRequestCheckOneOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


