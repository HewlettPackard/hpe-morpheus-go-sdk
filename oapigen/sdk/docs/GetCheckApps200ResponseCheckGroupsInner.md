# GetCheckApps200ResponseCheckGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetCheckApps200ResponseCheckGroupsInnerAccount**](GetCheckApps200ResponseCheckGroupsInnerAccount.md) |  | [optional] 
**Instance** | Pointer to [**GetCheckApps200ResponseCheckGroupsInnerInstance**](GetCheckApps200ResponseCheckGroupsInnerInstance.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**GetCheckApps200ResponseCheckGroupsInnerCreatedBy**](GetCheckApps200ResponseCheckGroupsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableFloat32** |  | [optional] 
**CheckType** | Pointer to [**GetCheckApps200ResponseCheckGroupsInnerCheckType**](GetCheckApps200ResponseCheckGroupsInnerCheckType.md) |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewGetCheckApps200ResponseCheckGroupsInner

`func NewGetCheckApps200ResponseCheckGroupsInner() *GetCheckApps200ResponseCheckGroupsInner`

NewGetCheckApps200ResponseCheckGroupsInner instantiates a new GetCheckApps200ResponseCheckGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCheckApps200ResponseCheckGroupsInnerWithDefaults

`func NewGetCheckApps200ResponseCheckGroupsInnerWithDefaults() *GetCheckApps200ResponseCheckGroupsInner`

NewGetCheckApps200ResponseCheckGroupsInnerWithDefaults instantiates a new GetCheckApps200ResponseCheckGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetAccount() GetCheckApps200ResponseCheckGroupsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetAccountOk() (*GetCheckApps200ResponseCheckGroupsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetAccount(v GetCheckApps200ResponseCheckGroupsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetInstance() GetCheckApps200ResponseCheckGroupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetInstanceOk() (*GetCheckApps200ResponseCheckGroupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetInstance(v GetCheckApps200ResponseCheckGroupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetOutageTime

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetMinHappy

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetSeverity

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetCreatedBy() GetCheckApps200ResponseCheckGroupsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetCreatedByOk() (*GetCheckApps200ResponseCheckGroupsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetCreatedBy(v GetCheckApps200ResponseCheckGroupsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *GetCheckApps200ResponseCheckGroupsInner) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetCheckType

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetCheckType() GetCheckApps200ResponseCheckGroupsInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetCheckTypeOk() (*GetCheckApps200ResponseCheckGroupsInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetCheckType(v GetCheckApps200ResponseCheckGroupsInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckApps200ResponseCheckGroupsInner) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckApps200ResponseCheckGroupsInner) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetCheckApps200ResponseCheckGroupsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


