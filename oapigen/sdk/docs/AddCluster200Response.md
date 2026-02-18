# AddCluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**AddCluster200ResponseAllOfCluster**](AddCluster200ResponseAllOfCluster.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCluster200Response

`func NewAddCluster200Response() *AddCluster200Response`

NewAddCluster200Response instantiates a new AddCluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCluster200ResponseWithDefaults

`func NewAddCluster200ResponseWithDefaults() *AddCluster200Response`

NewAddCluster200ResponseWithDefaults instantiates a new AddCluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *AddCluster200Response) GetCluster() AddCluster200ResponseAllOfCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AddCluster200Response) GetClusterOk() (*AddCluster200ResponseAllOfCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AddCluster200Response) SetCluster(v AddCluster200ResponseAllOfCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AddCluster200Response) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCluster200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCluster200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCluster200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCluster200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


