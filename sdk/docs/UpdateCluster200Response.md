# UpdateCluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**GetCluster200ResponseCluster**](GetCluster200ResponseCluster.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCluster200Response

`func NewUpdateCluster200Response() *UpdateCluster200Response`

NewUpdateCluster200Response instantiates a new UpdateCluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCluster200ResponseWithDefaults

`func NewUpdateCluster200ResponseWithDefaults() *UpdateCluster200Response`

NewUpdateCluster200ResponseWithDefaults instantiates a new UpdateCluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *UpdateCluster200Response) GetCluster() GetCluster200ResponseCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateCluster200Response) GetClusterOk() (*GetCluster200ResponseCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateCluster200Response) SetCluster(v GetCluster200ResponseCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *UpdateCluster200Response) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCluster200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCluster200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCluster200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCluster200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


