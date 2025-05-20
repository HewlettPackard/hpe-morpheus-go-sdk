# GetAppState200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | Pointer to [**[]GetAppState200ResponseAllOfWorkloadsInner**](GetAppState200ResponseAllOfWorkloadsInner.md) |  | [optional] 
**IacDrift** | Pointer to **bool** |  | [optional] 
**PlanResources** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Specs** | Pointer to [**[]GetAppState200ResponseAllOfSpecsInner**](GetAppState200ResponseAllOfSpecsInner.md) |  | [optional] 
**StateData** | Pointer to **string** |  | [optional] 
**PlanData** | Pointer to **string** |  | [optional] 
**Input** | Pointer to [**GetAppState200ResponseAllOfInput**](GetAppState200ResponseAllOfInput.md) |  | [optional] 
**Output** | Pointer to [**GetAppState200ResponseAllOfOutput**](GetAppState200ResponseAllOfOutput.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAppState200Response

`func NewGetAppState200Response() *GetAppState200Response`

NewGetAppState200Response instantiates a new GetAppState200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppState200ResponseWithDefaults

`func NewGetAppState200ResponseWithDefaults() *GetAppState200Response`

NewGetAppState200ResponseWithDefaults instantiates a new GetAppState200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *GetAppState200Response) GetWorkloads() []GetAppState200ResponseAllOfWorkloadsInner`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *GetAppState200Response) GetWorkloadsOk() (*[]GetAppState200ResponseAllOfWorkloadsInner, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *GetAppState200Response) SetWorkloads(v []GetAppState200ResponseAllOfWorkloadsInner)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *GetAppState200Response) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetIacDrift

`func (o *GetAppState200Response) GetIacDrift() bool`

GetIacDrift returns the IacDrift field if non-nil, zero value otherwise.

### GetIacDriftOk

`func (o *GetAppState200Response) GetIacDriftOk() (*bool, bool)`

GetIacDriftOk returns a tuple with the IacDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacDrift

`func (o *GetAppState200Response) SetIacDrift(v bool)`

SetIacDrift sets IacDrift field to given value.

### HasIacDrift

`func (o *GetAppState200Response) HasIacDrift() bool`

HasIacDrift returns a boolean if a field has been set.

### GetPlanResources

`func (o *GetAppState200Response) GetPlanResources() []map[string]interface{}`

GetPlanResources returns the PlanResources field if non-nil, zero value otherwise.

### GetPlanResourcesOk

`func (o *GetAppState200Response) GetPlanResourcesOk() (*[]map[string]interface{}, bool)`

GetPlanResourcesOk returns a tuple with the PlanResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResources

`func (o *GetAppState200Response) SetPlanResources(v []map[string]interface{})`

SetPlanResources sets PlanResources field to given value.

### HasPlanResources

`func (o *GetAppState200Response) HasPlanResources() bool`

HasPlanResources returns a boolean if a field has been set.

### GetSpecs

`func (o *GetAppState200Response) GetSpecs() []GetAppState200ResponseAllOfSpecsInner`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *GetAppState200Response) GetSpecsOk() (*[]GetAppState200ResponseAllOfSpecsInner, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *GetAppState200Response) SetSpecs(v []GetAppState200ResponseAllOfSpecsInner)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *GetAppState200Response) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetStateData

`func (o *GetAppState200Response) GetStateData() string`

GetStateData returns the StateData field if non-nil, zero value otherwise.

### GetStateDataOk

`func (o *GetAppState200Response) GetStateDataOk() (*string, bool)`

GetStateDataOk returns a tuple with the StateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateData

`func (o *GetAppState200Response) SetStateData(v string)`

SetStateData sets StateData field to given value.

### HasStateData

`func (o *GetAppState200Response) HasStateData() bool`

HasStateData returns a boolean if a field has been set.

### GetPlanData

`func (o *GetAppState200Response) GetPlanData() string`

GetPlanData returns the PlanData field if non-nil, zero value otherwise.

### GetPlanDataOk

`func (o *GetAppState200Response) GetPlanDataOk() (*string, bool)`

GetPlanDataOk returns a tuple with the PlanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanData

`func (o *GetAppState200Response) SetPlanData(v string)`

SetPlanData sets PlanData field to given value.

### HasPlanData

`func (o *GetAppState200Response) HasPlanData() bool`

HasPlanData returns a boolean if a field has been set.

### GetInput

`func (o *GetAppState200Response) GetInput() GetAppState200ResponseAllOfInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GetAppState200Response) GetInputOk() (*GetAppState200ResponseAllOfInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GetAppState200Response) SetInput(v GetAppState200ResponseAllOfInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *GetAppState200Response) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *GetAppState200Response) GetOutput() GetAppState200ResponseAllOfOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetAppState200Response) GetOutputOk() (*GetAppState200ResponseAllOfOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetAppState200Response) SetOutput(v GetAppState200ResponseAllOfOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetAppState200Response) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetSuccess

`func (o *GetAppState200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetAppState200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetAppState200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetAppState200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


