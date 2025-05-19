# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AttachmentType** | Pointer to **string** |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**Check** | Pointer to **string** |  | [optional] 
**CheckGroup** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**CheckStatus** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**Incident** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**LastCheckTime** | Pointer to **time.Time** |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**LastMessage** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SeverityId** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewIssue

`func NewIssue() *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Issue) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Issue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttachmentType

`func (o *Issue) GetAttachmentType() string`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *Issue) GetAttachmentTypeOk() (*string, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *Issue) SetAttachmentType(v string)`

SetAttachmentType sets AttachmentType field to given value.

### HasAttachmentType

`func (o *Issue) HasAttachmentType() bool`

HasAttachmentType returns a boolean if a field has been set.

### GetApp

`func (o *Issue) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Issue) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Issue) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *Issue) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAvailable

`func (o *Issue) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Issue) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Issue) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Issue) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCheck

`func (o *Issue) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *Issue) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *Issue) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *Issue) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckGroup

`func (o *Issue) GetCheckGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCheckGroup returns the CheckGroup field if non-nil, zero value otherwise.

### GetCheckGroupOk

`func (o *Issue) GetCheckGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCheckGroupOk returns a tuple with the CheckGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroup

`func (o *Issue) SetCheckGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCheckGroup sets CheckGroup field to given value.

### HasCheckGroup

`func (o *Issue) HasCheckGroup() bool`

HasCheckGroup returns a boolean if a field has been set.

### GetCheckStatus

`func (o *Issue) GetCheckStatus() string`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *Issue) GetCheckStatusOk() (*string, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *Issue) SetCheckStatus(v string)`

SetCheckStatus sets CheckStatus field to given value.

### HasCheckStatus

`func (o *Issue) HasCheckStatus() bool`

HasCheckStatus returns a boolean if a field has been set.

### GetEndDate

`func (o *Issue) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Issue) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Issue) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Issue) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetHealth

`func (o *Issue) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Issue) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Issue) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *Issue) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *Issue) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *Issue) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *Issue) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *Issue) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetIncident

`func (o *Issue) GetIncident() GetAlerts200ResponseAllOfChecksInnerAccount`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *Issue) GetIncidentOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *Issue) SetIncident(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *Issue) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *Issue) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *Issue) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *Issue) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *Issue) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *Issue) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *Issue) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *Issue) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *Issue) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *Issue) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *Issue) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *Issue) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *Issue) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetName

`func (o *Issue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Issue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Issue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Issue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *Issue) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Issue) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Issue) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Issue) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *Issue) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *Issue) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *Issue) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *Issue) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *Issue) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Issue) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Issue) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Issue) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *Issue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Issue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Issue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Issue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


