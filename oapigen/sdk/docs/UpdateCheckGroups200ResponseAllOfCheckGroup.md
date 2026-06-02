# UpdateCheckGroups200ResponseAllOfCheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**UpdateCheckGroups200ResponseAllOfCheckGroupAccount**](UpdateCheckGroups200ResponseAllOfCheckGroupAccount.md) |  | [optional] 
**Instance** | Pointer to [**UpdateCheckGroups200ResponseAllOfCheckGroupInstance**](UpdateCheckGroups200ResponseAllOfCheckGroupInstance.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**UpdateCheckGroups200ResponseAllOfCheckGroupCreatedBy**](UpdateCheckGroups200ResponseAllOfCheckGroupCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableFloat32** |  | [optional] 
**CheckType** | Pointer to [**UpdateCheckGroups200ResponseAllOfCheckGroupCheckType**](UpdateCheckGroups200ResponseAllOfCheckGroupCheckType.md) |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewUpdateCheckGroups200ResponseAllOfCheckGroup

`func NewUpdateCheckGroups200ResponseAllOfCheckGroup() *UpdateCheckGroups200ResponseAllOfCheckGroup`

NewUpdateCheckGroups200ResponseAllOfCheckGroup instantiates a new UpdateCheckGroups200ResponseAllOfCheckGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetAccount() UpdateCheckGroups200ResponseAllOfCheckGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetAccountOk() (*UpdateCheckGroups200ResponseAllOfCheckGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetAccount(v UpdateCheckGroups200ResponseAllOfCheckGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetInstance() UpdateCheckGroups200ResponseAllOfCheckGroupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetInstanceOk() (*UpdateCheckGroups200ResponseAllOfCheckGroupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetInstance(v UpdateCheckGroups200ResponseAllOfCheckGroupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetOutageTime

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetMinHappy

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetSeverity

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetCreatedBy() UpdateCheckGroups200ResponseAllOfCheckGroupCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetCreatedByOk() (*UpdateCheckGroups200ResponseAllOfCheckGroupCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetCreatedBy(v UpdateCheckGroups200ResponseAllOfCheckGroupCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetCheckType

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetCheckType() UpdateCheckGroups200ResponseAllOfCheckGroupCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetCheckTypeOk() (*UpdateCheckGroups200ResponseAllOfCheckGroupCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetCheckType(v UpdateCheckGroups200ResponseAllOfCheckGroupCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *UpdateCheckGroups200ResponseAllOfCheckGroup) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


