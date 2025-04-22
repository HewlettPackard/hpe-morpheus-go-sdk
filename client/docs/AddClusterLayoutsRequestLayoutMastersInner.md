# AddClusterLayoutsRequestLayoutMastersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of nodes | [optional] [default to 1]
**ContainerType** | [**AddClusterLayoutsRequestLayoutMastersInnerContainerType**](AddClusterLayoutsRequestLayoutMastersInnerContainerType.md) |  | 

## Methods

### NewAddClusterLayoutsRequestLayoutMastersInner

`func NewAddClusterLayoutsRequestLayoutMastersInner(containerType AddClusterLayoutsRequestLayoutMastersInnerContainerType, ) *AddClusterLayoutsRequestLayoutMastersInner`

NewAddClusterLayoutsRequestLayoutMastersInner instantiates a new AddClusterLayoutsRequestLayoutMastersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterLayoutsRequestLayoutMastersInnerWithDefaults

`func NewAddClusterLayoutsRequestLayoutMastersInnerWithDefaults() *AddClusterLayoutsRequestLayoutMastersInner`

NewAddClusterLayoutsRequestLayoutMastersInnerWithDefaults instantiates a new AddClusterLayoutsRequestLayoutMastersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *AddClusterLayoutsRequestLayoutMastersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterLayoutsRequestLayoutMastersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterLayoutsRequestLayoutMastersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterLayoutsRequestLayoutMastersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetContainerType

`func (o *AddClusterLayoutsRequestLayoutMastersInner) GetContainerType() AddClusterLayoutsRequestLayoutMastersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *AddClusterLayoutsRequestLayoutMastersInner) GetContainerTypeOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *AddClusterLayoutsRequestLayoutMastersInner) SetContainerType(v AddClusterLayoutsRequestLayoutMastersInnerContainerType)`

SetContainerType sets ContainerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


