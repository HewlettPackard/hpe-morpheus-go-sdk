# ImageBuildExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ImageBuild** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**BuildNumber** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StatusPercent** | Pointer to **int64** |  | [optional] 
**StatusEta** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy**](GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy.md) |  | [optional] 
**TempInstance** | Pointer to **string** |  | [optional] 
**VirtualImages** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewImageBuildExecution

`func NewImageBuildExecution() *ImageBuildExecution`

NewImageBuildExecution instantiates a new ImageBuildExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildExecutionWithDefaults

`func NewImageBuildExecutionWithDefaults() *ImageBuildExecution`

NewImageBuildExecutionWithDefaults instantiates a new ImageBuildExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageBuildExecution) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageBuildExecution) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageBuildExecution) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ImageBuildExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageBuild

`func (o *ImageBuildExecution) GetImageBuild() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *ImageBuildExecution) GetImageBuildOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *ImageBuildExecution) SetImageBuild(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *ImageBuildExecution) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetBuildNumber

`func (o *ImageBuildExecution) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *ImageBuildExecution) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *ImageBuildExecution) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *ImageBuildExecution) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *ImageBuildExecution) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ImageBuildExecution) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ImageBuildExecution) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ImageBuildExecution) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ImageBuildExecution) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ImageBuildExecution) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ImageBuildExecution) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ImageBuildExecution) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ImageBuildExecution) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ImageBuildExecution) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ImageBuildExecution) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ImageBuildExecution) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ImageBuildExecution) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ImageBuildExecution) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ImageBuildExecution) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ImageBuildExecution) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ImageBuildExecution) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ImageBuildExecution) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ImageBuildExecution) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ImageBuildExecution) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetStatus

`func (o *ImageBuildExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageBuildExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageBuildExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImageBuildExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ImageBuildExecution) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ImageBuildExecution) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ImageBuildExecution) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ImageBuildExecution) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ImageBuildExecution) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ImageBuildExecution) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ImageBuildExecution) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ImageBuildExecution) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTempInstance

`func (o *ImageBuildExecution) GetTempInstance() string`

GetTempInstance returns the TempInstance field if non-nil, zero value otherwise.

### GetTempInstanceOk

`func (o *ImageBuildExecution) GetTempInstanceOk() (*string, bool)`

GetTempInstanceOk returns a tuple with the TempInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempInstance

`func (o *ImageBuildExecution) SetTempInstance(v string)`

SetTempInstance sets TempInstance field to given value.

### HasTempInstance

`func (o *ImageBuildExecution) HasTempInstance() bool`

HasTempInstance returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ImageBuildExecution) GetVirtualImages() []map[string]interface{}`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ImageBuildExecution) GetVirtualImagesOk() (*[]map[string]interface{}, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ImageBuildExecution) SetVirtualImages(v []map[string]interface{})`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ImageBuildExecution) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


