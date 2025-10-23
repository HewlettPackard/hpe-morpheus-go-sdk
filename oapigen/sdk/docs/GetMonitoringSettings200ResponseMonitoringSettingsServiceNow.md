# GetMonitoringSettings200ResponseMonitoringSettingsServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**Integration** | Pointer to [**GetMonitoringSettings200ResponseMonitoringSettingsServiceNowIntegration**](GetMonitoringSettings200ResponseMonitoringSettingsServiceNowIntegration.md) |  | [optional] 
**NewIncidentAction** | Pointer to **string** | New Incident Action | [optional] 
**CloseIncidentAction** | Pointer to **string** | Close Incident Action | [optional] 
**InfoMapping** | Pointer to **string** | Info Mapping | [optional] 
**WarningMapping** | Pointer to **string** | Warning Mapping | [optional] 
**CriticalMapping** | Pointer to **string** | Critical Mapping | [optional] 

## Methods

### NewGetMonitoringSettings200ResponseMonitoringSettingsServiceNow

`func NewGetMonitoringSettings200ResponseMonitoringSettingsServiceNow() *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow`

NewGetMonitoringSettings200ResponseMonitoringSettingsServiceNow instantiates a new GetMonitoringSettings200ResponseMonitoringSettingsServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMonitoringSettings200ResponseMonitoringSettingsServiceNowWithDefaults

`func NewGetMonitoringSettings200ResponseMonitoringSettingsServiceNowWithDefaults() *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow`

NewGetMonitoringSettings200ResponseMonitoringSettingsServiceNowWithDefaults instantiates a new GetMonitoringSettings200ResponseMonitoringSettingsServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetIntegration() GetMonitoringSettings200ResponseMonitoringSettingsServiceNowIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetIntegrationOk() (*GetMonitoringSettings200ResponseMonitoringSettingsServiceNowIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetIntegration(v GetMonitoringSettings200ResponseMonitoringSettingsServiceNowIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetNewIncidentAction

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetNewIncidentAction() string`

GetNewIncidentAction returns the NewIncidentAction field if non-nil, zero value otherwise.

### GetNewIncidentActionOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetNewIncidentActionOk() (*string, bool)`

GetNewIncidentActionOk returns a tuple with the NewIncidentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIncidentAction

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetNewIncidentAction(v string)`

SetNewIncidentAction sets NewIncidentAction field to given value.

### HasNewIncidentAction

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasNewIncidentAction() bool`

HasNewIncidentAction returns a boolean if a field has been set.

### GetCloseIncidentAction

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetCloseIncidentAction() string`

GetCloseIncidentAction returns the CloseIncidentAction field if non-nil, zero value otherwise.

### GetCloseIncidentActionOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetCloseIncidentActionOk() (*string, bool)`

GetCloseIncidentActionOk returns a tuple with the CloseIncidentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseIncidentAction

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetCloseIncidentAction(v string)`

SetCloseIncidentAction sets CloseIncidentAction field to given value.

### HasCloseIncidentAction

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasCloseIncidentAction() bool`

HasCloseIncidentAction returns a boolean if a field has been set.

### GetInfoMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetInfoMapping() string`

GetInfoMapping returns the InfoMapping field if non-nil, zero value otherwise.

### GetInfoMappingOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetInfoMappingOk() (*string, bool)`

GetInfoMappingOk returns a tuple with the InfoMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetInfoMapping(v string)`

SetInfoMapping sets InfoMapping field to given value.

### HasInfoMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasInfoMapping() bool`

HasInfoMapping returns a boolean if a field has been set.

### GetWarningMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetWarningMapping() string`

GetWarningMapping returns the WarningMapping field if non-nil, zero value otherwise.

### GetWarningMappingOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetWarningMappingOk() (*string, bool)`

GetWarningMappingOk returns a tuple with the WarningMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetWarningMapping(v string)`

SetWarningMapping sets WarningMapping field to given value.

### HasWarningMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasWarningMapping() bool`

HasWarningMapping returns a boolean if a field has been set.

### GetCriticalMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetCriticalMapping() string`

GetCriticalMapping returns the CriticalMapping field if non-nil, zero value otherwise.

### GetCriticalMappingOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) GetCriticalMappingOk() (*string, bool)`

GetCriticalMappingOk returns a tuple with the CriticalMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) SetCriticalMapping(v string)`

SetCriticalMapping sets CriticalMapping field to given value.

### HasCriticalMapping

`func (o *GetMonitoringSettings200ResponseMonitoringSettingsServiceNow) HasCriticalMapping() bool`

HasCriticalMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


