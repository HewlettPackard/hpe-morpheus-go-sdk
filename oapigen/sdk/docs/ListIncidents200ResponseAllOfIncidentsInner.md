# ListIncidents200ResponseAllOfIncidentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListIncidents200ResponseAllOfIncidentsInnerAccount**](ListIncidents200ResponseAllOfIncidentsInnerAccount.md) |  | [optional] 
**App** | Pointer to **NullableString** |  | [optional] 
**AutoClose** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CheckGroups** | Pointer to [**[]ListIncidents200ResponseAllOfIncidentsInnerCheckGroupsInner**](ListIncidents200ResponseAllOfIncidentsInnerCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]ListIncidents200ResponseAllOfIncidentsInnerChecksInner**](ListIncidents200ResponseAllOfIncidentsInnerChecksInner.md) |  | [optional] 
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

### NewListIncidents200ResponseAllOfIncidentsInner

`func NewListIncidents200ResponseAllOfIncidentsInner() *ListIncidents200ResponseAllOfIncidentsInner`

NewListIncidents200ResponseAllOfIncidentsInner instantiates a new ListIncidents200ResponseAllOfIncidentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncidents200ResponseAllOfIncidentsInnerWithDefaults

`func NewListIncidents200ResponseAllOfIncidentsInnerWithDefaults() *ListIncidents200ResponseAllOfIncidentsInner`

NewListIncidents200ResponseAllOfIncidentsInnerWithDefaults instantiates a new ListIncidents200ResponseAllOfIncidentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetAccount() ListIncidents200ResponseAllOfIncidentsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetAccountOk() (*ListIncidents200ResponseAllOfIncidentsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetAccount(v ListIncidents200ResponseAllOfIncidentsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApp

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *ListIncidents200ResponseAllOfIncidentsInner) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetAutoClose

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetAutoClose() bool`

GetAutoClose returns the AutoClose field if non-nil, zero value otherwise.

### GetAutoCloseOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetAutoCloseOk() (*bool, bool)`

GetAutoCloseOk returns a tuple with the AutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClose

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetAutoClose(v bool)`

SetAutoClose sets AutoClose field to given value.

### HasAutoClose

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasAutoClose() bool`

HasAutoClose returns a boolean if a field has been set.

### GetChannelId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCheckGroups

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetCheckGroups() []ListIncidents200ResponseAllOfIncidentsInnerCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetCheckGroupsOk() (*[]ListIncidents200ResponseAllOfIncidentsInnerCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetCheckGroups(v []ListIncidents200ResponseAllOfIncidentsInnerCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetChecks() []ListIncidents200ResponseAllOfIncidentsInnerChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetChecksOk() (*[]ListIncidents200ResponseAllOfIncidentsInnerChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetChecks(v []ListIncidents200ResponseAllOfIncidentsInnerChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetComment

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ListIncidents200ResponseAllOfIncidentsInner) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDisplayName

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDuration

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ListIncidents200ResponseAllOfIncidentsInner) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndDate

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ListIncidents200ResponseAllOfIncidentsInner) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetInUptime

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetMuted

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *ListIncidents200ResponseAllOfIncidentsInner) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetName

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolution

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *ListIncidents200ResponseAllOfIncidentsInner) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSeverity

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListIncidents200ResponseAllOfIncidentsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListIncidents200ResponseAllOfIncidentsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListIncidents200ResponseAllOfIncidentsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


