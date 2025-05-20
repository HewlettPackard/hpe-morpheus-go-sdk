# GetClusterNamespaces200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | Pointer to [**[]GetClusterNamespaces200ResponseAllOfNamespacesInner**](GetClusterNamespaces200ResponseAllOfNamespacesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetClusterNamespaces200Response

`func NewGetClusterNamespaces200Response() *GetClusterNamespaces200Response`

NewGetClusterNamespaces200Response instantiates a new GetClusterNamespaces200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterNamespaces200ResponseWithDefaults

`func NewGetClusterNamespaces200ResponseWithDefaults() *GetClusterNamespaces200Response`

NewGetClusterNamespaces200ResponseWithDefaults instantiates a new GetClusterNamespaces200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaces

`func (o *GetClusterNamespaces200Response) GetNamespaces() []GetClusterNamespaces200ResponseAllOfNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *GetClusterNamespaces200Response) GetNamespacesOk() (*[]GetClusterNamespaces200ResponseAllOfNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *GetClusterNamespaces200Response) SetNamespaces(v []GetClusterNamespaces200ResponseAllOfNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *GetClusterNamespaces200Response) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetMeta

`func (o *GetClusterNamespaces200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetClusterNamespaces200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetClusterNamespaces200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetClusterNamespaces200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


