# ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult

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

### NewListImageBuilds200ResponseAllOfImageBuildsInnerLastResult

`func NewListImageBuilds200ResponseAllOfImageBuildsInnerLastResult() *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult`

NewListImageBuilds200ResponseAllOfImageBuildsInnerLastResult instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageBuilds200ResponseAllOfImageBuildsInnerLastResultWithDefaults

`func NewListImageBuilds200ResponseAllOfImageBuildsInnerLastResultWithDefaults() *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult`

NewListImageBuilds200ResponseAllOfImageBuildsInnerLastResultWithDefaults instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageBuild

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetImageBuild() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetImageBuildOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetImageBuild(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetBuildNumber

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetStatus

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTempInstance

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetTempInstance() string`

GetTempInstance returns the TempInstance field if non-nil, zero value otherwise.

### GetTempInstanceOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetTempInstanceOk() (*string, bool)`

GetTempInstanceOk returns a tuple with the TempInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempInstance

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetTempInstance(v string)`

SetTempInstance sets TempInstance field to given value.

### HasTempInstance

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasTempInstance() bool`

HasTempInstance returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetVirtualImages() []map[string]interface{}`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) GetVirtualImagesOk() (*[]map[string]interface{}, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) SetVirtualImages(v []map[string]interface{})`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


