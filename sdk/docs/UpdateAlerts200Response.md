# UpdateAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Alert** | Pointer to [**ListAlerts200ResponseAllOfAlertsInner**](ListAlerts200ResponseAllOfAlertsInner.md) |  | [optional] 

## Methods

### NewUpdateAlerts200Response

`func NewUpdateAlerts200Response() *UpdateAlerts200Response`

NewUpdateAlerts200Response instantiates a new UpdateAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlerts200ResponseWithDefaults

`func NewUpdateAlerts200ResponseWithDefaults() *UpdateAlerts200Response`

NewUpdateAlerts200ResponseWithDefaults instantiates a new UpdateAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateAlerts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateAlerts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateAlerts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateAlerts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAlert

`func (o *UpdateAlerts200Response) GetAlert() ListAlerts200ResponseAllOfAlertsInner`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *UpdateAlerts200Response) GetAlertOk() (*ListAlerts200ResponseAllOfAlertsInner, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *UpdateAlerts200Response) SetAlert(v ListAlerts200ResponseAllOfAlertsInner)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *UpdateAlerts200Response) HasAlert() bool`

HasAlert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


