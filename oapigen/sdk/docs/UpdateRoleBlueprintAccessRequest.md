# UpdateRoleBlueprintAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppTemplateId** | **int32** | &#x60;id&#x60; of the blueprint (appTemplate) | 
**Access** | **string** | The new access level. | 
**AllAppTemplates** | **bool** | Apply to all blueprints (appTemplate) | 

## Methods

### NewUpdateRoleBlueprintAccessRequest

`func NewUpdateRoleBlueprintAccessRequest(appTemplateId int32, access string, allAppTemplates bool, ) *UpdateRoleBlueprintAccessRequest`

NewUpdateRoleBlueprintAccessRequest instantiates a new UpdateRoleBlueprintAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleBlueprintAccessRequestWithDefaults

`func NewUpdateRoleBlueprintAccessRequestWithDefaults() *UpdateRoleBlueprintAccessRequest`

NewUpdateRoleBlueprintAccessRequestWithDefaults instantiates a new UpdateRoleBlueprintAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppTemplateId

`func (o *UpdateRoleBlueprintAccessRequest) GetAppTemplateId() int32`

GetAppTemplateId returns the AppTemplateId field if non-nil, zero value otherwise.

### GetAppTemplateIdOk

`func (o *UpdateRoleBlueprintAccessRequest) GetAppTemplateIdOk() (*int32, bool)`

GetAppTemplateIdOk returns a tuple with the AppTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplateId

`func (o *UpdateRoleBlueprintAccessRequest) SetAppTemplateId(v int32)`

SetAppTemplateId sets AppTemplateId field to given value.


### GetAccess

`func (o *UpdateRoleBlueprintAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleBlueprintAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleBlueprintAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllAppTemplates

`func (o *UpdateRoleBlueprintAccessRequest) GetAllAppTemplates() bool`

GetAllAppTemplates returns the AllAppTemplates field if non-nil, zero value otherwise.

### GetAllAppTemplatesOk

`func (o *UpdateRoleBlueprintAccessRequest) GetAllAppTemplatesOk() (*bool, bool)`

GetAllAppTemplatesOk returns a tuple with the AllAppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAppTemplates

`func (o *UpdateRoleBlueprintAccessRequest) SetAllAppTemplates(v bool)`

SetAllAppTemplates sets AllAppTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


