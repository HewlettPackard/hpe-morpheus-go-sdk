# UpdateAlerts200ResponseAllOfAlert

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
**Contacts** | Pointer to [**[]UpdateAlerts200ResponseAllOfAlertContactsInner**](UpdateAlerts200ResponseAllOfAlertContactsInner.md) |  | [optional] 

## Methods

### NewUpdateAlerts200ResponseAllOfAlert

`func NewUpdateAlerts200ResponseAllOfAlert() *UpdateAlerts200ResponseAllOfAlert`

NewUpdateAlerts200ResponseAllOfAlert instantiates a new UpdateAlerts200ResponseAllOfAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlerts200ResponseAllOfAlertWithDefaults

`func NewUpdateAlerts200ResponseAllOfAlertWithDefaults() *UpdateAlerts200ResponseAllOfAlert`

NewUpdateAlerts200ResponseAllOfAlertWithDefaults instantiates a new UpdateAlerts200ResponseAllOfAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAlerts200ResponseAllOfAlert) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAlerts200ResponseAllOfAlert) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAlerts200ResponseAllOfAlert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAlerts200ResponseAllOfAlert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAlerts200ResponseAllOfAlert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAlerts200ResponseAllOfAlert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllApps

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAllApps() bool`

GetAllApps returns the AllApps field if non-nil, zero value otherwise.

### GetAllAppsOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAllAppsOk() (*bool, bool)`

GetAllAppsOk returns a tuple with the AllApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllApps

`func (o *UpdateAlerts200ResponseAllOfAlert) SetAllApps(v bool)`

SetAllApps sets AllApps field to given value.

### HasAllApps

`func (o *UpdateAlerts200ResponseAllOfAlert) HasAllApps() bool`

HasAllApps returns a boolean if a field has been set.

### GetAllChecks

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAllChecks() bool`

GetAllChecks returns the AllChecks field if non-nil, zero value otherwise.

### GetAllChecksOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAllChecksOk() (*bool, bool)`

GetAllChecksOk returns a tuple with the AllChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllChecks

`func (o *UpdateAlerts200ResponseAllOfAlert) SetAllChecks(v bool)`

SetAllChecks sets AllChecks field to given value.

### HasAllChecks

`func (o *UpdateAlerts200ResponseAllOfAlert) HasAllChecks() bool`

HasAllChecks returns a boolean if a field has been set.

### GetAllGroups

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *UpdateAlerts200ResponseAllOfAlert) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *UpdateAlerts200ResponseAllOfAlert) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetActive

`func (o *UpdateAlerts200ResponseAllOfAlert) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateAlerts200ResponseAllOfAlert) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateAlerts200ResponseAllOfAlert) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMinSeverity

`func (o *UpdateAlerts200ResponseAllOfAlert) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *UpdateAlerts200ResponseAllOfAlert) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *UpdateAlerts200ResponseAllOfAlert) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetMinDuration

`func (o *UpdateAlerts200ResponseAllOfAlert) GetMinDuration() int64`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetMinDurationOk() (*int64, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *UpdateAlerts200ResponseAllOfAlert) SetMinDuration(v int64)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *UpdateAlerts200ResponseAllOfAlert) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateAlerts200ResponseAllOfAlert) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateAlerts200ResponseAllOfAlert) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateAlerts200ResponseAllOfAlert) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateAlerts200ResponseAllOfAlert) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateAlerts200ResponseAllOfAlert) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateAlerts200ResponseAllOfAlert) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetChecks

`func (o *UpdateAlerts200ResponseAllOfAlert) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *UpdateAlerts200ResponseAllOfAlert) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *UpdateAlerts200ResponseAllOfAlert) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### SetChecksNil

`func (o *UpdateAlerts200ResponseAllOfAlert) SetChecksNil(b bool)`

 SetChecksNil sets the value for Checks to be an explicit nil

### UnsetChecks
`func (o *UpdateAlerts200ResponseAllOfAlert) UnsetChecks()`

UnsetChecks ensures that no value is present for Checks, not even an explicit nil
### GetCheckGroups

`func (o *UpdateAlerts200ResponseAllOfAlert) GetCheckGroups() []int32`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetCheckGroupsOk() (*[]int32, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *UpdateAlerts200ResponseAllOfAlert) SetCheckGroups(v []int32)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *UpdateAlerts200ResponseAllOfAlert) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### SetCheckGroupsNil

`func (o *UpdateAlerts200ResponseAllOfAlert) SetCheckGroupsNil(b bool)`

 SetCheckGroupsNil sets the value for CheckGroups to be an explicit nil

### UnsetCheckGroups
`func (o *UpdateAlerts200ResponseAllOfAlert) UnsetCheckGroups()`

UnsetCheckGroups ensures that no value is present for CheckGroups, not even an explicit nil
### GetApps

`func (o *UpdateAlerts200ResponseAllOfAlert) GetApps() []int32`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetAppsOk() (*[]int32, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *UpdateAlerts200ResponseAllOfAlert) SetApps(v []int32)`

SetApps sets Apps field to given value.

### HasApps

`func (o *UpdateAlerts200ResponseAllOfAlert) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *UpdateAlerts200ResponseAllOfAlert) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *UpdateAlerts200ResponseAllOfAlert) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetContacts

`func (o *UpdateAlerts200ResponseAllOfAlert) GetContacts() []UpdateAlerts200ResponseAllOfAlertContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *UpdateAlerts200ResponseAllOfAlert) GetContactsOk() (*[]UpdateAlerts200ResponseAllOfAlertContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *UpdateAlerts200ResponseAllOfAlert) SetContacts(v []UpdateAlerts200ResponseAllOfAlertContactsInner)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *UpdateAlerts200ResponseAllOfAlert) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


