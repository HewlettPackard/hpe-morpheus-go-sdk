# LineItem

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

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LineItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *LineItem) GetInvoiceId() int64`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *LineItem) GetInvoiceIdOk() (*int64, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *LineItem) SetInvoiceId(v int64)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *LineItem) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetRefType

`func (o *LineItem) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *LineItem) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *LineItem) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *LineItem) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *LineItem) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *LineItem) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *LineItem) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *LineItem) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *LineItem) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *LineItem) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *LineItem) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *LineItem) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetStartDate

`func (o *LineItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LineItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LineItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LineItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LineItem) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LineItem) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LineItem) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LineItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetItemId

`func (o *LineItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *LineItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *LineItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *LineItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemType

`func (o *LineItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *LineItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *LineItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *LineItem) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetItemName

`func (o *LineItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *LineItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *LineItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *LineItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemDescription

`func (o *LineItem) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *LineItem) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *LineItem) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *LineItem) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetProductId

`func (o *LineItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *LineItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *LineItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *LineItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductCode

`func (o *LineItem) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *LineItem) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *LineItem) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *LineItem) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *LineItem) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *LineItem) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *LineItem) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *LineItem) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetItemSeller

`func (o *LineItem) GetItemSeller() string`

GetItemSeller returns the ItemSeller field if non-nil, zero value otherwise.

### GetItemSellerOk

`func (o *LineItem) GetItemSellerOk() (*string, bool)`

GetItemSellerOk returns a tuple with the ItemSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSeller

`func (o *LineItem) SetItemSeller(v string)`

SetItemSeller sets ItemSeller field to given value.

### HasItemSeller

`func (o *LineItem) HasItemSeller() bool`

HasItemSeller returns a boolean if a field has been set.

### GetItemAction

`func (o *LineItem) GetItemAction() string`

GetItemAction returns the ItemAction field if non-nil, zero value otherwise.

### GetItemActionOk

`func (o *LineItem) GetItemActionOk() (*string, bool)`

GetItemActionOk returns a tuple with the ItemAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAction

`func (o *LineItem) SetItemAction(v string)`

SetItemAction sets ItemAction field to given value.

### HasItemAction

`func (o *LineItem) HasItemAction() bool`

HasItemAction returns a boolean if a field has been set.

### GetExternalId

`func (o *LineItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *LineItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *LineItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *LineItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRateId

`func (o *LineItem) GetRateId() string`

GetRateId returns the RateId field if non-nil, zero value otherwise.

### GetRateIdOk

`func (o *LineItem) GetRateIdOk() (*string, bool)`

GetRateIdOk returns a tuple with the RateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateId

`func (o *LineItem) SetRateId(v string)`

SetRateId sets RateId field to given value.

### HasRateId

`func (o *LineItem) HasRateId() bool`

HasRateId returns a boolean if a field has been set.

### GetRateClass

`func (o *LineItem) GetRateClass() string`

GetRateClass returns the RateClass field if non-nil, zero value otherwise.

### GetRateClassOk

`func (o *LineItem) GetRateClassOk() (*string, bool)`

GetRateClassOk returns a tuple with the RateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateClass

`func (o *LineItem) SetRateClass(v string)`

SetRateClass sets RateClass field to given value.

### HasRateClass

`func (o *LineItem) HasRateClass() bool`

HasRateClass returns a boolean if a field has been set.

### GetRateUnit

`func (o *LineItem) GetRateUnit() string`

GetRateUnit returns the RateUnit field if non-nil, zero value otherwise.

### GetRateUnitOk

`func (o *LineItem) GetRateUnitOk() (*string, bool)`

GetRateUnitOk returns a tuple with the RateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnit

`func (o *LineItem) SetRateUnit(v string)`

SetRateUnit sets RateUnit field to given value.

### HasRateUnit

`func (o *LineItem) HasRateUnit() bool`

HasRateUnit returns a boolean if a field has been set.

### GetRateTerm

`func (o *LineItem) GetRateTerm() string`

GetRateTerm returns the RateTerm field if non-nil, zero value otherwise.

### GetRateTermOk

`func (o *LineItem) GetRateTermOk() (*string, bool)`

GetRateTermOk returns a tuple with the RateTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTerm

`func (o *LineItem) SetRateTerm(v string)`

SetRateTerm sets RateTerm field to given value.

### HasRateTerm

`func (o *LineItem) HasRateTerm() bool`

HasRateTerm returns a boolean if a field has been set.

### GetUsageType

`func (o *LineItem) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *LineItem) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *LineItem) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *LineItem) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUsageCategory

`func (o *LineItem) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *LineItem) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *LineItem) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *LineItem) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### GetUsageService

`func (o *LineItem) GetUsageService() string`

GetUsageService returns the UsageService field if non-nil, zero value otherwise.

### GetUsageServiceOk

`func (o *LineItem) GetUsageServiceOk() (*string, bool)`

GetUsageServiceOk returns a tuple with the UsageService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageService

`func (o *LineItem) SetUsageService(v string)`

SetUsageService sets UsageService field to given value.

### HasUsageService

`func (o *LineItem) HasUsageService() bool`

HasUsageService returns a boolean if a field has been set.

### GetItemUsage

`func (o *LineItem) GetItemUsage() int64`

GetItemUsage returns the ItemUsage field if non-nil, zero value otherwise.

### GetItemUsageOk

`func (o *LineItem) GetItemUsageOk() (*int64, bool)`

GetItemUsageOk returns a tuple with the ItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsage

`func (o *LineItem) SetItemUsage(v int64)`

SetItemUsage sets ItemUsage field to given value.

### HasItemUsage

`func (o *LineItem) HasItemUsage() bool`

HasItemUsage returns a boolean if a field has been set.

### GetItemRate

`func (o *LineItem) GetItemRate() float32`

GetItemRate returns the ItemRate field if non-nil, zero value otherwise.

### GetItemRateOk

`func (o *LineItem) GetItemRateOk() (*float32, bool)`

GetItemRateOk returns a tuple with the ItemRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRate

`func (o *LineItem) SetItemRate(v float32)`

SetItemRate sets ItemRate field to given value.

### HasItemRate

`func (o *LineItem) HasItemRate() bool`

HasItemRate returns a boolean if a field has been set.

### GetItemCost

`func (o *LineItem) GetItemCost() float32`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *LineItem) GetItemCostOk() (*float32, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *LineItem) SetItemCost(v float32)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *LineItem) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### GetItemPriceRate

