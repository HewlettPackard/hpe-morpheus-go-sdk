# SpecTemplateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Spec template name | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to [**UpdateSpecTemplateRequestSpecTemplateType**](UpdateSpecTemplateRequestSpecTemplateType.md) |  | [optional] 
**File** | Pointer to [**UpdateSpecTemplateRequestSpecTemplateFile**](UpdateSpecTemplateRequestSpecTemplateFile.md) |  | [optional] 
**Config** | Pointer to [**UpdateSpecTemplateRequestSpecTemplateConfig**](UpdateSpecTemplateRequestSpecTemplateConfig.md) |  | [optional] 

## Methods

### NewSpecTemplateUpdate

`func NewSpecTemplateUpdate() *SpecTemplateUpdate`

NewSpecTemplateUpdate instantiates a new SpecTemplateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecTemplateUpdateWithDefaults

`func NewSpecTemplateUpdateWithDefaults() *SpecTemplateUpdate`

NewSpecTemplateUpdateWithDefaults instantiates a new SpecTemplateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpecTemplateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpecTemplateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpecTemplateUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpecTemplateUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *SpecTemplateUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SpecTemplateUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SpecTemplateUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SpecTemplateUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *SpecTemplateUpdate) GetType() UpdateSpecTemplateRequestSpecTemplateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpecTemplateUpdate) GetTypeOk() (*UpdateSpecTemplateRequestSpecTemplateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpecTemplateUpdate) SetType(v UpdateSpecTemplateRequestSpecTemplateType)`

SetType sets Type field to given value.

### HasType

`func (o *SpecTemplateUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFile

`func (o *SpecTemplateUpdate) GetFile() UpdateSpecTemplateRequestSpecTemplateFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SpecTemplateUpdate) GetFileOk() (*UpdateSpecTemplateRequestSpecTemplateFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SpecTemplateUpdate) SetFile(v UpdateSpecTemplateRequestSpecTemplateFile)`

SetFile sets File field to given value.

### HasFile

`func (o *SpecTemplateUpdate) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *SpecTemplateUpdate) GetConfig() UpdateSpecTemplateRequestSpecTemplateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SpecTemplateUpdate) GetConfigOk() (*UpdateSpecTemplateRequestSpecTemplateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SpecTemplateUpdate) SetConfig(v UpdateSpecTemplateRequestSpecTemplateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SpecTemplateUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


