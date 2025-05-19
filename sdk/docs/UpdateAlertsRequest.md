# UpdateAlertsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | [**UpdateAlertsRequestAlert**](UpdateAlertsRequestAlert.md) |  | 

## Methods

### NewUpdateAlertsRequest

`func NewUpdateAlertsRequest(alert UpdateAlertsRequestAlert, ) *UpdateAlertsRequest`

NewUpdateAlertsRequest instantiates a new UpdateAlertsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertsRequestWithDefaults

`func NewUpdateAlertsRequestWithDefaults() *UpdateAlertsRequest`

NewUpdateAlertsRequestWithDefaults instantiates a new UpdateAlertsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *UpdateAlertsRequest) GetAlert() UpdateAlertsRequestAlert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *UpdateAlertsRequest) GetAlertOk() (*UpdateAlertsRequestAlert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *UpdateAlertsRequest) SetAlert(v UpdateAlertsRequestAlert)`

SetAlert sets Alert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


