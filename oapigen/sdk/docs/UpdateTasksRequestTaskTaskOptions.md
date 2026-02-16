# UpdateTasksRequestTaskTaskOptions

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

### NewUpdateTasksRequestTaskTaskOptions

`func NewUpdateTasksRequestTaskTaskOptions(operationalWorkflowId string, ) *UpdateTasksRequestTaskTaskOptions`

NewUpdateTasksRequestTaskTaskOptions instantiates a new UpdateTasksRequestTaskTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTasksRequestTaskTaskOptionsWithDefaults

`func NewUpdateTasksRequestTaskTaskOptionsWithDefaults() *UpdateTasksRequestTaskTaskOptions`

NewUpdateTasksRequestTaskTaskOptionsWithDefaults instantiates a new UpdateTasksRequestTaskTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleOptions

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleOptions() string`

GetAnsibleOptions returns the AnsibleOptions field if non-nil, zero value otherwise.

### GetAnsibleOptionsOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleOptionsOk() (*string, bool)`

GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleOptions

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleOptions(v string)`

SetAnsibleOptions sets AnsibleOptions field to given value.

### HasAnsibleOptions

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleOptions() bool`

HasAnsibleOptions returns a boolean if a field has been set.

### GetAnsiblePlaybook

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsiblePlaybook() string`

GetAnsiblePlaybook returns the AnsiblePlaybook field if non-nil, zero value otherwise.

### GetAnsiblePlaybookOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsiblePlaybookOk() (*string, bool)`

GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybook

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsiblePlaybook(v string)`

SetAnsiblePlaybook sets AnsiblePlaybook field to given value.

### HasAnsiblePlaybook

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsiblePlaybook() bool`

HasAnsiblePlaybook returns a boolean if a field has been set.

### GetSshKey

`func (o *UpdateTasksRequestTaskTaskOptions) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *UpdateTasksRequestTaskTaskOptions) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *UpdateTasksRequestTaskTaskOptions) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetPort

`func (o *UpdateTasksRequestTaskTaskOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateTasksRequestTaskTaskOptions) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateTasksRequestTaskTaskOptions) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLocalScriptGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateTasksRequestTaskTaskOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateTasksRequestTaskTaskOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateTasksRequestTaskTaskOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *UpdateTasksRequestTaskTaskOptions) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *UpdateTasksRequestTaskTaskOptions) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *UpdateTasksRequestTaskTaskOptions) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetLocalScriptGitId

`func (o *UpdateTasksRequestTaskTaskOptions) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *UpdateTasksRequestTaskTaskOptions) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *UpdateTasksRequestTaskTaskOptions) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### GetAnsibleGitId

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleGitId() string`

GetAnsibleGitId returns the AnsibleGitId field if non-nil, zero value otherwise.

### GetAnsibleGitIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleGitIdOk() (*string, bool)`

GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitId

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleGitId(v string)`

SetAnsibleGitId sets AnsibleGitId field to given value.

### HasAnsibleGitId

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleGitId() bool`

HasAnsibleGitId returns a boolean if a field has been set.

### GetHost

`func (o *UpdateTasksRequestTaskTaskOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateTasksRequestTaskTaskOptions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateTasksRequestTaskTaskOptions) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAnsibleSkipTags

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleSkipTags() string`

GetAnsibleSkipTags returns the AnsibleSkipTags field if non-nil, zero value otherwise.

### GetAnsibleSkipTagsOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleSkipTagsOk() (*string, bool)`

GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleSkipTags

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleSkipTags(v string)`

SetAnsibleSkipTags sets AnsibleSkipTags field to given value.

### HasAnsibleSkipTags

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleSkipTags() bool`

HasAnsibleSkipTags returns a boolean if a field has been set.

### GetAnsibleTags

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleTags() string`

GetAnsibleTags returns the AnsibleTags field if non-nil, zero value otherwise.

### GetAnsibleTagsOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleTagsOk() (*string, bool)`

GetAnsibleTagsOk returns a tuple with the AnsibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTags

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleTags(v string)`

SetAnsibleTags sets AnsibleTags field to given value.

### HasAnsibleTags

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleTags() bool`

HasAnsibleTags returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateTasksRequestTaskTaskOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateTasksRequestTaskTaskOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateTasksRequestTaskTaskOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAnsibleGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleGitRef() string`

GetAnsibleGitRef returns the AnsibleGitRef field if non-nil, zero value otherwise.

### GetAnsibleGitRefOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleGitRefOk() (*string, bool)`

GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleGitRef(v string)`

SetAnsibleGitRef sets AnsibleGitRef field to given value.

### HasAnsibleGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleGitRef() bool`

HasAnsibleGitRef returns a boolean if a field has been set.

### GetAnsibleTowerGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleTowerGitRef() string`

GetAnsibleTowerGitRef returns the AnsibleTowerGitRef field if non-nil, zero value otherwise.

### GetAnsibleTowerGitRefOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleTowerGitRefOk() (*string, bool)`

GetAnsibleTowerGitRefOk returns a tuple with the AnsibleTowerGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleTowerGitRef(v string)`

SetAnsibleTowerGitRef sets AnsibleTowerGitRef field to given value.

### HasAnsibleTowerGitRef

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleTowerGitRef() bool`

HasAnsibleTowerGitRef returns a boolean if a field has been set.

### GetAnsibleGroup

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleGroup() string`

GetAnsibleGroup returns the AnsibleGroup field if non-nil, zero value otherwise.

### GetAnsibleGroupOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleGroupOk() (*string, bool)`

GetAnsibleGroupOk returns a tuple with the AnsibleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroup

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleGroup(v string)`

SetAnsibleGroup sets AnsibleGroup field to given value.

### HasAnsibleGroup

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleGroup() bool`

HasAnsibleGroup returns a boolean if a field has been set.

### GetAnsibleTowerExecuteMode

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleTowerExecuteMode() string`

GetAnsibleTowerExecuteMode returns the AnsibleTowerExecuteMode field if non-nil, zero value otherwise.

### GetAnsibleTowerExecuteModeOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetAnsibleTowerExecuteModeOk() (*string, bool)`

GetAnsibleTowerExecuteModeOk returns a tuple with the AnsibleTowerExecuteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerExecuteMode

`func (o *UpdateTasksRequestTaskTaskOptions) SetAnsibleTowerExecuteMode(v string)`

SetAnsibleTowerExecuteMode sets AnsibleTowerExecuteMode field to given value.

### HasAnsibleTowerExecuteMode

`func (o *UpdateTasksRequestTaskTaskOptions) HasAnsibleTowerExecuteMode() bool`

HasAnsibleTowerExecuteMode returns a boolean if a field has been set.

### GetChefDataKey

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefDataKey() string`

GetChefDataKey returns the ChefDataKey field if non-nil, zero value otherwise.

### GetChefDataKeyOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefDataKeyOk() (*string, bool)`

GetChefDataKeyOk returns a tuple with the ChefDataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKey

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefDataKey(v string)`

SetChefDataKey sets ChefDataKey field to given value.

### HasChefDataKey

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefDataKey() bool`

HasChefDataKey returns a boolean if a field has been set.

### GetChefDataKeyHash

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefDataKeyHash() string`

GetChefDataKeyHash returns the ChefDataKeyHash field if non-nil, zero value otherwise.

### GetChefDataKeyHashOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefDataKeyHashOk() (*string, bool)`

GetChefDataKeyHashOk returns a tuple with the ChefDataKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyHash

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefDataKeyHash(v string)`

SetChefDataKeyHash sets ChefDataKeyHash field to given value.

### HasChefDataKeyHash

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefDataKeyHash() bool`

HasChefDataKeyHash returns a boolean if a field has been set.

### GetChefRunList

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefRunList() string`

GetChefRunList returns the ChefRunList field if non-nil, zero value otherwise.

### GetChefRunListOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefRunListOk() (*string, bool)`

GetChefRunListOk returns a tuple with the ChefRunList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefRunList

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefRunList(v string)`

SetChefRunList sets ChefRunList field to given value.

### HasChefRunList

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefRunList() bool`

HasChefRunList returns a boolean if a field has been set.

### GetChefDataKeyPath

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefDataKeyPath() string`

GetChefDataKeyPath returns the ChefDataKeyPath field if non-nil, zero value otherwise.

### GetChefDataKeyPathOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefDataKeyPathOk() (*string, bool)`

GetChefDataKeyPathOk returns a tuple with the ChefDataKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyPath

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefDataKeyPath(v string)`

SetChefDataKeyPath sets ChefDataKeyPath field to given value.

### HasChefDataKeyPath

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefDataKeyPath() bool`

HasChefDataKeyPath returns a boolean if a field has been set.

### GetChefEnv

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefEnv() string`

GetChefEnv returns the ChefEnv field if non-nil, zero value otherwise.

### GetChefEnvOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefEnvOk() (*string, bool)`

