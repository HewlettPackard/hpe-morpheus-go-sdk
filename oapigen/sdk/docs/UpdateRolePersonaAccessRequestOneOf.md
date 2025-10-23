# UpdateRolePersonaAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonaCode** | **string** | code of the Persona, eg. &#x60;standard&#x60; or &#x60;serviceCatalog&#x60; | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRolePersonaAccessRequestOneOf

`func NewUpdateRolePersonaAccessRequestOneOf(personaCode string, access string, ) *UpdateRolePersonaAccessRequestOneOf`

NewUpdateRolePersonaAccessRequestOneOf instantiates a new UpdateRolePersonaAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRolePersonaAccessRequestOneOfWithDefaults

`func NewUpdateRolePersonaAccessRequestOneOfWithDefaults() *UpdateRolePersonaAccessRequestOneOf`

NewUpdateRolePersonaAccessRequestOneOfWithDefaults instantiates a new UpdateRolePersonaAccessRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonaCode

`func (o *UpdateRolePersonaAccessRequestOneOf) GetPersonaCode() string`

GetPersonaCode returns the PersonaCode field if non-nil, zero value otherwise.

### GetPersonaCodeOk

`func (o *UpdateRolePersonaAccessRequestOneOf) GetPersonaCodeOk() (*string, bool)`

GetPersonaCodeOk returns a tuple with the PersonaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaCode

`func (o *UpdateRolePersonaAccessRequestOneOf) SetPersonaCode(v string)`

SetPersonaCode sets PersonaCode field to given value.


### GetAccess

`func (o *UpdateRolePersonaAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRolePersonaAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRolePersonaAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


