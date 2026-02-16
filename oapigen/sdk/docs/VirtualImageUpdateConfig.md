# VirtualImageUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publisher** | **string** | The name of the publisher in the Azure Marketplace | 
**Offer** | **string** | The name of the offer in the Azure Marketplace | 
**Sku** | **string** | The name of the sku in the Azure Marketplace | 
**Version** | **string** | The name of the version in the Azure Marketplace | 

## Methods

### NewVirtualImageUpdateConfig

`func NewVirtualImageUpdateConfig(publisher string, offer string, sku string, version string, ) *VirtualImageUpdateConfig`

NewVirtualImageUpdateConfig instantiates a new VirtualImageUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageUpdateConfigWithDefaults

`func NewVirtualImageUpdateConfigWithDefaults() *VirtualImageUpdateConfig`

NewVirtualImageUpdateConfigWithDefaults instantiates a new VirtualImageUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisher

`func (o *VirtualImageUpdateConfig) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *VirtualImageUpdateConfig) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *VirtualImageUpdateConfig) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetOffer

`func (o *VirtualImageUpdateConfig) GetOffer() string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *VirtualImageUpdateConfig) GetOfferOk() (*string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *VirtualImageUpdateConfig) SetOffer(v string)`

SetOffer sets Offer field to given value.


### GetSku

`func (o *VirtualImageUpdateConfig) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *VirtualImageUpdateConfig) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *VirtualImageUpdateConfig) SetSku(v string)`

SetSku sets Sku field to given value.


### GetVersion

`func (o *VirtualImageUpdateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualImageUpdateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualImageUpdateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