GetChefEnvOk returns a tuple with the ChefEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefEnv

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefEnv(v string)`

SetChefEnv sets ChefEnv field to given value.

### HasChefEnv

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefEnv() bool`

HasChefEnv returns a boolean if a field has been set.

### GetChefNodeName

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefNodeName() string`

GetChefNodeName returns the ChefNodeName field if non-nil, zero value otherwise.

### GetChefNodeNameOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefNodeNameOk() (*string, bool)`

GetChefNodeNameOk returns a tuple with the ChefNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefNodeName

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefNodeName(v string)`

SetChefNodeName sets ChefNodeName field to given value.

### HasChefNodeName

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefNodeName() bool`

HasChefNodeName returns a boolean if a field has been set.

### GetChefAttributes

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefAttributes() string`

GetChefAttributes returns the ChefAttributes field if non-nil, zero value otherwise.

### GetChefAttributesOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetChefAttributesOk() (*string, bool)`

GetChefAttributesOk returns a tuple with the ChefAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefAttributes

`func (o *UpdateTasksRequestTaskTaskOptions) SetChefAttributes(v string)`

SetChefAttributes sets ChefAttributes field to given value.

### HasChefAttributes

`func (o *UpdateTasksRequestTaskTaskOptions) HasChefAttributes() bool`

HasChefAttributes returns a boolean if a field has been set.

### GetConditionalScript

`func (o *UpdateTasksRequestTaskTaskOptions) GetConditionalScript() string`

GetConditionalScript returns the ConditionalScript field if non-nil, zero value otherwise.

### GetConditionalScriptOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetConditionalScriptOk() (*string, bool)`

GetConditionalScriptOk returns a tuple with the ConditionalScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalScript

`func (o *UpdateTasksRequestTaskTaskOptions) SetConditionalScript(v string)`

SetConditionalScript sets ConditionalScript field to given value.

### HasConditionalScript

`func (o *UpdateTasksRequestTaskTaskOptions) HasConditionalScript() bool`

HasConditionalScript returns a boolean if a field has been set.

### GetIfOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) GetIfOperationalWorkflowId() int64`

GetIfOperationalWorkflowId returns the IfOperationalWorkflowId field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetIfOperationalWorkflowIdOk() (*int64, bool)`

GetIfOperationalWorkflowIdOk returns a tuple with the IfOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) SetIfOperationalWorkflowId(v int64)`

SetIfOperationalWorkflowId sets IfOperationalWorkflowId field to given value.

### HasIfOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) HasIfOperationalWorkflowId() bool`

HasIfOperationalWorkflowId returns a boolean if a field has been set.

### GetIfOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) GetIfOperationalWorkflowName() string`

GetIfOperationalWorkflowName returns the IfOperationalWorkflowName field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowNameOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetIfOperationalWorkflowNameOk() (*string, bool)`

GetIfOperationalWorkflowNameOk returns a tuple with the IfOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) SetIfOperationalWorkflowName(v string)`

SetIfOperationalWorkflowName sets IfOperationalWorkflowName field to given value.

### HasIfOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) HasIfOperationalWorkflowName() bool`

HasIfOperationalWorkflowName returns a boolean if a field has been set.

### GetElseOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) GetElseOperationalWorkflowId() int64`

GetElseOperationalWorkflowId returns the ElseOperationalWorkflowId field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetElseOperationalWorkflowIdOk() (*int64, bool)`

GetElseOperationalWorkflowIdOk returns a tuple with the ElseOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) SetElseOperationalWorkflowId(v int64)`

SetElseOperationalWorkflowId sets ElseOperationalWorkflowId field to given value.

### HasElseOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) HasElseOperationalWorkflowId() bool`

HasElseOperationalWorkflowId returns a boolean if a field has been set.

### GetElseOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) GetElseOperationalWorkflowName() string`

GetElseOperationalWorkflowName returns the ElseOperationalWorkflowName field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowNameOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetElseOperationalWorkflowNameOk() (*string, bool)`

GetElseOperationalWorkflowNameOk returns a tuple with the ElseOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) SetElseOperationalWorkflowName(v string)`

SetElseOperationalWorkflowName sets ElseOperationalWorkflowName field to given value.

### HasElseOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) HasElseOperationalWorkflowName() bool`

HasElseOperationalWorkflowName returns a boolean if a field has been set.

### GetEmailSkipTemplate

`func (o *UpdateTasksRequestTaskTaskOptions) GetEmailSkipTemplate() string`

