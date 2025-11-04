# AlletraMPHVMDatastoreConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enableransomware** | Pointer to **string** | Enable ransomware protection for this datastore | [optional] 
**ProtocolType** | **string** | Storage protocol to use, can be &#39;iSCSI&#39; or &#39;FC&#39; | 

## Methods

### NewAlletraMPHVMDatastoreConfiguration

`func NewAlletraMPHVMDatastoreConfiguration(protocolType string, ) *AlletraMPHVMDatastoreConfiguration`

NewAlletraMPHVMDatastoreConfiguration instantiates a new AlletraMPHVMDatastoreConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlletraMPHVMDatastoreConfigurationWithDefaults

`func NewAlletraMPHVMDatastoreConfigurationWithDefaults() *AlletraMPHVMDatastoreConfiguration`

NewAlletraMPHVMDatastoreConfigurationWithDefaults instantiates a new AlletraMPHVMDatastoreConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableransomware

`func (o *AlletraMPHVMDatastoreConfiguration) GetEnableransomware() string`

GetEnableransomware returns the Enableransomware field if non-nil, zero value otherwise.

### GetEnableransomwareOk

`func (o *AlletraMPHVMDatastoreConfiguration) GetEnableransomwareOk() (*string, bool)`

GetEnableransomwareOk returns a tuple with the Enableransomware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableransomware

`func (o *AlletraMPHVMDatastoreConfiguration) SetEnableransomware(v string)`

SetEnableransomware sets Enableransomware field to given value.

### HasEnableransomware

`func (o *AlletraMPHVMDatastoreConfiguration) HasEnableransomware() bool`

HasEnableransomware returns a boolean if a field has been set.

### GetProtocolType

`func (o *AlletraMPHVMDatastoreConfiguration) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *AlletraMPHVMDatastoreConfiguration) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *AlletraMPHVMDatastoreConfiguration) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


