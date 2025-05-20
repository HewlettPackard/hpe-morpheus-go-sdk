# InstanceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IacDrift** | Pointer to **bool** |  | [optional] 
**PlanResources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Specs** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StateData** | Pointer to **string** |  | [optional] 
**PlanData** | Pointer to **string** |  | [optional] 
**Input** | Pointer to [**GetStateInstance200ResponseAllOfInstanceInput**](GetStateInstance200ResponseAllOfInstanceInput.md) |  | [optional] 
**Output** | Pointer to [**GetAppState200ResponseAllOfOutput**](GetAppState200ResponseAllOfOutput.md) |  | [optional] 

## Methods

### NewInstanceState

`func NewInstanceState() *InstanceState`

NewInstanceState instantiates a new InstanceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceStateWithDefaults

`func NewInstanceStateWithDefaults() *InstanceState`

NewInstanceStateWithDefaults instantiates a new InstanceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *InstanceState) GetWorkloads() []map[string]interface{}`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *InstanceState) GetWorkloadsOk() (*[]map[string]interface{}, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *InstanceState) SetWorkloads(v []map[string]interface{})`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *InstanceState) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetIacDrift

`func (o *InstanceState) GetIacDrift() bool`

GetIacDrift returns the IacDrift field if non-nil, zero value otherwise.

### GetIacDriftOk

`func (o *InstanceState) GetIacDriftOk() (*bool, bool)`

GetIacDriftOk returns a tuple with the IacDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacDrift

`func (o *InstanceState) SetIacDrift(v bool)`

SetIacDrift sets IacDrift field to given value.

### HasIacDrift

`func (o *InstanceState) HasIacDrift() bool`

HasIacDrift returns a boolean if a field has been set.

### GetPlanResources

`func (o *InstanceState) GetPlanResources() []map[string]interface{}`

GetPlanResources returns the PlanResources field if non-nil, zero value otherwise.

### GetPlanResourcesOk

`func (o *InstanceState) GetPlanResourcesOk() (*[]map[string]interface{}, bool)`

GetPlanResourcesOk returns a tuple with the PlanResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResources

`func (o *InstanceState) SetPlanResources(v []map[string]interface{})`

SetPlanResources sets PlanResources field to given value.

### HasPlanResources

`func (o *InstanceState) HasPlanResources() bool`

HasPlanResources returns a boolean if a field has been set.

### GetSpecs

`func (o *InstanceState) GetSpecs() []map[string]interface{}`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *InstanceState) GetSpecsOk() (*[]map[string]interface{}, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *InstanceState) SetSpecs(v []map[string]interface{})`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *InstanceState) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetStateData

`func (o *InstanceState) GetStateData() string`

GetStateData returns the StateData field if non-nil, zero value otherwise.

### GetStateDataOk

`func (o *InstanceState) GetStateDataOk() (*string, bool)`

GetStateDataOk returns a tuple with the StateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateData

`func (o *InstanceState) SetStateData(v string)`

SetStateData sets StateData field to given value.

### HasStateData

`func (o *InstanceState) HasStateData() bool`

HasStateData returns a boolean if a field has been set.

### GetPlanData

`func (o *InstanceState) GetPlanData() string`

GetPlanData returns the PlanData field if non-nil, zero value otherwise.

### GetPlanDataOk

`func (o *InstanceState) GetPlanDataOk() (*string, bool)`

GetPlanDataOk returns a tuple with the PlanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanData

`func (o *InstanceState) SetPlanData(v string)`

SetPlanData sets PlanData field to given value.

### HasPlanData

`func (o *InstanceState) HasPlanData() bool`

HasPlanData returns a boolean if a field has been set.

### GetInput

`func (o *InstanceState) GetInput() GetStateInstance200ResponseAllOfInstanceInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InstanceState) GetInputOk() (*GetStateInstance200ResponseAllOfInstanceInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InstanceState) SetInput(v GetStateInstance200ResponseAllOfInstanceInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *InstanceState) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *InstanceState) GetOutput() GetAppState200ResponseAllOfOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *InstanceState) GetOutputOk() (*GetAppState200ResponseAllOfOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *InstanceState) SetOutput(v GetAppState200ResponseAllOfOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *InstanceState) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


