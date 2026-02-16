# StorageDatastoreLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** | Type of the location, e.g. &#39;ComputeServer&#39; | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** | Status of the location, e.g. &#39;provisioned&#39;, &#39;provisioning&#39;, &#39;failed&#39;, &#39;warning&#39; | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewStorageDatastoreLocationsInner

`func NewStorageDatastoreLocationsInner() *StorageDatastoreLocationsInner`

NewStorageDatastoreLocationsInner instantiates a new StorageDatastoreLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDatastoreLocationsInnerWithDefaults

`func NewStorageDatastoreLocationsInnerWithDefaults() *StorageDatastoreLocationsInner`

NewStorageDatastoreLocationsInnerWithDefaults instantiates a new StorageDatastoreLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *StorageDatastoreLocationsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *StorageDatastoreLocationsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *StorageDatastoreLocationsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *StorageDatastoreLocationsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *StorageDatastoreLocationsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *StorageDatastoreLocationsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *StorageDatastoreLocationsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *StorageDatastoreLocationsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStatus

`func (o *StorageDatastoreLocationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageDatastoreLocationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageDatastoreLocationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageDatastoreLocationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *StorageDatastoreLocationsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *StorageDatastoreLocationsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *StorageDatastoreLocationsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *StorageDatastoreLocationsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


