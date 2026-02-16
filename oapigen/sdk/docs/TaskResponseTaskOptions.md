# TaskResponseTaskOptions

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

### NewTaskResponseTaskOptions

`func NewTaskResponseTaskOptions(operationalWorkflowId string, ) *TaskResponseTaskOptions`

NewTaskResponseTaskOptions instantiates a new TaskResponseTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseTaskOptionsWithDefaults

`func NewTaskResponseTaskOptionsWithDefaults() *TaskResponseTaskOptions`

NewTaskResponseTaskOptionsWithDefaults instantiates a new TaskResponseTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleOptions

`func (o *TaskResponseTaskOptions) GetAnsibleOptions() string`

GetAnsibleOptions returns the AnsibleOptions field if non-nil, zero value otherwise.

### GetAnsibleOptionsOk

`func (o *TaskResponseTaskOptions) GetAnsibleOptionsOk() (*string, bool)`

GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleOptions

`func (o *TaskResponseTaskOptions) SetAnsibleOptions(v string)`

SetAnsibleOptions sets AnsibleOptions field to given value.

### HasAnsibleOptions

`func (o *TaskResponseTaskOptions) HasAnsibleOptions() bool`

HasAnsibleOptions returns a boolean if a field has been set.

### GetAnsiblePlaybook

`func (o *TaskResponseTaskOptions) GetAnsiblePlaybook() string`

GetAnsiblePlaybook returns the AnsiblePlaybook field if non-nil, zero value otherwise.

### GetAnsiblePlaybookOk

`func (o *TaskResponseTaskOptions) GetAnsiblePlaybookOk() (*string, bool)`

GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybook

`func (o *TaskResponseTaskOptions) SetAnsiblePlaybook(v string)`

SetAnsiblePlaybook sets AnsiblePlaybook field to given value.

### HasAnsiblePlaybook

`func (o *TaskResponseTaskOptions) HasAnsiblePlaybook() bool`

HasAnsiblePlaybook returns a boolean if a field has been set.

### GetSshKey

`func (o *TaskResponseTaskOptions) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskResponseTaskOptions) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskResponseTaskOptions) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskResponseTaskOptions) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetPort

`func (o *TaskResponseTaskOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskResponseTaskOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskResponseTaskOptions) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskResponseTaskOptions) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLocalScriptGitRef

`func (o *TaskResponseTaskOptions) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskResponseTaskOptions) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskResponseTaskOptions) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskResponseTaskOptions) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### GetPassword

`func (o *TaskResponseTaskOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskResponseTaskOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskResponseTaskOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskResponseTaskOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *TaskResponseTaskOptions) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskResponseTaskOptions) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskResponseTaskOptions) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskResponseTaskOptions) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetLocalScriptGitId

`func (o *TaskResponseTaskOptions) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskResponseTaskOptions) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskResponseTaskOptions) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskResponseTaskOptions) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### GetAnsibleGitId

`func (o *TaskResponseTaskOptions) GetAnsibleGitId() string`

GetAnsibleGitId returns the AnsibleGitId field if non-nil, zero value otherwise.

### GetAnsibleGitIdOk

`func (o *TaskResponseTaskOptions) GetAnsibleGitIdOk() (*string, bool)`

GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitId

`func (o *TaskResponseTaskOptions) SetAnsibleGitId(v string)`

SetAnsibleGitId sets AnsibleGitId field to given value.

### HasAnsibleGitId

`func (o *TaskResponseTaskOptions) HasAnsibleGitId() bool`

HasAnsibleGitId returns a boolean if a field has been set.

### GetHost

`func (o *TaskResponseTaskOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskResponseTaskOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskResponseTaskOptions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskResponseTaskOptions) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAnsibleSkipTags

`func (o *TaskResponseTaskOptions) GetAnsibleSkipTags() string`

GetAnsibleSkipTags returns the AnsibleSkipTags field if non-nil, zero value otherwise.

### GetAnsibleSkipTagsOk

`func (o *TaskResponseTaskOptions) GetAnsibleSkipTagsOk() (*string, bool)`

GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleSkipTags

