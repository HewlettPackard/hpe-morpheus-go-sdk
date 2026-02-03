# AddTasks200ResponseAllOfTaskTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnsibleOptions** | Pointer to **string** |  | [optional] 
**AnsiblePlaybook** | Pointer to **string** |  | [optional] 
**SshKey** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**LocalScriptGitRef** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**LocalScriptGitId** | Pointer to **string** |  | [optional] 
**AnsibleGitId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**AnsibleSkipTags** | Pointer to **string** |  | [optional] 
**AnsibleTags** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**AnsibleGitRef** | Pointer to **string** |  | [optional] 
**AnsibleTowerGitRef** | Pointer to **string** |  | [optional] 
**AnsibleGroup** | Pointer to **string** |  | [optional] 
**AnsibleTowerExecuteMode** | Pointer to **string** |  | [optional] 
**ChefDataKey** | Pointer to **string** |  | [optional] 
**ChefDataKeyHash** | Pointer to **string** |  | [optional] 
**ChefRunList** | Pointer to **string** |  | [optional] 
**ChefDataKeyPath** | Pointer to **string** |  | [optional] 
**ChefEnv** | Pointer to **string** |  | [optional] 
**ChefNodeName** | Pointer to **string** |  | [optional] 
**ChefAttributes** | Pointer to **string** |  | [optional] 
**ConditionalScript** | Pointer to **string** | Allows the user to set JavaScript logic. If it resolves to true, HPE Morpheus Enterprise will run the Operational Workflow set as the “IF OPERATIONAL WORKFLOW” and if it resolves to false, HPE Morpheus Enterprise will run the “ELSE OPERATIONAL WORKFLOW” | [optional] 
**IfOperationalWorkflowId** | Pointer to **int64** | If Operational Workflow ID | [optional] 
**IfOperationalWorkflowName** | Pointer to **string** | If Operational Workflow Name | [optional] 
**ElseOperationalWorkflowId** | Pointer to **int64** | else Operational Workflow ID | [optional] 
**ElseOperationalWorkflowName** | Pointer to **string** | Else Operational Workflow Name | [optional] 
**EmailSkipTemplate** | Pointer to **string** |  | [optional] 
**EmailSubject** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**WebPassword** | Pointer to **string** |  | [optional] 
**WebPasswordHash** | Pointer to **string** |  | [optional] 
**WebUser** | Pointer to **string** |  | [optional] 
**WebBody** | Pointer to **string** |  | [optional] 
**WebHeaders** | Pointer to **string** |  | [optional] 
**IgnoreSSL** | Pointer to **string** |  | [optional] 
**WebMethod** | Pointer to **string** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**JsScript** | Pointer to **string** |  | [optional] 
**ContainerScriptId** | Pointer to **string** |  | [optional] 
**ContainerScript** | Pointer to **string** |  | [optional] 
**ContainerTemplateId** | Pointer to **string** |  | [optional] 
**ContainerTemplate** | Pointer to **string** |  | [optional] 
**OperationalWorkflowId** | **string** | Operational Workflow ID | 
**OperationalWorkflowName** | Pointer to **string** | Operational Workflow Name | [optional] 
**PuppetEnvironment** | Pointer to **string** |  | [optional] 
**PuppetNodeName** | Pointer to **string** |  | [optional] 
**PythonArgs** | Pointer to **string** |  | [optional] 
**PythonBinary** | Pointer to **string** |  | [optional] 
**PythonAdditionalPackages** | Pointer to **string** |  | [optional] 
**VroBody** | Pointer to **string** |  | [optional] 
**WriteAttributesAttributes** | Pointer to **string** |  | [optional] 

## Methods

### NewAddTasks200ResponseAllOfTaskTaskOptions

`func NewAddTasks200ResponseAllOfTaskTaskOptions(operationalWorkflowId string, ) *AddTasks200ResponseAllOfTaskTaskOptions`

NewAddTasks200ResponseAllOfTaskTaskOptions instantiates a new AddTasks200ResponseAllOfTaskTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTasks200ResponseAllOfTaskTaskOptionsWithDefaults

`func NewAddTasks200ResponseAllOfTaskTaskOptionsWithDefaults() *AddTasks200ResponseAllOfTaskTaskOptions`

