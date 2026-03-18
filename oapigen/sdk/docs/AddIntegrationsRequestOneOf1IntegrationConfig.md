# AddIntegrationsRequestOneOf1IntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBranch** | Pointer to **string** | default branch | [optional] 
**AnsiblePlaybooks** | Pointer to **string** | Playbooks path | [optional] 
**AnsibleRoles** | Pointer to **string** | Roles path | [optional] 
**AnsibleGroupVars** | Pointer to **string** | Group variables path | [optional] 
**AnsibleHostVars** | Pointer to **string** | Host variables path | [optional] 
**AnsibleGalaxyEnabled** | Pointer to **bool** | Use Ansible Galaxy | [optional] 
**AnsibleVerbose** | Pointer to **bool** | Use verbose logging | [optional] 
**AnsibleCommandBus** | Pointer to **bool** | Use Morpheus Agent Command Bus | [optional] 
**CacheEnabled** | Pointer to **bool** | Enable Git repository caching | [optional] 

## Methods

### NewAddIntegrationsRequestOneOf1IntegrationConfig

`func NewAddIntegrationsRequestOneOf1IntegrationConfig() *AddIntegrationsRequestOneOf1IntegrationConfig`

NewAddIntegrationsRequestOneOf1IntegrationConfig instantiates a new AddIntegrationsRequestOneOf1IntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrationsRequestOneOf1IntegrationConfigWithDefaults

`func NewAddIntegrationsRequestOneOf1IntegrationConfigWithDefaults() *AddIntegrationsRequestOneOf1IntegrationConfig`

NewAddIntegrationsRequestOneOf1IntegrationConfigWithDefaults instantiates a new AddIntegrationsRequestOneOf1IntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranch

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetAnsiblePlaybooks

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsiblePlaybooks() string`

GetAnsiblePlaybooks returns the AnsiblePlaybooks field if non-nil, zero value otherwise.

### GetAnsiblePlaybooksOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsiblePlaybooksOk() (*string, bool)`

GetAnsiblePlaybooksOk returns a tuple with the AnsiblePlaybooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybooks

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsiblePlaybooks(v string)`

SetAnsiblePlaybooks sets AnsiblePlaybooks field to given value.

### HasAnsiblePlaybooks

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsiblePlaybooks() bool`

HasAnsiblePlaybooks returns a boolean if a field has been set.

### GetAnsibleRoles

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleRoles() string`

GetAnsibleRoles returns the AnsibleRoles field if non-nil, zero value otherwise.

### GetAnsibleRolesOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleRolesOk() (*string, bool)`

GetAnsibleRolesOk returns a tuple with the AnsibleRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleRoles

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsibleRoles(v string)`

SetAnsibleRoles sets AnsibleRoles field to given value.

### HasAnsibleRoles

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsibleRoles() bool`

HasAnsibleRoles returns a boolean if a field has been set.

### GetAnsibleGroupVars

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleGroupVars() string`

GetAnsibleGroupVars returns the AnsibleGroupVars field if non-nil, zero value otherwise.

### GetAnsibleGroupVarsOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleGroupVarsOk() (*string, bool)`

GetAnsibleGroupVarsOk returns a tuple with the AnsibleGroupVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroupVars

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsibleGroupVars(v string)`

SetAnsibleGroupVars sets AnsibleGroupVars field to given value.

### HasAnsibleGroupVars

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsibleGroupVars() bool`

HasAnsibleGroupVars returns a boolean if a field has been set.

### GetAnsibleHostVars

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleHostVars() string`

GetAnsibleHostVars returns the AnsibleHostVars field if non-nil, zero value otherwise.

### GetAnsibleHostVarsOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleHostVarsOk() (*string, bool)`

GetAnsibleHostVarsOk returns a tuple with the AnsibleHostVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHostVars

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsibleHostVars(v string)`

SetAnsibleHostVars sets AnsibleHostVars field to given value.

### HasAnsibleHostVars

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsibleHostVars() bool`

HasAnsibleHostVars returns a boolean if a field has been set.

### GetAnsibleGalaxyEnabled

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleGalaxyEnabled() bool`

GetAnsibleGalaxyEnabled returns the AnsibleGalaxyEnabled field if non-nil, zero value otherwise.

### GetAnsibleGalaxyEnabledOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleGalaxyEnabledOk() (*bool, bool)`

GetAnsibleGalaxyEnabledOk returns a tuple with the AnsibleGalaxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGalaxyEnabled

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsibleGalaxyEnabled(v bool)`

SetAnsibleGalaxyEnabled sets AnsibleGalaxyEnabled field to given value.

### HasAnsibleGalaxyEnabled

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsibleGalaxyEnabled() bool`

HasAnsibleGalaxyEnabled returns a boolean if a field has been set.

### GetAnsibleVerbose

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleVerbose() bool`

GetAnsibleVerbose returns the AnsibleVerbose field if non-nil, zero value otherwise.

### GetAnsibleVerboseOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleVerboseOk() (*bool, bool)`

GetAnsibleVerboseOk returns a tuple with the AnsibleVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleVerbose

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsibleVerbose(v bool)`

SetAnsibleVerbose sets AnsibleVerbose field to given value.

### HasAnsibleVerbose

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsibleVerbose() bool`

HasAnsibleVerbose returns a boolean if a field has been set.

### GetAnsibleCommandBus

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleCommandBus() bool`

GetAnsibleCommandBus returns the AnsibleCommandBus field if non-nil, zero value otherwise.

### GetAnsibleCommandBusOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetAnsibleCommandBusOk() (*bool, bool)`

GetAnsibleCommandBusOk returns a tuple with the AnsibleCommandBus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleCommandBus

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetAnsibleCommandBus(v bool)`

SetAnsibleCommandBus sets AnsibleCommandBus field to given value.

### HasAnsibleCommandBus

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasAnsibleCommandBus() bool`

HasAnsibleCommandBus returns a boolean if a field has been set.

### GetCacheEnabled

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetCacheEnabled() bool`

GetCacheEnabled returns the CacheEnabled field if non-nil, zero value otherwise.

### GetCacheEnabledOk

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) GetCacheEnabledOk() (*bool, bool)`

GetCacheEnabledOk returns a tuple with the CacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheEnabled

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) SetCacheEnabled(v bool)`

SetCacheEnabled sets CacheEnabled field to given value.

### HasCacheEnabled

`func (o *AddIntegrationsRequestOneOf1IntegrationConfig) HasCacheEnabled() bool`

HasCacheEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


