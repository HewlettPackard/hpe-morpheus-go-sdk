# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AllApps** | Pointer to **bool** |  | [optional] 
**AllChecks** | Pointer to **bool** |  | [optional] 
**AllGroups** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**MinSeverity** | Pointer to **string** |  | [optional] 
**MinDuration** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Checks** | Pointer to **[]int32** |  | [optional] 
**CheckGroups** | Pointer to **[]int32** |  | [optional] 
**Apps** | Pointer to **[]int32** |  | [optional] 
**Contacts** | Pointer to [**[]ListAlerts200ResponseAllOfAlertsInnerContactsInner**](ListAlerts200ResponseAllOfAlertsInnerContactsInner.md) |  | [optional] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Alert) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Alert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Alert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllApps

`func (o *Alert) GetAllApps() bool`

GetAllApps returns the AllApps field if non-nil, zero value otherwise.

### GetAllAppsOk

`func (o *Alert) GetAllAppsOk() (*bool, bool)`

GetAllAppsOk returns a tuple with the AllApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllApps

`func (o *Alert) SetAllApps(v bool)`

SetAllApps sets AllApps field to given value.

### HasAllApps

`func (o *Alert) HasAllApps() bool`

HasAllApps returns a boolean if a field has been set.

### GetAllChecks

`func (o *Alert) GetAllChecks() bool`

GetAllChecks returns the AllChecks field if non-nil, zero value otherwise.

### GetAllChecksOk

`func (o *Alert) GetAllChecksOk() (*bool, bool)`

GetAllChecksOk returns a tuple with the AllChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllChecks

`func (o *Alert) SetAllChecks(v bool)`

SetAllChecks sets AllChecks field to given value.

### HasAllChecks

`func (o *Alert) HasAllChecks() bool`

HasAllChecks returns a boolean if a field has been set.

### GetAllGroups

`func (o *Alert) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *Alert) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *Alert) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *Alert) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetActive

`func (o *Alert) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Alert) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Alert) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Alert) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMinSeverity

`func (o *Alert) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *Alert) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *Alert) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *Alert) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetMinDuration

`func (o *Alert) GetMinDuration() int64`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *Alert) GetMinDurationOk() (*int64, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *Alert) SetMinDuration(v int64)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *Alert) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *Alert) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Alert) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Alert) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Alert) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Alert) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Alert) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Alert) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Alert) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetChecks

`func (o *Alert) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *Alert) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *Alert) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *Alert) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *Alert) GetCheckGroups() []int32`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *Alert) GetCheckGroupsOk() (*[]int32, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *Alert) SetCheckGroups(v []int32)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *Alert) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetApps

`func (o *Alert) GetApps() []int32`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *Alert) GetAppsOk() (*[]int32, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *Alert) SetApps(v []int32)`

SetApps sets Apps field to given value.

### HasApps

`func (o *Alert) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetContacts

`func (o *Alert) GetContacts() []ListAlerts200ResponseAllOfAlertsInnerContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Alert) GetContactsOk() (*[]ListAlerts200ResponseAllOfAlertsInnerContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Alert) SetContacts(v []ListAlerts200ResponseAllOfAlertsInnerContactsInner)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Alert) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


