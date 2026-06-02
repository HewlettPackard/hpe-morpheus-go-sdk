# UpdateChecks200ResponseAllOfCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**UpdateChecks200ResponseAllOfCheckAccount**](UpdateChecks200ResponseAllOfCheckAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckAgent** | Pointer to **NullableString** |  | [optional] 
**CheckInterval** | Pointer to **NullableInt64** |  | [optional] 
**CheckSpec** | Pointer to **NullableString** |  | [optional] 
**CheckType** | Pointer to [**UpdateChecks200ResponseAllOfCheckCheckType**](UpdateChecks200ResponseAllOfCheckCheckType.md) |  | [optional] 
**Config** | Pointer to [**UpdateChecks200ResponseAllOfCheckConfig**](UpdateChecks200ResponseAllOfCheckConfig.md) |  | [optional] 
**Container** | Pointer to [**UpdateChecks200ResponseAllOfCheckContainer**](UpdateChecks200ResponseAllOfCheckContainer.md) |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**UpdateChecks200ResponseAllOfCheckCreatedBy**](UpdateChecks200ResponseAllOfCheckCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastBoxStats** | Pointer to **NullableString** |  | [optional] 
**LastCheckStatus** | Pointer to **NullableString** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**LastErrorDate** | Pointer to **NullableTime** |  | [optional] 
**LastMessage** | Pointer to **NullableString** |  | [optional] 
**LastMetric** | Pointer to **NullableString** |  | [optional] 
**LastRunDate** | Pointer to **NullableTime** |  | [optional] 
**LastStats** | Pointer to **NullableString** |  | [optional] 
**LastSuccessDate** | Pointer to **NullableTime** |  | [optional] 
**LastTimer** | Pointer to **NullableInt64** |  | [optional] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] 
**LastWarningDate** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextRunDate** | Pointer to **NullableTime** |  | [optional] 
**OutageTime** | Pointer to **int64** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateChecks200ResponseAllOfCheck

`func NewUpdateChecks200ResponseAllOfCheck() *UpdateChecks200ResponseAllOfCheck`

NewUpdateChecks200ResponseAllOfCheck instantiates a new UpdateChecks200ResponseAllOfCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateChecks200ResponseAllOfCheck) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateChecks200ResponseAllOfCheck) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateChecks200ResponseAllOfCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateChecks200ResponseAllOfCheck) GetAccount() UpdateChecks200ResponseAllOfCheckAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetAccountOk() (*UpdateChecks200ResponseAllOfCheckAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateChecks200ResponseAllOfCheck) SetAccount(v UpdateChecks200ResponseAllOfCheckAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateChecks200ResponseAllOfCheck) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *UpdateChecks200ResponseAllOfCheck) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateChecks200ResponseAllOfCheck) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateChecks200ResponseAllOfCheck) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateChecks200ResponseAllOfCheck) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateChecks200ResponseAllOfCheck) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateChecks200ResponseAllOfCheck) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAvailability

`func (o *UpdateChecks200ResponseAllOfCheck) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *UpdateChecks200ResponseAllOfCheck) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *UpdateChecks200ResponseAllOfCheck) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckAgent

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckAgent() string`

GetCheckAgent returns the CheckAgent field if non-nil, zero value otherwise.

### GetCheckAgentOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckAgentOk() (*string, bool)`

GetCheckAgentOk returns a tuple with the CheckAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgent

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckAgent(v string)`

SetCheckAgent sets CheckAgent field to given value.

### HasCheckAgent

`func (o *UpdateChecks200ResponseAllOfCheck) HasCheckAgent() bool`

HasCheckAgent returns a boolean if a field has been set.

### SetCheckAgentNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckAgentNil(b bool)`

 SetCheckAgentNil sets the value for CheckAgent to be an explicit nil

### UnsetCheckAgent
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetCheckAgent()`

UnsetCheckAgent ensures that no value is present for CheckAgent, not even an explicit nil
### GetCheckInterval

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckIntervalOk() (*int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckInterval(v int64)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *UpdateChecks200ResponseAllOfCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### SetCheckIntervalNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckIntervalNil(b bool)`

 SetCheckIntervalNil sets the value for CheckInterval to be an explicit nil

### UnsetCheckInterval
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetCheckInterval()`

UnsetCheckInterval ensures that no value is present for CheckInterval, not even an explicit nil
### GetCheckSpec

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckSpec() string`

GetCheckSpec returns the CheckSpec field if non-nil, zero value otherwise.

### GetCheckSpecOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckSpecOk() (*string, bool)`

GetCheckSpecOk returns a tuple with the CheckSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSpec

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckSpec(v string)`

SetCheckSpec sets CheckSpec field to given value.

### HasCheckSpec

`func (o *UpdateChecks200ResponseAllOfCheck) HasCheckSpec() bool`

HasCheckSpec returns a boolean if a field has been set.

### SetCheckSpecNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckSpecNil(b bool)`

 SetCheckSpecNil sets the value for CheckSpec to be an explicit nil

