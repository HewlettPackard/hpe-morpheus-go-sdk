# ListExecuteSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to [**[]ListExecuteSchedules200ResponseAllOfSchedulesInner**](ListExecuteSchedules200ResponseAllOfSchedulesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListExecuteSchedules200Response

`func NewListExecuteSchedules200Response() *ListExecuteSchedules200Response`

NewListExecuteSchedules200Response instantiates a new ListExecuteSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExecuteSchedules200ResponseWithDefaults

`func NewListExecuteSchedules200ResponseWithDefaults() *ListExecuteSchedules200Response`

NewListExecuteSchedules200ResponseWithDefaults instantiates a new ListExecuteSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *ListExecuteSchedules200Response) GetSchedules() []ListExecuteSchedules200ResponseAllOfSchedulesInner`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ListExecuteSchedules200Response) GetSchedulesOk() (*[]ListExecuteSchedules200ResponseAllOfSchedulesInner, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ListExecuteSchedules200Response) SetSchedules(v []ListExecuteSchedules200ResponseAllOfSchedulesInner)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ListExecuteSchedules200Response) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetMeta

`func (o *ListExecuteSchedules200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListExecuteSchedules200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListExecuteSchedules200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListExecuteSchedules200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


