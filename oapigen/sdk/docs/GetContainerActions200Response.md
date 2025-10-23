# GetContainerActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerIds** | Pointer to **[]int64** |  | [optional] 
**Actions** | Pointer to [**[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner**](ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner.md) |  | [optional] 

## Methods

### NewGetContainerActions200Response

`func NewGetContainerActions200Response() *GetContainerActions200Response`

NewGetContainerActions200Response instantiates a new GetContainerActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerActions200ResponseWithDefaults

`func NewGetContainerActions200ResponseWithDefaults() *GetContainerActions200Response`

NewGetContainerActions200ResponseWithDefaults instantiates a new GetContainerActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerIds

`func (o *GetContainerActions200Response) GetContainerIds() []int64`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *GetContainerActions200Response) GetContainerIdsOk() (*[]int64, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *GetContainerActions200Response) SetContainerIds(v []int64)`

SetContainerIds sets ContainerIds field to given value.

### HasContainerIds

`func (o *GetContainerActions200Response) HasContainerIds() bool`

HasContainerIds returns a boolean if a field has been set.

### GetActions

`func (o *GetContainerActions200Response) GetActions() []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetContainerActions200Response) GetActionsOk() (*[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetContainerActions200Response) SetActions(v []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetContainerActions200Response) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


