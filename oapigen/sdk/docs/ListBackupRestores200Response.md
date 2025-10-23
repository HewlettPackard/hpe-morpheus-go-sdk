# ListBackupRestores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restores** | Pointer to [**[]ListBackupRestores200ResponseAllOfRestoresInner**](ListBackupRestores200ResponseAllOfRestoresInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBackupRestores200Response

`func NewListBackupRestores200Response() *ListBackupRestores200Response`

NewListBackupRestores200Response instantiates a new ListBackupRestores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupRestores200ResponseWithDefaults

`func NewListBackupRestores200ResponseWithDefaults() *ListBackupRestores200Response`

NewListBackupRestores200ResponseWithDefaults instantiates a new ListBackupRestores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestores

`func (o *ListBackupRestores200Response) GetRestores() []ListBackupRestores200ResponseAllOfRestoresInner`

GetRestores returns the Restores field if non-nil, zero value otherwise.

### GetRestoresOk

`func (o *ListBackupRestores200Response) GetRestoresOk() (*[]ListBackupRestores200ResponseAllOfRestoresInner, bool)`

GetRestoresOk returns a tuple with the Restores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestores

`func (o *ListBackupRestores200Response) SetRestores(v []ListBackupRestores200ResponseAllOfRestoresInner)`

SetRestores sets Restores field to given value.

### HasRestores

`func (o *ListBackupRestores200Response) HasRestores() bool`

HasRestores returns a boolean if a field has been set.

### GetMeta

`func (o *ListBackupRestores200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBackupRestores200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBackupRestores200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBackupRestores200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


