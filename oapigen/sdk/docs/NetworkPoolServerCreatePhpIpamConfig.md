# NetworkPoolServerCreatePhpIpamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | App ID (Your App name in phpIPAM) | 
**InventoryExisting** | Pointer to **string** | Inventory Existing | [optional] [default to "off"]

## Methods

### NewNetworkPoolServerCreatePhpIpamConfig

`func NewNetworkPoolServerCreatePhpIpamConfig(appId string, ) *NetworkPoolServerCreatePhpIpamConfig`

NewNetworkPoolServerCreatePhpIpamConfig instantiates a new NetworkPoolServerCreatePhpIpamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetAppId

`func (o *NetworkPoolServerCreatePhpIpamConfig) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *NetworkPoolServerCreatePhpIpamConfig) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *NetworkPoolServerCreatePhpIpamConfig) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInventoryExisting

`func (o *NetworkPoolServerCreatePhpIpamConfig) GetInventoryExisting() string`

GetInventoryExisting returns the InventoryExisting field if non-nil, zero value otherwise.

### GetInventoryExistingOk

`func (o *NetworkPoolServerCreatePhpIpamConfig) GetInventoryExistingOk() (*string, bool)`

GetInventoryExistingOk returns a tuple with the InventoryExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryExisting

`func (o *NetworkPoolServerCreatePhpIpamConfig) SetInventoryExisting(v string)`

SetInventoryExisting sets InventoryExisting field to given value.

### HasInventoryExisting

`func (o *NetworkPoolServerCreatePhpIpamConfig) HasInventoryExisting() bool`

HasInventoryExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


