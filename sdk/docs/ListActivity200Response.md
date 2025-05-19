# ListActivity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**[]ListActivity200ResponseAllOfActivityInner**](ListActivity200ResponseAllOfActivityInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListActivity200Response

`func NewListActivity200Response() *ListActivity200Response`

NewListActivity200Response instantiates a new ListActivity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActivity200ResponseWithDefaults

`func NewListActivity200ResponseWithDefaults() *ListActivity200Response`

NewListActivity200ResponseWithDefaults instantiates a new ListActivity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ListActivity200Response) GetActivity() []ListActivity200ResponseAllOfActivityInner`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ListActivity200Response) GetActivityOk() (*[]ListActivity200ResponseAllOfActivityInner, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ListActivity200Response) SetActivity(v []ListActivity200ResponseAllOfActivityInner)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ListActivity200Response) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMeta

`func (o *ListActivity200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListActivity200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListActivity200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListActivity200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


