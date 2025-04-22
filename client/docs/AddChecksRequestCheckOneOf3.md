# AddChecksRequestCheckOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the check | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**CheckType** | Pointer to [**AddChecksRequestCheckOneOf3CheckType**](AddChecksRequestCheckOneOf3CheckType.md) |  | [optional] 
**CheckInterval** | Pointer to **int32** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Config** | Pointer to [**AddChecksRequestCheckOneOf3Config**](AddChecksRequestCheckOneOf3Config.md) |  | [optional] 

## Methods

### NewAddChecksRequestCheckOneOf3

`func NewAddChecksRequestCheckOneOf3() *AddChecksRequestCheckOneOf3`

NewAddChecksRequestCheckOneOf3 instantiates a new AddChecksRequestCheckOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChecksRequestCheckOneOf3WithDefaults

`func NewAddChecksRequestCheckOneOf3WithDefaults() *AddChecksRequestCheckOneOf3`

NewAddChecksRequestCheckOneOf3WithDefaults instantiates a new AddChecksRequestCheckOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddChecksRequestCheckOneOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddChecksRequestCheckOneOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddChecksRequestCheckOneOf3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddChecksRequestCheckOneOf3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddChecksRequestCheckOneOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddChecksRequestCheckOneOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddChecksRequestCheckOneOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddChecksRequestCheckOneOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCheckType

`func (o *AddChecksRequestCheckOneOf3) GetCheckType() AddChecksRequestCheckOneOf3CheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *AddChecksRequestCheckOneOf3) GetCheckTypeOk() (*AddChecksRequestCheckOneOf3CheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *AddChecksRequestCheckOneOf3) SetCheckType(v AddChecksRequestCheckOneOf3CheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *AddChecksRequestCheckOneOf3) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetCheckInterval

`func (o *AddChecksRequestCheckOneOf3) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *AddChecksRequestCheckOneOf3) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *AddChecksRequestCheckOneOf3) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *AddChecksRequestCheckOneOf3) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *AddChecksRequestCheckOneOf3) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddChecksRequestCheckOneOf3) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddChecksRequestCheckOneOf3) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddChecksRequestCheckOneOf3) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *AddChecksRequestCheckOneOf3) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddChecksRequestCheckOneOf3) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddChecksRequestCheckOneOf3) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddChecksRequestCheckOneOf3) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *AddChecksRequestCheckOneOf3) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddChecksRequestCheckOneOf3) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddChecksRequestCheckOneOf3) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddChecksRequestCheckOneOf3) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetConfig

`func (o *AddChecksRequestCheckOneOf3) GetConfig() AddChecksRequestCheckOneOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddChecksRequestCheckOneOf3) GetConfigOk() (*AddChecksRequestCheckOneOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddChecksRequestCheckOneOf3) SetConfig(v AddChecksRequestCheckOneOf3Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddChecksRequestCheckOneOf3) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


