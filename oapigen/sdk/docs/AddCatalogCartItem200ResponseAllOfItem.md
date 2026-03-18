# AddCatalogCartItem200ResponseAllOfItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AddCatalogCartItem200ResponseAllOfItemType**](AddCatalogCartItem200ResponseAllOfItemType.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**AddCatalogCartItem200ResponseAllOfItemInstance**](AddCatalogCartItem200ResponseAllOfItemInstance.md) |  | [optional] 
**OrderDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddCatalogCartItem200ResponseAllOfItem

`func NewAddCatalogCartItem200ResponseAllOfItem() *AddCatalogCartItem200ResponseAllOfItem`

NewAddCatalogCartItem200ResponseAllOfItem instantiates a new AddCatalogCartItem200ResponseAllOfItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogCartItem200ResponseAllOfItemWithDefaults

`func NewAddCatalogCartItem200ResponseAllOfItemWithDefaults() *AddCatalogCartItem200ResponseAllOfItem`

NewAddCatalogCartItem200ResponseAllOfItemWithDefaults instantiates a new AddCatalogCartItem200ResponseAllOfItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetType() AddCatalogCartItem200ResponseAllOfItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetTypeOk() (*AddCatalogCartItem200ResponseAllOfItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetType(v AddCatalogCartItem200ResponseAllOfItemType)`

SetType sets Type field to given value.

### HasType

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddCatalogCartItem200ResponseAllOfItem) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetRefType

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetInstance

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetInstance() AddCatalogCartItem200ResponseAllOfItemInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetInstanceOk() (*AddCatalogCartItem200ResponseAllOfItemInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetInstance(v AddCatalogCartItem200ResponseAllOfItemInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetOrderDate

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddCatalogCartItem200ResponseAllOfItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddCatalogCartItem200ResponseAllOfItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddCatalogCartItem200ResponseAllOfItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


