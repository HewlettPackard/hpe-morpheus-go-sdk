# UserSourceCreateUserSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**UserSourceCreateUserSourceAccount**](UserSourceCreateUserSourceAccount.md) |  | [optional] 
**Name** | **string** | A name for the Identity Source | 
**Type** | **string** | IDM type code | 
**Description** | Pointer to **string** | description | [optional] 
**DefaultAccountRole** | Pointer to [**UserSourceCreateUserSourceDefaultAccountRole**](UserSourceCreateUserSourceDefaultAccountRole.md) |  | [optional] 
**RoleMappings** | Pointer to [**UserSourceCreateUserSourceRoleMappings**](UserSourceCreateUserSourceRoleMappings.md) |  | [optional] 
**RoleMappingNames** | Pointer to **map[string]string** | Map of Morpheus &#39;&#x60;Role ID&#x60;&#39;:&#39;&#x60;Role Name&#x60;&#39;.  | [optional] 
**AllowCustomMappings** | Pointer to **bool** | Enable Role Mapping Permission | [optional] 
**ManualRoleAssignment** | Pointer to **bool** | Manual Role Assignment | [optional] 
**Config** | Pointer to [**UserSourceCreateUserSourceConfig**](UserSourceCreateUserSourceConfig.md) |  | [optional] 

## Methods

### NewUserSourceCreateUserSource

`func NewUserSourceCreateUserSource(name string, type_ string, ) *UserSourceCreateUserSource`

NewUserSourceCreateUserSource instantiates a new UserSourceCreateUserSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateUserSourceWithDefaults

`func NewUserSourceCreateUserSourceWithDefaults() *UserSourceCreateUserSource`

NewUserSourceCreateUserSourceWithDefaults instantiates a new UserSourceCreateUserSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UserSourceCreateUserSource) GetAccount() UserSourceCreateUserSourceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserSourceCreateUserSource) GetAccountOk() (*UserSourceCreateUserSourceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserSourceCreateUserSource) SetAccount(v UserSourceCreateUserSourceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserSourceCreateUserSource) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *UserSourceCreateUserSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSourceCreateUserSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSourceCreateUserSource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UserSourceCreateUserSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSourceCreateUserSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSourceCreateUserSource) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *UserSourceCreateUserSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSourceCreateUserSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSourceCreateUserSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserSourceCreateUserSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *UserSourceCreateUserSource) GetDefaultAccountRole() UserSourceCreateUserSourceDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *UserSourceCreateUserSource) GetDefaultAccountRoleOk() (*UserSourceCreateUserSourceDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *UserSourceCreateUserSource) SetDefaultAccountRole(v UserSourceCreateUserSourceDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *UserSourceCreateUserSource) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetRoleMappings

`func (o *UserSourceCreateUserSource) GetRoleMappings() UserSourceCreateUserSourceRoleMappings`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *UserSourceCreateUserSource) GetRoleMappingsOk() (*UserSourceCreateUserSourceRoleMappings, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *UserSourceCreateUserSource) SetRoleMappings(v UserSourceCreateUserSourceRoleMappings)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *UserSourceCreateUserSource) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetRoleMappingNames

`func (o *UserSourceCreateUserSource) GetRoleMappingNames() map[string]string`

GetRoleMappingNames returns the RoleMappingNames field if non-nil, zero value otherwise.

### GetRoleMappingNamesOk

`func (o *UserSourceCreateUserSource) GetRoleMappingNamesOk() (*map[string]string, bool)`

GetRoleMappingNamesOk returns a tuple with the RoleMappingNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappingNames

`func (o *UserSourceCreateUserSource) SetRoleMappingNames(v map[string]string)`

SetRoleMappingNames sets RoleMappingNames field to given value.

### HasRoleMappingNames

`func (o *UserSourceCreateUserSource) HasRoleMappingNames() bool`

HasRoleMappingNames returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *UserSourceCreateUserSource) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *UserSourceCreateUserSource) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *UserSourceCreateUserSource) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *UserSourceCreateUserSource) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *UserSourceCreateUserSource) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *UserSourceCreateUserSource) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *UserSourceCreateUserSource) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *UserSourceCreateUserSource) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetConfig

`func (o *UserSourceCreateUserSource) GetConfig() UserSourceCreateUserSourceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UserSourceCreateUserSource) GetConfigOk() (*UserSourceCreateUserSourceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UserSourceCreateUserSource) SetConfig(v UserSourceCreateUserSourceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UserSourceCreateUserSource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


