# ClusterLayoutUpdateMastersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of nodes | [optional] [default to 1]
**ContainerType** | [**ClusterLayoutUpdateMastersInnerContainerType**](ClusterLayoutUpdateMastersInnerContainerType.md) |  | 

## Methods

### NewClusterLayoutUpdateMastersInner

`func NewClusterLayoutUpdateMastersInner(containerType ClusterLayoutUpdateMastersInnerContainerType, ) *ClusterLayoutUpdateMastersInner`

NewClusterLayoutUpdateMastersInner instantiates a new ClusterLayoutUpdateMastersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutUpdateMastersInnerWithDefaults

`func NewClusterLayoutUpdateMastersInnerWithDefaults() *ClusterLayoutUpdateMastersInner`

NewClusterLayoutUpdateMastersInnerWithDefaults instantiates a new ClusterLayoutUpdateMastersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterLayoutUpdateMastersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterLayoutUpdateMastersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterLayoutUpdateMastersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterLayoutUpdateMastersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetContainerType

`func (o *ClusterLayoutUpdateMastersInner) GetContainerType() ClusterLayoutUpdateMastersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ClusterLayoutUpdateMastersInner) GetContainerTypeOk() (*ClusterLayoutUpdateMastersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ClusterLayoutUpdateMastersInner) SetContainerType(v ClusterLayoutUpdateMastersInnerContainerType)`

SetContainerType sets ContainerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


