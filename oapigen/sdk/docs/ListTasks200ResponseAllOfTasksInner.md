# ListTasks200ResponseAllOfTasksInner

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

### NewListTasks200ResponseAllOfTasksInner

`func NewListTasks200ResponseAllOfTasksInner(operationalWorkflowId string, ) *ListTasks200ResponseAllOfTasksInner`

NewListTasks200ResponseAllOfTasksInner instantiates a new ListTasks200ResponseAllOfTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseAllOfTasksInnerWithDefaults

`func NewListTasks200ResponseAllOfTasksInnerWithDefaults() *ListTasks200ResponseAllOfTasksInner`

NewListTasks200ResponseAllOfTasksInnerWithDefaults instantiates a new ListTasks200ResponseAllOfTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsibleOptions

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleOptions() string`

GetAnsibleOptions returns the AnsibleOptions field if non-nil, zero value otherwise.

### GetAnsibleOptionsOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleOptionsOk() (*string, bool)`

GetAnsibleOptionsOk returns a tuple with the AnsibleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleOptions

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleOptions(v string)`

SetAnsibleOptions sets AnsibleOptions field to given value.

### HasAnsibleOptions

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleOptions() bool`

HasAnsibleOptions returns a boolean if a field has been set.

### GetAnsiblePlaybook

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsiblePlaybook() string`

GetAnsiblePlaybook returns the AnsiblePlaybook field if non-nil, zero value otherwise.

### GetAnsiblePlaybookOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsiblePlaybookOk() (*string, bool)`

GetAnsiblePlaybookOk returns a tuple with the AnsiblePlaybook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePlaybook

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsiblePlaybook(v string)`

SetAnsiblePlaybook sets AnsiblePlaybook field to given value.

### HasAnsiblePlaybook

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsiblePlaybook() bool`

HasAnsiblePlaybook returns a boolean if a field has been set.

### GetSshKey

`func (o *ListTasks200ResponseAllOfTasksInner) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *ListTasks200ResponseAllOfTasksInner) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *ListTasks200ResponseAllOfTasksInner) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetPort

`func (o *ListTasks200ResponseAllOfTasksInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListTasks200ResponseAllOfTasksInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ListTasks200ResponseAllOfTasksInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLocalScriptGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### GetPassword

`func (o *ListTasks200ResponseAllOfTasksInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ListTasks200ResponseAllOfTasksInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ListTasks200ResponseAllOfTasksInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *ListTasks200ResponseAllOfTasksInner) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ListTasks200ResponseAllOfTasksInner) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ListTasks200ResponseAllOfTasksInner) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetLocalScriptGitId

`func (o *ListTasks200ResponseAllOfTasksInner) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *ListTasks200ResponseAllOfTasksInner) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *ListTasks200ResponseAllOfTasksInner) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### GetAnsibleGitId

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleGitId() string`

GetAnsibleGitId returns the AnsibleGitId field if non-nil, zero value otherwise.

### GetAnsibleGitIdOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleGitIdOk() (*string, bool)`

GetAnsibleGitIdOk returns a tuple with the AnsibleGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitId

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleGitId(v string)`

SetAnsibleGitId sets AnsibleGitId field to given value.

### HasAnsibleGitId

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleGitId() bool`

HasAnsibleGitId returns a boolean if a field has been set.

### GetHost

`func (o *ListTasks200ResponseAllOfTasksInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ListTasks200ResponseAllOfTasksInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ListTasks200ResponseAllOfTasksInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAnsibleSkipTags

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleSkipTags() string`

GetAnsibleSkipTags returns the AnsibleSkipTags field if non-nil, zero value otherwise.

### GetAnsibleSkipTagsOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleSkipTagsOk() (*string, bool)`

GetAnsibleSkipTagsOk returns a tuple with the AnsibleSkipTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleSkipTags

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleSkipTags(v string)`

SetAnsibleSkipTags sets AnsibleSkipTags field to given value.

### HasAnsibleSkipTags

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleSkipTags() bool`

HasAnsibleSkipTags returns a boolean if a field has been set.