NewAddTasks200ResponseAllOfTaskTaskOptionsWithDefaults instantiates a new AddTasks200ResponseAllOfTaskTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleOptions

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleOptions() string`

GetAnsibleOptions returns the AnsibleOptions field if non-nil, zero value otherwise.

### GetAnsibleOptionsOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleOptionsOk() (*string, bool)`

GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleOptions

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleOptions(v string)`

SetAnsibleOptions sets AnsibleOptions field to given value.

### HasAnsibleOptions

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleOptions() bool`

HasAnsibleOptions returns a boolean if a field has been set.

### GetAnsiblePlaybook

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsiblePlaybook() string`

GetAnsiblePlaybook returns the AnsiblePlaybook field if non-nil, zero value otherwise.

### GetAnsiblePlaybookOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsiblePlaybookOk() (*string, bool)`

GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybook

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsiblePlaybook(v string)`

SetAnsiblePlaybook sets AnsiblePlaybook field to given value.

### HasAnsiblePlaybook

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsiblePlaybook() bool`

HasAnsiblePlaybook returns a boolean if a field has been set.

### GetSshKey

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetPort

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLocalScriptGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### GetPassword

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetLocalScriptGitId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### GetAnsibleGitId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleGitId() string`

GetAnsibleGitId returns the AnsibleGitId field if non-nil, zero value otherwise.

### GetAnsibleGitIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleGitIdOk() (*string, bool)`

GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleGitId(v string)`

SetAnsibleGitId sets AnsibleGitId field to given value.

### HasAnsibleGitId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleGitId() bool`

HasAnsibleGitId returns a boolean if a field has been set.

### GetHost

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAnsibleSkipTags

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleSkipTags() string`

GetAnsibleSkipTags returns the AnsibleSkipTags field if non-nil, zero value otherwise.

### GetAnsibleSkipTagsOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleSkipTagsOk() (*string, bool)`

GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleSkipTags

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleSkipTags(v string)`

SetAnsibleSkipTags sets AnsibleSkipTags field to given value.

### HasAnsibleSkipTags

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleSkipTags() bool`

HasAnsibleSkipTags returns a boolean if a field has been set.

### GetAnsibleTags

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleTags() string`

GetAnsibleTags returns the AnsibleTags field if non-nil, zero value otherwise.

### GetAnsibleTagsOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleTagsOk() (*string, bool)`

GetAnsibleTagsOk returns a tuple with the AnsibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTags

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleTags(v string)`

SetAnsibleTags sets AnsibleTags field to given value.

### HasAnsibleTags

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleTags() bool`

HasAnsibleTags returns a boolean if a field has been set.

### GetUsername

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAnsibleGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleGitRef() string`

GetAnsibleGitRef returns the AnsibleGitRef field if non-nil, zero value otherwise.

### GetAnsibleGitRefOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleGitRefOk() (*string, bool)`

GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleGitRef(v string)`

SetAnsibleGitRef sets AnsibleGitRef field to given value.

### HasAnsibleGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleGitRef() bool`

HasAnsibleGitRef returns a boolean if a field has been set.

### GetAnsibleTowerGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleTowerGitRef() string`

GetAnsibleTowerGitRef returns the AnsibleTowerGitRef field if non-nil, zero value otherwise.

### GetAnsibleTowerGitRefOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleTowerGitRefOk() (*string, bool)`

GetAnsibleTowerGitRefOk returns a tuple with the AnsibleTowerGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleTowerGitRef(v string)`

SetAnsibleTowerGitRef sets AnsibleTowerGitRef field to given value.

### HasAnsibleTowerGitRef

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleTowerGitRef() bool`

HasAnsibleTowerGitRef returns a boolean if a field has been set.

### GetAnsibleGroup

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleGroup() string`

GetAnsibleGroup returns the AnsibleGroup field if non-nil, zero value otherwise.

### GetAnsibleGroupOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleGroupOk() (*string, bool)`

GetAnsibleGroupOk returns a tuple with the AnsibleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroup

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleGroup(v string)`

SetAnsibleGroup sets AnsibleGroup field to given value.

### HasAnsibleGroup

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleGroup() bool`

HasAnsibleGroup returns a boolean if a field has been set.

### GetAnsibleTowerExecuteMode

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleTowerExecuteMode() string`

GetAnsibleTowerExecuteMode returns the AnsibleTowerExecuteMode field if non-nil, zero value otherwise.

### GetAnsibleTowerExecuteModeOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetAnsibleTowerExecuteModeOk() (*string, bool)`

