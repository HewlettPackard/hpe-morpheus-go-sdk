# AddCheckApps200ResponseAllOfCheckApp

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

### NewAddCheckApps200ResponseAllOfCheckApp

`func NewAddCheckApps200ResponseAllOfCheckApp() *AddCheckApps200ResponseAllOfCheckApp`

NewAddCheckApps200ResponseAllOfCheckApp instantiates a new AddCheckApps200ResponseAllOfCheckApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetAccount() AddCheckApps200ResponseAllOfCheckAppAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetAccountOk() (*AddCheckApps200ResponseAllOfCheckAppAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetAccount(v AddCheckApps200ResponseAllOfCheckAppAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApp

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetApp() AddCheckApps200ResponseAllOfCheckAppApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetAppOk() (*AddCheckApps200ResponseAllOfCheckAppApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetApp(v AddCheckApps200ResponseAllOfCheckAppApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetName

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastTimer

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetSeverity

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetCreatedBy() AddCheckApps200ResponseAllOfCheckAppCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetCreatedByOk() (*AddCheckApps200ResponseAllOfCheckAppCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetCreatedBy(v AddCheckApps200ResponseAllOfCheckAppCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *AddCheckApps200ResponseAllOfCheckApp) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetChecks

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetCheckGroups() []int64`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *AddCheckApps200ResponseAllOfCheckApp) GetCheckGroupsOk() (*[]int64, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *AddCheckApps200ResponseAllOfCheckApp) SetCheckGroups(v []int64)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *AddCheckApps200ResponseAllOfCheckApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


