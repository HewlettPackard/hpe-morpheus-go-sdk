# GetStateInstance200ResponseAllOfInstance

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

### NewGetStateInstance200ResponseAllOfInstance

`func NewGetStateInstance200ResponseAllOfInstance() *GetStateInstance200ResponseAllOfInstance`

NewGetStateInstance200ResponseAllOfInstance instantiates a new GetStateInstance200ResponseAllOfInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStateInstance200ResponseAllOfInstanceWithDefaults

`func NewGetStateInstance200ResponseAllOfInstanceWithDefaults() *GetStateInstance200ResponseAllOfInstance`

NewGetStateInstance200ResponseAllOfInstanceWithDefaults instantiates a new GetStateInstance200ResponseAllOfInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *GetStateInstance200ResponseAllOfInstance) GetWorkloads() []map[string]interface{}`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetWorkloadsOk() (*[]map[string]interface{}, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *GetStateInstance200ResponseAllOfInstance) SetWorkloads(v []map[string]interface{})`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *GetStateInstance200ResponseAllOfInstance) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetIacDrift

`func (o *GetStateInstance200ResponseAllOfInstance) GetIacDrift() bool`

GetIacDrift returns the IacDrift field if non-nil, zero value otherwise.

### GetIacDriftOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetIacDriftOk() (*bool, bool)`

GetIacDriftOk returns a tuple with the IacDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacDrift

`func (o *GetStateInstance200ResponseAllOfInstance) SetIacDrift(v bool)`

SetIacDrift sets IacDrift field to given value.

### HasIacDrift

`func (o *GetStateInstance200ResponseAllOfInstance) HasIacDrift() bool`

HasIacDrift returns a boolean if a field has been set.

### GetPlanResources

`func (o *GetStateInstance200ResponseAllOfInstance) GetPlanResources() []map[string]interface{}`

GetPlanResources returns the PlanResources field if non-nil, zero value otherwise.

### GetPlanResourcesOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetPlanResourcesOk() (*[]map[string]interface{}, bool)`

GetPlanResourcesOk returns a tuple with the PlanResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResources

`func (o *GetStateInstance200ResponseAllOfInstance) SetPlanResources(v []map[string]interface{})`

SetPlanResources sets PlanResources field to given value.

### HasPlanResources

`func (o *GetStateInstance200ResponseAllOfInstance) HasPlanResources() bool`

HasPlanResources returns a boolean if a field has been set.

### GetSpecs

`func (o *GetStateInstance200ResponseAllOfInstance) GetSpecs() []map[string]interface{}`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetSpecsOk() (*[]map[string]interface{}, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *GetStateInstance200ResponseAllOfInstance) SetSpecs(v []map[string]interface{})`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *GetStateInstance200ResponseAllOfInstance) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetStateData

`func (o *GetStateInstance200ResponseAllOfInstance) GetStateData() string`

GetStateData returns the StateData field if non-nil, zero value otherwise.

### GetStateDataOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetStateDataOk() (*string, bool)`

GetStateDataOk returns a tuple with the StateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateData

`func (o *GetStateInstance200ResponseAllOfInstance) SetStateData(v string)`

SetStateData sets StateData field to given value.

### HasStateData

`func (o *GetStateInstance200ResponseAllOfInstance) HasStateData() bool`

HasStateData returns a boolean if a field has been set.

### GetPlanData

`func (o *GetStateInstance200ResponseAllOfInstance) GetPlanData() string`

GetPlanData returns the PlanData field if non-nil, zero value otherwise.

### GetPlanDataOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetPlanDataOk() (*string, bool)`

GetPlanDataOk returns a tuple with the PlanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanData

`func (o *GetStateInstance200ResponseAllOfInstance) SetPlanData(v string)`

SetPlanData sets PlanData field to given value.

### HasPlanData

`func (o *GetStateInstance200ResponseAllOfInstance) HasPlanData() bool`

HasPlanData returns a boolean if a field has been set.

### GetInput

`func (o *GetStateInstance200ResponseAllOfInstance) GetInput() GetStateInstance200ResponseAllOfInstanceInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetInputOk() (*GetStateInstance200ResponseAllOfInstanceInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GetStateInstance200ResponseAllOfInstance) SetInput(v GetStateInstance200ResponseAllOfInstanceInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *GetStateInstance200ResponseAllOfInstance) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *GetStateInstance200ResponseAllOfInstance) GetOutput() GetAppState200ResponseAllOfOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetStateInstance200ResponseAllOfInstance) GetOutputOk() (*GetAppState200ResponseAllOfOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetStateInstance200ResponseAllOfInstance) SetOutput(v GetAppState200ResponseAllOfOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetStateInstance200ResponseAllOfInstance) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


