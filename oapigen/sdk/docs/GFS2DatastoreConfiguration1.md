# GFS2DatastoreConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDevice** | **string** | Block device for target GFS2. | 
**AllowReformat** | Pointer to **bool** | Allow reformatting of the device if it is already formatted. | [optional] 

## Methods

### NewGFS2DatastoreConfiguration1

`func NewGFS2DatastoreConfiguration1(blockDevice string, ) *GFS2DatastoreConfiguration1`

NewGFS2DatastoreConfiguration1 instantiates a new GFS2DatastoreConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFS2DatastoreConfiguration1WithDefaults

`func NewGFS2DatastoreConfiguration1WithDefaults() *GFS2DatastoreConfiguration1`

NewGFS2DatastoreConfiguration1WithDefaults instantiates a new GFS2DatastoreConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockDevice

`func (o *GFS2DatastoreConfiguration1) GetBlockDevice() string`

GetBlockDevice returns the BlockDevice field if non-nil, zero value otherwise.

### GetBlockDeviceOk

`func (o *GFS2DatastoreConfiguration1) GetBlockDeviceOk() (*string, bool)`

GetBlockDeviceOk returns a tuple with the BlockDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDevice

`func (o *GFS2DatastoreConfiguration1) SetBlockDevice(v string)`

SetBlockDevice sets BlockDevice field to given value.


### GetAllowReformat

`func (o *GFS2DatastoreConfiguration1) GetAllowReformat() bool`

GetAllowReformat returns the AllowReformat field if non-nil, zero value otherwise.

### GetAllowReformatOk

`func (o *GFS2DatastoreConfiguration1) GetAllowReformatOk() (*bool, bool)`

GetAllowReformatOk returns a tuple with the AllowReformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReformat

`func (o *GFS2DatastoreConfiguration1) SetAllowReformat(v bool)`

SetAllowReformat sets AllowReformat field to given value.

### HasAllowReformat

`func (o *GFS2DatastoreConfiguration1) HasAllowReformat() bool`

HasAllowReformat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


