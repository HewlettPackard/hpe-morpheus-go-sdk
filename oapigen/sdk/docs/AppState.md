# AppState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | Pointer to [**[]AppStateWorkloadsInner**](AppStateWorkloadsInner.md) |  | [optional] 
**IacDrift** | Pointer to **bool** |  | [optional] 
**PlanResources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Specs** | Pointer to [**[]AppStateSpecsInner**](AppStateSpecsInner.md) |  | [optional] 
**StateData** | Pointer to **string** |  | [optional] 
**PlanData** | Pointer to **string** |  | [optional] 
**Input** | Pointer to [**AppStateInput**](AppStateInput.md) |  | [optional] 
**Output** | Pointer to [**AppStateOutput**](AppStateOutput.md) |  | [optional] 

## Methods

### NewAppState

`func NewAppState() *AppState`

NewAppState instantiates a new AppState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStateWithDefaults

`func NewAppStateWithDefaults() *AppState`

NewAppStateWithDefaults instantiates a new AppState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *AppState) GetWorkloads() []AppStateWorkloadsInner`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *AppState) GetWorkloadsOk() (*[]AppStateWorkloadsInner, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *AppState) SetWorkloads(v []AppStateWorkloadsInner)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *AppState) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetIacDrift

`func (o *AppState) GetIacDrift() bool`

GetIacDrift returns the IacDrift field if non-nil, zero value otherwise.

### GetIacDriftOk

`func (o *AppState) GetIacDriftOk() (*bool, bool)`

GetIacDriftOk returns a tuple with the IacDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacDrift

`func (o *AppState) SetIacDrift(v bool)`

SetIacDrift sets IacDrift field to given value.

### HasIacDrift

`func (o *AppState) HasIacDrift() bool`

HasIacDrift returns a boolean if a field has been set.

### GetPlanResources

`func (o *AppState) GetPlanResources() []map[string]interface{}`

GetPlanResources returns the PlanResources field if non-nil, zero value otherwise.

### GetPlanResourcesOk

`func (o *AppState) GetPlanResourcesOk() (*[]map[string]interface{}, bool)`

GetPlanResourcesOk returns a tuple with the PlanResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResources

`func (o *AppState) SetPlanResources(v []map[string]interface{})`

SetPlanResources sets PlanResources field to given value.

### HasPlanResources

`func (o *AppState) HasPlanResources() bool`

HasPlanResources returns a boolean if a field has been set.

### GetSpecs

`func (o *AppState) GetSpecs() []AppStateSpecsInner`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *AppState) GetSpecsOk() (*[]AppStateSpecsInner, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *AppState) SetSpecs(v []AppStateSpecsInner)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *AppState) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetStateData

`func (o *AppState) GetStateData() string`

GetStateData returns the StateData field if non-nil, zero value otherwise.

### GetStateDataOk

`func (o *AppState) GetStateDataOk() (*string, bool)`

GetStateDataOk returns a tuple with the StateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateData

`func (o *AppState) SetStateData(v string)`

SetStateData sets StateData field to given value.

### HasStateData

`func (o *AppState) HasStateData() bool`

HasStateData returns a boolean if a field has been set.

### GetPlanData

`func (o *AppState) GetPlanData() string`

GetPlanData returns the PlanData field if non-nil, zero value otherwise.

### GetPlanDataOk

`func (o *AppState) GetPlanDataOk() (*string, bool)`

GetPlanDataOk returns a tuple with the PlanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanData

`func (o *AppState) SetPlanData(v string)`

SetPlanData sets PlanData field to given value.

### HasPlanData

`func (o *AppState) HasPlanData() bool`

HasPlanData returns a boolean if a field has been set.

### GetInput

`func (o *AppState) GetInput() AppStateInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AppState) GetInputOk() (*AppStateInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AppState) SetInput(v AppStateInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *AppState) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *AppState) GetOutput() AppStateOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AppState) GetOutputOk() (*AppStateOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AppState) SetOutput(v AppStateOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AppState) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


