# ListCatalogCart200ResponseCartItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListApps200ResponseAllOfAppsInnerBlueprint**](ListApps200ResponseAllOfAppsInnerBlueprint.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**ListCatalogCart200ResponseCartItemsInnerInstance**](ListCatalogCart200ResponseCartItemsInnerInstance.md) |  | [optional] 
**OrderDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListCatalogCart200ResponseCartItemsInner

`func NewListCatalogCart200ResponseCartItemsInner() *ListCatalogCart200ResponseCartItemsInner`

NewListCatalogCart200ResponseCartItemsInner instantiates a new ListCatalogCart200ResponseCartItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogCart200ResponseCartItemsInnerWithDefaults

`func NewListCatalogCart200ResponseCartItemsInnerWithDefaults() *ListCatalogCart200ResponseCartItemsInner`

NewListCatalogCart200ResponseCartItemsInnerWithDefaults instantiates a new ListCatalogCart200ResponseCartItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCatalogCart200ResponseCartItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCatalogCart200ResponseCartItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCatalogCart200ResponseCartItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCatalogCart200ResponseCartItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCatalogCart200ResponseCartItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCatalogCart200ResponseCartItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListCatalogCart200ResponseCartItemsInner) GetType() ListApps200ResponseAllOfAppsInnerBlueprint`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetTypeOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCatalogCart200ResponseCartItemsInner) SetType(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetType sets Type field to given value.

### HasType

`func (o *ListCatalogCart200ResponseCartItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *ListCatalogCart200ResponseCartItemsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListCatalogCart200ResponseCartItemsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ListCatalogCart200ResponseCartItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *ListCatalogCart200ResponseCartItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCatalogCart200ResponseCartItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListCatalogCart200ResponseCartItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListCatalogCart200ResponseCartItemsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListCatalogCart200ResponseCartItemsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListCatalogCart200ResponseCartItemsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetRefType

`func (o *ListCatalogCart200ResponseCartItemsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListCatalogCart200ResponseCartItemsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListCatalogCart200ResponseCartItemsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetInstance

`func (o *ListCatalogCart200ResponseCartItemsInner) GetInstance() ListCatalogCart200ResponseCartItemsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetInstanceOk() (*ListCatalogCart200ResponseCartItemsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListCatalogCart200ResponseCartItemsInner) SetInstance(v ListCatalogCart200ResponseCartItemsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListCatalogCart200ResponseCartItemsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetOrderDate

`func (o *ListCatalogCart200ResponseCartItemsInner) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *ListCatalogCart200ResponseCartItemsInner) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *ListCatalogCart200ResponseCartItemsInner) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCatalogCart200ResponseCartItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCatalogCart200ResponseCartItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCatalogCart200ResponseCartItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCatalogCart200ResponseCartItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCatalogCart200ResponseCartItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCatalogCart200ResponseCartItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCatalogCart200ResponseCartItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


