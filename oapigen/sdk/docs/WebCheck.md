# WebCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check | 
**Description** | Pointer to **NullableString** | Optional description field | [optional] 
**CheckInterval** | Pointer to **int32** | Number of milliseconds you want between check executions (minimum is 1 minute, depending on your subscription plan) | [optional] [default to 300000]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Active** | Pointer to **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**CheckType** | Pointer to [**WebCheckAllOfCheckType**](WebCheckAllOfCheckType.md) |  | [optional] 
**Config** | Pointer to [**WebCheckAllOfConfig**](WebCheckAllOfConfig.md) |  | [optional] 

## Methods

### NewWebCheck

`func NewWebCheck(name string, ) *WebCheck`

NewWebCheck instantiates a new WebCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebCheckWithDefaults

`func NewWebCheckWithDefaults() *WebCheck`

NewWebCheckWithDefaults instantiates a new WebCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebCheck) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WebCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebCheck) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebCheck) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCheckInterval

`func (o *WebCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *WebCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *WebCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *WebCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetInUptime

`func (o *WebCheck) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *WebCheck) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *WebCheck) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *WebCheck) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetActive

`func (o *WebCheck) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebCheck) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebCheck) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebCheck) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSeverity

`func (o *WebCheck) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WebCheck) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WebCheck) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WebCheck) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCheckType

`func (o *WebCheck) GetCheckType() WebCheckAllOfCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *WebCheck) GetCheckTypeOk() (*WebCheckAllOfCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *WebCheck) SetCheckType(v WebCheckAllOfCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *WebCheck) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *WebCheck) GetConfig() WebCheckAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebCheck) GetConfigOk() (*WebCheckAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebCheck) SetConfig(v WebCheckAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WebCheck) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


