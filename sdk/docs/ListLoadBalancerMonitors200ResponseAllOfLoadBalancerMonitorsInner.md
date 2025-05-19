# ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MonitorType** | Pointer to **string** |  | [optional] 
**MonitorInterval** | Pointer to **int64** |  | [optional] 
**MonitorTimeout** | Pointer to **int64** |  | [optional] 
**SendData** | Pointer to **string** |  | [optional] 
**SendVersion** | Pointer to **string** |  | [optional] 
**SendType** | Pointer to **string** |  | [optional] 
**ReceiveData** | Pointer to **string** |  | [optional] 
**ReceiveCode** | Pointer to **string** |  | [optional] 
**DisabledData** | Pointer to **string** |  | [optional] 
**MonitorUsername** | Pointer to **string** |  | [optional] 
**MonitorPassword** | Pointer to **string** |  | [optional] 
**MonitorDestination** | Pointer to **string** |  | [optional] 
**MonitorReverse** | Pointer to **bool** |  | [optional] 
**MonitorTransparent** | Pointer to **bool** |  | [optional] 
**MonitorAdaptive** | Pointer to **bool** |  | [optional] 
**AliasAddress** | Pointer to **string** |  | [optional] 
**AliasPort** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**MonitorSource** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MaxRetry** | Pointer to **int64** |  | [optional] 
**FallCount** | Pointer to **int64** |  | [optional] 
**RiseCount** | Pointer to **int64** |  | [optional] 
**DataLength** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner

`func NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner() *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner`

NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner instantiates a new ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerWithDefaults

`func NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerWithDefaults() *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner`

NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerWithDefaults instantiates a new ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetLoadBalancer() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetLoadBalancerOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetLoadBalancer(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVisibility

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorInterval

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorInterval() int64`

GetMonitorInterval returns the MonitorInterval field if non-nil, zero value otherwise.

### GetMonitorIntervalOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorIntervalOk() (*int64, bool)`

GetMonitorIntervalOk returns a tuple with the MonitorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorInterval

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorInterval(v int64)`

SetMonitorInterval sets MonitorInterval field to given value.

### HasMonitorInterval

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorInterval() bool`

HasMonitorInterval returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetSendData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetSendData() string`

GetSendData returns the SendData field if non-nil, zero value otherwise.

### GetSendDataOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetSendDataOk() (*string, bool)`

GetSendDataOk returns a tuple with the SendData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetSendData(v string)`

SetSendData sets SendData field to given value.

### HasSendData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasSendData() bool`

HasSendData returns a boolean if a field has been set.

### GetSendVersion

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetSendVersion() string`

GetSendVersion returns the SendVersion field if non-nil, zero value otherwise.

### GetSendVersionOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetSendVersionOk() (*string, bool)`

GetSendVersionOk returns a tuple with the SendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendVersion

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetSendVersion(v string)`

SetSendVersion sets SendVersion field to given value.

### HasSendVersion

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasSendVersion() bool`

HasSendVersion returns a boolean if a field has been set.

### GetSendType

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### GetReceiveData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetReceiveData() string`

GetReceiveData returns the ReceiveData field if non-nil, zero value otherwise.

### GetReceiveDataOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetReceiveDataOk() (*string, bool)`

GetReceiveDataOk returns a tuple with the ReceiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetReceiveData(v string)`

SetReceiveData sets ReceiveData field to given value.

### HasReceiveData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasReceiveData() bool`

HasReceiveData returns a boolean if a field has been set.

### GetReceiveCode

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### GetDisabledData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDisabledData() string`

GetDisabledData returns the DisabledData field if non-nil, zero value otherwise.

### GetDisabledDataOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDisabledDataOk() (*string, bool)`

GetDisabledDataOk returns a tuple with the DisabledData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetDisabledData(v string)`

SetDisabledData sets DisabledData field to given value.

### HasDisabledData

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasDisabledData() bool`

HasDisabledData returns a boolean if a field has been set.

### GetMonitorUsername

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorUsername() string`

GetMonitorUsername returns the MonitorUsername field if non-nil, zero value otherwise.

### GetMonitorUsernameOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorUsernameOk() (*string, bool)`

GetMonitorUsernameOk returns a tuple with the MonitorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsername

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorUsername(v string)`

SetMonitorUsername sets MonitorUsername field to given value.

### HasMonitorUsername

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorUsername() bool`

HasMonitorUsername returns a boolean if a field has been set.

### GetMonitorPassword

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorPassword() string`

GetMonitorPassword returns the MonitorPassword field if non-nil, zero value otherwise.

### GetMonitorPasswordOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorPasswordOk() (*string, bool)`

GetMonitorPasswordOk returns a tuple with the MonitorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPassword

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorPassword(v string)`

SetMonitorPassword sets MonitorPassword field to given value.

### HasMonitorPassword

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorPassword() bool`

HasMonitorPassword returns a boolean if a field has been set.

### GetMonitorDestination

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorDestination() string`

GetMonitorDestination returns the MonitorDestination field if non-nil, zero value otherwise.

### GetMonitorDestinationOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorDestinationOk() (*string, bool)`

GetMonitorDestinationOk returns a tuple with the MonitorDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDestination

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorDestination(v string)`

SetMonitorDestination sets MonitorDestination field to given value.

### HasMonitorDestination

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorDestination() bool`

HasMonitorDestination returns a boolean if a field has been set.

### GetMonitorReverse

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorReverse() bool`

GetMonitorReverse returns the MonitorReverse field if non-nil, zero value otherwise.

### GetMonitorReverseOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorReverseOk() (*bool, bool)`

GetMonitorReverseOk returns a tuple with the MonitorReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorReverse

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorReverse(v bool)`

SetMonitorReverse sets MonitorReverse field to given value.

### HasMonitorReverse

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorReverse() bool`

HasMonitorReverse returns a boolean if a field has been set.

### GetMonitorTransparent

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorTransparent() bool`

GetMonitorTransparent returns the MonitorTransparent field if non-nil, zero value otherwise.

### GetMonitorTransparentOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorTransparentOk() (*bool, bool)`

GetMonitorTransparentOk returns a tuple with the MonitorTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTransparent

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorTransparent(v bool)`

SetMonitorTransparent sets MonitorTransparent field to given value.

### HasMonitorTransparent

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorTransparent() bool`

HasMonitorTransparent returns a boolean if a field has been set.

### GetMonitorAdaptive

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorAdaptive() bool`

GetMonitorAdaptive returns the MonitorAdaptive field if non-nil, zero value otherwise.

### GetMonitorAdaptiveOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorAdaptiveOk() (*bool, bool)`

GetMonitorAdaptiveOk returns a tuple with the MonitorAdaptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAdaptive

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorAdaptive(v bool)`

SetMonitorAdaptive sets MonitorAdaptive field to given value.

### HasMonitorAdaptive

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorAdaptive() bool`

HasMonitorAdaptive returns a boolean if a field has been set.

### GetAliasAddress

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetAliasAddress() string`

GetAliasAddress returns the AliasAddress field if non-nil, zero value otherwise.

### GetAliasAddressOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetAliasAddressOk() (*string, bool)`

GetAliasAddressOk returns a tuple with the AliasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasAddress

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetAliasAddress(v string)`

SetAliasAddress sets AliasAddress field to given value.

### HasAliasAddress

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasAliasAddress() bool`

HasAliasAddress returns a boolean if a field has been set.

### GetAliasPort

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetAliasPort() int64`

GetAliasPort returns the AliasPort field if non-nil, zero value otherwise.

### GetAliasPortOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetAliasPortOk() (*int64, bool)`

GetAliasPortOk returns a tuple with the AliasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPort

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetAliasPort(v int64)`

SetAliasPort sets AliasPort field to given value.

### HasAliasPort

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasAliasPort() bool`

HasAliasPort returns a boolean if a field has been set.

### GetInternalId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMonitorSource

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorSource() string`

GetMonitorSource returns the MonitorSource field if non-nil, zero value otherwise.

### GetMonitorSourceOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMonitorSourceOk() (*string, bool)`

GetMonitorSourceOk returns a tuple with the MonitorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSource

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMonitorSource(v string)`

SetMonitorSource sets MonitorSource field to given value.

### HasMonitorSource

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMonitorSource() bool`

HasMonitorSource returns a boolean if a field has been set.

### GetStatus

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetEnabled

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxRetry

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.

### HasMaxRetry

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasMaxRetry() bool`

HasMaxRetry returns a boolean if a field has been set.

### GetFallCount

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetFallCount() int64`

GetFallCount returns the FallCount field if non-nil, zero value otherwise.

### GetFallCountOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetFallCountOk() (*int64, bool)`

GetFallCountOk returns a tuple with the FallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallCount

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetFallCount(v int64)`

SetFallCount sets FallCount field to given value.

### HasFallCount

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasFallCount() bool`

HasFallCount returns a boolean if a field has been set.

### GetRiseCount

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetRiseCount() int64`

GetRiseCount returns the RiseCount field if non-nil, zero value otherwise.

### GetRiseCountOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetRiseCountOk() (*int64, bool)`

GetRiseCountOk returns a tuple with the RiseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseCount

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetRiseCount(v int64)`

SetRiseCount sets RiseCount field to given value.

### HasRiseCount

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasRiseCount() bool`

HasRiseCount returns a boolean if a field has been set.

### GetDataLength

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDataLength() string`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDataLengthOk() (*string, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetDataLength(v string)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### GetConfig

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


