# UpdateIncidents200ResponseAllOfIncidentChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**UpdateIncidents200ResponseAllOfIncidentChecksInnerAccount**](UpdateIncidents200ResponseAllOfIncidentChecksInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckAgent** | Pointer to **NullableString** |  | [optional] 
**CheckInterval** | Pointer to **NullableInt64** |  | [optional] 
**CheckSpec** | Pointer to **NullableString** |  | [optional] 
**CheckType** | Pointer to [**UpdateIncidents200ResponseAllOfIncidentChecksInnerCheckType**](UpdateIncidents200ResponseAllOfIncidentChecksInnerCheckType.md) |  | [optional] 
**Config** | Pointer to [**UpdateIncidents200ResponseAllOfIncidentChecksInnerConfig**](UpdateIncidents200ResponseAllOfIncidentChecksInnerConfig.md) |  | [optional] 
**Container** | Pointer to [**UpdateIncidents200ResponseAllOfIncidentChecksInnerContainer**](UpdateIncidents200ResponseAllOfIncidentChecksInnerContainer.md) |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**UpdateIncidents200ResponseAllOfIncidentChecksInnerCreatedBy**](UpdateIncidents200ResponseAllOfIncidentChecksInnerCreatedBy.md) |  | [optional] 
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

### NewUpdateIncidents200ResponseAllOfIncidentChecksInner

`func NewUpdateIncidents200ResponseAllOfIncidentChecksInner() *UpdateIncidents200ResponseAllOfIncidentChecksInner`

NewUpdateIncidents200ResponseAllOfIncidentChecksInner instantiates a new UpdateIncidents200ResponseAllOfIncidentChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetAccount() UpdateIncidents200ResponseAllOfIncidentChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetAccountOk() (*UpdateIncidents200ResponseAllOfIncidentChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetAccount(v UpdateIncidents200ResponseAllOfIncidentChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAvailability

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckAgent

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckAgent() string`

GetCheckAgent returns the CheckAgent field if non-nil, zero value otherwise.

### GetCheckAgentOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckAgentOk() (*string, bool)`

GetCheckAgentOk returns a tuple with the CheckAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgent

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckAgent(v string)`

SetCheckAgent sets CheckAgent field to given value.

### HasCheckAgent

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasCheckAgent() bool`

HasCheckAgent returns a boolean if a field has been set.

### SetCheckAgentNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckAgentNil(b bool)`

 SetCheckAgentNil sets the value for CheckAgent to be an explicit nil

### UnsetCheckAgent
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetCheckAgent()`

UnsetCheckAgent ensures that no value is present for CheckAgent, not even an explicit nil
### GetCheckInterval

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckIntervalOk() (*int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckInterval(v int64)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### SetCheckIntervalNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckIntervalNil(b bool)`

 SetCheckIntervalNil sets the value for CheckInterval to be an explicit nil

### UnsetCheckInterval
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetCheckInterval()`

UnsetCheckInterval ensures that no value is present for CheckInterval, not even an explicit nil
### GetCheckSpec

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckSpec() string`

GetCheckSpec returns the CheckSpec field if non-nil, zero value otherwise.

### GetCheckSpecOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckSpecOk() (*string, bool)`

GetCheckSpecOk returns a tuple with the CheckSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSpec

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckSpec(v string)`

SetCheckSpec sets CheckSpec field to given value.

### HasCheckSpec

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasCheckSpec() bool`

HasCheckSpec returns a boolean if a field has been set.

### SetCheckSpecNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckSpecNil(b bool)`

 SetCheckSpecNil sets the value for CheckSpec to be an explicit nil

### UnsetCheckSpec
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetCheckSpec()`

UnsetCheckSpec ensures that no value is present for CheckSpec, not even an explicit nil
### GetCheckType

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckType() UpdateIncidents200ResponseAllOfIncidentChecksInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCheckTypeOk() (*UpdateIncidents200ResponseAllOfIncidentChecksInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCheckType(v UpdateIncidents200ResponseAllOfIncidentChecksInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetConfig() UpdateIncidents200ResponseAllOfIncidentChecksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetConfigOk() (*UpdateIncidents200ResponseAllOfIncidentChecksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetConfig(v UpdateIncidents200ResponseAllOfIncidentChecksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetContainer() UpdateIncidents200ResponseAllOfIncidentChecksInnerContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetContainerOk() (*UpdateIncidents200ResponseAllOfIncidentChecksInnerContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetContainer(v UpdateIncidents200ResponseAllOfIncidentChecksInnerContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreateIncident

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCreatedBy() UpdateIncidents200ResponseAllOfIncidentChecksInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetCreatedByOk() (*UpdateIncidents200ResponseAllOfIncidentChecksInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetCreatedBy(v UpdateIncidents200ResponseAllOfIncidentChecksInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetHealth

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastBoxStats

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastBoxStats() string`

GetLastBoxStats returns the LastBoxStats field if non-nil, zero value otherwise.

### GetLastBoxStatsOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastBoxStatsOk() (*string, bool)`

GetLastBoxStatsOk returns a tuple with the LastBoxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBoxStats

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastBoxStats(v string)`

SetLastBoxStats sets LastBoxStats field to given value.

### HasLastBoxStats

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastBoxStats() bool`

HasLastBoxStats returns a boolean if a field has been set.

### SetLastBoxStatsNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastBoxStatsNil(b bool)`

 SetLastBoxStatsNil sets the value for LastBoxStats to be an explicit nil

### UnsetLastBoxStats
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastBoxStats()`

UnsetLastBoxStats ensures that no value is present for LastBoxStats, not even an explicit nil
### GetLastCheckStatus

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastError

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastErrorDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastMessage

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetLastMetric

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetLastRunDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastStats

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### SetLastStatsNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastStatsNil(b bool)`

 SetLastStatsNil sets the value for LastStats to be an explicit nil

### UnsetLastStats
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastStats()`

UnsetLastStats ensures that no value is present for LastStats, not even an explicit nil
### GetLastSuccessDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastTimer

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### SetLastTimerNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastTimerNil(b bool)`

 SetLastTimerNil sets the value for LastTimer to be an explicit nil

### UnsetLastTimer
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastTimer()`

UnsetLastTimer ensures that no value is present for LastTimer, not even an explicit nil
### GetLastUpdated

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastWarningDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetName

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetOutageTime

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetSeverity

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDeleted

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UpdateIncidents200ResponseAllOfIncidentChecksInner) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


