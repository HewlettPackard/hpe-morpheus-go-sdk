# Script

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**ScriptVersion** | Pointer to **string** |  | [optional] 
**ScriptPhase** | Pointer to **string** |  | [optional] 
**ScriptType** | Pointer to **string** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**ScriptService** | Pointer to **string** |  | [optional] 
**ScriptMethod** | Pointer to **string** |  | [optional] 
**RunAsUser** | Pointer to **string** |  | [optional] 
**RunAsPassword** | Pointer to **string** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**FailOnError** | Pointer to **bool** |  | [optional] 

## Methods

### NewScript

`func NewScript() *Script`

NewScript instantiates a new Script object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWithDefaults

`func NewScriptWithDefaults() *Script`

NewScriptWithDefaults instantiates a new Script object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Script) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Script) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Script) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Script) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *Script) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Script) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Script) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Script) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *Script) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Script) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Script) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Script) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *Script) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Script) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Script) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Script) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *Script) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Script) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Script) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Script) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCategory

`func (o *Script) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Script) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Script) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Script) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSortOrder

`func (o *Script) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Script) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Script) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Script) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetScriptVersion

`func (o *Script) GetScriptVersion() string`

GetScriptVersion returns the ScriptVersion field if non-nil, zero value otherwise.

### GetScriptVersionOk

`func (o *Script) GetScriptVersionOk() (*string, bool)`

GetScriptVersionOk returns a tuple with the ScriptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVersion

`func (o *Script) SetScriptVersion(v string)`

SetScriptVersion sets ScriptVersion field to given value.

### HasScriptVersion

`func (o *Script) HasScriptVersion() bool`

HasScriptVersion returns a boolean if a field has been set.

### GetScriptPhase

`func (o *Script) GetScriptPhase() string`

GetScriptPhase returns the ScriptPhase field if non-nil, zero value otherwise.

### GetScriptPhaseOk

`func (o *Script) GetScriptPhaseOk() (*string, bool)`

GetScriptPhaseOk returns a tuple with the ScriptPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPhase

`func (o *Script) SetScriptPhase(v string)`

SetScriptPhase sets ScriptPhase field to given value.

### HasScriptPhase

`func (o *Script) HasScriptPhase() bool`

HasScriptPhase returns a boolean if a field has been set.

### GetScriptType

`func (o *Script) GetScriptType() string`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *Script) GetScriptTypeOk() (*string, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *Script) SetScriptType(v string)`

SetScriptType sets ScriptType field to given value.

### HasScriptType

`func (o *Script) HasScriptType() bool`

HasScriptType returns a boolean if a field has been set.

### GetScript

`func (o *Script) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Script) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Script) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *Script) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetScriptService

`func (o *Script) GetScriptService() string`

GetScriptService returns the ScriptService field if non-nil, zero value otherwise.

### GetScriptServiceOk

`func (o *Script) GetScriptServiceOk() (*string, bool)`

GetScriptServiceOk returns a tuple with the ScriptService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptService

`func (o *Script) SetScriptService(v string)`

SetScriptService sets ScriptService field to given value.

### HasScriptService

`func (o *Script) HasScriptService() bool`

HasScriptService returns a boolean if a field has been set.

### GetScriptMethod

`func (o *Script) GetScriptMethod() string`

GetScriptMethod returns the ScriptMethod field if non-nil, zero value otherwise.

### GetScriptMethodOk

`func (o *Script) GetScriptMethodOk() (*string, bool)`

GetScriptMethodOk returns a tuple with the ScriptMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptMethod

`func (o *Script) SetScriptMethod(v string)`

SetScriptMethod sets ScriptMethod field to given value.

### HasScriptMethod

`func (o *Script) HasScriptMethod() bool`

HasScriptMethod returns a boolean if a field has been set.

### GetRunAsUser

`func (o *Script) GetRunAsUser() string`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *Script) GetRunAsUserOk() (*string, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *Script) SetRunAsUser(v string)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *Script) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetRunAsPassword

`func (o *Script) GetRunAsPassword() string`

GetRunAsPassword returns the RunAsPassword field if non-nil, zero value otherwise.

### GetRunAsPasswordOk

`func (o *Script) GetRunAsPasswordOk() (*string, bool)`

GetRunAsPasswordOk returns a tuple with the RunAsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsPassword

`func (o *Script) SetRunAsPassword(v string)`

SetRunAsPassword sets RunAsPassword field to given value.

### HasRunAsPassword

`func (o *Script) HasRunAsPassword() bool`

HasRunAsPassword returns a boolean if a field has been set.

### GetSudoUser

`func (o *Script) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *Script) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *Script) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *Script) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetFailOnError

`func (o *Script) GetFailOnError() bool`

GetFailOnError returns the FailOnError field if non-nil, zero value otherwise.

### GetFailOnErrorOk

`func (o *Script) GetFailOnErrorOk() (*bool, bool)`

GetFailOnErrorOk returns a tuple with the FailOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnError

`func (o *Script) SetFailOnError(v bool)`

SetFailOnError sets FailOnError field to given value.

### HasFailOnError

`func (o *Script) HasFailOnError() bool`

HasFailOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


