# UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**Integration** | Pointer to [**UpdateMonitoringSettingsRequestMonitoringSettingsServiceNowIntegration**](UpdateMonitoringSettingsRequestMonitoringSettingsServiceNowIntegration.md) |  | [optional] 
**NewIncidentAction** | Pointer to **string** | New Incident Action | [optional] 
**CloseIncidentAction** | Pointer to **string** | Close Incident Action | [optional] 
**InfoMapping** | Pointer to **string** | Info Mapping | [optional] 
**WarningMapping** | Pointer to **string** | Warning Mapping | [optional] 
**CriticalMapping** | Pointer to **string** | Critical Mapping | [optional] 

## Methods

### NewUpdateMonitoringSettingsRequestMonitoringSettingsServiceNow

`func NewUpdateMonitoringSettingsRequestMonitoringSettingsServiceNow() *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow`

NewUpdateMonitoringSettingsRequestMonitoringSettingsServiceNow instantiates a new UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetEnabled

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetIntegration() UpdateMonitoringSettingsRequestMonitoringSettingsServiceNowIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetIntegrationOk() (*UpdateMonitoringSettingsRequestMonitoringSettingsServiceNowIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetIntegration(v UpdateMonitoringSettingsRequestMonitoringSettingsServiceNowIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetNewIncidentAction

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetNewIncidentAction() string`

GetNewIncidentAction returns the NewIncidentAction field if non-nil, zero value otherwise.

### GetNewIncidentActionOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetNewIncidentActionOk() (*string, bool)`

GetNewIncidentActionOk returns a tuple with the NewIncidentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIncidentAction

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetNewIncidentAction(v string)`

SetNewIncidentAction sets NewIncidentAction field to given value.

### HasNewIncidentAction

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasNewIncidentAction() bool`

HasNewIncidentAction returns a boolean if a field has been set.

### GetCloseIncidentAction

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetCloseIncidentAction() string`

GetCloseIncidentAction returns the CloseIncidentAction field if non-nil, zero value otherwise.

### GetCloseIncidentActionOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetCloseIncidentActionOk() (*string, bool)`

GetCloseIncidentActionOk returns a tuple with the CloseIncidentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseIncidentAction

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetCloseIncidentAction(v string)`

SetCloseIncidentAction sets CloseIncidentAction field to given value.

### HasCloseIncidentAction

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasCloseIncidentAction() bool`

HasCloseIncidentAction returns a boolean if a field has been set.

### GetInfoMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetInfoMapping() string`

GetInfoMapping returns the InfoMapping field if non-nil, zero value otherwise.

### GetInfoMappingOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetInfoMappingOk() (*string, bool)`

GetInfoMappingOk returns a tuple with the InfoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetInfoMapping(v string)`

SetInfoMapping sets InfoMapping field to given value.

### HasInfoMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasInfoMapping() bool`

HasInfoMapping returns a boolean if a field has been set.

### GetWarningMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetWarningMapping() string`

GetWarningMapping returns the WarningMapping field if non-nil, zero value otherwise.

### GetWarningMappingOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetWarningMappingOk() (*string, bool)`

GetWarningMappingOk returns a tuple with the WarningMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetWarningMapping(v string)`

SetWarningMapping sets WarningMapping field to given value.

### HasWarningMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasWarningMapping() bool`

HasWarningMapping returns a boolean if a field has been set.

### GetCriticalMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetCriticalMapping() string`

GetCriticalMapping returns the CriticalMapping field if non-nil, zero value otherwise.

### GetCriticalMappingOk

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) GetCriticalMappingOk() (*string, bool)`

GetCriticalMappingOk returns a tuple with the CriticalMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) SetCriticalMapping(v string)`

SetCriticalMapping sets CriticalMapping field to given value.

### HasCriticalMapping

`func (o *UpdateMonitoringSettingsRequestMonitoringSettingsServiceNow) HasCriticalMapping() bool`

HasCriticalMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


