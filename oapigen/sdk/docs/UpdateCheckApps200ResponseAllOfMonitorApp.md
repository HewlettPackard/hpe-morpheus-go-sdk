# UpdateCheckApps200ResponseAllOfMonitorApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**AddCheckApps200ResponseAllOfCheckAppAccount**](AddCheckApps200ResponseAllOfCheckAppAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**App** | Pointer to [**AddCheckApps200ResponseAllOfCheckAppApp**](AddCheckApps200ResponseAllOfCheckAppApp.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastCheckStatus** | Pointer to **NullableString** |  | [optional] 
**LastWarningDate** | Pointer to **NullableTime** |  | [optional] 
**LastErrorDate** | Pointer to **NullableTime** |  | [optional] 
**LastSuccessDate** | Pointer to **NullableTime** |  | [optional] 
**LastRunDate** | Pointer to **NullableTime** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**LastTimer** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**History** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**AddCheckApps200ResponseAllOfCheckAppCreatedBy**](AddCheckApps200ResponseAllOfCheckAppCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableString** |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 
**CheckGroups** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewUpdateCheckApps200ResponseAllOfMonitorApp

`func NewUpdateCheckApps200ResponseAllOfMonitorApp() *UpdateCheckApps200ResponseAllOfMonitorApp`

NewUpdateCheckApps200ResponseAllOfMonitorApp instantiates a new UpdateCheckApps200ResponseAllOfMonitorApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetAccount() AddCheckApps200ResponseAllOfCheckAppAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetAccountOk() (*AddCheckApps200ResponseAllOfCheckAppAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetAccount(v AddCheckApps200ResponseAllOfCheckAppAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApp

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetApp() AddCheckApps200ResponseAllOfCheckAppApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetAppOk() (*AddCheckApps200ResponseAllOfCheckAppApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetApp(v AddCheckApps200ResponseAllOfCheckAppApp)`

SetApp sets App field to given value.

### HasApp

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetName

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastTimer

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetSeverity

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetCreatedBy() AddCheckApps200ResponseAllOfCheckAppCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetCreatedByOk() (*AddCheckApps200ResponseAllOfCheckAppCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetCreatedBy(v AddCheckApps200ResponseAllOfCheckAppCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetChecks

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetCheckGroups() []int64`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) GetCheckGroupsOk() (*[]int64, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) SetCheckGroups(v []int64)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *UpdateCheckApps200ResponseAllOfMonitorApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


