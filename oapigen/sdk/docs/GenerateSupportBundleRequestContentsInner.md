# GenerateSupportBundleRequestContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Support bundle content type code. Preferred identifier for this API. Use &#x60;GET /api/options/supportBundles.contentTypes&#x60; to see available values. If provided, the code must resolve to a valid support bundle content type or the request returns &#x60;400&#x60;. | [optional] 
**TypeId** | Pointer to **int64** | Alternative support bundle content type ID returned by &#x60;GET /api/options/supportBundles.contentTypes&#x60;. Use this instead of &#x60;code&#x60; when the numeric type ID is known. If provided, the type ID must resolve to a valid support bundle content type or the request returns &#x60;400&#x60;. | [optional] 
**ResourceId** | Pointer to **int64** | Resource ID for resource-backed content types. Use &#x60;GET /api/options/supportBundles.contentTypeResources?contentTypeCode&#x3D;...&#x60; to see eligible values for a selected content type. Omit or set to null for standalone content types. | [optional] 

## Methods

### NewGenerateSupportBundleRequestContentsInner

`func NewGenerateSupportBundleRequestContentsInner() *GenerateSupportBundleRequestContentsInner`

NewGenerateSupportBundleRequestContentsInner instantiates a new GenerateSupportBundleRequestContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetCode

`func (o *GenerateSupportBundleRequestContentsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GenerateSupportBundleRequestContentsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GenerateSupportBundleRequestContentsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GenerateSupportBundleRequestContentsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTypeId

`func (o *GenerateSupportBundleRequestContentsInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *GenerateSupportBundleRequestContentsInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *GenerateSupportBundleRequestContentsInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *GenerateSupportBundleRequestContentsInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetResourceId

`func (o *GenerateSupportBundleRequestContentsInner) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *GenerateSupportBundleRequestContentsInner) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *GenerateSupportBundleRequestContentsInner) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *GenerateSupportBundleRequestContentsInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


