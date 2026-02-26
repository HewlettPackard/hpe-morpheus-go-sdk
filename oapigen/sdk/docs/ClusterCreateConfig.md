# ClusterCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DynamicPlacementMode** | Pointer to **string** | Dynamic Placement | [optional] 
**TemplateParameter** | Pointer to [**ClusterCreateConfigTemplateParameter**](ClusterCreateConfigTemplateParameter.md) |  | [optional] 

## Methods

### NewClusterCreateConfig

`func NewClusterCreateConfig() *ClusterCreateConfig`

NewClusterCreateConfig instantiates a new ClusterCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateConfigWithDefaults

`func NewClusterCreateConfigWithDefaults() *ClusterCreateConfig`

NewClusterCreateConfigWithDefaults instantiates a new ClusterCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamicPlacementMode

`func (o *ClusterCreateConfig) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *ClusterCreateConfig) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *ClusterCreateConfig) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *ClusterCreateConfig) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetTemplateParameter

`func (o *ClusterCreateConfig) GetTemplateParameter() ClusterCreateConfigTemplateParameter`

GetTemplateParameter returns the TemplateParameter field if non-nil, zero value otherwise.

### GetTemplateParameterOk

`func (o *ClusterCreateConfig) GetTemplateParameterOk() (*ClusterCreateConfigTemplateParameter, bool)`

GetTemplateParameterOk returns a tuple with the TemplateParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameter

`func (o *ClusterCreateConfig) SetTemplateParameter(v ClusterCreateConfigTemplateParameter)`

SetTemplateParameter sets TemplateParameter field to given value.

### HasTemplateParameter

`func (o *ClusterCreateConfig) HasTemplateParameter() bool`

HasTemplateParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


