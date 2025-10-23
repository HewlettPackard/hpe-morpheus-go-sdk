# ListScripts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerScripts** | Pointer to [**[]ListScripts200ResponseAllOfContainerScriptsInner**](ListScripts200ResponseAllOfContainerScriptsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListScripts200Response

`func NewListScripts200Response() *ListScripts200Response`

NewListScripts200Response instantiates a new ListScripts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListScripts200ResponseWithDefaults

`func NewListScripts200ResponseWithDefaults() *ListScripts200Response`

NewListScripts200ResponseWithDefaults instantiates a new ListScripts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerScripts

`func (o *ListScripts200Response) GetContainerScripts() []ListScripts200ResponseAllOfContainerScriptsInner`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *ListScripts200Response) GetContainerScriptsOk() (*[]ListScripts200ResponseAllOfContainerScriptsInner, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *ListScripts200Response) SetContainerScripts(v []ListScripts200ResponseAllOfContainerScriptsInner)`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *ListScripts200Response) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### GetMeta

`func (o *ListScripts200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListScripts200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListScripts200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListScripts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


