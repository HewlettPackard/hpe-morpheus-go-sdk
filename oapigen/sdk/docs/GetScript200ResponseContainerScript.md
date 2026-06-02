# GetScript200ResponseContainerScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**ScriptVersion** | Pointer to **string** |  | [optional] 
**ScriptPhase** | Pointer to **string** |  | [optional] 
**ScriptType** | Pointer to **string** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**ScriptService** | Pointer to **NullableString** |  | [optional] 
**ScriptMethod** | Pointer to **NullableString** |  | [optional] 
**RunAsUser** | Pointer to **NullableString** |  | [optional] 
**RunAsPassword** | Pointer to **NullableString** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**FailOnError** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetScript200ResponseContainerScript

`func NewGetScript200ResponseContainerScript() *GetScript200ResponseContainerScript`

NewGetScript200ResponseContainerScript instantiates a new GetScript200ResponseContainerScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetScript200ResponseContainerScript) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetScript200ResponseContainerScript) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetScript200ResponseContainerScript) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetScript200ResponseContainerScript) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetScript200ResponseContainerScript) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetScript200ResponseContainerScript) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetScript200ResponseContainerScript) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetScript200ResponseContainerScript) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *GetScript200ResponseContainerScript) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetScript200ResponseContainerScript) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetScript200ResponseContainerScript) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetScript200ResponseContainerScript) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetScript200ResponseContainerScript) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetScript200ResponseContainerScript) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetName

`func (o *GetScript200ResponseContainerScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetScript200ResponseContainerScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetScript200ResponseContainerScript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetScript200ResponseContainerScript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetScript200ResponseContainerScript) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetScript200ResponseContainerScript) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetScript200ResponseContainerScript) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetScript200ResponseContainerScript) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetScript200ResponseContainerScript) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetScript200ResponseContainerScript) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetCategory

`func (o *GetScript200ResponseContainerScript) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetScript200ResponseContainerScript) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetScript200ResponseContainerScript) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetScript200ResponseContainerScript) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetScript200ResponseContainerScript) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetScript200ResponseContainerScript) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetSortOrder

`func (o *GetScript200ResponseContainerScript) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetScript200ResponseContainerScript) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetScript200ResponseContainerScript) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetScript200ResponseContainerScript) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetScriptVersion

`func (o *GetScript200ResponseContainerScript) GetScriptVersion() string`

GetScriptVersion returns the ScriptVersion field if non-nil, zero value otherwise.

### GetScriptVersionOk

`func (o *GetScript200ResponseContainerScript) GetScriptVersionOk() (*string, bool)`

GetScriptVersionOk returns a tuple with the ScriptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptVersion

`func (o *GetScript200ResponseContainerScript) SetScriptVersion(v string)`

SetScriptVersion sets ScriptVersion field to given value.

### HasScriptVersion

`func (o *GetScript200ResponseContainerScript) HasScriptVersion() bool`

HasScriptVersion returns a boolean if a field has been set.

### GetScriptPhase

`func (o *GetScript200ResponseContainerScript) GetScriptPhase() string`

GetScriptPhase returns the ScriptPhase field if non-nil, zero value otherwise.

### GetScriptPhaseOk

`func (o *GetScript200ResponseContainerScript) GetScriptPhaseOk() (*string, bool)`

GetScriptPhaseOk returns a tuple with the ScriptPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPhase

`func (o *GetScript200ResponseContainerScript) SetScriptPhase(v string)`

SetScriptPhase sets ScriptPhase field to given value.

### HasScriptPhase

`func (o *GetScript200ResponseContainerScript) HasScriptPhase() bool`

HasScriptPhase returns a boolean if a field has been set.

### GetScriptType

`func (o *GetScript200ResponseContainerScript) GetScriptType() string`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *GetScript200ResponseContainerScript) GetScriptTypeOk() (*string, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *GetScript200ResponseContainerScript) SetScriptType(v string)`

SetScriptType sets ScriptType field to given value.

### HasScriptType

`func (o *GetScript200ResponseContainerScript) HasScriptType() bool`

HasScriptType returns a boolean if a field has been set.

### GetScript

