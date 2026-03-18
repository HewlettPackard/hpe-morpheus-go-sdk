# GetInvoices200ResponseAllOfInvoiceLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **int64** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ItemSeller** | Pointer to **NullableString** |  | [optional] 
**ItemAction** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RateId** | Pointer to **string** |  | [optional] 
**RateClass** | Pointer to **NullableString** |  | [optional] 
**RateUnit** | Pointer to **string** |  | [optional] 
**RateTerm** | Pointer to **NullableString** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 
**UsageCategory** | Pointer to **string** |  | [optional] 
**UsageService** | Pointer to **string** |  | [optional] 
**ItemUsage** | Pointer to **float32** |  | [optional] 
**ItemRate** | Pointer to **float32** |  | [optional] 
**ItemCost** | Pointer to **float32** |  | [optional] 
**ItemPrice** | Pointer to **float32** |  | [optional] 
**ItemTax** | Pointer to **int64** |  | [optional] 
**ItemTerm** | Pointer to **NullableString** |  | [optional] 
**TaxType** | Pointer to **NullableString** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetInvoices200ResponseAllOfInvoiceLineItemsInner

`func NewGetInvoices200ResponseAllOfInvoiceLineItemsInner() *GetInvoices200ResponseAllOfInvoiceLineItemsInner`

NewGetInvoices200ResponseAllOfInvoiceLineItemsInner instantiates a new GetInvoices200ResponseAllOfInvoiceLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoices200ResponseAllOfInvoiceLineItemsInnerWithDefaults

`func NewGetInvoices200ResponseAllOfInvoiceLineItemsInnerWithDefaults() *GetInvoices200ResponseAllOfInvoiceLineItemsInner`

NewGetInvoices200ResponseAllOfInvoiceLineItemsInnerWithDefaults instantiates a new GetInvoices200ResponseAllOfInvoiceLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetInvoiceId() int64`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetInvoiceIdOk() (*int64, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetInvoiceId(v int64)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetRefType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetStartDate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetItemId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetItemName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemDescription

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetProductId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductCode

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetItemSeller

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemSeller() string`

GetItemSeller returns the ItemSeller field if non-nil, zero value otherwise.

### GetItemSellerOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemSellerOk() (*string, bool)`

GetItemSellerOk returns a tuple with the ItemSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSeller

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemSeller(v string)`

SetItemSeller sets ItemSeller field to given value.

### HasItemSeller

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemSeller() bool`

HasItemSeller returns a boolean if a field has been set.

### SetItemSellerNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemSellerNil(b bool)`

 SetItemSellerNil sets the value for ItemSeller to be an explicit nil

### UnsetItemSeller
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetItemSeller()`

UnsetItemSeller ensures that no value is present for ItemSeller, not even an explicit nil
### GetItemAction

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemAction() string`

GetItemAction returns the ItemAction field if non-nil, zero value otherwise.

### GetItemActionOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemActionOk() (*string, bool)`

GetItemActionOk returns a tuple with the ItemAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAction

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemAction(v string)`

SetItemAction sets ItemAction field to given value.

### HasItemAction

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemAction() bool`

HasItemAction returns a boolean if a field has been set.

### SetItemActionNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemActionNil(b bool)`

 SetItemActionNil sets the value for ItemAction to be an explicit nil

### UnsetItemAction
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetItemAction()`

UnsetItemAction ensures that no value is present for ItemAction, not even an explicit nil
### GetExternalId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRateId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateId() string`

GetRateId returns the RateId field if non-nil, zero value otherwise.

### GetRateIdOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateIdOk() (*string, bool)`

GetRateIdOk returns a tuple with the RateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRateId(v string)`

SetRateId sets RateId field to given value.

### HasRateId

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRateId() bool`

HasRateId returns a boolean if a field has been set.

### GetRateClass

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateClass() string`

GetRateClass returns the RateClass field if non-nil, zero value otherwise.

### GetRateClassOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateClassOk() (*string, bool)`

GetRateClassOk returns a tuple with the RateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateClass

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRateClass(v string)`

SetRateClass sets RateClass field to given value.

### HasRateClass

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRateClass() bool`

HasRateClass returns a boolean if a field has been set.

### SetRateClassNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRateClassNil(b bool)`

 SetRateClassNil sets the value for RateClass to be an explicit nil

### UnsetRateClass
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetRateClass()`

UnsetRateClass ensures that no value is present for RateClass, not even an explicit nil
### GetRateUnit

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateUnit() string`

GetRateUnit returns the RateUnit field if non-nil, zero value otherwise.

### GetRateUnitOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateUnitOk() (*string, bool)`

GetRateUnitOk returns a tuple with the RateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnit

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRateUnit(v string)`

SetRateUnit sets RateUnit field to given value.

### HasRateUnit

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRateUnit() bool`

HasRateUnit returns a boolean if a field has been set.

### GetRateTerm

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateTerm() string`

GetRateTerm returns the RateTerm field if non-nil, zero value otherwise.

### GetRateTermOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRateTermOk() (*string, bool)`

GetRateTermOk returns a tuple with the RateTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTerm

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRateTerm(v string)`

SetRateTerm sets RateTerm field to given value.

### HasRateTerm

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRateTerm() bool`

HasRateTerm returns a boolean if a field has been set.

### SetRateTermNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRateTermNil(b bool)`

 SetRateTermNil sets the value for RateTerm to be an explicit nil

### UnsetRateTerm
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetRateTerm()`

UnsetRateTerm ensures that no value is present for RateTerm, not even an explicit nil
### GetUsageType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUsageCategory

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### GetUsageService

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetUsageService() string`

GetUsageService returns the UsageService field if non-nil, zero value otherwise.

### GetUsageServiceOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetUsageServiceOk() (*string, bool)`

GetUsageServiceOk returns a tuple with the UsageService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageService

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetUsageService(v string)`

SetUsageService sets UsageService field to given value.

### HasUsageService

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasUsageService() bool`

HasUsageService returns a boolean if a field has been set.

### GetItemUsage

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemUsage() float32`

GetItemUsage returns the ItemUsage field if non-nil, zero value otherwise.

### GetItemUsageOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemUsageOk() (*float32, bool)`

GetItemUsageOk returns a tuple with the ItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsage

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemUsage(v float32)`

SetItemUsage sets ItemUsage field to given value.

### HasItemUsage

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemUsage() bool`

HasItemUsage returns a boolean if a field has been set.

### GetItemRate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemRate() float32`

GetItemRate returns the ItemRate field if non-nil, zero value otherwise.

### GetItemRateOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemRateOk() (*float32, bool)`

GetItemRateOk returns a tuple with the ItemRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemRate(v float32)`

SetItemRate sets ItemRate field to given value.

### HasItemRate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemRate() bool`

HasItemRate returns a boolean if a field has been set.

### GetItemCost

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemCost() float32`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemCostOk() (*float32, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemCost(v float32)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### GetItemPrice

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemPrice() float32`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemPriceOk() (*float32, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemPrice(v float32)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetItemTax

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemTax() int64`

GetItemTax returns the ItemTax field if non-nil, zero value otherwise.

### GetItemTaxOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemTaxOk() (*int64, bool)`

GetItemTaxOk returns a tuple with the ItemTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTax

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemTax(v int64)`

SetItemTax sets ItemTax field to given value.

### HasItemTax

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemTax() bool`

HasItemTax returns a boolean if a field has been set.

### GetItemTerm

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemTerm() string`

GetItemTerm returns the ItemTerm field if non-nil, zero value otherwise.

### GetItemTermOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetItemTermOk() (*string, bool)`

GetItemTermOk returns a tuple with the ItemTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTerm

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemTerm(v string)`

SetItemTerm sets ItemTerm field to given value.

### HasItemTerm

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasItemTerm() bool`

HasItemTerm returns a boolean if a field has been set.

### SetItemTermNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetItemTermNil(b bool)`

 SetItemTermNil sets the value for ItemTerm to be an explicit nil

### UnsetItemTerm
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetItemTerm()`

UnsetItemTerm ensures that no value is present for ItemTerm, not even an explicit nil
### GetTaxType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetRegionCode

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCurrency

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetConversionRate() int64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetConversionRateOk() (*int64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetConversionRate(v int64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInvoices200ResponseAllOfInvoiceLineItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


