# UpdateScriptRequestContainerScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Script name | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Category** | Pointer to **string** | Script category | [optional] 
**ScriptVersion** | Pointer to **string** | Version of the script | [optional] [default to "1"]
**ScriptPhase** | Pointer to **string** | Phase for the script, provision, start, etc. | [optional] 
**ScriptType** | Pointer to **string** | Type for the script | [optional] [default to "bash"]
**Script** | Pointer to **string** | Script content, that is, the code itself. | [optional] 
**RunAsUser** | Pointer to **string** | Run as a specific user. | [optional] 
**SudoUser** | Pointer to **bool** | Sudo, whether or not to run with sudo. | [optional] [default to false]

## Methods

### NewUpdateScriptRequestContainerScript

`func NewUpdateScriptRequestContainerScript() *UpdateScriptRequestContainerScript`

NewUpdateScriptRequestContainerScript instantiates a new UpdateScriptRequestContainerScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScriptRequestContainerScriptWithDefaults

`func NewUpdateScriptRequestContainerScriptWithDefaults() *UpdateScriptRequestContainerScript`

NewUpdateScriptRequestContainerScriptWithDefaults instantiates a new UpdateScriptRequestContainerScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateScriptRequestContainerScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateScriptRequestContainerScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateScriptRequestContainerScript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateScriptRequestContainerScript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateScriptRequestContainerScript) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateScriptRequestContainerScript) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateScriptRequestContainerScript) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateScriptRequestContainerScript) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateScriptRequestContainerScript) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateScriptRequestContainerScript) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateScriptRequestContainerScript) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateScriptRequestContainerScript) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScriptVersion

`func (o *UpdateScriptRequestContainerScript) GetScriptVersion() string`

GetScriptVersion returns the ScriptVersion field if non-nil, zero value otherwise.

### GetScriptVersionOk

`func (o *UpdateScriptRequestContainerScript) GetScriptVersionOk() (*string, bool)`

GetScriptVersionOk returns a tuple with the ScriptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVersion

`func (o *UpdateScriptRequestContainerScript) SetScriptVersion(v string)`

SetScriptVersion sets ScriptVersion field to given value.

### HasScriptVersion

`func (o *UpdateScriptRequestContainerScript) HasScriptVersion() bool`

HasScriptVersion returns a boolean if a field has been set.

### GetScriptPhase

`func (o *UpdateScriptRequestContainerScript) GetScriptPhase() string`

GetScriptPhase returns the ScriptPhase field if non-nil, zero value otherwise.

### GetScriptPhaseOk

`func (o *UpdateScriptRequestContainerScript) GetScriptPhaseOk() (*string, bool)`

GetScriptPhaseOk returns a tuple with the ScriptPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPhase

`func (o *UpdateScriptRequestContainerScript) SetScriptPhase(v string)`

SetScriptPhase sets ScriptPhase field to given value.

### HasScriptPhase

`func (o *UpdateScriptRequestContainerScript) HasScriptPhase() bool`

HasScriptPhase returns a boolean if a field has been set.

### GetScriptType

`func (o *UpdateScriptRequestContainerScript) GetScriptType() string`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *UpdateScriptRequestContainerScript) GetScriptTypeOk() (*string, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *UpdateScriptRequestContainerScript) SetScriptType(v string)`

SetScriptType sets ScriptType field to given value.

### HasScriptType

`func (o *UpdateScriptRequestContainerScript) HasScriptType() bool`

HasScriptType returns a boolean if a field has been set.

### GetScript

`func (o *UpdateScriptRequestContainerScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *UpdateScriptRequestContainerScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *UpdateScriptRequestContainerScript) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *UpdateScriptRequestContainerScript) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetRunAsUser

`func (o *UpdateScriptRequestContainerScript) GetRunAsUser() string`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *UpdateScriptRequestContainerScript) GetRunAsUserOk() (*string, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *UpdateScriptRequestContainerScript) SetRunAsUser(v string)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *UpdateScriptRequestContainerScript) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSudoUser

`func (o *UpdateScriptRequestContainerScript) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *UpdateScriptRequestContainerScript) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *UpdateScriptRequestContainerScript) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *UpdateScriptRequestContainerScript) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


