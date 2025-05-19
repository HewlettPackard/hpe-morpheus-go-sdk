# AddWorkflowsRequestTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the workflow | 
**Description** | Pointer to **string** | A description of the workflow | [optional] 
**Labels** | Pointer to **[]string** | An array of Category labels for filtering | [optional] 
**Type** | Pointer to **string** | Workflow type | [optional] [default to "provision"]
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**OptionTypes** | Pointer to **[]int64** | List of option type IDs for use with operational workflow configuration. | [optional] 
**Tasks** | Pointer to [**AddWorkflowsRequestTaskSetTasks**](AddWorkflowsRequestTaskSetTasks.md) |  | [optional] 

## Methods

### NewAddWorkflowsRequestTaskSet

`func NewAddWorkflowsRequestTaskSet(name string, ) *AddWorkflowsRequestTaskSet`

NewAddWorkflowsRequestTaskSet instantiates a new AddWorkflowsRequestTaskSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWorkflowsRequestTaskSetWithDefaults

`func NewAddWorkflowsRequestTaskSetWithDefaults() *AddWorkflowsRequestTaskSet`

NewAddWorkflowsRequestTaskSetWithDefaults instantiates a new AddWorkflowsRequestTaskSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddWorkflowsRequestTaskSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddWorkflowsRequestTaskSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddWorkflowsRequestTaskSet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddWorkflowsRequestTaskSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddWorkflowsRequestTaskSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddWorkflowsRequestTaskSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddWorkflowsRequestTaskSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddWorkflowsRequestTaskSet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddWorkflowsRequestTaskSet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddWorkflowsRequestTaskSet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddWorkflowsRequestTaskSet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AddWorkflowsRequestTaskSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddWorkflowsRequestTaskSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddWorkflowsRequestTaskSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddWorkflowsRequestTaskSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *AddWorkflowsRequestTaskSet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddWorkflowsRequestTaskSet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddWorkflowsRequestTaskSet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddWorkflowsRequestTaskSet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetOptionTypes

`func (o *AddWorkflowsRequestTaskSet) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddWorkflowsRequestTaskSet) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddWorkflowsRequestTaskSet) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddWorkflowsRequestTaskSet) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTasks

`func (o *AddWorkflowsRequestTaskSet) GetTasks() AddWorkflowsRequestTaskSetTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AddWorkflowsRequestTaskSet) GetTasksOk() (*AddWorkflowsRequestTaskSetTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AddWorkflowsRequestTaskSet) SetTasks(v AddWorkflowsRequestTaskSetTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AddWorkflowsRequestTaskSet) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


