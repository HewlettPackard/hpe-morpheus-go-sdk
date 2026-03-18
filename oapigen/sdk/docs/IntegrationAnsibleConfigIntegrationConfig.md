# IntegrationAnsibleConfigIntegrationConfig

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

### NewIntegrationAnsibleConfigIntegrationConfig

`func NewIntegrationAnsibleConfigIntegrationConfig() *IntegrationAnsibleConfigIntegrationConfig`

NewIntegrationAnsibleConfigIntegrationConfig instantiates a new IntegrationAnsibleConfigIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationAnsibleConfigIntegrationConfigWithDefaults

`func NewIntegrationAnsibleConfigIntegrationConfigWithDefaults() *IntegrationAnsibleConfigIntegrationConfig`

NewIntegrationAnsibleConfigIntegrationConfigWithDefaults instantiates a new IntegrationAnsibleConfigIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranch

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetAnsiblePlaybooks

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsiblePlaybooks() string`

GetAnsiblePlaybooks returns the AnsiblePlaybooks field if non-nil, zero value otherwise.

### GetAnsiblePlaybooksOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsiblePlaybooksOk() (*string, bool)`

GetAnsiblePlaybooksOk returns a tuple with the AnsiblePlaybooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybooks

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsiblePlaybooks(v string)`

SetAnsiblePlaybooks sets AnsiblePlaybooks field to given value.

### HasAnsiblePlaybooks

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsiblePlaybooks() bool`

HasAnsiblePlaybooks returns a boolean if a field has been set.

### GetAnsibleRoles

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleRoles() string`

GetAnsibleRoles returns the AnsibleRoles field if non-nil, zero value otherwise.

### GetAnsibleRolesOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleRolesOk() (*string, bool)`

GetAnsibleRolesOk returns a tuple with the AnsibleRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleRoles

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleRoles(v string)`

SetAnsibleRoles sets AnsibleRoles field to given value.

### HasAnsibleRoles

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleRoles() bool`

HasAnsibleRoles returns a boolean if a field has been set.

### GetAnsibleGroupVars

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGroupVars() string`

GetAnsibleGroupVars returns the AnsibleGroupVars field if non-nil, zero value otherwise.

### GetAnsibleGroupVarsOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGroupVarsOk() (*string, bool)`

GetAnsibleGroupVarsOk returns a tuple with the AnsibleGroupVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroupVars

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleGroupVars(v string)`

SetAnsibleGroupVars sets AnsibleGroupVars field to given value.

### HasAnsibleGroupVars

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleGroupVars() bool`

HasAnsibleGroupVars returns a boolean if a field has been set.

### GetAnsibleHostVars

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleHostVars() string`

GetAnsibleHostVars returns the AnsibleHostVars field if non-nil, zero value otherwise.

### GetAnsibleHostVarsOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleHostVarsOk() (*string, bool)`

GetAnsibleHostVarsOk returns a tuple with the AnsibleHostVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHostVars

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleHostVars(v string)`

SetAnsibleHostVars sets AnsibleHostVars field to given value.

### HasAnsibleHostVars

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleHostVars() bool`

HasAnsibleHostVars returns a boolean if a field has been set.

### GetAnsibleGalaxyEnabled

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGalaxyEnabled() bool`

GetAnsibleGalaxyEnabled returns the AnsibleGalaxyEnabled field if non-nil, zero value otherwise.

### GetAnsibleGalaxyEnabledOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleGalaxyEnabledOk() (*bool, bool)`

GetAnsibleGalaxyEnabledOk returns a tuple with the AnsibleGalaxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGalaxyEnabled

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleGalaxyEnabled(v bool)`

SetAnsibleGalaxyEnabled sets AnsibleGalaxyEnabled field to given value.

### HasAnsibleGalaxyEnabled

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleGalaxyEnabled() bool`

HasAnsibleGalaxyEnabled returns a boolean if a field has been set.

### GetAnsibleVerbose

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleVerbose() bool`

GetAnsibleVerbose returns the AnsibleVerbose field if non-nil, zero value otherwise.

### GetAnsibleVerboseOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleVerboseOk() (*bool, bool)`

GetAnsibleVerboseOk returns a tuple with the AnsibleVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleVerbose

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleVerbose(v bool)`

SetAnsibleVerbose sets AnsibleVerbose field to given value.

### HasAnsibleVerbose

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleVerbose() bool`

HasAnsibleVerbose returns a boolean if a field has been set.

### GetAnsibleCommandBus

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleCommandBus() bool`

GetAnsibleCommandBus returns the AnsibleCommandBus field if non-nil, zero value otherwise.

### GetAnsibleCommandBusOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetAnsibleCommandBusOk() (*bool, bool)`

GetAnsibleCommandBusOk returns a tuple with the AnsibleCommandBus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleCommandBus

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetAnsibleCommandBus(v bool)`

SetAnsibleCommandBus sets AnsibleCommandBus field to given value.

### HasAnsibleCommandBus

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasAnsibleCommandBus() bool`

HasAnsibleCommandBus returns a boolean if a field has been set.

### GetCacheEnabled

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetCacheEnabled() bool`

GetCacheEnabled returns the CacheEnabled field if non-nil, zero value otherwise.

### GetCacheEnabledOk

`func (o *IntegrationAnsibleConfigIntegrationConfig) GetCacheEnabledOk() (*bool, bool)`

GetCacheEnabledOk returns a tuple with the CacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheEnabled

`func (o *IntegrationAnsibleConfigIntegrationConfig) SetCacheEnabled(v bool)`

SetCacheEnabled sets CacheEnabled field to given value.

### HasCacheEnabled

`func (o *IntegrationAnsibleConfigIntegrationConfig) HasCacheEnabled() bool`

HasCacheEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


