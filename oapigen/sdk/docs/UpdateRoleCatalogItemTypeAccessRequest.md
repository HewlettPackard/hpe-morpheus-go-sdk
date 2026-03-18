# UpdateRoleCatalogItemTypeAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogItemTypeId** | **int32** | &#x60;id&#x60; of the catalog item type | 
**Access** | **string** | The new access level. | 
**AllCatalogItemTypes** | **bool** | Apply to all catalog item types | 

## Methods

### NewUpdateRoleCatalogItemTypeAccessRequest

`func NewUpdateRoleCatalogItemTypeAccessRequest(catalogItemTypeId int32, access string, allCatalogItemTypes bool, ) *UpdateRoleCatalogItemTypeAccessRequest`

NewUpdateRoleCatalogItemTypeAccessRequest instantiates a new UpdateRoleCatalogItemTypeAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleCatalogItemTypeAccessRequestWithDefaults

`func NewUpdateRoleCatalogItemTypeAccessRequestWithDefaults() *UpdateRoleCatalogItemTypeAccessRequest`

NewUpdateRoleCatalogItemTypeAccessRequestWithDefaults instantiates a new UpdateRoleCatalogItemTypeAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogItemTypeId

`func (o *UpdateRoleCatalogItemTypeAccessRequest) GetCatalogItemTypeId() int32`

GetCatalogItemTypeId returns the CatalogItemTypeId field if non-nil, zero value otherwise.

### GetCatalogItemTypeIdOk

`func (o *UpdateRoleCatalogItemTypeAccessRequest) GetCatalogItemTypeIdOk() (*int32, bool)`

GetCatalogItemTypeIdOk returns a tuple with the CatalogItemTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypeId

`func (o *UpdateRoleCatalogItemTypeAccessRequest) SetCatalogItemTypeId(v int32)`

SetCatalogItemTypeId sets CatalogItemTypeId field to given value.


### GetAccess

`func (o *UpdateRoleCatalogItemTypeAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleCatalogItemTypeAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleCatalogItemTypeAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllCatalogItemTypes

`func (o *UpdateRoleCatalogItemTypeAccessRequest) GetAllCatalogItemTypes() bool`

GetAllCatalogItemTypes returns the AllCatalogItemTypes field if non-nil, zero value otherwise.

### GetAllCatalogItemTypesOk

`func (o *UpdateRoleCatalogItemTypeAccessRequest) GetAllCatalogItemTypesOk() (*bool, bool)`

GetAllCatalogItemTypesOk returns a tuple with the AllCatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCatalogItemTypes

`func (o *UpdateRoleCatalogItemTypeAccessRequest) SetAllCatalogItemTypes(v bool)`

SetAllCatalogItemTypes sets AllCatalogItemTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


