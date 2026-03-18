# MetaObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** | Offset records, the number of records to skip, can be used to paginate over results. | [optional] [default to 0]
**Max** | Pointer to **int64** | Max size, the maximum number of records to include in the response. | [optional] [default to 25]
**Size** | Pointer to **int64** | Number of records returned in the response | [optional] [default to 0]
**Total** | Pointer to **int64** | Total number of records found | [optional] [default to 0]

## Methods

### NewMetaObject

`func NewMetaObject() *MetaObject`

NewMetaObject instantiates a new MetaObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaObjectWithDefaults

`func NewMetaObjectWithDefaults() *MetaObject`

NewMetaObjectWithDefaults instantiates a new MetaObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *MetaObject) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MetaObject) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MetaObject) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *MetaObject) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMax

`func (o *MetaObject) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *MetaObject) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *MetaObject) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *MetaObject) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetSize

`func (o *MetaObject) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MetaObject) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MetaObject) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MetaObject) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *MetaObject) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetaObject) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetaObject) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MetaObject) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