`func (o *TaskResponseTaskOptions) SetAnsibleSkipTags(v string)`

SetAnsibleSkipTags sets AnsibleSkipTags field to given value.

### HasAnsibleSkipTags

`func (o *TaskResponseTaskOptions) HasAnsibleSkipTags() bool`

HasAnsibleSkipTags returns a boolean if a field has been set.

### GetAnsibleTags

`func (o *TaskResponseTaskOptions) GetAnsibleTags() string`

GetAnsibleTags returns the AnsibleTags field if non-nil, zero value otherwise.

### GetAnsibleTagsOk

`func (o *TaskResponseTaskOptions) GetAnsibleTagsOk() (*string, bool)`

GetAnsibleTagsOk returns a tuple with the AnsibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTags

`func (o *TaskResponseTaskOptions) SetAnsibleTags(v string)`

SetAnsibleTags sets AnsibleTags field to given value.

### HasAnsibleTags

`func (o *TaskResponseTaskOptions) HasAnsibleTags() bool`

HasAnsibleTags returns a boolean if a field has been set.

### GetUsername

`func (o *TaskResponseTaskOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskResponseTaskOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskResponseTaskOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskResponseTaskOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAnsibleGitRef

`func (o *TaskResponseTaskOptions) GetAnsibleGitRef() string`

GetAnsibleGitRef returns the AnsibleGitRef field if non-nil, zero value otherwise.

### GetAnsibleGitRefOk

`func (o *TaskResponseTaskOptions) GetAnsibleGitRefOk() (*string, bool)`

GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitRef

`func (o *TaskResponseTaskOptions) SetAnsibleGitRef(v string)`

SetAnsibleGitRef sets AnsibleGitRef field to given value.

### HasAnsibleGitRef

`func (o *TaskResponseTaskOptions) HasAnsibleGitRef() bool`

HasAnsibleGitRef returns a boolean if a field has been set.

### GetAnsibleTowerGitRef

`func (o *TaskResponseTaskOptions) GetAnsibleTowerGitRef() string`

GetAnsibleTowerGitRef returns the AnsibleTowerGitRef field if non-nil, zero value otherwise.

### GetAnsibleTowerGitRefOk

`func (o *TaskResponseTaskOptions) GetAnsibleTowerGitRefOk() (*string, bool)`

GetAnsibleTowerGitRefOk returns a tuple with the AnsibleTowerGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerGitRef

`func (o *TaskResponseTaskOptions) SetAnsibleTowerGitRef(v string)`

SetAnsibleTowerGitRef sets AnsibleTowerGitRef field to given value.

### HasAnsibleTowerGitRef

`func (o *TaskResponseTaskOptions) HasAnsibleTowerGitRef() bool`

HasAnsibleTowerGitRef returns a boolean if a field has been set.

### GetAnsibleGroup

`func (o *TaskResponseTaskOptions) GetAnsibleGroup() string`

GetAnsibleGroup returns the AnsibleGroup field if non-nil, zero value otherwise.

### GetAnsibleGroupOk

`func (o *TaskResponseTaskOptions) GetAnsibleGroupOk() (*string, bool)`

GetAnsibleGroupOk returns a tuple with the AnsibleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroup

`func (o *TaskResponseTaskOptions) SetAnsibleGroup(v string)`

SetAnsibleGroup sets AnsibleGroup field to given value.

### HasAnsibleGroup

`func (o *TaskResponseTaskOptions) HasAnsibleGroup() bool`

HasAnsibleGroup returns a boolean if a field has been set.

### GetAnsibleTowerExecuteMode

`func (o *TaskResponseTaskOptions) GetAnsibleTowerExecuteMode() string`

GetAnsibleTowerExecuteMode returns the AnsibleTowerExecuteMode field if non-nil, zero value otherwise.

### GetAnsibleTowerExecuteModeOk

`func (o *TaskResponseTaskOptions) GetAnsibleTowerExecuteModeOk() (*string, bool)`

GetAnsibleTowerExecuteModeOk returns a tuple with the AnsibleTowerExecuteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerExecuteMode

`func (o *TaskResponseTaskOptions) SetAnsibleTowerExecuteMode(v string)`

SetAnsibleTowerExecuteMode sets AnsibleTowerExecuteMode field to given value.

### HasAnsibleTowerExecuteMode

`func (o *TaskResponseTaskOptions) HasAnsibleTowerExecuteMode() bool`

HasAnsibleTowerExecuteMode returns a boolean if a field has been set.

### GetChefDataKey

`func (o *TaskResponseTaskOptions) GetChefDataKey() string`

GetChefDataKey returns the ChefDataKey field if non-nil, zero value otherwise.

### GetChefDataKeyOk

`func (o *TaskResponseTaskOptions) GetChefDataKeyOk() (*string, bool)`

GetChefDataKeyOk returns a tuple with the ChefDataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKey

`func (o *TaskResponseTaskOptions) SetChefDataKey(v string)`

SetChefDataKey sets ChefDataKey field to given value.

### HasChefDataKey

`func (o *TaskResponseTaskOptions) HasChefDataKey() bool`

HasChefDataKey returns a boolean if a field has been set.

### GetChefDataKeyHash

`func (o *TaskResponseTaskOptions) GetChefDataKeyHash() string`

GetChefDataKeyHash returns the ChefDataKeyHash field if non-nil, zero value otherwise.

### GetChefDataKeyHashOk

`func (o *TaskResponseTaskOptions) GetChefDataKeyHashOk() (*string, bool)`

GetChefDataKeyHashOk returns a tuple with the ChefDataKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyHash

`func (o *TaskResponseTaskOptions) SetChefDataKeyHash(v string)`

SetChefDataKeyHash sets ChefDataKeyHash field to given value.

### HasChefDataKeyHash

`func (o *TaskResponseTaskOptions) HasChefDataKeyHash() bool`

HasChefDataKeyHash returns a boolean if a field has been set.

### GetChefRunList

`func (o *TaskResponseTaskOptions) GetChefRunList() string`

GetChefRunList returns the ChefRunList field if non-nil, zero value otherwise.

### GetChefRunListOk

`func (o *TaskResponseTaskOptions) GetChefRunListOk() (*string, bool)`

GetChefRunListOk returns a tuple with the ChefRunList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefRunList

`func (o *TaskResponseTaskOptions) SetChefRunList(v string)`

SetChefRunList sets ChefRunList field to given value.

### HasChefRunList

`func (o *TaskResponseTaskOptions) HasChefRunList() bool`

HasChefRunList returns a boolean if a field has been set.

### GetChefDataKeyPath

`func (o *TaskResponseTaskOptions) GetChefDataKeyPath() string`

GetChefDataKeyPath returns the ChefDataKeyPath field if non-nil, zero value otherwise.

### GetChefDataKeyPathOk

`func (o *TaskResponseTaskOptions) GetChefDataKeyPathOk() (*string, bool)`

GetChefDataKeyPathOk returns a tuple with the ChefDataKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyPath

`func (o *TaskResponseTaskOptions) SetChefDataKeyPath(v string)`

SetChefDataKeyPath sets ChefDataKeyPath field to given value.

### HasChefDataKeyPath

`func (o *TaskResponseTaskOptions) HasChefDataKeyPath() bool`

HasChefDataKeyPath returns a boolean if a field has been set.

### GetChefEnv

`func (o *TaskResponseTaskOptions) GetChefEnv() string`

GetChefEnv returns the ChefEnv field if non-nil, zero value otherwise.

### GetChefEnvOk

`func (o *TaskResponseTaskOptions) GetChefEnvOk() (*string, bool)`

GetChefEnvOk returns a tuple with the ChefEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefEnv

`func (o *TaskResponseTaskOptions) SetChefEnv(v string)`

SetChefEnv sets ChefEnv field to given value.

### HasChefEnv

`func (o *TaskResponseTaskOptions) HasChefEnv() bool`

HasChefEnv returns a boolean if a field has been set.

### GetChefNodeName

`func (o *TaskResponseTaskOptions) GetChefNodeName() string`

GetChefNodeName returns the ChefNodeName field if non-nil, zero value otherwise.

### GetChefNodeNameOk

`func (o *TaskResponseTaskOptions) GetChefNodeNameOk() (*string, bool)`

GetChefNodeNameOk returns a tuple with the ChefNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefNodeName

`func (o *TaskResponseTaskOptions) SetChefNodeName(v string)`

SetChefNodeName sets ChefNodeName field to given value.

### HasChefNodeName

`func (o *TaskResponseTaskOptions) HasChefNodeName() bool`

HasChefNodeName returns a boolean if a field has been set.

### GetChefAttributes

`func (o *TaskResponseTaskOptions) GetChefAttributes() string`

GetChefAttributes returns the ChefAttributes field if non-nil, zero value otherwise.

### GetChefAttributesOk

`func (o *TaskResponseTaskOptions) GetChefAttributesOk() (*string, bool)`

GetChefAttributesOk returns a tuple with the ChefAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefAttributes

`func (o *TaskResponseTaskOptions) SetChefAttributes(v string)`

SetChefAttributes sets ChefAttributes field to given value.

### HasChefAttributes

`func (o *TaskResponseTaskOptions) HasChefAttributes() bool`

HasChefAttributes returns a boolean if a field has been set.

### GetConditionalScript

`func (o *TaskResponseTaskOptions) GetConditionalScript() string`

GetConditionalScript returns the ConditionalScript field if non-nil, zero value otherwise.

### GetConditionalScriptOk

`func (o *TaskResponseTaskOptions) GetConditionalScriptOk() (*string, bool)`

GetConditionalScriptOk returns a tuple with the ConditionalScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalScript

`func (o *TaskResponseTaskOptions) SetConditionalScript(v string)`

SetConditionalScript sets ConditionalScript field to given value.

### HasConditionalScript

`func (o *TaskResponseTaskOptions) HasConditionalScript() bool`

HasConditionalScript returns a boolean if a field has been set.

### GetIfOperationalWorkflowId

`func (o *TaskResponseTaskOptions) GetIfOperationalWorkflowId() int64`

GetIfOperationalWorkflowId returns the IfOperationalWorkflowId field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowIdOk

`func (o *TaskResponseTaskOptions) GetIfOperationalWorkflowIdOk() (*int64, bool)`

GetIfOperationalWorkflowIdOk returns a tuple with the IfOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowId

`func (o *TaskResponseTaskOptions) SetIfOperationalWorkflowId(v int64)`

SetIfOperationalWorkflowId sets IfOperationalWorkflowId field to given value.

### HasIfOperationalWorkflowId

`func (o *TaskResponseTaskOptions) HasIfOperationalWorkflowId() bool`

HasIfOperationalWorkflowId returns a boolean if a field has been set.

### GetIfOperationalWorkflowName

`func (o *TaskResponseTaskOptions) GetIfOperationalWorkflowName() string`

GetIfOperationalWorkflowName returns the IfOperationalWorkflowName field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowNameOk

`func (o *TaskResponseTaskOptions) GetIfOperationalWorkflowNameOk() (*string, bool)`

GetIfOperationalWorkflowNameOk returns a tuple with the IfOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowName

`func (o *TaskResponseTaskOptions) SetIfOperationalWorkflowName(v string)`

SetIfOperationalWorkflowName sets IfOperationalWorkflowName field to given value.

### HasIfOperationalWorkflowName

`func (o *TaskResponseTaskOptions) HasIfOperationalWorkflowName() bool`

HasIfOperationalWorkflowName returns a boolean if a field has been set.

### GetElseOperationalWorkflowId

`func (o *TaskResponseTaskOptions) GetElseOperationalWorkflowId() int64`

GetElseOperationalWorkflowId returns the ElseOperationalWorkflowId field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowIdOk

`func (o *TaskResponseTaskOptions) GetElseOperationalWorkflowIdOk() (*int64, bool)`

GetElseOperationalWorkflowIdOk returns a tuple with the ElseOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowId

`func (o *TaskResponseTaskOptions) SetElseOperationalWorkflowId(v int64)`

SetElseOperationalWorkflowId sets ElseOperationalWorkflowId field to given value.

### HasElseOperationalWorkflowId

`func (o *TaskResponseTaskOptions) HasElseOperationalWorkflowId() bool`

HasElseOperationalWorkflowId returns a boolean if a field has been set.

### GetElseOperationalWorkflowName

`func (o *TaskResponseTaskOptions) GetElseOperationalWorkflowName() string`

GetElseOperationalWorkflowName returns the ElseOperationalWorkflowName field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowNameOk

`func (o *TaskResponseTaskOptions) GetElseOperationalWorkflowNameOk() (*string, bool)`

GetElseOperationalWorkflowNameOk returns a tuple with the ElseOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowName

`func (o *TaskResponseTaskOptions) SetElseOperationalWorkflowName(v string)`

SetElseOperationalWorkflowName sets ElseOperationalWorkflowName field to given value.

### HasElseOperationalWorkflowName

`func (o *TaskResponseTaskOptions) HasElseOperationalWorkflowName() bool`

HasElseOperationalWorkflowName returns a boolean if a field has been set.

### GetEmailSkipTemplate

`func (o *TaskResponseTaskOptions) GetEmailSkipTemplate() string`

GetEmailSkipTemplate returns the EmailSkipTemplate field if non-nil, zero value otherwise.

### GetEmailSkipTemplateOk

`func (o *TaskResponseTaskOptions) GetEmailSkipTemplateOk() (*string, bool)`

GetEmailSkipTemplateOk returns a tuple with the EmailSkipTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSkipTemplate

`func (o *TaskResponseTaskOptions) SetEmailSkipTemplate(v string)`

SetEmailSkipTemplate sets EmailSkipTemplate field to given value.

### HasEmailSkipTemplate

`func (o *TaskResponseTaskOptions) HasEmailSkipTemplate() bool`

HasEmailSkipTemplate returns a boolean if a field has been set.

### GetEmailSubject

`func (o *TaskResponseTaskOptions) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *TaskResponseTaskOptions) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *TaskResponseTaskOptions) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *TaskResponseTaskOptions) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailAddress

