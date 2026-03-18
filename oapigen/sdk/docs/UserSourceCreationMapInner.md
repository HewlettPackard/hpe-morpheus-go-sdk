# UserSourceCreationMapInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappedRole** | Pointer to [**UserSourceCreationMapInnerMappedRole**](UserSourceCreationMapInnerMappedRole.md) |  | [optional] 
**SourceRoleFqn** | Pointer to **string** | Fully Qualified Name of the role in the identity source | [optional] 
**SourceRoleName** | Pointer to **string** | Name of the role in the identity source | [optional] 

## Methods

### NewUserSourceCreationMapInner

`func NewUserSourceCreationMapInner() *UserSourceCreationMapInner`

NewUserSourceCreationMapInner instantiates a new UserSourceCreationMapInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreationMapInnerWithDefaults

`func NewUserSourceCreationMapInnerWithDefaults() *UserSourceCreationMapInner`

NewUserSourceCreationMapInnerWithDefaults instantiates a new UserSourceCreationMapInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappedRole

`func (o *UserSourceCreationMapInner) GetMappedRole() UserSourceCreationMapInnerMappedRole`

GetMappedRole returns the MappedRole field if non-nil, zero value otherwise.

### GetMappedRoleOk

`func (o *UserSourceCreationMapInner) GetMappedRoleOk() (*UserSourceCreationMapInnerMappedRole, bool)`

GetMappedRoleOk returns a tuple with the MappedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedRole

`func (o *UserSourceCreationMapInner) SetMappedRole(v UserSourceCreationMapInnerMappedRole)`

SetMappedRole sets MappedRole field to given value.

### HasMappedRole

`func (o *UserSourceCreationMapInner) HasMappedRole() bool`

HasMappedRole returns a boolean if a field has been set.

### GetSourceRoleFqn

`func (o *UserSourceCreationMapInner) GetSourceRoleFqn() string`

GetSourceRoleFqn returns the SourceRoleFqn field if non-nil, zero value otherwise.

### GetSourceRoleFqnOk

`func (o *UserSourceCreationMapInner) GetSourceRoleFqnOk() (*string, bool)`

GetSourceRoleFqnOk returns a tuple with the SourceRoleFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoleFqn

`func (o *UserSourceCreationMapInner) SetSourceRoleFqn(v string)`

SetSourceRoleFqn sets SourceRoleFqn field to given value.

### HasSourceRoleFqn

`func (o *UserSourceCreationMapInner) HasSourceRoleFqn() bool`

HasSourceRoleFqn returns a boolean if a field has been set.

### GetSourceRoleName

`func (o *UserSourceCreationMapInner) GetSourceRoleName() string`

GetSourceRoleName returns the SourceRoleName field if non-nil, zero value otherwise.

### GetSourceRoleNameOk

`func (o *UserSourceCreationMapInner) GetSourceRoleNameOk() (*string, bool)`

GetSourceRoleNameOk returns a tuple with the SourceRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoleName

`func (o *UserSourceCreationMapInner) SetSourceRoleName(v string)`

SetSourceRoleName sets SourceRoleName field to given value.

### HasSourceRoleName

`func (o *UserSourceCreationMapInner) HasSourceRoleName() bool`

HasSourceRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


