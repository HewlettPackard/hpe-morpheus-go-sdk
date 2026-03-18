# GetWorkflows200ResponseAllOfTaskSet

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
**OptionTypes** | Pointer to [**[]GetWorkflows200ResponseAllOfTaskSetOptionTypesInner**](GetWorkflows200ResponseAllOfTaskSetOptionTypesInner.md) |  | [optional] 
**TaskSetTasks** | Pointer to [**[]GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInner**](GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInner.md) |  | [optional] 

## Methods

### NewGetWorkflows200ResponseAllOfTaskSet

`func NewGetWorkflows200ResponseAllOfTaskSet() *GetWorkflows200ResponseAllOfTaskSet`

NewGetWorkflows200ResponseAllOfTaskSet instantiates a new GetWorkflows200ResponseAllOfTaskSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflows200ResponseAllOfTaskSetWithDefaults

`func NewGetWorkflows200ResponseAllOfTaskSetWithDefaults() *GetWorkflows200ResponseAllOfTaskSet`

NewGetWorkflows200ResponseAllOfTaskSetWithDefaults instantiates a new GetWorkflows200ResponseAllOfTaskSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetWorkflows200ResponseAllOfTaskSet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetWorkflows200ResponseAllOfTaskSet) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDateCreated

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccountId

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GetWorkflows200ResponseAllOfTaskSet) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetVisibility

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetTasks

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetTasks() []int64`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetTasksOk() (*[]int64, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetTasks(v []int64)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetOptionTypes() []GetWorkflows200ResponseAllOfTaskSetOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetOptionTypesOk() (*[]GetWorkflows200ResponseAllOfTaskSetOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetOptionTypes(v []GetWorkflows200ResponseAllOfTaskSetOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *GetWorkflows200ResponseAllOfTaskSet) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetTaskSetTasks

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetTaskSetTasks() []GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInner`

GetTaskSetTasks returns the TaskSetTasks field if non-nil, zero value otherwise.

### GetTaskSetTasksOk

`func (o *GetWorkflows200ResponseAllOfTaskSet) GetTaskSetTasksOk() (*[]GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInner, bool)`

GetTaskSetTasksOk returns a tuple with the TaskSetTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetTasks

`func (o *GetWorkflows200ResponseAllOfTaskSet) SetTaskSetTasks(v []GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInner)`

SetTaskSetTasks sets TaskSetTasks field to given value.

### HasTaskSetTasks

`func (o *GetWorkflows200ResponseAllOfTaskSet) HasTaskSetTasks() bool`

HasTaskSetTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