### GetAnsibleTags

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleTags() string`

GetAnsibleTags returns the AnsibleTags field if non-nil, zero value otherwise.

### GetAnsibleTagsOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleTagsOk() (*string, bool)`

GetAnsibleTagsOk returns a tuple with the AnsibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTags

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleTags(v string)`

SetAnsibleTags sets AnsibleTags field to given value.

### HasAnsibleTags

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleTags() bool`

HasAnsibleTags returns a boolean if a field has been set.

### GetUsername

`func (o *ListTasks200ResponseAllOfTasksInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListTasks200ResponseAllOfTasksInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListTasks200ResponseAllOfTasksInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAnsibleGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleGitRef() string`

GetAnsibleGitRef returns the AnsibleGitRef field if non-nil, zero value otherwise.

### GetAnsibleGitRefOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleGitRefOk() (*string, bool)`

GetAnsibleGitRefOk returns a tuple with the AnsibleGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleGitRef(v string)`

SetAnsibleGitRef sets AnsibleGitRef field to given value.

### HasAnsibleGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleGitRef() bool`

HasAnsibleGitRef returns a boolean if a field has been set.

### GetAnsibleTowerGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleTowerGitRef() string`

GetAnsibleTowerGitRef returns the AnsibleTowerGitRef field if non-nil, zero value otherwise.

### GetAnsibleTowerGitRefOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleTowerGitRefOk() (*string, bool)`

GetAnsibleTowerGitRefOk returns a tuple with the AnsibleTowerGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleTowerGitRef(v string)`

SetAnsibleTowerGitRef sets AnsibleTowerGitRef field to given value.

### HasAnsibleTowerGitRef

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleTowerGitRef() bool`

HasAnsibleTowerGitRef returns a boolean if a field has been set.

### GetAnsibleGroup

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleGroup() string`

GetAnsibleGroup returns the AnsibleGroup field if non-nil, zero value otherwise.

### GetAnsibleGroupOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleGroupOk() (*string, bool)`

GetAnsibleGroupOk returns a tuple with the AnsibleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleGroup

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleGroup(v string)`

SetAnsibleGroup sets AnsibleGroup field to given value.

### HasAnsibleGroup

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleGroup() bool`

HasAnsibleGroup returns a boolean if a field has been set.

### GetAnsibleTowerExecuteMode

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleTowerExecuteMode() string`

GetAnsibleTowerExecuteMode returns the AnsibleTowerExecuteMode field if non-nil, zero value otherwise.

### GetAnsibleTowerExecuteModeOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetAnsibleTowerExecuteModeOk() (*string, bool)`

GetAnsibleTowerExecuteModeOk returns a tuple with the AnsibleTowerExecuteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleTowerExecuteMode

`func (o *ListTasks200ResponseAllOfTasksInner) SetAnsibleTowerExecuteMode(v string)`

SetAnsibleTowerExecuteMode sets AnsibleTowerExecuteMode field to given value.

### HasAnsibleTowerExecuteMode

`func (o *ListTasks200ResponseAllOfTasksInner) HasAnsibleTowerExecuteMode() bool`

HasAnsibleTowerExecuteMode returns a boolean if a field has been set.

### GetChefDataKey

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefDataKey() string`

GetChefDataKey returns the ChefDataKey field if non-nil, zero value otherwise.

### GetChefDataKeyOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefDataKeyOk() (*string, bool)`

GetChefDataKeyOk returns a tuple with the ChefDataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKey

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefDataKey(v string)`

SetChefDataKey sets ChefDataKey field to given value.

### HasChefDataKey

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefDataKey() bool`

HasChefDataKey returns a boolean if a field has been set.

### GetChefDataKeyHash

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefDataKeyHash() string`

GetChefDataKeyHash returns the ChefDataKeyHash field if non-nil, zero value otherwise.

### GetChefDataKeyHashOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefDataKeyHashOk() (*string, bool)`

GetChefDataKeyHashOk returns a tuple with the ChefDataKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyHash

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefDataKeyHash(v string)`

SetChefDataKeyHash sets ChefDataKeyHash field to given value.

### HasChefDataKeyHash

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefDataKeyHash() bool`

