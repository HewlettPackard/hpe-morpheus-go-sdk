# ListActivity200ResponseAllOfMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** | Offset records, the number of records to skip, can be used to paginate over results. | [optional] [default to 0]
**Max** | Pointer to **int64** | Max size, the maximum number of records to include in the response. | [optional] [default to 25]
**Size** | Pointer to **int64** | Number of records returned in the response | [optional] [default to 0]
**Total** | Pointer to **int64** | Total number of records found | [optional] [default to 0]

## Methods

### NewListActivity200ResponseAllOfMeta

`func NewListActivity200ResponseAllOfMeta() *ListActivity200ResponseAllOfMeta`

NewListActivity200ResponseAllOfMeta instantiates a new ListActivity200ResponseAllOfMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActivity200ResponseAllOfMetaWithDefaults

`func NewListActivity200ResponseAllOfMetaWithDefaults() *ListActivity200ResponseAllOfMeta`

NewListActivity200ResponseAllOfMetaWithDefaults instantiates a new ListActivity200ResponseAllOfMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *ListActivity200ResponseAllOfMeta) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListActivity200ResponseAllOfMeta) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListActivity200ResponseAllOfMeta) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListActivity200ResponseAllOfMeta) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMax

`func (o *ListActivity200ResponseAllOfMeta) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ListActivity200ResponseAllOfMeta) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ListActivity200ResponseAllOfMeta) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ListActivity200ResponseAllOfMeta) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetSize

`func (o *ListActivity200ResponseAllOfMeta) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListActivity200ResponseAllOfMeta) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListActivity200ResponseAllOfMeta) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListActivity200ResponseAllOfMeta) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *ListActivity200ResponseAllOfMeta) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListActivity200ResponseAllOfMeta) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListActivity200ResponseAllOfMeta) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListActivity200ResponseAllOfMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


