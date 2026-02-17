# ClusterLayoutCreateWorkersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of nodes | [optional] [default to 1]
**ContainerType** | [**ClusterLayoutCreateWorkersInnerContainerType**](ClusterLayoutCreateWorkersInnerContainerType.md) |  | 

## Methods

### NewClusterLayoutCreateWorkersInner

`func NewClusterLayoutCreateWorkersInner(containerType ClusterLayoutCreateWorkersInnerContainerType, ) *ClusterLayoutCreateWorkersInner`

NewClusterLayoutCreateWorkersInner instantiates a new ClusterLayoutCreateWorkersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutCreateWorkersInnerWithDefaults

`func NewClusterLayoutCreateWorkersInnerWithDefaults() *ClusterLayoutCreateWorkersInner`

NewClusterLayoutCreateWorkersInnerWithDefaults instantiates a new ClusterLayoutCreateWorkersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterLayoutCreateWorkersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterLayoutCreateWorkersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterLayoutCreateWorkersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterLayoutCreateWorkersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetContainerType

`func (o *ClusterLayoutCreateWorkersInner) GetContainerType() ClusterLayoutCreateWorkersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ClusterLayoutCreateWorkersInner) GetContainerTypeOk() (*ClusterLayoutCreateWorkersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ClusterLayoutCreateWorkersInner) SetContainerType(v ClusterLayoutCreateWorkersInnerContainerType)`

SetContainerType sets ContainerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


