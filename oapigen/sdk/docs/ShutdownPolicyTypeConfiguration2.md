# ShutdownPolicyTypeConfiguration2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShutdownType** | **string** |  | 
**ShutdownAge** | Pointer to **string** |  | [optional] 
**ShutdownRenewal** | Pointer to **string** |  | [optional] 
**ShutdownNotify** | Pointer to **string** |  | [optional] 
**ShutdownMessage** | Pointer to **string** |  | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** |  | [optional] 
**AccountIntegrationId** | Pointer to **string** |  | [optional] 
**ShutdownHideFixed** | Pointer to **bool** |  | [optional] 
**ShutdownWorkflowId** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 

## Methods

### NewShutdownPolicyTypeConfiguration2

`func NewShutdownPolicyTypeConfiguration2(shutdownType string, ) *ShutdownPolicyTypeConfiguration2`

NewShutdownPolicyTypeConfiguration2 instantiates a new ShutdownPolicyTypeConfiguration2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownPolicyTypeConfiguration2WithDefaults

`func NewShutdownPolicyTypeConfiguration2WithDefaults() *ShutdownPolicyTypeConfiguration2`

NewShutdownPolicyTypeConfiguration2WithDefaults instantiates a new ShutdownPolicyTypeConfiguration2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShutdownType

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.


### GetShutdownAge

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetAccountIntegrationId

`func (o *ShutdownPolicyTypeConfiguration2) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *ShutdownPolicyTypeConfiguration2) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *ShutdownPolicyTypeConfiguration2) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *ShutdownPolicyTypeConfiguration2) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetShutdownWorkflowId

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownWorkflowId() string`

GetShutdownWorkflowId returns the ShutdownWorkflowId field if non-nil, zero value otherwise.

### GetShutdownWorkflowIdOk

`func (o *ShutdownPolicyTypeConfiguration2) GetShutdownWorkflowIdOk() (*string, bool)`

GetShutdownWorkflowIdOk returns a tuple with the ShutdownWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWorkflowId

`func (o *ShutdownPolicyTypeConfiguration2) SetShutdownWorkflowId(v string)`

SetShutdownWorkflowId sets ShutdownWorkflowId field to given value.

### HasShutdownWorkflowId

`func (o *ShutdownPolicyTypeConfiguration2) HasShutdownWorkflowId() bool`

HasShutdownWorkflowId returns a boolean if a field has been set.

### GetFlowId

`func (o *ShutdownPolicyTypeConfiguration2) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *ShutdownPolicyTypeConfiguration2) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *ShutdownPolicyTypeConfiguration2) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *ShutdownPolicyTypeConfiguration2) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *ShutdownPolicyTypeConfiguration2) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *ShutdownPolicyTypeConfiguration2) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *ShutdownPolicyTypeConfiguration2) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *ShutdownPolicyTypeConfiguration2) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


