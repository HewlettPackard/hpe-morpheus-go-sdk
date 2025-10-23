# ListVdi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desktops** | Pointer to [**[]ListVdi200ResponseAllOfDesktopsInner**](ListVdi200ResponseAllOfDesktopsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVdi200Response

`func NewListVdi200Response() *ListVdi200Response`

NewListVdi200Response instantiates a new ListVdi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVdi200ResponseWithDefaults

`func NewListVdi200ResponseWithDefaults() *ListVdi200Response`

NewListVdi200ResponseWithDefaults instantiates a new ListVdi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesktops

`func (o *ListVdi200Response) GetDesktops() []ListVdi200ResponseAllOfDesktopsInner`

GetDesktops returns the Desktops field if non-nil, zero value otherwise.

### GetDesktopsOk

`func (o *ListVdi200Response) GetDesktopsOk() (*[]ListVdi200ResponseAllOfDesktopsInner, bool)`

GetDesktopsOk returns a tuple with the Desktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktops

`func (o *ListVdi200Response) SetDesktops(v []ListVdi200ResponseAllOfDesktopsInner)`

SetDesktops sets Desktops field to given value.

### HasDesktops

`func (o *ListVdi200Response) HasDesktops() bool`

HasDesktops returns a boolean if a field has been set.

### GetMeta

`func (o *ListVdi200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVdi200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVdi200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVdi200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


