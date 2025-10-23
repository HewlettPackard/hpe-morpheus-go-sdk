# ListBackupResults200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ListBackupResults200ResponseAllOfResultsInner**](ListBackupResults200ResponseAllOfResultsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBackupResults200Response

`func NewListBackupResults200Response() *ListBackupResults200Response`

NewListBackupResults200Response instantiates a new ListBackupResults200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupResults200ResponseWithDefaults

`func NewListBackupResults200ResponseWithDefaults() *ListBackupResults200Response`

NewListBackupResults200ResponseWithDefaults instantiates a new ListBackupResults200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListBackupResults200Response) GetResults() []ListBackupResults200ResponseAllOfResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListBackupResults200Response) GetResultsOk() (*[]ListBackupResults200ResponseAllOfResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListBackupResults200Response) SetResults(v []ListBackupResults200ResponseAllOfResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListBackupResults200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetMeta

`func (o *ListBackupResults200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBackupResults200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBackupResults200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBackupResults200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


