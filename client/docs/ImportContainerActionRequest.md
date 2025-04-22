# ImportContainerActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageProviderId** | Pointer to **int64** | Optional storage provider to use. | [optional] 

## Methods

### NewImportContainerActionRequest

`func NewImportContainerActionRequest() *ImportContainerActionRequest`

NewImportContainerActionRequest instantiates a new ImportContainerActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportContainerActionRequestWithDefaults

`func NewImportContainerActionRequestWithDefaults() *ImportContainerActionRequest`

NewImportContainerActionRequestWithDefaults instantiates a new ImportContainerActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageProviderId

`func (o *ImportContainerActionRequest) GetStorageProviderId() int64`

GetStorageProviderId returns the StorageProviderId field if non-nil, zero value otherwise.

### GetStorageProviderIdOk

`func (o *ImportContainerActionRequest) GetStorageProviderIdOk() (*int64, bool)`

GetStorageProviderIdOk returns a tuple with the StorageProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderId

`func (o *ImportContainerActionRequest) SetStorageProviderId(v int64)`

SetStorageProviderId sets StorageProviderId field to given value.

### HasStorageProviderId

`func (o *ImportContainerActionRequest) HasStorageProviderId() bool`

HasStorageProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


