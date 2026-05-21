# GenerateSupportBundleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageProviderId** | Pointer to **int64** | ID of the storage provider (storage bucket) where the bundle should be saved. Defaults to the storage bucket configured as the default support bundle target. If no default is configured, the bundle is saved to the appliance&#39;s local storage. If provided, this ID must reference an existing storage provider or the request returns &#x60;404&#x60;. | [optional] 
**StartDate** | **time.Time** | ISO 8601 start of the log collection window (e.g. &#x60;2026-01-15T00:00:00Z&#x60;). Required. Appliance logs are collected starting from this time. | 
**EndDate** | Pointer to **time.Time** | ISO 8601 end of the log collection window (e.g. &#x60;2026-01-15T23:59:59Z&#x60;). Defaults to the time the request is processed when omitted. | [optional] 
**Contents** | Pointer to [**[]GenerateSupportBundleRequestContentsInner**](GenerateSupportBundleRequestContentsInner.md) | Flat list of support bundle content entries to include. Resource-backed entries should include &#x60;resourceId&#x60;. Standalone entries omit it. If omitted or empty, all eligible content entries are included. | [optional] 

## Methods

### NewGenerateSupportBundleRequest

`func NewGenerateSupportBundleRequest(startDate time.Time, ) *GenerateSupportBundleRequest`

NewGenerateSupportBundleRequest instantiates a new GenerateSupportBundleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSupportBundleRequestWithDefaults

`func NewGenerateSupportBundleRequestWithDefaults() *GenerateSupportBundleRequest`

NewGenerateSupportBundleRequestWithDefaults instantiates a new GenerateSupportBundleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageProviderId

`func (o *GenerateSupportBundleRequest) GetStorageProviderId() int64`

GetStorageProviderId returns the StorageProviderId field if non-nil, zero value otherwise.

### GetStorageProviderIdOk

`func (o *GenerateSupportBundleRequest) GetStorageProviderIdOk() (*int64, bool)`

GetStorageProviderIdOk returns a tuple with the StorageProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderId

`func (o *GenerateSupportBundleRequest) SetStorageProviderId(v int64)`

SetStorageProviderId sets StorageProviderId field to given value.

### HasStorageProviderId

`func (o *GenerateSupportBundleRequest) HasStorageProviderId() bool`

HasStorageProviderId returns a boolean if a field has been set.

### GetStartDate

`func (o *GenerateSupportBundleRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GenerateSupportBundleRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GenerateSupportBundleRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GenerateSupportBundleRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GenerateSupportBundleRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GenerateSupportBundleRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GenerateSupportBundleRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetContents

`func (o *GenerateSupportBundleRequest) GetContents() []GenerateSupportBundleRequestContentsInner`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *GenerateSupportBundleRequest) GetContentsOk() (*[]GenerateSupportBundleRequestContentsInner, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *GenerateSupportBundleRequest) SetContents(v []GenerateSupportBundleRequestContentsInner)`

SetContents sets Contents field to given value.

### HasContents

`func (o *GenerateSupportBundleRequest) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


