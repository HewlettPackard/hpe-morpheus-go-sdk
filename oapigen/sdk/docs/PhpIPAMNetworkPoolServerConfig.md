# PhpIPAMNetworkPoolServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | App ID (Your App name in phpIPAM) | 
**InventoryExisting** | Pointer to **string** | Inventory Existing | [optional] [default to "off"]

## Methods

### NewPhpIPAMNetworkPoolServerConfig

`func NewPhpIPAMNetworkPoolServerConfig(appId string, ) *PhpIPAMNetworkPoolServerConfig`

NewPhpIPAMNetworkPoolServerConfig instantiates a new PhpIPAMNetworkPoolServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhpIPAMNetworkPoolServerConfigWithDefaults

`func NewPhpIPAMNetworkPoolServerConfigWithDefaults() *PhpIPAMNetworkPoolServerConfig`

NewPhpIPAMNetworkPoolServerConfigWithDefaults instantiates a new PhpIPAMNetworkPoolServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *PhpIPAMNetworkPoolServerConfig) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PhpIPAMNetworkPoolServerConfig) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PhpIPAMNetworkPoolServerConfig) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetInventoryExisting

`func (o *PhpIPAMNetworkPoolServerConfig) GetInventoryExisting() string`

GetInventoryExisting returns the InventoryExisting field if non-nil, zero value otherwise.

### GetInventoryExistingOk

`func (o *PhpIPAMNetworkPoolServerConfig) GetInventoryExistingOk() (*string, bool)`

GetInventoryExistingOk returns a tuple with the InventoryExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryExisting

`func (o *PhpIPAMNetworkPoolServerConfig) SetInventoryExisting(v string)`

SetInventoryExisting sets InventoryExisting field to given value.

### HasInventoryExisting

`func (o *PhpIPAMNetworkPoolServerConfig) HasInventoryExisting() bool`

HasInventoryExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


