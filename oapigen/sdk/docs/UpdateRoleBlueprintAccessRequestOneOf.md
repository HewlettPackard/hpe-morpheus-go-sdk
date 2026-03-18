# UpdateRoleBlueprintAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppTemplateId** | **int32** | &#x60;id&#x60; of the blueprint (appTemplate) | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleBlueprintAccessRequestOneOf

`func NewUpdateRoleBlueprintAccessRequestOneOf(appTemplateId int32, access string, ) *UpdateRoleBlueprintAccessRequestOneOf`

NewUpdateRoleBlueprintAccessRequestOneOf instantiates a new UpdateRoleBlueprintAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleBlueprintAccessRequestOneOfWithDefaults

`func NewUpdateRoleBlueprintAccessRequestOneOfWithDefaults() *UpdateRoleBlueprintAccessRequestOneOf`

NewUpdateRoleBlueprintAccessRequestOneOfWithDefaults instantiates a new UpdateRoleBlueprintAccessRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppTemplateId

`func (o *UpdateRoleBlueprintAccessRequestOneOf) GetAppTemplateId() int32`

GetAppTemplateId returns the AppTemplateId field if non-nil, zero value otherwise.

### GetAppTemplateIdOk

`func (o *UpdateRoleBlueprintAccessRequestOneOf) GetAppTemplateIdOk() (*int32, bool)`

GetAppTemplateIdOk returns a tuple with the AppTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplateId

`func (o *UpdateRoleBlueprintAccessRequestOneOf) SetAppTemplateId(v int32)`

SetAppTemplateId sets AppTemplateId field to given value.


### GetAccess

`func (o *UpdateRoleBlueprintAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleBlueprintAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleBlueprintAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


