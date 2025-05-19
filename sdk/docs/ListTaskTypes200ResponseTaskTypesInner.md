# ListTaskTypes200ResponseTaskTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Scriptable** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HasResults** | Pointer to **bool** |  | [optional] 
**AllowExecuteLocal** | Pointer to **bool** |  | [optional] 
**AllowExecuteRemote** | Pointer to **bool** |  | [optional] 
**AllowExecuteResource** | Pointer to **bool** |  | [optional] 
**AllowLocalRepo** | Pointer to **bool** |  | [optional] 
**AllowRemoteKeyAuth** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner**](ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewListTaskTypes200ResponseTaskTypesInner

`func NewListTaskTypes200ResponseTaskTypesInner() *ListTaskTypes200ResponseTaskTypesInner`

NewListTaskTypes200ResponseTaskTypesInner instantiates a new ListTaskTypes200ResponseTaskTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskTypes200ResponseTaskTypesInnerWithDefaults

`func NewListTaskTypes200ResponseTaskTypesInnerWithDefaults() *ListTaskTypes200ResponseTaskTypesInner`

NewListTaskTypes200ResponseTaskTypesInnerWithDefaults instantiates a new ListTaskTypes200ResponseTaskTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScriptable

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetScriptable() bool`

GetScriptable returns the Scriptable field if non-nil, zero value otherwise.

### GetScriptableOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetScriptableOk() (*bool, bool)`

GetScriptableOk returns a tuple with the Scriptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptable

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetScriptable(v bool)`

SetScriptable sets Scriptable field to given value.

### HasScriptable

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasScriptable() bool`

HasScriptable returns a boolean if a field has been set.

### GetEnabled

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasResults

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetHasResults() bool`

GetHasResults returns the HasResults field if non-nil, zero value otherwise.

### GetHasResultsOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetHasResultsOk() (*bool, bool)`

GetHasResultsOk returns a tuple with the HasResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResults

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetHasResults(v bool)`

SetHasResults sets HasResults field to given value.

### HasHasResults

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasHasResults() bool`

HasHasResults returns a boolean if a field has been set.

### GetAllowExecuteLocal

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteLocal() bool`

GetAllowExecuteLocal returns the AllowExecuteLocal field if non-nil, zero value otherwise.

### GetAllowExecuteLocalOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteLocalOk() (*bool, bool)`

GetAllowExecuteLocalOk returns a tuple with the AllowExecuteLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteLocal

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowExecuteLocal(v bool)`

SetAllowExecuteLocal sets AllowExecuteLocal field to given value.

### HasAllowExecuteLocal

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasAllowExecuteLocal() bool`

HasAllowExecuteLocal returns a boolean if a field has been set.

### GetAllowExecuteRemote

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteRemote() bool`

GetAllowExecuteRemote returns the AllowExecuteRemote field if non-nil, zero value otherwise.

### GetAllowExecuteRemoteOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteRemoteOk() (*bool, bool)`

GetAllowExecuteRemoteOk returns a tuple with the AllowExecuteRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteRemote

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowExecuteRemote(v bool)`

SetAllowExecuteRemote sets AllowExecuteRemote field to given value.

### HasAllowExecuteRemote

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasAllowExecuteRemote() bool`

HasAllowExecuteRemote returns a boolean if a field has been set.

### GetAllowExecuteResource

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteResource() bool`

GetAllowExecuteResource returns the AllowExecuteResource field if non-nil, zero value otherwise.

### GetAllowExecuteResourceOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteResourceOk() (*bool, bool)`

GetAllowExecuteResourceOk returns a tuple with the AllowExecuteResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteResource

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowExecuteResource(v bool)`

SetAllowExecuteResource sets AllowExecuteResource field to given value.

### HasAllowExecuteResource

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasAllowExecuteResource() bool`

HasAllowExecuteResource returns a boolean if a field has been set.

### GetAllowLocalRepo

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowLocalRepo() bool`

GetAllowLocalRepo returns the AllowLocalRepo field if non-nil, zero value otherwise.

### GetAllowLocalRepoOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowLocalRepoOk() (*bool, bool)`

GetAllowLocalRepoOk returns a tuple with the AllowLocalRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalRepo

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowLocalRepo(v bool)`

SetAllowLocalRepo sets AllowLocalRepo field to given value.

### HasAllowLocalRepo

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasAllowLocalRepo() bool`

HasAllowLocalRepo returns a boolean if a field has been set.

### GetAllowRemoteKeyAuth

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowRemoteKeyAuth() bool`

GetAllowRemoteKeyAuth returns the AllowRemoteKeyAuth field if non-nil, zero value otherwise.

### GetAllowRemoteKeyAuthOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowRemoteKeyAuthOk() (*bool, bool)`

GetAllowRemoteKeyAuthOk returns a tuple with the AllowRemoteKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteKeyAuth

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowRemoteKeyAuth(v bool)`

SetAllowRemoteKeyAuth sets AllowRemoteKeyAuth field to given value.

### HasAllowRemoteKeyAuth

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasAllowRemoteKeyAuth() bool`

HasAllowRemoteKeyAuth returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetOptionTypes() []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListTaskTypes200ResponseTaskTypesInner) GetOptionTypesOk() (*[]ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListTaskTypes200ResponseTaskTypesInner) SetOptionTypes(v []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListTaskTypes200ResponseTaskTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


