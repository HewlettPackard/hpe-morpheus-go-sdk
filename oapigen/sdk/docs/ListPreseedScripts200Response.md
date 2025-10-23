# ListPreseedScripts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreseedScripts** | Pointer to [**[]ListPreseedScripts200ResponseAllOfPreseedScriptsInner**](ListPreseedScripts200ResponseAllOfPreseedScriptsInner.md) |  | [optional] 
**PreseedScriptCount** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPreseedScripts200Response

`func NewListPreseedScripts200Response() *ListPreseedScripts200Response`

NewListPreseedScripts200Response instantiates a new ListPreseedScripts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPreseedScripts200ResponseWithDefaults

`func NewListPreseedScripts200ResponseWithDefaults() *ListPreseedScripts200Response`

NewListPreseedScripts200ResponseWithDefaults instantiates a new ListPreseedScripts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreseedScripts

`func (o *ListPreseedScripts200Response) GetPreseedScripts() []ListPreseedScripts200ResponseAllOfPreseedScriptsInner`

GetPreseedScripts returns the PreseedScripts field if non-nil, zero value otherwise.

### GetPreseedScriptsOk

`func (o *ListPreseedScripts200Response) GetPreseedScriptsOk() (*[]ListPreseedScripts200ResponseAllOfPreseedScriptsInner, bool)`

GetPreseedScriptsOk returns a tuple with the PreseedScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScripts

`func (o *ListPreseedScripts200Response) SetPreseedScripts(v []ListPreseedScripts200ResponseAllOfPreseedScriptsInner)`

SetPreseedScripts sets PreseedScripts field to given value.

### HasPreseedScripts

`func (o *ListPreseedScripts200Response) HasPreseedScripts() bool`

HasPreseedScripts returns a boolean if a field has been set.

### GetPreseedScriptCount

`func (o *ListPreseedScripts200Response) GetPreseedScriptCount() int64`

GetPreseedScriptCount returns the PreseedScriptCount field if non-nil, zero value otherwise.

### GetPreseedScriptCountOk

`func (o *ListPreseedScripts200Response) GetPreseedScriptCountOk() (*int64, bool)`

GetPreseedScriptCountOk returns a tuple with the PreseedScriptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScriptCount

`func (o *ListPreseedScripts200Response) SetPreseedScriptCount(v int64)`

SetPreseedScriptCount sets PreseedScriptCount field to given value.

### HasPreseedScriptCount

`func (o *ListPreseedScripts200Response) HasPreseedScriptCount() bool`

HasPreseedScriptCount returns a boolean if a field has been set.

### GetMeta

`func (o *ListPreseedScripts200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPreseedScripts200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPreseedScripts200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPreseedScripts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


