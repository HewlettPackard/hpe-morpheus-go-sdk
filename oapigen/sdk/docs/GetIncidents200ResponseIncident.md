# GetIncidents200ResponseIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetIncidents200ResponseIncidentAccount**](GetIncidents200ResponseIncidentAccount.md) |  | [optional] 
**App** | Pointer to **NullableString** |  | [optional] 
**AutoClose** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CheckGroups** | Pointer to [**[]GetIncidents200ResponseIncidentCheckGroupsInner**](GetIncidents200ResponseIncidentCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]GetIncidents200ResponseIncidentChecksInner**](GetIncidents200ResponseIncidentChecksInner.md) |  | [optional] 
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

### NewGetIncidents200ResponseIncident

`func NewGetIncidents200ResponseIncident() *GetIncidents200ResponseIncident`

NewGetIncidents200ResponseIncident instantiates a new GetIncidents200ResponseIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncidents200ResponseIncidentWithDefaults

`func NewGetIncidents200ResponseIncidentWithDefaults() *GetIncidents200ResponseIncident`

NewGetIncidents200ResponseIncidentWithDefaults instantiates a new GetIncidents200ResponseIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIncidents200ResponseIncident) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIncidents200ResponseIncident) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIncidents200ResponseIncident) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIncidents200ResponseIncident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetIncidents200ResponseIncident) GetAccount() GetIncidents200ResponseIncidentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetIncidents200ResponseIncident) GetAccountOk() (*GetIncidents200ResponseIncidentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetIncidents200ResponseIncident) SetAccount(v GetIncidents200ResponseIncidentAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetIncidents200ResponseIncident) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApp

`func (o *GetIncidents200ResponseIncident) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GetIncidents200ResponseIncident) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GetIncidents200ResponseIncident) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *GetIncidents200ResponseIncident) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *GetIncidents200ResponseIncident) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *GetIncidents200ResponseIncident) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetAutoClose

`func (o *GetIncidents200ResponseIncident) GetAutoClose() bool`

GetAutoClose returns the AutoClose field if non-nil, zero value otherwise.

### GetAutoCloseOk

`func (o *GetIncidents200ResponseIncident) GetAutoCloseOk() (*bool, bool)`

GetAutoCloseOk returns a tuple with the AutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClose

`func (o *GetIncidents200ResponseIncident) SetAutoClose(v bool)`

SetAutoClose sets AutoClose field to given value.

### HasAutoClose

`func (o *GetIncidents200ResponseIncident) HasAutoClose() bool`

HasAutoClose returns a boolean if a field has been set.

### GetChannelId

`func (o *GetIncidents200ResponseIncident) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GetIncidents200ResponseIncident) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GetIncidents200ResponseIncident) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *GetIncidents200ResponseIncident) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetIncidents200ResponseIncident) GetCheckGroups() []GetIncidents200ResponseIncidentCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetIncidents200ResponseIncident) GetCheckGroupsOk() (*[]GetIncidents200ResponseIncidentCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetIncidents200ResponseIncident) SetCheckGroups(v []GetIncidents200ResponseIncidentCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetIncidents200ResponseIncident) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *GetIncidents200ResponseIncident) GetChecks() []GetIncidents200ResponseIncidentChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetIncidents200ResponseIncident) GetChecksOk() (*[]GetIncidents200ResponseIncidentChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetIncidents200ResponseIncident) SetChecks(v []GetIncidents200ResponseIncidentChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetIncidents200ResponseIncident) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetComment

`func (o *GetIncidents200ResponseIncident) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIncidents200ResponseIncident) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIncidents200ResponseIncident) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIncidents200ResponseIncident) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *GetIncidents200ResponseIncident) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *GetIncidents200ResponseIncident) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDisplayName

`func (o *GetIncidents200ResponseIncident) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetIncidents200ResponseIncident) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetIncidents200ResponseIncident) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetIncidents200ResponseIncident) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDuration

`func (o *GetIncidents200ResponseIncident) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetIncidents200ResponseIncident) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetIncidents200ResponseIncident) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetIncidents200ResponseIncident) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *GetIncidents200ResponseIncident) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *GetIncidents200ResponseIncident) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndDate

`func (o *GetIncidents200ResponseIncident) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetIncidents200ResponseIncident) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetIncidents200ResponseIncident) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetIncidents200ResponseIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetIncidents200ResponseIncident) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetIncidents200ResponseIncident) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetInUptime

`func (o *GetIncidents200ResponseIncident) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetIncidents200ResponseIncident) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetIncidents200ResponseIncident) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetIncidents200ResponseIncident) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetMuted

`func (o *GetIncidents200ResponseIncident) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetIncidents200ResponseIncident) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetIncidents200ResponseIncident) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetIncidents200ResponseIncident) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *GetIncidents200ResponseIncident) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *GetIncidents200ResponseIncident) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *GetIncidents200ResponseIncident) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *GetIncidents200ResponseIncident) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *GetIncidents200ResponseIncident) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetIncidents200ResponseIncident) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetIncidents200ResponseIncident) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetIncidents200ResponseIncident) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *GetIncidents200ResponseIncident) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *GetIncidents200ResponseIncident) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *GetIncidents200ResponseIncident) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *GetIncidents200ResponseIncident) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *GetIncidents200ResponseIncident) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *GetIncidents200ResponseIncident) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetName

`func (o *GetIncidents200ResponseIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIncidents200ResponseIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIncidents200ResponseIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIncidents200ResponseIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolution

`func (o *GetIncidents200ResponseIncident) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *GetIncidents200ResponseIncident) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *GetIncidents200ResponseIncident) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *GetIncidents200ResponseIncident) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *GetIncidents200ResponseIncident) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *GetIncidents200ResponseIncident) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSeverity

`func (o *GetIncidents200ResponseIncident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetIncidents200ResponseIncident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetIncidents200ResponseIncident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetIncidents200ResponseIncident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *GetIncidents200ResponseIncident) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *GetIncidents200ResponseIncident) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *GetIncidents200ResponseIncident) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *GetIncidents200ResponseIncident) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *GetIncidents200ResponseIncident) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetIncidents200ResponseIncident) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetIncidents200ResponseIncident) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetIncidents200ResponseIncident) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetIncidents200ResponseIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIncidents200ResponseIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIncidents200ResponseIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIncidents200ResponseIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *GetIncidents200ResponseIncident) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetIncidents200ResponseIncident) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetIncidents200ResponseIncident) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetIncidents200ResponseIncident) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


