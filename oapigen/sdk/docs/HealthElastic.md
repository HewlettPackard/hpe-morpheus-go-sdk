# HealthElastic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Master** | Pointer to [**HealthElasticMaster**](HealthElasticMaster.md) |  | [optional] 
**Nodes** | Pointer to [**[]HealthElasticNodesInner**](HealthElasticNodesInner.md) |  | [optional] 
**Stats** | Pointer to [**HealthElasticStats**](HealthElasticStats.md) |  | [optional] 
**Indices** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BadIndices** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewHealthElastic

`func NewHealthElastic() *HealthElastic`

NewHealthElastic instantiates a new HealthElastic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthElasticWithDefaults

`func NewHealthElasticWithDefaults() *HealthElastic`

NewHealthElasticWithDefaults instantiates a new HealthElastic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *HealthElastic) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HealthElastic) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HealthElastic) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HealthElastic) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetStatus

`func (o *HealthElastic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthElastic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthElastic) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthElastic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaster

`func (o *HealthElastic) GetMaster() HealthElasticMaster`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *HealthElastic) GetMasterOk() (*HealthElasticMaster, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *HealthElastic) SetMaster(v HealthElasticMaster)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *HealthElastic) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetNodes

`func (o *HealthElastic) GetNodes() []HealthElasticNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HealthElastic) GetNodesOk() (*[]HealthElasticNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HealthElastic) SetNodes(v []HealthElasticNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HealthElastic) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetStats

`func (o *HealthElastic) GetStats() HealthElasticStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *HealthElastic) GetStatsOk() (*HealthElasticStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *HealthElastic) SetStats(v HealthElasticStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *HealthElastic) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIndices

`func (o *HealthElastic) GetIndices() []map[string]interface{}`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *HealthElastic) GetIndicesOk() (*[]map[string]interface{}, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *HealthElastic) SetIndices(v []map[string]interface{})`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *HealthElastic) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### SetIndicesNil

`func (o *HealthElastic) SetIndicesNil(b bool)`

 SetIndicesNil sets the value for Indices to be an explicit nil

### UnsetIndices
`func (o *HealthElastic) UnsetIndices()`

UnsetIndices ensures that no value is present for Indices, not even an explicit nil
### GetBadIndices

`func (o *HealthElastic) GetBadIndices() []map[string]interface{}`

GetBadIndices returns the BadIndices field if non-nil, zero value otherwise.

### GetBadIndicesOk

`func (o *HealthElastic) GetBadIndicesOk() (*[]map[string]interface{}, bool)`

GetBadIndicesOk returns a tuple with the BadIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadIndices

`func (o *HealthElastic) SetBadIndices(v []map[string]interface{})`

SetBadIndices sets BadIndices field to given value.

### HasBadIndices

`func (o *HealthElastic) HasBadIndices() bool`

HasBadIndices returns a boolean if a field has been set.

### SetBadIndicesNil

`func (o *HealthElastic) SetBadIndicesNil(b bool)`

 SetBadIndicesNil sets the value for BadIndices to be an explicit nil

### UnsetBadIndices
`func (o *HealthElastic) UnsetBadIndices()`

UnsetBadIndices ensures that no value is present for BadIndices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


