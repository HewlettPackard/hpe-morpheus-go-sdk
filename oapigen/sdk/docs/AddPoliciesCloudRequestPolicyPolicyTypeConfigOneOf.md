# AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShutdownType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict shutdown) | [optional] 
**ShutdownAge** | Pointer to **string** | Shutdown Age (days) | [optional] 
**ShutdownRenewal** | Pointer to **string** | If the instance is renewed, this is the number of day increments the shutdown date is increased by | [optional] 
**ShutdownNotify** | Pointer to **string** | Days before shutdown to notify via email | [optional] 
**ShutdownMessage** | Pointer to **string** | Notification message | [optional] 
**ShutdownAutoRenew** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownAllowExtend** | Pointer to **string** |  | [optional] [default to "off"]
**ShutdownExtensionsBeforeApproval** | Pointer to **string** | Number of extensions before requiring approval | [optional] 
**AccountIntegrationId** | Pointer to **string** | ID of your ServiceNow or approval integration | [optional] 
**ShutdownHideFixed** | Pointer to **bool** | Hide fixed shutdown from users | [optional] 
**ShutdownWorkflowId** | Pointer to **string** | ID of legacy ServiceNow workflow (set if workflowType is &#39;workflow&#39;) | [optional] 
**FlowId** | Pointer to **string** | ID of ServiceNow Flow (set if workflowType is &#39;flow&#39;) | [optional] 
**WorkflowType** | Pointer to **string** | Options: \&quot;workflow\&quot; (legacy workflow), \&quot;flow\&quot; (ServiceNow Flow) | [optional] 

## Methods

### NewAddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf

`func NewAddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf() *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf`

NewAddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf instantiates a new AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPoliciesCloudRequestPolicyPolicyTypeConfigOneOfWithDefaults

`func NewAddPoliciesCloudRequestPolicyPolicyTypeConfigOneOfWithDefaults() *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf`

NewAddPoliciesCloudRequestPolicyPolicyTypeConfigOneOfWithDefaults instantiates a new AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShutdownType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownType() string`

GetShutdownType returns the ShutdownType field if non-nil, zero value otherwise.

### GetShutdownTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownTypeOk() (*string, bool)`

GetShutdownTypeOk returns a tuple with the ShutdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownType(v string)`

SetShutdownType sets ShutdownType field to given value.

### HasShutdownType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownType() bool`

HasShutdownType returns a boolean if a field has been set.

### GetShutdownAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownAge() string`

GetShutdownAge returns the ShutdownAge field if non-nil, zero value otherwise.

### GetShutdownAgeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownAgeOk() (*string, bool)`

GetShutdownAgeOk returns a tuple with the ShutdownAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownAge(v string)`

SetShutdownAge sets ShutdownAge field to given value.

### HasShutdownAge

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownAge() bool`

HasShutdownAge returns a boolean if a field has been set.

### GetShutdownRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownRenewal() string`

GetShutdownRenewal returns the ShutdownRenewal field if non-nil, zero value otherwise.

### GetShutdownRenewalOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownRenewalOk() (*string, bool)`

GetShutdownRenewalOk returns a tuple with the ShutdownRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownRenewal(v string)`

SetShutdownRenewal sets ShutdownRenewal field to given value.

### HasShutdownRenewal

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownRenewal() bool`

HasShutdownRenewal returns a boolean if a field has been set.

### GetShutdownNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownNotify() string`

GetShutdownNotify returns the ShutdownNotify field if non-nil, zero value otherwise.

### GetShutdownNotifyOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownNotifyOk() (*string, bool)`

GetShutdownNotifyOk returns a tuple with the ShutdownNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownNotify(v string)`

SetShutdownNotify sets ShutdownNotify field to given value.

### HasShutdownNotify

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownNotify() bool`

HasShutdownNotify returns a boolean if a field has been set.

### GetShutdownMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownMessage() string`

GetShutdownMessage returns the ShutdownMessage field if non-nil, zero value otherwise.

### GetShutdownMessageOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownMessageOk() (*string, bool)`

GetShutdownMessageOk returns a tuple with the ShutdownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownMessage(v string)`

SetShutdownMessage sets ShutdownMessage field to given value.

### HasShutdownMessage

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownMessage() bool`

HasShutdownMessage returns a boolean if a field has been set.

### GetShutdownAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownAutoRenew() string`

GetShutdownAutoRenew returns the ShutdownAutoRenew field if non-nil, zero value otherwise.

### GetShutdownAutoRenewOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownAutoRenewOk() (*string, bool)`

GetShutdownAutoRenewOk returns a tuple with the ShutdownAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownAutoRenew(v string)`

SetShutdownAutoRenew sets ShutdownAutoRenew field to given value.

### HasShutdownAutoRenew

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownAutoRenew() bool`

HasShutdownAutoRenew returns a boolean if a field has been set.

### GetShutdownAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownAllowExtend() string`

GetShutdownAllowExtend returns the ShutdownAllowExtend field if non-nil, zero value otherwise.

### GetShutdownAllowExtendOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownAllowExtendOk() (*string, bool)`

GetShutdownAllowExtendOk returns a tuple with the ShutdownAllowExtend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownAllowExtend(v string)`

SetShutdownAllowExtend sets ShutdownAllowExtend field to given value.

### HasShutdownAllowExtend

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownAllowExtend() bool`

HasShutdownAllowExtend returns a boolean if a field has been set.

### GetShutdownExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownExtensionsBeforeApproval() string`

GetShutdownExtensionsBeforeApproval returns the ShutdownExtensionsBeforeApproval field if non-nil, zero value otherwise.

### GetShutdownExtensionsBeforeApprovalOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownExtensionsBeforeApprovalOk() (*string, bool)`

GetShutdownExtensionsBeforeApprovalOk returns a tuple with the ShutdownExtensionsBeforeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownExtensionsBeforeApproval(v string)`

SetShutdownExtensionsBeforeApproval sets ShutdownExtensionsBeforeApproval field to given value.

### HasShutdownExtensionsBeforeApproval

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownExtensionsBeforeApproval() bool`

HasShutdownExtensionsBeforeApproval returns a boolean if a field has been set.

### GetAccountIntegrationId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetShutdownHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownHideFixed() bool`

GetShutdownHideFixed returns the ShutdownHideFixed field if non-nil, zero value otherwise.

### GetShutdownHideFixedOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownHideFixedOk() (*bool, bool)`

GetShutdownHideFixedOk returns a tuple with the ShutdownHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownHideFixed(v bool)`

SetShutdownHideFixed sets ShutdownHideFixed field to given value.

### HasShutdownHideFixed

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownHideFixed() bool`

HasShutdownHideFixed returns a boolean if a field has been set.

### GetShutdownWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownWorkflowId() string`

GetShutdownWorkflowId returns the ShutdownWorkflowId field if non-nil, zero value otherwise.

### GetShutdownWorkflowIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetShutdownWorkflowIdOk() (*string, bool)`

GetShutdownWorkflowIdOk returns a tuple with the ShutdownWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetShutdownWorkflowId(v string)`

SetShutdownWorkflowId sets ShutdownWorkflowId field to given value.

### HasShutdownWorkflowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasShutdownWorkflowId() bool`

HasShutdownWorkflowId returns a boolean if a field has been set.

### GetFlowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *AddPoliciesCloudRequestPolicyPolicyTypeConfigOneOf) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


