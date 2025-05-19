# UpdateCheckApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApp** | Pointer to [**GetAlerts200ResponseAllOfAppsInner**](GetAlerts200ResponseAllOfAppsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCheckApps200Response

`func NewUpdateCheckApps200Response() *UpdateCheckApps200Response`

NewUpdateCheckApps200Response instantiates a new UpdateCheckApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCheckApps200ResponseWithDefaults

`func NewUpdateCheckApps200ResponseWithDefaults() *UpdateCheckApps200Response`

NewUpdateCheckApps200ResponseWithDefaults instantiates a new UpdateCheckApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorApp

`func (o *UpdateCheckApps200Response) GetMonitorApp() GetAlerts200ResponseAllOfAppsInner`

GetMonitorApp returns the MonitorApp field if non-nil, zero value otherwise.

### GetMonitorAppOk

`func (o *UpdateCheckApps200Response) GetMonitorAppOk() (*GetAlerts200ResponseAllOfAppsInner, bool)`

GetMonitorAppOk returns a tuple with the MonitorApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApp

`func (o *UpdateCheckApps200Response) SetMonitorApp(v GetAlerts200ResponseAllOfAppsInner)`

SetMonitorApp sets MonitorApp field to given value.

### HasMonitorApp

`func (o *UpdateCheckApps200Response) HasMonitorApp() bool`

HasMonitorApp returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCheckApps200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCheckApps200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCheckApps200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCheckApps200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


