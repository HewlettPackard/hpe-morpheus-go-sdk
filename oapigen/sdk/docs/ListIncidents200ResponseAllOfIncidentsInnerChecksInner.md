# ListIncidents200ResponseAllOfIncidentsInnerChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListIncidents200ResponseAllOfIncidentsInnerChecksInnerAccount**](ListIncidents200ResponseAllOfIncidentsInnerChecksInnerAccount.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**CheckAgent** | Pointer to **NullableString** |  | [optional] 
**CheckInterval** | Pointer to **NullableInt64** |  | [optional] 
**CheckSpec** | Pointer to **NullableString** |  | [optional] 
**CheckType** | Pointer to [**ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCheckType**](ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCheckType.md) |  | [optional] 
**Config** | Pointer to [**ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig**](ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig.md) |  | [optional] 
**Container** | Pointer to [**ListIncidents200ResponseAllOfIncidentsInnerChecksInnerContainer**](ListIncidents200ResponseAllOfIncidentsInnerChecksInnerContainer.md) |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCreatedBy**](ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCreatedBy.md) |  | [optional] 
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

### NewListIncidents200ResponseAllOfIncidentsInnerChecksInner

`func NewListIncidents200ResponseAllOfIncidentsInnerChecksInner() *ListIncidents200ResponseAllOfIncidentsInnerChecksInner`

NewListIncidents200ResponseAllOfIncidentsInnerChecksInner instantiates a new ListIncidents200ResponseAllOfIncidentsInnerChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerWithDefaults

`func NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerWithDefaults() *ListIncidents200ResponseAllOfIncidentsInnerChecksInner`

NewListIncidents200ResponseAllOfIncidentsInnerChecksInnerWithDefaults instantiates a new ListIncidents200ResponseAllOfIncidentsInnerChecksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetAccount() ListIncidents200ResponseAllOfIncidentsInnerChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetAccountOk() (*ListIncidents200ResponseAllOfIncidentsInnerChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetAccount(v ListIncidents200ResponseAllOfIncidentsInnerChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiKey

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAvailability

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCheckAgent

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckAgent() string`

GetCheckAgent returns the CheckAgent field if non-nil, zero value otherwise.

### GetCheckAgentOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckAgentOk() (*string, bool)`

GetCheckAgentOk returns a tuple with the CheckAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgent

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckAgent(v string)`

SetCheckAgent sets CheckAgent field to given value.

### HasCheckAgent

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasCheckAgent() bool`

HasCheckAgent returns a boolean if a field has been set.

### SetCheckAgentNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckAgentNil(b bool)`

 SetCheckAgentNil sets the value for CheckAgent to be an explicit nil

### UnsetCheckAgent
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetCheckAgent()`

UnsetCheckAgent ensures that no value is present for CheckAgent, not even an explicit nil
### GetCheckInterval

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckIntervalOk() (*int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckInterval(v int64)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### SetCheckIntervalNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckIntervalNil(b bool)`

 SetCheckIntervalNil sets the value for CheckInterval to be an explicit nil

### UnsetCheckInterval
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetCheckInterval()`

UnsetCheckInterval ensures that no value is present for CheckInterval, not even an explicit nil
### GetCheckSpec

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckSpec() string`

GetCheckSpec returns the CheckSpec field if non-nil, zero value otherwise.

### GetCheckSpecOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckSpecOk() (*string, bool)`

GetCheckSpecOk returns a tuple with the CheckSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSpec

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckSpec(v string)`

SetCheckSpec sets CheckSpec field to given value.

### HasCheckSpec

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasCheckSpec() bool`

HasCheckSpec returns a boolean if a field has been set.

### SetCheckSpecNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckSpecNil(b bool)`

 SetCheckSpecNil sets the value for CheckSpec to be an explicit nil

### UnsetCheckSpec
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetCheckSpec()`

UnsetCheckSpec ensures that no value is present for CheckSpec, not even an explicit nil
### GetCheckType

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckType() ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCheckTypeOk() (*ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCheckType(v ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetConfig

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetConfig() ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetConfigOk() (*ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetConfig(v ListIncidents200ResponseAllOfIncidentsInnerChecksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetContainer() ListIncidents200ResponseAllOfIncidentsInnerChecksInnerContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetContainerOk() (*ListIncidents200ResponseAllOfIncidentsInnerChecksInnerContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetContainer(v ListIncidents200ResponseAllOfIncidentsInnerChecksInnerContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreateIncident

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCreatedBy() ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetCreatedByOk() (*ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetCreatedBy(v ListIncidents200ResponseAllOfIncidentsInnerChecksInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetHealth

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastBoxStats

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastBoxStats() string`

GetLastBoxStats returns the LastBoxStats field if non-nil, zero value otherwise.

### GetLastBoxStatsOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastBoxStatsOk() (*string, bool)`

GetLastBoxStatsOk returns a tuple with the LastBoxStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBoxStats

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastBoxStats(v string)`

SetLastBoxStats sets LastBoxStats field to given value.

### HasLastBoxStats

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastBoxStats() bool`

HasLastBoxStats returns a boolean if a field has been set.

### SetLastBoxStatsNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastBoxStatsNil(b bool)`

 SetLastBoxStatsNil sets the value for LastBoxStats to be an explicit nil

### UnsetLastBoxStats
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastBoxStats()`

UnsetLastBoxStats ensures that no value is present for LastBoxStats, not even an explicit nil
### GetLastCheckStatus

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastError

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastErrorDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastMessage

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetLastMetric

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetLastRunDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastStats

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### SetLastStatsNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastStatsNil(b bool)`

 SetLastStatsNil sets the value for LastStats to be an explicit nil

### UnsetLastStats
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastStats()`

UnsetLastStats ensures that no value is present for LastStats, not even an explicit nil
### GetLastSuccessDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastTimer

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### SetLastTimerNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastTimerNil(b bool)`

 SetLastTimerNil sets the value for LastTimer to be an explicit nil

### UnsetLastTimer
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastTimer()`

UnsetLastTimer ensures that no value is present for LastTimer, not even an explicit nil
### GetLastUpdated

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastWarningDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetOutageTime

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetSeverity

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStartDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDeleted

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ListIncidents200ResponseAllOfIncidentsInnerChecksInner) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


