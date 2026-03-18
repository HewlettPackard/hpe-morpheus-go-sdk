# GetTaskTypes200ResponseTaskType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Scriptable** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HasResults** | Pointer to **bool** |  | [optional] 
**AllowExecuteLocal** | Pointer to **NullableBool** |  | [optional] 
**AllowExecuteRemote** | Pointer to **NullableBool** |  | [optional] 
**AllowExecuteResource** | Pointer to **NullableBool** |  | [optional] 
**AllowLocalRepo** | Pointer to **NullableBool** |  | [optional] 
**AllowRemoteKeyAuth** | Pointer to **NullableBool** |  | [optional] 
**OptionTypes** | Pointer to [**[]GetTaskTypes200ResponseTaskTypeOptionTypesInner**](GetTaskTypes200ResponseTaskTypeOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetTaskTypes200ResponseTaskType

`func NewGetTaskTypes200ResponseTaskType() *GetTaskTypes200ResponseTaskType`

NewGetTaskTypes200ResponseTaskType instantiates a new GetTaskTypes200ResponseTaskType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaskTypes200ResponseTaskTypeWithDefaults

`func NewGetTaskTypes200ResponseTaskTypeWithDefaults() *GetTaskTypes200ResponseTaskType`

NewGetTaskTypes200ResponseTaskTypeWithDefaults instantiates a new GetTaskTypes200ResponseTaskType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTaskTypes200ResponseTaskType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTaskTypes200ResponseTaskType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTaskTypes200ResponseTaskType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTaskTypes200ResponseTaskType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetTaskTypes200ResponseTaskType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTaskTypes200ResponseTaskType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTaskTypes200ResponseTaskType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetTaskTypes200ResponseTaskType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetTaskTypes200ResponseTaskType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTaskTypes200ResponseTaskType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTaskTypes200ResponseTaskType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTaskTypes200ResponseTaskType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *GetTaskTypes200ResponseTaskType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetTaskTypes200ResponseTaskType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetTaskTypes200ResponseTaskType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetTaskTypes200ResponseTaskType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *GetTaskTypes200ResponseTaskType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTaskTypes200ResponseTaskType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTaskTypes200ResponseTaskType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTaskTypes200ResponseTaskType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetTaskTypes200ResponseTaskType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetTaskTypes200ResponseTaskType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScriptable

`func (o *GetTaskTypes200ResponseTaskType) GetScriptable() bool`

GetScriptable returns the Scriptable field if non-nil, zero value otherwise.

### GetScriptableOk

`func (o *GetTaskTypes200ResponseTaskType) GetScriptableOk() (*bool, bool)`

GetScriptableOk returns a tuple with the Scriptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptable

`func (o *GetTaskTypes200ResponseTaskType) SetScriptable(v bool)`

SetScriptable sets Scriptable field to given value.

### HasScriptable

`func (o *GetTaskTypes200ResponseTaskType) HasScriptable() bool`

HasScriptable returns a boolean if a field has been set.

### GetEnabled

`func (o *GetTaskTypes200ResponseTaskType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetTaskTypes200ResponseTaskType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetTaskTypes200ResponseTaskType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetTaskTypes200ResponseTaskType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasResults

`func (o *GetTaskTypes200ResponseTaskType) GetHasResults() bool`

GetHasResults returns the HasResults field if non-nil, zero value otherwise.

### GetHasResultsOk

`func (o *GetTaskTypes200ResponseTaskType) GetHasResultsOk() (*bool, bool)`

GetHasResultsOk returns a tuple with the HasResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResults

`func (o *GetTaskTypes200ResponseTaskType) SetHasResults(v bool)`

SetHasResults sets HasResults field to given value.

### HasHasResults

`func (o *GetTaskTypes200ResponseTaskType) HasHasResults() bool`

HasHasResults returns a boolean if a field has been set.

### GetAllowExecuteLocal

`func (o *GetTaskTypes200ResponseTaskType) GetAllowExecuteLocal() bool`

GetAllowExecuteLocal returns the AllowExecuteLocal field if non-nil, zero value otherwise.

### GetAllowExecuteLocalOk

`func (o *GetTaskTypes200ResponseTaskType) GetAllowExecuteLocalOk() (*bool, bool)`

GetAllowExecuteLocalOk returns a tuple with the AllowExecuteLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteLocal

`func (o *GetTaskTypes200ResponseTaskType) SetAllowExecuteLocal(v bool)`

SetAllowExecuteLocal sets AllowExecuteLocal field to given value.

### HasAllowExecuteLocal

`func (o *GetTaskTypes200ResponseTaskType) HasAllowExecuteLocal() bool`

HasAllowExecuteLocal returns a boolean if a field has been set.

### SetAllowExecuteLocalNil

`func (o *GetTaskTypes200ResponseTaskType) SetAllowExecuteLocalNil(b bool)`

 SetAllowExecuteLocalNil sets the value for AllowExecuteLocal to be an explicit nil

### UnsetAllowExecuteLocal
`func (o *GetTaskTypes200ResponseTaskType) UnsetAllowExecuteLocal()`

UnsetAllowExecuteLocal ensures that no value is present for AllowExecuteLocal, not even an explicit nil
### GetAllowExecuteRemote

`func (o *GetTaskTypes200ResponseTaskType) GetAllowExecuteRemote() bool`

GetAllowExecuteRemote returns the AllowExecuteRemote field if non-nil, zero value otherwise.

### GetAllowExecuteRemoteOk

`func (o *GetTaskTypes200ResponseTaskType) GetAllowExecuteRemoteOk() (*bool, bool)`

GetAllowExecuteRemoteOk returns a tuple with the AllowExecuteRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteRemote

`func (o *GetTaskTypes200ResponseTaskType) SetAllowExecuteRemote(v bool)`

SetAllowExecuteRemote sets AllowExecuteRemote field to given value.

### HasAllowExecuteRemote

`func (o *GetTaskTypes200ResponseTaskType) HasAllowExecuteRemote() bool`

HasAllowExecuteRemote returns a boolean if a field has been set.

### SetAllowExecuteRemoteNil

`func (o *GetTaskTypes200ResponseTaskType) SetAllowExecuteRemoteNil(b bool)`

 SetAllowExecuteRemoteNil sets the value for AllowExecuteRemote to be an explicit nil

### UnsetAllowExecuteRemote
`func (o *GetTaskTypes200ResponseTaskType) UnsetAllowExecuteRemote()`

UnsetAllowExecuteRemote ensures that no value is present for AllowExecuteRemote, not even an explicit nil
### GetAllowExecuteResource

`func (o *GetTaskTypes200ResponseTaskType) GetAllowExecuteResource() bool`

GetAllowExecuteResource returns the AllowExecuteResource field if non-nil, zero value otherwise.

### GetAllowExecuteResourceOk

`func (o *GetTaskTypes200ResponseTaskType) GetAllowExecuteResourceOk() (*bool, bool)`

GetAllowExecuteResourceOk returns a tuple with the AllowExecuteResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteResource

`func (o *GetTaskTypes200ResponseTaskType) SetAllowExecuteResource(v bool)`

SetAllowExecuteResource sets AllowExecuteResource field to given value.

### HasAllowExecuteResource

`func (o *GetTaskTypes200ResponseTaskType) HasAllowExecuteResource() bool`

HasAllowExecuteResource returns a boolean if a field has been set.

### SetAllowExecuteResourceNil

`func (o *GetTaskTypes200ResponseTaskType) SetAllowExecuteResourceNil(b bool)`

 SetAllowExecuteResourceNil sets the value for AllowExecuteResource to be an explicit nil

### UnsetAllowExecuteResource
`func (o *GetTaskTypes200ResponseTaskType) UnsetAllowExecuteResource()`

UnsetAllowExecuteResource ensures that no value is present for AllowExecuteResource, not even an explicit nil
### GetAllowLocalRepo

`func (o *GetTaskTypes200ResponseTaskType) GetAllowLocalRepo() bool`

GetAllowLocalRepo returns the AllowLocalRepo field if non-nil, zero value otherwise.

### GetAllowLocalRepoOk

`func (o *GetTaskTypes200ResponseTaskType) GetAllowLocalRepoOk() (*bool, bool)`

GetAllowLocalRepoOk returns a tuple with the AllowLocalRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalRepo

`func (o *GetTaskTypes200ResponseTaskType) SetAllowLocalRepo(v bool)`

SetAllowLocalRepo sets AllowLocalRepo field to given value.

### HasAllowLocalRepo

`func (o *GetTaskTypes200ResponseTaskType) HasAllowLocalRepo() bool`

HasAllowLocalRepo returns a boolean if a field has been set.

### SetAllowLocalRepoNil

`func (o *GetTaskTypes200ResponseTaskType) SetAllowLocalRepoNil(b bool)`

 SetAllowLocalRepoNil sets the value for AllowLocalRepo to be an explicit nil

### UnsetAllowLocalRepo
`func (o *GetTaskTypes200ResponseTaskType) UnsetAllowLocalRepo()`

UnsetAllowLocalRepo ensures that no value is present for AllowLocalRepo, not even an explicit nil
### GetAllowRemoteKeyAuth

`func (o *GetTaskTypes200ResponseTaskType) GetAllowRemoteKeyAuth() bool`

GetAllowRemoteKeyAuth returns the AllowRemoteKeyAuth field if non-nil, zero value otherwise.

### GetAllowRemoteKeyAuthOk

`func (o *GetTaskTypes200ResponseTaskType) GetAllowRemoteKeyAuthOk() (*bool, bool)`

GetAllowRemoteKeyAuthOk returns a tuple with the AllowRemoteKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteKeyAuth

`func (o *GetTaskTypes200ResponseTaskType) SetAllowRemoteKeyAuth(v bool)`

SetAllowRemoteKeyAuth sets AllowRemoteKeyAuth field to given value.

### HasAllowRemoteKeyAuth

`func (o *GetTaskTypes200ResponseTaskType) HasAllowRemoteKeyAuth() bool`

HasAllowRemoteKeyAuth returns a boolean if a field has been set.

### SetAllowRemoteKeyAuthNil

`func (o *GetTaskTypes200ResponseTaskType) SetAllowRemoteKeyAuthNil(b bool)`

 SetAllowRemoteKeyAuthNil sets the value for AllowRemoteKeyAuth to be an explicit nil

### UnsetAllowRemoteKeyAuth
`func (o *GetTaskTypes200ResponseTaskType) UnsetAllowRemoteKeyAuth()`

UnsetAllowRemoteKeyAuth ensures that no value is present for AllowRemoteKeyAuth, not even an explicit nil
### GetOptionTypes

`func (o *GetTaskTypes200ResponseTaskType) GetOptionTypes() []GetTaskTypes200ResponseTaskTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetTaskTypes200ResponseTaskType) GetOptionTypesOk() (*[]GetTaskTypes200ResponseTaskTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetTaskTypes200ResponseTaskType) SetOptionTypes(v []GetTaskTypes200ResponseTaskTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetTaskTypes200ResponseTaskType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


