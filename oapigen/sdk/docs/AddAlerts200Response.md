# AddAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to [**ListAlerts200ResponseAllOfAlertsInner**](ListAlerts200ResponseAllOfAlertsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddAlerts200Response

`func NewAddAlerts200Response() *AddAlerts200Response`

NewAddAlerts200Response instantiates a new AddAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlerts200ResponseWithDefaults

`func NewAddAlerts200ResponseWithDefaults() *AddAlerts200Response`

NewAddAlerts200ResponseWithDefaults instantiates a new AddAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *AddAlerts200Response) GetAlert() ListAlerts200ResponseAllOfAlertsInner`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *AddAlerts200Response) GetAlertOk() (*ListAlerts200ResponseAllOfAlertsInner, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *AddAlerts200Response) SetAlert(v ListAlerts200ResponseAllOfAlertsInner)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *AddAlerts200Response) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetSuccess

`func (o *AddAlerts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddAlerts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddAlerts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddAlerts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


