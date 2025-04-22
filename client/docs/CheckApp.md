# CheckApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**App** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastCheckStatus** | Pointer to **string** |  | [optional] 
**LastWarningDate** | Pointer to **time.Time** |  | [optional] 
**LastErrorDate** | Pointer to **time.Time** |  | [optional] 
**LastSuccessDate** | Pointer to **time.Time** |  | [optional] 
**LastRunDate** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastTimer** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**History** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **string** |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 
**CheckGroups** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCheckApp

`func NewCheckApp() *CheckApp`

NewCheckApp instantiates a new CheckApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAppWithDefaults

`func NewCheckAppWithDefaults() *CheckApp`

NewCheckAppWithDefaults instantiates a new CheckApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckApp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckApp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckApp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CheckApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *CheckApp) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CheckApp) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CheckApp) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CheckApp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *CheckApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApp

`func (o *CheckApp) GetApp() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CheckApp) GetAppOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CheckApp) SetApp(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetApp sets App field to given value.

### HasApp

`func (o *CheckApp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetName

`func (o *CheckApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CheckApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInUptime

`func (o *CheckApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *CheckApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *CheckApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *CheckApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *CheckApp) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *CheckApp) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *CheckApp) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *CheckApp) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### GetLastWarningDate

`func (o *CheckApp) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *CheckApp) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *CheckApp) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *CheckApp) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### GetLastErrorDate

`func (o *CheckApp) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *CheckApp) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *CheckApp) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *CheckApp) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### GetLastSuccessDate

`func (o *CheckApp) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *CheckApp) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *CheckApp) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *CheckApp) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### GetLastRunDate

`func (o *CheckApp) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *CheckApp) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *CheckApp) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *CheckApp) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastError

`func (o *CheckApp) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *CheckApp) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *CheckApp) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *CheckApp) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastTimer

`func (o *CheckApp) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *CheckApp) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *CheckApp) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *CheckApp) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *CheckApp) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CheckApp) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CheckApp) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CheckApp) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *CheckApp) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CheckApp) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CheckApp) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CheckApp) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetSeverity

`func (o *CheckApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CheckApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CheckApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CheckApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *CheckApp) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *CheckApp) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *CheckApp) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *CheckApp) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *CheckApp) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *CheckApp) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *CheckApp) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *CheckApp) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CheckApp) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CheckApp) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CheckApp) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CheckApp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *CheckApp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CheckApp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CheckApp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CheckApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CheckApp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CheckApp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CheckApp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CheckApp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *CheckApp) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CheckApp) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CheckApp) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *CheckApp) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetChecks

`func (o *CheckApp) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *CheckApp) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *CheckApp) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *CheckApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *CheckApp) GetCheckGroups() []int64`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *CheckApp) GetCheckGroupsOk() (*[]int64, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *CheckApp) SetCheckGroups(v []int64)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *CheckApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


