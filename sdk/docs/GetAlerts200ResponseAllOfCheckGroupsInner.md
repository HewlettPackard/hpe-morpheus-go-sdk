# GetAlerts200ResponseAllOfCheckGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Instance** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastCheckStatus** | Pointer to **string** |  | [optional] 
**LastWarningDate** | Pointer to **time.Time** |  | [optional] 
**LastErrorDate** | Pointer to **time.Time** |  | [optional] 
**LastSuccessDate** | Pointer to **time.Time** |  | [optional] 
**LastRunDate** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**OutageTime** | Pointer to **int64** |  | [optional] 
**LastTimer** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**History** | Pointer to **string** |  | [optional] 
**MinHappy** | Pointer to **int64** |  | [optional] 
**LastMetric** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckType** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCheckType**](GetAlerts200ResponseAllOfChecksInnerCheckType.md) |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewGetAlerts200ResponseAllOfCheckGroupsInner

`func NewGetAlerts200ResponseAllOfCheckGroupsInner() *GetAlerts200ResponseAllOfCheckGroupsInner`

NewGetAlerts200ResponseAllOfCheckGroupsInner instantiates a new GetAlerts200ResponseAllOfCheckGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfCheckGroupsInnerWithDefaults

`func NewGetAlerts200ResponseAllOfCheckGroupsInnerWithDefaults() *GetAlerts200ResponseAllOfCheckGroupsInner`

NewGetAlerts200ResponseAllOfCheckGroupsInnerWithDefaults instantiates a new GetAlerts200ResponseAllOfCheckGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetInstance() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetInstanceOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetInstance(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInUptime

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### GetLastWarningDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### GetLastErrorDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### GetLastSuccessDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### GetLastRunDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastError

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetOutageTime

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMinHappy

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### GetSeverity

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckType

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetCheckType() GetAlerts200ResponseAllOfChecksInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetCheckTypeOk() (*GetAlerts200ResponseAllOfChecksInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetCheckType(v GetAlerts200ResponseAllOfChecksInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetAlerts200ResponseAllOfCheckGroupsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


