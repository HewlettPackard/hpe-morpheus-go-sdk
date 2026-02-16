# GFS2DatastoreConfiguration3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDevice** | **string** | Block device for target GFS2. | 
**AllowReformat** | Pointer to **bool** | Allow reformatting of the device if it is already formatted. | [optional] 

## Methods

### NewGFS2DatastoreConfiguration3

`func NewGFS2DatastoreConfiguration3(blockDevice string, ) *GFS2DatastoreConfiguration3`

NewGFS2DatastoreConfiguration3 instantiates a new GFS2DatastoreConfiguration3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFS2DatastoreConfiguration3WithDefaults

`func NewGFS2DatastoreConfiguration3WithDefaults() *GFS2DatastoreConfiguration3`

NewGFS2DatastoreConfiguration3WithDefaults instantiates a new GFS2DatastoreConfiguration3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDevice

`func (o *GFS2DatastoreConfiguration3) GetBlockDevice() string`

GetBlockDevice returns the BlockDevice field if non-nil, zero value otherwise.

### GetBlockDeviceOk

`func (o *GFS2DatastoreConfiguration3) GetBlockDeviceOk() (*string, bool)`

GetBlockDeviceOk returns a tuple with the BlockDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDevice

`func (o *GFS2DatastoreConfiguration3) SetBlockDevice(v string)`

SetBlockDevice sets BlockDevice field to given value.


### GetAllowReformat

`func (o *GFS2DatastoreConfiguration3) GetAllowReformat() bool`

GetAllowReformat returns the AllowReformat field if non-nil, zero value otherwise.

### GetAllowReformatOk

`func (o *GFS2DatastoreConfiguration3) GetAllowReformatOk() (*bool, bool)`

GetAllowReformatOk returns a tuple with the AllowReformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReformat

`func (o *GFS2DatastoreConfiguration3) SetAllowReformat(v bool)`

SetAllowReformat sets AllowReformat field to given value.

### HasAllowReformat

`func (o *GFS2DatastoreConfiguration3) HasAllowReformat() bool`

HasAllowReformat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


