# AddSecurityPackages200ResponseAllOfSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AddSecurityPackages200ResponseAllOfSecurityPackageType**](AddSecurityPackages200ResponseAllOfSecurityPackageType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddSecurityPackages200ResponseAllOfSecurityPackage

`func NewAddSecurityPackages200ResponseAllOfSecurityPackage() *AddSecurityPackages200ResponseAllOfSecurityPackage`

NewAddSecurityPackages200ResponseAllOfSecurityPackage instantiates a new AddSecurityPackages200ResponseAllOfSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityPackages200ResponseAllOfSecurityPackageWithDefaults

`func NewAddSecurityPackages200ResponseAllOfSecurityPackageWithDefaults() *AddSecurityPackages200ResponseAllOfSecurityPackage`

NewAddSecurityPackages200ResponseAllOfSecurityPackageWithDefaults instantiates a new AddSecurityPackages200ResponseAllOfSecurityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetType() AddSecurityPackages200ResponseAllOfSecurityPackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetTypeOk() (*AddSecurityPackages200ResponseAllOfSecurityPackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetType(v AddSecurityPackages200ResponseAllOfSecurityPackageType)`

SetType sets Type field to given value.

### HasType

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddSecurityPackages200ResponseAllOfSecurityPackage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


