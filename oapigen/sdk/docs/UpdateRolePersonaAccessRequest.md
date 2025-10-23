# UpdateRolePersonaAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonaCode** | **string** | code of the Persona, eg. &#x60;standard&#x60; or &#x60;serviceCatalog&#x60; | 
**Access** | **string** | The new access level. | 
**AllPersonas** | **bool** | Apply to all personas | 

## Methods

### NewUpdateRolePersonaAccessRequest

`func NewUpdateRolePersonaAccessRequest(personaCode string, access string, allPersonas bool, ) *UpdateRolePersonaAccessRequest`

NewUpdateRolePersonaAccessRequest instantiates a new UpdateRolePersonaAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRolePersonaAccessRequestWithDefaults

`func NewUpdateRolePersonaAccessRequestWithDefaults() *UpdateRolePersonaAccessRequest`

NewUpdateRolePersonaAccessRequestWithDefaults instantiates a new UpdateRolePersonaAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonaCode

`func (o *UpdateRolePersonaAccessRequest) GetPersonaCode() string`

GetPersonaCode returns the PersonaCode field if non-nil, zero value otherwise.

### GetPersonaCodeOk

`func (o *UpdateRolePersonaAccessRequest) GetPersonaCodeOk() (*string, bool)`

GetPersonaCodeOk returns a tuple with the PersonaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaCode

`func (o *UpdateRolePersonaAccessRequest) SetPersonaCode(v string)`

SetPersonaCode sets PersonaCode field to given value.


### GetAccess

`func (o *UpdateRolePersonaAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRolePersonaAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRolePersonaAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllPersonas

`func (o *UpdateRolePersonaAccessRequest) GetAllPersonas() bool`

GetAllPersonas returns the AllPersonas field if non-nil, zero value otherwise.

### GetAllPersonasOk

`func (o *UpdateRolePersonaAccessRequest) GetAllPersonasOk() (*bool, bool)`

GetAllPersonasOk returns a tuple with the AllPersonas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPersonas

`func (o *UpdateRolePersonaAccessRequest) SetAllPersonas(v bool)`

SetAllPersonas sets AllPersonas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


