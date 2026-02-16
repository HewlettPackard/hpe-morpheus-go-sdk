# VirtualImageCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publisher** | **string** | The name of the publisher in the Azure Marketplace | 
**Offer** | **string** | The name of the offer in the Azure Marketplace | 
**Sku** | **string** | The name of the sku in the Azure Marketplace | 
**Version** | **string** | The name of the version in the Azure Marketplace | 

## Methods

### NewVirtualImageCreateConfig

`func NewVirtualImageCreateConfig(publisher string, offer string, sku string, version string, ) *VirtualImageCreateConfig`

NewVirtualImageCreateConfig instantiates a new VirtualImageCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageCreateConfigWithDefaults

`func NewVirtualImageCreateConfigWithDefaults() *VirtualImageCreateConfig`

NewVirtualImageCreateConfigWithDefaults instantiates a new VirtualImageCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisher

`func (o *VirtualImageCreateConfig) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *VirtualImageCreateConfig) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *VirtualImageCreateConfig) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetOffer

`func (o *VirtualImageCreateConfig) GetOffer() string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *VirtualImageCreateConfig) GetOfferOk() (*string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *VirtualImageCreateConfig) SetOffer(v string)`

SetOffer sets Offer field to given value.


### GetSku

`func (o *VirtualImageCreateConfig) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *VirtualImageCreateConfig) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *VirtualImageCreateConfig) SetSku(v string)`

SetSku sets Sku field to given value.


### GetVersion

`func (o *VirtualImageCreateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualImageCreateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualImageCreateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


