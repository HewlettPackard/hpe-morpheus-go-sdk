# AppStateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**[]AppStateInputVariablesInner**](AppStateInputVariablesInner.md) |  | [optional] 
**Providers** | Pointer to [**[]AppStateInputProvidersInner**](AppStateInputProvidersInner.md) |  | [optional] 
**Data** | Pointer to [**[]AppStateInputDataInner**](AppStateInputDataInner.md) |  | [optional] 

## Methods

### NewAppStateInput

`func NewAppStateInput() *AppStateInput`

NewAppStateInput instantiates a new AppStateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStateInputWithDefaults

`func NewAppStateInputWithDefaults() *AppStateInput`

NewAppStateInputWithDefaults instantiates a new AppStateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *AppStateInput) GetVariables() []AppStateInputVariablesInner`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *AppStateInput) GetVariablesOk() (*[]AppStateInputVariablesInner, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *AppStateInput) SetVariables(v []AppStateInputVariablesInner)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *AppStateInput) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetProviders

`func (o *AppStateInput) GetProviders() []AppStateInputProvidersInner`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *AppStateInput) GetProvidersOk() (*[]AppStateInputProvidersInner, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *AppStateInput) SetProviders(v []AppStateInputProvidersInner)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *AppStateInput) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetData

`func (o *AppStateInput) GetData() []AppStateInputDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStateInput) GetDataOk() (*[]AppStateInputDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStateInput) SetData(v []AppStateInputDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AppStateInput) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


