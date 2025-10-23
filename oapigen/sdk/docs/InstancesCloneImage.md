# InstancesCloneImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateName** | Pointer to **string** | Image Template Name | [optional] [default to "{server.name}-{timestamp}"]
**ZoneFolder** | Pointer to **string** | Zone Folder externalId. This is required for VMware | [optional] 

## Methods

### NewInstancesCloneImage

`func NewInstancesCloneImage() *InstancesCloneImage`

NewInstancesCloneImage instantiates a new InstancesCloneImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesCloneImageWithDefaults

`func NewInstancesCloneImageWithDefaults() *InstancesCloneImage`

NewInstancesCloneImageWithDefaults instantiates a new InstancesCloneImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateName

`func (o *InstancesCloneImage) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *InstancesCloneImage) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *InstancesCloneImage) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *InstancesCloneImage) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetZoneFolder

`func (o *InstancesCloneImage) GetZoneFolder() string`

GetZoneFolder returns the ZoneFolder field if non-nil, zero value otherwise.

### GetZoneFolderOk

`func (o *InstancesCloneImage) GetZoneFolderOk() (*string, bool)`

GetZoneFolderOk returns a tuple with the ZoneFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFolder

`func (o *InstancesCloneImage) SetZoneFolder(v string)`

SetZoneFolder sets ZoneFolder field to given value.

### HasZoneFolder

`func (o *InstancesCloneImage) HasZoneFolder() bool`

HasZoneFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


