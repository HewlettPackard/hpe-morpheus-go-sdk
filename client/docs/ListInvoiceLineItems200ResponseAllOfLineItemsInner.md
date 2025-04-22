# ListInvoiceLineItems200ResponseAllOfLineItemsInner

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
**ItemType** | Pointer to **string** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ItemSeller** | Pointer to **string** |  | [optional] 
**ItemAction** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RateId** | Pointer to **string** |  | [optional] 
**RateClass** | Pointer to **string** |  | [optional] 
**RateUnit** | Pointer to **string** |  | [optional] 
**RateTerm** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 
**UsageCategory** | Pointer to **string** |  | [optional] 
**UsageService** | Pointer to **string** |  | [optional] 
**ItemUsage** | Pointer to **int64** |  | [optional] 
**ItemRate** | Pointer to **float32** |  | [optional] 
**ItemCost** | Pointer to **float32** |  | [optional] 
**ItemPriceRate** | Pointer to **float32** |  | [optional] 
**ItemPrice** | Pointer to **float32** |  | [optional] 
**ItemTax** | Pointer to **int64** |  | [optional] 
**ItemTerm** | Pointer to **string** |  | [optional] 
**TaxType** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListInvoiceLineItems200ResponseAllOfLineItemsInner

`func NewListInvoiceLineItems200ResponseAllOfLineItemsInner() *ListInvoiceLineItems200ResponseAllOfLineItemsInner`

NewListInvoiceLineItems200ResponseAllOfLineItemsInner instantiates a new ListInvoiceLineItems200ResponseAllOfLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoiceLineItems200ResponseAllOfLineItemsInnerWithDefaults

`func NewListInvoiceLineItems200ResponseAllOfLineItemsInnerWithDefaults() *ListInvoiceLineItems200ResponseAllOfLineItemsInner`

NewListInvoiceLineItems200ResponseAllOfLineItemsInnerWithDefaults instantiates a new ListInvoiceLineItems200ResponseAllOfLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetInvoiceId() int64`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetInvoiceIdOk() (*int64, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetInvoiceId(v int64)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetRefType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetStartDate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetItemId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetItemName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemDescription

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetProductId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductCode

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetItemSeller

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemSeller() string`

GetItemSeller returns the ItemSeller field if non-nil, zero value otherwise.

### GetItemSellerOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemSellerOk() (*string, bool)`

GetItemSellerOk returns a tuple with the ItemSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSeller

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemSeller(v string)`

SetItemSeller sets ItemSeller field to given value.

### HasItemSeller

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemSeller() bool`

HasItemSeller returns a boolean if a field has been set.

### GetItemAction

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemAction() string`

GetItemAction returns the ItemAction field if non-nil, zero value otherwise.

### GetItemActionOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemActionOk() (*string, bool)`

GetItemActionOk returns a tuple with the ItemAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAction

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemAction(v string)`

SetItemAction sets ItemAction field to given value.

### HasItemAction

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemAction() bool`

HasItemAction returns a boolean if a field has been set.

### GetExternalId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRateId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateId() string`

GetRateId returns the RateId field if non-nil, zero value otherwise.

### GetRateIdOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateIdOk() (*string, bool)`

GetRateIdOk returns a tuple with the RateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRateId(v string)`

SetRateId sets RateId field to given value.

### HasRateId

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRateId() bool`

HasRateId returns a boolean if a field has been set.

### GetRateClass

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateClass() string`

GetRateClass returns the RateClass field if non-nil, zero value otherwise.

### GetRateClassOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateClassOk() (*string, bool)`

GetRateClassOk returns a tuple with the RateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateClass

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRateClass(v string)`

SetRateClass sets RateClass field to given value.

### HasRateClass

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRateClass() bool`

HasRateClass returns a boolean if a field has been set.

### GetRateUnit

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateUnit() string`

GetRateUnit returns the RateUnit field if non-nil, zero value otherwise.

### GetRateUnitOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateUnitOk() (*string, bool)`

GetRateUnitOk returns a tuple with the RateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnit

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRateUnit(v string)`

SetRateUnit sets RateUnit field to given value.

### HasRateUnit

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRateUnit() bool`

HasRateUnit returns a boolean if a field has been set.

### GetRateTerm

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateTerm() string`

GetRateTerm returns the RateTerm field if non-nil, zero value otherwise.

### GetRateTermOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRateTermOk() (*string, bool)`

GetRateTermOk returns a tuple with the RateTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTerm

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRateTerm(v string)`

SetRateTerm sets RateTerm field to given value.

### HasRateTerm

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRateTerm() bool`

HasRateTerm returns a boolean if a field has been set.

### GetUsageType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUsageCategory

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### GetUsageService

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetUsageService() string`

GetUsageService returns the UsageService field if non-nil, zero value otherwise.

### GetUsageServiceOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetUsageServiceOk() (*string, bool)`

GetUsageServiceOk returns a tuple with the UsageService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageService

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetUsageService(v string)`

SetUsageService sets UsageService field to given value.

### HasUsageService

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasUsageService() bool`

HasUsageService returns a boolean if a field has been set.

### GetItemUsage

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemUsage() int64`

GetItemUsage returns the ItemUsage field if non-nil, zero value otherwise.

### GetItemUsageOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemUsageOk() (*int64, bool)`

GetItemUsageOk returns a tuple with the ItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsage

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemUsage(v int64)`

SetItemUsage sets ItemUsage field to given value.

### HasItemUsage

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemUsage() bool`

HasItemUsage returns a boolean if a field has been set.

### GetItemRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemRate() float32`

GetItemRate returns the ItemRate field if non-nil, zero value otherwise.

### GetItemRateOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemRateOk() (*float32, bool)`

GetItemRateOk returns a tuple with the ItemRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemRate(v float32)`

SetItemRate sets ItemRate field to given value.

### HasItemRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemRate() bool`

HasItemRate returns a boolean if a field has been set.

### GetItemCost

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemCost() float32`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemCostOk() (*float32, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemCost(v float32)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### GetItemPriceRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemPriceRate() float32`

GetItemPriceRate returns the ItemPriceRate field if non-nil, zero value otherwise.

### GetItemPriceRateOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemPriceRateOk() (*float32, bool)`

GetItemPriceRateOk returns a tuple with the ItemPriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemPriceRate(v float32)`

SetItemPriceRate sets ItemPriceRate field to given value.

### HasItemPriceRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemPriceRate() bool`

HasItemPriceRate returns a boolean if a field has been set.

### GetItemPrice

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemPrice() float32`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemPriceOk() (*float32, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemPrice(v float32)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetItemTax

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemTax() int64`

GetItemTax returns the ItemTax field if non-nil, zero value otherwise.

### GetItemTaxOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemTaxOk() (*int64, bool)`

GetItemTaxOk returns a tuple with the ItemTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTax

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemTax(v int64)`

SetItemTax sets ItemTax field to given value.

### HasItemTax

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemTax() bool`

HasItemTax returns a boolean if a field has been set.

### GetItemTerm

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemTerm() string`

GetItemTerm returns the ItemTerm field if non-nil, zero value otherwise.

### GetItemTermOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetItemTermOk() (*string, bool)`

GetItemTermOk returns a tuple with the ItemTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTerm

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetItemTerm(v string)`

SetItemTerm sets ItemTerm field to given value.

### HasItemTerm

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasItemTerm() bool`

HasItemTerm returns a boolean if a field has been set.

### GetTaxType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetRegionCode

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCurrency

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetConversionRate() int64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetConversionRateOk() (*int64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetConversionRate(v int64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListInvoiceLineItems200ResponseAllOfLineItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


