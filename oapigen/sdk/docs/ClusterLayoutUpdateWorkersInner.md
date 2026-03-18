# ClusterLayoutUpdateWorkersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of nodes | [optional] [default to 1]
**ContainerType** | [**ClusterLayoutUpdateWorkersInnerContainerType**](ClusterLayoutUpdateWorkersInnerContainerType.md) |  | 

## Methods

### NewClusterLayoutUpdateWorkersInner

`func NewClusterLayoutUpdateWorkersInner(containerType ClusterLayoutUpdateWorkersInnerContainerType, ) *ClusterLayoutUpdateWorkersInner`

NewClusterLayoutUpdateWorkersInner instantiates a new ClusterLayoutUpdateWorkersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutUpdateWorkersInnerWithDefaults

`func NewClusterLayoutUpdateWorkersInnerWithDefaults() *ClusterLayoutUpdateWorkersInner`

NewClusterLayoutUpdateWorkersInnerWithDefaults instantiates a new ClusterLayoutUpdateWorkersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterLayoutUpdateWorkersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterLayoutUpdateWorkersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterLayoutUpdateWorkersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterLayoutUpdateWorkersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetContainerType

`func (o *ClusterLayoutUpdateWorkersInner) GetContainerType() ClusterLayoutUpdateWorkersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ClusterLayoutUpdateWorkersInner) GetContainerTypeOk() (*ClusterLayoutUpdateWorkersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ClusterLayoutUpdateWorkersInner) SetContainerType(v ClusterLayoutUpdateWorkersInnerContainerType)`

SetContainerType sets ContainerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


