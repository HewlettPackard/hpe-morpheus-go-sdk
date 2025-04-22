# Workflow

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

### NewWorkflow

`func NewWorkflow() *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workflow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Workflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Workflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Workflow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Workflow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Workflow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Workflow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Workflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Workflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Workflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Workflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *Workflow) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Workflow) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Workflow) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Workflow) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDateCreated

`func (o *Workflow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Workflow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Workflow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Workflow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Workflow) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Workflow) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Workflow) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Workflow) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccountId

`func (o *Workflow) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Workflow) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Workflow) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Workflow) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *Workflow) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Workflow) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Workflow) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Workflow) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVisibility

`func (o *Workflow) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Workflow) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Workflow) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Workflow) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *Workflow) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *Workflow) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *Workflow) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *Workflow) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetTasks

`func (o *Workflow) GetTasks() []int64`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Workflow) GetTasksOk() (*[]int64, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Workflow) SetTasks(v []int64)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Workflow) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetOptionTypes

`func (o *Workflow) GetOptionTypes() []ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *Workflow) GetOptionTypesOk() (*[]ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *Workflow) SetOptionTypes(v []ListWorkflows200ResponseAllOfTaskSetsInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *Workflow) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTaskSetTasks

`func (o *Workflow) GetTaskSetTasks() []ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner`

GetTaskSetTasks returns the TaskSetTasks field if non-nil, zero value otherwise.

### GetTaskSetTasksOk

`func (o *Workflow) GetTaskSetTasksOk() (*[]ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner, bool)`

GetTaskSetTasksOk returns a tuple with the TaskSetTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetTasks

`func (o *Workflow) SetTaskSetTasks(v []ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInner)`

SetTaskSetTasks sets TaskSetTasks field to given value.

### HasTaskSetTasks

`func (o *Workflow) HasTaskSetTasks() bool`

HasTaskSetTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


