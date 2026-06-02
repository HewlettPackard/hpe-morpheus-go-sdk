# AddCheckGroups200ResponseAllOfCheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**AddCheckGroups200ResponseAllOfCheckGroupAccount**](AddCheckGroups200ResponseAllOfCheckGroupAccount.md) |  | [optional] 
**Instance** | Pointer to [**AddCheckGroups200ResponseAllOfCheckGroupInstance**](AddCheckGroups200ResponseAllOfCheckGroupInstance.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**AddCheckGroups200ResponseAllOfCheckGroupCreatedBy**](AddCheckGroups200ResponseAllOfCheckGroupCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableFloat32** |  | [optional] 
**CheckType** | Pointer to [**AddCheckGroups200ResponseAllOfCheckGroupCheckType**](AddCheckGroups200ResponseAllOfCheckGroupCheckType.md) |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewAddCheckGroups200ResponseAllOfCheckGroup

`func NewAddCheckGroups200ResponseAllOfCheckGroup() *AddCheckGroups200ResponseAllOfCheckGroup`

NewAddCheckGroups200ResponseAllOfCheckGroup instantiates a new AddCheckGroups200ResponseAllOfCheckGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetAccount() AddCheckGroups200ResponseAllOfCheckGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetAccountOk() (*AddCheckGroups200ResponseAllOfCheckGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetAccount(v AddCheckGroups200ResponseAllOfCheckGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetInstance() AddCheckGroups200ResponseAllOfCheckGroupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetInstanceOk() (*AddCheckGroups200ResponseAllOfCheckGroupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetInstance(v AddCheckGroups200ResponseAllOfCheckGroupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetOutageTime

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetMinHappy

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetSeverity

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetCreatedBy() AddCheckGroups200ResponseAllOfCheckGroupCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetCreatedByOk() (*AddCheckGroups200ResponseAllOfCheckGroupCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetCreatedBy(v AddCheckGroups200ResponseAllOfCheckGroupCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *AddCheckGroups200ResponseAllOfCheckGroup) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetCheckType

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetCheckType() AddCheckGroups200ResponseAllOfCheckGroupCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetCheckTypeOk() (*AddCheckGroups200ResponseAllOfCheckGroupCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetCheckType(v AddCheckGroups200ResponseAllOfCheckGroupCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *AddCheckGroups200ResponseAllOfCheckGroup) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


