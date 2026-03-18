# GetCheckTypes200ResponseCheckType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultInterval** | Pointer to **int64** |  | [optional] 
**MetricName** | Pointer to **string** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**PushOnly** | Pointer to **bool** |  | [optional] 
**TunnelSupported** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetCheckTypes200ResponseCheckType

`func NewGetCheckTypes200ResponseCheckType() *GetCheckTypes200ResponseCheckType`

NewGetCheckTypes200ResponseCheckType instantiates a new GetCheckTypes200ResponseCheckType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCheckTypes200ResponseCheckTypeWithDefaults

`func NewGetCheckTypes200ResponseCheckTypeWithDefaults() *GetCheckTypes200ResponseCheckType`

NewGetCheckTypes200ResponseCheckTypeWithDefaults instantiates a new GetCheckTypes200ResponseCheckType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCheckTypes200ResponseCheckType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCheckTypes200ResponseCheckType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCheckTypes200ResponseCheckType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCheckTypes200ResponseCheckType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetCheckTypes200ResponseCheckType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCheckTypes200ResponseCheckType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCheckTypes200ResponseCheckType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCheckTypes200ResponseCheckType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetCheckTypes200ResponseCheckType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCheckTypes200ResponseCheckType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCheckTypes200ResponseCheckType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCheckTypes200ResponseCheckType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultInterval

`func (o *GetCheckTypes200ResponseCheckType) GetDefaultInterval() int64`

GetDefaultInterval returns the DefaultInterval field if non-nil, zero value otherwise.

### GetDefaultIntervalOk

`func (o *GetCheckTypes200ResponseCheckType) GetDefaultIntervalOk() (*int64, bool)`

GetDefaultIntervalOk returns a tuple with the DefaultInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInterval

`func (o *GetCheckTypes200ResponseCheckType) SetDefaultInterval(v int64)`

SetDefaultInterval sets DefaultInterval field to given value.

### HasDefaultInterval

`func (o *GetCheckTypes200ResponseCheckType) HasDefaultInterval() bool`

HasDefaultInterval returns a boolean if a field has been set.

### GetMetricName

`func (o *GetCheckTypes200ResponseCheckType) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *GetCheckTypes200ResponseCheckType) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *GetCheckTypes200ResponseCheckType) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *GetCheckTypes200ResponseCheckType) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetInUptime

`func (o *GetCheckTypes200ResponseCheckType) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *GetCheckTypes200ResponseCheckType) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *GetCheckTypes200ResponseCheckType) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *GetCheckTypes200ResponseCheckType) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetCreateIncident

`func (o *GetCheckTypes200ResponseCheckType) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *GetCheckTypes200ResponseCheckType) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *GetCheckTypes200ResponseCheckType) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *GetCheckTypes200ResponseCheckType) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetPushOnly

`func (o *GetCheckTypes200ResponseCheckType) GetPushOnly() bool`

GetPushOnly returns the PushOnly field if non-nil, zero value otherwise.

### GetPushOnlyOk

`func (o *GetCheckTypes200ResponseCheckType) GetPushOnlyOk() (*bool, bool)`

GetPushOnlyOk returns a tuple with the PushOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushOnly

`func (o *GetCheckTypes200ResponseCheckType) SetPushOnly(v bool)`

SetPushOnly sets PushOnly field to given value.

### HasPushOnly

`func (o *GetCheckTypes200ResponseCheckType) HasPushOnly() bool`

HasPushOnly returns a boolean if a field has been set.

### GetTunnelSupported

`func (o *GetCheckTypes200ResponseCheckType) GetTunnelSupported() bool`

GetTunnelSupported returns the TunnelSupported field if non-nil, zero value otherwise.

### GetTunnelSupportedOk

`func (o *GetCheckTypes200ResponseCheckType) GetTunnelSupportedOk() (*bool, bool)`

GetTunnelSupportedOk returns a tuple with the TunnelSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelSupported

`func (o *GetCheckTypes200ResponseCheckType) SetTunnelSupported(v bool)`

SetTunnelSupported sets TunnelSupported field to given value.

### HasTunnelSupported

`func (o *GetCheckTypes200ResponseCheckType) HasTunnelSupported() bool`

HasTunnelSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


