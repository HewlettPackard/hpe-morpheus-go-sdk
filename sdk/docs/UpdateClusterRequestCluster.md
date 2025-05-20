# UpdateClusterRequestCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster name | [optional] 
**Description** | Pointer to **string** | Cluster description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Cluster enabled | [optional] 
**UseAgent** | Pointer to **bool** | Use the Agent to relay communications for the Kubernetes API instead of direct. | [optional] 
**ServiceUrl** | Pointer to **string** | Cluster API Url | [optional] 
**ServiceToken** | Pointer to **string** | Cluster API token | [optional] 
**Refresh** | Pointer to **bool** | Queue cluster refresh | [optional] 
**Managed** | Pointer to **bool** | Cluster managed | [optional] 
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] 
**Integrations** | Pointer to [**[]UpdateClusterRequestClusterIntegrationsInner**](UpdateClusterRequestClusterIntegrationsInner.md) | Cluster integrations | [optional] 

## Methods

### NewUpdateClusterRequestCluster

`func NewUpdateClusterRequestCluster() *UpdateClusterRequestCluster`

NewUpdateClusterRequestCluster instantiates a new UpdateClusterRequestCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterRequestClusterWithDefaults

`func NewUpdateClusterRequestClusterWithDefaults() *UpdateClusterRequestCluster`

NewUpdateClusterRequestClusterWithDefaults instantiates a new UpdateClusterRequestCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterRequestCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterRequestCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterRequestCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterRequestCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateClusterRequestCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClusterRequestCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClusterRequestCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClusterRequestCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateClusterRequestCluster) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateClusterRequestCluster) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateClusterRequestCluster) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateClusterRequestCluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateClusterRequestCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateClusterRequestCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateClusterRequestCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateClusterRequestCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUseAgent

`func (o *UpdateClusterRequestCluster) GetUseAgent() bool`

GetUseAgent returns the UseAgent field if non-nil, zero value otherwise.

### GetUseAgentOk

`func (o *UpdateClusterRequestCluster) GetUseAgentOk() (*bool, bool)`

GetUseAgentOk returns a tuple with the UseAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgent

`func (o *UpdateClusterRequestCluster) SetUseAgent(v bool)`

SetUseAgent sets UseAgent field to given value.

### HasUseAgent

`func (o *UpdateClusterRequestCluster) HasUseAgent() bool`

HasUseAgent returns a boolean if a field has been set.

### GetServiceUrl

`func (o *UpdateClusterRequestCluster) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateClusterRequestCluster) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateClusterRequestCluster) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateClusterRequestCluster) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceToken

`func (o *UpdateClusterRequestCluster) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *UpdateClusterRequestCluster) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *UpdateClusterRequestCluster) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *UpdateClusterRequestCluster) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetRefresh

`func (o *UpdateClusterRequestCluster) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *UpdateClusterRequestCluster) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *UpdateClusterRequestCluster) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *UpdateClusterRequestCluster) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetManaged

`func (o *UpdateClusterRequestCluster) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *UpdateClusterRequestCluster) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *UpdateClusterRequestCluster) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *UpdateClusterRequestCluster) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *UpdateClusterRequestCluster) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *UpdateClusterRequestCluster) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *UpdateClusterRequestCluster) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *UpdateClusterRequestCluster) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetIntegrations

`func (o *UpdateClusterRequestCluster) GetIntegrations() []UpdateClusterRequestClusterIntegrationsInner`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *UpdateClusterRequestCluster) GetIntegrationsOk() (*[]UpdateClusterRequestClusterIntegrationsInner, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *UpdateClusterRequestCluster) SetIntegrations(v []UpdateClusterRequestClusterIntegrationsInner)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *UpdateClusterRequestCluster) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


