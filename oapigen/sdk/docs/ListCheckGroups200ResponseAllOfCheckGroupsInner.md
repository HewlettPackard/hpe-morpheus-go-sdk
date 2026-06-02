# ListCheckGroups200ResponseAllOfCheckGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListCheckGroups200ResponseAllOfCheckGroupsInnerAccount**](ListCheckGroups200ResponseAllOfCheckGroupsInnerAccount.md) |  | [optional] 
**Instance** | Pointer to [**ListCheckGroups200ResponseAllOfCheckGroupsInnerInstance**](ListCheckGroups200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastCheckStatus** | Pointer to **NullableString** |  | [optional] 
**LastWarningDate** | Pointer to **NullableTime** |  | [optional] 
**LastErrorDate** | Pointer to **NullableTime** |  | [optional] 
**LastSuccessDate** | Pointer to **NullableTime** |  | [optional] 
**LastRunDate** | Pointer to **NullableTime** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**OutageTime** | Pointer to **int64** |  | [optional] 
**LastTimer** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**History** | Pointer to **NullableString** |  | [optional] 
**MinHappy** | Pointer to **int64** |  | [optional] 
**LastMetric** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**ListCheckGroups200ResponseAllOfCheckGroupsInnerCreatedBy**](ListCheckGroups200ResponseAllOfCheckGroupsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableFloat32** |  | [optional] 
**CheckType** | Pointer to [**ListCheckGroups200ResponseAllOfCheckGroupsInnerCheckType**](ListCheckGroups200ResponseAllOfCheckGroupsInnerCheckType.md) |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewListCheckGroups200ResponseAllOfCheckGroupsInner

`func NewListCheckGroups200ResponseAllOfCheckGroupsInner() *ListCheckGroups200ResponseAllOfCheckGroupsInner`

NewListCheckGroups200ResponseAllOfCheckGroupsInner instantiates a new ListCheckGroups200ResponseAllOfCheckGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetAccount() ListCheckGroups200ResponseAllOfCheckGroupsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetAccountOk() (*ListCheckGroups200ResponseAllOfCheckGroupsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetAccount(v ListCheckGroups200ResponseAllOfCheckGroupsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetInstance() ListCheckGroups200ResponseAllOfCheckGroupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetInstanceOk() (*ListCheckGroups200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetInstance(v ListCheckGroups200ResponseAllOfCheckGroupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetOutageTime

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetMinHappy

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetSeverity

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetCreatedBy() ListCheckGroups200ResponseAllOfCheckGroupsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetCreatedByOk() (*ListCheckGroups200ResponseAllOfCheckGroupsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetCreatedBy(v ListCheckGroups200ResponseAllOfCheckGroupsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetCheckType

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetCheckType() ListCheckGroups200ResponseAllOfCheckGroupsInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetCheckTypeOk() (*ListCheckGroups200ResponseAllOfCheckGroupsInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetCheckType(v ListCheckGroups200ResponseAllOfCheckGroupsInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ListCheckGroups200ResponseAllOfCheckGroupsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


