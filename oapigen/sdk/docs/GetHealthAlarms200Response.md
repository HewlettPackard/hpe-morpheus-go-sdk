# GetHealthAlarms200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alarm** | Pointer to [**GetHealthAlarms200ResponseAllOfAlarm**](GetHealthAlarms200ResponseAllOfAlarm.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetHealthAlarms200Response

`func NewGetHealthAlarms200Response() *GetHealthAlarms200Response`

NewGetHealthAlarms200Response instantiates a new GetHealthAlarms200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetAlarm

`func (o *GetHealthAlarms200Response) GetAlarm() GetHealthAlarms200ResponseAllOfAlarm`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *GetHealthAlarms200Response) GetAlarmOk() (*GetHealthAlarms200ResponseAllOfAlarm, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *GetHealthAlarms200Response) SetAlarm(v GetHealthAlarms200ResponseAllOfAlarm)`

SetAlarm sets Alarm field to given value.

### HasAlarm

`func (o *GetHealthAlarms200Response) HasAlarm() bool`

HasAlarm returns a boolean if a field has been set.

### GetMeta

`func (o *GetHealthAlarms200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetHealthAlarms200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetHealthAlarms200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetHealthAlarms200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


