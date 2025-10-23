# Log

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

## Methods

### NewLog

`func NewLog() *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSort

`func (o *Log) GetSort() ListLogs200ResponseAllOfSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Log) GetSortOk() (*ListLogs200ResponseAllOfSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Log) SetSort(v ListLogs200ResponseAllOfSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Log) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetOffset

`func (o *Log) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Log) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Log) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Log) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStart

`func (o *Log) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Log) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Log) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *Log) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Log) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Log) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Log) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Log) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetData

`func (o *Log) GetData() []ListLogs200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Log) GetDataOk() (*[]ListLogs200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Log) SetData(v []ListLogs200ResponseAllOfDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *Log) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMax

`func (o *Log) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Log) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Log) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *Log) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetGrandTotal

`func (o *Log) GetGrandTotal() int64`

GetGrandTotal returns the GrandTotal field if non-nil, zero value otherwise.

### GetGrandTotalOk

`func (o *Log) GetGrandTotalOk() (*int64, bool)`

GetGrandTotalOk returns a tuple with the GrandTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandTotal

`func (o *Log) SetGrandTotal(v int64)`

SetGrandTotal sets GrandTotal field to given value.

### HasGrandTotal

`func (o *Log) HasGrandTotal() bool`

HasGrandTotal returns a boolean if a field has been set.

### GetTotal

`func (o *Log) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Log) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Log) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Log) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *Log) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Log) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Log) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Log) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCount

`func (o *Log) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Log) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Log) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Log) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


