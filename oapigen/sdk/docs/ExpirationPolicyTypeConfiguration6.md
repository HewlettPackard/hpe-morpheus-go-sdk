# ExpirationPolicyTypeConfiguration6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifecycleType** | **string** |  | 
**LifecycleAge** | Pointer to **string** |  | [optional] 
**LifecycleRenewal** | Pointer to **string** |  | [optional] 
**LifecycleNotify** | Pointer to **string** |  | [optional] 
**LifecycleMessage** | Pointer to **string** |  | [optional] 
**LifecycleAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**LifecycleExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**AccountIntegrationId** | Pointer to **string** |  | [optional] 
**LifecycleWorkflowId** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**LifecycleHideFixed** | Pointer to **bool** |  | [optional] 

## Methods

### NewExpirationPolicyTypeConfiguration6

`func NewExpirationPolicyTypeConfiguration6(lifecycleType string, ) *ExpirationPolicyTypeConfiguration6`

NewExpirationPolicyTypeConfiguration6 instantiates a new ExpirationPolicyTypeConfiguration6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationPolicyTypeConfiguration6WithDefaults

`func NewExpirationPolicyTypeConfiguration6WithDefaults() *ExpirationPolicyTypeConfiguration6`

NewExpirationPolicyTypeConfiguration6WithDefaults instantiates a new ExpirationPolicyTypeConfiguration6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifecycleType

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleType() string`

GetLifecycleType returns the LifecycleType field if non-nil, zero value otherwise.

### GetLifecycleTypeOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleTypeOk() (*string, bool)`

GetLifecycleTypeOk returns a tuple with the LifecycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleType

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleType(v string)`

SetLifecycleType sets LifecycleType field to given value.


### GetLifecycleAge

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleAge() string`

GetLifecycleAge returns the LifecycleAge field if non-nil, zero value otherwise.

### GetLifecycleAgeOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleAgeOk() (*string, bool)`

GetLifecycleAgeOk returns a tuple with the LifecycleAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAge

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleAge(v string)`

SetLifecycleAge sets LifecycleAge field to given value.

### HasLifecycleAge

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleAge() bool`

HasLifecycleAge returns a boolean if a field has been set.

### GetLifecycleRenewal

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleRenewal() string`

GetLifecycleRenewal returns the LifecycleRenewal field if non-nil, zero value otherwise.

### GetLifecycleRenewalOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleRenewalOk() (*string, bool)`

GetLifecycleRenewalOk returns a tuple with the LifecycleRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRenewal

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleRenewal(v string)`

SetLifecycleRenewal sets LifecycleRenewal field to given value.

### HasLifecycleRenewal

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleRenewal() bool`

HasLifecycleRenewal returns a boolean if a field has been set.

### GetLifecycleNotify

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleNotify() string`

GetLifecycleNotify returns the LifecycleNotify field if non-nil, zero value otherwise.

### GetLifecycleNotifyOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleNotifyOk() (*string, bool)`

GetLifecycleNotifyOk returns a tuple with the LifecycleNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotify

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleNotify(v string)`

SetLifecycleNotify sets LifecycleNotify field to given value.

### HasLifecycleNotify

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleNotify() bool`

HasLifecycleNotify returns a boolean if a field has been set.

### GetLifecycleMessage

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleMessage() string`

GetLifecycleMessage returns the LifecycleMessage field if non-nil, zero value otherwise.

### GetLifecycleMessageOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleMessageOk() (*string, bool)`

GetLifecycleMessageOk returns a tuple with the LifecycleMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleMessage

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleMessage(v string)`

SetLifecycleMessage sets LifecycleMessage field to given value.

### HasLifecycleMessage

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleMessage() bool`

HasLifecycleMessage returns a boolean if a field has been set.

### GetLifecycleAutoRenew

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleAutoRenew() string`

GetLifecycleAutoRenew returns the LifecycleAutoRenew field if non-nil, zero value otherwise.

### GetLifecycleAutoRenewOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleAutoRenewOk() (*string, bool)`

