# ListPowerSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to [**[]ListPowerSchedules200ResponseAllOfSchedulesInner**](ListPowerSchedules200ResponseAllOfSchedulesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPowerSchedules200Response

`func NewListPowerSchedules200Response() *ListPowerSchedules200Response`

NewListPowerSchedules200Response instantiates a new ListPowerSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPowerSchedules200ResponseWithDefaults

`func NewListPowerSchedules200ResponseWithDefaults() *ListPowerSchedules200Response`

NewListPowerSchedules200ResponseWithDefaults instantiates a new ListPowerSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *ListPowerSchedules200Response) GetSchedules() []ListPowerSchedules200ResponseAllOfSchedulesInner`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ListPowerSchedules200Response) GetSchedulesOk() (*[]ListPowerSchedules200ResponseAllOfSchedulesInner, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ListPowerSchedules200Response) SetSchedules(v []ListPowerSchedules200ResponseAllOfSchedulesInner)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ListPowerSchedules200Response) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetMeta

`func (o *ListPowerSchedules200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPowerSchedules200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPowerSchedules200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPowerSchedules200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


