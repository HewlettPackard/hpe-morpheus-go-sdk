# ListCheckTypes200ResponseAllOfCheckTypesInner

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

### NewListCheckTypes200ResponseAllOfCheckTypesInner

`func NewListCheckTypes200ResponseAllOfCheckTypesInner() *ListCheckTypes200ResponseAllOfCheckTypesInner`

NewListCheckTypes200ResponseAllOfCheckTypesInner instantiates a new ListCheckTypes200ResponseAllOfCheckTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCheckTypes200ResponseAllOfCheckTypesInnerWithDefaults

`func NewListCheckTypes200ResponseAllOfCheckTypesInnerWithDefaults() *ListCheckTypes200ResponseAllOfCheckTypesInner`

NewListCheckTypes200ResponseAllOfCheckTypesInnerWithDefaults instantiates a new ListCheckTypes200ResponseAllOfCheckTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultInterval

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetDefaultInterval() int64`

GetDefaultInterval returns the DefaultInterval field if non-nil, zero value otherwise.

### GetDefaultIntervalOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetDefaultIntervalOk() (*int64, bool)`

GetDefaultIntervalOk returns a tuple with the DefaultInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInterval

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetDefaultInterval(v int64)`

SetDefaultInterval sets DefaultInterval field to given value.

### HasDefaultInterval

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasDefaultInterval() bool`

HasDefaultInterval returns a boolean if a field has been set.

### GetMetricName

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetInUptime

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetCreateIncident

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetPushOnly

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetPushOnly() bool`

GetPushOnly returns the PushOnly field if non-nil, zero value otherwise.

### GetPushOnlyOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetPushOnlyOk() (*bool, bool)`

GetPushOnlyOk returns a tuple with the PushOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushOnly

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetPushOnly(v bool)`

SetPushOnly sets PushOnly field to given value.

### HasPushOnly

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasPushOnly() bool`

HasPushOnly returns a boolean if a field has been set.

### GetTunnelSupported

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetTunnelSupported() bool`

GetTunnelSupported returns the TunnelSupported field if non-nil, zero value otherwise.

### GetTunnelSupportedOk

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) GetTunnelSupportedOk() (*bool, bool)`

GetTunnelSupportedOk returns a tuple with the TunnelSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelSupported

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) SetTunnelSupported(v bool)`

SetTunnelSupported sets TunnelSupported field to given value.

### HasTunnelSupported

`func (o *ListCheckTypes200ResponseAllOfCheckTypesInner) HasTunnelSupported() bool`

HasTunnelSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