HasChefDataKeyHash returns a boolean if a field has been set.

### GetChefRunList

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefRunList() string`

GetChefRunList returns the ChefRunList field if non-nil, zero value otherwise.

### GetChefRunListOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefRunListOk() (*string, bool)`

GetChefRunListOk returns a tuple with the ChefRunList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefRunList

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefRunList(v string)`

SetChefRunList sets ChefRunList field to given value.

### HasChefRunList

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefRunList() bool`

HasChefRunList returns a boolean if a field has been set.

### GetChefDataKeyPath

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefDataKeyPath() string`

GetChefDataKeyPath returns the ChefDataKeyPath field if non-nil, zero value otherwise.

### GetChefDataKeyPathOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefDataKeyPathOk() (*string, bool)`

GetChefDataKeyPathOk returns a tuple with the ChefDataKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefDataKeyPath

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefDataKeyPath(v string)`

SetChefDataKeyPath sets ChefDataKeyPath field to given value.

### HasChefDataKeyPath

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefDataKeyPath() bool`

HasChefDataKeyPath returns a boolean if a field has been set.

### GetChefEnv

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefEnv() string`

GetChefEnv returns the ChefEnv field if non-nil, zero value otherwise.

### GetChefEnvOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefEnvOk() (*string, bool)`

GetChefEnvOk returns a tuple with the ChefEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefEnv

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefEnv(v string)`

SetChefEnv sets ChefEnv field to given value.

### HasChefEnv

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefEnv() bool`

HasChefEnv returns a boolean if a field has been set.

### GetChefNodeName

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefNodeName() string`

GetChefNodeName returns the ChefNodeName field if non-nil, zero value otherwise.

### GetChefNodeNameOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefNodeNameOk() (*string, bool)`

GetChefNodeNameOk returns a tuple with the ChefNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefNodeName

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefNodeName(v string)`

SetChefNodeName sets ChefNodeName field to given value.

### HasChefNodeName

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefNodeName() bool`

HasChefNodeName returns a boolean if a field has been set.

### GetChefAttributes

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefAttributes() string`

GetChefAttributes returns the ChefAttributes field if non-nil, zero value otherwise.

### GetChefAttributesOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetChefAttributesOk() (*string, bool)`

GetChefAttributesOk returns a tuple with the ChefAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefAttributes

`func (o *ListTasks200ResponseAllOfTasksInner) SetChefAttributes(v string)`

SetChefAttributes sets ChefAttributes field to given value.

### HasChefAttributes

`func (o *ListTasks200ResponseAllOfTasksInner) HasChefAttributes() bool`

HasChefAttributes returns a boolean if a field has been set.

### GetEmailSkipTemplate

`func (o *ListTasks200ResponseAllOfTasksInner) GetEmailSkipTemplate() string`

GetEmailSkipTemplate returns the EmailSkipTemplate field if non-nil, zero value otherwise.

### GetEmailSkipTemplateOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetEmailSkipTemplateOk() (*string, bool)`

GetEmailSkipTemplateOk returns a tuple with the EmailSkipTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSkipTemplate

`func (o *ListTasks200ResponseAllOfTasksInner) SetEmailSkipTemplate(v string)`

SetEmailSkipTemplate sets EmailSkipTemplate field to given value.

### HasEmailSkipTemplate

`func (o *ListTasks200ResponseAllOfTasksInner) HasEmailSkipTemplate() bool`

HasEmailSkipTemplate returns a boolean if a field has been set.

### GetEmailSubject

`func (o *ListTasks200ResponseAllOfTasksInner) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *ListTasks200ResponseAllOfTasksInner) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *ListTasks200ResponseAllOfTasksInner) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ListTasks200ResponseAllOfTasksInner) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ListTasks200ResponseAllOfTasksInner) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ListTasks200ResponseAllOfTasksInner) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetWebPassword

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebPassword() string`

GetWebPassword returns the WebPassword field if non-nil, zero value otherwise.

### GetWebPasswordOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebPasswordOk() (*string, bool)`

GetWebPasswordOk returns a tuple with the WebPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPassword

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebPassword(v string)`

SetWebPassword sets WebPassword field to given value.

### HasWebPassword

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebPassword() bool`