`func (o *TaskResponseTaskOptions) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *TaskResponseTaskOptions) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *TaskResponseTaskOptions) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *TaskResponseTaskOptions) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetWebPassword

`func (o *TaskResponseTaskOptions) GetWebPassword() string`

GetWebPassword returns the WebPassword field if non-nil, zero value otherwise.

### GetWebPasswordOk

`func (o *TaskResponseTaskOptions) GetWebPasswordOk() (*string, bool)`

GetWebPasswordOk returns a tuple with the WebPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPassword

`func (o *TaskResponseTaskOptions) SetWebPassword(v string)`

SetWebPassword sets WebPassword field to given value.

### HasWebPassword

`func (o *TaskResponseTaskOptions) HasWebPassword() bool`

HasWebPassword returns a boolean if a field has been set.

### GetWebPasswordHash

`func (o *TaskResponseTaskOptions) GetWebPasswordHash() string`

GetWebPasswordHash returns the WebPasswordHash field if non-nil, zero value otherwise.

### GetWebPasswordHashOk

`func (o *TaskResponseTaskOptions) GetWebPasswordHashOk() (*string, bool)`

GetWebPasswordHashOk returns a tuple with the WebPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPasswordHash

`func (o *TaskResponseTaskOptions) SetWebPasswordHash(v string)`

SetWebPasswordHash sets WebPasswordHash field to given value.

### HasWebPasswordHash

`func (o *TaskResponseTaskOptions) HasWebPasswordHash() bool`

HasWebPasswordHash returns a boolean if a field has been set.

### GetWebUser

`func (o *TaskResponseTaskOptions) GetWebUser() string`

GetWebUser returns the WebUser field if non-nil, zero value otherwise.

### GetWebUserOk

`func (o *TaskResponseTaskOptions) GetWebUserOk() (*string, bool)`

GetWebUserOk returns a tuple with the WebUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUser

`func (o *TaskResponseTaskOptions) SetWebUser(v string)`

SetWebUser sets WebUser field to given value.

### HasWebUser

`func (o *TaskResponseTaskOptions) HasWebUser() bool`

HasWebUser returns a boolean if a field has been set.

### GetWebBody

`func (o *TaskResponseTaskOptions) GetWebBody() string`

GetWebBody returns the WebBody field if non-nil, zero value otherwise.

### GetWebBodyOk

`func (o *TaskResponseTaskOptions) GetWebBodyOk() (*string, bool)`

GetWebBodyOk returns a tuple with the WebBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBody

`func (o *TaskResponseTaskOptions) SetWebBody(v string)`

SetWebBody sets WebBody field to given value.

### HasWebBody

`func (o *TaskResponseTaskOptions) HasWebBody() bool`

HasWebBody returns a boolean if a field has been set.

### GetWebHeaders

`func (o *TaskResponseTaskOptions) GetWebHeaders() string`

GetWebHeaders returns the WebHeaders field if non-nil, zero value otherwise.

### GetWebHeadersOk

`func (o *TaskResponseTaskOptions) GetWebHeadersOk() (*string, bool)`

GetWebHeadersOk returns a tuple with the WebHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHeaders

`func (o *TaskResponseTaskOptions) SetWebHeaders(v string)`

SetWebHeaders sets WebHeaders field to given value.

### HasWebHeaders

`func (o *TaskResponseTaskOptions) HasWebHeaders() bool`

HasWebHeaders returns a boolean if a field has been set.

### GetIgnoreSSL

`func (o *TaskResponseTaskOptions) GetIgnoreSSL() string`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *TaskResponseTaskOptions) GetIgnoreSSLOk() (*string, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *TaskResponseTaskOptions) SetIgnoreSSL(v string)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *TaskResponseTaskOptions) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetWebMethod

`func (o *TaskResponseTaskOptions) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *TaskResponseTaskOptions) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *TaskResponseTaskOptions) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.

