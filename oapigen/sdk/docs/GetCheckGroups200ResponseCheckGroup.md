# GetCheckGroups200ResponseCheckGroup

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

### NewGetCheckGroups200ResponseCheckGroup

`func NewGetCheckGroups200ResponseCheckGroup() *GetCheckGroups200ResponseCheckGroup`

NewGetCheckGroups200ResponseCheckGroup instantiates a new GetCheckGroups200ResponseCheckGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCheckGroups200ResponseCheckGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCheckGroups200ResponseCheckGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCheckGroups200ResponseCheckGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetCheckGroups200ResponseCheckGroup) GetAccount() AddCheckGroups200ResponseAllOfCheckGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetAccountOk() (*AddCheckGroups200ResponseAllOfCheckGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCheckGroups200ResponseCheckGroup) SetAccount(v AddCheckGroups200ResponseAllOfCheckGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCheckGroups200ResponseCheckGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *GetCheckGroups200ResponseCheckGroup) GetInstance() AddCheckGroups200ResponseAllOfCheckGroupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetInstanceOk() (*AddCheckGroups200ResponseAllOfCheckGroupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetCheckGroups200ResponseCheckGroup) SetInstance(v AddCheckGroups200ResponseAllOfCheckGroupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetCheckGroups200ResponseCheckGroup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *GetCheckGroups200ResponseCheckGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCheckGroups200ResponseCheckGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCheckGroups200ResponseCheckGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCheckGroups200ResponseCheckGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCheckGroups200ResponseCheckGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCheckGroups200ResponseCheckGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *GetCheckGroups200ResponseCheckGroup) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetCheckGroups200ResponseCheckGroup) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetCheckGroups200ResponseCheckGroup) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetOutageTime

`func (o *GetCheckGroups200ResponseCheckGroup) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *GetCheckGroups200ResponseCheckGroup) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *GetCheckGroups200ResponseCheckGroup) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *GetCheckGroups200ResponseCheckGroup) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetCheckGroups200ResponseCheckGroup) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetCheckGroups200ResponseCheckGroup) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *GetCheckGroups200ResponseCheckGroup) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetCheckGroups200ResponseCheckGroup) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GetCheckGroups200ResponseCheckGroup) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetMinHappy

`func (o *GetCheckGroups200ResponseCheckGroup) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *GetCheckGroups200ResponseCheckGroup) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *GetCheckGroups200ResponseCheckGroup) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetSeverity

`func (o *GetCheckGroups200ResponseCheckGroup) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetCheckGroups200ResponseCheckGroup) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetCheckGroups200ResponseCheckGroup) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetCheckGroups200ResponseCheckGroup) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetCheckGroups200ResponseCheckGroup) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetCheckGroups200ResponseCheckGroup) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetCheckGroups200ResponseCheckGroup) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetCheckGroups200ResponseCheckGroup) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetCheckGroups200ResponseCheckGroup) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetCheckGroups200ResponseCheckGroup) GetCreatedBy() AddCheckGroups200ResponseAllOfCheckGroupCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetCreatedByOk() (*AddCheckGroups200ResponseAllOfCheckGroupCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCheckGroups200ResponseCheckGroup) SetCreatedBy(v AddCheckGroups200ResponseAllOfCheckGroupCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCheckGroups200ResponseCheckGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCheckGroups200ResponseCheckGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCheckGroups200ResponseCheckGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCheckGroups200ResponseCheckGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCheckGroups200ResponseCheckGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCheckGroups200ResponseCheckGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *GetCheckGroups200ResponseCheckGroup) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetCheckGroups200ResponseCheckGroup) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetCheckGroups200ResponseCheckGroup) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *GetCheckGroups200ResponseCheckGroup) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *GetCheckGroups200ResponseCheckGroup) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetCheckType

`func (o *GetCheckGroups200ResponseCheckGroup) GetCheckType() AddCheckGroups200ResponseAllOfCheckGroupCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetCheckTypeOk() (*AddCheckGroups200ResponseAllOfCheckGroupCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *GetCheckGroups200ResponseCheckGroup) SetCheckType(v AddCheckGroups200ResponseAllOfCheckGroupCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *GetCheckGroups200ResponseCheckGroup) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *GetCheckGroups200ResponseCheckGroup) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckGroups200ResponseCheckGroup) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckGroups200ResponseCheckGroup) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetCheckGroups200ResponseCheckGroup) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


