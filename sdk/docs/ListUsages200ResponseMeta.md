# ListUsages200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** | Offset records, the number of records to skip, can be used to paginate over results. | [optional] [default to 0]
**Max** | Pointer to **int64** | Max size, the maximum number of records to include in the response. | [optional] [default to 25]
**Size** | Pointer to **int64** | Number of records returned in the response | [optional] [default to 0]
**Total** | Pointer to **int64** | Total number of records found | [optional] [default to 0]
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListUsages200ResponseMeta

`func NewListUsages200ResponseMeta() *ListUsages200ResponseMeta`

NewListUsages200ResponseMeta instantiates a new ListUsages200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsages200ResponseMetaWithDefaults

`func NewListUsages200ResponseMetaWithDefaults() *ListUsages200ResponseMeta`

NewListUsages200ResponseMetaWithDefaults instantiates a new ListUsages200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *ListUsages200ResponseMeta) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListUsages200ResponseMeta) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListUsages200ResponseMeta) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListUsages200ResponseMeta) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMax

`func (o *ListUsages200ResponseMeta) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ListUsages200ResponseMeta) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ListUsages200ResponseMeta) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ListUsages200ResponseMeta) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetSize

`func (o *ListUsages200ResponseMeta) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUsages200ResponseMeta) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUsages200ResponseMeta) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListUsages200ResponseMeta) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *ListUsages200ResponseMeta) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListUsages200ResponseMeta) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListUsages200ResponseMeta) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListUsages200ResponseMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetStartDate

`func (o *ListUsages200ResponseMeta) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListUsages200ResponseMeta) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListUsages200ResponseMeta) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListUsages200ResponseMeta) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListUsages200ResponseMeta) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListUsages200ResponseMeta) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListUsages200ResponseMeta) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListUsages200ResponseMeta) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