GetAnsibleTowerExecuteModeOk returns a tuple with the AnsibleTowerExecuteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerExecuteMode

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetAnsibleTowerExecuteMode(v string)`

SetAnsibleTowerExecuteMode sets AnsibleTowerExecuteMode field to given value.

### HasAnsibleTowerExecuteMode

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasAnsibleTowerExecuteMode() bool`

HasAnsibleTowerExecuteMode returns a boolean if a field has been set.

### GetChefDataKey

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefDataKey() string`

GetChefDataKey returns the ChefDataKey field if non-nil, zero value otherwise.

### GetChefDataKeyOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefDataKeyOk() (*string, bool)`

GetChefDataKeyOk returns a tuple with the ChefDataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKey

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefDataKey(v string)`

SetChefDataKey sets ChefDataKey field to given value.

### HasChefDataKey

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefDataKey() bool`

HasChefDataKey returns a boolean if a field has been set.

### GetChefDataKeyHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefDataKeyHash() string`

GetChefDataKeyHash returns the ChefDataKeyHash field if non-nil, zero value otherwise.

### GetChefDataKeyHashOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefDataKeyHashOk() (*string, bool)`

GetChefDataKeyHashOk returns a tuple with the ChefDataKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefDataKeyHash(v string)`

SetChefDataKeyHash sets ChefDataKeyHash field to given value.

### HasChefDataKeyHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefDataKeyHash() bool`

HasChefDataKeyHash returns a boolean if a field has been set.

### GetChefRunList

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefRunList() string`

GetChefRunList returns the ChefRunList field if non-nil, zero value otherwise.

### GetChefRunListOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefRunListOk() (*string, bool)`

GetChefRunListOk returns a tuple with the ChefRunList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefRunList

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefRunList(v string)`

SetChefRunList sets ChefRunList field to given value.

### HasChefRunList

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefRunList() bool`

HasChefRunList returns a boolean if a field has been set.

### GetChefDataKeyPath

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefDataKeyPath() string`

GetChefDataKeyPath returns the ChefDataKeyPath field if non-nil, zero value otherwise.

### GetChefDataKeyPathOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefDataKeyPathOk() (*string, bool)`

GetChefDataKeyPathOk returns a tuple with the ChefDataKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyPath

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefDataKeyPath(v string)`

SetChefDataKeyPath sets ChefDataKeyPath field to given value.

### HasChefDataKeyPath

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefDataKeyPath() bool`

HasChefDataKeyPath returns a boolean if a field has been set.

### GetChefEnv

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefEnv() string`

GetChefEnv returns the ChefEnv field if non-nil, zero value otherwise.

### GetChefEnvOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefEnvOk() (*string, bool)`

GetChefEnvOk returns a tuple with the ChefEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefEnv

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefEnv(v string)`

SetChefEnv sets ChefEnv field to given value.

### HasChefEnv

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefEnv() bool`

HasChefEnv returns a boolean if a field has been set.

### GetChefNodeName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefNodeName() string`

GetChefNodeName returns the ChefNodeName field if non-nil, zero value otherwise.

### GetChefNodeNameOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefNodeNameOk() (*string, bool)`

GetChefNodeNameOk returns a tuple with the ChefNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefNodeName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefNodeName(v string)`

SetChefNodeName sets ChefNodeName field to given value.

### HasChefNodeName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefNodeName() bool`

HasChefNodeName returns a boolean if a field has been set.

### GetChefAttributes

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefAttributes() string`

GetChefAttributes returns the ChefAttributes field if non-nil, zero value otherwise.

### GetChefAttributesOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetChefAttributesOk() (*string, bool)`

GetChefAttributesOk returns a tuple with the ChefAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefAttributes

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetChefAttributes(v string)`

SetChefAttributes sets ChefAttributes field to given value.

### HasChefAttributes

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasChefAttributes() bool`

HasChefAttributes returns a boolean if a field has been set.

### GetConditionalScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetConditionalScript() string`

GetConditionalScript returns the ConditionalScript field if non-nil, zero value otherwise.

### GetConditionalScriptOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetConditionalScriptOk() (*string, bool)`

GetConditionalScriptOk returns a tuple with the ConditionalScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetConditionalScript(v string)`

SetConditionalScript sets ConditionalScript field to given value.

### HasConditionalScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasConditionalScript() bool`

