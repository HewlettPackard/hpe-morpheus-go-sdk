# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckAgent** | Pointer to **string** |  | [optional] 
**CheckInterval** | Pointer to **int64** |  | [optional] 
**CheckSpec** | Pointer to **string** |  | [optional] 
**CheckType** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCheckType**](GetAlerts200ResponseAllOfChecksInnerCheckType.md) |  | [optional] 
**Config** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerConfig**](GetAlerts200ResponseAllOfChecksInnerConfig.md) |  | [optional] 
**Container** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerContainer**](GetAlerts200ResponseAllOfChecksInnerContainer.md) |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCreatedBy**](GetAlerts200ResponseAllOfChecksInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastBoxStats** | Pointer to **string** |  | [optional] 
**LastCheckStatus** | Pointer to **string** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastErrorDate** | Pointer to **time.Time** |  | [optional] 
**LastMessage** | Pointer to **string** |  | [optional] 
**LastMetric** | Pointer to **string** |  | [optional] 
**LastRunDate** | Pointer to **time.Time** |  | [optional] 
**LastStats** | Pointer to **string** |  | [optional] 
**LastSuccessDate** | Pointer to **time.Time** |  | [optional] 
**LastTimer** | Pointer to **int64** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastWarningDate** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextRunDate** | Pointer to **time.Time** |  | [optional] 
**OutageTime** | Pointer to **int64** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheck

`func NewCheck() *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Check) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Check) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Check) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Check) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *Check) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Check) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Check) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Check) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *Check) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Check) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Check) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Check) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiKey

`func (o *Check) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Check) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Check) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Check) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAvailability

`func (o *Check) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Check) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Check) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *Check) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckAgent

`func (o *Check) GetCheckAgent() string`

GetCheckAgent returns the CheckAgent field if non-nil, zero value otherwise.

### GetCheckAgentOk

`func (o *Check) GetCheckAgentOk() (*string, bool)`

GetCheckAgentOk returns a tuple with the CheckAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgent

`func (o *Check) SetCheckAgent(v string)`

SetCheckAgent sets CheckAgent field to given value.

### HasCheckAgent

`func (o *Check) HasCheckAgent() bool`

HasCheckAgent returns a boolean if a field has been set.

### GetCheckInterval

`func (o *Check) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *Check) GetCheckIntervalOk() (*int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *Check) SetCheckInterval(v int64)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *Check) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetCheckSpec

`func (o *Check) GetCheckSpec() string`

GetCheckSpec returns the CheckSpec field if non-nil, zero value otherwise.

### GetCheckSpecOk

`func (o *Check) GetCheckSpecOk() (*string, bool)`

GetCheckSpecOk returns a tuple with the CheckSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSpec

`func (o *Check) SetCheckSpec(v string)`

SetCheckSpec sets CheckSpec field to given value.

### HasCheckSpec

`func (o *Check) HasCheckSpec() bool`

HasCheckSpec returns a boolean if a field has been set.

### GetCheckType

`func (o *Check) GetCheckType() GetAlerts200ResponseAllOfChecksInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *Check) GetCheckTypeOk() (*GetAlerts200ResponseAllOfChecksInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *Check) SetCheckType(v GetAlerts200ResponseAllOfChecksInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *Check) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *Check) GetConfig() GetAlerts200ResponseAllOfChecksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Check) GetConfigOk() (*GetAlerts200ResponseAllOfChecksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Check) SetConfig(v GetAlerts200ResponseAllOfChecksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Check) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *Check) GetContainer() GetAlerts200ResponseAllOfChecksInnerContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *Check) GetContainerOk() (*GetAlerts200ResponseAllOfChecksInnerContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *Check) SetContainer(v GetAlerts200ResponseAllOfChecksInnerContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *Check) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreateIncident

`func (o *Check) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *Check) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *Check) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *Check) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *Check) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Check) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Check) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *Check) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Check) GetCreatedBy() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Check) GetCreatedByOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Check) SetCreatedBy(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Check) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *Check) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Check) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Check) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Check) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Check) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Check) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Check) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Check) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *Check) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Check) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Check) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Check) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetHealth

`func (o *Check) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Check) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Check) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *Check) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *Check) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *Check) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *Check) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *Check) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastBoxStats

`func (o *Check) GetLastBoxStats() string`

GetLastBoxStats returns the LastBoxStats field if non-nil, zero value otherwise.

### GetLastBoxStatsOk

`func (o *Check) GetLastBoxStatsOk() (*string, bool)`

GetLastBoxStatsOk returns a tuple with the LastBoxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBoxStats

`func (o *Check) SetLastBoxStats(v string)`

SetLastBoxStats sets LastBoxStats field to given value.

### HasLastBoxStats

`func (o *Check) HasLastBoxStats() bool`

HasLastBoxStats returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *Check) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *Check) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *Check) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *Check) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### GetLastError

`func (o *Check) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *Check) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *Check) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *Check) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastErrorDate

`func (o *Check) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *Check) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *Check) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *Check) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### GetLastMessage

`func (o *Check) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *Check) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *Check) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *Check) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetLastMetric

`func (o *Check) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *Check) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *Check) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *Check) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### GetLastRunDate

`func (o *Check) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *Check) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *Check) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *Check) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastStats

`func (o *Check) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *Check) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *Check) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *Check) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### GetLastSuccessDate

`func (o *Check) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *Check) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *Check) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *Check) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### GetLastTimer

`func (o *Check) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *Check) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *Check) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *Check) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Check) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Check) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Check) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Check) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastWarningDate

`func (o *Check) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *Check) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *Check) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *Check) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### GetName

`func (o *Check) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Check) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Check) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Check) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunDate

`func (o *Check) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *Check) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *Check) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *Check) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### GetOutageTime

`func (o *Check) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *Check) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *Check) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *Check) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetSeverity

`func (o *Check) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Check) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Check) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Check) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *Check) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Check) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Check) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Check) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDeleted

`func (o *Check) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Check) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Check) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Check) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


