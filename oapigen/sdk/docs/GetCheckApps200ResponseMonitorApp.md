# GetCheckApps200ResponseMonitorApp

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

### NewGetCheckApps200ResponseMonitorApp

`func NewGetCheckApps200ResponseMonitorApp() *GetCheckApps200ResponseMonitorApp`

NewGetCheckApps200ResponseMonitorApp instantiates a new GetCheckApps200ResponseMonitorApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCheckApps200ResponseMonitorApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCheckApps200ResponseMonitorApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCheckApps200ResponseMonitorApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCheckApps200ResponseMonitorApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetCheckApps200ResponseMonitorApp) GetAccount() AddCheckApps200ResponseAllOfCheckAppAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCheckApps200ResponseMonitorApp) GetAccountOk() (*AddCheckApps200ResponseAllOfCheckAppAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCheckApps200ResponseMonitorApp) SetAccount(v AddCheckApps200ResponseAllOfCheckAppAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCheckApps200ResponseMonitorApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *GetCheckApps200ResponseMonitorApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetCheckApps200ResponseMonitorApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetCheckApps200ResponseMonitorApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetCheckApps200ResponseMonitorApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApp

`func (o *GetCheckApps200ResponseMonitorApp) GetApp() AddCheckApps200ResponseAllOfCheckAppApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GetCheckApps200ResponseMonitorApp) GetAppOk() (*AddCheckApps200ResponseAllOfCheckAppApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GetCheckApps200ResponseMonitorApp) SetApp(v AddCheckApps200ResponseAllOfCheckAppApp)`

SetApp sets App field to given value.

### HasApp

`func (o *GetCheckApps200ResponseMonitorApp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetName

`func (o *GetCheckApps200ResponseMonitorApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCheckApps200ResponseMonitorApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCheckApps200ResponseMonitorApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCheckApps200ResponseMonitorApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCheckApps200ResponseMonitorApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCheckApps200ResponseMonitorApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCheckApps200ResponseMonitorApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCheckApps200ResponseMonitorApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCheckApps200ResponseMonitorApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCheckApps200ResponseMonitorApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *GetCheckApps200ResponseMonitorApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetCheckApps200ResponseMonitorApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetCheckApps200ResponseMonitorApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetCheckApps200ResponseMonitorApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *GetCheckApps200ResponseMonitorApp) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetCheckApps200ResponseMonitorApp) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetCheckApps200ResponseMonitorApp) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *GetCheckApps200ResponseMonitorApp) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *GetCheckApps200ResponseMonitorApp) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *GetCheckApps200ResponseMonitorApp) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetCheckApps200ResponseMonitorApp) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetCheckApps200ResponseMonitorApp) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *GetCheckApps200ResponseMonitorApp) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *GetCheckApps200ResponseMonitorApp) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *GetCheckApps200ResponseMonitorApp) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetCheckApps200ResponseMonitorApp) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetCheckApps200ResponseMonitorApp) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *GetCheckApps200ResponseMonitorApp) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *GetCheckApps200ResponseMonitorApp) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *GetCheckApps200ResponseMonitorApp) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetCheckApps200ResponseMonitorApp) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetCheckApps200ResponseMonitorApp) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *GetCheckApps200ResponseMonitorApp) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *GetCheckApps200ResponseMonitorApp) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *GetCheckApps200ResponseMonitorApp) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetCheckApps200ResponseMonitorApp) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetCheckApps200ResponseMonitorApp) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *GetCheckApps200ResponseMonitorApp) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *GetCheckApps200ResponseMonitorApp) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *GetCheckApps200ResponseMonitorApp) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetCheckApps200ResponseMonitorApp) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetCheckApps200ResponseMonitorApp) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GetCheckApps200ResponseMonitorApp) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetCheckApps200ResponseMonitorApp) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastTimer

`func (o *GetCheckApps200ResponseMonitorApp) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetCheckApps200ResponseMonitorApp) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetCheckApps200ResponseMonitorApp) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *GetCheckApps200ResponseMonitorApp) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetCheckApps200ResponseMonitorApp) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetCheckApps200ResponseMonitorApp) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetCheckApps200ResponseMonitorApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *GetCheckApps200ResponseMonitorApp) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetCheckApps200ResponseMonitorApp) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetCheckApps200ResponseMonitorApp) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GetCheckApps200ResponseMonitorApp) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *GetCheckApps200ResponseMonitorApp) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *GetCheckApps200ResponseMonitorApp) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetSeverity

`func (o *GetCheckApps200ResponseMonitorApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetCheckApps200ResponseMonitorApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetCheckApps200ResponseMonitorApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetCheckApps200ResponseMonitorApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetCheckApps200ResponseMonitorApp) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetCheckApps200ResponseMonitorApp) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetCheckApps200ResponseMonitorApp) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetCheckApps200ResponseMonitorApp) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetCheckApps200ResponseMonitorApp) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetCheckApps200ResponseMonitorApp) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetCheckApps200ResponseMonitorApp) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetCheckApps200ResponseMonitorApp) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetCheckApps200ResponseMonitorApp) GetCreatedBy() AddCheckApps200ResponseAllOfCheckAppCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCheckApps200ResponseMonitorApp) GetCreatedByOk() (*AddCheckApps200ResponseAllOfCheckAppCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCheckApps200ResponseMonitorApp) SetCreatedBy(v AddCheckApps200ResponseAllOfCheckAppCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCheckApps200ResponseMonitorApp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCheckApps200ResponseMonitorApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCheckApps200ResponseMonitorApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCheckApps200ResponseMonitorApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCheckApps200ResponseMonitorApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCheckApps200ResponseMonitorApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCheckApps200ResponseMonitorApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCheckApps200ResponseMonitorApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCheckApps200ResponseMonitorApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *GetCheckApps200ResponseMonitorApp) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetCheckApps200ResponseMonitorApp) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetCheckApps200ResponseMonitorApp) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetCheckApps200ResponseMonitorApp) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *GetCheckApps200ResponseMonitorApp) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *GetCheckApps200ResponseMonitorApp) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetChecks

`func (o *GetCheckApps200ResponseMonitorApp) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckApps200ResponseMonitorApp) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckApps200ResponseMonitorApp) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetCheckApps200ResponseMonitorApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetCheckApps200ResponseMonitorApp) GetCheckGroups() []int64`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetCheckApps200ResponseMonitorApp) GetCheckGroupsOk() (*[]int64, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetCheckApps200ResponseMonitorApp) SetCheckGroups(v []int64)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetCheckApps200ResponseMonitorApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


