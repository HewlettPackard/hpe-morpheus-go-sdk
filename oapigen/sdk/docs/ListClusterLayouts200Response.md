# ListClusterLayouts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layouts** | Pointer to [**[]ListClusterLayouts200ResponseAllOfLayoutsInner**](ListClusterLayouts200ResponseAllOfLayoutsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterLayouts200Response

`func NewListClusterLayouts200Response() *ListClusterLayouts200Response`

NewListClusterLayouts200Response instantiates a new ListClusterLayouts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetLayouts

`func (o *ListClusterLayouts200Response) GetLayouts() []ListClusterLayouts200ResponseAllOfLayoutsInner`

GetLayouts returns the Layouts field if non-nil, zero value otherwise.

### GetLayoutsOk

`func (o *ListClusterLayouts200Response) GetLayoutsOk() (*[]ListClusterLayouts200ResponseAllOfLayoutsInner, bool)`

GetLayoutsOk returns a tuple with the Layouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayouts

`func (o *ListClusterLayouts200Response) SetLayouts(v []ListClusterLayouts200ResponseAllOfLayoutsInner)`

SetLayouts sets Layouts field to given value.

### HasLayouts

`func (o *ListClusterLayouts200Response) HasLayouts() bool`

HasLayouts returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterLayouts200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterLayouts200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterLayouts200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterLayouts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


