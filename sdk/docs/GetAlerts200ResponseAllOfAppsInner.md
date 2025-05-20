# GetAlerts200ResponseAllOfAppsInner

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

### NewGetAlerts200ResponseAllOfAppsInner

`func NewGetAlerts200ResponseAllOfAppsInner() *GetAlerts200ResponseAllOfAppsInner`

NewGetAlerts200ResponseAllOfAppsInner instantiates a new GetAlerts200ResponseAllOfAppsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfAppsInnerWithDefaults

`func NewGetAlerts200ResponseAllOfAppsInnerWithDefaults() *GetAlerts200ResponseAllOfAppsInner`

NewGetAlerts200ResponseAllOfAppsInnerWithDefaults instantiates a new GetAlerts200ResponseAllOfAppsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAlerts200ResponseAllOfAppsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAlerts200ResponseAllOfAppsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAlerts200ResponseAllOfAppsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetAlerts200ResponseAllOfAppsInner) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetAlerts200ResponseAllOfAppsInner) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetAlerts200ResponseAllOfAppsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *GetAlerts200ResponseAllOfAppsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetAlerts200ResponseAllOfAppsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetAlerts200ResponseAllOfAppsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApp

`func (o *GetAlerts200ResponseAllOfAppsInner) GetApp() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetAppOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GetAlerts200ResponseAllOfAppsInner) SetApp(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetApp sets App field to given value.

### HasApp

`func (o *GetAlerts200ResponseAllOfAppsInner) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetName

`func (o *GetAlerts200ResponseAllOfAppsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAlerts200ResponseAllOfAppsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAlerts200ResponseAllOfAppsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetAlerts200ResponseAllOfAppsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAlerts200ResponseAllOfAppsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAlerts200ResponseAllOfAppsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInUptime

`func (o *GetAlerts200ResponseAllOfAppsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetAlerts200ResponseAllOfAppsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetAlerts200ResponseAllOfAppsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### GetLastWarningDate

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### GetLastErrorDate

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### GetLastSuccessDate

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### GetLastRunDate

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastError

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastTimer

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *GetAlerts200ResponseAllOfAppsInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetAlerts200ResponseAllOfAppsInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetAlerts200ResponseAllOfAppsInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *GetAlerts200ResponseAllOfAppsInner) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetAlerts200ResponseAllOfAppsInner) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GetAlerts200ResponseAllOfAppsInner) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetSeverity

`func (o *GetAlerts200ResponseAllOfAppsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetAlerts200ResponseAllOfAppsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetAlerts200ResponseAllOfAppsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetAlerts200ResponseAllOfAppsInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetAlerts200ResponseAllOfAppsInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetAlerts200ResponseAllOfAppsInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetAlerts200ResponseAllOfAppsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetAlerts200ResponseAllOfAppsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetAlerts200ResponseAllOfAppsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetAlerts200ResponseAllOfAppsInner) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetAlerts200ResponseAllOfAppsInner) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetAlerts200ResponseAllOfAppsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetAlerts200ResponseAllOfAppsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetAlerts200ResponseAllOfAppsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetAlerts200ResponseAllOfAppsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetAlerts200ResponseAllOfAppsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetAlerts200ResponseAllOfAppsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *GetAlerts200ResponseAllOfAppsInner) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetAlerts200ResponseAllOfAppsInner) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetAlerts200ResponseAllOfAppsInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetChecks

`func (o *GetAlerts200ResponseAllOfAppsInner) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetAlerts200ResponseAllOfAppsInner) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetAlerts200ResponseAllOfAppsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetAlerts200ResponseAllOfAppsInner) GetCheckGroups() []int64`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetAlerts200ResponseAllOfAppsInner) GetCheckGroupsOk() (*[]int64, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetAlerts200ResponseAllOfAppsInner) SetCheckGroups(v []int64)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetAlerts200ResponseAllOfAppsInner) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