`func (o *LineItem) GetItemPriceRate() float32`

GetItemPriceRate returns the ItemPriceRate field if non-nil, zero value otherwise.

### GetItemPriceRateOk

`func (o *LineItem) GetItemPriceRateOk() (*float32, bool)`

GetItemPriceRateOk returns a tuple with the ItemPriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceRate

`func (o *LineItem) SetItemPriceRate(v float32)`

SetItemPriceRate sets ItemPriceRate field to given value.

### HasItemPriceRate

`func (o *LineItem) HasItemPriceRate() bool`

HasItemPriceRate returns a boolean if a field has been set.

### GetItemPrice

`func (o *LineItem) GetItemPrice() float32`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *LineItem) GetItemPriceOk() (*float32, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *LineItem) SetItemPrice(v float32)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *LineItem) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetItemTax

`func (o *LineItem) GetItemTax() int64`

GetItemTax returns the ItemTax field if non-nil, zero value otherwise.

### GetItemTaxOk

`func (o *LineItem) GetItemTaxOk() (*int64, bool)`

GetItemTaxOk returns a tuple with the ItemTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTax

`func (o *LineItem) SetItemTax(v int64)`

SetItemTax sets ItemTax field to given value.

### HasItemTax

`func (o *LineItem) HasItemTax() bool`

HasItemTax returns a boolean if a field has been set.

### GetItemTerm

`func (o *LineItem) GetItemTerm() string`

GetItemTerm returns the ItemTerm field if non-nil, zero value otherwise.

### GetItemTermOk

`func (o *LineItem) GetItemTermOk() (*string, bool)`

GetItemTermOk returns a tuple with the ItemTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTerm

`func (o *LineItem) SetItemTerm(v string)`

SetItemTerm sets ItemTerm field to given value.

### HasItemTerm

`func (o *LineItem) HasItemTerm() bool`

HasItemTerm returns a boolean if a field has been set.

### GetTaxType

`func (o *LineItem) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *LineItem) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *LineItem) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *LineItem) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetRegionCode

`func (o *LineItem) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *LineItem) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *LineItem) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *LineItem) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCurrency

`func (o *LineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *LineItem) GetConversionRate() int64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *LineItem) GetConversionRateOk() (*int64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *LineItem) SetConversionRate(v int64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *LineItem) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDateCreated

`func (o *LineItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *LineItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *LineItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *LineItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *LineItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LineItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LineItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *LineItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