GetEmailSkipTemplate returns the EmailSkipTemplate field if non-nil, zero value otherwise.

### GetEmailSkipTemplateOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetEmailSkipTemplateOk() (*string, bool)`

GetEmailSkipTemplateOk returns a tuple with the EmailSkipTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSkipTemplate

`func (o *UpdateTasksRequestTaskTaskOptions) SetEmailSkipTemplate(v string)`

SetEmailSkipTemplate sets EmailSkipTemplate field to given value.

### HasEmailSkipTemplate

`func (o *UpdateTasksRequestTaskTaskOptions) HasEmailSkipTemplate() bool`

HasEmailSkipTemplate returns a boolean if a field has been set.

### GetEmailSubject

`func (o *UpdateTasksRequestTaskTaskOptions) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *UpdateTasksRequestTaskTaskOptions) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *UpdateTasksRequestTaskTaskOptions) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailAddress

`func (o *UpdateTasksRequestTaskTaskOptions) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UpdateTasksRequestTaskTaskOptions) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UpdateTasksRequestTaskTaskOptions) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetWebPassword

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebPassword() string`

GetWebPassword returns the WebPassword field if non-nil, zero value otherwise.

### GetWebPasswordOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebPasswordOk() (*string, bool)`

GetWebPasswordOk returns a tuple with the WebPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPassword

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebPassword(v string)`

SetWebPassword sets WebPassword field to given value.

### HasWebPassword

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebPassword() bool`

HasWebPassword returns a boolean if a field has been set.

### GetWebPasswordHash

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebPasswordHash() string`

GetWebPasswordHash returns the WebPasswordHash field if non-nil, zero value otherwise.

### GetWebPasswordHashOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebPasswordHashOk() (*string, bool)`

GetWebPasswordHashOk returns a tuple with the WebPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPasswordHash

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebPasswordHash(v string)`

SetWebPasswordHash sets WebPasswordHash field to given value.

### HasWebPasswordHash

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebPasswordHash() bool`

HasWebPasswordHash returns a boolean if a field has been set.

### GetWebUser

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebUser() string`

GetWebUser returns the WebUser field if non-nil, zero value otherwise.

### GetWebUserOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebUserOk() (*string, bool)`

GetWebUserOk returns a tuple with the WebUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUser

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebUser(v string)`

SetWebUser sets WebUser field to given value.

### HasWebUser

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebUser() bool`

HasWebUser returns a boolean if a field has been set.

### GetWebBody

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebBody() string`

GetWebBody returns the WebBody field if non-nil, zero value otherwise.

### GetWebBodyOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebBodyOk() (*string, bool)`

GetWebBodyOk returns a tuple with the WebBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBody

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebBody(v string)`

SetWebBody sets WebBody field to given value.

### HasWebBody

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebBody() bool`

HasWebBody returns a boolean if a field has been set.

### GetWebHeaders

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebHeaders() string`

GetWebHeaders returns the WebHeaders field if non-nil, zero value otherwise.

### GetWebHeadersOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebHeadersOk() (*string, bool)`

GetWebHeadersOk returns a tuple with the WebHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHeaders

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebHeaders(v string)`

SetWebHeaders sets WebHeaders field to given value.

### HasWebHeaders

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebHeaders() bool`

HasWebHeaders returns a boolean if a field has been set.

### GetIgnoreSSL

`func (o *UpdateTasksRequestTaskTaskOptions) GetIgnoreSSL() string`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetIgnoreSSLOk() (*string, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *UpdateTasksRequestTaskTaskOptions) SetIgnoreSSL(v string)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *UpdateTasksRequestTaskTaskOptions) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetWebMethod

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.

### HasWebMethod

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebMethod() bool`

HasWebMethod returns a boolean if a field has been set.

### GetWebUrl

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *UpdateTasksRequestTaskTaskOptions) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *UpdateTasksRequestTaskTaskOptions) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetJsScript

`func (o *UpdateTasksRequestTaskTaskOptions) GetJsScript() string`

GetJsScript returns the JsScript field if non-nil, zero value otherwise.

### GetJsScriptOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetJsScriptOk() (*string, bool)`

GetJsScriptOk returns a tuple with the JsScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsScript

`func (o *UpdateTasksRequestTaskTaskOptions) SetJsScript(v string)`

SetJsScript sets JsScript field to given value.

### HasJsScript

`func (o *UpdateTasksRequestTaskTaskOptions) HasJsScript() bool`

