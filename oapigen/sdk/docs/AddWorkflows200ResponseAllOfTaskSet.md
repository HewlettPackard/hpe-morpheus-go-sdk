# AddWorkflows200ResponseAllOfTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Tasks** | Pointer to **[]int64** |  | [optional] 
**OptionTypes** | Pointer to [**[]AddWorkflows200ResponseAllOfTaskSetOptionTypesInner**](AddWorkflows200ResponseAllOfTaskSetOptionTypesInner.md) |  | [optional] 
**TaskSetTasks** | Pointer to [**[]AddWorkflows200ResponseAllOfTaskSetTaskSetTasksInner**](AddWorkflows200ResponseAllOfTaskSetTaskSetTasksInner.md) |  | [optional] 

## Methods

### NewAddWorkflows200ResponseAllOfTaskSet

`func NewAddWorkflows200ResponseAllOfTaskSet() *AddWorkflows200ResponseAllOfTaskSet`

NewAddWorkflows200ResponseAllOfTaskSet instantiates a new AddWorkflows200ResponseAllOfTaskSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddWorkflows200ResponseAllOfTaskSet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AddWorkflows200ResponseAllOfTaskSet) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDateCreated

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccountId

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *AddWorkflows200ResponseAllOfTaskSet) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetVisibility

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetTasks

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetTasks() []int64`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetTasksOk() (*[]int64, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetTasks(v []int64)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOptionTypes

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetOptionTypes() []AddWorkflows200ResponseAllOfTaskSetOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetOptionTypesOk() (*[]AddWorkflows200ResponseAllOfTaskSetOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetOptionTypes(v []AddWorkflows200ResponseAllOfTaskSetOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *AddWorkflows200ResponseAllOfTaskSet) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetTaskSetTasks

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetTaskSetTasks() []AddWorkflows200ResponseAllOfTaskSetTaskSetTasksInner`

GetTaskSetTasks returns the TaskSetTasks field if non-nil, zero value otherwise.

### GetTaskSetTasksOk

`func (o *AddWorkflows200ResponseAllOfTaskSet) GetTaskSetTasksOk() (*[]AddWorkflows200ResponseAllOfTaskSetTaskSetTasksInner, bool)`

GetTaskSetTasksOk returns a tuple with the TaskSetTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetTasks

`func (o *AddWorkflows200ResponseAllOfTaskSet) SetTaskSetTasks(v []AddWorkflows200ResponseAllOfTaskSetTaskSetTasksInner)`

SetTaskSetTasks sets TaskSetTasks field to given value.

### HasTaskSetTasks

`func (o *AddWorkflows200ResponseAllOfTaskSet) HasTaskSetTasks() bool`

HasTaskSetTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


