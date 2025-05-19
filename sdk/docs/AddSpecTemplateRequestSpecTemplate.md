# AddSpecTemplateRequestSpecTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Spec template name | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | [**AddSpecTemplateRequestSpecTemplateType**](AddSpecTemplateRequestSpecTemplateType.md) |  | 
**File** | [**AddSpecTemplateRequestSpecTemplateFile**](AddSpecTemplateRequestSpecTemplateFile.md) |  | 
**Config** | Pointer to [**AddSpecTemplateRequestSpecTemplateConfig**](AddSpecTemplateRequestSpecTemplateConfig.md) |  | [optional] 

## Methods

### NewAddSpecTemplateRequestSpecTemplate

`func NewAddSpecTemplateRequestSpecTemplate(name string, type_ AddSpecTemplateRequestSpecTemplateType, file AddSpecTemplateRequestSpecTemplateFile, ) *AddSpecTemplateRequestSpecTemplate`

NewAddSpecTemplateRequestSpecTemplate instantiates a new AddSpecTemplateRequestSpecTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSpecTemplateRequestSpecTemplateWithDefaults

`func NewAddSpecTemplateRequestSpecTemplateWithDefaults() *AddSpecTemplateRequestSpecTemplate`

NewAddSpecTemplateRequestSpecTemplateWithDefaults instantiates a new AddSpecTemplateRequestSpecTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddSpecTemplateRequestSpecTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSpecTemplateRequestSpecTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSpecTemplateRequestSpecTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *AddSpecTemplateRequestSpecTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddSpecTemplateRequestSpecTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddSpecTemplateRequestSpecTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddSpecTemplateRequestSpecTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AddSpecTemplateRequestSpecTemplate) GetType() AddSpecTemplateRequestSpecTemplateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddSpecTemplateRequestSpecTemplate) GetTypeOk() (*AddSpecTemplateRequestSpecTemplateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddSpecTemplateRequestSpecTemplate) SetType(v AddSpecTemplateRequestSpecTemplateType)`

SetType sets Type field to given value.


### GetFile

`func (o *AddSpecTemplateRequestSpecTemplate) GetFile() AddSpecTemplateRequestSpecTemplateFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AddSpecTemplateRequestSpecTemplate) GetFileOk() (*AddSpecTemplateRequestSpecTemplateFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AddSpecTemplateRequestSpecTemplate) SetFile(v AddSpecTemplateRequestSpecTemplateFile)`

SetFile sets File field to given value.


### GetConfig

`func (o *AddSpecTemplateRequestSpecTemplate) GetConfig() AddSpecTemplateRequestSpecTemplateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddSpecTemplateRequestSpecTemplate) GetConfigOk() (*AddSpecTemplateRequestSpecTemplateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddSpecTemplateRequestSpecTemplate) SetConfig(v AddSpecTemplateRequestSpecTemplateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddSpecTemplateRequestSpecTemplate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