HasJsScript returns a boolean if a field has been set.

### GetContainerScriptId

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerScriptId() string`

GetContainerScriptId returns the ContainerScriptId field if non-nil, zero value otherwise.

### GetContainerScriptIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerScriptIdOk() (*string, bool)`

GetContainerScriptIdOk returns a tuple with the ContainerScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScriptId

`func (o *UpdateTasksRequestTaskTaskOptions) SetContainerScriptId(v string)`

SetContainerScriptId sets ContainerScriptId field to given value.

### HasContainerScriptId

`func (o *UpdateTasksRequestTaskTaskOptions) HasContainerScriptId() bool`

HasContainerScriptId returns a boolean if a field has been set.

### GetContainerScript

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerScript() string`

GetContainerScript returns the ContainerScript field if non-nil, zero value otherwise.

### GetContainerScriptOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerScriptOk() (*string, bool)`

GetContainerScriptOk returns a tuple with the ContainerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScript

`func (o *UpdateTasksRequestTaskTaskOptions) SetContainerScript(v string)`

SetContainerScript sets ContainerScript field to given value.

### HasContainerScript

`func (o *UpdateTasksRequestTaskTaskOptions) HasContainerScript() bool`

HasContainerScript returns a boolean if a field has been set.

### GetContainerTemplateId

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerTemplateId() string`

GetContainerTemplateId returns the ContainerTemplateId field if non-nil, zero value otherwise.

### GetContainerTemplateIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerTemplateIdOk() (*string, bool)`

GetContainerTemplateIdOk returns a tuple with the ContainerTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplateId

`func (o *UpdateTasksRequestTaskTaskOptions) SetContainerTemplateId(v string)`

SetContainerTemplateId sets ContainerTemplateId field to given value.

### HasContainerTemplateId

`func (o *UpdateTasksRequestTaskTaskOptions) HasContainerTemplateId() bool`

HasContainerTemplateId returns a boolean if a field has been set.

### GetContainerTemplate

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerTemplate() string`

GetContainerTemplate returns the ContainerTemplate field if non-nil, zero value otherwise.

### GetContainerTemplateOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetContainerTemplateOk() (*string, bool)`

GetContainerTemplateOk returns a tuple with the ContainerTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplate

`func (o *UpdateTasksRequestTaskTaskOptions) SetContainerTemplate(v string)`

SetContainerTemplate sets ContainerTemplate field to given value.

### HasContainerTemplate

`func (o *UpdateTasksRequestTaskTaskOptions) HasContainerTemplate() bool`

HasContainerTemplate returns a boolean if a field has been set.

### GetOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) GetOperationalWorkflowId() string`

GetOperationalWorkflowId returns the OperationalWorkflowId field if non-nil, zero value otherwise.

### GetOperationalWorkflowIdOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetOperationalWorkflowIdOk() (*string, bool)`

GetOperationalWorkflowIdOk returns a tuple with the OperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowId

`func (o *UpdateTasksRequestTaskTaskOptions) SetOperationalWorkflowId(v string)`

SetOperationalWorkflowId sets OperationalWorkflowId field to given value.


### GetOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) GetOperationalWorkflowName() string`

GetOperationalWorkflowName returns the OperationalWorkflowName field if non-nil, zero value otherwise.

### GetOperationalWorkflowNameOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetOperationalWorkflowNameOk() (*string, bool)`

GetOperationalWorkflowNameOk returns a tuple with the OperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) SetOperationalWorkflowName(v string)`

SetOperationalWorkflowName sets OperationalWorkflowName field to given value.

### HasOperationalWorkflowName

`func (o *UpdateTasksRequestTaskTaskOptions) HasOperationalWorkflowName() bool`

HasOperationalWorkflowName returns a boolean if a field has been set.

### GetPuppetEnvironment

`func (o *UpdateTasksRequestTaskTaskOptions) GetPuppetEnvironment() string`

GetPuppetEnvironment returns the PuppetEnvironment field if non-nil, zero value otherwise.

### GetPuppetEnvironmentOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPuppetEnvironmentOk() (*string, bool)`

GetPuppetEnvironmentOk returns a tuple with the PuppetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetEnvironment

`func (o *UpdateTasksRequestTaskTaskOptions) SetPuppetEnvironment(v string)`

SetPuppetEnvironment sets PuppetEnvironment field to given value.

### HasPuppetEnvironment

`func (o *UpdateTasksRequestTaskTaskOptions) HasPuppetEnvironment() bool`

HasPuppetEnvironment returns a boolean if a field has been set.