HasConditionalScript returns a boolean if a field has been set.

### GetIfOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetIfOperationalWorkflowId() int64`

GetIfOperationalWorkflowId returns the IfOperationalWorkflowId field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetIfOperationalWorkflowIdOk() (*int64, bool)`

GetIfOperationalWorkflowIdOk returns a tuple with the IfOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetIfOperationalWorkflowId(v int64)`

SetIfOperationalWorkflowId sets IfOperationalWorkflowId field to given value.

### HasIfOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasIfOperationalWorkflowId() bool`

HasIfOperationalWorkflowId returns a boolean if a field has been set.

### GetIfOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetIfOperationalWorkflowName() string`

GetIfOperationalWorkflowName returns the IfOperationalWorkflowName field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowNameOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetIfOperationalWorkflowNameOk() (*string, bool)`

GetIfOperationalWorkflowNameOk returns a tuple with the IfOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetIfOperationalWorkflowName(v string)`

SetIfOperationalWorkflowName sets IfOperationalWorkflowName field to given value.

### HasIfOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasIfOperationalWorkflowName() bool`

HasIfOperationalWorkflowName returns a boolean if a field has been set.

### GetElseOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetElseOperationalWorkflowId() int64`

GetElseOperationalWorkflowId returns the ElseOperationalWorkflowId field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetElseOperationalWorkflowIdOk() (*int64, bool)`

GetElseOperationalWorkflowIdOk returns a tuple with the ElseOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetElseOperationalWorkflowId(v int64)`

SetElseOperationalWorkflowId sets ElseOperationalWorkflowId field to given value.

### HasElseOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasElseOperationalWorkflowId() bool`

HasElseOperationalWorkflowId returns a boolean if a field has been set.

### GetElseOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetElseOperationalWorkflowName() string`

GetElseOperationalWorkflowName returns the ElseOperationalWorkflowName field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowNameOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetElseOperationalWorkflowNameOk() (*string, bool)`

GetElseOperationalWorkflowNameOk returns a tuple with the ElseOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetElseOperationalWorkflowName(v string)`

SetElseOperationalWorkflowName sets ElseOperationalWorkflowName field to given value.

### HasElseOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasElseOperationalWorkflowName() bool`

HasElseOperationalWorkflowName returns a boolean if a field has been set.

### GetEmailSkipTemplate

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetEmailSkipTemplate() string`

GetEmailSkipTemplate returns the EmailSkipTemplate field if non-nil, zero value otherwise.

### GetEmailSkipTemplateOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetEmailSkipTemplateOk() (*string, bool)`

GetEmailSkipTemplateOk returns a tuple with the EmailSkipTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSkipTemplate

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetEmailSkipTemplate(v string)`

SetEmailSkipTemplate sets EmailSkipTemplate field to given value.

### HasEmailSkipTemplate

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasEmailSkipTemplate() bool`

HasEmailSkipTemplate returns a boolean if a field has been set.

### GetEmailSubject

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailAddress

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetWebPassword

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebPassword() string`

GetWebPassword returns the WebPassword field if non-nil, zero value otherwise.

### GetWebPasswordOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebPasswordOk() (*string, bool)`

GetWebPasswordOk returns a tuple with the WebPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPassword

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebPassword(v string)`

SetWebPassword sets WebPassword field to given value.

### HasWebPassword

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebPassword() bool`

HasWebPassword returns a boolean if a field has been set.

### GetWebPasswordHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebPasswordHash() string`

GetWebPasswordHash returns the WebPasswordHash field if non-nil, zero value otherwise.

### GetWebPasswordHashOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebPasswordHashOk() (*string, bool)`

GetWebPasswordHashOk returns a tuple with the WebPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPasswordHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebPasswordHash(v string)`

SetWebPasswordHash sets WebPasswordHash field to given value.

### HasWebPasswordHash

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebPasswordHash() bool`

HasWebPasswordHash returns a boolean if a field has been set.

### GetWebUser

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebUser() string`

GetWebUser returns the WebUser field if non-nil, zero value otherwise.

### GetWebUserOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebUserOk() (*string, bool)`

GetWebUserOk returns a tuple with the WebUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUser

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebUser(v string)`

SetWebUser sets WebUser field to given value.

### HasWebUser

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebUser() bool`

