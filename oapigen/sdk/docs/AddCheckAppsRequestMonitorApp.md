# AddCheckAppsRequestMonitorApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check app | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level of incidents that are created when this check fails | [optional] [default to "critical"]
**Active** | Pointer to **bool** | Used to determine if check app is active | [optional] [default to true]
**Checks** | Pointer to **[]int32** |  | [optional] 
**CheckGroups** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewAddCheckAppsRequestMonitorApp

`func NewAddCheckAppsRequestMonitorApp(name string, ) *AddCheckAppsRequestMonitorApp`

NewAddCheckAppsRequestMonitorApp instantiates a new AddCheckAppsRequestMonitorApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCheckAppsRequestMonitorAppWithDefaults

`func NewAddCheckAppsRequestMonitorAppWithDefaults() *AddCheckAppsRequestMonitorApp`

NewAddCheckAppsRequestMonitorAppWithDefaults instantiates a new AddCheckAppsRequestMonitorApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCheckAppsRequestMonitorApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCheckAppsRequestMonitorApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCheckAppsRequestMonitorApp) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCheckAppsRequestMonitorApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCheckAppsRequestMonitorApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCheckAppsRequestMonitorApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCheckAppsRequestMonitorApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInUptime

`func (o *AddCheckAppsRequestMonitorApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddCheckAppsRequestMonitorApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddCheckAppsRequestMonitorApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddCheckAppsRequestMonitorApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetSeverity

`func (o *AddCheckAppsRequestMonitorApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddCheckAppsRequestMonitorApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddCheckAppsRequestMonitorApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddCheckAppsRequestMonitorApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetActive

`func (o *AddCheckAppsRequestMonitorApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddCheckAppsRequestMonitorApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddCheckAppsRequestMonitorApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddCheckAppsRequestMonitorApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetChecks

`func (o *AddCheckAppsRequestMonitorApp) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *AddCheckAppsRequestMonitorApp) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *AddCheckAppsRequestMonitorApp) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *AddCheckAppsRequestMonitorApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *AddCheckAppsRequestMonitorApp) GetCheckGroups() []int32`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *AddCheckAppsRequestMonitorApp) GetCheckGroupsOk() (*[]int32, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *AddCheckAppsRequestMonitorApp) SetCheckGroups(v []int32)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *AddCheckAppsRequestMonitorApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


