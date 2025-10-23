# ListBootScripts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootScripts** | Pointer to [**[]ListBootScripts200ResponseAllOfBootScriptsInner**](ListBootScripts200ResponseAllOfBootScriptsInner.md) |  | [optional] 
**BootScriptCount** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBootScripts200Response

`func NewListBootScripts200Response() *ListBootScripts200Response`

NewListBootScripts200Response instantiates a new ListBootScripts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBootScripts200ResponseWithDefaults

`func NewListBootScripts200ResponseWithDefaults() *ListBootScripts200Response`

NewListBootScripts200ResponseWithDefaults instantiates a new ListBootScripts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootScripts

`func (o *ListBootScripts200Response) GetBootScripts() []ListBootScripts200ResponseAllOfBootScriptsInner`

GetBootScripts returns the BootScripts field if non-nil, zero value otherwise.

### GetBootScriptsOk

`func (o *ListBootScripts200Response) GetBootScriptsOk() (*[]ListBootScripts200ResponseAllOfBootScriptsInner, bool)`

GetBootScriptsOk returns a tuple with the BootScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScripts

`func (o *ListBootScripts200Response) SetBootScripts(v []ListBootScripts200ResponseAllOfBootScriptsInner)`

SetBootScripts sets BootScripts field to given value.

### HasBootScripts

`func (o *ListBootScripts200Response) HasBootScripts() bool`

HasBootScripts returns a boolean if a field has been set.

### GetBootScriptCount

`func (o *ListBootScripts200Response) GetBootScriptCount() int64`

GetBootScriptCount returns the BootScriptCount field if non-nil, zero value otherwise.

### GetBootScriptCountOk

`func (o *ListBootScripts200Response) GetBootScriptCountOk() (*int64, bool)`

GetBootScriptCountOk returns a tuple with the BootScriptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScriptCount

`func (o *ListBootScripts200Response) SetBootScriptCount(v int64)`

SetBootScriptCount sets BootScriptCount field to given value.

### HasBootScriptCount

`func (o *ListBootScripts200Response) HasBootScriptCount() bool`

HasBootScriptCount returns a boolean if a field has been set.

### GetMeta

`func (o *ListBootScripts200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBootScripts200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBootScripts200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBootScripts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


