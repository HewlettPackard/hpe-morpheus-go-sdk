# ListBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | Pointer to [**[]ListBackups200ResponseAllOfBackupsInner**](ListBackups200ResponseAllOfBackupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBackups200Response

`func NewListBackups200Response() *ListBackups200Response`

NewListBackups200Response instantiates a new ListBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackups200ResponseWithDefaults

`func NewListBackups200ResponseWithDefaults() *ListBackups200Response`

NewListBackups200ResponseWithDefaults instantiates a new ListBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackups

`func (o *ListBackups200Response) GetBackups() []ListBackups200ResponseAllOfBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *ListBackups200Response) GetBackupsOk() (*[]ListBackups200ResponseAllOfBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *ListBackups200Response) SetBackups(v []ListBackups200ResponseAllOfBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *ListBackups200Response) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetMeta

`func (o *ListBackups200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBackups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBackups200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBackups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


