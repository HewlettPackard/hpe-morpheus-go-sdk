# ListOptionLists200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionTypes** | Pointer to [**[]ListOptionLists200ResponseAllOfOptionTypesInner**](ListOptionLists200ResponseAllOfOptionTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListOptionLists200Response

`func NewListOptionLists200Response() *ListOptionLists200Response`

NewListOptionLists200Response instantiates a new ListOptionLists200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionLists200ResponseWithDefaults

`func NewListOptionLists200ResponseWithDefaults() *ListOptionLists200Response`

NewListOptionLists200ResponseWithDefaults instantiates a new ListOptionLists200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionTypes

`func (o *ListOptionLists200Response) GetOptionTypes() []ListOptionLists200ResponseAllOfOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListOptionLists200Response) GetOptionTypesOk() (*[]ListOptionLists200ResponseAllOfOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListOptionLists200Response) SetOptionTypes(v []ListOptionLists200ResponseAllOfOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListOptionLists200Response) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListOptionLists200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListOptionLists200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListOptionLists200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListOptionLists200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


