# UpdateClusterAssignedLayout200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**UpdateCluster200ResponseAllOfCluster**](UpdateCluster200ResponseAllOfCluster.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateClusterAssignedLayout200Response

`func NewUpdateClusterAssignedLayout200Response() *UpdateClusterAssignedLayout200Response`

NewUpdateClusterAssignedLayout200Response instantiates a new UpdateClusterAssignedLayout200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterAssignedLayout200ResponseWithDefaults

`func NewUpdateClusterAssignedLayout200ResponseWithDefaults() *UpdateClusterAssignedLayout200Response`

NewUpdateClusterAssignedLayout200ResponseWithDefaults instantiates a new UpdateClusterAssignedLayout200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *UpdateClusterAssignedLayout200Response) GetCluster() UpdateCluster200ResponseAllOfCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateClusterAssignedLayout200Response) GetClusterOk() (*UpdateCluster200ResponseAllOfCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateClusterAssignedLayout200Response) SetCluster(v UpdateCluster200ResponseAllOfCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *UpdateClusterAssignedLayout200Response) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateClusterAssignedLayout200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateClusterAssignedLayout200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateClusterAssignedLayout200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateClusterAssignedLayout200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


