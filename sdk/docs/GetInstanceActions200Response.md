# GetInstanceActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceIds** | Pointer to **[]int64** |  | [optional] 
**Actions** | Pointer to [**[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner**](ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner.md) |  | [optional] 

## Methods

### NewGetInstanceActions200Response

`func NewGetInstanceActions200Response() *GetInstanceActions200Response`

NewGetInstanceActions200Response instantiates a new GetInstanceActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceActions200ResponseWithDefaults

`func NewGetInstanceActions200ResponseWithDefaults() *GetInstanceActions200Response`

NewGetInstanceActions200ResponseWithDefaults instantiates a new GetInstanceActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceIds

`func (o *GetInstanceActions200Response) GetInstanceIds() []int64`

GetInstanceIds returns the InstanceIds field if non-nil, zero value otherwise.

### GetInstanceIdsOk

`func (o *GetInstanceActions200Response) GetInstanceIdsOk() (*[]int64, bool)`

GetInstanceIdsOk returns a tuple with the InstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIds

`func (o *GetInstanceActions200Response) SetInstanceIds(v []int64)`

SetInstanceIds sets InstanceIds field to given value.

### HasInstanceIds

`func (o *GetInstanceActions200Response) HasInstanceIds() bool`

HasInstanceIds returns a boolean if a field has been set.

### GetActions

`func (o *GetInstanceActions200Response) GetActions() []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetInstanceActions200Response) GetActionsOk() (*[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetInstanceActions200Response) SetActions(v []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetInstanceActions200Response) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


