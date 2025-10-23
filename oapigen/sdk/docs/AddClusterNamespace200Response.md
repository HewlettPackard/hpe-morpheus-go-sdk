# AddClusterNamespace200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to [**AddClusterNamespace200ResponseAllOfNamespace**](AddClusterNamespace200ResponseAllOfNamespace.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddClusterNamespace200Response

`func NewAddClusterNamespace200Response() *AddClusterNamespace200Response`

NewAddClusterNamespace200Response instantiates a new AddClusterNamespace200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterNamespace200ResponseWithDefaults

`func NewAddClusterNamespace200ResponseWithDefaults() *AddClusterNamespace200Response`

NewAddClusterNamespace200ResponseWithDefaults instantiates a new AddClusterNamespace200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *AddClusterNamespace200Response) GetNamespace() AddClusterNamespace200ResponseAllOfNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AddClusterNamespace200Response) GetNamespaceOk() (*AddClusterNamespace200ResponseAllOfNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AddClusterNamespace200Response) SetNamespace(v AddClusterNamespace200ResponseAllOfNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AddClusterNamespace200Response) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSuccess

`func (o *AddClusterNamespace200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddClusterNamespace200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddClusterNamespace200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddClusterNamespace200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


