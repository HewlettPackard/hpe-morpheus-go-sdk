# GetInstanceThreshold200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**InstanceThreshold** | Pointer to [**GetInstanceThreshold200ResponseInstanceThreshold**](GetInstanceThreshold200ResponseInstanceThreshold.md) |  | [optional] 
**InstanceSchedules** | Pointer to [**[]GetInstanceThreshold200ResponseInstanceSchedulesInner**](GetInstanceThreshold200ResponseInstanceSchedulesInner.md) |  | [optional] 

## Methods

### NewGetInstanceThreshold200Response

`func NewGetInstanceThreshold200Response() *GetInstanceThreshold200Response`

NewGetInstanceThreshold200Response instantiates a new GetInstanceThreshold200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceThreshold200ResponseWithDefaults

`func NewGetInstanceThreshold200ResponseWithDefaults() *GetInstanceThreshold200Response`

NewGetInstanceThreshold200ResponseWithDefaults instantiates a new GetInstanceThreshold200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *GetInstanceThreshold200Response) GetInstance() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetInstanceThreshold200Response) GetInstanceOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetInstanceThreshold200Response) SetInstance(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetInstanceThreshold200Response) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetInstanceThreshold

`func (o *GetInstanceThreshold200Response) GetInstanceThreshold() GetInstanceThreshold200ResponseInstanceThreshold`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *GetInstanceThreshold200Response) GetInstanceThresholdOk() (*GetInstanceThreshold200ResponseInstanceThreshold, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *GetInstanceThreshold200Response) SetInstanceThreshold(v GetInstanceThreshold200ResponseInstanceThreshold)`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *GetInstanceThreshold200Response) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### GetInstanceSchedules

`func (o *GetInstanceThreshold200Response) GetInstanceSchedules() []GetInstanceThreshold200ResponseInstanceSchedulesInner`

GetInstanceSchedules returns the InstanceSchedules field if non-nil, zero value otherwise.

### GetInstanceSchedulesOk

`func (o *GetInstanceThreshold200Response) GetInstanceSchedulesOk() (*[]GetInstanceThreshold200ResponseInstanceSchedulesInner, bool)`

GetInstanceSchedulesOk returns a tuple with the InstanceSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSchedules

`func (o *GetInstanceThreshold200Response) SetInstanceSchedules(v []GetInstanceThreshold200ResponseInstanceSchedulesInner)`

SetInstanceSchedules sets InstanceSchedules field to given value.

### HasInstanceSchedules

`func (o *GetInstanceThreshold200Response) HasInstanceSchedules() bool`

HasInstanceSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


