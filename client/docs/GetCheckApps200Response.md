# GetCheckApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApp** | Pointer to [**GetAlerts200ResponseAllOfAppsInner**](GetAlerts200ResponseAllOfAppsInner.md) |  | [optional] 
**CheckGroups** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInner**](GetAlerts200ResponseAllOfCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInner**](GetAlerts200ResponseAllOfChecksInner.md) |  | [optional] 
**OpenIncidents** | Pointer to [**[]GetCheckApps200ResponseOpenIncidentsInner**](GetCheckApps200ResponseOpenIncidentsInner.md) |  | [optional] 

## Methods

### NewGetCheckApps200Response

`func NewGetCheckApps200Response() *GetCheckApps200Response`

NewGetCheckApps200Response instantiates a new GetCheckApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCheckApps200ResponseWithDefaults

`func NewGetCheckApps200ResponseWithDefaults() *GetCheckApps200Response`

NewGetCheckApps200ResponseWithDefaults instantiates a new GetCheckApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorApp

`func (o *GetCheckApps200Response) GetMonitorApp() GetAlerts200ResponseAllOfAppsInner`

GetMonitorApp returns the MonitorApp field if non-nil, zero value otherwise.

### GetMonitorAppOk

`func (o *GetCheckApps200Response) GetMonitorAppOk() (*GetAlerts200ResponseAllOfAppsInner, bool)`

GetMonitorAppOk returns a tuple with the MonitorApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApp

`func (o *GetCheckApps200Response) SetMonitorApp(v GetAlerts200ResponseAllOfAppsInner)`

SetMonitorApp sets MonitorApp field to given value.

### HasMonitorApp

`func (o *GetCheckApps200Response) HasMonitorApp() bool`

HasMonitorApp returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetCheckApps200Response) GetCheckGroups() []GetAlerts200ResponseAllOfCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetCheckApps200Response) GetCheckGroupsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetCheckApps200Response) SetCheckGroups(v []GetAlerts200ResponseAllOfCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetCheckApps200Response) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *GetCheckApps200Response) GetChecks() []GetAlerts200ResponseAllOfChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckApps200Response) GetChecksOk() (*[]GetAlerts200ResponseAllOfChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckApps200Response) SetChecks(v []GetAlerts200ResponseAllOfChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetCheckApps200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetOpenIncidents

`func (o *GetCheckApps200Response) GetOpenIncidents() []GetCheckApps200ResponseOpenIncidentsInner`

GetOpenIncidents returns the OpenIncidents field if non-nil, zero value otherwise.

### GetOpenIncidentsOk

`func (o *GetCheckApps200Response) GetOpenIncidentsOk() (*[]GetCheckApps200ResponseOpenIncidentsInner, bool)`

GetOpenIncidentsOk returns a tuple with the OpenIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIncidents

`func (o *GetCheckApps200Response) SetOpenIncidents(v []GetCheckApps200ResponseOpenIncidentsInner)`

SetOpenIncidents sets OpenIncidents field to given value.

### HasOpenIncidents

`func (o *GetCheckApps200Response) HasOpenIncidents() bool`

HasOpenIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


