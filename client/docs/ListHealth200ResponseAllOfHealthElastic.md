# ListHealth200ResponseAllOfHealthElastic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Master** | Pointer to [**ListHealth200ResponseAllOfHealthElasticMaster**](ListHealth200ResponseAllOfHealthElasticMaster.md) |  | [optional] 
**Nodes** | Pointer to [**[]ListHealth200ResponseAllOfHealthElasticNodesInner**](ListHealth200ResponseAllOfHealthElasticNodesInner.md) |  | [optional] 
**Stats** | Pointer to [**ListHealth200ResponseAllOfHealthElasticStats**](ListHealth200ResponseAllOfHealthElasticStats.md) |  | [optional] 
**Indices** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BadIndices** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListHealth200ResponseAllOfHealthElastic

`func NewListHealth200ResponseAllOfHealthElastic() *ListHealth200ResponseAllOfHealthElastic`

NewListHealth200ResponseAllOfHealthElastic instantiates a new ListHealth200ResponseAllOfHealthElastic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseAllOfHealthElasticWithDefaults

`func NewListHealth200ResponseAllOfHealthElasticWithDefaults() *ListHealth200ResponseAllOfHealthElastic`

NewListHealth200ResponseAllOfHealthElasticWithDefaults instantiates a new ListHealth200ResponseAllOfHealthElastic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListHealth200ResponseAllOfHealthElastic) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListHealth200ResponseAllOfHealthElastic) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListHealth200ResponseAllOfHealthElastic) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetStatus

`func (o *ListHealth200ResponseAllOfHealthElastic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHealth200ResponseAllOfHealthElastic) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHealth200ResponseAllOfHealthElastic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMaster

`func (o *ListHealth200ResponseAllOfHealthElastic) GetMaster() ListHealth200ResponseAllOfHealthElasticMaster`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetMasterOk() (*ListHealth200ResponseAllOfHealthElasticMaster, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ListHealth200ResponseAllOfHealthElastic) SetMaster(v ListHealth200ResponseAllOfHealthElasticMaster)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ListHealth200ResponseAllOfHealthElastic) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetNodes

`func (o *ListHealth200ResponseAllOfHealthElastic) GetNodes() []ListHealth200ResponseAllOfHealthElasticNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetNodesOk() (*[]ListHealth200ResponseAllOfHealthElasticNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ListHealth200ResponseAllOfHealthElastic) SetNodes(v []ListHealth200ResponseAllOfHealthElasticNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ListHealth200ResponseAllOfHealthElastic) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetStats

`func (o *ListHealth200ResponseAllOfHealthElastic) GetStats() ListHealth200ResponseAllOfHealthElasticStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetStatsOk() (*ListHealth200ResponseAllOfHealthElasticStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListHealth200ResponseAllOfHealthElastic) SetStats(v ListHealth200ResponseAllOfHealthElasticStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListHealth200ResponseAllOfHealthElastic) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIndices

`func (o *ListHealth200ResponseAllOfHealthElastic) GetIndices() []map[string]interface{}`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetIndicesOk() (*[]map[string]interface{}, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *ListHealth200ResponseAllOfHealthElastic) SetIndices(v []map[string]interface{})`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *ListHealth200ResponseAllOfHealthElastic) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetBadIndices

`func (o *ListHealth200ResponseAllOfHealthElastic) GetBadIndices() []map[string]interface{}`

GetBadIndices returns the BadIndices field if non-nil, zero value otherwise.

### GetBadIndicesOk

`func (o *ListHealth200ResponseAllOfHealthElastic) GetBadIndicesOk() (*[]map[string]interface{}, bool)`

GetBadIndicesOk returns a tuple with the BadIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadIndices

`func (o *ListHealth200ResponseAllOfHealthElastic) SetBadIndices(v []map[string]interface{})`

SetBadIndices sets BadIndices field to given value.

### HasBadIndices

`func (o *ListHealth200ResponseAllOfHealthElastic) HasBadIndices() bool`

HasBadIndices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


