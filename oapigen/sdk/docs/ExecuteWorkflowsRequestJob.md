# ExecuteWorkflowsRequestJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the execution job. Can be used to find execution results with &#x60;/api/processes?name&#x3D;&#x60; | [optional] 
**TargetType** | Pointer to **string** | The target context for task execution. This is required for tasks with &#x60;executeTarget&#x60; set to &#x60;resource&#x60;. | [optional] 
**Instances** | Pointer to **[]int64** | Array of Instance IDs. Only applicable if &#x60;targetType&#x60; is instance. | [optional] 
**Servers** | Pointer to **[]int64** | Array of Server IDs. Only applicable if &#x60;targetType&#x60; is &#x60;server&#x60;. | [optional] 
**InstanceLabel** | Pointer to **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | Pointer to **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the task. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | String of custom configuration values as JSON. | [optional] 

## Methods

### NewExecuteWorkflowsRequestJob

`func NewExecuteWorkflowsRequestJob() *ExecuteWorkflowsRequestJob`

NewExecuteWorkflowsRequestJob instantiates a new ExecuteWorkflowsRequestJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteWorkflowsRequestJobWithDefaults

`func NewExecuteWorkflowsRequestJobWithDefaults() *ExecuteWorkflowsRequestJob`

NewExecuteWorkflowsRequestJobWithDefaults instantiates a new ExecuteWorkflowsRequestJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExecuteWorkflowsRequestJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecuteWorkflowsRequestJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecuteWorkflowsRequestJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecuteWorkflowsRequestJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetType

`func (o *ExecuteWorkflowsRequestJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ExecuteWorkflowsRequestJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ExecuteWorkflowsRequestJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ExecuteWorkflowsRequestJob) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetInstances

`func (o *ExecuteWorkflowsRequestJob) GetInstances() []int64`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ExecuteWorkflowsRequestJob) GetInstancesOk() (*[]int64, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ExecuteWorkflowsRequestJob) SetInstances(v []int64)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ExecuteWorkflowsRequestJob) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetServers

`func (o *ExecuteWorkflowsRequestJob) GetServers() []int64`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ExecuteWorkflowsRequestJob) GetServersOk() (*[]int64, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ExecuteWorkflowsRequestJob) SetServers(v []int64)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ExecuteWorkflowsRequestJob) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *ExecuteWorkflowsRequestJob) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ExecuteWorkflowsRequestJob) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ExecuteWorkflowsRequestJob) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ExecuteWorkflowsRequestJob) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetServerLabel

`func (o *ExecuteWorkflowsRequestJob) GetServerLabel() string`

GetServerLabel returns the ServerLabel field if non-nil, zero value otherwise.

### GetServerLabelOk

`func (o *ExecuteWorkflowsRequestJob) GetServerLabelOk() (*string, bool)`

GetServerLabelOk returns a tuple with the ServerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLabel

`func (o *ExecuteWorkflowsRequestJob) SetServerLabel(v string)`

SetServerLabel sets ServerLabel field to given value.

### HasServerLabel

`func (o *ExecuteWorkflowsRequestJob) HasServerLabel() bool`

HasServerLabel returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ExecuteWorkflowsRequestJob) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ExecuteWorkflowsRequestJob) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ExecuteWorkflowsRequestJob) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ExecuteWorkflowsRequestJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ExecuteWorkflowsRequestJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ExecuteWorkflowsRequestJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ExecuteWorkflowsRequestJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ExecuteWorkflowsRequestJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


