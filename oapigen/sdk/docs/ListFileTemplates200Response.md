# ListFileTemplates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerTemplates** | Pointer to [**[]ListFileTemplates200ResponseAllOfContainerTemplatesInner**](ListFileTemplates200ResponseAllOfContainerTemplatesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListFileTemplates200Response

`func NewListFileTemplates200Response() *ListFileTemplates200Response`

NewListFileTemplates200Response instantiates a new ListFileTemplates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFileTemplates200ResponseWithDefaults

`func NewListFileTemplates200ResponseWithDefaults() *ListFileTemplates200Response`

NewListFileTemplates200ResponseWithDefaults instantiates a new ListFileTemplates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerTemplates

`func (o *ListFileTemplates200Response) GetContainerTemplates() []ListFileTemplates200ResponseAllOfContainerTemplatesInner`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *ListFileTemplates200Response) GetContainerTemplatesOk() (*[]ListFileTemplates200ResponseAllOfContainerTemplatesInner, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *ListFileTemplates200Response) SetContainerTemplates(v []ListFileTemplates200ResponseAllOfContainerTemplatesInner)`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *ListFileTemplates200Response) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### GetMeta

`func (o *ListFileTemplates200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFileTemplates200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFileTemplates200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFileTemplates200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


