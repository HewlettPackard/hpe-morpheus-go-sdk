# UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorLoadBalancer**](UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorLoadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MonitorType** | Pointer to **string** |  | [optional] 
**MonitorInterval** | Pointer to **int64** |  | [optional] 
**MonitorTimeout** | Pointer to **int64** |  | [optional] 
**SendData** | Pointer to **NullableString** |  | [optional] 
**SendVersion** | Pointer to **string** |  | [optional] 
**SendType** | Pointer to **string** |  | [optional] 
**ReceiveData** | Pointer to **NullableString** |  | [optional] 
**ReceiveCode** | Pointer to **string** |  | [optional] 
**DisabledData** | Pointer to **NullableString** |  | [optional] 
**MonitorUsername** | Pointer to **NullableString** |  | [optional] 
**MonitorPassword** | Pointer to **NullableString** |  | [optional] 
**MonitorPasswordHash** | Pointer to **NullableString** |  | [optional] 
**MonitorDestination** | Pointer to **string** |  | [optional] 
**MonitorReverse** | Pointer to **bool** |  | [optional] 
**MonitorTransparent** | Pointer to **bool** |  | [optional] 
**MonitorAdaptive** | Pointer to **bool** |  | [optional] 
**AliasAddress** | Pointer to **NullableString** |  | [optional] 
**AliasPort** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**MonitorSource** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MaxRetry** | Pointer to **int64** |  | [optional] 
**FallCount** | Pointer to **int64** |  | [optional] 
**RiseCount** | Pointer to **int64** |  | [optional] 
**DataLength** | Pointer to **NullableInt64** |  | [optional] 
**Config** | Pointer to [**UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig**](UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig.md) |  | [optional] 
**CreatedBy** | Pointer to [**UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorCreatedBy**](UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor

`func NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor() *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor`

NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor instantiates a new UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetLoadBalancer() UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetLoadBalancerOk() (*UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetLoadBalancer(v UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorInterval

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorInterval() int64`

GetMonitorInterval returns the MonitorInterval field if non-nil, zero value otherwise.

### GetMonitorIntervalOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorIntervalOk() (*int64, bool)`

GetMonitorIntervalOk returns a tuple with the MonitorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorInterval

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorInterval(v int64)`

SetMonitorInterval sets MonitorInterval field to given value.

### HasMonitorInterval

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorInterval() bool`

HasMonitorInterval returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetSendData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetSendData() string`

GetSendData returns the SendData field if non-nil, zero value otherwise.

### GetSendDataOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetSendDataOk() (*string, bool)`

GetSendDataOk returns a tuple with the SendData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetSendData(v string)`

SetSendData sets SendData field to given value.

### HasSendData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasSendData() bool`

HasSendData returns a boolean if a field has been set.

### SetSendDataNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetSendDataNil(b bool)`

 SetSendDataNil sets the value for SendData to be an explicit nil

### UnsetSendData
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetSendData()`

UnsetSendData ensures that no value is present for SendData, not even an explicit nil
### GetSendVersion

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetSendVersion() string`

GetSendVersion returns the SendVersion field if non-nil, zero value otherwise.

### GetSendVersionOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetSendVersionOk() (*string, bool)`

GetSendVersionOk returns a tuple with the SendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendVersion

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetSendVersion(v string)`

SetSendVersion sets SendVersion field to given value.

### HasSendVersion

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasSendVersion() bool`

HasSendVersion returns a boolean if a field has been set.

### GetSendType

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### GetReceiveData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetReceiveData() string`

GetReceiveData returns the ReceiveData field if non-nil, zero value otherwise.

### GetReceiveDataOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetReceiveDataOk() (*string, bool)`

GetReceiveDataOk returns a tuple with the ReceiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetReceiveData(v string)`

SetReceiveData sets ReceiveData field to given value.

### HasReceiveData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasReceiveData() bool`

HasReceiveData returns a boolean if a field has been set.

### SetReceiveDataNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetReceiveDataNil(b bool)`

 SetReceiveDataNil sets the value for ReceiveData to be an explicit nil

### UnsetReceiveData
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetReceiveData()`

UnsetReceiveData ensures that no value is present for ReceiveData, not even an explicit nil
### GetReceiveCode

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### GetDisabledData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDisabledData() string`

GetDisabledData returns the DisabledData field if non-nil, zero value otherwise.

### GetDisabledDataOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDisabledDataOk() (*string, bool)`

GetDisabledDataOk returns a tuple with the DisabledData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetDisabledData(v string)`

SetDisabledData sets DisabledData field to given value.

### HasDisabledData

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasDisabledData() bool`

HasDisabledData returns a boolean if a field has been set.

### SetDisabledDataNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetDisabledDataNil(b bool)`

 SetDisabledDataNil sets the value for DisabledData to be an explicit nil

### UnsetDisabledData
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetDisabledData()`

UnsetDisabledData ensures that no value is present for DisabledData, not even an explicit nil
### GetMonitorUsername

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorUsername() string`

GetMonitorUsername returns the MonitorUsername field if non-nil, zero value otherwise.

### GetMonitorUsernameOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorUsernameOk() (*string, bool)`

GetMonitorUsernameOk returns a tuple with the MonitorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsername

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorUsername(v string)`

SetMonitorUsername sets MonitorUsername field to given value.

### HasMonitorUsername

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorUsername() bool`

HasMonitorUsername returns a boolean if a field has been set.

### SetMonitorUsernameNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorUsernameNil(b bool)`

 SetMonitorUsernameNil sets the value for MonitorUsername to be an explicit nil

### UnsetMonitorUsername
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetMonitorUsername()`

UnsetMonitorUsername ensures that no value is present for MonitorUsername, not even an explicit nil
### GetMonitorPassword

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorPassword() string`

GetMonitorPassword returns the MonitorPassword field if non-nil, zero value otherwise.

### GetMonitorPasswordOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorPasswordOk() (*string, bool)`

GetMonitorPasswordOk returns a tuple with the MonitorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPassword

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorPassword(v string)`

SetMonitorPassword sets MonitorPassword field to given value.

### HasMonitorPassword

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorPassword() bool`

HasMonitorPassword returns a boolean if a field has been set.

### SetMonitorPasswordNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorPasswordNil(b bool)`

 SetMonitorPasswordNil sets the value for MonitorPassword to be an explicit nil

### UnsetMonitorPassword
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetMonitorPassword()`

UnsetMonitorPassword ensures that no value is present for MonitorPassword, not even an explicit nil
### GetMonitorPasswordHash

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorPasswordHash() string`

GetMonitorPasswordHash returns the MonitorPasswordHash field if non-nil, zero value otherwise.

### GetMonitorPasswordHashOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorPasswordHashOk() (*string, bool)`

GetMonitorPasswordHashOk returns a tuple with the MonitorPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPasswordHash

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorPasswordHash(v string)`

SetMonitorPasswordHash sets MonitorPasswordHash field to given value.

### HasMonitorPasswordHash

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorPasswordHash() bool`

HasMonitorPasswordHash returns a boolean if a field has been set.

### SetMonitorPasswordHashNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorPasswordHashNil(b bool)`

 SetMonitorPasswordHashNil sets the value for MonitorPasswordHash to be an explicit nil

### UnsetMonitorPasswordHash
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetMonitorPasswordHash()`

UnsetMonitorPasswordHash ensures that no value is present for MonitorPasswordHash, not even an explicit nil
### GetMonitorDestination

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorDestination() string`

GetMonitorDestination returns the MonitorDestination field if non-nil, zero value otherwise.

### GetMonitorDestinationOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorDestinationOk() (*string, bool)`

GetMonitorDestinationOk returns a tuple with the MonitorDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDestination

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorDestination(v string)`

SetMonitorDestination sets MonitorDestination field to given value.

### HasMonitorDestination

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorDestination() bool`

HasMonitorDestination returns a boolean if a field has been set.

### GetMonitorReverse

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorReverse() bool`

GetMonitorReverse returns the MonitorReverse field if non-nil, zero value otherwise.

### GetMonitorReverseOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorReverseOk() (*bool, bool)`

GetMonitorReverseOk returns a tuple with the MonitorReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorReverse

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorReverse(v bool)`

SetMonitorReverse sets MonitorReverse field to given value.

### HasMonitorReverse

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorReverse() bool`

HasMonitorReverse returns a boolean if a field has been set.

### GetMonitorTransparent

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorTransparent() bool`

GetMonitorTransparent returns the MonitorTransparent field if non-nil, zero value otherwise.

### GetMonitorTransparentOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorTransparentOk() (*bool, bool)`

GetMonitorTransparentOk returns a tuple with the MonitorTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTransparent

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorTransparent(v bool)`

SetMonitorTransparent sets MonitorTransparent field to given value.

### HasMonitorTransparent

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorTransparent() bool`

HasMonitorTransparent returns a boolean if a field has been set.

### GetMonitorAdaptive

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorAdaptive() bool`

GetMonitorAdaptive returns the MonitorAdaptive field if non-nil, zero value otherwise.

### GetMonitorAdaptiveOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorAdaptiveOk() (*bool, bool)`

GetMonitorAdaptiveOk returns a tuple with the MonitorAdaptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAdaptive

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorAdaptive(v bool)`

SetMonitorAdaptive sets MonitorAdaptive field to given value.

### HasMonitorAdaptive

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorAdaptive() bool`

HasMonitorAdaptive returns a boolean if a field has been set.

### GetAliasAddress

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetAliasAddress() string`

GetAliasAddress returns the AliasAddress field if non-nil, zero value otherwise.

### GetAliasAddressOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetAliasAddressOk() (*string, bool)`

GetAliasAddressOk returns a tuple with the AliasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasAddress

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetAliasAddress(v string)`

SetAliasAddress sets AliasAddress field to given value.

### HasAliasAddress

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasAliasAddress() bool`

HasAliasAddress returns a boolean if a field has been set.

### SetAliasAddressNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetAliasAddressNil(b bool)`

 SetAliasAddressNil sets the value for AliasAddress to be an explicit nil

### UnsetAliasAddress
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetAliasAddress()`

UnsetAliasAddress ensures that no value is present for AliasAddress, not even an explicit nil
### GetAliasPort

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetAliasPort() int64`

GetAliasPort returns the AliasPort field if non-nil, zero value otherwise.

### GetAliasPortOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetAliasPortOk() (*int64, bool)`

GetAliasPortOk returns a tuple with the AliasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPort

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetAliasPort(v int64)`

SetAliasPort sets AliasPort field to given value.

### HasAliasPort

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasAliasPort() bool`

HasAliasPort returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMonitorSource

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorSource() string`

GetMonitorSource returns the MonitorSource field if non-nil, zero value otherwise.

### GetMonitorSourceOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMonitorSourceOk() (*string, bool)`

GetMonitorSourceOk returns a tuple with the MonitorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSource

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMonitorSource(v string)`

SetMonitorSource sets MonitorSource field to given value.

### HasMonitorSource

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMonitorSource() bool`

HasMonitorSource returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetEnabled

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxRetry

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.

### HasMaxRetry

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasMaxRetry() bool`

HasMaxRetry returns a boolean if a field has been set.

### GetFallCount

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetFallCount() int64`

GetFallCount returns the FallCount field if non-nil, zero value otherwise.

### GetFallCountOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetFallCountOk() (*int64, bool)`

GetFallCountOk returns a tuple with the FallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallCount

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetFallCount(v int64)`

SetFallCount sets FallCount field to given value.

### HasFallCount

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasFallCount() bool`

HasFallCount returns a boolean if a field has been set.

### GetRiseCount

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetRiseCount() int64`

GetRiseCount returns the RiseCount field if non-nil, zero value otherwise.

### GetRiseCountOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetRiseCountOk() (*int64, bool)`

GetRiseCountOk returns a tuple with the RiseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseCount

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetRiseCount(v int64)`

SetRiseCount sets RiseCount field to given value.

### HasRiseCount

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasRiseCount() bool`

HasRiseCount returns a boolean if a field has been set.

### GetDataLength

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDataLength() int64`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDataLengthOk() (*int64, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetDataLength(v int64)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### SetDataLengthNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetDataLengthNil(b bool)`

 SetDataLengthNil sets the value for DataLength to be an explicit nil

### UnsetDataLength
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) UnsetDataLength()`

UnsetDataLength ensures that no value is present for DataLength, not even an explicit nil
### GetConfig

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetConfig() UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetConfigOk() (*UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetConfig(v UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetCreatedBy() UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetCreatedByOk() (*UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetCreatedBy(v UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


