# ConvertImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewConvertImageRequest

`func NewConvertImageRequest() *ConvertImageRequest`

NewConvertImageRequest instantiates a new ConvertImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertImageRequestWithDefaults

`func NewConvertImageRequestWithDefaults() *ConvertImageRequest`

NewConvertImageRequestWithDefaults instantiates a new ConvertImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConvertImageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConvertImageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConvertImageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConvertImageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *ConvertImageRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ConvertImageRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ConvertImageRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ConvertImageRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ConvertImageRequest) GetStorageProvider() GetAlerts200ResponseAllOfChecksInnerAccount`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ConvertImageRequest) GetStorageProviderOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ConvertImageRequest) SetStorageProvider(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ConvertImageRequest) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


