# AddIncident200ResponseAllOfIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**AddIncident200ResponseAllOfIncidentAccount**](AddIncident200ResponseAllOfIncidentAccount.md) |  | [optional] 
**App** | Pointer to **NullableString** |  | [optional] 
**AutoClose** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CheckGroups** | Pointer to [**[]AddIncident200ResponseAllOfIncidentCheckGroupsInner**](AddIncident200ResponseAllOfIncidentCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]AddIncident200ResponseAllOfIncidentChecksInner**](AddIncident200ResponseAllOfIncidentChecksInner.md) |  | [optional] 
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

### NewAddIncident200ResponseAllOfIncident

`func NewAddIncident200ResponseAllOfIncident() *AddIncident200ResponseAllOfIncident`

NewAddIncident200ResponseAllOfIncident instantiates a new AddIncident200ResponseAllOfIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIncident200ResponseAllOfIncidentWithDefaults

`func NewAddIncident200ResponseAllOfIncidentWithDefaults() *AddIncident200ResponseAllOfIncident`

NewAddIncident200ResponseAllOfIncidentWithDefaults instantiates a new AddIncident200ResponseAllOfIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIncident200ResponseAllOfIncident) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIncident200ResponseAllOfIncident) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIncident200ResponseAllOfIncident) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIncident200ResponseAllOfIncident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AddIncident200ResponseAllOfIncident) GetAccount() AddIncident200ResponseAllOfIncidentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddIncident200ResponseAllOfIncident) GetAccountOk() (*AddIncident200ResponseAllOfIncidentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddIncident200ResponseAllOfIncident) SetAccount(v AddIncident200ResponseAllOfIncidentAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddIncident200ResponseAllOfIncident) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApp

`func (o *AddIncident200ResponseAllOfIncident) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AddIncident200ResponseAllOfIncident) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AddIncident200ResponseAllOfIncident) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *AddIncident200ResponseAllOfIncident) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *AddIncident200ResponseAllOfIncident) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *AddIncident200ResponseAllOfIncident) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetAutoClose

`func (o *AddIncident200ResponseAllOfIncident) GetAutoClose() bool`

GetAutoClose returns the AutoClose field if non-nil, zero value otherwise.

### GetAutoCloseOk

`func (o *AddIncident200ResponseAllOfIncident) GetAutoCloseOk() (*bool, bool)`

GetAutoCloseOk returns a tuple with the AutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClose

`func (o *AddIncident200ResponseAllOfIncident) SetAutoClose(v bool)`

SetAutoClose sets AutoClose field to given value.

### HasAutoClose

`func (o *AddIncident200ResponseAllOfIncident) HasAutoClose() bool`

HasAutoClose returns a boolean if a field has been set.

### GetChannelId

`func (o *AddIncident200ResponseAllOfIncident) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AddIncident200ResponseAllOfIncident) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AddIncident200ResponseAllOfIncident) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *AddIncident200ResponseAllOfIncident) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCheckGroups

`func (o *AddIncident200ResponseAllOfIncident) GetCheckGroups() []AddIncident200ResponseAllOfIncidentCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *AddIncident200ResponseAllOfIncident) GetCheckGroupsOk() (*[]AddIncident200ResponseAllOfIncidentCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *AddIncident200ResponseAllOfIncident) SetCheckGroups(v []AddIncident200ResponseAllOfIncidentCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *AddIncident200ResponseAllOfIncident) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *AddIncident200ResponseAllOfIncident) GetChecks() []AddIncident200ResponseAllOfIncidentChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *AddIncident200ResponseAllOfIncident) GetChecksOk() (*[]AddIncident200ResponseAllOfIncidentChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *AddIncident200ResponseAllOfIncident) SetChecks(v []AddIncident200ResponseAllOfIncidentChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *AddIncident200ResponseAllOfIncident) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetComment

`func (o *AddIncident200ResponseAllOfIncident) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddIncident200ResponseAllOfIncident) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddIncident200ResponseAllOfIncident) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddIncident200ResponseAllOfIncident) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *AddIncident200ResponseAllOfIncident) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *AddIncident200ResponseAllOfIncident) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDisplayName

`func (o *AddIncident200ResponseAllOfIncident) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddIncident200ResponseAllOfIncident) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddIncident200ResponseAllOfIncident) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddIncident200ResponseAllOfIncident) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDuration

`func (o *AddIncident200ResponseAllOfIncident) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AddIncident200ResponseAllOfIncident) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AddIncident200ResponseAllOfIncident) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AddIncident200ResponseAllOfIncident) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AddIncident200ResponseAllOfIncident) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AddIncident200ResponseAllOfIncident) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndDate

`func (o *AddIncident200ResponseAllOfIncident) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AddIncident200ResponseAllOfIncident) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AddIncident200ResponseAllOfIncident) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AddIncident200ResponseAllOfIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *AddIncident200ResponseAllOfIncident) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *AddIncident200ResponseAllOfIncident) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetInUptime

`func (o *AddIncident200ResponseAllOfIncident) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *AddIncident200ResponseAllOfIncident) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *AddIncident200ResponseAllOfIncident) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *AddIncident200ResponseAllOfIncident) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetMuted

`func (o *AddIncident200ResponseAllOfIncident) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *AddIncident200ResponseAllOfIncident) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *AddIncident200ResponseAllOfIncident) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *AddIncident200ResponseAllOfIncident) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *AddIncident200ResponseAllOfIncident) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *AddIncident200ResponseAllOfIncident) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *AddIncident200ResponseAllOfIncident) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *AddIncident200ResponseAllOfIncident) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *AddIncident200ResponseAllOfIncident) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *AddIncident200ResponseAllOfIncident) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *AddIncident200ResponseAllOfIncident) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *AddIncident200ResponseAllOfIncident) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *AddIncident200ResponseAllOfIncident) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *AddIncident200ResponseAllOfIncident) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *AddIncident200ResponseAllOfIncident) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *AddIncident200ResponseAllOfIncident) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *AddIncident200ResponseAllOfIncident) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *AddIncident200ResponseAllOfIncident) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetName

`func (o *AddIncident200ResponseAllOfIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIncident200ResponseAllOfIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIncident200ResponseAllOfIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIncident200ResponseAllOfIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolution

`func (o *AddIncident200ResponseAllOfIncident) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AddIncident200ResponseAllOfIncident) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AddIncident200ResponseAllOfIncident) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AddIncident200ResponseAllOfIncident) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *AddIncident200ResponseAllOfIncident) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *AddIncident200ResponseAllOfIncident) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSeverity

`func (o *AddIncident200ResponseAllOfIncident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AddIncident200ResponseAllOfIncident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AddIncident200ResponseAllOfIncident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AddIncident200ResponseAllOfIncident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *AddIncident200ResponseAllOfIncident) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *AddIncident200ResponseAllOfIncident) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *AddIncident200ResponseAllOfIncident) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *AddIncident200ResponseAllOfIncident) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *AddIncident200ResponseAllOfIncident) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddIncident200ResponseAllOfIncident) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AddIncident200ResponseAllOfIncident) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AddIncident200ResponseAllOfIncident) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *AddIncident200ResponseAllOfIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddIncident200ResponseAllOfIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddIncident200ResponseAllOfIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddIncident200ResponseAllOfIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *AddIncident200ResponseAllOfIncident) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddIncident200ResponseAllOfIncident) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddIncident200ResponseAllOfIncident) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddIncident200ResponseAllOfIncident) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


