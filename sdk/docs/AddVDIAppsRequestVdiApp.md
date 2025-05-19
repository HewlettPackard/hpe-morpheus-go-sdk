# AddVDIAppsRequestVdiApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | VDI App name | 
**Description** | Pointer to **string** | Description | [optional] 
**IconPath** | Pointer to ***os.File** | Icon Path. A relative location of an icon image | [optional] 
**LaunchPrefix** | Pointer to **string** | The RDS App Name Prefix.  Must start with &#39;||&#39; | [optional] 

## Methods

### NewAddVDIAppsRequestVdiApp

`func NewAddVDIAppsRequestVdiApp(name string, ) *AddVDIAppsRequestVdiApp`

NewAddVDIAppsRequestVdiApp instantiates a new AddVDIAppsRequestVdiApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVDIAppsRequestVdiAppWithDefaults

`func NewAddVDIAppsRequestVdiAppWithDefaults() *AddVDIAppsRequestVdiApp`

NewAddVDIAppsRequestVdiAppWithDefaults instantiates a new AddVDIAppsRequestVdiApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddVDIAppsRequestVdiApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVDIAppsRequestVdiApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVDIAppsRequestVdiApp) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddVDIAppsRequestVdiApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVDIAppsRequestVdiApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVDIAppsRequestVdiApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVDIAppsRequestVdiApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconPath

`func (o *AddVDIAppsRequestVdiApp) GetIconPath() *os.File`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *AddVDIAppsRequestVdiApp) GetIconPathOk() (**os.File, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *AddVDIAppsRequestVdiApp) SetIconPath(v *os.File)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *AddVDIAppsRequestVdiApp) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLaunchPrefix

`func (o *AddVDIAppsRequestVdiApp) GetLaunchPrefix() string`

GetLaunchPrefix returns the LaunchPrefix field if non-nil, zero value otherwise.

### GetLaunchPrefixOk

`func (o *AddVDIAppsRequestVdiApp) GetLaunchPrefixOk() (*string, bool)`

GetLaunchPrefixOk returns a tuple with the LaunchPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchPrefix

`func (o *AddVDIAppsRequestVdiApp) SetLaunchPrefix(v string)`

SetLaunchPrefix sets LaunchPrefix field to given value.

### HasLaunchPrefix

`func (o *AddVDIAppsRequestVdiApp) HasLaunchPrefix() bool`

HasLaunchPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