### GetPuppetNodeName

`func (o *UpdateTasksRequestTaskTaskOptions) GetPuppetNodeName() string`

GetPuppetNodeName returns the PuppetNodeName field if non-nil, zero value otherwise.

### GetPuppetNodeNameOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPuppetNodeNameOk() (*string, bool)`

GetPuppetNodeNameOk returns a tuple with the PuppetNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetNodeName

`func (o *UpdateTasksRequestTaskTaskOptions) SetPuppetNodeName(v string)`

SetPuppetNodeName sets PuppetNodeName field to given value.

### HasPuppetNodeName

`func (o *UpdateTasksRequestTaskTaskOptions) HasPuppetNodeName() bool`

HasPuppetNodeName returns a boolean if a field has been set.

### GetPythonArgs

`func (o *UpdateTasksRequestTaskTaskOptions) GetPythonArgs() string`

GetPythonArgs returns the PythonArgs field if non-nil, zero value otherwise.

### GetPythonArgsOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPythonArgsOk() (*string, bool)`

GetPythonArgsOk returns a tuple with the PythonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonArgs

`func (o *UpdateTasksRequestTaskTaskOptions) SetPythonArgs(v string)`

SetPythonArgs sets PythonArgs field to given value.

### HasPythonArgs

`func (o *UpdateTasksRequestTaskTaskOptions) HasPythonArgs() bool`

HasPythonArgs returns a boolean if a field has been set.

### GetPythonBinary

`func (o *UpdateTasksRequestTaskTaskOptions) GetPythonBinary() string`

GetPythonBinary returns the PythonBinary field if non-nil, zero value otherwise.

### GetPythonBinaryOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPythonBinaryOk() (*string, bool)`

GetPythonBinaryOk returns a tuple with the PythonBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonBinary

`func (o *UpdateTasksRequestTaskTaskOptions) SetPythonBinary(v string)`

SetPythonBinary sets PythonBinary field to given value.

### HasPythonBinary

`func (o *UpdateTasksRequestTaskTaskOptions) HasPythonBinary() bool`

HasPythonBinary returns a boolean if a field has been set.

### GetPythonAdditionalPackages

`func (o *UpdateTasksRequestTaskTaskOptions) GetPythonAdditionalPackages() string`

GetPythonAdditionalPackages returns the PythonAdditionalPackages field if non-nil, zero value otherwise.

### GetPythonAdditionalPackagesOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetPythonAdditionalPackagesOk() (*string, bool)`

GetPythonAdditionalPackagesOk returns a tuple with the PythonAdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonAdditionalPackages

`func (o *UpdateTasksRequestTaskTaskOptions) SetPythonAdditionalPackages(v string)`

SetPythonAdditionalPackages sets PythonAdditionalPackages field to given value.

### HasPythonAdditionalPackages

`func (o *UpdateTasksRequestTaskTaskOptions) HasPythonAdditionalPackages() bool`

HasPythonAdditionalPackages returns a boolean if a field has been set.

### GetVroBody

`func (o *UpdateTasksRequestTaskTaskOptions) GetVroBody() string`

GetVroBody returns the VroBody field if non-nil, zero value otherwise.

### GetVroBodyOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetVroBodyOk() (*string, bool)`

GetVroBodyOk returns a tuple with the VroBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVroBody

`func (o *UpdateTasksRequestTaskTaskOptions) SetVroBody(v string)`

SetVroBody sets VroBody field to given value.

### HasVroBody

`func (o *UpdateTasksRequestTaskTaskOptions) HasVroBody() bool`

HasVroBody returns a boolean if a field has been set.

### GetWriteAttributesAttributes

`func (o *UpdateTasksRequestTaskTaskOptions) GetWriteAttributesAttributes() string`

GetWriteAttributesAttributes returns the WriteAttributesAttributes field if non-nil, zero value otherwise.

### GetWriteAttributesAttributesOk

`func (o *UpdateTasksRequestTaskTaskOptions) GetWriteAttributesAttributesOk() (*string, bool)`

GetWriteAttributesAttributesOk returns a tuple with the WriteAttributesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAttributesAttributes

`func (o *UpdateTasksRequestTaskTaskOptions) SetWriteAttributesAttributes(v string)`

SetWriteAttributesAttributes sets WriteAttributesAttributes field to given value.

### HasWriteAttributesAttributes

`func (o *UpdateTasksRequestTaskTaskOptions) HasWriteAttributesAttributes() bool`

HasWriteAttributesAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


