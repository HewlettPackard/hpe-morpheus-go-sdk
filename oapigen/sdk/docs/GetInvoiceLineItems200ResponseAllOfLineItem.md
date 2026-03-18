# GetInvoiceLineItems200ResponseAllOfLineItem

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
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**ItemName** | Pointer to **NullableString** |  | [optional] 
**ItemDescription** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**ProductCode** | Pointer to **NullableString** |  | [optional] 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**ItemSeller** | Pointer to **NullableString** |  | [optional] 
**ItemAction** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RateId** | Pointer to **NullableString** |  | [optional] 
**RateClass** | Pointer to **NullableString** |  | [optional] 
**RateUnit** | Pointer to **string** |  | [optional] 
**RateTerm** | Pointer to **NullableString** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 
**UsageCategory** | Pointer to **string** |  | [optional] 
**UsageService** | Pointer to **NullableString** |  | [optional] 
**ItemUsage** | Pointer to **int64** |  | [optional] 
**ItemRate** | Pointer to **float32** |  | [optional] 
**ItemCost** | Pointer to **float32** |  | [optional] 
**ItemPriceRate** | Pointer to **float32** |  | [optional] 
**ItemPrice** | Pointer to **float32** |  | [optional] 
**ItemTax** | Pointer to **int64** |  | [optional] 
**ItemTerm** | Pointer to **NullableString** |  | [optional] 
**TaxType** | Pointer to **NullableString** |  | [optional] 
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetInvoiceLineItems200ResponseAllOfLineItem

`func NewGetInvoiceLineItems200ResponseAllOfLineItem() *GetInvoiceLineItems200ResponseAllOfLineItem`

NewGetInvoiceLineItems200ResponseAllOfLineItem instantiates a new GetInvoiceLineItems200ResponseAllOfLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceLineItems200ResponseAllOfLineItemWithDefaults

`func NewGetInvoiceLineItems200ResponseAllOfLineItemWithDefaults() *GetInvoiceLineItems200ResponseAllOfLineItem`

NewGetInvoiceLineItems200ResponseAllOfLineItemWithDefaults instantiates a new GetInvoiceLineItems200ResponseAllOfLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetInvoiceId() int64`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetInvoiceIdOk() (*int64, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetInvoiceId(v int64)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetRefType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetStartDate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetItemId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetItemName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### SetItemNameNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemNameNil(b bool)`

 SetItemNameNil sets the value for ItemName to be an explicit nil

### UnsetItemName
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemName()`

UnsetItemName ensures that no value is present for ItemName, not even an explicit nil
### GetItemDescription

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### SetItemDescriptionNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemDescriptionNil(b bool)`

 SetItemDescriptionNil sets the value for ItemDescription to be an explicit nil

### UnsetItemDescription
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemDescription()`

UnsetItemDescription ensures that no value is present for ItemDescription, not even an explicit nil
### GetProductId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductCode

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCodeNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetProductCodeNil(b bool)`

 SetProductCodeNil sets the value for ProductCode to be an explicit nil

### UnsetProductCode
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetProductCode()`

UnsetProductCode ensures that no value is present for ProductCode, not even an explicit nil
### GetProductName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetItemSeller

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemSeller() string`

GetItemSeller returns the ItemSeller field if non-nil, zero value otherwise.

### GetItemSellerOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemSellerOk() (*string, bool)`

GetItemSellerOk returns a tuple with the ItemSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSeller

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemSeller(v string)`

SetItemSeller sets ItemSeller field to given value.

### HasItemSeller

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemSeller() bool`

HasItemSeller returns a boolean if a field has been set.

### SetItemSellerNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemSellerNil(b bool)`

 SetItemSellerNil sets the value for ItemSeller to be an explicit nil

### UnsetItemSeller
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemSeller()`

UnsetItemSeller ensures that no value is present for ItemSeller, not even an explicit nil
### GetItemAction

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemAction() string`

GetItemAction returns the ItemAction field if non-nil, zero value otherwise.

### GetItemActionOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemActionOk() (*string, bool)`

GetItemActionOk returns a tuple with the ItemAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAction

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemAction(v string)`

SetItemAction sets ItemAction field to given value.

### HasItemAction

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemAction() bool`

HasItemAction returns a boolean if a field has been set.

### SetItemActionNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemActionNil(b bool)`

 SetItemActionNil sets the value for ItemAction to be an explicit nil

### UnsetItemAction
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemAction()`

UnsetItemAction ensures that no value is present for ItemAction, not even an explicit nil
### GetExternalId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRateId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateId() string`

GetRateId returns the RateId field if non-nil, zero value otherwise.

### GetRateIdOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateIdOk() (*string, bool)`

GetRateIdOk returns a tuple with the RateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateId(v string)`

SetRateId sets RateId field to given value.

### HasRateId

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRateId() bool`

HasRateId returns a boolean if a field has been set.

### SetRateIdNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateIdNil(b bool)`

 SetRateIdNil sets the value for RateId to be an explicit nil

### UnsetRateId
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetRateId()`

UnsetRateId ensures that no value is present for RateId, not even an explicit nil
### GetRateClass

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateClass() string`

GetRateClass returns the RateClass field if non-nil, zero value otherwise.

### GetRateClassOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateClassOk() (*string, bool)`

GetRateClassOk returns a tuple with the RateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateClass

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateClass(v string)`

SetRateClass sets RateClass field to given value.

### HasRateClass

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRateClass() bool`

HasRateClass returns a boolean if a field has been set.

### SetRateClassNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateClassNil(b bool)`

 SetRateClassNil sets the value for RateClass to be an explicit nil

### UnsetRateClass
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetRateClass()`

UnsetRateClass ensures that no value is present for RateClass, not even an explicit nil
### GetRateUnit

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateUnit() string`

GetRateUnit returns the RateUnit field if non-nil, zero value otherwise.

### GetRateUnitOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateUnitOk() (*string, bool)`

GetRateUnitOk returns a tuple with the RateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnit

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateUnit(v string)`

SetRateUnit sets RateUnit field to given value.

### HasRateUnit

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRateUnit() bool`

HasRateUnit returns a boolean if a field has been set.

### GetRateTerm

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateTerm() string`

GetRateTerm returns the RateTerm field if non-nil, zero value otherwise.

### GetRateTermOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRateTermOk() (*string, bool)`

GetRateTermOk returns a tuple with the RateTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTerm

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateTerm(v string)`

SetRateTerm sets RateTerm field to given value.

### HasRateTerm

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRateTerm() bool`

HasRateTerm returns a boolean if a field has been set.

### SetRateTermNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRateTermNil(b bool)`

 SetRateTermNil sets the value for RateTerm to be an explicit nil

### UnsetRateTerm
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetRateTerm()`

UnsetRateTerm ensures that no value is present for RateTerm, not even an explicit nil
### GetUsageType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUsageCategory

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### GetUsageService

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetUsageService() string`

GetUsageService returns the UsageService field if non-nil, zero value otherwise.

### GetUsageServiceOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetUsageServiceOk() (*string, bool)`

GetUsageServiceOk returns a tuple with the UsageService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageService

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetUsageService(v string)`

SetUsageService sets UsageService field to given value.

### HasUsageService

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasUsageService() bool`

HasUsageService returns a boolean if a field has been set.

### SetUsageServiceNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetUsageServiceNil(b bool)`

 SetUsageServiceNil sets the value for UsageService to be an explicit nil

### UnsetUsageService
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetUsageService()`

UnsetUsageService ensures that no value is present for UsageService, not even an explicit nil
### GetItemUsage

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemUsage() int64`

GetItemUsage returns the ItemUsage field if non-nil, zero value otherwise.

### GetItemUsageOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemUsageOk() (*int64, bool)`

GetItemUsageOk returns a tuple with the ItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsage

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemUsage(v int64)`

SetItemUsage sets ItemUsage field to given value.

### HasItemUsage

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemUsage() bool`

HasItemUsage returns a boolean if a field has been set.

### GetItemRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemRate() float32`

GetItemRate returns the ItemRate field if non-nil, zero value otherwise.

### GetItemRateOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemRateOk() (*float32, bool)`

GetItemRateOk returns a tuple with the ItemRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemRate(v float32)`

SetItemRate sets ItemRate field to given value.

### HasItemRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemRate() bool`

HasItemRate returns a boolean if a field has been set.

### GetItemCost

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemCost() float32`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemCostOk() (*float32, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemCost(v float32)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### GetItemPriceRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemPriceRate() float32`

GetItemPriceRate returns the ItemPriceRate field if non-nil, zero value otherwise.

### GetItemPriceRateOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemPriceRateOk() (*float32, bool)`

GetItemPriceRateOk returns a tuple with the ItemPriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemPriceRate(v float32)`

SetItemPriceRate sets ItemPriceRate field to given value.

### HasItemPriceRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemPriceRate() bool`

HasItemPriceRate returns a boolean if a field has been set.

### GetItemPrice

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemPrice() float32`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemPriceOk() (*float32, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemPrice(v float32)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetItemTax

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemTax() int64`

GetItemTax returns the ItemTax field if non-nil, zero value otherwise.

### GetItemTaxOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemTaxOk() (*int64, bool)`

GetItemTaxOk returns a tuple with the ItemTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTax

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemTax(v int64)`

SetItemTax sets ItemTax field to given value.

### HasItemTax

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemTax() bool`

HasItemTax returns a boolean if a field has been set.

### GetItemTerm

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemTerm() string`

GetItemTerm returns the ItemTerm field if non-nil, zero value otherwise.

### GetItemTermOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetItemTermOk() (*string, bool)`

GetItemTermOk returns a tuple with the ItemTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTerm

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemTerm(v string)`

SetItemTerm sets ItemTerm field to given value.

### HasItemTerm

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasItemTerm() bool`

HasItemTerm returns a boolean if a field has been set.

### SetItemTermNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetItemTermNil(b bool)`

 SetItemTermNil sets the value for ItemTerm to be an explicit nil

### UnsetItemTerm
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetItemTerm()`

UnsetItemTerm ensures that no value is present for ItemTerm, not even an explicit nil
### GetTaxType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetRegionCode

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetCurrency

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetConversionRate() int64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetConversionRateOk() (*int64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetConversionRate(v int64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInvoiceLineItems200ResponseAllOfLineItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


