# ListCheckApps200ResponseAllOfMonitorAppsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListCheckApps200ResponseAllOfMonitorAppsInnerAccount**](ListCheckApps200ResponseAllOfMonitorAppsInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**App** | Pointer to [**ListCheckApps200ResponseAllOfMonitorAppsInnerApp**](ListCheckApps200ResponseAllOfMonitorAppsInnerApp.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**ListCheckApps200ResponseAllOfMonitorAppsInnerCreatedBy**](ListCheckApps200ResponseAllOfMonitorAppsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableString** |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 
**CheckGroups** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewListCheckApps200ResponseAllOfMonitorAppsInner

`func NewListCheckApps200ResponseAllOfMonitorAppsInner() *ListCheckApps200ResponseAllOfMonitorAppsInner`

NewListCheckApps200ResponseAllOfMonitorAppsInner instantiates a new ListCheckApps200ResponseAllOfMonitorAppsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetAccount() ListCheckApps200ResponseAllOfMonitorAppsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetAccountOk() (*ListCheckApps200ResponseAllOfMonitorAppsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetAccount(v ListCheckApps200ResponseAllOfMonitorAppsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApp

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetApp() ListCheckApps200ResponseAllOfMonitorAppsInnerApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetAppOk() (*ListCheckApps200ResponseAllOfMonitorAppsInnerApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetApp(v ListCheckApps200ResponseAllOfMonitorAppsInnerApp)`

SetApp sets App field to given value.

### HasApp

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetName

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastTimer

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetSeverity

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetCreatedBy() ListCheckApps200ResponseAllOfMonitorAppsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetCreatedByOk() (*ListCheckApps200ResponseAllOfMonitorAppsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetCreatedBy(v ListCheckApps200ResponseAllOfMonitorAppsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetChecks

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetCheckGroups() []int64`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) GetCheckGroupsOk() (*[]int64, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) SetCheckGroups(v []int64)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *ListCheckApps200ResponseAllOfMonitorAppsInner) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


