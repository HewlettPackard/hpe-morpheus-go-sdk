# CloudCreateConfigAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | The URL used by workloads provisioned in the cloud for interacting with the Morpheus appliance. | [optional] 
**DatacenterName** | Pointer to **string** | A custom name used to reference the datacenter for the cloud. | [optional] 
**ExternalId** | Pointer to **NullableString** | The external id of the cloud | [optional] 
**InventoryLevel** | Pointer to **string** | Whether to import existing virtual machines. | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout to use for the console | [optional] 
**SubscriberId** | Pointer to **string** | Azure subscriber id | [optional] 
**TenantId** | Pointer to **string** | Azure tenant id | [optional] 
**ClientId** | Pointer to **string** | Azure client id | [optional] 
**ClientSecret** | Pointer to **string** | Azure client secret | [optional] 
**ResourceGroup** | Pointer to **string** | Azure resource group | [optional] 
**RpcMode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCloudCreateConfigAzure

`func NewCloudCreateConfigAzure() *CloudCreateConfigAzure`

NewCloudCreateConfigAzure instantiates a new CloudCreateConfigAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateConfigAzureWithDefaults

`func NewCloudCreateConfigAzureWithDefaults() *CloudCreateConfigAzure`

NewCloudCreateConfigAzureWithDefaults instantiates a new CloudCreateConfigAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *CloudCreateConfigAzure) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *CloudCreateConfigAzure) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *CloudCreateConfigAzure) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *CloudCreateConfigAzure) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *CloudCreateConfigAzure) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CloudCreateConfigAzure) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CloudCreateConfigAzure) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *CloudCreateConfigAzure) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudCreateConfigAzure) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudCreateConfigAzure) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudCreateConfigAzure) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudCreateConfigAzure) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CloudCreateConfigAzure) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CloudCreateConfigAzure) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *CloudCreateConfigAzure) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *CloudCreateConfigAzure) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *CloudCreateConfigAzure) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *CloudCreateConfigAzure) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *CloudCreateConfigAzure) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *CloudCreateConfigAzure) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *CloudCreateConfigAzure) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *CloudCreateConfigAzure) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetSubscriberId

`func (o *CloudCreateConfigAzure) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *CloudCreateConfigAzure) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *CloudCreateConfigAzure) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *CloudCreateConfigAzure) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTenantId

`func (o *CloudCreateConfigAzure) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudCreateConfigAzure) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudCreateConfigAzure) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudCreateConfigAzure) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *CloudCreateConfigAzure) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CloudCreateConfigAzure) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CloudCreateConfigAzure) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CloudCreateConfigAzure) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CloudCreateConfigAzure) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CloudCreateConfigAzure) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CloudCreateConfigAzure) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CloudCreateConfigAzure) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResourceGroup

`func (o *CloudCreateConfigAzure) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *CloudCreateConfigAzure) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *CloudCreateConfigAzure) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *CloudCreateConfigAzure) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetRpcMode

`func (o *CloudCreateConfigAzure) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *CloudCreateConfigAzure) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *CloudCreateConfigAzure) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *CloudCreateConfigAzure) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### SetRpcModeNil

`func (o *CloudCreateConfigAzure) SetRpcModeNil(b bool)`

 SetRpcModeNil sets the value for RpcMode to be an explicit nil

### UnsetRpcMode
`func (o *CloudCreateConfigAzure) UnsetRpcMode()`

UnsetRpcMode ensures that no value is present for RpcMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


