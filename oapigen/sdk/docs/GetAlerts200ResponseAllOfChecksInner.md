# GetAlerts200ResponseAllOfChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckAgent** | Pointer to **NullableString** |  | [optional] 
**CheckInterval** | Pointer to **NullableInt64** |  | [optional] 
**CheckSpec** | Pointer to **NullableString** |  | [optional] 
**CheckType** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCheckType**](GetAlerts200ResponseAllOfChecksInnerCheckType.md) |  | [optional] 
**Config** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerConfig**](GetAlerts200ResponseAllOfChecksInnerConfig.md) |  | [optional] 
**Container** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
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

### NewGetAlerts200ResponseAllOfChecksInner

`func NewGetAlerts200ResponseAllOfChecksInner() *GetAlerts200ResponseAllOfChecksInner`

NewGetAlerts200ResponseAllOfChecksInner instantiates a new GetAlerts200ResponseAllOfChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseAllOfChecksInnerWithDefaults

`func NewGetAlerts200ResponseAllOfChecksInnerWithDefaults() *GetAlerts200ResponseAllOfChecksInner`

NewGetAlerts200ResponseAllOfChecksInnerWithDefaults instantiates a new GetAlerts200ResponseAllOfChecksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAlerts200ResponseAllOfChecksInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAlerts200ResponseAllOfChecksInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAlerts200ResponseAllOfChecksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetAlerts200ResponseAllOfChecksInner) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetAlerts200ResponseAllOfChecksInner) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetAlerts200ResponseAllOfChecksInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *GetAlerts200ResponseAllOfChecksInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetAlerts200ResponseAllOfChecksInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetAlerts200ResponseAllOfChecksInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiKey

