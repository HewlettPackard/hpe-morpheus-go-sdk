# AddIntegrationSnowObjectsRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name to display | [optional] 
**Type** | **string** | Integration Object Type Code | 
**CatalogItemType** | **int64** | Catalog Item Type ID | 

## Methods

### NewAddIntegrationSnowObjectsRequestObject

`func NewAddIntegrationSnowObjectsRequestObject(type_ string, catalogItemType int64, ) *AddIntegrationSnowObjectsRequestObject`

NewAddIntegrationSnowObjectsRequestObject instantiates a new AddIntegrationSnowObjectsRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrationSnowObjectsRequestObjectWithDefaults

`func NewAddIntegrationSnowObjectsRequestObjectWithDefaults() *AddIntegrationSnowObjectsRequestObject`

NewAddIntegrationSnowObjectsRequestObjectWithDefaults instantiates a new AddIntegrationSnowObjectsRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddIntegrationSnowObjectsRequestObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrationSnowObjectsRequestObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrationSnowObjectsRequestObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIntegrationSnowObjectsRequestObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddIntegrationSnowObjectsRequestObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrationSnowObjectsRequestObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrationSnowObjectsRequestObject) SetType(v string)`

SetType sets Type field to given value.


### GetCatalogItemType

`func (o *AddIntegrationSnowObjectsRequestObject) GetCatalogItemType() int64`

GetCatalogItemType returns the CatalogItemType field if non-nil, zero value otherwise.

### GetCatalogItemTypeOk

`func (o *AddIntegrationSnowObjectsRequestObject) GetCatalogItemTypeOk() (*int64, bool)`

GetCatalogItemTypeOk returns a tuple with the CatalogItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemType

`func (o *AddIntegrationSnowObjectsRequestObject) SetCatalogItemType(v int64)`

SetCatalogItemType sets CatalogItemType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


