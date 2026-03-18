# ListCatalogItems200ResponseAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListCatalogItems200ResponseAllOfItemsInnerType**](ListCatalogItems200ResponseAllOfItemsInnerType.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**ListCatalogItems200ResponseAllOfItemsInnerInstance**](ListCatalogItems200ResponseAllOfItemsInnerInstance.md) |  | [optional] 
**OrderDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListCatalogItems200ResponseAllOfItemsInner

`func NewListCatalogItems200ResponseAllOfItemsInner() *ListCatalogItems200ResponseAllOfItemsInner`

NewListCatalogItems200ResponseAllOfItemsInner instantiates a new ListCatalogItems200ResponseAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogItems200ResponseAllOfItemsInnerWithDefaults

`func NewListCatalogItems200ResponseAllOfItemsInnerWithDefaults() *ListCatalogItems200ResponseAllOfItemsInner`

NewListCatalogItems200ResponseAllOfItemsInnerWithDefaults instantiates a new ListCatalogItems200ResponseAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetType() ListCatalogItems200ResponseAllOfItemsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetTypeOk() (*ListCatalogItems200ResponseAllOfItemsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetType(v ListCatalogItems200ResponseAllOfItemsInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListCatalogItems200ResponseAllOfItemsInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetRefType

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetInstance

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetInstance() ListCatalogItems200ResponseAllOfItemsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetInstanceOk() (*ListCatalogItems200ResponseAllOfItemsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetInstance(v ListCatalogItems200ResponseAllOfItemsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetOrderDate

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCatalogItems200ResponseAllOfItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCatalogItems200ResponseAllOfItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCatalogItems200ResponseAllOfItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