HasWebUser returns a boolean if a field has been set.

### GetWebBody

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebBody() string`

GetWebBody returns the WebBody field if non-nil, zero value otherwise.

### GetWebBodyOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebBodyOk() (*string, bool)`

GetWebBodyOk returns a tuple with the WebBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBody

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebBody(v string)`

SetWebBody sets WebBody field to given value.

### HasWebBody

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebBody() bool`

HasWebBody returns a boolean if a field has been set.

### GetWebHeaders

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebHeaders() string`

GetWebHeaders returns the WebHeaders field if non-nil, zero value otherwise.

### GetWebHeadersOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebHeadersOk() (*string, bool)`

GetWebHeadersOk returns a tuple with the WebHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHeaders

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebHeaders(v string)`

SetWebHeaders sets WebHeaders field to given value.

### HasWebHeaders

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebHeaders() bool`

HasWebHeaders returns a boolean if a field has been set.

### GetIgnoreSSL

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetIgnoreSSL() string`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetIgnoreSSLOk() (*string, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetIgnoreSSL(v string)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetWebMethod

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.

### HasWebMethod

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebMethod() bool`

HasWebMethod returns a boolean if a field has been set.

### GetWebUrl

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetJsScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetJsScript() string`

GetJsScript returns the JsScript field if non-nil, zero value otherwise.

### GetJsScriptOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetJsScriptOk() (*string, bool)`

GetJsScriptOk returns a tuple with the JsScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetJsScript(v string)`

SetJsScript sets JsScript field to given value.

### HasJsScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasJsScript() bool`

HasJsScript returns a boolean if a field has been set.

### GetContainerScriptId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerScriptId() string`

GetContainerScriptId returns the ContainerScriptId field if non-nil, zero value otherwise.

### GetContainerScriptIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerScriptIdOk() (*string, bool)`

GetContainerScriptIdOk returns a tuple with the ContainerScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScriptId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetContainerScriptId(v string)`

SetContainerScriptId sets ContainerScriptId field to given value.

### HasContainerScriptId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasContainerScriptId() bool`

HasContainerScriptId returns a boolean if a field has been set.

### GetContainerScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerScript() string`

GetContainerScript returns the ContainerScript field if non-nil, zero value otherwise.

### GetContainerScriptOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerScriptOk() (*string, bool)`

GetContainerScriptOk returns a tuple with the ContainerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetContainerScript(v string)`

SetContainerScript sets ContainerScript field to given value.

### HasContainerScript

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasContainerScript() bool`

HasContainerScript returns a boolean if a field has been set.

### GetContainerTemplateId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerTemplateId() string`

GetContainerTemplateId returns the ContainerTemplateId field if non-nil, zero value otherwise.

### GetContainerTemplateIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerTemplateIdOk() (*string, bool)`

GetContainerTemplateIdOk returns a tuple with the ContainerTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplateId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetContainerTemplateId(v string)`

SetContainerTemplateId sets ContainerTemplateId field to given value.

### HasContainerTemplateId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasContainerTemplateId() bool`

HasContainerTemplateId returns a boolean if a field has been set.

### GetContainerTemplate

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerTemplate() string`

GetContainerTemplate returns the ContainerTemplate field if non-nil, zero value otherwise.

### GetContainerTemplateOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetContainerTemplateOk() (*string, bool)`

GetContainerTemplateOk returns a tuple with the ContainerTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplate

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetContainerTemplate(v string)`

SetContainerTemplate sets ContainerTemplate field to given value.

### HasContainerTemplate

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasContainerTemplate() bool`

HasContainerTemplate returns a boolean if a field has been set.

### GetOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetOperationalWorkflowId() string`

GetOperationalWorkflowId returns the OperationalWorkflowId field if non-nil, zero value otherwise.

### GetOperationalWorkflowIdOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetOperationalWorkflowIdOk() (*string, bool)`

GetOperationalWorkflowIdOk returns a tuple with the OperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowId

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetOperationalWorkflowId(v string)`

SetOperationalWorkflowId sets OperationalWorkflowId field to given value.


### GetOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetOperationalWorkflowName() string`

GetOperationalWorkflowName returns the OperationalWorkflowName field if non-nil, zero value otherwise.

### GetOperationalWorkflowNameOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetOperationalWorkflowNameOk() (*string, bool)`

GetOperationalWorkflowNameOk returns a tuple with the OperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetOperationalWorkflowName(v string)`

SetOperationalWorkflowName sets OperationalWorkflowName field to given value.

### HasOperationalWorkflowName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasOperationalWorkflowName() bool`

HasOperationalWorkflowName returns a boolean if a field has been set.

### GetPuppetEnvironment

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPuppetEnvironment() string`

GetPuppetEnvironment returns the PuppetEnvironment field if non-nil, zero value otherwise.

### GetPuppetEnvironmentOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPuppetEnvironmentOk() (*string, bool)`

GetPuppetEnvironmentOk returns a tuple with the PuppetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetEnvironment

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPuppetEnvironment(v string)`

SetPuppetEnvironment sets PuppetEnvironment field to given value.

### HasPuppetEnvironment

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPuppetEnvironment() bool`

HasPuppetEnvironment returns a boolean if a field has been set.

### GetPuppetNodeName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPuppetNodeName() string`

GetPuppetNodeName returns the PuppetNodeName field if non-nil, zero value otherwise.

### GetPuppetNodeNameOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPuppetNodeNameOk() (*string, bool)`

GetPuppetNodeNameOk returns a tuple with the PuppetNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetNodeName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPuppetNodeName(v string)`

SetPuppetNodeName sets PuppetNodeName field to given value.

### HasPuppetNodeName

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPuppetNodeName() bool`

HasPuppetNodeName returns a boolean if a field has been set.

### GetPythonArgs

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPythonArgs() string`

GetPythonArgs returns the PythonArgs field if non-nil, zero value otherwise.

### GetPythonArgsOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPythonArgsOk() (*string, bool)`

GetPythonArgsOk returns a tuple with the PythonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonArgs

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPythonArgs(v string)`

SetPythonArgs sets PythonArgs field to given value.

### HasPythonArgs

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPythonArgs() bool`

HasPythonArgs returns a boolean if a field has been set.

### GetPythonBinary

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPythonBinary() string`

GetPythonBinary returns the PythonBinary field if non-nil, zero value otherwise.

### GetPythonBinaryOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPythonBinaryOk() (*string, bool)`

GetPythonBinaryOk returns a tuple with the PythonBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonBinary

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPythonBinary(v string)`

SetPythonBinary sets PythonBinary field to given value.

### HasPythonBinary

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPythonBinary() bool`

HasPythonBinary returns a boolean if a field has been set.

### GetPythonAdditionalPackages

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPythonAdditionalPackages() string`

GetPythonAdditionalPackages returns the PythonAdditionalPackages field if non-nil, zero value otherwise.

### GetPythonAdditionalPackagesOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetPythonAdditionalPackagesOk() (*string, bool)`

GetPythonAdditionalPackagesOk returns a tuple with the PythonAdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonAdditionalPackages

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetPythonAdditionalPackages(v string)`

SetPythonAdditionalPackages sets PythonAdditionalPackages field to given value.

### HasPythonAdditionalPackages

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasPythonAdditionalPackages() bool`

HasPythonAdditionalPackages returns a boolean if a field has been set.

### GetVroBody

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetVroBody() string`

GetVroBody returns the VroBody field if non-nil, zero value otherwise.

### GetVroBodyOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetVroBodyOk() (*string, bool)`

GetVroBodyOk returns a tuple with the VroBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVroBody

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetVroBody(v string)`

SetVroBody sets VroBody field to given value.

### HasVroBody

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasVroBody() bool`

HasVroBody returns a boolean if a field has been set.

### GetWriteAttributesAttributes

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWriteAttributesAttributes() string`

GetWriteAttributesAttributes returns the WriteAttributesAttributes field if non-nil, zero value otherwise.

### GetWriteAttributesAttributesOk

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) GetWriteAttributesAttributesOk() (*string, bool)`

GetWriteAttributesAttributesOk returns a tuple with the WriteAttributesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAttributesAttributes

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) SetWriteAttributesAttributes(v string)`

SetWriteAttributesAttributes sets WriteAttributesAttributes field to given value.

### HasWriteAttributesAttributes

`func (o *AddTasks200ResponseAllOfTaskTaskOptions) HasWriteAttributesAttributes() bool`

HasWriteAttributesAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


