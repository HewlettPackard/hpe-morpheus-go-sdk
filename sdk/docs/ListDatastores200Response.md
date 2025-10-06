# ListDatastores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to [**[]ListDatastores200ResponseAllOfDatastoresInner**](ListDatastores200ResponseAllOfDatastoresInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListDatastores200Response

`func NewListDatastores200Response() *ListDatastores200Response`

NewListDatastores200Response instantiates a new ListDatastores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatastores200ResponseWithDefaults

`func NewListDatastores200ResponseWithDefaults() *ListDatastores200Response`

NewListDatastores200ResponseWithDefaults instantiates a new ListDatastores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *ListDatastores200Response) GetDatastores() []ListDatastores200ResponseAllOfDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ListDatastores200Response) GetDatastoresOk() (*[]ListDatastores200ResponseAllOfDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ListDatastores200Response) SetDatastores(v []ListDatastores200ResponseAllOfDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ListDatastores200Response) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetMeta

`func (o *ListDatastores200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDatastores200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDatastores200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDatastores200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


