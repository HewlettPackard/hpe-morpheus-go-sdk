# ListAlerts200ResponseAllOfAlertsInner

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

### NewListAlerts200ResponseAllOfAlertsInner

`func NewListAlerts200ResponseAllOfAlertsInner() *ListAlerts200ResponseAllOfAlertsInner`

NewListAlerts200ResponseAllOfAlertsInner instantiates a new ListAlerts200ResponseAllOfAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlerts200ResponseAllOfAlertsInnerWithDefaults

`func NewListAlerts200ResponseAllOfAlertsInnerWithDefaults() *ListAlerts200ResponseAllOfAlertsInner`

NewListAlerts200ResponseAllOfAlertsInnerWithDefaults instantiates a new ListAlerts200ResponseAllOfAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllApps

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAllApps() bool`

GetAllApps returns the AllApps field if non-nil, zero value otherwise.

### GetAllAppsOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAllAppsOk() (*bool, bool)`

GetAllAppsOk returns a tuple with the AllApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllApps

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetAllApps(v bool)`

SetAllApps sets AllApps field to given value.

### HasAllApps

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasAllApps() bool`

HasAllApps returns a boolean if a field has been set.

### GetAllChecks

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAllChecks() bool`

GetAllChecks returns the AllChecks field if non-nil, zero value otherwise.

### GetAllChecksOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAllChecksOk() (*bool, bool)`

GetAllChecksOk returns a tuple with the AllChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllChecks

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetAllChecks(v bool)`

SetAllChecks sets AllChecks field to given value.

### HasAllChecks

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasAllChecks() bool`

HasAllChecks returns a boolean if a field has been set.

### GetAllGroups

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetActive

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMinSeverity

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetMinDuration

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetMinDuration() int64`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetMinDurationOk() (*int64, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetMinDuration(v int64)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetChecks

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetCheckGroups() []int32`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetCheckGroupsOk() (*[]int32, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetCheckGroups(v []int32)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetApps

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetApps() []int32`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetAppsOk() (*[]int32, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetApps(v []int32)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetContacts

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetContacts() []ListAlerts200ResponseAllOfAlertsInnerContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ListAlerts200ResponseAllOfAlertsInner) GetContactsOk() (*[]ListAlerts200ResponseAllOfAlertsInnerContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ListAlerts200ResponseAllOfAlertsInner) SetContacts(v []ListAlerts200ResponseAllOfAlertsInnerContactsInner)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ListAlerts200ResponseAllOfAlertsInner) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


