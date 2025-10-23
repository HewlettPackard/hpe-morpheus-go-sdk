# ListLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sort** | Pointer to [**ListLogs200ResponseAllOfSort**](ListLogs200ResponseAllOfSort.md) |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**Data** | Pointer to [**[]ListLogs200ResponseAllOfDataInner**](ListLogs200ResponseAllOfDataInner.md) |  | [optional] 
**Max** | Pointer to **int64** |  | [optional] 
**GrandTotal** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLogs200Response

`func NewListLogs200Response() *ListLogs200Response`

NewListLogs200Response instantiates a new ListLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseWithDefaults

`func NewListLogs200ResponseWithDefaults() *ListLogs200Response`

NewListLogs200ResponseWithDefaults instantiates a new ListLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSort

`func (o *ListLogs200Response) GetSort() ListLogs200ResponseAllOfSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListLogs200Response) GetSortOk() (*ListLogs200ResponseAllOfSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListLogs200Response) SetSort(v ListLogs200ResponseAllOfSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListLogs200Response) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetOffset

`func (o *ListLogs200Response) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListLogs200Response) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListLogs200Response) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListLogs200Response) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStart

`func (o *ListLogs200Response) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListLogs200Response) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListLogs200Response) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListLogs200Response) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ListLogs200Response) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListLogs200Response) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListLogs200Response) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListLogs200Response) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetData

`func (o *ListLogs200Response) GetData() []ListLogs200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLogs200Response) GetDataOk() (*[]ListLogs200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLogs200Response) SetData(v []ListLogs200ResponseAllOfDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListLogs200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMax

`func (o *ListLogs200Response) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ListLogs200Response) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ListLogs200Response) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ListLogs200Response) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetGrandTotal

`func (o *ListLogs200Response) GetGrandTotal() int64`

GetGrandTotal returns the GrandTotal field if non-nil, zero value otherwise.

### GetGrandTotalOk

`func (o *ListLogs200Response) GetGrandTotalOk() (*int64, bool)`

GetGrandTotalOk returns a tuple with the GrandTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandTotal

`func (o *ListLogs200Response) SetGrandTotal(v int64)`

SetGrandTotal sets GrandTotal field to given value.

### HasGrandTotal

`func (o *ListLogs200Response) HasGrandTotal() bool`

HasGrandTotal returns a boolean if a field has been set.

### GetTotal

`func (o *ListLogs200Response) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListLogs200Response) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListLogs200Response) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListLogs200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *ListLogs200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListLogs200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListLogs200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListLogs200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCount

`func (o *ListLogs200Response) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListLogs200Response) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListLogs200Response) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListLogs200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMeta

`func (o *ListLogs200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLogs200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLogs200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLogs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