GetLifecycleAutoRenewOk returns a tuple with the LifecycleAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAutoRenew

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleAutoRenew(v string)`

SetLifecycleAutoRenew sets LifecycleAutoRenew field to given value.

### HasLifecycleAutoRenew

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleAutoRenew() bool`

HasLifecycleAutoRenew returns a boolean if a field has been set.

### GetLifecycleAllowExtend

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleAllowExtend() string`

GetLifecycleAllowExtend returns the LifecycleAllowExtend field if non-nil, zero value otherwise.

### GetLifecycleAllowExtendOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleAllowExtendOk() (*string, bool)`

GetLifecycleAllowExtendOk returns a tuple with the LifecycleAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAllowExtend

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleAllowExtend(v string)`

SetLifecycleAllowExtend sets LifecycleAllowExtend field to given value.

### HasLifecycleAllowExtend

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleAllowExtend() bool`

HasLifecycleAllowExtend returns a boolean if a field has been set.

### GetLifecycleExtensionsBeforeApproval

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleExtensionsBeforeApproval() string`

GetLifecycleExtensionsBeforeApproval returns the LifecycleExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetLifecycleExtensionsBeforeApprovalOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleExtensionsBeforeApprovalOk() (*string, bool)`

GetLifecycleExtensionsBeforeApprovalOk returns a tuple with the LifecycleExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExtensionsBeforeApproval

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleExtensionsBeforeApproval(v string)`

SetLifecycleExtensionsBeforeApproval sets LifecycleExtensionsBeforeApproval field to given value.

### HasLifecycleExtensionsBeforeApproval

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleExtensionsBeforeApproval() bool`

HasLifecycleExtensionsBeforeApproval returns a boolean if a field has been set.

### GetAccountIntegrationId

`func (o *ExpirationPolicyTypeConfiguration6) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *ExpirationPolicyTypeConfiguration6) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *ExpirationPolicyTypeConfiguration6) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *ExpirationPolicyTypeConfiguration6) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetLifecycleWorkflowId

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleWorkflowId() string`

GetLifecycleWorkflowId returns the LifecycleWorkflowId field if non-nil, zero value otherwise.

### GetLifecycleWorkflowIdOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleWorkflowIdOk() (*string, bool)`

GetLifecycleWorkflowIdOk returns a tuple with the LifecycleWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleWorkflowId

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleWorkflowId(v string)`

SetLifecycleWorkflowId sets LifecycleWorkflowId field to given value.

### HasLifecycleWorkflowId

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleWorkflowId() bool`

HasLifecycleWorkflowId returns a boolean if a field has been set.

### GetFlowId

`func (o *ExpirationPolicyTypeConfiguration6) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *ExpirationPolicyTypeConfiguration6) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *ExpirationPolicyTypeConfiguration6) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *ExpirationPolicyTypeConfiguration6) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *ExpirationPolicyTypeConfiguration6) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *ExpirationPolicyTypeConfiguration6) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *ExpirationPolicyTypeConfiguration6) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *ExpirationPolicyTypeConfiguration6) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetLifecycleHideFixed

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleHideFixed() bool`

GetLifecycleHideFixed returns the LifecycleHideFixed field if non-nil, zero value otherwise.

### GetLifecycleHideFixedOk

`func (o *ExpirationPolicyTypeConfiguration6) GetLifecycleHideFixedOk() (*bool, bool)`

GetLifecycleHideFixedOk returns a tuple with the LifecycleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleHideFixed

`func (o *ExpirationPolicyTypeConfiguration6) SetLifecycleHideFixed(v bool)`

SetLifecycleHideFixed sets LifecycleHideFixed field to given value.

### HasLifecycleHideFixed

`func (o *ExpirationPolicyTypeConfiguration6) HasLifecycleHideFixed() bool`

HasLifecycleHideFixed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


