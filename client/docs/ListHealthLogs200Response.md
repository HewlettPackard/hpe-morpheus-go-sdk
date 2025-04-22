# ListHealthLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]ListHealthLogs200ResponseAllOfLogsInner**](ListHealthLogs200ResponseAllOfLogsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListHealthLogs200Response

`func NewListHealthLogs200Response() *ListHealthLogs200Response`

NewListHealthLogs200Response instantiates a new ListHealthLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealthLogs200ResponseWithDefaults

`func NewListHealthLogs200ResponseWithDefaults() *ListHealthLogs200Response`

NewListHealthLogs200ResponseWithDefaults instantiates a new ListHealthLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ListHealthLogs200Response) GetLogs() []ListHealthLogs200ResponseAllOfLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ListHealthLogs200Response) GetLogsOk() (*[]ListHealthLogs200ResponseAllOfLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ListHealthLogs200Response) SetLogs(v []ListHealthLogs200ResponseAllOfLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ListHealthLogs200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMeta

`func (o *ListHealthLogs200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListHealthLogs200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListHealthLogs200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListHealthLogs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


