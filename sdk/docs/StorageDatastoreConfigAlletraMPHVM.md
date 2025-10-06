# StorageDatastoreConfigAlletraMPHVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRansomware** | **bool** | Enable ransomware protection for this datastore | 
**ProtocolType** | **string** | Storage protocol to use, can be &#39;iSCSI&#39; or &#39;FC&#39; | 

## Methods

### NewStorageDatastoreConfigAlletraMPHVM

`func NewStorageDatastoreConfigAlletraMPHVM(enableRansomware bool, protocolType string, ) *StorageDatastoreConfigAlletraMPHVM`

NewStorageDatastoreConfigAlletraMPHVM instantiates a new StorageDatastoreConfigAlletraMPHVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDatastoreConfigAlletraMPHVMWithDefaults

`func NewStorageDatastoreConfigAlletraMPHVMWithDefaults() *StorageDatastoreConfigAlletraMPHVM`

NewStorageDatastoreConfigAlletraMPHVMWithDefaults instantiates a new StorageDatastoreConfigAlletraMPHVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableRansomware

`func (o *StorageDatastoreConfigAlletraMPHVM) GetEnableRansomware() bool`

GetEnableRansomware returns the EnableRansomware field if non-nil, zero value otherwise.

### GetEnableRansomwareOk

`func (o *StorageDatastoreConfigAlletraMPHVM) GetEnableRansomwareOk() (*bool, bool)`

GetEnableRansomwareOk returns a tuple with the EnableRansomware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRansomware

`func (o *StorageDatastoreConfigAlletraMPHVM) SetEnableRansomware(v bool)`

SetEnableRansomware sets EnableRansomware field to given value.


### GetProtocolType

`func (o *StorageDatastoreConfigAlletraMPHVM) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *StorageDatastoreConfigAlletraMPHVM) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *StorageDatastoreConfigAlletraMPHVM) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


