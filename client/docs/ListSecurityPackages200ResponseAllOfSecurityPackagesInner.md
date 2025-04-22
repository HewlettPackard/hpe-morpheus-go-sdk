# ListSecurityPackages200ResponseAllOfSecurityPackagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListSecurityPackages200ResponseAllOfSecurityPackagesInner

`func NewListSecurityPackages200ResponseAllOfSecurityPackagesInner() *ListSecurityPackages200ResponseAllOfSecurityPackagesInner`

NewListSecurityPackages200ResponseAllOfSecurityPackagesInner instantiates a new ListSecurityPackages200ResponseAllOfSecurityPackagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityPackages200ResponseAllOfSecurityPackagesInnerWithDefaults

`func NewListSecurityPackages200ResponseAllOfSecurityPackagesInnerWithDefaults() *ListSecurityPackages200ResponseAllOfSecurityPackagesInner`

NewListSecurityPackages200ResponseAllOfSecurityPackagesInnerWithDefaults instantiates a new ListSecurityPackages200ResponseAllOfSecurityPackagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


