# GetHealthAlarms200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alarm** | Pointer to [**ListHealthAlarms200ResponseAllOfAlarmInner**](ListHealthAlarms200ResponseAllOfAlarmInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetHealthAlarms200Response

`func NewGetHealthAlarms200Response() *GetHealthAlarms200Response`

NewGetHealthAlarms200Response instantiates a new GetHealthAlarms200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHealthAlarms200ResponseWithDefaults

`func NewGetHealthAlarms200ResponseWithDefaults() *GetHealthAlarms200Response`

NewGetHealthAlarms200ResponseWithDefaults instantiates a new GetHealthAlarms200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarm

`func (o *GetHealthAlarms200Response) GetAlarm() ListHealthAlarms200ResponseAllOfAlarmInner`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *GetHealthAlarms200Response) GetAlarmOk() (*ListHealthAlarms200ResponseAllOfAlarmInner, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *GetHealthAlarms200Response) SetAlarm(v ListHealthAlarms200ResponseAllOfAlarmInner)`

SetAlarm sets Alarm field to given value.

### HasAlarm

`func (o *GetHealthAlarms200Response) HasAlarm() bool`

HasAlarm returns a boolean if a field has been set.

### GetMeta

`func (o *GetHealthAlarms200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetHealthAlarms200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetHealthAlarms200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetHealthAlarms200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


