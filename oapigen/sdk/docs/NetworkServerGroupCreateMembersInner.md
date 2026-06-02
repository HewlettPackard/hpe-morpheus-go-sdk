# NetworkServerGroupCreateMembersInner

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

### NewNetworkServerGroupCreateMembersInner

`func NewNetworkServerGroupCreateMembersInner() *NetworkServerGroupCreateMembersInner`

NewNetworkServerGroupCreateMembersInner instantiates a new NetworkServerGroupCreateMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *NetworkServerGroupCreateMembersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkServerGroupCreateMembersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkServerGroupCreateMembersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkServerGroupCreateMembersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *NetworkServerGroupCreateMembersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NetworkServerGroupCreateMembersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NetworkServerGroupCreateMembersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NetworkServerGroupCreateMembersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *NetworkServerGroupCreateMembersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkServerGroupCreateMembersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkServerGroupCreateMembersInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkServerGroupCreateMembersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMemberName

`func (o *NetworkServerGroupCreateMembersInner) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *NetworkServerGroupCreateMembersInner) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *NetworkServerGroupCreateMembersInner) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *NetworkServerGroupCreateMembersInner) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### SetMemberNameNil

`func (o *NetworkServerGroupCreateMembersInner) SetMemberNameNil(b bool)`

 SetMemberNameNil sets the value for MemberName to be an explicit nil

### UnsetMemberName
`func (o *NetworkServerGroupCreateMembersInner) UnsetMemberName()`

UnsetMemberName ensures that no value is present for MemberName, not even an explicit nil
### GetMemberType

`func (o *NetworkServerGroupCreateMembersInner) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *NetworkServerGroupCreateMembersInner) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *NetworkServerGroupCreateMembersInner) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *NetworkServerGroupCreateMembersInner) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### SetMemberTypeNil

`func (o *NetworkServerGroupCreateMembersInner) SetMemberTypeNil(b bool)`

 SetMemberTypeNil sets the value for MemberType to be an explicit nil

### UnsetMemberType
`func (o *NetworkServerGroupCreateMembersInner) UnsetMemberType()`

UnsetMemberType ensures that no value is present for MemberType, not even an explicit nil
### GetMemberValue

`func (o *NetworkServerGroupCreateMembersInner) GetMemberValue() string`

GetMemberValue returns the MemberValue field if non-nil, zero value otherwise.

### GetMemberValueOk

`func (o *NetworkServerGroupCreateMembersInner) GetMemberValueOk() (*string, bool)`

GetMemberValueOk returns a tuple with the MemberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberValue

`func (o *NetworkServerGroupCreateMembersInner) SetMemberValue(v string)`

SetMemberValue sets MemberValue field to given value.

### HasMemberValue

`func (o *NetworkServerGroupCreateMembersInner) HasMemberValue() bool`

HasMemberValue returns a boolean if a field has been set.

### SetMemberValueNil

`func (o *NetworkServerGroupCreateMembersInner) SetMemberValueNil(b bool)`

 SetMemberValueNil sets the value for MemberValue to be an explicit nil

### UnsetMemberValue
`func (o *NetworkServerGroupCreateMembersInner) UnsetMemberValue()`

UnsetMemberValue ensures that no value is present for MemberValue, not even an explicit nil
### GetMemberExpression

`func (o *NetworkServerGroupCreateMembersInner) GetMemberExpression() string`

GetMemberExpression returns the MemberExpression field if non-nil, zero value otherwise.

### GetMemberExpressionOk

`func (o *NetworkServerGroupCreateMembersInner) GetMemberExpressionOk() (*string, bool)`

GetMemberExpressionOk returns a tuple with the MemberExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberExpression

`func (o *NetworkServerGroupCreateMembersInner) SetMemberExpression(v string)`

SetMemberExpression sets MemberExpression field to given value.

### HasMemberExpression

`func (o *NetworkServerGroupCreateMembersInner) HasMemberExpression() bool`

HasMemberExpression returns a boolean if a field has been set.

### SetMemberExpressionNil

`func (o *NetworkServerGroupCreateMembersInner) SetMemberExpressionNil(b bool)`

 SetMemberExpressionNil sets the value for MemberExpression to be an explicit nil

### UnsetMemberExpression
`func (o *NetworkServerGroupCreateMembersInner) UnsetMemberExpression()`

UnsetMemberExpression ensures that no value is present for MemberExpression, not even an explicit nil
### GetDisplayOrder

`func (o *NetworkServerGroupCreateMembersInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *NetworkServerGroupCreateMembersInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *NetworkServerGroupCreateMembersInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *NetworkServerGroupCreateMembersInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkServerGroupCreateMembersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkServerGroupCreateMembersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkServerGroupCreateMembersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkServerGroupCreateMembersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *NetworkServerGroupCreateMembersInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *NetworkServerGroupCreateMembersInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *NetworkServerGroupCreateMembersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkServerGroupCreateMembersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkServerGroupCreateMembersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkServerGroupCreateMembersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMembers

`func (o *NetworkServerGroupCreateMembersInner) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *NetworkServerGroupCreateMembersInner) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *NetworkServerGroupCreateMembersInner) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *NetworkServerGroupCreateMembersInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


