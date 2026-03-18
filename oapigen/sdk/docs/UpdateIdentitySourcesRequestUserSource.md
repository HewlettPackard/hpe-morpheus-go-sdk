# UpdateIdentitySourcesRequestUserSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**UpdateIdentitySourcesRequestUserSourceAccount**](UpdateIdentitySourcesRequestUserSourceAccount.md) |  | [optional] 
**Name** | **string** | A name for the Identity Source | 
**Type** | **string** | IDM type code | 
**Description** | Pointer to **string** | description | [optional] 
**DefaultAccountRole** | Pointer to [**UpdateIdentitySourcesRequestUserSourceDefaultAccountRole**](UpdateIdentitySourcesRequestUserSourceDefaultAccountRole.md) |  | [optional] 
**RoleMappings** | Pointer to [**UpdateIdentitySourcesRequestUserSourceRoleMappings**](UpdateIdentitySourcesRequestUserSourceRoleMappings.md) |  | [optional] 
**RoleMappingNames** | Pointer to **map[string]string** | Map of Morpheus &#39;&#x60;Role ID&#x60;&#39;:&#39;&#x60;Role Name&#x60;&#39;.  | [optional] 
**AllowCustomMappings** | Pointer to **bool** | Enable Role Mapping Permission | [optional] 
**ManualRoleAssignment** | Pointer to **bool** | Manual Role Assignment | [optional] 
**Config** | Pointer to [**UpdateIdentitySourcesRequestUserSourceConfig**](UpdateIdentitySourcesRequestUserSourceConfig.md) |  | [optional] 

## Methods

### NewUpdateIdentitySourcesRequestUserSource

`func NewUpdateIdentitySourcesRequestUserSource(name string, type_ string, ) *UpdateIdentitySourcesRequestUserSource`

NewUpdateIdentitySourcesRequestUserSource instantiates a new UpdateIdentitySourcesRequestUserSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentitySourcesRequestUserSourceWithDefaults

`func NewUpdateIdentitySourcesRequestUserSourceWithDefaults() *UpdateIdentitySourcesRequestUserSource`

NewUpdateIdentitySourcesRequestUserSourceWithDefaults instantiates a new UpdateIdentitySourcesRequestUserSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UpdateIdentitySourcesRequestUserSource) GetAccount() UpdateIdentitySourcesRequestUserSourceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetAccountOk() (*UpdateIdentitySourcesRequestUserSourceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateIdentitySourcesRequestUserSource) SetAccount(v UpdateIdentitySourcesRequestUserSourceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateIdentitySourcesRequestUserSource) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *UpdateIdentitySourcesRequestUserSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIdentitySourcesRequestUserSource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateIdentitySourcesRequestUserSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIdentitySourcesRequestUserSource) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *UpdateIdentitySourcesRequestUserSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIdentitySourcesRequestUserSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIdentitySourcesRequestUserSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultAccountRole

`func (o *UpdateIdentitySourcesRequestUserSource) GetDefaultAccountRole() UpdateIdentitySourcesRequestUserSourceDefaultAccountRole`

GetDefaultAccountRole returns the DefaultAccountRole field if non-nil, zero value otherwise.

### GetDefaultAccountRoleOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetDefaultAccountRoleOk() (*UpdateIdentitySourcesRequestUserSourceDefaultAccountRole, bool)`

GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccountRole

`func (o *UpdateIdentitySourcesRequestUserSource) SetDefaultAccountRole(v UpdateIdentitySourcesRequestUserSourceDefaultAccountRole)`

SetDefaultAccountRole sets DefaultAccountRole field to given value.

### HasDefaultAccountRole

`func (o *UpdateIdentitySourcesRequestUserSource) HasDefaultAccountRole() bool`

HasDefaultAccountRole returns a boolean if a field has been set.

### GetRoleMappings

`func (o *UpdateIdentitySourcesRequestUserSource) GetRoleMappings() UpdateIdentitySourcesRequestUserSourceRoleMappings`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetRoleMappingsOk() (*UpdateIdentitySourcesRequestUserSourceRoleMappings, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *UpdateIdentitySourcesRequestUserSource) SetRoleMappings(v UpdateIdentitySourcesRequestUserSourceRoleMappings)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *UpdateIdentitySourcesRequestUserSource) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetRoleMappingNames

`func (o *UpdateIdentitySourcesRequestUserSource) GetRoleMappingNames() map[string]string`

GetRoleMappingNames returns the RoleMappingNames field if non-nil, zero value otherwise.

### GetRoleMappingNamesOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetRoleMappingNamesOk() (*map[string]string, bool)`

GetRoleMappingNamesOk returns a tuple with the RoleMappingNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappingNames

`func (o *UpdateIdentitySourcesRequestUserSource) SetRoleMappingNames(v map[string]string)`

SetRoleMappingNames sets RoleMappingNames field to given value.

### HasRoleMappingNames

`func (o *UpdateIdentitySourcesRequestUserSource) HasRoleMappingNames() bool`

HasRoleMappingNames returns a boolean if a field has been set.

### GetAllowCustomMappings

`func (o *UpdateIdentitySourcesRequestUserSource) GetAllowCustomMappings() bool`

GetAllowCustomMappings returns the AllowCustomMappings field if non-nil, zero value otherwise.

### GetAllowCustomMappingsOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetAllowCustomMappingsOk() (*bool, bool)`

GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMappings

`func (o *UpdateIdentitySourcesRequestUserSource) SetAllowCustomMappings(v bool)`

SetAllowCustomMappings sets AllowCustomMappings field to given value.

### HasAllowCustomMappings

`func (o *UpdateIdentitySourcesRequestUserSource) HasAllowCustomMappings() bool`

HasAllowCustomMappings returns a boolean if a field has been set.

### GetManualRoleAssignment

`func (o *UpdateIdentitySourcesRequestUserSource) GetManualRoleAssignment() bool`

GetManualRoleAssignment returns the ManualRoleAssignment field if non-nil, zero value otherwise.

### GetManualRoleAssignmentOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetManualRoleAssignmentOk() (*bool, bool)`

GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRoleAssignment

`func (o *UpdateIdentitySourcesRequestUserSource) SetManualRoleAssignment(v bool)`

SetManualRoleAssignment sets ManualRoleAssignment field to given value.

### HasManualRoleAssignment

`func (o *UpdateIdentitySourcesRequestUserSource) HasManualRoleAssignment() bool`

HasManualRoleAssignment returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIdentitySourcesRequestUserSource) GetConfig() UpdateIdentitySourcesRequestUserSourceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIdentitySourcesRequestUserSource) GetConfigOk() (*UpdateIdentitySourcesRequestUserSourceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIdentitySourcesRequestUserSource) SetConfig(v UpdateIdentitySourcesRequestUserSourceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIdentitySourcesRequestUserSource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