`func (o *GetAlerts200ResponseAllOfChecksInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetAlerts200ResponseAllOfChecksInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetAlerts200ResponseAllOfChecksInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAvailability

`func (o *GetAlerts200ResponseAllOfChecksInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetAlerts200ResponseAllOfChecksInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetAlerts200ResponseAllOfChecksInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckAgent

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckAgent() string`

GetCheckAgent returns the CheckAgent field if non-nil, zero value otherwise.

### GetCheckAgentOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckAgentOk() (*string, bool)`

GetCheckAgentOk returns a tuple with the CheckAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgent

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckAgent(v string)`

SetCheckAgent sets CheckAgent field to given value.

### HasCheckAgent

`func (o *GetAlerts200ResponseAllOfChecksInner) HasCheckAgent() bool`

HasCheckAgent returns a boolean if a field has been set.

### SetCheckAgentNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckAgentNil(b bool)`

 SetCheckAgentNil sets the value for CheckAgent to be an explicit nil

### UnsetCheckAgent
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetCheckAgent()`

UnsetCheckAgent ensures that no value is present for CheckAgent, not even an explicit nil
### GetCheckInterval

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckIntervalOk() (*int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckInterval(v int64)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *GetAlerts200ResponseAllOfChecksInner) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### SetCheckIntervalNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckIntervalNil(b bool)`

 SetCheckIntervalNil sets the value for CheckInterval to be an explicit nil

### UnsetCheckInterval
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetCheckInterval()`

UnsetCheckInterval ensures that no value is present for CheckInterval, not even an explicit nil
### GetCheckSpec

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckSpec() string`

GetCheckSpec returns the CheckSpec field if non-nil, zero value otherwise.

### GetCheckSpecOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckSpecOk() (*string, bool)`

GetCheckSpecOk returns a tuple with the CheckSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSpec

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckSpec(v string)`

SetCheckSpec sets CheckSpec field to given value.

### HasCheckSpec

`func (o *GetAlerts200ResponseAllOfChecksInner) HasCheckSpec() bool`

HasCheckSpec returns a boolean if a field has been set.

### SetCheckSpecNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckSpecNil(b bool)`

 SetCheckSpecNil sets the value for CheckSpec to be an explicit nil

### UnsetCheckSpec
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetCheckSpec()`

UnsetCheckSpec ensures that no value is present for CheckSpec, not even an explicit nil
### GetCheckType

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckType() GetAlerts200ResponseAllOfChecksInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCheckTypeOk() (*GetAlerts200ResponseAllOfChecksInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCheckType(v GetAlerts200ResponseAllOfChecksInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *GetAlerts200ResponseAllOfChecksInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *GetAlerts200ResponseAllOfChecksInner) GetConfig() GetAlerts200ResponseAllOfChecksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetConfigOk() (*GetAlerts200ResponseAllOfChecksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetAlerts200ResponseAllOfChecksInner) SetConfig(v GetAlerts200ResponseAllOfChecksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetAlerts200ResponseAllOfChecksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *GetAlerts200ResponseAllOfChecksInner) GetContainer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetContainerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GetAlerts200ResponseAllOfChecksInner) SetContainer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GetAlerts200ResponseAllOfChecksInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetAlerts200ResponseAllOfChecksInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *GetAlerts200ResponseAllOfChecksInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetAlerts200ResponseAllOfChecksInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetAlerts200ResponseAllOfChecksInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetAlerts200ResponseAllOfChecksInner) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetAlerts200ResponseAllOfChecksInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetAlerts200ResponseAllOfChecksInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetAlerts200ResponseAllOfChecksInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetAlerts200ResponseAllOfChecksInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *GetAlerts200ResponseAllOfChecksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAlerts200ResponseAllOfChecksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAlerts200ResponseAllOfChecksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetHealth

`func (o *GetAlerts200ResponseAllOfChecksInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetAlerts200ResponseAllOfChecksInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetAlerts200ResponseAllOfChecksInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *GetAlerts200ResponseAllOfChecksInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetAlerts200ResponseAllOfChecksInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetAlerts200ResponseAllOfChecksInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastBoxStats

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastBoxStats() string`

GetLastBoxStats returns the LastBoxStats field if non-nil, zero value otherwise.

### GetLastBoxStatsOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastBoxStatsOk() (*string, bool)`

GetLastBoxStatsOk returns a tuple with the LastBoxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBoxStats

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastBoxStats(v string)`

SetLastBoxStats sets LastBoxStats field to given value.

### HasLastBoxStats

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastBoxStats() bool`

HasLastBoxStats returns a boolean if a field has been set.

### SetLastBoxStatsNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastBoxStatsNil(b bool)`

 SetLastBoxStatsNil sets the value for LastBoxStats to be an explicit nil

### UnsetLastBoxStats
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastBoxStats()`

UnsetLastBoxStats ensures that no value is present for LastBoxStats, not even an explicit nil
### GetLastCheckStatus

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastError

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastErrorDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastMessage

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetLastMetric

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetLastRunDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastStats

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### SetLastStatsNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastStatsNil(b bool)`

 SetLastStatsNil sets the value for LastStats to be an explicit nil

### UnsetLastStats
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastStats()`

UnsetLastStats ensures that no value is present for LastStats, not even an explicit nil
### GetLastSuccessDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastTimer

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### SetLastTimerNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastTimerNil(b bool)`

 SetLastTimerNil sets the value for LastTimer to be an explicit nil

### UnsetLastTimer
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastTimer()`

UnsetLastTimer ensures that no value is present for LastTimer, not even an explicit nil
### GetLastUpdated

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastWarningDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetName

`func (o *GetAlerts200ResponseAllOfChecksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAlerts200ResponseAllOfChecksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAlerts200ResponseAllOfChecksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetOutageTime

`func (o *GetAlerts200ResponseAllOfChecksInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *GetAlerts200ResponseAllOfChecksInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *GetAlerts200ResponseAllOfChecksInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetSeverity

`func (o *GetAlerts200ResponseAllOfChecksInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetAlerts200ResponseAllOfChecksInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetAlerts200ResponseAllOfChecksInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *GetAlerts200ResponseAllOfChecksInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetAlerts200ResponseAllOfChecksInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetAlerts200ResponseAllOfChecksInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GetAlerts200ResponseAllOfChecksInner) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GetAlerts200ResponseAllOfChecksInner) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDeleted

`func (o *GetAlerts200ResponseAllOfChecksInner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetAlerts200ResponseAllOfChecksInner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetAlerts200ResponseAllOfChecksInner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *GetAlerts200ResponseAllOfChecksInner) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


