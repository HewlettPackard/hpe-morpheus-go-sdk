# UpdateRoleCatalogItemTypeAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogItemTypeId** | **int32** | &#x60;id&#x60; of the catalog item type | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleCatalogItemTypeAccessRequestOneOf

`func NewUpdateRoleCatalogItemTypeAccessRequestOneOf(catalogItemTypeId int32, access string, ) *UpdateRoleCatalogItemTypeAccessRequestOneOf`

NewUpdateRoleCatalogItemTypeAccessRequestOneOf instantiates a new UpdateRoleCatalogItemTypeAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleCatalogItemTypeAccessRequestOneOfWithDefaults

`func NewUpdateRoleCatalogItemTypeAccessRequestOneOfWithDefaults() *UpdateRoleCatalogItemTypeAccessRequestOneOf`

NewUpdateRoleCatalogItemTypeAccessRequestOneOfWithDefaults instantiates a new UpdateRoleCatalogItemTypeAccessRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogItemTypeId

`func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf) GetCatalogItemTypeId() int32`

GetCatalogItemTypeId returns the CatalogItemTypeId field if non-nil, zero value otherwise.

### GetCatalogItemTypeIdOk

`func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf) GetCatalogItemTypeIdOk() (*int32, bool)`

GetCatalogItemTypeIdOk returns a tuple with the CatalogItemTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypeId

`func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf) SetCatalogItemTypeId(v int32)`

SetCatalogItemTypeId sets CatalogItemTypeId field to given value.


### GetAccess

`func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


