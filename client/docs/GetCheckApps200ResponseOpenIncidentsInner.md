# GetCheckApps200ResponseOpenIncidentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**AutoClose** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CheckGroups** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Checks** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInner**](GetAlerts200ResponseAllOfChecksInner.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**LastCheckTime** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastMessage** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SeverityId** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCheckApps200ResponseOpenIncidentsInner

`func NewGetCheckApps200ResponseOpenIncidentsInner() *GetCheckApps200ResponseOpenIncidentsInner`

NewGetCheckApps200ResponseOpenIncidentsInner instantiates a new GetCheckApps200ResponseOpenIncidentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCheckApps200ResponseOpenIncidentsInnerWithDefaults

`func NewGetCheckApps200ResponseOpenIncidentsInnerWithDefaults() *GetCheckApps200ResponseOpenIncidentsInner`

NewGetCheckApps200ResponseOpenIncidentsInnerWithDefaults instantiates a new GetCheckApps200ResponseOpenIncidentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApp

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAutoClose

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetAutoClose() bool`

GetAutoClose returns the AutoClose field if non-nil, zero value otherwise.

### GetAutoCloseOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetAutoCloseOk() (*bool, bool)`

GetAutoCloseOk returns a tuple with the AutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClose

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetAutoClose(v bool)`

SetAutoClose sets AutoClose field to given value.

### HasAutoClose

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasAutoClose() bool`

HasAutoClose returns a boolean if a field has been set.

### GetChannelId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetCheckGroups() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetCheckGroupsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetCheckGroups(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetChecks() []GetAlerts200ResponseAllOfChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetChecksOk() (*[]GetAlerts200ResponseAllOfChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetChecks(v []GetAlerts200ResponseAllOfChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetComment

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDuration

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndDate

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInUptime

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetMuted

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetName

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolution

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeverity

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetCheckApps200ResponseOpenIncidentsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetCheckApps200ResponseOpenIncidentsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetCheckApps200ResponseOpenIncidentsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


