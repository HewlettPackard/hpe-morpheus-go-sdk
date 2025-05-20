# AddIdentitySourcesRequestUserSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AddIdentitySourcesRequestUserSourceAccount**](AddIdentitySourcesRequestUserSourceAccount.md) |  | [optional] 
**Name** | **string** | A name for the Identity Source | 
**Type** | **string** | IDM type code | 
**Description** | Pointer to **string** | description | [optional] 
**DefaultAccountRole** | Pointer to [**AddIdentitySourcesRequestUserSourceDefaultAccountRole**](AddIdentitySourcesRequestUserSourceDefaultAccountRole.md) |  | [optional] 
**RoleMappings** | Pointer to [**AddIdentitySourcesRequestUserSourceRoleMappings**](AddIdentitySourcesRequestUserSourceRoleMappings.md) |  | [optional] 
**RoleMappingNames** | Pointer to **map[string]string** | Map of Morpheus &#39;&#x60;Role ID&#x60;&#39;:&#39;&#x60;Role Name&#x60;&#39;.  | [optional] 
**AllowCustomMappings** | Pointer to **bool** | Enable Role Mapping Permission | [optional] 
**ManualRoleAssignment** | Pointer to **bool** | Manual Role Assignment | [optional] 
**Config** | Pointer to [**AddIdentitySourcesRequestUserSourceConfig**](AddIdentitySourcesRequestUserSourceConfig.md) |  | [optional] 

## Methods

### NewAddIdentitySourcesRequestUserSource

`func NewAddIdentitySourcesRequestUserSource(name string, type_ string, ) *AddIdentitySourcesRequestUserSource`

NewAddIdentitySourcesRequestUserSource instantiates a new AddIdentitySourcesRequestUserSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdentitySourcesRequestUserSourceWithDefaults

`func NewAddIdentitySourcesRequestUserSourceWithDefaults() *AddIdentitySourcesRequestUserSource`

NewAddIdentitySourcesRequestUserSourceWithDefaults instantiates a new AddIdentitySourcesRequestUserSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AddIdentitySourcesRequestUserSource) GetAccount() AddIdentitySourcesRequestUserSourceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddIdentitySourcesRequestUserSource) GetAccountOk() (*AddIdentitySourcesRequestUserSourceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddIdentitySourcesRequestUserSource) SetAccount(v AddIdentitySourcesRequestUserSourceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddIdentitySourcesRequestUserSource) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *AddIdentitySourcesRequestUserSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIdentitySourcesRequestUserSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIdentitySourcesRequestUserSource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddIdentitySourcesRequestUserSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIdentitySourcesRequestUserSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIdentitySourcesRequestUserSource) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *AddIdentitySourcesRequestUserSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIdentitySourcesRequestUserSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIdentitySourcesRequestUserSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIdentitySourcesRequestUserSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *AddIdentitySourcesRequestUserSource) GetDefaultAccountRole() AddIdentitySourcesRequestUserSourceDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *AddIdentitySourcesRequestUserSource) GetDefaultAccountRoleOk() (*AddIdentitySourcesRequestUserSourceDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *AddIdentitySourcesRequestUserSource) SetDefaultAccountRole(v AddIdentitySourcesRequestUserSourceDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *AddIdentitySourcesRequestUserSource) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetRoleMappings

`func (o *AddIdentitySourcesRequestUserSource) GetRoleMappings() AddIdentitySourcesRequestUserSourceRoleMappings`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *AddIdentitySourcesRequestUserSource) GetRoleMappingsOk() (*AddIdentitySourcesRequestUserSourceRoleMappings, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *AddIdentitySourcesRequestUserSource) SetRoleMappings(v AddIdentitySourcesRequestUserSourceRoleMappings)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *AddIdentitySourcesRequestUserSource) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetRoleMappingNames

`func (o *AddIdentitySourcesRequestUserSource) GetRoleMappingNames() map[string]string`

GetRoleMappingNames returns the RoleMappingNames field if non-nil, zero value otherwise.

### GetRoleMappingNamesOk

`func (o *AddIdentitySourcesRequestUserSource) GetRoleMappingNamesOk() (*map[string]string, bool)`

GetRoleMappingNamesOk returns a tuple with the RoleMappingNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappingNames

`func (o *AddIdentitySourcesRequestUserSource) SetRoleMappingNames(v map[string]string)`

SetRoleMappingNames sets RoleMappingNames field to given value.

### HasRoleMappingNames

`func (o *AddIdentitySourcesRequestUserSource) HasRoleMappingNames() bool`

HasRoleMappingNames returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *AddIdentitySourcesRequestUserSource) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *AddIdentitySourcesRequestUserSource) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *AddIdentitySourcesRequestUserSource) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *AddIdentitySourcesRequestUserSource) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *AddIdentitySourcesRequestUserSource) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *AddIdentitySourcesRequestUserSource) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *AddIdentitySourcesRequestUserSource) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *AddIdentitySourcesRequestUserSource) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetConfig

`func (o *AddIdentitySourcesRequestUserSource) GetConfig() AddIdentitySourcesRequestUserSourceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIdentitySourcesRequestUserSource) GetConfigOk() (*AddIdentitySourcesRequestUserSourceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIdentitySourcesRequestUserSource) SetConfig(v AddIdentitySourcesRequestUserSourceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIdentitySourcesRequestUserSource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


