# GetCheckApps200ResponseChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetCheckApps200ResponseChecksInnerAccount**](GetCheckApps200ResponseChecksInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckAgent** | Pointer to **NullableString** |  | [optional] 
**CheckInterval** | Pointer to **NullableInt64** |  | [optional] 
**CheckSpec** | Pointer to **NullableString** |  | [optional] 
**CheckType** | Pointer to [**GetCheckApps200ResponseChecksInnerCheckType**](GetCheckApps200ResponseChecksInnerCheckType.md) |  | [optional] 
**Config** | Pointer to [**GetCheckApps200ResponseChecksInnerConfig**](GetCheckApps200ResponseChecksInnerConfig.md) |  | [optional] 
**Container** | Pointer to [**GetCheckApps200ResponseChecksInnerContainer**](GetCheckApps200ResponseChecksInnerContainer.md) |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**GetCheckApps200ResponseChecksInnerCreatedBy**](GetCheckApps200ResponseChecksInnerCreatedBy.md) |  | [optional] 
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

### NewGetCheckApps200ResponseChecksInner

`func NewGetCheckApps200ResponseChecksInner() *GetCheckApps200ResponseChecksInner`

NewGetCheckApps200ResponseChecksInner instantiates a new GetCheckApps200ResponseChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCheckApps200ResponseChecksInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCheckApps200ResponseChecksInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCheckApps200ResponseChecksInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCheckApps200ResponseChecksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetCheckApps200ResponseChecksInner) GetAccount() GetCheckApps200ResponseChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCheckApps200ResponseChecksInner) GetAccountOk() (*GetCheckApps200ResponseChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCheckApps200ResponseChecksInner) SetAccount(v GetCheckApps200ResponseChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCheckApps200ResponseChecksInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *GetCheckApps200ResponseChecksInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetCheckApps200ResponseChecksInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetCheckApps200ResponseChecksInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetCheckApps200ResponseChecksInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiKey

`func (o *GetCheckApps200ResponseChecksInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetCheckApps200ResponseChecksInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetCheckApps200ResponseChecksInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetCheckApps200ResponseChecksInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAvailability

`func (o *GetCheckApps200ResponseChecksInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetCheckApps200ResponseChecksInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetCheckApps200ResponseChecksInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetCheckApps200ResponseChecksInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckAgent

`func (o *GetCheckApps200ResponseChecksInner) GetCheckAgent() string`

GetCheckAgent returns the CheckAgent field if non-nil, zero value otherwise.

### GetCheckAgentOk

`func (o *GetCheckApps200ResponseChecksInner) GetCheckAgentOk() (*string, bool)`

GetCheckAgentOk returns a tuple with the CheckAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgent

`func (o *GetCheckApps200ResponseChecksInner) SetCheckAgent(v string)`

SetCheckAgent sets CheckAgent field to given value.

### HasCheckAgent

`func (o *GetCheckApps200ResponseChecksInner) HasCheckAgent() bool`

HasCheckAgent returns a boolean if a field has been set.

### SetCheckAgentNil

`func (o *GetCheckApps200ResponseChecksInner) SetCheckAgentNil(b bool)`

 SetCheckAgentNil sets the value for CheckAgent to be an explicit nil

### UnsetCheckAgent
`func (o *GetCheckApps200ResponseChecksInner) UnsetCheckAgent()`

UnsetCheckAgent ensures that no value is present for CheckAgent, not even an explicit nil
### GetCheckInterval

`func (o *GetCheckApps200ResponseChecksInner) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *GetCheckApps200ResponseChecksInner) GetCheckIntervalOk() (*int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *GetCheckApps200ResponseChecksInner) SetCheckInterval(v int64)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *GetCheckApps200ResponseChecksInner) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### SetCheckIntervalNil

`func (o *GetCheckApps200ResponseChecksInner) SetCheckIntervalNil(b bool)`

 SetCheckIntervalNil sets the value for CheckInterval to be an explicit nil

### UnsetCheckInterval
`func (o *GetCheckApps200ResponseChecksInner) UnsetCheckInterval()`

UnsetCheckInterval ensures that no value is present for CheckInterval, not even an explicit nil
### GetCheckSpec

`func (o *GetCheckApps200ResponseChecksInner) GetCheckSpec() string`

GetCheckSpec returns the CheckSpec field if non-nil, zero value otherwise.

### GetCheckSpecOk

`func (o *GetCheckApps200ResponseChecksInner) GetCheckSpecOk() (*string, bool)`

GetCheckSpecOk returns a tuple with the CheckSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSpec

`func (o *GetCheckApps200ResponseChecksInner) SetCheckSpec(v string)`

SetCheckSpec sets CheckSpec field to given value.

### HasCheckSpec

`func (o *GetCheckApps200ResponseChecksInner) HasCheckSpec() bool`

HasCheckSpec returns a boolean if a field has been set.

### SetCheckSpecNil

`func (o *GetCheckApps200ResponseChecksInner) SetCheckSpecNil(b bool)`

 SetCheckSpecNil sets the value for CheckSpec to be an explicit nil

### UnsetCheckSpec
`func (o *GetCheckApps200ResponseChecksInner) UnsetCheckSpec()`

UnsetCheckSpec ensures that no value is present for CheckSpec, not even an explicit nil
### GetCheckType

`func (o *GetCheckApps200ResponseChecksInner) GetCheckType() GetCheckApps200ResponseChecksInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *GetCheckApps200ResponseChecksInner) GetCheckTypeOk() (*GetCheckApps200ResponseChecksInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *GetCheckApps200ResponseChecksInner) SetCheckType(v GetCheckApps200ResponseChecksInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *GetCheckApps200ResponseChecksInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *GetCheckApps200ResponseChecksInner) GetConfig() GetCheckApps200ResponseChecksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetCheckApps200ResponseChecksInner) GetConfigOk() (*GetCheckApps200ResponseChecksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetCheckApps200ResponseChecksInner) SetConfig(v GetCheckApps200ResponseChecksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetCheckApps200ResponseChecksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *GetCheckApps200ResponseChecksInner) GetContainer() GetCheckApps200ResponseChecksInnerContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GetCheckApps200ResponseChecksInner) GetContainerOk() (*GetCheckApps200ResponseChecksInnerContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GetCheckApps200ResponseChecksInner) SetContainer(v GetCheckApps200ResponseChecksInnerContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GetCheckApps200ResponseChecksInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetCheckApps200ResponseChecksInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetCheckApps200ResponseChecksInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetCheckApps200ResponseChecksInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetCheckApps200ResponseChecksInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetCheckApps200ResponseChecksInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetCheckApps200ResponseChecksInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetCheckApps200ResponseChecksInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetCheckApps200ResponseChecksInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetCheckApps200ResponseChecksInner) GetCreatedBy() GetCheckApps200ResponseChecksInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCheckApps200ResponseChecksInner) GetCreatedByOk() (*GetCheckApps200ResponseChecksInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCheckApps200ResponseChecksInner) SetCreatedBy(v GetCheckApps200ResponseChecksInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCheckApps200ResponseChecksInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCheckApps200ResponseChecksInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCheckApps200ResponseChecksInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCheckApps200ResponseChecksInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCheckApps200ResponseChecksInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *GetCheckApps200ResponseChecksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCheckApps200ResponseChecksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCheckApps200ResponseChecksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCheckApps200ResponseChecksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCheckApps200ResponseChecksInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCheckApps200ResponseChecksInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDate

`func (o *GetCheckApps200ResponseChecksInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetCheckApps200ResponseChecksInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetCheckApps200ResponseChecksInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetHealth

`func (o *GetCheckApps200ResponseChecksInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetCheckApps200ResponseChecksInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetCheckApps200ResponseChecksInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetCheckApps200ResponseChecksInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *GetCheckApps200ResponseChecksInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetCheckApps200ResponseChecksInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetCheckApps200ResponseChecksInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetCheckApps200ResponseChecksInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastBoxStats

`func (o *GetCheckApps200ResponseChecksInner) GetLastBoxStats() string`

GetLastBoxStats returns the LastBoxStats field if non-nil, zero value otherwise.

### GetLastBoxStatsOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastBoxStatsOk() (*string, bool)`

GetLastBoxStatsOk returns a tuple with the LastBoxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBoxStats

`func (o *GetCheckApps200ResponseChecksInner) SetLastBoxStats(v string)`

SetLastBoxStats sets LastBoxStats field to given value.

### HasLastBoxStats

`func (o *GetCheckApps200ResponseChecksInner) HasLastBoxStats() bool`

HasLastBoxStats returns a boolean if a field has been set.

### SetLastBoxStatsNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastBoxStatsNil(b bool)`

 SetLastBoxStatsNil sets the value for LastBoxStats to be an explicit nil

### UnsetLastBoxStats
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastBoxStats()`

UnsetLastBoxStats ensures that no value is present for LastBoxStats, not even an explicit nil
### GetLastCheckStatus

`func (o *GetCheckApps200ResponseChecksInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetCheckApps200ResponseChecksInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetCheckApps200ResponseChecksInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastError

`func (o *GetCheckApps200ResponseChecksInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetCheckApps200ResponseChecksInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetCheckApps200ResponseChecksInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastErrorDate

`func (o *GetCheckApps200ResponseChecksInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetCheckApps200ResponseChecksInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetCheckApps200ResponseChecksInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastMessage

`func (o *GetCheckApps200ResponseChecksInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *GetCheckApps200ResponseChecksInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *GetCheckApps200ResponseChecksInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetLastMetric

`func (o *GetCheckApps200ResponseChecksInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *GetCheckApps200ResponseChecksInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *GetCheckApps200ResponseChecksInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetLastRunDate

`func (o *GetCheckApps200ResponseChecksInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetCheckApps200ResponseChecksInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetCheckApps200ResponseChecksInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastStats

`func (o *GetCheckApps200ResponseChecksInner) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *GetCheckApps200ResponseChecksInner) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *GetCheckApps200ResponseChecksInner) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### SetLastStatsNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastStatsNil(b bool)`

 SetLastStatsNil sets the value for LastStats to be an explicit nil

### UnsetLastStats
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastStats()`

UnsetLastStats ensures that no value is present for LastStats, not even an explicit nil
### GetLastSuccessDate

`func (o *GetCheckApps200ResponseChecksInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetCheckApps200ResponseChecksInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetCheckApps200ResponseChecksInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastTimer

`func (o *GetCheckApps200ResponseChecksInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetCheckApps200ResponseChecksInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetCheckApps200ResponseChecksInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### SetLastTimerNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastTimerNil(b bool)`

 SetLastTimerNil sets the value for LastTimer to be an explicit nil

### UnsetLastTimer
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastTimer()`

UnsetLastTimer ensures that no value is present for LastTimer, not even an explicit nil
### GetLastUpdated

`func (o *GetCheckApps200ResponseChecksInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCheckApps200ResponseChecksInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCheckApps200ResponseChecksInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastWarningDate

`func (o *GetCheckApps200ResponseChecksInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetCheckApps200ResponseChecksInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetCheckApps200ResponseChecksInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetName

`func (o *GetCheckApps200ResponseChecksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCheckApps200ResponseChecksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCheckApps200ResponseChecksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCheckApps200ResponseChecksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunDate

`func (o *GetCheckApps200ResponseChecksInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *GetCheckApps200ResponseChecksInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *GetCheckApps200ResponseChecksInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetOutageTime

`func (o *GetCheckApps200ResponseChecksInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *GetCheckApps200ResponseChecksInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *GetCheckApps200ResponseChecksInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *GetCheckApps200ResponseChecksInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetSeverity

`func (o *GetCheckApps200ResponseChecksInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetCheckApps200ResponseChecksInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetCheckApps200ResponseChecksInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetCheckApps200ResponseChecksInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *GetCheckApps200ResponseChecksInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetCheckApps200ResponseChecksInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetCheckApps200ResponseChecksInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetCheckApps200ResponseChecksInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GetCheckApps200ResponseChecksInner) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GetCheckApps200ResponseChecksInner) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDeleted

`func (o *GetCheckApps200ResponseChecksInner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetCheckApps200ResponseChecksInner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetCheckApps200ResponseChecksInner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *GetCheckApps200ResponseChecksInner) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


