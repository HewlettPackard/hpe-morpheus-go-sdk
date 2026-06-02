# CreateNetworkServerGroupRequestGroupMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MemberName** | Pointer to **NullableString** |  | [optional] 
**MemberType** | Pointer to **NullableString** |  | [optional] 
**MemberValue** | Pointer to **NullableString** |  | [optional] 
**MemberExpression** | Pointer to **NullableString** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewCreateNetworkServerGroupRequestGroupMembersInner

`func NewCreateNetworkServerGroupRequestGroupMembersInner() *CreateNetworkServerGroupRequestGroupMembersInner`

NewCreateNetworkServerGroupRequestGroupMembersInner instantiates a new CreateNetworkServerGroupRequestGroupMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMemberName

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### SetMemberNameNil

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberNameNil(b bool)`

 SetMemberNameNil sets the value for MemberName to be an explicit nil

### UnsetMemberName
`func (o *CreateNetworkServerGroupRequestGroupMembersInner) UnsetMemberName()`

UnsetMemberName ensures that no value is present for MemberName, not even an explicit nil
### GetMemberType

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### SetMemberTypeNil

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberTypeNil(b bool)`

 SetMemberTypeNil sets the value for MemberType to be an explicit nil

### UnsetMemberType
`func (o *CreateNetworkServerGroupRequestGroupMembersInner) UnsetMemberType()`

UnsetMemberType ensures that no value is present for MemberType, not even an explicit nil
### GetMemberValue

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberValue() string`

GetMemberValue returns the MemberValue field if non-nil, zero value otherwise.

### GetMemberValueOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberValueOk() (*string, bool)`

GetMemberValueOk returns a tuple with the MemberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberValue

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberValue(v string)`

SetMemberValue sets MemberValue field to given value.

### HasMemberValue

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasMemberValue() bool`

HasMemberValue returns a boolean if a field has been set.

### SetMemberValueNil

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberValueNil(b bool)`

 SetMemberValueNil sets the value for MemberValue to be an explicit nil

### UnsetMemberValue
`func (o *CreateNetworkServerGroupRequestGroupMembersInner) UnsetMemberValue()`

UnsetMemberValue ensures that no value is present for MemberValue, not even an explicit nil
### GetMemberExpression

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberExpression() string`

GetMemberExpression returns the MemberExpression field if non-nil, zero value otherwise.

### GetMemberExpressionOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMemberExpressionOk() (*string, bool)`

GetMemberExpressionOk returns a tuple with the MemberExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberExpression

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberExpression(v string)`

SetMemberExpression sets MemberExpression field to given value.

### HasMemberExpression

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasMemberExpression() bool`

HasMemberExpression returns a boolean if a field has been set.

### SetMemberExpressionNil

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMemberExpressionNil(b bool)`

 SetMemberExpressionNil sets the value for MemberExpression to be an explicit nil

### UnsetMemberExpression
`func (o *CreateNetworkServerGroupRequestGroupMembersInner) UnsetMemberExpression()`

UnsetMemberExpression ensures that no value is present for MemberExpression, not even an explicit nil
### GetDisplayOrder

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *CreateNetworkServerGroupRequestGroupMembersInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMembers

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CreateNetworkServerGroupRequestGroupMembersInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


