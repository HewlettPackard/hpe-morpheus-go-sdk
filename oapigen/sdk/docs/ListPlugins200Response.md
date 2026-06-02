# ListPlugins200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]ListPlugins200ResponseAllOfPluginsInner**](ListPlugins200ResponseAllOfPluginsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPlugins200Response

`func NewListPlugins200Response() *ListPlugins200Response`

NewListPlugins200Response instantiates a new ListPlugins200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetPlugins

`func (o *ListPlugins200Response) GetPlugins() []ListPlugins200ResponseAllOfPluginsInner`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ListPlugins200Response) GetPluginsOk() (*[]ListPlugins200ResponseAllOfPluginsInner, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ListPlugins200Response) SetPlugins(v []ListPlugins200ResponseAllOfPluginsInner)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ListPlugins200Response) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetMeta

`func (o *ListPlugins200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPlugins200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPlugins200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPlugins200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


