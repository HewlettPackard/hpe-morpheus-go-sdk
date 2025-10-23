# ListSpecTemplates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecTemplates** | Pointer to [**[]ListSpecTemplates200ResponseAllOfSpecTemplatesInner**](ListSpecTemplates200ResponseAllOfSpecTemplatesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSpecTemplates200Response

`func NewListSpecTemplates200Response() *ListSpecTemplates200Response`

NewListSpecTemplates200Response instantiates a new ListSpecTemplates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpecTemplates200ResponseWithDefaults

`func NewListSpecTemplates200ResponseWithDefaults() *ListSpecTemplates200Response`

NewListSpecTemplates200ResponseWithDefaults instantiates a new ListSpecTemplates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecTemplates

`func (o *ListSpecTemplates200Response) GetSpecTemplates() []ListSpecTemplates200ResponseAllOfSpecTemplatesInner`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *ListSpecTemplates200Response) GetSpecTemplatesOk() (*[]ListSpecTemplates200ResponseAllOfSpecTemplatesInner, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *ListSpecTemplates200Response) SetSpecTemplates(v []ListSpecTemplates200ResponseAllOfSpecTemplatesInner)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *ListSpecTemplates200Response) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### GetMeta

`func (o *ListSpecTemplates200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSpecTemplates200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSpecTemplates200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSpecTemplates200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


