# ListIncidents200ResponseAllOfIssuesInner

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

### NewListIncidents200ResponseAllOfIssuesInner

`func NewListIncidents200ResponseAllOfIssuesInner() *ListIncidents200ResponseAllOfIssuesInner`

NewListIncidents200ResponseAllOfIssuesInner instantiates a new ListIncidents200ResponseAllOfIssuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncidents200ResponseAllOfIssuesInnerWithDefaults

`func NewListIncidents200ResponseAllOfIssuesInnerWithDefaults() *ListIncidents200ResponseAllOfIssuesInner`

NewListIncidents200ResponseAllOfIssuesInnerWithDefaults instantiates a new ListIncidents200ResponseAllOfIssuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttachmentType

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetAttachmentType() string`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetAttachmentTypeOk() (*string, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetAttachmentType(v string)`

SetAttachmentType sets AttachmentType field to given value.

### HasAttachmentType

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasAttachmentType() bool`

HasAttachmentType returns a boolean if a field has been set.

### GetApp

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAvailable

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCheck

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckGroup

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetCheckGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCheckGroup returns the CheckGroup field if non-nil, zero value otherwise.

### GetCheckGroupOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetCheckGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCheckGroupOk returns a tuple with the CheckGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroup

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetCheckGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCheckGroup sets CheckGroup field to given value.

### HasCheckGroup

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasCheckGroup() bool`

HasCheckGroup returns a boolean if a field has been set.

### GetCheckStatus

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetCheckStatus() string`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetCheckStatusOk() (*string, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetCheckStatus(v string)`

SetCheckStatus sets CheckStatus field to given value.

### HasCheckStatus

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasCheckStatus() bool`

HasCheckStatus returns a boolean if a field has been set.

### GetEndDate

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetHealth

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInUptime

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetIncident

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetIncident() GetAlerts200ResponseAllOfChecksInnerAccount`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetIncidentOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetIncident(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetLastCheckTime

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.

### HasLastCheckTime

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasLastCheckTime() bool`

HasLastCheckTime returns a boolean if a field has been set.

### GetLastError

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetLastMessage

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetName

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityId

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetSeverityId() int64`

GetSeverityId returns the SeverityId field if non-nil, zero value otherwise.

### GetSeverityIdOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetSeverityIdOk() (*int64, bool)`

GetSeverityIdOk returns a tuple with the SeverityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityId

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetSeverityId(v int64)`

SetSeverityId sets SeverityId field to given value.

### HasSeverityId

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasSeverityId() bool`

HasSeverityId returns a boolean if a field has been set.

### GetStartDate

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListIncidents200ResponseAllOfIssuesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListIncidents200ResponseAllOfIssuesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListIncidents200ResponseAllOfIssuesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