HasWebPassword returns a boolean if a field has been set.

### GetWebPasswordHash

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebPasswordHash() string`

GetWebPasswordHash returns the WebPasswordHash field if non-nil, zero value otherwise.

### GetWebPasswordHashOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebPasswordHashOk() (*string, bool)`

GetWebPasswordHashOk returns a tuple with the WebPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPasswordHash

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebPasswordHash(v string)`

SetWebPasswordHash sets WebPasswordHash field to given value.

### HasWebPasswordHash

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebPasswordHash() bool`

HasWebPasswordHash returns a boolean if a field has been set.

### GetWebUser

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebUser() string`

GetWebUser returns the WebUser field if non-nil, zero value otherwise.

### GetWebUserOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebUserOk() (*string, bool)`

GetWebUserOk returns a tuple with the WebUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUser

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebUser(v string)`

SetWebUser sets WebUser field to given value.

### HasWebUser

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebUser() bool`

HasWebUser returns a boolean if a field has been set.

### GetWebBody

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebBody() string`

GetWebBody returns the WebBody field if non-nil, zero value otherwise.

### GetWebBodyOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebBodyOk() (*string, bool)`

GetWebBodyOk returns a tuple with the WebBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBody

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebBody(v string)`

SetWebBody sets WebBody field to given value.

### HasWebBody

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebBody() bool`

HasWebBody returns a boolean if a field has been set.

### GetWebHeaders

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebHeaders() string`

GetWebHeaders returns the WebHeaders field if non-nil, zero value otherwise.

### GetWebHeadersOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebHeadersOk() (*string, bool)`

GetWebHeadersOk returns a tuple with the WebHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHeaders

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebHeaders(v string)`

SetWebHeaders sets WebHeaders field to given value.

### HasWebHeaders

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebHeaders() bool`

HasWebHeaders returns a boolean if a field has been set.

### GetIgnoreSSL

`func (o *ListTasks200ResponseAllOfTasksInner) GetIgnoreSSL() string`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetIgnoreSSLOk() (*string, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *ListTasks200ResponseAllOfTasksInner) SetIgnoreSSL(v string)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *ListTasks200ResponseAllOfTasksInner) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetWebMethod

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.

### HasWebMethod

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebMethod() bool`

HasWebMethod returns a boolean if a field has been set.

### GetWebUrl

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ListTasks200ResponseAllOfTasksInner) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ListTasks200ResponseAllOfTasksInner) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetJsScript

`func (o *ListTasks200ResponseAllOfTasksInner) GetJsScript() string`

GetJsScript returns the JsScript field if non-nil, zero value otherwise.

### GetJsScriptOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetJsScriptOk() (*string, bool)`

GetJsScriptOk returns a tuple with the JsScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsScript

`func (o *ListTasks200ResponseAllOfTasksInner) SetJsScript(v string)`

SetJsScript sets JsScript field to given value.

### HasJsScript

`func (o *ListTasks200ResponseAllOfTasksInner) HasJsScript() bool`

HasJsScript returns a boolean if a field has been set.

### GetContainerScriptId

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerScriptId() string`

GetContainerScriptId returns the ContainerScriptId field if non-nil, zero value otherwise.

### GetContainerScriptIdOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerScriptIdOk() (*string, bool)`

GetContainerScriptIdOk returns a tuple with the ContainerScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScriptId

`func (o *ListTasks200ResponseAllOfTasksInner) SetContainerScriptId(v string)`

SetContainerScriptId sets ContainerScriptId field to given value.

### HasContainerScriptId

`func (o *ListTasks200ResponseAllOfTasksInner) HasContainerScriptId() bool`

HasContainerScriptId returns a boolean if a field has been set.

### GetContainerScript

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerScript() string`

GetContainerScript returns the ContainerScript field if non-nil, zero value otherwise.

### GetContainerScriptOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerScriptOk() (*string, bool)`

GetContainerScriptOk returns a tuple with the ContainerScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScript

`func (o *ListTasks200ResponseAllOfTasksInner) SetContainerScript(v string)`

