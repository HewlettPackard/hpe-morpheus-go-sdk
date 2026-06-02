# GetCheckApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApp** | Pointer to [**GetCheckApps200ResponseMonitorApp**](GetCheckApps200ResponseMonitorApp.md) |  | [optional] 
**CheckGroups** | Pointer to [**[]GetCheckApps200ResponseCheckGroupsInner**](GetCheckApps200ResponseCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]GetCheckApps200ResponseChecksInner**](GetCheckApps200ResponseChecksInner.md) |  | [optional] 
**OpenIncidents** | Pointer to [**[]GetCheckApps200ResponseOpenIncidentsInner**](GetCheckApps200ResponseOpenIncidentsInner.md) |  | [optional] 

## Methods

### NewGetCheckApps200Response

`func NewGetCheckApps200Response() *GetCheckApps200Response`

NewGetCheckApps200Response instantiates a new GetCheckApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetMonitorApp

`func (o *GetCheckApps200Response) GetMonitorApp() GetCheckApps200ResponseMonitorApp`

GetMonitorApp returns the MonitorApp field if non-nil, zero value otherwise.

### GetMonitorAppOk

`func (o *GetCheckApps200Response) GetMonitorAppOk() (*GetCheckApps200ResponseMonitorApp, bool)`

GetMonitorAppOk returns a tuple with the MonitorApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApp

`func (o *GetCheckApps200Response) SetMonitorApp(v GetCheckApps200ResponseMonitorApp)`

SetMonitorApp sets MonitorApp field to given value.

### HasMonitorApp

`func (o *GetCheckApps200Response) HasMonitorApp() bool`

HasMonitorApp returns a boolean if a field has been set.

### GetCheckGroups

`func (o *GetCheckApps200Response) GetCheckGroups() []GetCheckApps200ResponseCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *GetCheckApps200Response) GetCheckGroupsOk() (*[]GetCheckApps200ResponseCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *GetCheckApps200Response) SetCheckGroups(v []GetCheckApps200ResponseCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *GetCheckApps200Response) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetChecks

`func (o *GetCheckApps200Response) GetChecks() []GetCheckApps200ResponseChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckApps200Response) GetChecksOk() (*[]GetCheckApps200ResponseChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckApps200Response) SetChecks(v []GetCheckApps200ResponseChecksInner)`

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


