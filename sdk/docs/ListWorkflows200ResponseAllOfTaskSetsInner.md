# ListWorkflows200ResponseAllOfTaskSetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Tasks** | Pointer to **[]int64** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner**](ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner.md) |  | [optional] 
**TaskSetTasks** | Pointer to [**[]ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner**](ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner.md) |  | [optional] 

## Methods

### NewListWorkflows200ResponseAllOfTaskSetsInner

`func NewListWorkflows200ResponseAllOfTaskSetsInner() *ListWorkflows200ResponseAllOfTaskSetsInner`

NewListWorkflows200ResponseAllOfTaskSetsInner instantiates a new ListWorkflows200ResponseAllOfTaskSetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflows200ResponseAllOfTaskSetsInnerWithDefaults

`func NewListWorkflows200ResponseAllOfTaskSetsInnerWithDefaults() *ListWorkflows200ResponseAllOfTaskSetsInner`

NewListWorkflows200ResponseAllOfTaskSetsInnerWithDefaults instantiates a new ListWorkflows200ResponseAllOfTaskSetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccountId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVisibility

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetTasks

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetTasks() []int64`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetTasksOk() (*[]int64, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetTasks(v []int64)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetOptionTypes() []ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetOptionTypesOk() (*[]ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetOptionTypes(v []ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTaskSetTasks

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetTaskSetTasks() []ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner`

GetTaskSetTasks returns the TaskSetTasks field if non-nil, zero value otherwise.

### GetTaskSetTasksOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) GetTaskSetTasksOk() (*[]ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner, bool)`

GetTaskSetTasksOk returns a tuple with the TaskSetTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetTasks

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) SetTaskSetTasks(v []ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner)`

SetTaskSetTasks sets TaskSetTasks field to given value.

### HasTaskSetTasks

`func (o *ListWorkflows200ResponseAllOfTaskSetsInner) HasTaskSetTasks() bool`

HasTaskSetTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


