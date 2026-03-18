# UpdateIncidents200ResponseAllOfIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**UpdateIncidents200ResponseAllOfIncidentAccount**](UpdateIncidents200ResponseAllOfIncidentAccount.md) |  | [optional] 
**App** | Pointer to **NullableString** |  | [optional] 
**AutoClose** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CheckGroups** | Pointer to [**[]UpdateIncidents200ResponseAllOfIncidentCheckGroupsInner**](UpdateIncidents200ResponseAllOfIncidentCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]UpdateIncidents200ResponseAllOfIncidentChecksInner**](UpdateIncidents200ResponseAllOfIncidentChecksInner.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**LastCheckTime** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastMessage** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SeverityId** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateIncidents200ResponseAllOfIncident

`func NewUpdateIncidents200ResponseAllOfIncident() *UpdateIncidents200ResponseAllOfIncident`

NewUpdateIncidents200ResponseAllOfIncident instantiates a new UpdateIncidents200ResponseAllOfIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIncidents200ResponseAllOfIncidentWithDefaults

`func NewUpdateIncidents200ResponseAllOfIncidentWithDefaults() *UpdateIncidents200ResponseAllOfIncident`

NewUpdateIncidents200ResponseAllOfIncidentWithDefaults instantiates a new UpdateIncidents200ResponseAllOfIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateIncidents200ResponseAllOfIncident) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateIncidents200ResponseAllOfIncident) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateIncidents200ResponseAllOfIncident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateIncidents200ResponseAllOfIncident) GetAccount() UpdateIncidents200ResponseAllOfIncidentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetAccountOk() (*UpdateIncidents200ResponseAllOfIncidentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIncidents200ResponseAllOfIncident) SetAccount(v UpdateIncidents200ResponseAllOfIncidentAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIncidents200ResponseAllOfIncident) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApp

`func (o *UpdateIncidents200ResponseAllOfIncident) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *UpdateIncidents200ResponseAllOfIncident) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *UpdateIncidents200ResponseAllOfIncident) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *UpdateIncidents200ResponseAllOfIncident) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *UpdateIncidents200ResponseAllOfIncident) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetAutoClose

`func (o *UpdateIncidents200ResponseAllOfIncident) GetAutoClose() bool`

GetAutoClose returns the AutoClose field if non-nil, zero value otherwise.

### GetAutoCloseOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetAutoCloseOk() (*bool, bool)`

GetAutoCloseOk returns a tuple with the AutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClose

`func (o *UpdateIncidents200ResponseAllOfIncident) SetAutoClose(v bool)`

SetAutoClose sets AutoClose field to given value.

### HasAutoClose

`func (o *UpdateIncidents200ResponseAllOfIncident) HasAutoClose() bool`

HasAutoClose returns a boolean if a field has been set.

### GetChannelId

`func (o *UpdateIncidents200ResponseAllOfIncident) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateIncidents200ResponseAllOfIncident) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateIncidents200ResponseAllOfIncident) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCheckGroups

`func (o *UpdateIncidents200ResponseAllOfIncident) GetCheckGroups() []UpdateIncidents200ResponseAllOfIncidentCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetCheckGroupsOk() (*[]UpdateIncidents200ResponseAllOfIncidentCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *UpdateIncidents200ResponseAllOfIncident) SetCheckGroups(v []UpdateIncidents200ResponseAllOfIncidentCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *UpdateIncidents200ResponseAllOfIncident) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *UpdateIncidents200ResponseAllOfIncident) GetChecks() []UpdateIncidents200ResponseAllOfIncidentChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetChecksOk() (*[]UpdateIncidents200ResponseAllOfIncidentChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *UpdateIncidents200ResponseAllOfIncident) SetChecks(v []UpdateIncidents200ResponseAllOfIncidentChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *UpdateIncidents200ResponseAllOfIncident) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetComment

`func (o *UpdateIncidents200ResponseAllOfIncident) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateIncidents200ResponseAllOfIncident) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateIncidents200ResponseAllOfIncident) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *UpdateIncidents200ResponseAllOfIncident) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UpdateIncidents200ResponseAllOfIncident) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDisplayName

`func (o *UpdateIncidents200ResponseAllOfIncident) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateIncidents200ResponseAllOfIncident) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateIncidents200ResponseAllOfIncident) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDuration

`func (o *UpdateIncidents200ResponseAllOfIncident) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateIncidents200ResponseAllOfIncident) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UpdateIncidents200ResponseAllOfIncident) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *UpdateIncidents200ResponseAllOfIncident) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *UpdateIncidents200ResponseAllOfIncident) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndDate

`func (o *UpdateIncidents200ResponseAllOfIncident) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateIncidents200ResponseAllOfIncident) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateIncidents200ResponseAllOfIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UpdateIncidents200ResponseAllOfIncident) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UpdateIncidents200ResponseAllOfIncident) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetInUptime

`func (o *UpdateIncidents200ResponseAllOfIncident) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateIncidents200ResponseAllOfIncident) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateIncidents200ResponseAllOfIncident) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetMuted

`func (o *UpdateIncidents200ResponseAllOfIncident) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *UpdateIncidents200ResponseAllOfIncident) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *UpdateIncidents200ResponseAllOfIncident) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *UpdateIncidents200ResponseAllOfIncident) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *UpdateIncidents200ResponseAllOfIncident) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *UpdateIncidents200ResponseAllOfIncident) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *UpdateIncidents200ResponseAllOfIncident) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *UpdateIncidents200ResponseAllOfIncident) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *UpdateIncidents200ResponseAllOfIncident) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *UpdateIncidents200ResponseAllOfIncident) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *UpdateIncidents200ResponseAllOfIncident) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *UpdateIncidents200ResponseAllOfIncident) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *UpdateIncidents200ResponseAllOfIncident) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *UpdateIncidents200ResponseAllOfIncident) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetName

`func (o *UpdateIncidents200ResponseAllOfIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIncidents200ResponseAllOfIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIncidents200ResponseAllOfIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolution

`func (o *UpdateIncidents200ResponseAllOfIncident) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *UpdateIncidents200ResponseAllOfIncident) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *UpdateIncidents200ResponseAllOfIncident) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *UpdateIncidents200ResponseAllOfIncident) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *UpdateIncidents200ResponseAllOfIncident) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSeverity

`func (o *UpdateIncidents200ResponseAllOfIncident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateIncidents200ResponseAllOfIncident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateIncidents200ResponseAllOfIncident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *UpdateIncidents200ResponseAllOfIncident) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *UpdateIncidents200ResponseAllOfIncident) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *UpdateIncidents200ResponseAllOfIncident) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateIncidents200ResponseAllOfIncident) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateIncidents200ResponseAllOfIncident) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateIncidents200ResponseAllOfIncident) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateIncidents200ResponseAllOfIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateIncidents200ResponseAllOfIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateIncidents200ResponseAllOfIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateIncidents200ResponseAllOfIncident) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateIncidents200ResponseAllOfIncident) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateIncidents200ResponseAllOfIncident) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateIncidents200ResponseAllOfIncident) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


