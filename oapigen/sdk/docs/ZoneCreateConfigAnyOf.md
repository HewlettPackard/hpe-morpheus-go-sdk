# ZoneCreateConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | The URL used by workloads provisioned in the cloud for interacting with the Morpheus appliance. | [optional] 
**DatacenterName** | Pointer to **string** | A custom name used to reference the datacenter for the cloud. | [optional] 
**ExternalId** | Pointer to **NullableString** | The external id of the cloud | [optional] 
**InventoryLevel** | Pointer to **string** | Whether to import existing virtual machines. | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout to use for the console | [optional] 
**Endpoint** | **string** | AWS endpoint | 
**AccessKey** | Pointer to **string** | AWS access key | [optional] 
**SecretKey** | Pointer to **string** | AWS secret key | [optional] 
**UseHostCredentials** | Pointer to **string** | Whether to use the IAM profile associated with the Morpheus server or not | [optional] [default to "on"]
**EbsEncryption** | Pointer to **string** | Determines whether to configure default EBS volume encryption or not | [optional] [default to "on"]
**StsAssumeRole** | Pointer to **string** | The AWS IAM role ARN to assume for authentication | [optional] 
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the AWS cloud | [optional] 
**Vpc** | Pointer to **string** | The VPC ID for a specific VPC | [optional] 

## Methods

### NewZoneCreateConfigAnyOf

`func NewZoneCreateConfigAnyOf(endpoint string, ) *ZoneCreateConfigAnyOf`

NewZoneCreateConfigAnyOf instantiates a new ZoneCreateConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateConfigAnyOfWithDefaults

`func NewZoneCreateConfigAnyOfWithDefaults() *ZoneCreateConfigAnyOf`

NewZoneCreateConfigAnyOfWithDefaults instantiates a new ZoneCreateConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ZoneCreateConfigAnyOf) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneCreateConfigAnyOf) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneCreateConfigAnyOf) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneCreateConfigAnyOf) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneCreateConfigAnyOf) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneCreateConfigAnyOf) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneCreateConfigAnyOf) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneCreateConfigAnyOf) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *ZoneCreateConfigAnyOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ZoneCreateConfigAnyOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ZoneCreateConfigAnyOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ZoneCreateConfigAnyOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ZoneCreateConfigAnyOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ZoneCreateConfigAnyOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *ZoneCreateConfigAnyOf) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneCreateConfigAnyOf) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneCreateConfigAnyOf) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneCreateConfigAnyOf) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ZoneCreateConfigAnyOf) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ZoneCreateConfigAnyOf) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ZoneCreateConfigAnyOf) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ZoneCreateConfigAnyOf) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetEndpoint

`func (o *ZoneCreateConfigAnyOf) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ZoneCreateConfigAnyOf) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ZoneCreateConfigAnyOf) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *ZoneCreateConfigAnyOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ZoneCreateConfigAnyOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ZoneCreateConfigAnyOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ZoneCreateConfigAnyOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *ZoneCreateConfigAnyOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ZoneCreateConfigAnyOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ZoneCreateConfigAnyOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ZoneCreateConfigAnyOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *ZoneCreateConfigAnyOf) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *ZoneCreateConfigAnyOf) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *ZoneCreateConfigAnyOf) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *ZoneCreateConfigAnyOf) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *ZoneCreateConfigAnyOf) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *ZoneCreateConfigAnyOf) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *ZoneCreateConfigAnyOf) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *ZoneCreateConfigAnyOf) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *ZoneCreateConfigAnyOf) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *ZoneCreateConfigAnyOf) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *ZoneCreateConfigAnyOf) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *ZoneCreateConfigAnyOf) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneCreateConfigAnyOf) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneCreateConfigAnyOf) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneCreateConfigAnyOf) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneCreateConfigAnyOf) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *ZoneCreateConfigAnyOf) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ZoneCreateConfigAnyOf) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ZoneCreateConfigAnyOf) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *ZoneCreateConfigAnyOf) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