`func (o *GetScript200ResponseContainerScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *GetScript200ResponseContainerScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *GetScript200ResponseContainerScript) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *GetScript200ResponseContainerScript) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetScriptService

`func (o *GetScript200ResponseContainerScript) GetScriptService() string`

GetScriptService returns the ScriptService field if non-nil, zero value otherwise.

### GetScriptServiceOk

`func (o *GetScript200ResponseContainerScript) GetScriptServiceOk() (*string, bool)`

GetScriptServiceOk returns a tuple with the ScriptService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptService

`func (o *GetScript200ResponseContainerScript) SetScriptService(v string)`

SetScriptService sets ScriptService field to given value.

### HasScriptService

`func (o *GetScript200ResponseContainerScript) HasScriptService() bool`

HasScriptService returns a boolean if a field has been set.

### SetScriptServiceNil

`func (o *GetScript200ResponseContainerScript) SetScriptServiceNil(b bool)`

 SetScriptServiceNil sets the value for ScriptService to be an explicit nil

### UnsetScriptService
`func (o *GetScript200ResponseContainerScript) UnsetScriptService()`

UnsetScriptService ensures that no value is present for ScriptService, not even an explicit nil
### GetScriptMethod

`func (o *GetScript200ResponseContainerScript) GetScriptMethod() string`

GetScriptMethod returns the ScriptMethod field if non-nil, zero value otherwise.

### GetScriptMethodOk

`func (o *GetScript200ResponseContainerScript) GetScriptMethodOk() (*string, bool)`

GetScriptMethodOk returns a tuple with the ScriptMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptMethod

`func (o *GetScript200ResponseContainerScript) SetScriptMethod(v string)`

SetScriptMethod sets ScriptMethod field to given value.

### HasScriptMethod

`func (o *GetScript200ResponseContainerScript) HasScriptMethod() bool`

HasScriptMethod returns a boolean if a field has been set.

### SetScriptMethodNil

`func (o *GetScript200ResponseContainerScript) SetScriptMethodNil(b bool)`

 SetScriptMethodNil sets the value for ScriptMethod to be an explicit nil

### UnsetScriptMethod
`func (o *GetScript200ResponseContainerScript) UnsetScriptMethod()`

UnsetScriptMethod ensures that no value is present for ScriptMethod, not even an explicit nil
### GetRunAsUser

`func (o *GetScript200ResponseContainerScript) GetRunAsUser() string`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *GetScript200ResponseContainerScript) GetRunAsUserOk() (*string, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *GetScript200ResponseContainerScript) SetRunAsUser(v string)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *GetScript200ResponseContainerScript) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### SetRunAsUserNil

`func (o *GetScript200ResponseContainerScript) SetRunAsUserNil(b bool)`

 SetRunAsUserNil sets the value for RunAsUser to be an explicit nil

### UnsetRunAsUser
`func (o *GetScript200ResponseContainerScript) UnsetRunAsUser()`

UnsetRunAsUser ensures that no value is present for RunAsUser, not even an explicit nil
### GetRunAsPassword

`func (o *GetScript200ResponseContainerScript) GetRunAsPassword() string`

GetRunAsPassword returns the RunAsPassword field if non-nil, zero value otherwise.

### GetRunAsPasswordOk

`func (o *GetScript200ResponseContainerScript) GetRunAsPasswordOk() (*string, bool)`

GetRunAsPasswordOk returns a tuple with the RunAsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsPassword

`func (o *GetScript200ResponseContainerScript) SetRunAsPassword(v string)`

SetRunAsPassword sets RunAsPassword field to given value.

### HasRunAsPassword

`func (o *GetScript200ResponseContainerScript) HasRunAsPassword() bool`

HasRunAsPassword returns a boolean if a field has been set.

### SetRunAsPasswordNil

`func (o *GetScript200ResponseContainerScript) SetRunAsPasswordNil(b bool)`

 SetRunAsPasswordNil sets the value for RunAsPassword to be an explicit nil

### UnsetRunAsPassword
`func (o *GetScript200ResponseContainerScript) UnsetRunAsPassword()`

UnsetRunAsPassword ensures that no value is present for RunAsPassword, not even an explicit nil
### GetSudoUser

`func (o *GetScript200ResponseContainerScript) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *GetScript200ResponseContainerScript) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *GetScript200ResponseContainerScript) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *GetScript200ResponseContainerScript) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetFailOnError

`func (o *GetScript200ResponseContainerScript) GetFailOnError() bool`

GetFailOnError returns the FailOnError field if non-nil, zero value otherwise.

### GetFailOnErrorOk

`func (o *GetScript200ResponseContainerScript) GetFailOnErrorOk() (*bool, bool)`

GetFailOnErrorOk returns a tuple with the FailOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnError

`func (o *GetScript200ResponseContainerScript) SetFailOnError(v bool)`

SetFailOnError sets FailOnError field to given value.

### HasFailOnError

`func (o *GetScript200ResponseContainerScript) HasFailOnError() bool`

HasFailOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


