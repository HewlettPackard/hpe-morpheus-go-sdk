# RolePermissionPersona

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonaCode** | **string** | code of the Persona, eg. &#x60;standard&#x60; or &#x60;serviceCatalog&#x60; | 
**Access** | **string** | The new access level. | 

## Methods

### NewRolePermissionPersona

`func NewRolePermissionPersona(personaCode string, access string, ) *RolePermissionPersona`

NewRolePermissionPersona instantiates a new RolePermissionPersona object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionPersonaWithDefaults

`func NewRolePermissionPersonaWithDefaults() *RolePermissionPersona`

NewRolePermissionPersonaWithDefaults instantiates a new RolePermissionPersona object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonaCode

`func (o *RolePermissionPersona) GetPersonaCode() string`

GetPersonaCode returns the PersonaCode field if non-nil, zero value otherwise.

### GetPersonaCodeOk

`func (o *RolePermissionPersona) GetPersonaCodeOk() (*string, bool)`

GetPersonaCodeOk returns a tuple with the PersonaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaCode

`func (o *RolePermissionPersona) SetPersonaCode(v string)`

SetPersonaCode sets PersonaCode field to given value.


### GetAccess

`func (o *RolePermissionPersona) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RolePermissionPersona) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RolePermissionPersona) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