SetContainerScript sets ContainerScript field to given value.

### HasContainerScript

`func (o *ListTasks200ResponseAllOfTasksInner) HasContainerScript() bool`

HasContainerScript returns a boolean if a field has been set.

### GetContainerTemplateId

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerTemplateId() string`

GetContainerTemplateId returns the ContainerTemplateId field if non-nil, zero value otherwise.

### GetContainerTemplateIdOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerTemplateIdOk() (*string, bool)`

GetContainerTemplateIdOk returns a tuple with the ContainerTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplateId

`func (o *ListTasks200ResponseAllOfTasksInner) SetContainerTemplateId(v string)`

SetContainerTemplateId sets ContainerTemplateId field to given value.

### HasContainerTemplateId

`func (o *ListTasks200ResponseAllOfTasksInner) HasContainerTemplateId() bool`

HasContainerTemplateId returns a boolean if a field has been set.

### GetContainerTemplate

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerTemplate() string`

GetContainerTemplate returns the ContainerTemplate field if non-nil, zero value otherwise.

### GetContainerTemplateOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetContainerTemplateOk() (*string, bool)`

GetContainerTemplateOk returns a tuple with the ContainerTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplate

`func (o *ListTasks200ResponseAllOfTasksInner) SetContainerTemplate(v string)`

SetContainerTemplate sets ContainerTemplate field to given value.

### HasContainerTemplate

`func (o *ListTasks200ResponseAllOfTasksInner) HasContainerTemplate() bool`

HasContainerTemplate returns a boolean if a field has been set.

### GetOperationalWorkflowId

`func (o *ListTasks200ResponseAllOfTasksInner) GetOperationalWorkflowId() string`

GetOperationalWorkflowId returns the OperationalWorkflowId field if non-nil, zero value otherwise.

### GetOperationalWorkflowIdOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetOperationalWorkflowIdOk() (*string, bool)`

GetOperationalWorkflowIdOk returns a tuple with the OperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowId

`func (o *ListTasks200ResponseAllOfTasksInner) SetOperationalWorkflowId(v string)`

SetOperationalWorkflowId sets OperationalWorkflowId field to given value.


### GetOperationalWorkflowName

`func (o *ListTasks200ResponseAllOfTasksInner) GetOperationalWorkflowName() string`

GetOperationalWorkflowName returns the OperationalWorkflowName field if non-nil, zero value otherwise.

### GetOperationalWorkflowNameOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetOperationalWorkflowNameOk() (*string, bool)`

GetOperationalWorkflowNameOk returns a tuple with the OperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalWorkflowName

`func (o *ListTasks200ResponseAllOfTasksInner) SetOperationalWorkflowName(v string)`

SetOperationalWorkflowName sets OperationalWorkflowName field to given value.

### HasOperationalWorkflowName

`func (o *ListTasks200ResponseAllOfTasksInner) HasOperationalWorkflowName() bool`

HasOperationalWorkflowName returns a boolean if a field has been set.

### GetPuppetEnvironment

`func (o *ListTasks200ResponseAllOfTasksInner) GetPuppetEnvironment() string`

GetPuppetEnvironment returns the PuppetEnvironment field if non-nil, zero value otherwise.

### GetPuppetEnvironmentOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPuppetEnvironmentOk() (*string, bool)`

GetPuppetEnvironmentOk returns a tuple with the PuppetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetEnvironment

`func (o *ListTasks200ResponseAllOfTasksInner) SetPuppetEnvironment(v string)`

SetPuppetEnvironment sets PuppetEnvironment field to given value.

### HasPuppetEnvironment

`func (o *ListTasks200ResponseAllOfTasksInner) HasPuppetEnvironment() bool`

HasPuppetEnvironment returns a boolean if a field has been set.

### GetPuppetNodeName

`func (o *ListTasks200ResponseAllOfTasksInner) GetPuppetNodeName() string`

GetPuppetNodeName returns the PuppetNodeName field if non-nil, zero value otherwise.

### GetPuppetNodeNameOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPuppetNodeNameOk() (*string, bool)`