### UnsetCheckSpec
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetCheckSpec()`

UnsetCheckSpec ensures that no value is present for CheckSpec, not even an explicit nil
### GetCheckType

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckType() UpdateChecks200ResponseAllOfCheckCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetCheckTypeOk() (*UpdateChecks200ResponseAllOfCheckCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *UpdateChecks200ResponseAllOfCheck) SetCheckType(v UpdateChecks200ResponseAllOfCheckCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *UpdateChecks200ResponseAllOfCheck) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateChecks200ResponseAllOfCheck) GetConfig() UpdateChecks200ResponseAllOfCheckConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetConfigOk() (*UpdateChecks200ResponseAllOfCheckConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateChecks200ResponseAllOfCheck) SetConfig(v UpdateChecks200ResponseAllOfCheckConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateChecks200ResponseAllOfCheck) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *UpdateChecks200ResponseAllOfCheck) GetContainer() UpdateChecks200ResponseAllOfCheckContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetContainerOk() (*UpdateChecks200ResponseAllOfCheckContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *UpdateChecks200ResponseAllOfCheck) SetContainer(v UpdateChecks200ResponseAllOfCheckContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *UpdateChecks200ResponseAllOfCheck) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreateIncident

`func (o *UpdateChecks200ResponseAllOfCheck) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *UpdateChecks200ResponseAllOfCheck) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *UpdateChecks200ResponseAllOfCheck) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *UpdateChecks200ResponseAllOfCheck) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *UpdateChecks200ResponseAllOfCheck) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *UpdateChecks200ResponseAllOfCheck) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateChecks200ResponseAllOfCheck) GetCreatedBy() UpdateChecks200ResponseAllOfCheckCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetCreatedByOk() (*UpdateChecks200ResponseAllOfCheckCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateChecks200ResponseAllOfCheck) SetCreatedBy(v UpdateChecks200ResponseAllOfCheckCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateChecks200ResponseAllOfCheck) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateChecks200ResponseAllOfCheck) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateChecks200ResponseAllOfCheck) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateChecks200ResponseAllOfCheck) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateChecks200ResponseAllOfCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateChecks200ResponseAllOfCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateChecks200ResponseAllOfCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetHealth

`func (o *UpdateChecks200ResponseAllOfCheck) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *UpdateChecks200ResponseAllOfCheck) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *UpdateChecks200ResponseAllOfCheck) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *UpdateChecks200ResponseAllOfCheck) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateChecks200ResponseAllOfCheck) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateChecks200ResponseAllOfCheck) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastBoxStats

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastBoxStats() string`

GetLastBoxStats returns the LastBoxStats field if non-nil, zero value otherwise.

### GetLastBoxStatsOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastBoxStatsOk() (*string, bool)`

GetLastBoxStatsOk returns a tuple with the LastBoxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBoxStats

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastBoxStats(v string)`

SetLastBoxStats sets LastBoxStats field to given value.

### HasLastBoxStats

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastBoxStats() bool`

HasLastBoxStats returns a boolean if a field has been set.

### SetLastBoxStatsNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastBoxStatsNil(b bool)`

 SetLastBoxStatsNil sets the value for LastBoxStats to be an explicit nil

### UnsetLastBoxStats
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastBoxStats()`

UnsetLastBoxStats ensures that no value is present for LastBoxStats, not even an explicit nil
### GetLastCheckStatus

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastError

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastErrorDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastMessage

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetLastMetric

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetLastRunDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastStats

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### SetLastStatsNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastStatsNil(b bool)`

 SetLastStatsNil sets the value for LastStats to be an explicit nil

### UnsetLastStats
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastStats()`

UnsetLastStats ensures that no value is present for LastStats, not even an explicit nil
### GetLastSuccessDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastTimer

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### SetLastTimerNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastTimerNil(b bool)`

 SetLastTimerNil sets the value for LastTimer to be an explicit nil

### UnsetLastTimer
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastTimer()`

UnsetLastTimer ensures that no value is present for LastTimer, not even an explicit nil
### GetLastUpdated

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastWarningDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetName

`func (o *UpdateChecks200ResponseAllOfCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateChecks200ResponseAllOfCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateChecks200ResponseAllOfCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetOutageTime

`func (o *UpdateChecks200ResponseAllOfCheck) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *UpdateChecks200ResponseAllOfCheck) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *UpdateChecks200ResponseAllOfCheck) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetSeverity

`func (o *UpdateChecks200ResponseAllOfCheck) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateChecks200ResponseAllOfCheck) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateChecks200ResponseAllOfCheck) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateChecks200ResponseAllOfCheck) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateChecks200ResponseAllOfCheck) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateChecks200ResponseAllOfCheck) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *UpdateChecks200ResponseAllOfCheck) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *UpdateChecks200ResponseAllOfCheck) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDeleted

`func (o *UpdateChecks200ResponseAllOfCheck) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UpdateChecks200ResponseAllOfCheck) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UpdateChecks200ResponseAllOfCheck) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UpdateChecks200ResponseAllOfCheck) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


