# GetCatalogItem200ResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetCatalogItem200ResponseItemType**](GetCatalogItem200ResponseItemType.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**GetCatalogItem200ResponseItemInstance**](GetCatalogItem200ResponseItemInstance.md) |  | [optional] 
**OrderDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetCatalogItem200ResponseItem

`func NewGetCatalogItem200ResponseItem() *GetCatalogItem200ResponseItem`

NewGetCatalogItem200ResponseItem instantiates a new GetCatalogItem200ResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCatalogItem200ResponseItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogItem200ResponseItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogItem200ResponseItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCatalogItem200ResponseItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCatalogItem200ResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCatalogItem200ResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCatalogItem200ResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCatalogItem200ResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetCatalogItem200ResponseItem) GetType() GetCatalogItem200ResponseItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogItem200ResponseItem) GetTypeOk() (*GetCatalogItem200ResponseItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogItem200ResponseItem) SetType(v GetCatalogItem200ResponseItemType)`

SetType sets Type field to given value.

### HasType

`func (o *GetCatalogItem200ResponseItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *GetCatalogItem200ResponseItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetCatalogItem200ResponseItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetCatalogItem200ResponseItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GetCatalogItem200ResponseItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *GetCatalogItem200ResponseItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCatalogItem200ResponseItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCatalogItem200ResponseItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCatalogItem200ResponseItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetCatalogItem200ResponseItem) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetCatalogItem200ResponseItem) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetCatalogItem200ResponseItem) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetCatalogItem200ResponseItem) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetCatalogItem200ResponseItem) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetCatalogItem200ResponseItem) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetRefType

`func (o *GetCatalogItem200ResponseItem) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetCatalogItem200ResponseItem) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetCatalogItem200ResponseItem) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetCatalogItem200ResponseItem) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetInstance

`func (o *GetCatalogItem200ResponseItem) GetInstance() GetCatalogItem200ResponseItemInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetCatalogItem200ResponseItem) GetInstanceOk() (*GetCatalogItem200ResponseItemInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetCatalogItem200ResponseItem) SetInstance(v GetCatalogItem200ResponseItemInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetCatalogItem200ResponseItem) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetOrderDate

`func (o *GetCatalogItem200ResponseItem) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *GetCatalogItem200ResponseItem) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *GetCatalogItem200ResponseItem) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *GetCatalogItem200ResponseItem) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCatalogItem200ResponseItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCatalogItem200ResponseItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCatalogItem200ResponseItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCatalogItem200ResponseItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCatalogItem200ResponseItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCatalogItem200ResponseItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCatalogItem200ResponseItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCatalogItem200ResponseItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