GetPuppetNodeNameOk returns a tuple with the PuppetNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetNodeName

`func (o *ListTasks200ResponseAllOfTasksInner) SetPuppetNodeName(v string)`

SetPuppetNodeName sets PuppetNodeName field to given value.

### HasPuppetNodeName

`func (o *ListTasks200ResponseAllOfTasksInner) HasPuppetNodeName() bool`

HasPuppetNodeName returns a boolean if a field has been set.

### GetPythonArgs

`func (o *ListTasks200ResponseAllOfTasksInner) GetPythonArgs() string`

GetPythonArgs returns the PythonArgs field if non-nil, zero value otherwise.

### GetPythonArgsOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPythonArgsOk() (*string, bool)`

GetPythonArgsOk returns a tuple with the PythonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonArgs

`func (o *ListTasks200ResponseAllOfTasksInner) SetPythonArgs(v string)`

SetPythonArgs sets PythonArgs field to given value.

### HasPythonArgs

`func (o *ListTasks200ResponseAllOfTasksInner) HasPythonArgs() bool`

HasPythonArgs returns a boolean if a field has been set.

### GetPythonBinary

`func (o *ListTasks200ResponseAllOfTasksInner) GetPythonBinary() string`

GetPythonBinary returns the PythonBinary field if non-nil, zero value otherwise.

### GetPythonBinaryOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPythonBinaryOk() (*string, bool)`

GetPythonBinaryOk returns a tuple with the PythonBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonBinary

`func (o *ListTasks200ResponseAllOfTasksInner) SetPythonBinary(v string)`

SetPythonBinary sets PythonBinary field to given value.

### HasPythonBinary

`func (o *ListTasks200ResponseAllOfTasksInner) HasPythonBinary() bool`

HasPythonBinary returns a boolean if a field has been set.

### GetPythonAdditionalPackages

`func (o *ListTasks200ResponseAllOfTasksInner) GetPythonAdditionalPackages() string`

GetPythonAdditionalPackages returns the PythonAdditionalPackages field if non-nil, zero value otherwise.

### GetPythonAdditionalPackagesOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetPythonAdditionalPackagesOk() (*string, bool)`

GetPythonAdditionalPackagesOk returns a tuple with the PythonAdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonAdditionalPackages

`func (o *ListTasks200ResponseAllOfTasksInner) SetPythonAdditionalPackages(v string)`

SetPythonAdditionalPackages sets PythonAdditionalPackages field to given value.

### HasPythonAdditionalPackages

`func (o *ListTasks200ResponseAllOfTasksInner) HasPythonAdditionalPackages() bool`

HasPythonAdditionalPackages returns a boolean if a field has been set.

### GetVroBody

`func (o *ListTasks200ResponseAllOfTasksInner) GetVroBody() string`

GetVroBody returns the VroBody field if non-nil, zero value otherwise.

### GetVroBodyOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetVroBodyOk() (*string, bool)`

GetVroBodyOk returns a tuple with the VroBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVroBody

`func (o *ListTasks200ResponseAllOfTasksInner) SetVroBody(v string)`

SetVroBody sets VroBody field to given value.

### HasVroBody

`func (o *ListTasks200ResponseAllOfTasksInner) HasVroBody() bool`

HasVroBody returns a boolean if a field has been set.

### GetWriteAttributesAttributes

`func (o *ListTasks200ResponseAllOfTasksInner) GetWriteAttributesAttributes() string`

GetWriteAttributesAttributes returns the WriteAttributesAttributes field if non-nil, zero value otherwise.

### GetWriteAttributesAttributesOk

`func (o *ListTasks200ResponseAllOfTasksInner) GetWriteAttributesAttributesOk() (*string, bool)`

GetWriteAttributesAttributesOk returns a tuple with the WriteAttributesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAttributesAttributes

`func (o *ListTasks200ResponseAllOfTasksInner) SetWriteAttributesAttributes(v string)`

SetWriteAttributesAttributes sets WriteAttributesAttributes field to given value.

### HasWriteAttributesAttributes

`func (o *ListTasks200ResponseAllOfTasksInner) HasWriteAttributesAttributes() bool`

HasWriteAttributesAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