### HasWebMethod

`func (o *TaskResponseTaskOptions) HasWebMethod() bool`

HasWebMethod returns a boolean if a field has been set.

### GetWebUrl

`func (o *TaskResponseTaskOptions) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *TaskResponseTaskOptions) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *TaskResponseTaskOptions) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *TaskResponseTaskOptions) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetJsScript

`func (o *TaskResponseTaskOptions) GetJsScript() string`

GetJsScript returns the JsScript field if non-nil, zero value otherwise.

### GetJsScriptOk

`func (o *TaskResponseTaskOptions) GetJsScriptOk() (*string, bool)`

GetJsScriptOk returns a tuple with the JsScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsScript

`func (o *TaskResponseTaskOptions) SetJsScript(v string)`

SetJsScript sets JsScript field to given value.

### HasJsScript

`func (o *TaskResponseTaskOptions) HasJsScript() bool`

HasJsScript returns a boolean if a field has been set.

### GetContainerScriptId

`func (o *TaskResponseTaskOptions) GetContainerScriptId() string`

GetContainerScriptId returns the ContainerScriptId field if non-nil, zero value otherwise.

### GetContainerScriptIdOk

`func (o *TaskResponseTaskOptions) GetContainerScriptIdOk() (*string, bool)`

GetContainerScriptIdOk returns a tuple with the ContainerScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScriptId

`func (o *TaskResponseTaskOptions) SetContainerScriptId(v string)`

SetContainerScriptId sets ContainerScriptId field to given value.

### HasContainerScriptId

`func (o *TaskResponseTaskOptions) HasContainerScriptId() bool`

HasContainerScriptId returns a boolean if a field has been set.

### GetContainerScript

`func (o *TaskResponseTaskOptions) GetContainerScript() string`

GetContainerScript returns the ContainerScript field if non-nil, zero value otherwise.

### GetContainerScriptOk

`func (o *TaskResponseTaskOptions) GetContainerScriptOk() (*string, bool)`

GetContainerScriptOk returns a tuple with the ContainerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScript

`func (o *TaskResponseTaskOptions) SetContainerScript(v string)`

SetContainerScript sets ContainerScript field to given value.

### HasContainerScript

`func (o *TaskResponseTaskOptions) HasContainerScript() bool`

HasContainerScript returns a boolean if a field has been set.

### GetContainerTemplateId

`func (o *TaskResponseTaskOptions) GetContainerTemplateId() string`

GetContainerTemplateId returns the ContainerTemplateId field if non-nil, zero value otherwise.

### GetContainerTemplateIdOk

`func (o *TaskResponseTaskOptions) GetContainerTemplateIdOk() (*string, bool)`

GetContainerTemplateIdOk returns a tuple with the ContainerTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplateId

`func (o *TaskResponseTaskOptions) SetContainerTemplateId(v string)`

SetContainerTemplateId sets ContainerTemplateId field to given value.

### HasContainerTemplateId

`func (o *TaskResponseTaskOptions) HasContainerTemplateId() bool`

HasContainerTemplateId returns a boolean if a field has been set.

### GetContainerTemplate

`func (o *TaskResponseTaskOptions) GetContainerTemplate() string`

GetContainerTemplate returns the ContainerTemplate field if non-nil, zero value otherwise.

### GetContainerTemplateOk

`func (o *TaskResponseTaskOptions) GetContainerTemplateOk() (*string, bool)`

GetContainerTemplateOk returns a tuple with the ContainerTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplate

`func (o *TaskResponseTaskOptions) SetContainerTemplate(v string)`

SetContainerTemplate sets ContainerTemplate field to given value.

### HasContainerTemplate

`func (o *TaskResponseTaskOptions) HasContainerTemplate() bool`

HasContainerTemplate returns a boolean if a field has been set.

### GetOperationalWorkflowId

`func (o *TaskResponseTaskOptions) GetOperationalWorkflowId() string`

GetOperationalWorkflowId returns the OperationalWorkflowId field if non-nil, zero value otherwise.

### GetOperationalWorkflowIdOk

`func (o *TaskResponseTaskOptions) GetOperationalWorkflowIdOk() (*string, bool)`

GetOperationalWorkflowIdOk returns a tuple with the OperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowId

`func (o *TaskResponseTaskOptions) SetOperationalWorkflowId(v string)`

SetOperationalWorkflowId sets OperationalWorkflowId field to given value.


### GetOperationalWorkflowName

`func (o *TaskResponseTaskOptions) GetOperationalWorkflowName() string`

GetOperationalWorkflowName returns the OperationalWorkflowName field if non-nil, zero value otherwise.

### GetOperationalWorkflowNameOk

`func (o *TaskResponseTaskOptions) GetOperationalWorkflowNameOk() (*string, bool)`

GetOperationalWorkflowNameOk returns a tuple with the OperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowName

`func (o *TaskResponseTaskOptions) SetOperationalWorkflowName(v string)`

SetOperationalWorkflowName sets OperationalWorkflowName field to given value.

### HasOperationalWorkflowName

`func (o *TaskResponseTaskOptions) HasOperationalWorkflowName() bool`

HasOperationalWorkflowName returns a boolean if a field has been set.

### GetPuppetEnvironment

`func (o *TaskResponseTaskOptions) GetPuppetEnvironment() string`

GetPuppetEnvironment returns the PuppetEnvironment field if non-nil, zero value otherwise.

### GetPuppetEnvironmentOk

`func (o *TaskResponseTaskOptions) GetPuppetEnvironmentOk() (*string, bool)`

GetPuppetEnvironmentOk returns a tuple with the PuppetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetEnvironment

`func (o *TaskResponseTaskOptions) SetPuppetEnvironment(v string)`

SetPuppetEnvironment sets PuppetEnvironment field to given value.

### HasPuppetEnvironment

`func (o *TaskResponseTaskOptions) HasPuppetEnvironment() bool`

HasPuppetEnvironment returns a boolean if a field has been set.

### GetPuppetNodeName

`func (o *TaskResponseTaskOptions) GetPuppetNodeName() string`

GetPuppetNodeName returns the PuppetNodeName field if non-nil, zero value otherwise.

### GetPuppetNodeNameOk

`func (o *TaskResponseTaskOptions) GetPuppetNodeNameOk() (*string, bool)`

GetPuppetNodeNameOk returns a tuple with the PuppetNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetNodeName

`func (o *TaskResponseTaskOptions) SetPuppetNodeName(v string)`

SetPuppetNodeName sets PuppetNodeName field to given value.

### HasPuppetNodeName

`func (o *TaskResponseTaskOptions) HasPuppetNodeName() bool`

HasPuppetNodeName returns a boolean if a field has been set.

### GetPythonArgs

`func (o *TaskResponseTaskOptions) GetPythonArgs() string`

GetPythonArgs returns the PythonArgs field if non-nil, zero value otherwise.

### GetPythonArgsOk

`func (o *TaskResponseTaskOptions) GetPythonArgsOk() (*string, bool)`

GetPythonArgsOk returns a tuple with the PythonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonArgs

`func (o *TaskResponseTaskOptions) SetPythonArgs(v string)`

SetPythonArgs sets PythonArgs field to given value.

### HasPythonArgs

`func (o *TaskResponseTaskOptions) HasPythonArgs() bool`

HasPythonArgs returns a boolean if a field has been set.

### GetPythonBinary

`func (o *TaskResponseTaskOptions) GetPythonBinary() string`

GetPythonBinary returns the PythonBinary field if non-nil, zero value otherwise.

### GetPythonBinaryOk

`func (o *TaskResponseTaskOptions) GetPythonBinaryOk() (*string, bool)`

GetPythonBinaryOk returns a tuple with the PythonBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonBinary

`func (o *TaskResponseTaskOptions) SetPythonBinary(v string)`

SetPythonBinary sets PythonBinary field to given value.

### HasPythonBinary

`func (o *TaskResponseTaskOptions) HasPythonBinary() bool`

HasPythonBinary returns a boolean if a field has been set.

### GetPythonAdditionalPackages

`func (o *TaskResponseTaskOptions) GetPythonAdditionalPackages() string`

GetPythonAdditionalPackages returns the PythonAdditionalPackages field if non-nil, zero value otherwise.

### GetPythonAdditionalPackagesOk

`func (o *TaskResponseTaskOptions) GetPythonAdditionalPackagesOk() (*string, bool)`

GetPythonAdditionalPackagesOk returns a tuple with the PythonAdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonAdditionalPackages

`func (o *TaskResponseTaskOptions) SetPythonAdditionalPackages(v string)`

SetPythonAdditionalPackages sets PythonAdditionalPackages field to given value.

### HasPythonAdditionalPackages

`func (o *TaskResponseTaskOptions) HasPythonAdditionalPackages() bool`

HasPythonAdditionalPackages returns a boolean if a field has been set.

### GetVroBody

`func (o *TaskResponseTaskOptions) GetVroBody() string`

GetVroBody returns the VroBody field if non-nil, zero value otherwise.

### GetVroBodyOk

`func (o *TaskResponseTaskOptions) GetVroBodyOk() (*string, bool)`

GetVroBodyOk returns a tuple with the VroBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVroBody

`func (o *TaskResponseTaskOptions) SetVroBody(v string)`

SetVroBody sets VroBody field to given value.

### HasVroBody

`func (o *TaskResponseTaskOptions) HasVroBody() bool`

HasVroBody returns a boolean if a field has been set.

### GetWriteAttributesAttributes

`func (o *TaskResponseTaskOptions) GetWriteAttributesAttributes() string`

GetWriteAttributesAttributes returns the WriteAttributesAttributes field if non-nil, zero value otherwise.

### GetWriteAttributesAttributesOk

`func (o *TaskResponseTaskOptions) GetWriteAttributesAttributesOk() (*string, bool)`

GetWriteAttributesAttributesOk returns a tuple with the WriteAttributesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAttributesAttributes

`func (o *TaskResponseTaskOptions) SetWriteAttributesAttributes(v string)`

SetWriteAttributesAttributes sets WriteAttributesAttributes field to given value.

### HasWriteAttributesAttributes

`func (o *TaskResponseTaskOptions) HasWriteAttributesAttributes() bool`

HasWriteAttributesAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


