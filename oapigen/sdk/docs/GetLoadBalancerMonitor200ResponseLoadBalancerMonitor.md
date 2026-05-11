# GetLoadBalancerMonitor200ResponseLoadBalancerMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**GetLoadBalancerMonitor200ResponseLoadBalancerMonitorLoadBalancer**](GetLoadBalancerMonitor200ResponseLoadBalancerMonitorLoadBalancer.md) |  | [optional] 
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
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to [**GetLoadBalancerMonitor200ResponseLoadBalancerMonitorCreatedBy**](GetLoadBalancerMonitor200ResponseLoadBalancerMonitorCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitor

`func NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitor() *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor`

NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitor instantiates a new GetLoadBalancerMonitor200ResponseLoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitorWithDefaults

`func NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitorWithDefaults() *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor`

NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitorWithDefaults instantiates a new GetLoadBalancerMonitor200ResponseLoadBalancerMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetLoadBalancer() GetLoadBalancerMonitor200ResponseLoadBalancerMonitorLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetLoadBalancerOk() (*GetLoadBalancerMonitor200ResponseLoadBalancerMonitorLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetLoadBalancer(v GetLoadBalancerMonitor200ResponseLoadBalancerMonitorLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorInterval

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorInterval() int64`

GetMonitorInterval returns the MonitorInterval field if non-nil, zero value otherwise.

### GetMonitorIntervalOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorIntervalOk() (*int64, bool)`

GetMonitorIntervalOk returns a tuple with the MonitorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorInterval

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorInterval(v int64)`

SetMonitorInterval sets MonitorInterval field to given value.

### HasMonitorInterval

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorInterval() bool`

HasMonitorInterval returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetSendData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetSendData() string`

GetSendData returns the SendData field if non-nil, zero value otherwise.

### GetSendDataOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetSendDataOk() (*string, bool)`

GetSendDataOk returns a tuple with the SendData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetSendData(v string)`

SetSendData sets SendData field to given value.

### HasSendData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasSendData() bool`

HasSendData returns a boolean if a field has been set.

### SetSendDataNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetSendDataNil(b bool)`

 SetSendDataNil sets the value for SendData to be an explicit nil

### UnsetSendData
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetSendData()`

UnsetSendData ensures that no value is present for SendData, not even an explicit nil
### GetSendVersion

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetSendVersion() string`

GetSendVersion returns the SendVersion field if non-nil, zero value otherwise.

### GetSendVersionOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetSendVersionOk() (*string, bool)`

GetSendVersionOk returns a tuple with the SendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendVersion

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetSendVersion(v string)`

SetSendVersion sets SendVersion field to given value.

### HasSendVersion

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasSendVersion() bool`

HasSendVersion returns a boolean if a field has been set.

### GetSendType

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### GetReceiveData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetReceiveData() string`

GetReceiveData returns the ReceiveData field if non-nil, zero value otherwise.

### GetReceiveDataOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetReceiveDataOk() (*string, bool)`

GetReceiveDataOk returns a tuple with the ReceiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetReceiveData(v string)`

SetReceiveData sets ReceiveData field to given value.

### HasReceiveData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasReceiveData() bool`

HasReceiveData returns a boolean if a field has been set.

### SetReceiveDataNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetReceiveDataNil(b bool)`

 SetReceiveDataNil sets the value for ReceiveData to be an explicit nil

### UnsetReceiveData
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetReceiveData()`

UnsetReceiveData ensures that no value is present for ReceiveData, not even an explicit nil
### GetReceiveCode

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### GetDisabledData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDisabledData() string`

GetDisabledData returns the DisabledData field if non-nil, zero value otherwise.

### GetDisabledDataOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDisabledDataOk() (*string, bool)`

GetDisabledDataOk returns a tuple with the DisabledData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetDisabledData(v string)`

SetDisabledData sets DisabledData field to given value.

### HasDisabledData

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasDisabledData() bool`

HasDisabledData returns a boolean if a field has been set.

### SetDisabledDataNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetDisabledDataNil(b bool)`

 SetDisabledDataNil sets the value for DisabledData to be an explicit nil

### UnsetDisabledData
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetDisabledData()`

UnsetDisabledData ensures that no value is present for DisabledData, not even an explicit nil
### GetMonitorUsername

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorUsername() string`

GetMonitorUsername returns the MonitorUsername field if non-nil, zero value otherwise.

### GetMonitorUsernameOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorUsernameOk() (*string, bool)`

GetMonitorUsernameOk returns a tuple with the MonitorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsername

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorUsername(v string)`

SetMonitorUsername sets MonitorUsername field to given value.

### HasMonitorUsername

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorUsername() bool`

HasMonitorUsername returns a boolean if a field has been set.

### SetMonitorUsernameNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorUsernameNil(b bool)`

 SetMonitorUsernameNil sets the value for MonitorUsername to be an explicit nil

### UnsetMonitorUsername
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetMonitorUsername()`

UnsetMonitorUsername ensures that no value is present for MonitorUsername, not even an explicit nil
### GetMonitorPassword

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorPassword() string`

GetMonitorPassword returns the MonitorPassword field if non-nil, zero value otherwise.

### GetMonitorPasswordOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorPasswordOk() (*string, bool)`

GetMonitorPasswordOk returns a tuple with the MonitorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPassword

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorPassword(v string)`

SetMonitorPassword sets MonitorPassword field to given value.

### HasMonitorPassword

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorPassword() bool`

HasMonitorPassword returns a boolean if a field has been set.

### SetMonitorPasswordNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorPasswordNil(b bool)`

 SetMonitorPasswordNil sets the value for MonitorPassword to be an explicit nil

### UnsetMonitorPassword
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetMonitorPassword()`

UnsetMonitorPassword ensures that no value is present for MonitorPassword, not even an explicit nil
### GetMonitorPasswordHash

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorPasswordHash() string`

GetMonitorPasswordHash returns the MonitorPasswordHash field if non-nil, zero value otherwise.

### GetMonitorPasswordHashOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorPasswordHashOk() (*string, bool)`

GetMonitorPasswordHashOk returns a tuple with the MonitorPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPasswordHash

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorPasswordHash(v string)`

SetMonitorPasswordHash sets MonitorPasswordHash field to given value.

### HasMonitorPasswordHash

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorPasswordHash() bool`

HasMonitorPasswordHash returns a boolean if a field has been set.

### SetMonitorPasswordHashNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorPasswordHashNil(b bool)`

 SetMonitorPasswordHashNil sets the value for MonitorPasswordHash to be an explicit nil

### UnsetMonitorPasswordHash
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetMonitorPasswordHash()`

UnsetMonitorPasswordHash ensures that no value is present for MonitorPasswordHash, not even an explicit nil
### GetMonitorDestination

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorDestination() string`

GetMonitorDestination returns the MonitorDestination field if non-nil, zero value otherwise.

### GetMonitorDestinationOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorDestinationOk() (*string, bool)`

GetMonitorDestinationOk returns a tuple with the MonitorDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDestination

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorDestination(v string)`

SetMonitorDestination sets MonitorDestination field to given value.

### HasMonitorDestination

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorDestination() bool`

HasMonitorDestination returns a boolean if a field has been set.

### GetMonitorReverse

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorReverse() bool`

GetMonitorReverse returns the MonitorReverse field if non-nil, zero value otherwise.

### GetMonitorReverseOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorReverseOk() (*bool, bool)`

GetMonitorReverseOk returns a tuple with the MonitorReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorReverse

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorReverse(v bool)`

SetMonitorReverse sets MonitorReverse field to given value.

### HasMonitorReverse

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorReverse() bool`

HasMonitorReverse returns a boolean if a field has been set.

### GetMonitorTransparent

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorTransparent() bool`

GetMonitorTransparent returns the MonitorTransparent field if non-nil, zero value otherwise.

### GetMonitorTransparentOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorTransparentOk() (*bool, bool)`

GetMonitorTransparentOk returns a tuple with the MonitorTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTransparent

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorTransparent(v bool)`

SetMonitorTransparent sets MonitorTransparent field to given value.

### HasMonitorTransparent

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorTransparent() bool`

HasMonitorTransparent returns a boolean if a field has been set.

### GetMonitorAdaptive

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorAdaptive() bool`

GetMonitorAdaptive returns the MonitorAdaptive field if non-nil, zero value otherwise.

### GetMonitorAdaptiveOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorAdaptiveOk() (*bool, bool)`

GetMonitorAdaptiveOk returns a tuple with the MonitorAdaptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAdaptive

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorAdaptive(v bool)`

SetMonitorAdaptive sets MonitorAdaptive field to given value.

### HasMonitorAdaptive

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorAdaptive() bool`

HasMonitorAdaptive returns a boolean if a field has been set.

### GetAliasAddress

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetAliasAddress() string`

GetAliasAddress returns the AliasAddress field if non-nil, zero value otherwise.

### GetAliasAddressOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetAliasAddressOk() (*string, bool)`

GetAliasAddressOk returns a tuple with the AliasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasAddress

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetAliasAddress(v string)`

SetAliasAddress sets AliasAddress field to given value.

### HasAliasAddress

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasAliasAddress() bool`

HasAliasAddress returns a boolean if a field has been set.

### SetAliasAddressNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetAliasAddressNil(b bool)`

 SetAliasAddressNil sets the value for AliasAddress to be an explicit nil

### UnsetAliasAddress
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetAliasAddress()`

UnsetAliasAddress ensures that no value is present for AliasAddress, not even an explicit nil
### GetAliasPort

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetAliasPort() int64`

GetAliasPort returns the AliasPort field if non-nil, zero value otherwise.

### GetAliasPortOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetAliasPortOk() (*int64, bool)`

GetAliasPortOk returns a tuple with the AliasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPort

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetAliasPort(v int64)`

SetAliasPort sets AliasPort field to given value.

### HasAliasPort

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasAliasPort() bool`

HasAliasPort returns a boolean if a field has been set.

### GetInternalId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMonitorSource

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorSource() string`

GetMonitorSource returns the MonitorSource field if non-nil, zero value otherwise.

### GetMonitorSourceOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMonitorSourceOk() (*string, bool)`

GetMonitorSourceOk returns a tuple with the MonitorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSource

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMonitorSource(v string)`

SetMonitorSource sets MonitorSource field to given value.

### HasMonitorSource

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMonitorSource() bool`

HasMonitorSource returns a boolean if a field has been set.

### GetStatus

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetEnabled

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxRetry

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.

### HasMaxRetry

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasMaxRetry() bool`

HasMaxRetry returns a boolean if a field has been set.

### GetFallCount

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetFallCount() int64`

GetFallCount returns the FallCount field if non-nil, zero value otherwise.

### GetFallCountOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetFallCountOk() (*int64, bool)`

GetFallCountOk returns a tuple with the FallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallCount

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetFallCount(v int64)`

SetFallCount sets FallCount field to given value.

### HasFallCount

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasFallCount() bool`

HasFallCount returns a boolean if a field has been set.

### GetRiseCount

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetRiseCount() int64`

GetRiseCount returns the RiseCount field if non-nil, zero value otherwise.

### GetRiseCountOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetRiseCountOk() (*int64, bool)`

GetRiseCountOk returns a tuple with the RiseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseCount

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetRiseCount(v int64)`

SetRiseCount sets RiseCount field to given value.

### HasRiseCount

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasRiseCount() bool`

HasRiseCount returns a boolean if a field has been set.

### GetDataLength

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDataLength() int64`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDataLengthOk() (*int64, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetDataLength(v int64)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### SetDataLengthNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetDataLengthNil(b bool)`

 SetDataLengthNil sets the value for DataLength to be an explicit nil

### UnsetDataLength
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) UnsetDataLength()`

UnsetDataLength ensures that no value is present for DataLength, not even an explicit nil
### GetConfig

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetCreatedBy() GetLoadBalancerMonitor200ResponseLoadBalancerMonitorCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetCreatedByOk() (*GetLoadBalancerMonitor200ResponseLoadBalancerMonitorCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetCreatedBy(v GetLoadBalancerMonitor200ResponseLoadBalancerMonitorCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


