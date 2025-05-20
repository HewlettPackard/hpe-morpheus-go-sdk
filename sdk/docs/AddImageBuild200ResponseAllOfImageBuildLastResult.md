# AddImageBuild200ResponseAllOfImageBuildLastResult

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
**VirtualImages** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewAddImageBuild200ResponseAllOfImageBuildLastResult

`func NewAddImageBuild200ResponseAllOfImageBuildLastResult() *AddImageBuild200ResponseAllOfImageBuildLastResult`

NewAddImageBuild200ResponseAllOfImageBuildLastResult instantiates a new AddImageBuild200ResponseAllOfImageBuildLastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddImageBuild200ResponseAllOfImageBuildLastResultWithDefaults

`func NewAddImageBuild200ResponseAllOfImageBuildLastResultWithDefaults() *AddImageBuild200ResponseAllOfImageBuildLastResult`

NewAddImageBuild200ResponseAllOfImageBuildLastResultWithDefaults instantiates a new AddImageBuild200ResponseAllOfImageBuildLastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageBuild

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetImageBuild() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetImageBuildOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetImageBuild(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetBuildNumber

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusPercent

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetStatus

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTempInstance

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetTempInstance() string`

GetTempInstance returns the TempInstance field if non-nil, zero value otherwise.

### GetTempInstanceOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetTempInstanceOk() (*string, bool)`

GetTempInstanceOk returns a tuple with the TempInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempInstance

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetTempInstance(v string)`

SetTempInstance sets TempInstance field to given value.

### HasTempInstance

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasTempInstance() bool`

HasTempInstance returns a boolean if a field has been set.

### GetVirtualImages

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetVirtualImages() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) GetVirtualImagesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) SetVirtualImages(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *AddImageBuild200ResponseAllOfImageBuildLastResult) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


