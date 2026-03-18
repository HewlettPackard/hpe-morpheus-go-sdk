# UpdateClusterNamespace200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to [**UpdateClusterNamespace200ResponseAllOfNamespace**](UpdateClusterNamespace200ResponseAllOfNamespace.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateClusterNamespace200Response

`func NewUpdateClusterNamespace200Response() *UpdateClusterNamespace200Response`

NewUpdateClusterNamespace200Response instantiates a new UpdateClusterNamespace200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterNamespace200ResponseWithDefaults

`func NewUpdateClusterNamespace200ResponseWithDefaults() *UpdateClusterNamespace200Response`

NewUpdateClusterNamespace200ResponseWithDefaults instantiates a new UpdateClusterNamespace200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *UpdateClusterNamespace200Response) GetNamespace() UpdateClusterNamespace200ResponseAllOfNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateClusterNamespace200Response) GetNamespaceOk() (*UpdateClusterNamespace200ResponseAllOfNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateClusterNamespace200Response) SetNamespace(v UpdateClusterNamespace200ResponseAllOfNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateClusterNamespace200Response) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateClusterNamespace200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateClusterNamespace200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateClusterNamespace200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateClusterNamespace200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


