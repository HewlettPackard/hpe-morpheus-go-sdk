# ListVirtualImages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualImages** | Pointer to [**[]ListVirtualImages200ResponseAllOfVirtualImagesInner**](ListVirtualImages200ResponseAllOfVirtualImagesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVirtualImages200Response

`func NewListVirtualImages200Response() *ListVirtualImages200Response`

NewListVirtualImages200Response instantiates a new ListVirtualImages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVirtualImages200ResponseWithDefaults

`func NewListVirtualImages200ResponseWithDefaults() *ListVirtualImages200Response`

NewListVirtualImages200ResponseWithDefaults instantiates a new ListVirtualImages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualImages

`func (o *ListVirtualImages200Response) GetVirtualImages() []ListVirtualImages200ResponseAllOfVirtualImagesInner`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ListVirtualImages200Response) GetVirtualImagesOk() (*[]ListVirtualImages200ResponseAllOfVirtualImagesInner, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ListVirtualImages200Response) SetVirtualImages(v []ListVirtualImages200ResponseAllOfVirtualImagesInner)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ListVirtualImages200Response) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetMeta

`func (o *ListVirtualImages200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVirtualImages200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVirtualImages200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVirtualImages200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


