# AddScriptRequestContainerScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Script name | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Category** | Pointer to **string** | Script category | [optional] 
**ScriptVersion** | Pointer to **string** | Version of the script | [optional] [default to "1"]
**ScriptPhase** | Pointer to **string** | Phase for the script, provision, start, etc. | [optional] 
**ScriptType** | Pointer to **string** | Type for the script | [optional] [default to "bash"]
**Script** | Pointer to **string** | Script content, that is, the code itself. | [optional] 
**RunAsUser** | Pointer to **string** | Run as a specific user. | [optional] 
**SudoUser** | Pointer to **bool** | Sudo, whether or not to run with sudo. | [optional] [default to false]

## Methods

### NewAddScriptRequestContainerScript

`func NewAddScriptRequestContainerScript(name string, ) *AddScriptRequestContainerScript`

NewAddScriptRequestContainerScript instantiates a new AddScriptRequestContainerScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScriptRequestContainerScriptWithDefaults

`func NewAddScriptRequestContainerScriptWithDefaults() *AddScriptRequestContainerScript`

NewAddScriptRequestContainerScriptWithDefaults instantiates a new AddScriptRequestContainerScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddScriptRequestContainerScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddScriptRequestContainerScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddScriptRequestContainerScript) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *AddScriptRequestContainerScript) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddScriptRequestContainerScript) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddScriptRequestContainerScript) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddScriptRequestContainerScript) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCategory

`func (o *AddScriptRequestContainerScript) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddScriptRequestContainerScript) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddScriptRequestContainerScript) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddScriptRequestContainerScript) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScriptVersion

`func (o *AddScriptRequestContainerScript) GetScriptVersion() string`

GetScriptVersion returns the ScriptVersion field if non-nil, zero value otherwise.

### GetScriptVersionOk

`func (o *AddScriptRequestContainerScript) GetScriptVersionOk() (*string, bool)`

GetScriptVersionOk returns a tuple with the ScriptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVersion

`func (o *AddScriptRequestContainerScript) SetScriptVersion(v string)`

SetScriptVersion sets ScriptVersion field to given value.

### HasScriptVersion

`func (o *AddScriptRequestContainerScript) HasScriptVersion() bool`

HasScriptVersion returns a boolean if a field has been set.

### GetScriptPhase

`func (o *AddScriptRequestContainerScript) GetScriptPhase() string`

GetScriptPhase returns the ScriptPhase field if non-nil, zero value otherwise.

### GetScriptPhaseOk

`func (o *AddScriptRequestContainerScript) GetScriptPhaseOk() (*string, bool)`

GetScriptPhaseOk returns a tuple with the ScriptPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPhase

`func (o *AddScriptRequestContainerScript) SetScriptPhase(v string)`

SetScriptPhase sets ScriptPhase field to given value.

### HasScriptPhase

`func (o *AddScriptRequestContainerScript) HasScriptPhase() bool`

HasScriptPhase returns a boolean if a field has been set.

### GetScriptType

`func (o *AddScriptRequestContainerScript) GetScriptType() string`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *AddScriptRequestContainerScript) GetScriptTypeOk() (*string, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *AddScriptRequestContainerScript) SetScriptType(v string)`

SetScriptType sets ScriptType field to given value.

### HasScriptType

`func (o *AddScriptRequestContainerScript) HasScriptType() bool`

HasScriptType returns a boolean if a field has been set.

### GetScript

`func (o *AddScriptRequestContainerScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *AddScriptRequestContainerScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *AddScriptRequestContainerScript) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *AddScriptRequestContainerScript) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetRunAsUser

`func (o *AddScriptRequestContainerScript) GetRunAsUser() string`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *AddScriptRequestContainerScript) GetRunAsUserOk() (*string, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *AddScriptRequestContainerScript) SetRunAsUser(v string)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *AddScriptRequestContainerScript) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSudoUser

`func (o *AddScriptRequestContainerScript) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *AddScriptRequestContainerScript) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *AddScriptRequestContainerScript) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *AddScriptRequestContainerScript) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


