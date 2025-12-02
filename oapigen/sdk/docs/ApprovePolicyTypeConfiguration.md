# ApprovePolicyTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | **string** | ID of your ServiceNow or approval integration | 
**WorkflowId** | Pointer to **string** | ID of legacy ServiceNow workflow (set if workflowType is &#39;workflow&#39;) | [optional] 
**FlowId** | Pointer to **string** | ID of ServiceNow Flow (set if workflowType is &#39;flow&#39;) | [optional] 
**WorkflowType** | Pointer to **string** | Options: \&quot;workflow\&quot; (legacy workflow), \&quot;flow\&quot; (ServiceNow Flow) | [optional] 

## Methods

### NewApprovePolicyTypeConfiguration

`func NewApprovePolicyTypeConfiguration(accountIntegrationId string, ) *ApprovePolicyTypeConfiguration`

NewApprovePolicyTypeConfiguration instantiates a new ApprovePolicyTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovePolicyTypeConfigurationWithDefaults

`func NewApprovePolicyTypeConfigurationWithDefaults() *ApprovePolicyTypeConfiguration`

NewApprovePolicyTypeConfigurationWithDefaults instantiates a new ApprovePolicyTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *ApprovePolicyTypeConfiguration) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *ApprovePolicyTypeConfiguration) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *ApprovePolicyTypeConfiguration) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.


### GetWorkflowId

`func (o *ApprovePolicyTypeConfiguration) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ApprovePolicyTypeConfiguration) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ApprovePolicyTypeConfiguration) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ApprovePolicyTypeConfiguration) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetFlowId

`func (o *ApprovePolicyTypeConfiguration) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *ApprovePolicyTypeConfiguration) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *ApprovePolicyTypeConfiguration) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *ApprovePolicyTypeConfiguration) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *ApprovePolicyTypeConfiguration) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *ApprovePolicyTypeConfiguration) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *ApprovePolicyTypeConfiguration) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *ApprovePolicyTypeConfiguration) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


